<template>
  <div class="my-ln-vw-nft-dtls">
      <h3> My Loans </h3>
      <p style="color:red"> To search by loanId set bank Id as 0, and vice versa. To see all loans search by bankId and loanId as 0 </p>
      <form @submit.prevent>
        <label > <b>Loan Id</b> </label> <input type="number" v-model="slid"/>
        <label> <b>Bank Id</b> </label> <input type="number"  v-model="sbid"/>
        <Button class="frm-ln-btn" @click="searchClick" :disabled="disablSearch"> Search </Button>

      </form>
      <hr>
      <div class="card-container">
        <LoanAndNFTDetailsCard v-for="(card, index) in cards" :key="index" :data="card" @click="loanDetailsViewRender(card.bankToWhichItisPledged, card.pledger, card.loanId)"/>
    </div>
  </div>
</template>

<script>
import LoanAndNFTDetailsCard from '@/components/LoanAndNFTDetailsCard.vue'
import blockBorrowerAbi from '../contract_abis/blockborrower_abi.json'
import web3 from 'web3';
import {EVM_SIDECHAIN_RPC, BLOCKBORROWER_CONTRACT_ADDRESS} from '../env/env.js';
// @ is an alias to /src
export default {
  name: 'MyLoanCustomerView',
  components: {
    LoanAndNFTDetailsCard,
  },


  data() {
      return {
          cards: [
          ],
          copyCards: [

          ],
          sbid: 0,
          slid: 0,
          disablSearch: false,
      }
  },

  created() {
    this.getMyLoanDetails();
  },

  methods: {
    async getMyLoanDetails() {

      let walletAddress = localStorage.getItem("walletAddress")
      this.cards = []
      this.copyCards = []

      try {
        // get all bank ids of customer
        let bankIdsOfCustomer = await this.getAllBankIdsOfCustomer(walletAddress);

        // for each bankId getAllNFT holdings of customer
        for(let i = 0; i < bankIdsOfCustomer.length; i++) {
          let nftHoldings = await this.getNFTHoldingOfCustomerForBankId(bankIdsOfCustomer[i], walletAddress);
          this.cards = [...this.cards, ...nftHoldings]
        }

        console.log(this.cards)
        this.copyCards = JSON.parse(JSON.stringify(this.cards));

      } catch (error) {
        console.log(error)
      }


    },

   async getAllBankIdsOfCustomer(customerAddress) {
      const web3i = new web3(EVM_SIDECHAIN_RPC);
      try {
          const contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
          let result = await contract.methods.getAllBanksOfCustomer().call({from:customerAddress})
          return result
      } catch (error) {
          console.log(error);
          return [];
      }
    },

    async getNFTHoldingOfCustomerForBankId(bankId, customerAddress) {
      console.log(bankId, customerAddress);
      const web3i = new web3(EVM_SIDECHAIN_RPC);
      try {
          const contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
          console.log("INitialized contract:: ")
          let result = await contract.methods.getAllNftDepositsOfBank(bankId, customerAddress).call()
          console.log("i got result")
          const jsonString = JSON.stringify(result, (key, value) => (typeof value === 'bigint' ? value.toString() : value));

          return JSON.parse(jsonString);
      } catch (error) {
          console.log(error);
          return [];
      }
    },

    async searchClick() {
      this.disablSearch = true;
      try {
        this.cards = []
        const li = Number(this.slid)
        const bi = Number(this.sbid)
        console.log(li, bi)

        if(li == 0 && bi != 0) {
          // search by bank id
          for(let i = 0; i < this.copyCards.length; i++) {
            if(this.copyCards[i].bankToWhichItisPledged == bi) {
              this.cards = [this.copyCards[i]]
            }
          }

        } else if (li != 0 && bi == 0) {
          // search by loanid
          for(let i = 0; i < this.copyCards.length; i++) {
            if(this.copyCards[i].loanId == li) {
              this.cards = [ this.copyCards[i] ]
            }
          }

        } else if (li == 0 && bi == 0) {
          // search for all loans
          await this.getMyLoanDetails();

        } else if (li != 0 && bi != 0) {
          // search by both
          for(let i = 0; i < this.copyCards.length; i++) {
            if(this.copyCards[i].loanId == li && this.copyCards[i].bankToWhichItisPledged == bi) {
              this.cards = [ this.copyCards[i] ]
            }
          }
        }
      } catch(error) {
        this.disablSearch = false;
      }
      
      this.disablSearch = false;

    },

    loanDetailsViewRender(bankId, customerAddress, loanId) {
      this.$router.push("/loan_details_view/"+String(bankId) + "/" + String(loanId)+ "/" + String(customerAddress))
    }

  },
}
</script>

<style scoped>
.card-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
}

.card-container .loan-and-nft-details-card {
  flex-basis: calc(50% - 10px); /* 10px for margin between cards */
  margin-bottom: 10px;
  box-sizing: border-box;
}

.loan-and-nft-details-card {
    cursor: pointer;
    transition: transform 0.3s ease; /* Add a smooth transition effect */
  /* Initial styling or transformation */
     transform: scale(1);
}

.loan-and-nft-details-card:hover {
    transform: scale(1.03); /* Increase the scale to 1.2 on hover */
}

input {
  margin-right: 30px;
}

form {
  margin-bottom: 30px;
}

.frm-ln-btn {
  border-radius: 5px;
  background-image: linear-gradient(to right, aliceblue, skyblue);
  cursor: pointer;
}
.frm-ln-btn:enabled:hover {
  background-image: linear-gradient(to left, aliceblue, skyblue);
}
</style>
