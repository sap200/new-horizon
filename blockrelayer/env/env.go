package env

const (
	LOCKER_CONTRACT_ADDRESS                              = "0x1f8871584c79D0B87C242757D2d2c88915f04929" // locker is deployed on Ethereum goerli testnet
	MINTER_CONTRACT_ADDRESS                              = "0x5134847Dd630F1ed4B5Cc0fc08328a204CEB2D4D" // minter is on EVM XRPL sidechain
	PRIVATE_KEY_OF_ACCOUNT_THAT_DEPLOYED_LOCKER_CONTRACT = "<<PRI_KEY>>"                                // private key that deployed locker smart contract on Ethereum goerli testnet
	PUBLIC_KEY_OF_ACCOUNT_THAT_DEPLOYED_LOCKER_CONTRACT  = "<<PUB_KEY>>"                                // public key of the same
	PRIVATE_KEY_OF_ACCOUNT_THAT_DEPLOYED_MINTER_CONTRACT = "<<PRI_KEY>>"
	PUBLIC_KEY_OF_ACCOUNT_THAT_DEPLOYED_MINTER_CONTRACT  = "<<PUB_KEY>>"
	GOERLI_TESTNET_RPC                                   = "<<URL>>"
	EVM_XRPL_SIDECHAIN_RPC                               = "https://rpc-evm-sidechain.xrpl.org/"
	GOERLI_TESTNET_RPC_CLIENT                            = "<<URL>>"
	RETRIES                                              = 3
	GAP_BTWN_RETRIES                                     = 5000
	POLLING_SECONDS                                      = 30
	BLOCK_OFFSET                                         = 20
	LOG_DB_NAME                                          = "UNIQUE_LOG_DB"
	LOG_DB_VALUE_OF_KEYS                                 = "EXISTS"
)
