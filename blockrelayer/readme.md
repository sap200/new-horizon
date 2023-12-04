## Block Relayer 
### For Asset Transfer Between Sepolia Blockchain and EVM XRPL Sidechain

### Introduction
This project introduces a relayer designed to seamlessly facilitate asset transfers between the Sepolia blockchain and an EVM-compatible XRPL (XRP Ledger) sidechain. By listening to events from smart contracts, the relayer efficiently processes and relays information, ensuring smooth and secure cross-chain transactions.

### Key Features
- **Event Listening**: Continuously monitors smart contract events on the Sepolia blockchain.
- **Message Relaying**: Processes and relays messages between Sepolia and the EVM XRPL sidechain.
- **Cross-Chain Asset Transfer**: Enables users to transfer assets between two different blockchain architectures.
- **Security and Efficiency**: Focuses on secure and efficient communication between blockchains.

### Technology Stack
- **Blockchain Networks**: Sepolia and EVM-compatible XRPL Sidechain.
- **Smart Contract Development**: [Programming languages and frameworks used, e.g., Solidity].
- **Backend**: [Technologies used for backend development, e.g., Node.js, Web3.js].
- **Additional Tools**: [Any additional tools or technologies used in the project].

### Getting Started
Guidelines on setting up and running the relayer locally:
1. Clone the repository: `git clone https://github.com/sap200/blockrelayer`.
4. Start the relayer service: `go run main.go` or an equivalent command.

## Configuration
- deploy the Minter and Locker Smart contract in EVM XRPL compatible sidechain and sepolia chain respectively.
- Fill in these details in env/env.go
```go
const (
	LOCKER_CONTRACT_ADDRESS                              = "<<LOCKER CONTRACT ADDRESS>>"                       // locker is deployed on Ethereum goerli testnet
	MINTER_CONTRACT_ADDRESS                              = "<<MINTER CONTRACT ADDRESS>>"                       // minter is on EVM XRPL sidechain
	PRIVATE_KEY_OF_ACCOUNT_THAT_DEPLOYED_LOCKER_CONTRACT = "<<PRIVATE KEY>>" // private key that deployed locker smart contract on Ethereum goerli testnet
	PRIVATE_KEY_OF_ACCOUNT_THAT_DEPLOYED_MINTER_CONTRACT = "<<PRIVATE KEY>>"
	GOERLI_TESTNET_RPC                                   = "<<WSS URL FOR EVENT LISTENING>>"
	EVM_XRPL_SIDECHAIN_RPC                               = "https://rpc-evm-sidechain.xrpl.org/"
	GOERLI_TESTNET_RPC_CLIENT                            = "<<SEPOLIA RPC URL>>"
	LOG_DB_NAME                                          = "UNIQUE_LOG_DB"
	LOG_DB_VALUE_OF_KEYS                                 = "EXISTS"
)

```

## Contributing
Provide guidelines for those who wish to contribute to the project:
- Fork the repository.
- Follow coding standards and commit guidelines.
- Submit pull requests with detailed descriptions of changes and improvements.

