<template>
  <h2 id="bl-v"> <u>All Bank List</u> </h2>
  <p id="bl-ins"><i>{{msgData}}</i></p>
  <div id="bnk-srch-all-bnk">
    <form id="all-bnk-srch-frm" @submit.prevent >
        <label style="color: purple; margin-right: 10px;"> <b>Bank Id</b> </label> <input type="number" min="0" v-model="bankNumberSearch"/> 
        <Button id="srch-btn-frm" @click="bankNumberSearchHere"> Search </Button>
    </form>
  </div>
  <div class="card-container">
    <CardComponent v-for="card in cards" :key="card.id" :data="card" @click="routeToBankDetails(card.bankId)" @custom-event="handleCustomEvent"/>
  </div>
</template>

<script>
import CardComponent from '../components/CardComponent.vue';
import blockBorrowerAbi from '../contract_abis/blockborrower_abi.json'
import web3 from 'web3';
import {EVM_SIDECHAIN_RPC, BLOCKBORROWER_CONTRACT_ADDRESS} from '../env/env.js';


export default {
  components: {
    CardComponent,
  },
  data() {
    return {
      cards:[],
      copyCard: [],

      msgData: "Loading...",
      bankNumberSearch: 0,
    };
  },

  created() {
     this.getAllBanks()
  },



  methods: {

    async handleCustomEvent() {
      this.cards = []
      this.copyCard = []
      this.msgData = "Loading..."
      this.bankNumberSearch = 0
      await this.getAllBanks()
    },

      bankNumberSearchHere() {
          this.msgData = "Searching..."
          console.log("HEre")
          this.cards = []
          if(this.bankNumberSearch == 0) {
              this.getAllBanks();
              return;
          }
          for (var i = 0; i < this.copyCard.length; i++) {
              if(this.copyCard[i].bankId == this.bankNumberSearch) {
                  this.cards = [this.copyCard[i]];
                  this.msgData = "Put 0 and search to get all banks again..."
                  return;
              }
          }

        this.msgData = "Put 0 and search to get all banks again..."

      },

      async getAllBanks() {
          var enableJoinButton = false;

          const bankCount = await this.getTotalCountOfBank();
          if(bankCount == 0) {
              this.msgData = "No Banks Found"
              return;
          }

          var i = 1;

          this.cards = []
          this.msgData = "Loading..."
          for(i = 1; i <= bankCount; i++) {
              const bankDetail = await this.getBankDetails(i);
                bankDetail.bankId = i;
                bankDetail.indexi = i-1;
                bankDetail.shouldShowJoin = enableJoinButton;
                this.cards.push(bankDetail)
          }

          this.msgData = "Put 0 and search to get all banks again..."
          this.copyCard = JSON.parse(JSON.stringify(this.cards));



          //console.log(JSON.stringify(this.cards))
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

      routeToBankDetails(bankId) {
          console.log("Inside details")
        this.$router.push('/personal_bank_view/' + bankId);
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

.card-container .card-component {
  flex-basis: calc(50% - 10px); /* 10px for margin between cards */
  margin-bottom: 10px;
  box-sizing: border-box;
}

.card-component {
    cursor: pointer;
    transition: transform 0.3s ease; /* Add a smooth transition effect */
  /* Initial styling or transformation */
     transform: scale(1);
}

.card-component:hover {
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
</style>
