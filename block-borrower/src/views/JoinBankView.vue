<template>
  <h2 id="bl-v"> <u>Bank List</u> </h2>
  <p id="bl-ins"><i>{{msgData}}</i></p>
  <div class="card-container">
    <CardComponent v-for="card in cards" :key="card.id" :data="card" />
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

      msgData: "Loading...",
    };
  },

  created() {
     this.getAllBanks()
  },

  methods: {
      async getAllBanks() {
          const a = await this.getBankIdForUser(localStorage.getItem("walletAddress"))
          var enableJoinButton = false;
        console.log("A", a)
        if (a == 0 || a==-1) {
              enableJoinButton = true;
          }
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
              if(bankDetail.functional) {
                bankDetail.bankId = i;
                bankDetail.indexi = i-1;
                bankDetail.shouldShowJoin = enableJoinButton;
                this.cards.push(bankDetail)
              }
          }

          this.msgData = ""


          console.log(JSON.stringify(this.cards))
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

        async getBankIdForUser(addr) {
            console.log("inside get bank for user id");
            const web3i = new web3(EVM_SIDECHAIN_RPC);
            try {
                const contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                const result = await contract.methods.lenderAddressToBankMapping(addr).call({})
                console.log(result);
                return result;
            } catch (error) {
                console.log(error);
                return -1;
            }
        },
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
</style>
