<template>
  <div id="nav-id">
    <Button class="connect-btn-nav" @click="connectWallet"> <b>{{btnText}}</b> </Button><br><br><hr><br>
    <div id="nav-heading">
        <label class="nav-left-item" id="nav-head-text"> <b>{{dashboardName}}</b> </label>
        <label class="switch nav-right-item">
            <input type="checkbox" id="toggleSwitch" v-model="isChecked" @click="changeDashboardSetting" checked>
            <span class="slider"></span>
        </label>
    </div>
    <br>
    <hr>
    <br>
    <div id="all-rt-lnk-nav">
        <router-link to="/create_bank_view" class="rt-lnk" v-if="isChecked"><b>🏦 Create Bank</b></router-link><hr  v-if="isChecked">
        <router-link to="/join_bank_view" class="rt-lnk" v-if="isChecked"><b>💼 Join Bank</b></router-link><hr  v-if="isChecked">
        <router-link to="/my_bank_view" class="rt-lnk" v-if="isChecked"><b>🪙 My Bank</b></router-link><hr  v-if="isChecked">
                <router-link to="/my_loan_bank_view" class="rt-lnk" v-if="isChecked"><b>💶 Loans Given</b></router-link><hr  v-if="isChecked">


        <router-link to="/take_loan_customer_view" class="rt-lnk" v-if="!isChecked"><b>💰 Take Loan</b></router-link><hr  v-if="!isChecked">
        <router-link to="/my_loan_customer_view" class="rt-lnk" v-if="!isChecked"><b>💲 My Loans</b></router-link><hr  v-if="!isChecked">
        <router-link to="/all_bank_view" class="rt-lnk" ><b>💵 All Banks</b></router-link><hr>
        <router-link to="/auction_list_view" class="rt-lnk" ><b>🎟️ Auction</b></router-link><hr>
        <a href="https://strong-dusk-b787a1.netlify.app/" class="rt-lnk" target="_blank"><b>🌉 EVM-XRPL NFT Bridge</b></a><hr>
        <a href="https://bridge.devnet.xrpl.org/" class="rt-lnk" target="_blank"><b>🖾 EVM-XRPL Bridge</b></a><hr>




    </div>

    <alert-modal :is-visible="isModalVisible" :message="alertMessage" @close="onCloseModal" />





  </div>
</template>

<script>
import AlertModal from './AlertModal.vue';

