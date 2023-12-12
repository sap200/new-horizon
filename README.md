# New Horizon Hackathon Judging Repository

#### This repository is a collection of four distinct projects developed for the New Horizon Hackathon. Each project plays a specific role in our comprehensive blockchain application.

<div align="center">
<img src="./assets/blockBorrower_logo.png" alt="Block Borrower Logo" width="400" height="300">
</div>




## Repositories Overview

### block-borrower
- [Try it out](https://funny-salmiakki-254d51.netlify.app/)
- **Functionality**: Front-end interface for user interactions.
- **Details**: Further information can be found in the repository's README.

### block-borrower-contracts
- **Functionality**: Contains the smart contracts for the application.
- **Contained Contracts**:
  - `Auction`: Handles auction logic.
  - `Minter`: Manages token minting.
  - `Locker`: Responsible for locking mechanisms.
  - `BlockBorrower`: Core contract for borrowing operations.
  - `DogNFT`: A test NFT contract for demonstration purposes.
- **Details**: Additional specifics are available in the repository's README.

### sx-relayer
- [Try it out](https://friendly-blini-b6d679.netlify.app/)
- **Note**: Use Mint Button to mint test DogNFTs for testing purpose.
- **Functionality**: Primarily serves as a front-end bridge for transferring NFTs.
- **Details**: More information about its capabilities and usage can be found in the repository's README.


### blockrelayer
- **Functionality**: Implements the relayer logic for our application.
- **Details**: Further details are available in the repository's README.

## Running Instruction
- clone this repository using `git clone https://github.com/sap200/new-horizon.git`
- Inside block-borrower-contracts/contracts deploy Auction.sol, BlockBorrower.sol, Minter.sol and Locker.sol 
- The first 3 would be deployed in EVM XRPL Sidechain (Auction.sol, BlockBorrower.sol, Minter.sol), Locker.sol will be deployed in Sepolia chain.
- Run the blockrelayer but before change this variables, which are self explanatory.
```go

const (
	LOCKER_CONTRACT_ADDRESS                              = "<<MINTER CONTRACT ADDRESS>>" // locker is deployed on Ethereum goerli testnet
	MINTER_CONTRACT_ADDRESS                              = "<<LOCKER CONTRACT ADDRESS>>" // minter is on EVM XRPL sidechain
	PRIVATE_KEY_OF_ACCOUNT_THAT_DEPLOYED_LOCKER_CONTRACT = "<<PRI_KEY>>"                                // private key that deployed locker smart contract on Ethereum goerli testnet
	PUBLIC_KEY_OF_ACCOUNT_THAT_DEPLOYED_LOCKER_CONTRACT  = "<<PUB_KEY>>"                                // public key of the same
	PRIVATE_KEY_OF_ACCOUNT_THAT_DEPLOYED_MINTER_CONTRACT = "<<PRI_KEY>>"
	PUBLIC_KEY_OF_ACCOUNT_THAT_DEPLOYED_MINTER_CONTRACT  = "<<PUB_KEY>>"
	GOERLI_TESTNET_RPC                                   = "<<URL>>"
	GOERLI_TESTNET_RPC_CLIENT                            = "<<URL>>"
)
```

- Use `cd blockrelayer` and `go run main.go` to run the blockrelayer
- use `cd sxrelayer` and `npm install` and `npm run serve` to run the sxrelayer
- Now go inside the blockborrower directory and type `npm install`
- Inside block-borrower/src/env change this constant values which are again self explanatory
``` js
export const AUCTION_CONTRACT_ADDRESS = "0x0B2fcD8736E5726917efb8ab476AFd63e1ae2068";
export const MINTER_CONTRACT_ADDRESS =  "0x5134847Dd630F1ed4B5Cc0fc08328a204CEB2D4D";
export const BLOCKBORROWER_CONTRACT_ADDRESS = "0x94cd50Fc61Bb74DBdFCEc715fF44207faF80Ae54";
```
- Now `npm run serve` to run the project, but before in block-borrower/src/components/NavigationComponents.vue change the link in `line no 25` `<a href="<<change this link to sxrelayer's link>>" class="rt-lnk" target="_blank"><b>🌉 EVM-XRPL NFT Bridge</b></a><hr>`
---

Each sub-repository contains its own README for more detailed information about its specific functionality and implementation.
