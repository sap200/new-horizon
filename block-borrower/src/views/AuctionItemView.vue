<template>
  <div class="au-itm">
      <h3> Auction Item Detail </h3>
      <div id="btn-cntnr" v-if="data != null">
          <Button class="au-btns" @click="withdrawAmount"> Withdraw Amount </Button>

          <Button class="au-btns" v-if="data != null" @click="checkIfAuctionHasEnded" :disabled="checkAuctionStateBtnDisabled || this.data.auctionState != 1"> Click If auction Ended </Button>
       </div>
       <hr>
        <div id="img-div" v-if="data != null">
            <img :src="imagingURL" />
        </div>
        <div id = "timer" style="color:red" v-if="data != null && data.auctionState == 1">
            Auction ends in - {{this.formattedTime}} minutes
        </div>
    <div id="place-a-bid" v-if="data != null && data.auctionState == 1">
        <label> <b>Bid Amount</b> </label> <input type="number" v-model="bidValue"/> XRP 
        <Button class="au-btns" id="bid-btn" @click="placeABid" :disabled="disableBidBtn"> Place a Bid </Button> 
    </div>
    <hr>
    <div id="auc-end-tm" style="color:blue;" v-if="data != null">
        <p> <b> Auction State </b> : {{getAuctionState(Number(data.auctionState))}} </p>
        <p> <b> Auction Start Time </b>: {{getDate(data.auctionStartTime)}}</p>
        <p> <b> Auction End Time </b> : {{getDate(Number(data.auctionStartTime) + 20*60)}} </p>
        <p> <b> Higesht Bid </b> : {{getInXRP(data.highestBindingBid)}} XRP </p>
        <p> <b> Highest Bidder </b> : {{data.highestBidder}} </p>
    </div>
    <hr>
    <h4> Details </h4>
    <div id="tbl-vw" v-if="data != null">
        <table>
            <tr>
                <td> <b>Auction Id</b> </td>
                <td> {{data.auctionId}} </td>
                <td> <b>Loan Id</b> </td>
                <td> {{data.loanId}} </td>
            </tr>

            <tr>
                <td> <b> Bank Id </b> </td>
                <td> {{data.bankToWhichItisPledged}} </td>
                <td> <b> Bank Address </b> </td>
                <td> {{data.bankAddress}} </td>
            </tr>
            <tr>
                <td> <b>Bank Admin</b> </td>
                <td> {{data.bankAdmin}} </td>
                <td><b> Contract Address </b> </td>
                <td> {{data.nftContractAddress}} </td>
            </tr>

            <tr>
                <td> <b> Symbol </b> </td>
                <td> {{data.symbol}} </td>
                <td> <b> Token ID </b> </td>
                <td> {{data.tokenId}} </td>
            </tr>
        
        </table>
    </div>
    <hr>
    <h4> Bidder's List </h4>
    <div id="tbl-vw-x" v-if="biddersList != null">
        <table>
            <thead>
                <th> Bidder's Address </th>
                <th> Bid Amount </th>
            </thead>
            <tbody>
                <tr v-for="(item,index) in biddersList" :key="index"> 
                     <td> {{item.bidder}} </td>
                    <td> {{getInXRP(item.bid)}} XRP </td>
                </tr>
            </tbody>

        </table>

    </div>

    <alert-modal :is-visible="isModalVisible" :message="alertMessage" @close="onCloseModal" />



  </div>
</template>

<script>
// @ is an alias to /src

// import blockBorrowerAbi from '../contract_abis/blockborrower_abi.json'
import auctionAbi from '../contract_abis/auction_abi.json'
import erc721Abi from '../contract_abis/erc721_abi.json'
import web3 from 'web3';
import {EVM_SIDECHAIN_RPC, AUCTION_CONTRACT_ADDRESS, BLOCKBORROWER_CONTRACT_ADDRESS, EVM_SIDECHAIN_BLOCK_EXPLORER} from '../env/env.js';
import {createAlchemyWeb3} from "@alch/alchemy-web3"
import blockBorrowerAbi from '../contract_abis/blockborrower_abi.json'
import AlertModal from '../components/AlertModal.vue';


