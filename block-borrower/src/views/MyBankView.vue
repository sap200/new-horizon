<template>
  <h2> Bank Detail </h2>
  <div class="prnt-div-my-bnk">
      <div id="upper-div">
         <label style="color:blue"> Amount : </label><input type="number" v-model="depositAmount"/>
         <Button class="mn-btn" style="margin-left: 20px;" @click="depositToBank"> Deposit </Button>
        <Button class="mn-btn" :disabled="deactivateClose" v-if="String(this.bankData.admin).toLowerCase() == getMyWalletAddress().toLowerCase()" @click="closeMyBank"> Close Bank (Admin) </Button>
        <Button class="mn-btn" :disabled="deActivateLeaveOrClose" v-if="String(this.bankData.admin).toLowerCase() != getMyWalletAddress().toLowerCase()" @click="leaveBank"> Leave Bank (Lender) </Button>
    </div>
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
                    <td class="td-heading"> <b> Defaulting Time Limit </b></td>
                    <td class="td-value"> {{bankData.defaultingTimeLimit/(60*60)}} hours </td>
                </tr>
            </tbody>
        </table>
    </div>
    <br v-if="this.bankData.functional">
    <hr v-if="this.bankData.functional">
    <div v-if="this.bankData.functional"> 
        <h3> Functionalities </h3>
        <form id="irt-tnr-frm" @submit.prevent>
            <span v-if="String(this.bankData.admin).toLowerCase() == getMyWalletAddress().toLowerCase()">
            <label> <b>Interest Rate :</b> </label> <input type="number" min="1" v-model="interestRate"/> 
            <Button class="btn-intr" @click="setInterestRate"> Change Interest </Button>
            </span>
            <span v-if="String(this.bankData.admin).toLowerCase() == getMyWalletAddress().toLowerCase()">
            <label> <b>Tenure</b> </label> <input type="number" min="1" v-model="tenure"/>
            <Button class="btn-intr" @click="setTenure"> Change Tenure </Button>
            </span>
            <span>
            <label> <b>Withdrawl Amount (XRP)</b> </label> <input type="number" min="1" v-model="withdrawlAmount"/> 
            <Button class="btn-intr" @click="withdrawAmount" :disabled="withdrawButtonDisable"> Withdraw </Button>
            </span>
        </form>
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


    <alert-modal :is-visible="isModalVisible" :message="alertMessage" @close="onCloseModal" />

  </div>
</template>

<script>
// @ is an alias to /src
import blockBorrowerAbi from '../contract_abis/blockborrower_abi.json'
import web3 from 'web3';
import {EVM_SIDECHAIN_RPC, BLOCKBORROWER_CONTRACT_ADDRESS, EVM_SIDECHAIN_BLOCK_EXPLORER} from '../env/env.js';
import {createAlchemyWeb3} from "@alch/alchemy-web3"
import AlertModal from '../components/AlertModal.vue';



