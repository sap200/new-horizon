<template>
  <div class="card-component">
    <p v-if="data != null"> <b>Highest Binding Bid : </b> {{getInXRP(data.highestBindingBid)}} XRP </p>
    <p> <b> Highest Bidder : </b> {{data.highestBidder}} </p>
    <p> <b> NFT Contract Address : </b> {{data.nftContractAddress}} </p>
    <p> <b> Bank Id : </b> {{data.bankToWhichItisPledged}} </p>
    <hr>
    <div id="img-div">
        <img :src="myImageURL" />
    </div>
    <hr>

    <table>
        <tbody>
            <tr>
                <td> <b>Auction Id</b> </td>
                <td> {{data.auctionId}} </td>
                <td> <b>Loan Id</b> </td>
                <td> {{data.loanId}} </td>
            </tr>
            <tr>
                <td> <b>Symbol</b> </td>
                <td> {{data.symbol}} </td>
                <td> <b>Token Id</b> </td>
                <td> {{data.tokenId}} </td>
            </tr>
            <tr>
                <td> <b>Start Time</b> </td>
                <td> {{getDate(data.auctionStartTime)}} </td>
                <td> <b>End Time</b> </td>
                <td> {{getDate(Number(data.auctionStartTime) + (7*24*60*60))}} </td>
            </tr>
        </tbody>
    </table>
    <!-- Add more content or customize as needed -->
  </div>
</template>

<script>
import web3 from 'web3';
import erc721Abi from '../contract_abis/erc721_abi.json'
import {EVM_SIDECHAIN_RPC} from '../env/env.js';

export default {
  name: "AuctionCard",

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
        myImageURL: "https://developers.elementor.com/docs/assets/img/elementor-placeholder-image.png",


      }
  },

  created() {
      console.log(this.data)
      this.getAndLoadImage()
  },

  methods: {
      getInXRP(amount) {
          return web3.utils.fromWei(String(amount), "ether")
      },

    async getAndLoadImage() {
     const turi = await this.getTokenURI(this.data.nftContractAddress, this.data.tokenId)
     const imobj = await this.decodeTheTokenURIForImage(turi)
     this.myImageURL = imobj.image

    },

    getDate(sec) {
        let date = new Date(0);
        date.setUTCSeconds(sec)
         const options = {
            year: '2-digit',
            month: '2-digit',
            day: 'numeric',
            hour: '2-digit',
            minute: 'numeric',
            timeZoneName: 'short',
        };

        return date.toLocaleString('en-US', options);
    },

    async decodeTheTokenURIForImage(imageURL) {
            const x = {
                "image": imageURL
            }

         try {
            const response = await fetch(imageURL);

            if (!response.ok) {
                return x
            }

            const data = await response.json();
            console.log(data);
            return data;
        } catch (error) {
            //console.error('Error fetching JSON:', error);
            return x
        }
    },

    async getTokenURI(contractAddress, tokenId) {

           const web3i = new web3(EVM_SIDECHAIN_RPC);
           try {
                const contract = new web3i.eth.Contract(erc721Abi, contractAddress);
                const tokenURI = await contract.methods.tokenURI(tokenId).call({})
                return tokenURI
                
           } catch (error) {
               console.log(error);
               return "https://developers.elementor.com/docs/assets/img/elementor-placeholder-image.png";
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

table {
      width: 100%;
      border-collapse: collapse;
}

th, td {
      border: 1px solid black;
      padding: 8px;
      text-align: left;
      background-color: lightcoral;
}

img {
    width: 200px;
}

#img-div {
    text-align: center;
}

</style>
