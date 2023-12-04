<template>
  <h2> Bank Detail </h2>
  <div class="prnt-div-my-bnk">
    <p id="fn-pr"> <b> {{bankStatus}} </b> </p>
    <div>
        <table>
            <tbody>
                <tr>
                    <td class="td-heading"> <b>Bank Name</b> </td>
                    <td class="td-value"> {{bankData.bankName}} </td>
                    <td class="td-heading"> <b>Bank ID</b> </td>
                    <td class="td-value"> {{bankData.bankId}} </td>
                    <td class="td-heading"> <b> Admin </b></td>
                    <td class="td-value"> {{bankData.admin}} </td>
                </tr>
                 <tr>
                    <td class="td-heading"> <b>Interest Rate</b> </td>
                    <td class="td-value"> {{bankData.interestRate}}% </td>
                    <td class="td-heading"> <b>Tenure</b> </td>
                    <td class="td-value"> {{bankData.tenure}} years </td>
                    <td class="td-heading"> <b> Total Balance </b></td>
                    <td class="td-value"> {{bankData.totalBalance/1000000000000000000}} XRP </td>
                </tr>
                <tr>
                    <td class="td-heading"> <b>Loan Percentage</b> </td>
                    <td class="td-value"> {{bankData.loanPercentage}}% </td>
                    <td class="td-heading"> <b>Defaulting Limit</b> </td>
                    <td class="td-value"> {{bankData.defaultingLimit}} </td>
                    <td class="td-heading"> <b> Defaulting TIme Limit </b></td>
                    <td class="td-value"> {{bankData.defaultingTimeLimit/(60*60)}} hours </td>
                </tr>
            </tbody>
        </table>
    </div>
    <br>
    <hr>

    <h3> All Lenders </h3>
    <div>
         <table >
                <thead>
                    <th> Lender Address </th>
                    <th> Balance </th>
                </thead>
            <tbody v-if="this.lenders.length != 0">
                <tr v-for="(item, index) in lenders" :key="index">
                    <td class="td-heading"> <b>{{item.address}}</b> </td>
                    <td class="td-value"> {{item.balance/1000000000000000000}} XRP</td>
                </tr>
            </tbody>
        </table>
    </div>

    <br>
    <hr>
    <h3> All Customers </h3>
    <table>
        <tbody>
            <thead>
                <th> Customer's Address </th>
            </thead>
            <tr v-for="(item, index) in customers" :key="index" >
                <td class="td-heading"> {{item}} </td>
            </tr>
        </tbody>
    </table>



  </div>
</template>

<script>
// @ is an alias to /src
import blockBorrowerAbi from '../contract_abis/blockborrower_abi.json'
import web3 from 'web3';
import {EVM_SIDECHAIN_RPC, BLOCKBORROWER_CONTRACT_ADDRESS} from '../env/env.js';

export default {
  name: 'PersonalBankView',

  props: {
        bankId:Object
  },

  data() {
      return {
          bankData: {},
          bankStatus: "",
          lenders: [],
          customers: [],
      }
  },

  created() {
      const bankId = this.$route.params.bankId
      /* eslint-disable */
      let bId = BigInt(bankId)
      console.log("In created of personal bank view", bId);
      this.fetchMyBankDetails(bId);
  },

  methods: {

    async fetchMyBankDetails(bankId) {
        if(Number(bankId) == 0 || Number(bankId) == -1) {
            console.log("No bank found")
            this.bankStatus = "No Bank Found !!"
            return;
        }

        await this.getBankDetails(bankId);
        await this.getAllLendersAndBalance(bankId);
        await this.getAllCustomerOfBank(bankId);
    },

    async getAllLendersAndBalance(bankId) {
            console.log("Inside lenders")
           const web3i = new web3(EVM_SIDECHAIN_RPC);
            try {
                const contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                console.log("Here")
                const result = await contract.methods.getAllLendersOfBank(bankId).call({})
                //console.log("Lender's result", result)
                for(var i = 0; i < result.length; i++) {
                    const statusOfLen = await contract.methods.getStatusOfLender(bankId, result[i]).call({});
                    console.log(result[i], "status: ",statusOfLen)
                    if(!statusOfLen) {
                        continue;
                    }
                    const balance = await contract.methods.getMyBankBalance(bankId, result[i]).call({});
                    const data = {
                        "address": result[i],
                        "balance": Number(balance),
                    }
                    this.lenders.push(data);

                    console.log(this.lenders);
                }

            } catch(error) {
                 console.log("Error here", error);
             }

    },

    async getBankDetails(bankId) {
           console.log("Bank Details");
           const web3i = new web3(EVM_SIDECHAIN_RPC);
           try {
                const contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                console.log("Here")
                const result = await contract.methods.bankList(bankId).call({})
                console.log("here")
                let jsonString = ((obj) => JSON.stringify(obj, (key, value) => (typeof value === 'bigint' ? value.toString() : value), 2))(result);

                 this.bankData = JSON.parse(jsonString);
                 //console.log(jsonString)
                 this.bankData.bankId = bankId;
                 if(this.bankData.functional) {
                     this.bankStatus = "Bank is Functional !"
                 } else {
                     this.bankStatus = "Bank is not Functional !"
                 }
           } catch (error) {
               console.log(error);
           }
      },

        async getAllCustomerOfBank(bankId) {
           console.log("Bank Details");
           const web3i = new web3(EVM_SIDECHAIN_RPC);
           try {
                const contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                console.log("Here")
                const result = await contract.methods.getAllCustomersOfBank(bankId).call({})

                 this.customers = result;
                 //console.log(jsonString)
           } catch (error) {
               console.log(error);
           }
      },
  },


}
</script>

<style scoped>
  table {
      border-collapse: collapse;
      width: 100%;
    table-layout: fixed; /* Ensures equal width for all columns */

    }

    th, td {
      border: 1px solid black;
      padding: 8px;
      text-align: left;
    width: 33.33%; /* Equal width for each column */
    overflow-x: scroll;

    }

    th {
      background-color: #f2f2f2;
    }

    td:hover {
      background-color: aliceblue;
    }

    .td-heading {
        background-color: skyblue
    }

    .td-value {
        background-color: lightgreen;

    }
    td:nth-child(2n) td:not(:last-child) {
      border-right: 5px solid black;
    }

    #fn-pr {
        color: red;
    }
</style>
