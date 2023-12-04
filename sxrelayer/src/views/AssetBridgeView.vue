<template>
    <div id="parent-div">
    <p id="ins"> If you transferred asset cross-chain, please wait here to until you get the notification for the success of the transaction </p>
    <div class="menu-bar">
        <Button id="nft-mint-btn" @click="takeToMint"> <b>Mint NFT</b></Button>
        <Button id="lm-view" @click="routeToView"> <b> Locker View </b> </Button>
        <Button id="switch" @click="changeNetwork"> <b> Switch </b> </Button>
        <Button id="metamask-button" @click="connectWallet"> <b>{{btnText}}</b> </Button>
    </div>
        <div class="heading">
            <h2> <span id="from">{{from}}</span> ⮀ <span id="to">{{to}}</span> ASSET BRIDGE </h2>
        </div>

        <div class="forms" v-if="this.switch==1">
            <form class="sepolia-xrpl" @submit.prevent>

               <p class="item"> <b> Provide <span id="approv-col">Approval </span> </b> </p>
                <p class="item normal"> NFT's Contract Address </p> <input type="text" v-model="nftContractAddressForApproval"/> 
               <p class="item normal"> Token Id </p> <input type="number" min="0" v-model="nftIdForApproval"/> 
               <br><br>
               <hr>
               <Button class="frm-btn" @click="approveLockerContract" :disabled="this.nftContractAddressForApproval == '' || this.nftIdForApproval == 0"> Approve </Button>
               <br><br>
               <hr>
               <p class="item space-from"> <b> From <span class="fm-from"> {{from}} </span> </b> </p>
               <p class="item normal"> NFT's Contract Address </p> <input type="text" v-model="nftContractAddressForSendingToXRPL" />
               <p class="item normal"> Token Id </p> <input type="number" min="0" v-model="nftIdForSendingToXRPL"/>
               <br><br>
               <hr>
               <p class="item space-from"> <b>To <span class="fm-to"> {{to}} </span></b> </p>
               <p class="item normal"> Receipent's Address EVM_XRPL Side Chain </p> <input type="text" v-model="receipentAddressInDestinationChain"/>
               <br><br>
               <hr>
               <div id="cnt-dv">
                <Button class="frm-btn" @click="transferNFTToEVMCosmosSidechain" :disabled="this.nftContractAddressForSendingToXRPL == '' || this.nftIdForSendingToXRPL == 0 || this.receipentAddressInDestinationChain == ''"> Transfer </Button>
               </div>
            </form>
        </div> 

        <div class="forms" v-if="this.switch==2">
            <form  class="sepolia-xrpl fm-z" @submit.prevent>

               <p class="item space-from"> <b> From <span class="fm-to"> {{from}} </span> </b> </p>
               <p class="item normal"> NFT's Contract Address </p> <input type="text" v-model="nftContractAddressForSendingToSepolia"/>
               <p class="item normal"> Token Id </p> <input type="number" min="0" v-model="nftIdForSendingToSepolia"/>
               <br><br>
               <hr>
               <p class="item space-from"> <b>To <span class="fm-from"> {{to}} </span></b> </p>
               <p class="item normal"> Receiver Address in Sepolia Chain </p> <input type="text" v-model="receipentAddressInSepoliaChain"/>
               <br><br>
               <hr>
               <div id="cnt-dv">
                <Button class="frm-btn" @click="transferFromXrplToSepolia" :disabled="this.nftContractAddressForSendingToSepolia == '' || this.nftIdForSendingToSepolia == 0 || this.receipentAddressInSepoliaChain == ''"> Transfer </Button>
               </div>
            </form>
        </div> 
        <br><br>
        <hr v-if="this.switch==1">
        <h2 id="wt-head" v-if="this.switch==1"> Withdraw Asset </h2>

        <div class="forms" v-if="this.switch==1">
            <form  class="sepolia-xrpl fm-wt" @submit.prevent>
                <p class="item"> <b> Asset <span id="withdraw-col">Withdraw </span> </b> </p>
                <p class="item normal"> NFT's Contract Address: </p> <input type="text" v-model="nftContractAddressToWithdraw"/> 
                <p class="item normal"> Token Id </p> <input type="number" min="0" v-model="nftTokenIdToWithdraw"/> 
                <br><br>
                <hr>
               <Button class="frm-btn" @click="withdrawNFT" :disabled="this.nftContractAddressToWithdraw == '' || this.nftTokenIdToWithdraw == 0"> Withdraw </Button>
            </form>

        </div>
        <br><br>
        <hr>
        <div id="footer">
           <i> &copy; Created by zed !! 😉 Just Kidding, Feel free to copy the site and violate the copyright 🙂 !! </i>
        </div>

         <alert-modal :is-visible="isModalVisible" :message="alertMessage" @close="onCloseModal" :noCloseButton="noCloseButton"/>

    </div>
