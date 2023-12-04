<template>
  <div class="ln-dt-vw">
      <h2> <b>Loan Details View</b> </h2>
      <hr>
      <div id="btn-cntnr">
          <Button class="clickable-btn" @click="withdrawNFT" :disabled="data.isActiveLoan || Number(data.auctionState != 0) || withdrawNFTButtonDisabled" v-if="data != null"> Withdraw NFT </Button>
          <Button class="clickable-btn" @click="claimBalance" :disabled="data.isActiveLoan || data.hasClaimed" v-if="data != null"> Claim Balance </Button>
          <Button class="clickable-btn" @click="refreshLoan" :disabled="refreshLoanBtnDisabled"> Refresh Loan </Button>
      </div>
      <hr>
      <p v-if="data != null"> <b>NFT Pledger's Address</b> : {{data.pledger}} </p>
       <p v-if="data != null"> <b>NFT Contract's Address</b> : {{data.nftContractAddress}} </p>
        <p v-if="data != null"> <b>Current Owner (last-known)</b> : {{data.lastKnownCurrentOwner}} </p>


      <div id="ln-dtls-vw">
                <div id="img-div">
                    <img :src="imagingURL"/>
                </div>
          <table v-if="data != null"> 

              <thead>
                  <th>
                  <i><b>Loan Data</b></i> {{getEmojiActive(data.isActiveLoan)}}
                  </th>
              </thead>
              <tbody>
                  <tr>
                    <td> <b> Loan Id </b> </td>
                    <td> {{data.loanId}} </td>
                    <td> <b>Token Id </b></td>
                    <td> {{data.tokenId}} </td>
                    <td> <b>Loan Amount</b> </td>
                    <td> {{getInXRP(String(data.loanAmount))}} XRP </td>
                    <td> <b>Interest Rate</b> </td>
                    <td> {{String(data.interestRate)}}% </td>
                    <td> <b>Tenure</b> </td>
                    <td> {{String(data.tenure)}} years </td>
                  </tr>
                  <tr>
                    <td> <b> Repayment Amount </b> </td>
                    <td> {{getInXRP(String(data.repaymentAmount))}} XRP </td>
                    <td> <b>Interest</b> </td>
                    <td> {{getInXRP(String(data.interest))}} XRP </td>
                    <td> <b>Amount Repaid </b></td>
                    <td> {{getInXRP(String(data.amountRepaid))}} XRP </td>
                    <td> <b>Pledged Bank</b> </td>
                    <td>  {{data.bankToWhichItisPledged}} </td>
                    <td> <b>Refunded Amount</b> </td>
                    <td> {{getInXRP(String(data.amountReturnedToUser))}} XRP </td>
                  </tr>
                  <tr>
                    <td> <b> Auction State </b> </td>
                    <td> {{getAuctionState(data.auctionState)}} </td>
                    <td> <b>Auction Id</b> </td>
                    <td> {{data.auctionId}} </td>
                    <td> <b>Auction profit</b> </td>
                    <td> {{getInXRP(String(data.amountEarnedFromAuction))}} XRP </td>
                    <td> <b> Ready for auction </b></td>
                    <td>  {{data.canBeAuctioned}} </td>
                    <td> <b> NFT symbol </b></td>
                    <td> {{data.nftSymbol}} </td>
                  </tr>
               </tbody>
          </table>
      </div>
      <br>
      <hr>
      <h3> Repayment Schedule </h3>
      <div id="rps-tbl">
          <table v-if="this.rpsLoaded">
              <thead>
                  <th> <i> RPS Index </i></th>
                  <th> <i>Amount</i> </th>
                  <th> <i>Due Date</i> </th>
                  <th> <i> Is Due </i> </th>
                  <th> <i> LPI </i> </th>
                  <th> <i>Amount Repaid</i> </th>
                  <th> <i>Action</i> </th>
              </thead>
              <tbody>
                  <tr v-for="(item,index) in rps" :key="item.rpsIndex"> 
                      <td> {{item.rpsIndex}} </td>
                      <td> {{getInXRP(item.amount)}} XRP </td>
                      <td> {{getDate(item.dueDate)}} </td>
                      <td> {{item.isDue}} </td>
                      <td> {{item.lpiCharge}} XRP </td>

                      <td> {{getInXRP(item.amountRepaid)}} XRP </td>
                      <td> <Button class="tbl-py" :disabled="(item.amountRepaid >= item.amount) || repayBtnDisable[index] || data.auctionState != 0" @click="repayEMI(index)"> Pay </Button> </td>
                  </tr>
              </tbody>
          </table>
      </div>

     <alert-modal :is-visible="isModalVisible" :message="alertMessage" @close="onCloseModal" />


    
  </div>