export default {
  name: 'AuctionItemView',

  components: {
      AlertModal,
  },

  data() {
      return {
          data: null,
          imagingURL:  "https://developers.elementor.com/docs/assets/img/elementor-placeholder-image.png",
          currentTime: 0,
          endTime: 0, 
          intervalId: null,
          biddersList: null,
          evmxrplChainId: "0x15f902",
          checkAuctionStateBtnDisabled: false,
          isModalVisible: false,
          alertMessage: "",
          bidValue: 1,
          disableBidBtn: false,


      }
  },

  created() {
      this.performInitialization()
  },

  destroy() {
      clearInterval(this.intervalId);
  },

  computed: {
        formattedTime() {
            const hours = Math.floor(this.currentTime / 3600);
            const minutes = Math.floor((this.currentTime % 3600) / 60);
            const seconds = this.currentTime % 60;

            return `${String(hours).padStart(2, '0')}:${String(minutes).padStart(2, '0')}:${String(seconds).padStart(2, '0')}`;
        },
  },

  methods: {

      async setCurrentTimer() {
        const web3i = new web3(EVM_SIDECHAIN_RPC);
        const latestBlockNumber = await web3i.eth.getBlockNumber();
        const latestBlock = await web3i.eth.getBlock(latestBlockNumber);

        const ts = Number(latestBlock.timestamp)
        this.currentTime = ts;
        // need to change
        //this.endTime = (Number(this.data.auctionStartTime) + 7*24*60*60)*1000
                this.endTime = (Number(this.data.auctionStartTime) + 20*60)*1000


      },

      updateTimer() {
        const now = new Date().getTime();
        //console.log(now, this.endTime)
        const timeDifference = Math.floor((this.endTime - now) / 1000);

        if (timeDifference <= 0) {
          clearInterval(this.intervalId);
          this.currentTime = 0;
        } else {
          this.currentTime = timeDifference;
        }
      },

      async performInitialization() {
          await this.getAuctionItem()
          await this.getAndLoadImage()
          await this.setCurrentTimer();
          this.intervalId = setInterval(this.updateTimer, 1000);
          await this.fetchBiddersList()

      },

      async getAuctionItem() {
        const loanId = this.$route.params.loanId;
        const web3i = new web3(EVM_SIDECHAIN_RPC);
        try {
            const contract = new web3i.eth.Contract(auctionAbi, AUCTION_CONTRACT_ADDRESS);
            const result = await contract.methods.auction(loanId).call({})
            const jsonString = JSON.stringify(result, (key, value) => (typeof value === 'bigint' ? value.toString() : value));

            this.data = JSON.parse(jsonString);
            this.bidValue = Number(this.getInXRP(String(this.data.highestBindingBid))) + 6
            console.log(this.data)
        } catch (error) {
            console.log(error);
        }
      },

    async getAndLoadImage() {
     const turi = await this.getTokenURI(this.data.nftContractAddress, this.data.tokenId)
     const imobj = await this.decodeTheTokenURIForImage(turi)
     this.imagingURL = imobj.image

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

    async getTokenURI(contractAddress, tokenId) {

           const web3i = new web3(EVM_SIDECHAIN_RPC);
           try {
                const contract = new web3i.eth.Contract(erc721Abi, contractAddress);
                const tokenURI = await contract.methods.tokenURI(tokenId).call({})
                return tokenURI
                
           } catch (error) {
               console.log(error);
               return "https://developers.elementor.com/docs/assets/img/elementor-placeholder-image.png";
           }

    },

    
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
                return "⏳ Not Started"
            } else if (x == 1) {
                return "⏰ Running"
            } else if (x == 2) {
                return "✅ Finalized"
            } else if (x == 3) {
                return "❌ Cancelled"
            } else if (x == 4) {
                return "✅ Ended"
            }
  
    },

    getInXRP(amt) {
        console.log(amt)
        return web3.utils.fromWei(amt, "ether");
    }, 

    async fetchBiddersList() {
        const web3i = new web3(EVM_SIDECHAIN_RPC);
        try {
            const contract = new web3i.eth.Contract(auctionAbi, AUCTION_CONTRACT_ADDRESS);
            const result = await contract.methods.getAllBiddersList(this.data.auctionId).call({})
            let newArr = []
            for(let i = 0; i < result.length; i++) {
                let bibid = await contract.methods.bidsForAuction(this.data.auctionId, result[i]).call()
                let nobj = {
                    bidder: result[i],
                    bid: Number(bibid),
                }

                newArr.push(nobj);
            }
            if(newArr.length > 0) {
                this.biddersList = newArr
            } else {
                this.biddersList = null
            }
            console.log("Printing Bidders List", this.biddersList)
        } catch (error) {
            console.log(error);
            this.biddersList = null;
        }
    },

    async checkIfAuctionHasEnded() {
        this.checkAuctionStateBtnDisabled = true;
          // approve the blockborrower contract address
          const chainId = await this.getCurrentChainId();
          if(chainId != this.evmxrplChainId) {
              this.showAlert("Switch to EVM XRPL chain");
              return;
          }
        const bankId = this.data.bankToWhichItisPledged;
        const loanId = this.data.loanId;
        

        const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL

        try {
                window.contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                const sender = localStorage.getItem("walletAddress");

                let txnParams = {
                    to: BLOCKBORROWER_CONTRACT_ADDRESS,
                    from: sender,
                    data: window.contract.methods.checkIfAuctionHasExpired(bankId,  loanId).encodeABI(),
                }

                const txnHash = await window.ethereum.request({
                        method: 'eth_sendTransaction',
                        params: [txnParams],
                })

                this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
                console.log('Transaction hash:', txnHash)
                setTimeout(() => {
                    this.getTxnReceiptAction(txnHash, "Auction status updated !! Please refresh the page")
                }, 5000);

        } catch (error) {
            console.log(error)
        }

        this.performInitialization()

        this.checkAuctionStateBtnDisabled = false;

    },

    
        async getTxnReceiptAction(txnHash, message) {
            try {
                const web3i = new web3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL

                const receipt = await web3i.eth.getTransactionReceipt(txnHash);
                if (receipt && receipt.status) {
                    this.showAlert(message)
                    console.log(receipt)
                    await this.getAuctionItem()
                    await this.fetchBiddersList()
                } else {
                    this.showAlert("Transaction Failed. check it in blockexplorer::" + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash);
                }
            } catch(err) {
                console.log(err)
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

        async placeABid() {
            this.disableBidBtn = true;
            const uadr = localStorage.getItem("walletAddress")
            const aucId = this.data.auctionId;

             const chainId = await this.getCurrentChainId();
             if(chainId != this.evmxrplChainId) {
                 this.showAlert("Switch to EVM XRPL chain");
                 this.disableBidBtn = false;
                 return;
             }

            if(Number(web3.utils.toWei(this.bidValue, "ether")) < Number(this.data.highestBindingBid) + Number(web3.utils.toWei('5', 'ether')) ) {
                this.showAlert("Bid value should be highest binding bid + 5 ether")
                this.disableBidBtn = false;
                return;
            }
            // find blackListed Number
            let isBlackListed = await this.findBlackListed(aucId, uadr);
            if(isBlackListed) {
                this.showAlert("you were blackListed by the contract");
                this.disableBidBtn = false;
                return;
            } 
            // find user's bid value
            let ubid = await this.findUserBid(aucId, uadr);
            let amtToSend = Number(Number(web3.utils.toWei(this.bidValue, 'ether')) - ubid).toString(16);
            console.log("Amt to send", amtToSend)
            const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL

            try {
                window.contract = new web3i.eth.Contract(auctionAbi, AUCTION_CONTRACT_ADDRESS);

                let txnParams = {
                    to: AUCTION_CONTRACT_ADDRESS,
                    from: uadr,
                    value: amtToSend,
                    data: window.contract.methods.placeBidForAuction(this.data.loanId).encodeABI(),
                }

                const txnHash = await window.ethereum.request({
                        method: 'eth_sendTransaction',
                        params: [txnParams],
                })

                this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
                console.log('Transaction hash:', txnHash)
                setTimeout(() => {
                    this.getTxnReceiptAction(txnHash, "Bid Placed successfully !! Please refresh the page")
                }, 5000);

            } catch(error) {
                console.log(error)
                this.disableBidBtn = false;

            }
            this.disableBidBtn = false;
            await this.getAuctionItem()
            await this.fetchBiddersList()


            
            // input - bidvalue + 1gwei txn to be done.
        },

        async findUserBid(auctionId, addrs) {
            const web3i = new web3(EVM_SIDECHAIN_RPC);
            try {
                const contract = new web3i.eth.Contract(auctionAbi, AUCTION_CONTRACT_ADDRESS);
                const result = await contract.methods.bidsForAuction(auctionId, addrs).call({})
                return Number(result)
            } catch (error) {
                console.log(error);
                return -1;
            }
        },

        async findBlackListed(auctionId, addrs) {
            const web3i = new web3(EVM_SIDECHAIN_RPC);
            try {
                const contract = new web3i.eth.Contract(auctionAbi, AUCTION_CONTRACT_ADDRESS);
                const result = await contract.methods.blackListed(auctionId, addrs).call({})
                return result
            } catch (error) {
                console.log(error);
                return false;
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

    async withdrawAmount() {
        if(this.data.auctionState == 1) {
            this.showAlert("Auction is running ! Withdrawn once it")
            return;
        }
        const loanId = this.$route.params.loanId;
        const uadr = localStorage.getItem("walletAddress");
        const chainId = await this.getCurrentChainId();
        if(chainId != this.evmxrplChainId) {
            this.showAlert("Switch to EVM XRPL chain");
            this.disableBidBtn = false;
            return;
        }

        try {
            const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL
            window.contract = new web3i.eth.Contract(auctionAbi, AUCTION_CONTRACT_ADDRESS);

            let txnParams = {
                to: AUCTION_CONTRACT_ADDRESS,
                from: uadr,
                data: window.contract.methods.withdrawMoney(loanId).encodeABI(),
            }

            const txnHash = await window.ethereum.request({
                    method: 'eth_sendTransaction',
                    params: [txnParams],
            })

            this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
            console.log('Transaction hash:', txnHash)
            setTimeout(() => {
                this.getTxnReceiptAction(txnHash, "Withdrawn successfully !! Please refresh the page")
            }, 5000);

        } catch (error) {
            console.log(error);
        }
    }


  }

}
</script>

<style scoped>
#btn-cntnr {
    text-align: right;
    padding-right: 20px;
}

.au-btns {
    padding: 10px;
    border-radius: 5px;
    background-image: linear-gradient(to right, lightgreen, lightseagreen);
    cursor: pointer;
    font-family: monospace;
    margin-right: 15px;
}

.au-btns:hover:enabled {
    background-image: linear-gradient(to left, lightgreen, lightseagreen);

}

img {
    width: 400px;
    border-radius: 5px;
    border: 2px solid black;
}

#bid-btn {
    padding: 2px;
    width: 150px;
}

#place-a-bid {
    margin-top: 20px;
}

table {
    width: 100%;
    border-collapse: collapse;
    margin-top: 20px;
    background-color: skyblue;
}

th, td {
    border: 1px solid black;
    padding: 8px;
    text-align: left;

}
tr:hover {
   background-color: aliceblue;

}

#auc-end-tm {
    text-align: left;
}
</style>