</template>

<script>
import AlertModal from '../components/AlertModal.vue';
import {createAlchemyWeb3} from "@alch/alchemy-web3"
import contractABI from '../contract_abis/nft_abi.json'; 
import lockerContractAbi from '../contract_abis/locker_abi.json'
import minterContractAbi from '../contract_abis/minter_abi.json'
import {ETHEREUM_SEPOLIA_RPC, LOCKER_CONTRACT_ADDRESS, SEPOLIA_BLOCK_EXPLORER, EVM_XRPL_CHAIN_NAME, MINTER_CONTRACT_ADDRESS, ETHEREUM_XRPL_RPC, EVM_XRPL_BLOCK_EXPLORER, EVENT_LISTENER_URL} from '../env/env'


export default {
    name: "AssetBridgeView", 
    components: {
        AlertModal,
    },

    data() {
        return {
            to: "EVM_XRPL",
            from: "SEPOLIA",
            btnText: "Connect Your wallet",
            walletAddress: "",
            switch: 1,
            sepoliaChainId: "0xaa36a7",
            evmXrplChainId: "0x15f902",
            alertMessage: "",
            isModalVisible: false,
            nftContractAddressForApproval: "",
            nftIdForApproval: 0,
            nftIdForSendingToXRPL: 0,
            nftContractAddressForSendingToXRPL: "",
            receipentAddressInDestinationChain: "",
            nftContractAddressForSendingToSepolia: "",
            nftIdForSendingToSepolia: 0,
            receipentAddressInSepoliaChain: "",
            nftContractAddressToWithdraw: "",
            nftTokenIdToWithdraw: 0,
            noCloseButton: false,
        }
    },

    created() {

        this.connectWallet();
        this.currentChainId();
        if (window.ethereum) {
            window.ethereum.on('accountsChanged', this.handleAccountsChanged);
            window.ethereum.on('chainChanged', this.handleChainChanged);
        }
        this.listenForEvents()
    },

    methods: {
        changeNetwork() {
            if(this.switch==1) {
                this.switch=2; 
                this.from = "EVM_XRPL"
                this.to = "SEPOLIA"
            } else {
                this.switch=1
                this.from = "SEPOLOIA"
                this.to = "EVM_XRPL"
            }
            
            this.currentChainId();

        },

        showAlert(alertMsg, shouldShowCloseButton) {
            this.isModalVisible = true;
            this.noCloseButton = shouldShowCloseButton;
            this.alertMessage = alertMsg;
        },

        onCloseModal() {
            this.isModalVisible = false;
            this.alertMessage = "";
        },

        async connectWallet() {
            if (window.ethereum) {
                try {
                    const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
                    this.account = accounts[0];
                    this.walletAddress = this.account
                    this.setButtonText()
                } catch (error) {
                    console.error('User denied account access');
                    this.showAlert('Please connect to MetaMask.', false);
                }
            } else {
                console.error("Metamask not existing")
                this.showAlert("MetaMask is not available. Please install and configure MetaMask.", false);
            }
        },

        setButtonText() {
            if(this.walletAddress.length != 0) {
                this.btnText = this.walletAddress.substring(0,9) + "..." + this.walletAddress.substring(this.walletAddress.length-7, this.walletAddress.length-1)
            }
        },

        handleAccountsChanged(accounts) {
            if (accounts.length === 0) {
                console.log('Please connect to MetaMask.');
                this.showAlert('Please connect to MetaMask.', false);
            } else {
                this.account = accounts[0];
                this.walletAddress = this.account;
                this.setButtonText();
            }
        },

        takeToMint() {
            window.open("https://waiting-pot.surge.sh/", "_blank");
        },

        routeToView() {
            this.$router.push("/lmview/"+this.walletAddress)
        },

        handleChainChanged(chainId) {
        // Handle the chain change
            console.log(chainId);

            if(this.switch == 1) {
                if(chainId != this.sepoliaChainId) {
                    this.showAlert("Please change chain to Sepolia !! Else transfer won't work", false)
                }
            } else {
                if(chainId != this.evmXrplChainId) {
                    this.showAlert("Please change chain to EVM XRPL Sidechain !! ", false)
                }

            }
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
                    this.showAlert(`Error getting chainId: ${error}`, false)
                }
            } else {
                console.error('MetaMask is not available. Please install and configure MetaMask.');
                this.showAlert('MetaMask is not available. Please install and configure MetaMask.', false)

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
                    this.showAlert(`Error getting chainId: ${error}`, false)
                    return "0"
                }
            } else {
                console.error('MetaMask is not available. Please install and configure MetaMask.');
                this.showAlert('MetaMask is not available. Please install and configure MetaMask.', false)
                return "0"
            }
        },



        // ---------------------- WEB3 Functions --------------------------------------------------------------------------------

        async approveLockerContract() {
            const chainId = await this.getCurrentChainId()
            if(chainId != this.sepoliaChainId) {
                this.showAlert("Please switch to sepolia chain Id", false);
                console.log("Invalid chainId , its not sepolia");
                return;
            }
            const web3i = createAlchemyWeb3(ETHEREUM_SEPOLIA_RPC); // Replace with your Ethereum node URL
            try {
                window.contract = new web3i.eth.Contract(contractABI, this.nftContractAddressForApproval);
                // Account with unlocked wallet in MetaMask or other provider
                const sender = this.walletAddress
                // make a txn
                let txnParams = {
                    to: this.nftContractAddressForApproval,
                    from: sender,
                    data: window.contract.methods.approve(LOCKER_CONTRACT_ADDRESS, this.nftIdForApproval).encodeABI(),
                }

                const txnHash = await window.ethereum.request({
                        method: 'eth_sendTransaction',
                        params: [txnParams],
                })

                this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on sepolia block explorer : " + SEPOLIA_BLOCK_EXPLORER + txnHash, false)
                console.log('Transaction hash:', txnHash);
        } catch (error) {
                this.showAlert(JSON.stringify(error), false)
                console.error('Error:', error);
        }

        },

        async transferNFTToEVMCosmosSidechain() {
            const chainId = await this.getCurrentChainId()
            if(chainId != this.sepoliaChainId) {
                this.showAlert("Please switch to sepolia chain Id");
                console.log("Invalid chainId , its not sepolia");
                return;
            }

            const web3i = createAlchemyWeb3(ETHEREUM_SEPOLIA_RPC); // Replace with your Ethereum node URL

            try {
                window.contract = new web3i.eth.Contract(lockerContractAbi, LOCKER_CONTRACT_ADDRESS);
                // Account with unlocked wallet in MetaMask or other provider
                const sender = this.walletAddress
                // make a txn
                let txnParams = {
                    to: LOCKER_CONTRACT_ADDRESS,
                    from: sender,
                    data: window.contract.methods.sendNFT(this.nftIdForSendingToXRPL, this.nftContractAddressForSendingToXRPL, this.receipentAddressInDestinationChain, EVM_XRPL_CHAIN_NAME).encodeABI(),
                }

                const txnHash = await window.ethereum.request({
                        method: 'eth_sendTransaction',
                        params: [txnParams],
                })

                this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on sepolia block explorer : " + SEPOLIA_BLOCK_EXPLORER + txnHash, false)
                console.log('Transaction hash:', txnHash);
                this.showAlert("Transaction is in progress...", true);
                // Example using Web3.js


            } catch (error) {
                this.showAlert(JSON.stringify(error), false)
                console.error('Error:', error);
            }

        },

        async transferFromXrplToSepolia() {
            const chainId = await this.getCurrentChainId()
            if(chainId != this.evmXrplChainId) {
                this.showAlert("Please switch to EVM XRPL side chain Id");
                console.log("Invalid chainId , its not EVM XRPL side chain");
                return;
            }
            
             const web3i = createAlchemyWeb3(ETHEREUM_XRPL_RPC); // Replace with your Ethereum node URL
             try {
                window.contract = new web3i.eth.Contract(minterContractAbi, MINTER_CONTRACT_ADDRESS);
                // Account with unlocked wallet in MetaMask or other provider
                const sender = this.walletAddress
                // make a txn
                let txnParams = {
                    to: MINTER_CONTRACT_ADDRESS,
                    from: sender,
                    data: window.contract.methods.sendBackNFT(this.nftContractAddressForSendingToSepolia, this.nftIdForSendingToSepolia, this.receipentAddressInSepoliaChain).encodeABI(),
                }

                const txnHash = await window.ethereum.request({
                        method: 'eth_sendTransaction',
                        params: [txnParams],
                })

                this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on EVM_XRPL block explorer : " + EVM_XRPL_BLOCK_EXPLORER + txnHash, false)
                console.log('Transaction hash:', txnHash);
                // Example using Web3.js
                this.showAlert("Transaction is in progress...", true);

            } catch (error) {
                this.showAlert(JSON.stringify(error))
                console.error('Error:', error);
            }

        },

        async withdrawNFT() {
            const chainId = await this.getCurrentChainId()
            if(chainId != this.sepoliaChainId) {
                this.showAlert("Please switch to sepolia chain Id");
                console.log("Invalid chainId , its not sepolia");
                return;
            }

            const web3i = createAlchemyWeb3(ETHEREUM_SEPOLIA_RPC); // Replace with your Ethereum node URL

            try {
                window.contract = new web3i.eth.Contract(lockerContractAbi, LOCKER_CONTRACT_ADDRESS);
                // Account with unlocked wallet in MetaMask or other provider
                const sender = this.walletAddress
                // make a txn
                let txnParams = {
                    to: LOCKER_CONTRACT_ADDRESS,
                    from: sender,
                    data: window.contract.methods.withdrawNFT(this.nftContractAddressToWithdraw, this.nftTokenIdToWithdraw).encodeABI(),
                }

                const txnHash = await window.ethereum.request({
                        method: 'eth_sendTransaction',
                        params: [txnParams],
                })

                this.showAlert("Txn Hash is: " + txnHash + ". Please check it out on sepolia block explorer : " + SEPOLIA_BLOCK_EXPLORER + txnHash)
                console.log('Transaction hash:', txnHash);
                // Example using Web3.js
                this.showAlert("Transaction is processing ...", true);
            } catch (error) {
                this.showAlert(JSON.stringify(error), false)
                console.error('Error:', error);
            }

        },

        listenForEvents() {
            const web3i = createAlchemyWeb3(EVENT_LISTENER_URL); // Replace with your Ethereum node URL
            const contract = new web3i.eth.Contract(lockerContractAbi, LOCKER_CONTRACT_ADDRESS);

            // Replace 'YourEventName' with the actual name of the event you want to listen to
            const event = contract.events.LockedNFT();
            const unlockEvent = contract.events.UnlockedNFT();
            const withdrawEvent = contract.events.WithdrawnNFT();

            // Listen for events
            event.on('data', (eventData) => {
                this.showAlert("✅ Transfer to EVM_XRPL Chain successful. ", false)
                console.log('Event data:', eventData.returnValues);
            })
            .on('error', (error) => {
                this.showAlert(JSON.stringify(error), false)
                console.error('Error:', error);
            });

            unlockEvent.on('data', (eventData) => {
                this.showAlert("✅ Your NFT is unlocked please proceed to withdraw", false)
                console.log('Event data:', eventData.returnValues);
            })
            .on('error', (error) => {
                this.showAlert(JSON.stringify(error), false)
                console.error('Error:', error);
            });

            withdrawEvent.on('data', (eventData) => {
                this.showAlert("✅ NFT withdrawn, proceed to opensea to validate !!", false)
                console.log('Event data:', eventData.returnValues);
            })
            .on('error', (error) => {
                this.showAlert(JSON.stringify(error), false)
                console.error('Error:', error);
            });
        }

    }
}