</template>

<script>
// @ is an alias to /src
import blockBorrowerAbi from '../contract_abis/blockborrower_abi.json'
import erc721Abi from '../contract_abis/erc721_abi.json'
import web3 from 'web3';
import {EVM_SIDECHAIN_RPC, BLOCKBORROWER_CONTRACT_ADDRESS, EVM_SIDECHAIN_BLOCK_EXPLORER} from '../env/env.js';
import AlertModal from '../components/AlertModal.vue';
import {createAlchemyWeb3} from "@alch/alchemy-web3"

export default {
  name: 'LoanDetailsView',

  components: {
      AlertModal,
  },

  data() {
      return {
          data: null,
          rps: null,
          evmXrplChainId: "0x15f902",
          isModalVisible: false,
          alertMessage: "",
          repayBtnDisable: [],
          refreshLoanBtnDisabled: false,
          rpsLoaded: false,
          imagingURL: "https://developers.elementor.com/docs/assets/img/elementor-placeholder-image.png",
          withdrawNFTButtonDisabled: false


      }
  },


  mounted() {
    // eslint-disable-next-line
    const bankId = Number(this.$route.params.bankId)
    // eslint-disable-next-line
    const customerAddress = String(this.$route.params.customerAddress);
    const loanId = Number(this.$route.params.loanId)

    console.log("Bank ID and Customer Address: ", bankId, customerAddress)

    this.getNFTByLoanId(bankId, customerAddress, loanId);

  },

  methods: {

      getEmojiActive(x) {
          if(x) {
              return "✅"
          } else {
              return "❌"
          }
      },

    getDate(sec) {
        let date = new Date(0);
        date.setUTCSeconds(sec)
         const options = {
            year: 'numeric',
            month: 'long',
            day: 'numeric',
            hour: 'numeric',
            minute: 'numeric',
            second: 'numeric',
            timeZoneName: 'short',
        };

        return date.toLocaleString('en-US', options);
    },

    getAuctionState(x) {

            if(x == 0) {
                return "Not Started"
            } else if (x == 1) {
                return "Running"
            } else if (x == 2) {
                return "Finalized"
            }
  
    },

    getInXRP(amt) {
        console.log(amt)
        return web3.utils.fromWei(amt, "ether");
    }, 

    async getNFTByLoanId(bankId, customerAddress, loanId) {
        let result = await this.getNFTHoldingOfCustomerForBankId(bankId, customerAddress);
        for(let i = 0; i < result.length; i++) {
            if(Number(result[i].loanId) == loanId) {
                    this.data = result[i]
                    this.getImage(result[i].nftContractAddress, result[i].tokenId);
                    break
            }
        }

        await this.getRepaymentSchedules(this.data.loanId);
        await this.appendDueAndLPI();
        
        this.rpsLoaded = true;
    },

    async getNFTHoldingOfCustomerForBankId(bankId, customerAddress) {
            //console.log(bankId, customerAddress);
            const web3i = new web3(EVM_SIDECHAIN_RPC);
            try {
                const contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                let result = await contract.methods.getAllNftDepositsOfBank(bankId, customerAddress).call({from:customerAddress})
                const jsonString = JSON.stringify(result, (key, value) => (typeof value === 'bigint' ? value.toString() : value));
                //console.log(jsonString)
                return JSON.parse(jsonString);
            } catch (error) {
                console.log(error);
                return [];
            }
    },

    async getRepaymentSchedules(loanId) {
        this.repayBtnDisable = []
        const web3i = new web3(EVM_SIDECHAIN_RPC);
        try {
            //console.log("Inside RPS")
            const contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
            let result = await contract.methods.getRepaymentSchedules(loanId).call({})
            const jsonString = JSON.stringify(result, (key, value) => (typeof value === 'bigint' ? value.toString() : value));
            //console.log("RPS executed")
            console.log(JSON.parse(jsonString))
            this.rps = JSON.parse(jsonString)
            for(let i = 0; i < this.rps.length; i++) {
                this.repayBtnDisable[i] = this.rps[i].amountRepaid >= this.amount;
            }
        } catch(error) {
            console.log(error)
        }
    },

    async appendDueAndLPI() {
        const web3i = new web3(EVM_SIDECHAIN_RPC);
        const latestBlockNumber = await web3i.eth.getBlockNumber();
        const latestBlock = await web3i.eth.getBlock(latestBlockNumber);

        const ts = Number(latestBlock.timestamp)
        for(let i = 0; i < this.rps.length; i++) {
            if(ts > this.rps[i].dueDate) {
                // due
                this.rps[i].isDue = true
                this.rps[i].lpiCharge = web3.utils.toWei("1", "ether");
            } else {
                // no due
                this.rps[i].isDue = false
                this.rps[i].lpiCharge = "0"
            }
        }

    },

    async getCurrentChainId() {
        if (window.ethereum) {
            const ethereum = window.ethereum;
            try {
                const chainId = await ethereum.request({ method: 'eth_chainId' });
                return chainId
            } catch (error) {
                console.error(`Error getting chainId: ${error}`);
                this.showAlert(`Error getting chainId: ${error}`)
                return "0"
            }
        } else {
            console.error('MetaMask is not available. Please install and configure MetaMask.');
            this.showAlert('MetaMask is not available. Please install and configure MetaMask.')
            return "0"
        }
    },

        showAlert(alertMsg) {
            this.isModalVisible = true;
            this.alertMessage = alertMsg;
        },
        
        onCloseModal() {
            this.isModalVisible = false;
            this.alertMessage = "";
        },

    

    async repayEMI(idx) {
        this.repayBtnDisable[idx] = true;
        const chainId = await this.getCurrentChainId();
        if(chainId != this.evmXrplChainId) {
                this.showAlert("Switch to EVM XRPL Sidechain");
                console.log(chainId);
                return;
        }

        try {
            const sender = localStorage.getItem("walletAddress")
            const bankId = this.data.bankToWhichItisPledged;
            const loanId = this.data.loanId;
            const repayAmt = Number(this.rps[idx].amount) + Number(this.rps[idx].lpiCharge);
            //console.log("Repay Amy", repayAmt);
            if(this.rps[idx].amountRepaid > repayAmt) {
                this.showAlert("Please do not tamper");
                return;
            }

            const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL



            window.contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
            // for gas charge
            const one_ether = Number(web3.utils.toWei("1", "Gwei"))
            var weiValue = (Number(repayAmt) + one_ether).toString(16);
            //console.log("Wei Value", weiValue);
            //console.log(this.rps, idx)
            let txnParams = {
                to: BLOCKBORROWER_CONTRACT_ADDRESS,
                from: sender,
                value: weiValue,
                data: window.contract.methods.repayEmi(bankId, loanId, idx).encodeABI(),
            }

             const txnHash = await window.ethereum.request({
                            method: 'eth_sendTransaction',
                            params: [txnParams],
                    })

            this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
            console.log('Transaction hash:', txnHash)
            setTimeout(() => {
                this.getTxnReceiptAction(txnHash, "Repayment done successfully, refresh the page to check updated loan details !!")
            }, 5000);

        } catch(error) {
            this.showAlert(JSON.stringify(error));
            this.repayBtnDisable[idx] = false;
            console.log(error);
        }

    },

        async getTxnReceiptAction(txnHash, message) {
            try {
                const web3i = new web3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL

                const receipt = await web3i.eth.getTransactionReceipt(txnHash);
                if (receipt && receipt.status) {
                    this.showAlert(message)
                    console.log(receipt)
                } else {
                    this.showAlert("Transaction Failed. check it in blockexplorer::" + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash);
                    this.withdrawNFTButtonDisabled = false;
                }
            } catch(err) {
                console.log(err)
            }
        },

        async checkAuctionable(bankId, loanId) {

            const chainId = await this.getCurrentChainId();
            if(chainId != this.evmXrplChainId) {
                    this.showAlert("Switch to EVM XRPL Sidechain");
                    console.log(chainId);
                    return;
            }

            const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL

            try {
                window.contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                const sender = localStorage.getItem('walletAddress');
                let txnParams = {
                    to: BLOCKBORROWER_CONTRACT_ADDRESS,
                    from: sender,
                    data: window.contract.methods.checkIfAuctionable(bankId, loanId).encodeABI(),
                }

                const txnHash = await window.ethereum.request({
                        method: 'eth_sendTransaction',
                        params: [txnParams],
                })

                this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
                console.log('Transaction hash:', txnHash)
                setTimeout(() => {
                    this.getTxnReceiptAction(txnHash, "Check auction done !!")
                }, 5000);

            } catch(error) {
                console.log(error)
            }
        },

        async checkIfLoanIsClosed(bankId, loanId, userAddress) {

            const sender = localStorage.getItem('walletAddress')

            const chainId = await this.getCurrentChainId();
            if(chainId != this.evmXrplChainId) {
                    this.showAlert("Switch to EVM XRPL Sidechain");
                    console.log(chainId);
                    return;
            }

            const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL

            try {
                window.contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                let txnParams = {
                    to: BLOCKBORROWER_CONTRACT_ADDRESS,
                    from: sender,
                    data: window.contract.methods.isLoanClosed(bankId, loanId, userAddress).encodeABI(),
                }

                const txnHash = await window.ethereum.request({
                        method: 'eth_sendTransaction',
                        params: [txnParams],
                })

                this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
                console.log('Transaction hash:', txnHash)
                setTimeout(() => {
                    this.getTxnReceiptAction(txnHash, "Check Loan close done !!")
                }, 5000);

            } catch(error) {
                console.log(error)
            }

        },

        async refreshLoan() {
            this.refreshLoanBtnDisabled = true;
            await this.checkAuctionable(this.data.bankToWhichItisPledged, this.data.loanId);
            await this.checkIfLoanIsClosed(this.data.bankToWhichItisPledged, this.data.loanId, this.data.pledger);
            this.refreshLoanBtnDisabled = false;
        },

        async claimBalance() {
            const loanId = this.data.loanId
            const bankId = this.data.bankToWhichItisPledged   
            const sender = localStorage.getItem('walletAddress')

            const chainId = await this.getCurrentChainId();
            if(chainId != this.evmXrplChainId) {
                    this.showAlert("Switch to EVM XRPL Sidechain");
                    console.log(chainId);
                    return;
            }

            const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL

            try {
                window.contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                let txnParams = {
                    to: BLOCKBORROWER_CONTRACT_ADDRESS,
                    from: sender,
                    data: window.contract.methods.claimBalanceFromLoan(bankId, loanId).encodeABI(),
                }

                const txnHash = await window.ethereum.request({
                        method: 'eth_sendTransaction',
                        params: [txnParams],
                })

                this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
                console.log('Transaction hash:', txnHash)
                setTimeout(() => {
                    this.getTxnReceiptAction(txnHash, "Balance claimed, refresh the page and check your balance in metamask !!")
                }, 5000);

            } catch(error) {
                console.log(error)
            }

        },

        async withdrawNFT() {
            console.log("Withdrawing NFT")
            this.withdrawNFTButtonDisabled = true;
            const loanId = this.data.loanId
            const bankId = this.data.bankToWhichItisPledged   
            const sender = localStorage.getItem('walletAddress')
            const nftContractAddress = this.data.nftContractAddress;
            const chainId = await this.getCurrentChainId();

            if(chainId != this.evmXrplChainId) {
                    this.withdrawNFTButtonDisabled = false
                    this.showAlert("Switch to EVM XRPL Sidechain");
                    console.log(chainId);
                    return;
            }

            const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL
            try {
                window.contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                let txnParams = {
                    to: BLOCKBORROWER_CONTRACT_ADDRESS,
                    from: sender,
                    data: window.contract.methods.withdrawNFT(bankId, nftContractAddress,loanId).encodeABI(),
                }

                const txnHash = await window.ethereum.request({
                        method: 'eth_sendTransaction',
                        params: [txnParams],
                })

                this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
                console.log('Transaction hash:', txnHash)
                setTimeout(() => {
                    this.getTxnReceiptAction(txnHash, "NFT Withdrawn, please proceed ahead to do a cross-chain transfer if required.!!")
                }, 5000);

            } catch(error) {
                this.withdrawNFTButtonDisabled = false;
                console.log(error)
            }


        },

        async getTokenURI(contractAddress, tokenId) {
           const web3i = new web3(EVM_SIDECHAIN_RPC);
           try {
                const contract = new web3i.eth.Contract(erc721Abi, contractAddress);
                // eslint-disable-next-line
                const tokenURI = await contract.methods.tokenURI(BigInt(tokenId)).call({})

                return tokenURI;
           } catch (error) {
               console.log(error);
               return null;
           }
      },

    async decodeTheTokenURIForImage(imageURL) {
            const x = {
                "image": imageURL
            }

         try {
            const response = await fetch(imageURL);

            if (!response.ok) {
                return x
            }

            const data = await response.json();
            console.log(data);
            return data;
        } catch (error) {
            //console.error('Error fetching JSON:', error);
            return x
        }
    },

    async getImage(contractAddress, tokenId) {
        const res = await this.getTokenURI(contractAddress, tokenId);
        const iurl = await this.decodeTheTokenURIForImage(res);
        this.imagingURL = iurl.image
    }
  }
}
</script>

