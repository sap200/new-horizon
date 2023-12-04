<template>
  <div class="card-component">
   
      <h4> Loan Details {{getActiveEmoji(data.isActiveLoan)}} </h4>
      <hr>
    <!-- Add more content or customize as needed -->
    <p> <b><u>Contract Address </u> : {{data.nftContractAddress}}</b></p>
    <p><b><u>Customer Address </u> : {{data.pledger}}</b>  </p>
    <hr>
    <div id="img-cntnr">
        <img :src="imagingURL"  />
    </div>
    <hr>

    <table>
        <tbody>
            <tr>
                <td> <b> Loan Id </b> </td>
                <td> {{data.loanId}} </td>
                <td> <b>Token Id</b> </td>
                <td> {{data.tokenId}} </td>

            </tr>
            <tr> 
                <td> <b>Bank Id</b> </td>
                <td> {{data.bankToWhichItisPledged}} </td>
                <td> <b>Loan Amount</b> </td>
                <td> {{getInXRP(data.loanAmount)}} XRP </td>
            </tr>
        </tbody>
    </table>
  </div>
</template>

<script>
import web3 from 'web3';
import erc721Abi from '../contract_abis/erc721_abi.json'
import {EVM_SIDECHAIN_RPC} from '../env/env.js'


export default {
  name: "LoanAndNFTDetailsCard",

  props: {
    data: {
      type: Object,
      required: true,
    },
  },

  data() {
      return {
        evmXrplChainId: "0x15f902",
        imagingURL: "https://developers.elementor.com/docs/assets/img/elementor-placeholder-image.png"
      }
  },

    created() {
        this.getImage(this.data.nftContractAddress, this.data.tokenId);
    },

  methods: {

      getInXRP(amt) {
        return web3.utils.fromWei(amt, "ether")
      },
        
      getActiveEmoji(functional) {
          if(functional) {
              return "✅"
          } else {
              return "🔴"
          }
      },

      async getTokenURI(contractAddress, tokenId) {
           const web3i = new web3(EVM_SIDECHAIN_RPC);
           try {
                const contract = new web3i.eth.Contract(erc721Abi, contractAddress);
                // eslint-disable-next-line
                const tokenURI = await contract.methods.tokenURI(BigInt(tokenId)).call({})

                return tokenURI;
           } catch (error) {
               console.log(error);
               return null;
           }
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

    async getImage(contractAddress, tokenId) {
        const res = await this.getTokenURI(contractAddress, tokenId);
        const iurl = await this.decodeTheTokenURIForImage(res);
        this.imagingURL = iurl.image
    }

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

h4 {
    text-align: center;
    color: darkorchid;
}

p {
    color: darkslategrey;
}

  table {
    width: 100%;
    border-collapse: collapse; /* Ensure borders collapse into a single border */
  }

  th, td {
    border: 1px solid black; /* Set border for table cells */
    padding: 8px; /* Add padding for better readability */
    text-align: left; /* Align text to the left */
    background-color: skyblue;
  }

 th:hover, td:hover {
    background-color: aliceblue;
  }

  th {
    background-color: black; /* Set background color for table header */
  }
  #img-cntnr {
      text-align: center;
      margin-bottom: 5px;

  }

  img {
            border: 2px solid #333; /* 2px solid border with color #333 */
      border-radius: 8px; /* Optional: Add rounded corners for a better look */
      display: block; /* Remove default bottom margin for inline elements */
      margin: 0 auto; /* Center the image horizontally */
      height: auto;
      width: 150px;
  }
</style>