</script>

<style scoped>
#parent-div {
    background-image: linear-gradient(-45deg,pink, black,black, black, skyblue);
    width: 100%;
    font-family:monospace;
    overflow:auto;
    padding-bottom: 20px;
    padding-top: 20px;
    height: 100vh;
}

#ins {
    color:violet;
    background-color: black;
    font-size: 15px;
    padding: 5px;
}

#from {
    margin-right: 10px;
    color: pink;
}

#to {
    margin-left: 10px;
    color: palegreen;
}

.heading {
    color: white;
    padding: 2px;
    text-align: center; /* Center text horizontally within container */
}

#wt-head {
    color:rgb(237, 175, 237);
    padding: 2px;
    text-align: center; /* Center text horizontally within container */
    margin-top: 50px;
}
.menu-bar {
    text-align: right;
    padding-right: 20px;
}

#metamask-button {
    height: 40px;
    padding: 5px;
    width: 200px;
    border-radius: 5px;
    color:black;
    background-image: linear-gradient(to right, skyblue, aliceblue);
    cursor: pointer;
    font-family: monospace;

}

#nft-mint-btn {
    height: 40px;
    padding: 5px;
    width: 120px;
    border-radius: 5px;
    color:black;
    background-image: linear-gradient(to right, skyblue, aliceblue);
    margin-right: 15px;
    cursor: pointer;
    font-family: monospace;
}

