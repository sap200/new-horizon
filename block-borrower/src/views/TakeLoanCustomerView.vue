<template>
  <div class="tk-ln-cst-view">
      <h2> <b><u>Take Loan</u></b></h2>
      <p style="color:blue;"> If you have transferred your NFT to EVN XRPL sidechain using our NFT Bridge, here is the contract address: {{mntCtrAdrs}} </p>
    <hr>
    <div id="evalt-nft">
        <h3 style="color:green;text-align:left;"> Step 1: Evaluate your NFT </h3>

        <label style="color:darkgreen"> <b>Please check if it is sepolia NFT, and it has been transferred using our NFT Bridge, If yes then please put in tokenId and contract of sepolia network</b>  
        <input type="checkbox" v-model="isChecked" />  </label> <br>

        <form class="apprv-frm" @submit.prevent>
            <p> <b>NFT Contract Address :</b> </p>
            <input type="text" v-model="nftContractAddress"/>
            <p> <b>Token Id :</b> </p>
            <input type="number" min="1" v-model="nftTokenId"/>

            <div class="btn-ctr" >
            <Button class="tk-ln-btn" @click="evaluateNFT" :disabled="this.nftContractAddress == ''"> Evaluate </Button>
            </div>    
            <br>
            <div id="nft-cntnr">
                <img width="200" height="200" :src="myImageURL" />
                <p> {{nftSymbol}} - {{nftName}}</p>
                <p style="color:red"> {{ownerOrNot}} </p>
            </div>
        </form>
        <p>{{getDisplayablePrice()}}</p>
        <p> Loan Offer: <span style="color:red; font-size: 20px;"> <b> {{this.nftPrice + " XRP"}} </b> </span></p>
    </div>
    <br>
    <hr>
    <div id="crd-vw">
    <h3 style="color:green;text-align:left;"> Step 2: Select Lender from whom you want to take offer, you can select only one ! The selected one appears in step 4 section </h3>
        <p style="color:red" v-if="this.cards.length == 0"> No offers Found !! </p>

        <div class="card-container" >
            <loan-offer-card v-for="(card, idx) in cards" :key="idx" :data="card" @click="selectTheOffer(idx)"/>
        </div>
    </div>

    <hr>

    <div id="ln-apprv-ctrct">
        <h3 style="color:green;text-align:left;"> Step 3: Approve the bank to take right of the NFT </h3>
        <p style="color:green"> <b>{{this.approvalProvidedText}}</b> </p>
        <form class="apprv-frm" @submit.prevent>
            <div class="btn-ctr" >
            <Button class="tk-ln-btn" @click="approveTheLender" :disabled="this.cardi == null"> Approve </Button>
            </div>
        </form>
    </div>
    <br>
    <hr>
    <div id="finlz-ln-ofr">
        <h3 style="color:green; text-align:left;"> step 4: Pledge your NFT and claim your loan offer </h3>
        <loan-offer-card :data="this.cardi" v-if="this.cardi != null"/>
        <Button class="tk-ln-btn" :disabled="!this.approvalStageDone" @click="sendNFTToContract"> Pledge and Claim </Button>

    </div>


  </div>

<alert-modal :is-visible="isModalVisible" :message="alertMessage" @close="onCloseModal" />

</template>