export default {
  name: 'MyBankView',

  components: {
      AlertModal
  },

  data() {
      return {
          bankData: {},
          bankStatus: "",
          lenders: [],
          customers: [],
          interestRate: 1,
          tenure: 1,
          evmXrplChainId: "0x15f902",
          isModalVisible: false,
          alertMessage: "",
          isLender: false,
          deActivateLeaveOrClose: true,
          withdrawlAmount: 1,
          withdrawButtonDisable: true,
          depositAmount: 20,
          deactivateClose: true,

      }
  },

  created() {
      this.fetchMyBankDetails();
  },

  methods: {

    async fetchMyBankDetails() {
        const id = await this.getUsersBankIdByAddress();
        if(Number(id) == 0 || Number(id) == -1) {
            console.log("No bank found")
            this.bankStatus = "No Bank Found !!"
            return;
        }

        await this.getBankDetails(id);
        await this.getAllLendersAndBalance(id);
        await this.getAllCustomerOfBank(id);
    },


    getMyWalletAddress() {
        return localStorage.getItem("walletAddress")
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
                    if(Number(balance) > 0 && result[i].toLowerCase() == localStorage.getItem("walletAddress").toLowerCase()) {
                        this.withdrawButtonDisable = false;
                    }
                    
                    if(Number(balance) <= Number(web3.utils.toWei('2', 'ether')) && result[i].toLowerCase() == localStorage.getItem("walletAddress").toLowerCase()) {
                        this.deActivateLeaveOrClose = false;
                    }


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

                 if(this.bankData.totalBalance <= Number(web3.utils.toWei('2', 'ether'))) {
                     this.deactivateClose = false;
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

      async setInterestRate() {
          const chainId = await this.getCurrentChainId();
          if(chainId != this.evmXrplChainId) {
              return;
          }

          const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL
          try {
                    window.contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                    const sender = localStorage.getItem("walletAddress");

                  let txnParams = {
                        to: BLOCKBORROWER_CONTRACT_ADDRESS,
                        from: sender,
                        data: window.contract.methods.setInterestRate(this.bankData.bankId, this.interestRate).encodeABI(),
                    }

                    const txnHash = await window.ethereum.request({
                            method: 'eth_sendTransaction',
                            params: [txnParams],
                    })

                    this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
                    console.log('Transaction hash:', txnHash)
                    setTimeout(() => {
                        this.getTxnReceiptAction(txnHash, "Interest rate set successfully, please reload the page !!")
                    }, 5000);
              } catch(error) {
                    console.log(error);
                    this.showAlert(JSON.stringify(error))
              }
      },

      async setTenure() {
          const chainId = await this.getCurrentChainId();
          if(chainId != this.evmXrplChainId) {
              return;
          }

          const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL
          try {
                    window.contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                    const sender = localStorage.getItem("walletAddress");

                  let txnParams = {
                        to: BLOCKBORROWER_CONTRACT_ADDRESS,
                        from: sender,
                        data: window.contract.methods.setTenure(this.bankData.bankId, this.tenure).encodeABI(),
                    }

                    const txnHash = await window.ethereum.request({
                            method: 'eth_sendTransaction',
                            params: [txnParams],
                    })

                    this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
                    console.log('Transaction hash:', txnHash)
                    setTimeout(() => {
                        this.getTxnReceiptAction(txnHash, "Tenure changed successfully, please reload the page !!")
                    }, 5000);
              } catch(error) {
                    console.log(error);
                    this.showAlert(JSON.stringify(error))
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

        showAlert(alertMsg) {
            this.isModalVisible = true;
            this.alertMessage = alertMsg;
        },
        
        onCloseModal() {
            this.isModalVisible = false;
            this.alertMessage = "";
        },

        async withdrawAmount() {
          const chainId = await this.getCurrentChainId();
          if(chainId != this.evmXrplChainId) {
              return;
          }

          console.log("Inside withdraw amount")

          const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL
          try {
                    window.contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                    const sender = localStorage.getItem("walletAddress");
                    /* eslint-disable */
                    var weiValue = BigInt(web3.utils.toWei(String(this.withdrawlAmount), "ether"))

                  let txnParams = {
                        to: BLOCKBORROWER_CONTRACT_ADDRESS,
                        from: sender,
                        data: window.contract.methods.withdrawBalance(weiValue).encodeABI(),
                    }

                    const txnHash = await window.ethereum.request({
                            method: 'eth_sendTransaction',
                            params: [txnParams],
                    })

                    this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
                    console.log('Transaction hash:', txnHash)
                    setTimeout(() => {
                        this.getTxnReceiptAction(txnHash, "Withdrawl successful, please reload the page and check your balance!!")
                    }, 5000);
              } catch(error) {
                    console.log(error);
                    this.showAlert(JSON.stringify(error))
              }
        },

        async depositToBank() {
            console.log("Inside deposit amount")
            const chainId = await this.getCurrentChainId();
            if(chainId != this.evmXrplChainId) {
                this.showAlert("Chain Id is not evm xrpl chain id ! Please change")
                return;
            }

            console.log("Inside Deposit amount")
            const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL
            try {
                        window.contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                        const sender = localStorage.getItem("walletAddress");
                        /* eslint-disable */
                        var weiValue = Number(web3.utils.toWei(String(this.depositAmount), "ether")).toString(16)


                    let txnParams = {
                            to: BLOCKBORROWER_CONTRACT_ADDRESS,
                            from: sender,
                            value: weiValue,
                            data: window.contract.methods.depositToBank(this.bankData.bankId).encodeABI(),
                        }

                        console.log(txnParams)

                        const txnHash = await window.ethereum.request({
                                method: 'eth_sendTransaction',
                                params: [txnParams],
                        })

                        console.log(txnHash)


                        this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
                        console.log('Transaction hash:', txnHash)
                        setTimeout(() => {
                            this.getTxnReceiptAction(txnHash, "Deposit successful, please reload the page and check your balance!!")
                        }, 5000);
                } catch(error) {
                        console.log(error);
                        this.showAlert(JSON.stringify(error))
                }



        },

        async leaveBank() {
          const chainId = await this.getCurrentChainId();
          if(chainId != this.evmXrplChainId) {
              return;
          }

          console.log("Inside leave Bank")

          const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL
          try {
                    window.contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                    const sender = localStorage.getItem("walletAddress");
                    /* eslint-disable */

                  let txnParams = {
                        to: BLOCKBORROWER_CONTRACT_ADDRESS,
                        from: sender,
                        data: window.contract.methods.LeaveBank().encodeABI(),
                    }

                    const txnHash = await window.ethereum.request({
                            method: 'eth_sendTransaction',
                            params: [txnParams],
                    })

                    this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
                    console.log('Transaction hash:', txnHash)
                    setTimeout(() => {
                        this.getTxnReceiptAction(txnHash, "Bank Left ! Reload the page !!")
                    }, 5000);
              } catch(error) {
                    console.log(error);
                    this.showAlert(JSON.stringify(error))
              }

        },

        async closeMyBank() {
          const chainId = await this.getCurrentChainId();
          if(chainId != this.evmXrplChainId) {
              return;
          }

          console.log("Inside close Bank")

          const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL
          try {
                    window.contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                    const sender = localStorage.getItem("walletAddress");
                    /* eslint-disable */

                  let txnParams = {
                        to: BLOCKBORROWER_CONTRACT_ADDRESS,
                        from: sender,
                        data: window.contract.methods.closeBank(BigInt(this.bankData.bankId)).encodeABI(),
                    }

                    const txnHash = await window.ethereum.request({
                            method: 'eth_sendTransaction',
                            params: [txnParams],
                    })

                    this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
                    console.log('Transaction hash:', txnHash)
                    setTimeout(() => {
                        this.getTxnReceiptAction(txnHash, "Bank Permanently shutdown ! Reload the page !!")
                    }, 5000);
              } catch(error) {
                    console.log(error);
                    this.showAlert(JSON.stringify(error))
              }

        }
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

    .btn-intr {
        margin-right: 25px;
        margin-left: 10px;
        border-radius: 5px;
        background-image: linear-gradient(to right, skyblue, aliceblue);
        cursor: pointer;
        height: 30px;
        font-family: monospace;
    }

    .btn-intr:hover:enabled {

        background-image: linear-gradient(to left, skyblue, aliceblue);

    }

    
    .btn-intr:disabled {

        background-image: linear-gradient(to left, grey, lightgrey);

    }

    #upper-div {
        text-align: right;
    }

    .mn-btn {
        font-family: monospace;
        margin-right: 25px;
        padding: 5px;
        background-image: linear-gradient(to right, salmon, red);
        cursor: pointer;
        border-radius: 5px;
        color: white;
    }

    .mn-btn:hover:enabled {

        background-image: linear-gradient(to left, salmon, red);

    }

    .mn-btn:disabled {
        background-image: linear-gradient(to left, grey, lightgrey);
        color: black;
    }
</style>