#lm-view {
    height: 40px;
    padding: 5px;
    width: 120px;
    border-radius: 5px;
    color:black;
    background-image: linear-gradient(to right, skyblue, aliceblue);
    margin-right: 15px;
    cursor: pointer;
    font-family: monospace;
}

#nft-withdraw {
    height: 40px;
    padding: 5px;
    width: 120px;
    border-radius: 5px;
    color:black;
    background-image: linear-gradient(to right, skyblue, aliceblue);
    margin-right: 15px;
    cursor: pointer;
    font-family: monospace;
}

#metamask-button:hover:enabled { 
    background-image: linear-gradient(to left, skyblue, aliceblue);
    box-shadow: 1px 1px 10px pink; /* Add the box shadow effect on hover */

}

#switch:hover:enabled { 
    background-image: linear-gradient(to left, skyblue, aliceblue);
    box-shadow: 1px 1px 10px pink; /* Add the box shadow effect on hover */

}

#nft-mint-btn:hover:enabled { 
    background-image: linear-gradient(to left, skyblue, aliceblue);
    box-shadow: 1px 1px 10px pink; /* Add the box shadow effect on hover */

}

#nft-withdraw:hover:enabled { 
    background-image: linear-gradient(to left, skyblue, aliceblue);
    box-shadow: 1px 1px 10px pink; /* Add the box shadow effect on hover */

}


