<template>
    <div id="p-cont"> 
        <h2> Locker View </h2>
        <hr>

        <Button id="flexi-de-btn" @click="getAllTokensUnderUser"> View All My Contract Lists </Button>
       <hr> 
        <h3> CONTRACT,TOKENID => DETAILS MAP </h3>
        <p id="ins"> To get all tokens under user in the SEPOLIA CHAIN, select SEPOLIA CHAIN and put NFT contract Address and pass TokenId as 0 ! Also make sure to switch network when querying. </p>
        <div class="ox-flexi">
            <form @submit.prevent>
                <label class="itm-txt"> NFT Contract Address : </label>
                <select id="sel-id" v-model="commandType">
                    <option value="0"> SEPOLIA CHAIN </option>
                    <option value="1"> EVM XRPL CHAIN </option>
                </select>

                <label class="itm-txt"> NFT Contract Address : </label>
                <input type="text" id="ext-id" v-model="nftContractAddress"/>

                <label class="itm-txt"> Token Id in SEPIOLA CHAIN : </label>
                <input type="number" id="num-id" min="0" v-model="tokenId"/>

                <Button id="fm-btn" @click="takeDecision" :disabled="this.tokenId == 0 && this.nftContractAddress == ''"> Find </Button>
            </form>
        </div>
    <hr>
    <br>
        <p v-if="resultObj.length == 0"> EMPTY LIST RETURNED </p>
        <div id="tbl">
            <table>
                <tbody>
                <tr  v-for="(item, index) in resultObj" :key="index">
                    <td>
                        <pre>{{item}}</pre>
                    </td>
                </tr>
                </tbody>
            </table>
        </div>
    </div>

</template>

<script>
import {createAlchemyWeb3} from "@alch/alchemy-web3"
import lockerContractAbi from '../contract_abis/locker_abi.json'
import minterContractAbi from '../contract_abis/minter_abi.json'
import web3 from 'web3';
import {ETHEREUM_SEPOLIA_RPC, LOCKER_CONTRACT_ADDRESS, ETHEREUM_XRPL_RPC, MINTER_CONTRACT_ADDRESS} from '../env/env'

export default {
    name: "LMView",


    data() {
        return {
            commandType: 0,
            nftContractAddress: "",
            tokenId: 0,
            resultObj: {}
        }
    },

    methods: {
       async getAllTokensUnderUser() {
           console.log("inside")
            const web3i = createAlchemyWeb3(ETHEREUM_SEPOLIA_RPC); // Replace with your Ethereum node URL
            try {
                const contract = new web3i.eth.Contract(lockerContractAbi, LOCKER_CONTRACT_ADDRESS);
                // Account with unlocked wallet in MetaMask or other provider
                console.log(this.$route.params.myAddress)
                const result = await contract.methods.returnAllMyContracts().call({from: this.$route.params.myAddress})
                this.resultObj = result;
                console.log(result);

            } catch (error) {
                console.error('Error:', error);
            }
        },

        async returnAllMyTokens() {
            console.log("inside")
            const web3i = createAlchemyWeb3(ETHEREUM_SEPOLIA_RPC); // Replace with your Ethereum node URL
            try {
                const contract = new web3i.eth.Contract(lockerContractAbi, LOCKER_CONTRACT_ADDRESS);
                // Account with unlocked wallet in MetaMask or other provider
                console.log(this.$route.params.myAddress)
                const result = await contract.methods.returnAllMyTokens(this.nftContractAddress).call({from: this.$route.params.myAddress})
                this.resultObj = result;
                console.log(result);

            } catch (error) {
                console.error('Error:', error);
            }
        },

        async getTokenToDetailsMapSepolia() {
            console.log("inside")
            const web3i = createAlchemyWeb3(ETHEREUM_SEPOLIA_RPC); // Replace with your Ethereum node URL
            try {
                const contract = new web3i.eth.Contract(lockerContractAbi, LOCKER_CONTRACT_ADDRESS);
                // Account with unlocked wallet in MetaMask or other provider
                console.log(this.$route.params.myAddress)
                const result = await contract.methods.tokenIdToNFTMap(this.nftContractAddress, this.tokenId).call()
                this.resultObj = [result];
                console.log(result);

            } catch (error) {
                console.error('Error:', error);
            }
        },

        async getTokenToDetailsXRPL() {
            console.log("inside")
            const web3i = new web3(ETHEREUM_XRPL_RPC); // Replace with your Ethereum node URL
            try {
                const contract = new web3i.eth.Contract(minterContractAbi, MINTER_CONTRACT_ADDRESS);
                // Account with unlocked wallet in MetaMask or other provider
                console.log(this.$route.params.myAddress)
                const result = await contract.methods.tokenIdToNFTDetailsMapping(this.nftContractAddress, this.tokenId).call({})
                // convert bigInt to JSON
                const jsonString = JSON.stringify(result, (key, value) => {
                    if (typeof value === 'bigint') {
                        return value.toString();
                    } 

                    return value
                });                
                this.resultObj = [JSON.parse(jsonString)]
                console.log(jsonString)

            } catch (error) {
                console.error('Error:', error);
            }
        },

        async takeDecision() {
            if (this.tokenId == 0) {
                this.returnAllMyTokens();
            } else if (this.commandType == 0 && this.tokenId != 0) {
                this.getTokenToDetailsMapSepolia();
            } else if (this.commandType == 1 && this.tokenId != 0) {
                this.getTokenToDetailsXRPL();
            }
        }

    }
}
</script>