<script>
import erc721Abi from '../contract_abis/erc721_abi.json'
import blockBorrowerAbi from '../contract_abis/blockborrower_abi.json'
import AlertModal from '../components/AlertModal.vue';
import {createAlchemyWeb3} from "@alch/alchemy-web3"
import web3 from 'web3';
import {EVM_SIDECHAIN_RPC, MINTER_CONTRACT_ADDRESS, BLOCKBORROWER_CONTRACT_ADDRESS, EVM_SIDECHAIN_BLOCK_EXPLORER} from '../env/env.js';
// @ is an alias to /src
import {MIN_LN_AMT_RND, MAX_LN_AMT_RND} from '@/env/env.js'
import LoanOfferCard from '@/components/LoanOfferCard.vue'
export default {
  name: 'TakeLoanCustomerView',

  components: {
      LoanOfferCard,
      AlertModal,
  },

  data() {
      return {
          nftPrice: 0,
          nftContractAddress: "",
          nftTokenId: 1,
          cards: [
              
          ],
          cardi: null,
          isModalVisible: false,
          alertMessage: "",
          nftName: "",
          nftSymbol: "",
          ownerOrNot: "",
          ownerOfNFT: "",
          tokenURI: "",
          isChecked: true,
          mntCtrAdrs: MINTER_CONTRACT_ADDRESS,
          myImageURL: "https://developers.elementor.com/docs/assets/img/elementor-placeholder-image.png",
          approvalProvidedText: "",
          evmxrplChainId: "0x15f902",
          originalNFTId: 1,
          originalNFTContractAddress: "",
          approvalStageDone: false,
          evaluationStageDone: false,
          loanReceivedText: "pledge to claim your loan"
      }
  },


  methods: {

      selectTheOffer(idx) {
          this.cardi = this.cards[idx];
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

     async approveTheLender() {
         if(this.cardi == null) {
             this.showAlert("Please select a offer before proceeding !")
             return;
         }
          // approve the blockborrower contract address
          const chainId = await this.getCurrentChainId();
          if(chainId != this.evmxrplChainId) {
              this.showAlert("Switch to EVM XRPL chain");
              return;
          }
          let contractAddress = MINTER_CONTRACT_ADDRESS;
          if(!this.isChecked) {
              contractAddress = this.nftContractAddress.trim();
          }

            const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL
            try {
                window.contract = new web3i.eth.Contract(erc721Abi, contractAddress);
                const sender = localStorage.getItem("walletAddress");

                let txnParams = {
                    to: contractAddress,
                    from: sender,
                    data: window.contract.methods.approve(BLOCKBORROWER_CONTRACT_ADDRESS,  this.originalNFTId).encodeABI(),
                }

                const txnHash = await window.ethereum.request({
                        method: 'eth_sendTransaction',
                        params: [txnParams],
                })

                this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
                console.log('Transaction hash:', txnHash)
                setTimeout(() => {
                    this.getTxnReceiptAction(txnHash, "Approval provided, Proceed to Pledge !!")
                }, 5000);

            } catch (error) {
                this.showAlert(JSON.stringify(error))
                console.error('Error:', error);
            }
      },

      async sendNFTToContract() {
           if(this.cardi == null) {
             this.showAlert("Please select a offer before proceeding !")
             return;
         }
          // approve the blockborrower contract address
          const chainId = await this.getCurrentChainId();
          if(chainId != this.evmxrplChainId) {
              this.showAlert("Switch to EVM XRPL chain");
              return;
          }

        const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL
        try {
            window.contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
            const sender = localStorage.getItem("walletAddress");
            let txnParams = {
                to: BLOCKBORROWER_CONTRACT_ADDRESS,
                from: sender,
                // eslint-disable-next-line
                data: window.contract.methods.depositNFT(this.cardi.bankId, this.originalNFTContractAddress, this.originalNFTId, BigInt(this.cardi.nftPrice)).encodeABI(),
            }
            const txnHash = await window.ethereum.request({
                        method: 'eth_sendTransaction',
                        params: [txnParams],
            })

            this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
             console.log('Transaction hash:', txnHash)
            setTimeout(() => {
                    this.getTxnReceiptAction(txnHash, "NFT Pledged successfully, check your bank account in 5-10 seconds, you might have received the loan amount!!")
            }, 5000);

        } catch (error) {
                this.showAlert(JSON.stringify(error))
                console.error('Error:', error);
        }



      },

        async getTxnReceiptAction(txnHash, message) {
            try {
                const web3i = new web3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL

                const receipt = await web3i.eth.getTransactionReceipt(txnHash);
                if (receipt && receipt.status) {
                    this.approvalProvidedText = "Approval Provided ! Please proceed to pledge NFT !"
                    this.approvalStageDone = true;
                    this.showAlert(message)
                    console.log(receipt)
                } else {
                    this.showAlert("Transaction Failed. check it in blockexplorer::" + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash);
                }
            } catch(err) {
                console.log(err)
            }
        },

      async evaluateNFT() {
          
          this.nftSymbol = ""
          this.nftSymbol = ""
          this.ownerOrNot = ""
          this.ownerOfNFT = ""
          this.tokenURI = ""
          this.cards = []
          this.cardi = null
          this.myImageURL = "https://developers.elementor.com/docs/assets/img/elementor-placeholder-image.png";
          if(this.nftContractAddress == "") {
              return;
          }



          try {
                        // Get the symbol, owner and tokenURI of NFT
                if(this.isChecked) {
                    const myTokenId = await this.getTokenToDetailsXRPL(this.nftContractAddress.trim(), this.nftTokenId);
                    console.log(myTokenId.nftId);
                    this.originalNFTId = myTokenId.nftId;
                    this.originalNFTContractAddress = MINTER_CONTRACT_ADDRESS
                    await this.getSymbolNameOwnerAndTokenURIOfNFT(MINTER_CONTRACT_ADDRESS.trim(), Number(myTokenId.nftId))
                } else {

                    await this.getSymbolNameOwnerAndTokenURIOfNFT(this.nftContractAddress.trim(), this.nftTokenId)
                    this.originalNFTId = this.nftTokenId;
                    this.originalNFTContract = this.nftContractAddress.trim();
                }
                console.log("NFT DETAILS::", this.nftSymbol, this.nftName, this.ownerOfNFT, this.tokenURI);
                if(this.ownerOfNFT.toLowerCase() != localStorage.getItem("walletAddress").toLowerCase()) {
                    this.ownerOrNot = "You are not the owner of NFT !! Owner is " + this.ownerOfNFT;
                    return;
                }

                const xder = await this.decodeTheTokenURIForImage(this.tokenURI)
                this.myImageURL = xder.image;


                const priceGen = this.getRandomNumber(MIN_LN_AMT_RND, MAX_LN_AMT_RND)
                this.nftPrice = priceGen;
                console.log(this.nftPrice)

                this.getAllBanksWithOffer(this.nftPrice);
                this.evaluationStageDone = true;


                } catch(error) {
                    this.showAlert(error)
                    console.log(error)
                }

      },

      async getAllBanksWithOffer(amount) {
          console.log(amount)
          // get all banks that are functional and its totalBalance > amount, return the list of banks append bankId with it.
          const bankCount = await this.getTotalCountOfBank();
          if(bankCount == 0) {
              return;
          }

          var i = 1;

          this.cards = []
          for(i = 1; i <= bankCount; i++) {
              let bankDetail = await this.getBankDetails(i);
              const bankBalance = web3.utils.fromWei(bankDetail.totalBalance, "ether");
              console.log("Bank Balance: ", bankBalance);
              if(Number(bankBalance) >= amount) {
                bankDetail.bankId = i;
                bankDetail.indexi = i-1;
                bankDetail.nftPrice = web3.utils.toWei(String(amount), "ether");
                this.cards.push(bankDetail)
              }
          }

      },

      async getTotalCountOfBank() {
            console.log("Get Total Count of Bank");
            const web3i = new web3(EVM_SIDECHAIN_RPC);
            try {
                const contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                const result = await contract.methods.getTotalNumberOfBanks().call({})
                console.log(result);
                return Number(result)
            } catch (error) {
                console.log(error);
                return -1;
            }
      },

      async getBankDetails(bankId) {
           console.log("Bank Details");
           const web3i = new web3(EVM_SIDECHAIN_RPC);
           try {
                const contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                const result = await contract.methods.bankList(bankId).call({})
                let jsonString = ((obj) => JSON.stringify(obj, (key, value) => (typeof value === 'bigint' ? value.toString() : value), 2))(result);

                return JSON.parse(jsonString);
           } catch (error) {
               console.log(error);
               return null;
           }
      },

      async getSymbolNameOwnerAndTokenURIOfNFT(contractAddress, tokenId) {

           const web3i = new web3(EVM_SIDECHAIN_RPC);
           try {
                const contract = new web3i.eth.Contract(erc721Abi, contractAddress);
                const symbol = await contract.methods.symbol().call({})
                this.nftSymbol = symbol;
                const name = await contract.methods.name().call({})
                this.nftName = name;

                /* eslint-disable */                
                const onr = await contract.methods.ownerOf(BigInt(tokenId)).call({})
                this.ownerOfNFT = onr
                const tokenURI = await contract.methods.tokenURI(BigInt(tokenId)).call({})
                this.tokenURI = tokenURI
                
           } catch (error) {
               this.showAlert(JSON.stringify(error))
               console.log(error);
               return null;
           }

      },

      getDisplayablePrice() {
          if(this.nftPrice == 0) {
              return ""
            } else {
                return " Your NFT is worth " +  this.nftPrice + " XRP!! Please proceed with offer selection in case you want to avail loan below " +  this.nftPrice  + " XRP"
            }
      },

    getRandomNumber(min, max) {
        return Math.floor(Math.random() * (max - min + 1)) + min;
    } ,

    showAlert(alertMsg) {
        this.isModalVisible = true;
        this.alertMessage = alertMsg;
    },
    
    onCloseModal() {
        this.isModalVisible = false;
        this.alertMessage = "";
    },

    async getTokenToDetailsXRPL(nftContractAddress, tokenId) {
        console.log("inside")
        const web3i = new web3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL
        try {
            const contract = new web3i.eth.Contract(erc721Abi, MINTER_CONTRACT_ADDRESS);
            const result = await contract.methods.tokenIdToNFTDetailsMapping(nftContractAddress, tokenId).call({})
            // convert bigInt to JSON
            const jsonString = JSON.stringify(result, (key, value) => {
                if (typeof value === 'bigint') {
                    return value.toString();
                } 

                return value
            });                
            return JSON.parse(jsonString)

        } catch (error) {
            console.error('Error:', error);
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
    }

    
    
  }
}
</script>

<style scoped>
.tk-ln-btn {
    width: 150px;
    height: 30px;
    background-image: linear-gradient(to right, aliceblue, skyblue);
    border-radius: 5px;
    cursor: pointer;
    margin-left: 25px;
    font-family: monospace;
}
.tk-ln-btn:hover:enabled {
    background-image: linear-gradient(to left, aliceblue, skyblue);
}

.tk-ln-btn:disabled {
    background-image: linear-gradient(to left, rgb(221, 217, 217), white);
}

input {
  border: none;          /* Remove default border */
  border-bottom: 1px solid #000;  /* Add bottom border */
  outline: none;         /* Remove default focus outline */
  width: 100%;           /* Make the input full width */
  box-sizing: border-box; /* Include padding and border in width calculation */
  background: transparent;
  font-size: 15px;
  font-family: monospace;

}

.btn-ctr {
    width: 100%;
    text-align: center;
    margin-top: 25px;
    font-family: monospace;
}

.apprv-frm {
    text-align: left;
}

hr {
      border: 0;            /* Remove the default border */
      height: 3px;          /* Set the height of the line */
      background-color: orange; /* Set the color of the line */
      margin: 20px 0;        /* Set top and bottom margin */
    }

.card-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
}

.card-container .loan-offer-card {
  flex-basis: calc(50% - 10px); /* 10px for margin between cards */
  margin-bottom: 10px;
  box-sizing: border-box;
}

.loan-offer-card {
    transition: transform 0.3s ease; /* Add a smooth transition effect */
  /* Initial styling or transformation */
     transform: scale(1);
}

.loan-offer-card:hover {
    transform: scale(1.03); /* Increase the scale to 1.2 on hover */
}

#nft-cntnr {
    text-align: center;
}

</style>
