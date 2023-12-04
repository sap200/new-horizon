<template>
  <div class="my-ln-vw-nft-dtls">
      <h3> Loans Provided </h3>
      <p style="color:red"> To search by loanId set customer address as empty, and vice versa. To see all loans search by customer address as empty and loanId as 0 </p>
      <form @submit.prevent>
        <label > <b>Loan Id</b> </label> <input type="number" v-model="slid"/>
        <label> <b>Customer's Address</b> </label> <input type="text"  v-model="cadr"/>
        <Button class="frm-ln-btn" @click="searchClick" :disabled="disablSearch"> Search </Button>

      </form>
      <hr>
      <div class="card-container">
        <LoanAndNFTDetailsCard v-for="(card, index) in cards" :key="index" :data="card" @click="routeToBankViewOfLoanDetails(index)"/>
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
          slid: 0,
          cadr: "",
          disablSearch: false,
      }
  },

  created() {
    this.getBankLoanDetails();
  },

  methods: {
    async getBankLoanDetails() {
      this.cards = []
      this.copyCards = []

      try {
        // determine my bank
        const bankId = await this.getUsersBankIdByAddress();
        console.log("Bank ID", bankId)
        // get all bank ids of customer
        let customersOfBank = await this.getAllCustomersOfBank(bankId);

        // for each bankId getAllNFT holdings of customer
        for(let i = 0; i < customersOfBank.length; i++) {
          let nftHoldings = await this.getNFTHoldingOfCustomerForBankId(bankId, customersOfBank[i]);
          this.cards = [...this.cards, ...nftHoldings]
        }

        console.log(this.cards)
        this.copyCards = JSON.parse(JSON.stringify(this.cards));

      } catch (error) {
        console.log(error)
      }
    },

    async getUsersBankIdByAddress() {
        const addr = localStorage.getItem("walletAddress");
        const web3i = new web3(EVM_SIDECHAIN_RPC);
         try {
                const contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                const result = await contract.methods.lenderAddressToBankMapping(addr).call({})
                return result;
           } catch (error) {
               console.log(error);
               return -1;
           }
    },

   async getAllCustomersOfBank(bankId) {
      const web3i = new web3(EVM_SIDECHAIN_RPC);
      try {
          const contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
          let result = await contract.methods.getAllCustomersOfBank(Number(bankId)).call({})
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
          let result = await contract.methods.getAllNftDepositsOfBank(bankId, customerAddress).call({from:customerAddress})
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
        const cadr =this.cadr
        console.log(li, cadr)

        if(li == 0 && cadr != "") {
          // search by bank id
          for(let i = 0; i < this.copyCards.length; i++) {
            if(this.copyCards[i].pledger.toLowerCase() == cadr.toLowerCase()) {
              this.cards = [this.copyCards[i]]
            }
          }

        } else if (li != 0 && cadr == "") {
          // search by loanid
          for(let i = 0; i < this.copyCards.length; i++) {
            if(this.copyCards[i].loanId == li) {
              this.cards = [ this.copyCards[i] ]
            }
          }

        } else if (li == 0 && cadr == "") {
          // search for all loans
          await this.getBankLoanDetails();

        } else if (li != 0 && cadr != "") {
          // search by both
          for(let i = 0; i < this.copyCards.length; i++) {
            if(this.copyCards[i].loanId == li && this.copyCards[i].pledger == cadr) {
              this.cards = [ this.copyCards[i] ]
            }
          }
        }
      } catch(error) {
        this.disablSearch = false;

      }

      this.disablSearch = false;

      

    },

    routeToBankViewOfLoanDetails(idx) {
      this.$router.push("/bank_loan_details_View/" + this.cards[idx].bankToWhichItisPledged + "/" + this.cards[idx].loanId + "/" + this.cards[idx].pledger)
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