<style scoped>

#p-cont {
    background-image: linear-gradient(-45deg,pink, black,black, black, skyblue);
    width: 100%;
    font-family:monospace;
    overflow:scroll;
    padding-bottom: 20px;
    padding-top: 20px;
    height: 100vh;
    color: white;
}

.ox-flexi  {
    display: flex;
    text-align: center;
    justify-content: center;
    align-items: center;
    height: 40px;
    margin-bottom: 20px;

}

#tbl {
    display: flex;
    align-items: center;
    justify-content: center;
}

.itm-txt {
    color: white;
}

h2 {
    color: greenyellow;
}

#ins {
    color:violet;
}

h3 {
    color: skyblue;
}


input {
    width: 350px;
    font-family: monospace;
}

#fm-btn {
    height: 25px;
    padding: 2px;
    width: 120px;
    border-radius: 5px;
    color:black;
    background-image: linear-gradient(to right, skyblue, aliceblue);
    margin-right: 15px;
    cursor: pointer;
    font-family: monospace;
    margin-left: 15px;

}

#flexi-de-btn {
    height: 25px;
    padding: 2px;
    width: 300px;
    border-radius: 5px;
    color:black;
    background-image: linear-gradient(to right, skyblue, aliceblue);
    margin-right: 15px;
    cursor: pointer;
    font-family: monospace;
    margin-left: 15px;
}

#ap-d {
    text-align: right;
}

#addr {
    height: 25px;
    padding: 2px;
    width: 200px;
    border-radius: 5px;
    color:black;
    background-image: linear-gradient(to right, skyblue, aliceblue);
    margin-right: 15px;
    cursor: pointer;
    font-family: monospace;
    margin-left: 15px;
}

#fm-btn:hover:enabled { 
    background-image: linear-gradient(to left, skyblue, aliceblue);
    box-shadow: 1px 1px 10px pink; /* Add the box shadow effect on hover */

}

.disbld-btn {
    background-image: linear-gradient(to left, grey, grey);

}

#flexi-de-btn:hover:enabled { 
    background-image: linear-gradient(to left, skyblue, aliceblue);
    box-shadow: 1px 1px 10px pink; /* Add the box shadow effect on hover */

}

#sel-id {
    margin-right: 30px;
    font-family: monospace;
}
#num-id {
    width: 120px;
}

table {
  width: 40%;
  border-collapse: collapse;
  margin-left: 10px;
}

table th, table td {
  border: 1px solid #dddddd;
  padding: 8px;
  text-align: left;
}

#fm-btn:disabled {
    background-image: linear-gradient(to right, gray, white);
}
    
</style>