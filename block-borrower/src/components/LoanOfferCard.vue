<template>
  <div class="card-component">
    <h3>🏦 {{data.bankId}}. {{data.bankName.trim() }} </h3>
    <h4>🏧 Loan Offer:  {{this.getLoanAmount(data)}} XRP</h4>
    <h4>💲 Repayment Amount: {{this.getRepaymentAmount(data)}} XRP</h4>
    <h4>📝 EMI : {{this.getEMI(data)}} XRP </h4>
    <h4>🛑 Defaulting Limit : {{data.defaultingLimit}} </h4>
    <h4>⌛ Defaulting Time Limit : {{data.defaultingTimeLimit}} </h4>
    <h4>❗ Loan Percentage : {{data.loanPercentage}} </h4>
    <h4>📈 ROI : {{data.interestRate}} </h4>
    <h4>📅 Tenure : {{data.tenure}} </h4>
    <h4>👨🏻‍💼 Admin : {{data.admin}} </h4>
    <!-- Add more content or customize as needed -->
  </div>
</template>

<script>
import web3 from 'web3';
import blockBorrowerAbi from '../contract_abis/blockborrower_abi.json'
import {EVM_SIDECHAIN_RPC, BLOCKBORROWER_CONTRACT_ADDRESS, MIN_VAL_JOIN_BANK} from '../env/env.js';
import {createAlchemyWeb3} from "@alch/alchemy-web3"

export default {
  name: "LoanOfferCard",

  props: {
    data: {
      type: Object,
      required: true,
    },
  },

  data() {
      return {
        evmXrplChainId: "0x15f902",
        chainMsg: '',
        clickedJoin: false,

      }
  },

  created() {
      console.log("Inside Loan offer card")
  },

  methods: {
      getLoanAmount(data) {
          console.log(data)
          var lnprcnt = Number(data.nftPrice) * Number(data.loanPercentage) / 100;
          return this.getInXRP(String(lnprcnt));
      },

      getRepaymentAmount(data) {
        var lnprcnt = Number(data.nftPrice) * Number(data.loanPercentage) / 100;
        let repayment = lnprcnt + (lnprcnt*Number(data.interestRate)*Number(data.tenure)/100)
        return this.getInXRP(String(repayment))
      },

      getEMI(data) {
        var lnprcnt = Number(data.nftPrice) * Number(data.loanPercentage) / 100;
        let repayment = lnprcnt + (lnprcnt*Number(data.interestRate)*Number(data.tenure)/100)
        const emi = Number(repayment/ (Number(data.tenure)*12))
        return this.getInXRP(String(emi));
},

      getFunctionalEmoji(functional) {
          if(functional) {
              return "✅"
          } else {
              return "🔴"
          }
      },

      getInXRP(amt) {
          return web3.utils.fromWei(amt, "ether")
      },

      async joinBank() {
          const chainId = await this.getCurrentChainId()
          if(chainId != this.evmXrplChainId) {
            this.chainMsg = "Switch to EVM XRPL chain";
            return;
          } else {
              this.chainMsg = ""
          }

            console.log(this.data.bankId);

            const web3i = new createAlchemyWeb3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL
            try {
                window.contract = new web3i.eth.Contract(blockBorrowerAbi, BLOCKBORROWER_CONTRACT_ADDRESS);
                const sender = localStorage.getItem("walletAddress");

                var weiValue = Number(web3.utils.toWei(MIN_VAL_JOIN_BANK, "ether")).toString(16);
                console.log(weiValue)
                let txnParams = {
                    to: BLOCKBORROWER_CONTRACT_ADDRESS,
                    from: sender,
                    value: weiValue,
                    data: window.contract.methods.joinBank(this.data.bankId).encodeABI(),
                }

                const txnHash = await window.ethereum.request({
                        method: 'eth_sendTransaction',
                        params: [txnParams],
                })

                this.chainMsg = "Txn Hash: " + txnHash;
                console.log('Transaction hash:', txnHash)
                setTimeout(() => {
                    this.getTxnReceiptAction(txnHash, "Bank Joined successfully !!")
                }, 5000);
            } catch (error) {
                console.log(error);
                this.clickedJoin = false;
                this.chainMsg = "Unable to Join Bank !"
            }
      },

      
        async getTxnReceiptAction(txnHash, message) {
            try {
                const web3i = new web3(EVM_SIDECHAIN_RPC); // Replace with your Ethereum node URL

                const receipt = await web3i.eth.getTransactionReceipt(txnHash);
                if (receipt && receipt.status) {
                    this.chainMsg = message;
                    console.log(receipt)
                } else {
                    this.chainMsg = "Txn Failed: " + txnHash;
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
  }
};
</script>

<style scoped>
.card-component {
  border: 1px solid #ccc;
  padding: 20px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  margin-bottom: 20px;
  width: 100%; /* Adjust the width based on your design */
  box-sizing: border-box;
  text-align: left;
}

h3 {
    text-align: center;
    color: darkorchid;
}

p {
    color: darkslategrey;
}

.ttl-bnk-jn {
    color: green;
}

#card-jn-btn {
    width: 100px;
    height: 30px;
    border-radius: 5px;
    background-image: linear-gradient(to right, gold, rgb(194, 183, 123));
    font-family: monospace;
    cursor: pointer;
}

#card-jn-btn:hover:enabled {
    background-image: linear-gradient(to left, gold, rgb(194, 183, 123));

}

#card-jn-btn:disabled {
    background-image: linear-gradient(to left, grey, lightgrey);
    cursor:auto;

}

#ftm-brn-crd {
    width: 100%;
    text-align:right;
}

#err-msg-crd {
    text-align: center;
    color: red;
}

</style>
