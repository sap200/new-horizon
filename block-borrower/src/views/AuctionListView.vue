<template>
  <h2 id="bl-v"> <u>All Auction List</u> </h2>
  <p id="bl-ins"><i>{{msgData}}</i></p>
  <div id="bnk-srch-all-bnk">
    <form id="all-bnk-srch-frm" @submit.prevent >
        <label style="color: purple; margin-right: 10px;"> <b>Loan Id</b> </label> <input type="number" min="0" v-model="loanIdSearch"/> 
        <Button id="srch-btn-frm" :disabled="disableLoanSearchBtn" @click="loanNumberSearchHere"> Search </Button>
    </form>
  </div>
  <div class="card-container" v-if="cards.length != 0">
    <AuctionCard v-for="card in cards" :key="card.id" :data="card" @click="routeToAuctionDetails(card.loanId)"/>
  </div>
</template>

<script>
import AuctionCard from '../components/AuctionCard.vue';
import blockBorrowerAbi from '../contract_abis/blockborrower_abi.json'
import auctionAbi from '../contract_abis/auction_abi.json'
import web3 from 'web3';
import {EVM_SIDECHAIN_RPC, BLOCKBORROWER_CONTRACT_ADDRESS, AUCTION_CONTRACT_ADDRESS} from '../env/env.js';


export default {
  components: {
    AuctionCard,
  },
  data() {
    return {
      cards:[],
      copyCards: [],

      msgData: "Loading...",
      loanIdSearch: 0,
      disableLoanSearchBtn: true,
      mlid: 0,
    };
  },

  created() {
      this.getAllAuctions()
  },



  methods: {

 

      loanNumberSearchHere() {
          this.disableLoanSearchBtn = true;
          this.msgData = "Searching..."
          this.cards = []
          if(this.loanIdSearch == 0) {
              this.getAllAuctions();
              return;
          }
          for (var i = 0; i < this.copyCards.length; i++) {
              if(this.copyCards[i].loanId == this.loanIdSearch) {
                  this.cards = [this.copyCards[i]];
                  this.msgData = "Put 0 and search to get all banks again..."
                  break;
              }
          }

        this.msgData = "Put 0 and search to get all banks again..."
        this.disableLoanSearchBtn = false;

      },

      async getAllAuctions() {
          this.cards = []
          this.copyCards = []
          const bankCount = await this.getTotalCountOfBank();
          if(bankCount == 0) {
              this.msgData = "No Banks Found"
              return;
          }

          console.log(bankCount);

          console.log("Crossed Total Bank Count")

          this.msgData = "Loading..."

          for(let i = 1; i <= bankCount; i++) {
              const allCustomers = await this.getAllCustomersOfBank(i);
              for (let j = 0; j < allCustomers.length; j++) {
                  const myBankId = i;
                  const myCustomerAddress = allCustomers[j];
                  //console.log(myBankId, myCustomerAddress)
                  const myNFTDeposits = await this.getAllNFTDepositsOfBankByCustomer(myBankId, myCustomerAddress);
                  //console.log(myNFTDeposits);
                  for(let k = 0; k < myNFTDeposits.length; k++) {
                    const myLoanId = myNFTDeposits[k].loanId;
                    const result = await this.getAuctionDetails(myLoanId)
                    //console.log(result)
                    const jsonString = JSON.stringify(result, (key, value) => (typeof value === 'bigint' ? value.toString() : value));
                    const myResult = JSON.parse(jsonString);
                    if(myResult.auctionId != 0) {
                        this.cards.push(JSON.parse(jsonString));
                    }
                  }
              }
          }

        this.copyCards = JSON.parse(JSON.stringify(this.cards));
        console.log(this.cards)
        this.disableLoanSearchBtn = false;
        
        console.log("Function ended");
        this.msgData = "Please put 0 and search to search for all loans in auction"

    },

      async getAuctionDetails(loanId) {
            const web3i = new web3(EVM_SIDECHAIN_RPC);
            try {
                const contract = new web3i.eth.Contract(auctionAbi, AUCTION_CONTRACT_ADDRESS);
                const result = await contract.methods.auction(loanId).call({})
                return result
            } catch (error) {
                console.log(error);
                return null;
            }
      },

        async getAllNFTDepositsOfBankByCustomer(bankId, customerAddress) {
            const web3i = new web3(EVM_SIDECHAIN_RPC);
            try {
                const contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                const result = await contract.methods.getAllNftDepositsOfBank(bankId, customerAddress).call({})
                return result
            } catch (error) {
                console.log(error);
                return null;
            }
      },

      async getAllCustomersOfBank(bankId) {
            const web3i = new web3(EVM_SIDECHAIN_RPC);
            try {
                const contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                const result = await contract.methods.getAllCustomersOfBank(bankId).call({})
                return result
            } catch (error) {
                console.log(error);
                return null;
            }

      },

      async getTotalCountOfBank() {
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

      routeToAuctionDetails(loanId) {
          this.$router.push("/auction_item_view/"+loanId);

      }

  }
};
</script>

<style>
.card-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
}

.card-container .auction-card {
  flex-basis: calc(50% - 10px); /* 10px for margin between cards */
  margin-bottom: 10px;
  box-sizing: border-box;
}

.auction-card  {
    cursor: pointer;
    transition: transform 0.3s ease; /* Add a smooth transition effect */
  /* Initial styling or transformation */
     transform: scale(1);
}

.auction-card:hover {
    transform: scale(1.03); /* Increase the scale to 1.2 on hover */
}

#bl-v {
    color: blue;
}

#bl-ins {
    color: green;
}

#bnk-srch-all-bnk {
    margin-bottom: 20px;
    padding: 10px;
}

#srch-btn-frm {
    font-family: monospace;
    margin-left: 20px;
    border-radius: 5px;
    background-image: linear-gradient(to right, pink, violet);
    cursor: pointer;
    width: 100px;
}

#btn-crnt {
    text-align: right;
}

#btn-crnt-1 {
    padding: 7px;
    font-family: monospace;
    border-radius: 5px;
    background-image: linear-gradient(to right, pink, lightpink);
    cursor: pointer;
}

#btn-crnt-1:hover {
        background-image: linear-gradient(to left, pink, lightpink);

}
</style>