#lm-view:hover:enabled { 
    background-image: linear-gradient(to left, skyblue, aliceblue);
    box-shadow: 1px 1px 10px pink; /* Add the box shadow effect on hover */

}


#switch {
    height: 40px;
    padding: 5px;
    width: 120px;
    border-radius: 5px;
    color:black;
    background-image: linear-gradient(to right, skyblue, aliceblue);
    margin-left: auto;
    margin-right: 15px;
    cursor: pointer;
    font-family: monospace;

}

.frm-btn {
    height: 40px;
    padding: 5px;
    width: 120px;
    color:black;
    background-image: linear-gradient(to right, skyblue, aliceblue);
    cursor: pointer;
    margin-top: 15px;
    font-family: monospace;
    width: 100%;
    border-radius: 5px;
}

.frm-btn:hover:enabled {
    background-image: linear-gradient(to left, skyblue, aliceblue);
    box-shadow: 1px 1px 10px pink; /* Add the box shadow effect on hover */
}

.frm-btn:disabled {
    background-image: linear-gradient(to left, grey, white);

}

.forms {
    display: flex;
    justify-content: center;
    
}

.sepolia-xrpl {
    background-color: rgba(51, 121, 151, 0.3); /* Your desired background color */
    padding: 20px;
    border-radius: 5px;
    text-align: left;
    width: 50%;
    height: 740px;
    margin-top: 40px;
    box-shadow: 0 7px 5px 5px rgba(99, 119, 146, 0.5); 
}

#footer {
    color: white;
    text-align: center;
    margin-top: 40px;
    color:greenyellow;
}

.item {
    color: white;
    font-size: 18px;
}

input {
    width: 100%;
    box-sizing: border-box; /* Include padding and border in the element's width */
    padding: 8px; /* Example padding */
    border: 1px solid #ccc; /* Example border */
    margin: 0; /* Remove default margin */
    font-family: monospace;
    border-radius: 5px;
}

.normal {
    font-size: 13px;
}

.fm-from {
    color: pink;
}

.fm-to {
    color: palegreen;

}

#cnt-dv {
    text-align: center;
}

.fm-z {
    height: 440px;
}

#approv-col {
    color:salmon
}

#withdraw-col {
    color: salmon
}

.fm-wt {
    height: 295px;

}
</style>