export default {
    name: "NavigationComponent",

    components: {
        AlertModal,
    },

    data() {
        return {
            dashboardName: "Lender's Dashboard",
            switchIdentifier: 1,
            isChecked: true,
            walletAddress: "",
            btnText: "Connect Metamask",
            evmXrplChainId: "0x15f902",
            isModalVisible: false,
            alertMessage: "",
        }
    },

    mounted() {
        this.determineDashboardTypeOnRefresh()
    },

     created() {
        this.currentChainId();
        this.connectWallet();
        if (window.ethereum) {
            window.ethereum.on('accountsChanged', this.handleAccountsChanged);
            window.ethereum.on('chainChanged', this.handleChainChanged);
        }
    },

    methods: {
        changeDashboardSetting() {
            console.log(this.isChecked);
            if(!this.isChecked) {
                this.dashboardName = "Lender's Dashboard"
                this.switchIdentifier = 1;
                localStorage.setItem("dashboardType", "lender");

            } else {
                this.dashboardName = "Customer's Dashboard"
                this.switchIdentifier = 2;
                localStorage.setItem("dashboardType", "customer");

            }
                            
            this.$router.push("/")

        },

        async connectWallet() {
            console.log("I am inside connect wallet function")
            if (window.ethereum) {
                console.log("I am inside window ether")
                try {
                    const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
                    this.account = accounts[0];
                    this.walletAddress = this.account
                    localStorage.setItem("walletAddress", this.walletAddress);
                    this.setButtonText()
                } catch (error) {
                    console.error('User denied account access');
                    this.showAlert('Please connect to MetaMask.')
                }
            } else {
                console.error("Metamask not existing")
                this.showAlert("MetaMask is not available. Please install and configure MetaMask.")
            }

            console.log("I am done")
        },

        
        // Check if MetaMask is installed and the Ethereum provider is available
        async currentChainId() {
            if (window.ethereum) {
                const ethereum = window.ethereum;
                try {
                    const chainId = await ethereum.request({ method: 'eth_chainId' });
                    this.handleChainChanged(chainId);
                } catch (error) {
                    console.error(`Error getting chainId: ${error}`);
                    this.showAlert(`Error getting chainId: ${error}`)
                }
            } else {
                console.error('MetaMask is not available. Please install and configure MetaMask.');
                this.showAlert('MetaMask is not available. Please install and configure MetaMask.')

            }
        },

        setButtonText() {
            if(this.walletAddress.length != 0) {
                this.btnText = this.walletAddress.substring(0,9) + "..." + this.walletAddress.substring(this.walletAddress.length-7, this.walletAddress.length-1)
            }
        },

        handleChainChanged(chainId) {
            // Handle the chain change
            console.log(chainId);

            if(chainId != this.evmXrplChainId) {
                 this.showAlert("Please change chain to EVM XRPL Sidechain !! ")
            } else {
                this.onCloseModal()
            }
        },

        handleAccountsChanged(accounts) {
            if (accounts.length === 0) {
                console.log('Please connect to MetaMask.');
                this.showAlert('Please connect to MetaMask.')
            } else {
                this.account = accounts[0];
                this.walletAddress = this.account;
                localStorage.setItem("walletAddress", this.walletAddress);
                this.setButtonText();
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

        determineDashboardTypeOnRefresh() {
            try {
                const dType = localStorage.getItem("dashboardType");
                if(dType == "customer") {
                    this.isChecked = false;
                    this.dashboardName = "Customer's Dashboard"
                    this.switchIdentifier = 2;
                    localStorage.setItem("dashboardType", "customer");
                } else if (dType == "lender") {
                    this.isChecked = true;
                    this.dashboardName = "Lender's Dashboard"
                    this.switchIdentifier = 1;
                    localStorage.setItem("dashboardType", "lender");
                } else {
                    this.isChecked = true;
                    this.dashboardName = "Lender's Dashboard"
                    this.switchIdentifier = 1;
                    localStorage.setItem("dashboardType", "lender");
                }
            } catch(err) {
                this.isChecked = true;
                this.dashboardName = "Lender's Dashboard"
                this.switchIdentifier = 1;
                localStorage.setItem("dashboardType", "lender");
                console.log(err);
            }

        }


    }

}

</script>

<style scoped>
#nav-id {
    text-align: left;
    padding: 10px;
}

#nav-heading {
  display: flex;
  justify-content: space-between;
}

.nav-left-item {
    text-align: left;
}

.nav-right-item {
    text-align: right;
}


#switch-btn {
    margin-left: 2px;
}

#nav-head-text {
    font-size: 18px;
    color:green;
}

.connect-btn-nav {
    width: 100%;
    height: 30px;
    border-radius: 5px;
    font-family: monospace;
    cursor: pointer;
    background-image:  linear-gradient(to right, lightgreen, lightseagreen);
}

.connect-btn-nav:hover {
    background-image:  linear-gradient(to left, lightgreen, lightseagreen );
    box-shadow: 0 2px 0 skyblue;
    transform: translateZ(1px);

}

.rt-lnk {
  text-decoration: none;
  display: inline-block;
  padding: 8px; /* Adjust padding as needed */
  color: black; /* Button text color */
  border: none;
  border-radius: 2px; /* Rounded corners */
  cursor: pointer;
  width: 100%;
  margin-bottom: 5px;
transition: transform 0.3s ease; /* Smooth transition effect */

}

.rt-lnk:hover {
  background-color: lightgreen; /* Change color on hover */
  box-shadow: 0 0 2px 0 salmon;
  transform: scale(1.05); /* Add a subtle scale effect */
}

#all-rt-lnk-nav {
    padding-right: 15px;
}










/* CSS FOR SWITCH */
.switch {
  position: relative;
  display: inline-block;
  width: 40px; /* Reduced width */
  height: 20px; /* Reduced height */
}

.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color:blue;
  transition: 0.4s;
  border-radius: 10px; /* Reduced border-radius */
}

.slider:before {
  position: absolute;
  content: "";
  height: 14px; /* Reduced height */
  width: 14px; /* Reduced width */
  left: 3px; /* Adjusted left position */
  bottom: 3px; /* Adjusted bottom position */
  background-color: white;
  transition: 0.4s;
  border-radius: 50%;
}

input:checked + .slider {
  background-color: blueviolet;
}

input:focus + .slider {
  box-shadow: 0 0 1px blueviolet;
}

input:checked + .slider:before {
  transform: translateX(20px); /* Reduced translation */
}

</style>
