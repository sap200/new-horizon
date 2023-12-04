<template>
<div>
    <h2> <u>Create a Bank</u> </h2>
    <h4 class="ins-cr-bn"> <b>IMPORTANT !! This Txn will cost you 450 XRP (MIN AMOUNT TO CREATE BANK)!!</b> </h4>

    <div id="cr-bnk-div"> 
        <form id="cr-bnk-frm" @submit.prevent>
            <p class="frm-p-nm"> <b>Bank Name</b> </p>
            <input type="text"  class="styled-input" v-model="bankName"/>
            <p class="frm-p-nm"> <b>Interest Rate</b> </p>
            <input type="number" min="1" class="styled-input" v-model="interestRate"/>
            <p class="frm-p-nm"> <b>Tenure</b> </p>
            <input type="number" min="1" class="styled-input" v-model="tenure"/>
            <p class="frm-p-nm"> <b>Loan Percentage from original value</b> </p>
            <input type="number" min="10" class="styled-input" v-model="loanPercentage"/>
            <p class="frm-p-nm"> <b>Number of Defaults after which auction can start</b> </p>
            <input type="number" min="2" class="styled-input" v-model="defaultingLimit"/>
            <p class="frm-p-nm"> <b>Time Limit for full payment after Loan closure, if exceeded you can go for auction (In days) </b> </p>
            <input type="number" min="1" class="styled-input" v-model="defaultingTimeLimitInDays"/>
            <br><br>
            <div id="btn-ctr">
                <Button id="cr-frm-btn" @click="createBank" :disabled="createButtonDisable || this.bankName == ''"> <b>Create</b> </Button>
            </div>
        </form>
    </div>

    <alert-modal :is-visible="isModalVisible" :message="alertMessage" @close="onCloseModal" />
</div>
    
</template>

<script>
import blockBorrowerAbi from '../contract_abis/blockborrower_abi.json'
import web3 from 'web3';
import {EVM_SIDECHAIN_RPC, BLOCKBORROWER_CONTRACT_ADDRESS, EVM_SIDECHAIN_BLOCK_EXPLORER, MIN_VAL_CREATE_BANK} from '../env/env.js';
import AlertModal from '../components/AlertModal.vue';
import {createAlchemyWeb3} from "@alch/alchemy-web3"



export default {
    name: "CreateBankView",

    components: {
        AlertModal,
    },

    data() {
        return {
            bankName: "",
            interestRate: 1,
            tenure: 1,
            loanPercentage: 10,
            defaultingLimit: 2,
            defaultingTimeLimitInDays: 1,
            isModalVisible: false,
            alertMessage: "",
            createButtonDisable: false,
            evmXrplChainId: "0x15f902",

        }
    },

    created() {
         this.getBankIdForUser(localStorage.getItem("walletAddress"))
    },

    methods: {

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

        async createBank() {
            const chainId = await this.getCurrentChainId();
            if(chainId != this.evmXrplChainId) {
                this.showAlert("Switch to EVM XRPL Sidechain");
                console.log(chainId);
                return;
            }

            const walletAddress = localStorage.getItem("walletAddress")
            const bankId = await this.getBankIdForUser(walletAddress)
            const numBankId = Number(bankId)
            console.log(numBankId);

            if(numBankId == 0) {
                this.createButtonDisable = true;
                const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL
                try {
                    window.contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                    const sender = localStorage.getItem("walletAddress");

                  var weiValue = Number(web3.utils.toWei(MIN_VAL_CREATE_BANK, "ether")).toString(16);
                  console.log(weiValue)
                  let txnParams = {
                        to: BLOCKBORROWER_CONTRACT_ADDRESS,
                        from: sender,
                        value: weiValue,
                        data: window.contract.methods.createBank(this.bankName, this.interestRate, this.tenure, this.loanPercentage, this.defaultingLimit, this.defaultingTimeLimitInDays*86400).encodeABI(),
                    }

                    const txnHash = await window.ethereum.request({
                            method: 'eth_sendTransaction',
                            params: [txnParams],
                    })

                    this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_SIDECHAIN_BLOCK_EXPLORER + txnHash)
                    console.log('Transaction hash:', txnHash)
                    setTimeout(() => {
                        this.getTxnReceiptAction(txnHash, "Bank creation successful, Proceed to my banks to see your bank !!")
                    }, 5000);

                } catch (error) {
                    this.createButtonDisable = false;
                    this.showAlert(JSON.stringify(error))
                    console.error('Error:', error);
                }

            } else if (numBankId == -1) {
                this.showAlert("Unable to make a call to smart contract !");
            } else {
                this.showAlert("You already have a bank !");
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

        async getBankIdForUser(addr) {
            console.log("inside get bank for user id");
            const web3i = new web3(EVM_SIDECHAIN_RPC);
            try {
                const contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                const result = await contract.methods.lenderAddressToBankMapping(addr).call({})
                console.log(result);
                if(result != 0) {
                    this.createButtonDisable = true;
                } else {
                    this.createButtonDisable = false;
                }
                return result;
            } catch (error) {
                console.log(error);
                return -1;
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
    }
}
</script>

<style scoped>
#cr-bnk-div {
    text-align: center;
}

#cr-bnk-frm {
    display: inline-block;
    width: 500px;
    text-align: left;
}
.frm-p-nm {
    color: darkred;
}

/* Reset some default styles for the input */
input {
    border: none;
    border-bottom: 1px solid black; /* Set the color and thickness of the bottom line */
    outline: none;
    background-color: transparent;
    font-family: monospace;
}

/* Style the bottom line */
.styled-input {
    position: relative;
    width: 100%;
}

.styled-input::before {
    content: "";
    position: absolute;
    bottom: 0;
    left: 0;
    width: 100%;
    height: 2px; /* Adjust the height of the line as needed */
    background-color: #333; /* Color of the line */
    transition: background-color 0.3s ease; /* Add a transition effect */
}

.styled-input input:focus::before {
    background-color: #007BFF; /* Change the color when the input is focused */
}

#btn-ctr {
    text-align: center;
}

#cr-frm-btn {
    width: 50%;
    margin-top: 20px;
    height: 30px;
    border-radius: 5px;
    background-image: linear-gradient(to right, aliceblue, skyblue, pink);
    font-family: monospace;
    color: black;
    cursor: pointer;
}

#cr-frm-btn:hover:enabled {
    background-image: linear-gradient(to left, aliceblue, skyblue, pink);
    box-shadow: 0 0 0 5px skyblue;

}

#cr-frm-btn:disabled {
    background-image: linear-gradient(to left, grey, lightgrey);

}

.ins-cr-bn {
    color: red;
    text-align: center;
}
</style>