<style scoped>
#btn-cntnr {
    text-align: right;
}
.clickable-btn {
    background-image: linear-gradient(to right, aliceblue, skyblue);
    margin-right: 20px;
    border-radius: 5px;
    padding: 6px;
    font-family: monospace;
    cursor: pointer;
}

.clickable-btn:hover:enabled {
    background-image: linear-gradient(to left, aliceblue, skyblue);
}

.clickable-btn:disabled {
    background-image: linear-gradient(to right, lightgray, gray);

}

p {
    text-align: left;
    background-color: lightcoral;
    padding: 5px;
}

 table {
    width: 100%;
    border-collapse: collapse; /* Ensure borders collapse into a single border */
  }

  th, td {
    border: 1px solid black; /* Set border for table cells */
    padding: 8px; /* Add padding for better readability */
    text-align: left; /* Align text to the left */
  }

  td {
    background-color: rgb(136, 242, 136); /* Set background color for table header */
  }

  td:hover {
    background-color:aliceblue; /* Set background color for table header */

  }

  th {
      background-color: lightpink;
  }

  .tbl-py {
      width: 100%;
      padding: 5px;
      border-radius: 5px;
      background-color: lightsalmon;
      font-family: monospace;
      cursor: pointer;
  }

  #img-div {
      text-align: center;
  }

  img {
      width: 150px;
  }

</style>