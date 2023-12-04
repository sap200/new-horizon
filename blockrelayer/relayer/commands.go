package relayer

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sap200/abigens"
	"github.com/sap200/env"
	t "github.com/sap200/types"
)

func performMint(sourceAddress, destinationAddress, tokenURI, symbol, nftContractAddress string, sTokenId uint) (string, error) {
	// Connect to the Ethereum node
	client, err := ethclient.Dial(env.EVM_XRPL_SIDECHAIN_RPC)
	if err != nil {
		return "", err
	}

	// Convert private key hex to crypto.PrivateKey
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(env.PRIVATE_KEY_OF_ACCOUNT_THAT_DEPLOYED_MINTER_CONTRACT, "0x"))
	if err != nil {
		return "", err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", err
	}

	// Create a new contract instance
	contract, err := abigens.NewMinterContract(common.HexToAddress(env.MINTER_CONTRACT_ADDRESS), client)
	if err != nil {
		return "", err
	}

	// Estimate gas for the function cal

	auth.GasLimit = 3000000

	// Call the contract function
	txn, err := contract.StartMintNFT(
		auth,
		common.HexToAddress(sourceAddress),
		common.HexToAddress(destinationAddress),
		t.GOERLI_CHAIN_NAME,
		tokenURI,
		symbol,
		common.HexToAddress(nftContractAddress),
		big.NewInt(int64(sTokenId)),
	)
	if err != nil {
		return "", err
	}

	return txn.Hash().Hex(), nil

}

func PerformMintWithRetry(sourceAddress, destinationAddress, tokenURI, symbol, nftContractAddress string, sTokenId, retries uint) (string, error) {
	txn, err := performMint(sourceAddress, destinationAddress, tokenURI, symbol, nftContractAddress, sTokenId)
	if err != nil {
		for i := 0; i < int(retries); i++ {
			fmt.Println("Retrying txn minting:", retries, "more time::")
			txn, err := performMint(sourceAddress, destinationAddress, tokenURI, symbol, nftContractAddress, sTokenId)
			if err == nil {
				return txn, err
			}
			time.Sleep(env.GAP_BTWN_RETRIES)
		}

		return txn, err
	} else {
		return txn, err
	}
}

func performUnlocking(nftContractAddress, receipentAddress string, tokenId uint) (string, error) {
	// Connect to the Ethereum node
	client, err := ethclient.Dial(env.GOERLI_TESTNET_RPC_CLIENT)
	if err != nil {
		return "", err
	}

	// Convert private key hex to crypto.PrivateKey
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(env.PRIVATE_KEY_OF_ACCOUNT_THAT_DEPLOYED_LOCKER_CONTRACT, "0x"))
	if err != nil {
		return "", err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		fmt.Println(err)
	}

	// Create a new contract instance
	contract, err := abigens.NewLockerContract(common.HexToAddress(env.LOCKER_CONTRACT_ADDRESS), client)
	if err != nil {
		return "", err
	}

	// Estimate gas for the function cal

	auth.GasLimit = 3000000

	// Call the contract function
	txn, err := contract.FinalizeUnlocking(
		auth,
		common.HexToAddress(nftContractAddress),
		common.HexToAddress(receipentAddress),
		big.NewInt(int64(tokenId)),
	)
	if err != nil {
		return "", err
	}

	return txn.Hash().Hex(), err
}

func PerformUnlockingWithRetry(nftContractAddress, receipentAddress string, tokenId, retries uint) (string, error) {
	txn, err := performUnlocking(nftContractAddress, receipentAddress, tokenId)
	if err != nil {
		for i := 0; i < int(retries); i++ {
			fmt.Println("Retrying txn unlocking:", retries, "more time::")
			txn, err := performUnlocking(nftContractAddress, receipentAddress, tokenId)
			if err == nil {
				return txn, err
			}
			time.Sleep(env.GAP_BTWN_RETRIES)
		}

		return txn, err
	} else {
		return txn, err
	}
}

func performFallbackUnlocking(nftContractAddress, receipentAddress string, tokenId uint) (string, error) {
	// Connect to the Ethereum node
	client, err := ethclient.Dial(env.GOERLI_TESTNET_RPC_CLIENT)
	if err != nil {
		return "", err
	}

	// Convert private key hex to crypto.PrivateKey
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(env.PRIVATE_KEY_OF_ACCOUNT_THAT_DEPLOYED_LOCKER_CONTRACT, "0x"))
	if err != nil {
		return "", err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		fmt.Println(err)
	}

	// Create a new contract instance
	contract, err := abigens.NewLockerContract(common.HexToAddress(env.LOCKER_CONTRACT_ADDRESS), client)
	if err != nil {
		return "", err
	}

	// Estimate gas for the function cal

	auth.GasLimit = 3000000

	// Call the contract function
	txn, err := contract.FallBackUnlocking(
		auth,
		common.HexToAddress(nftContractAddress),
		common.HexToAddress(receipentAddress),
		big.NewInt(int64(tokenId)),
	)
	if err != nil {
		return "", err
	}

	return txn.Hash().Hex(), err
}

func PerformFallbackUnlockingWithRetry(nftContractAddress, receipentAddress string, tokenId, retries uint) (string, error) {
	txn, err := performFallbackUnlocking(nftContractAddress, receipentAddress, tokenId)
	if err != nil {
		for i := 0; i < int(retries); i++ {
			fmt.Println("Retrying txn unlocking:", retries, "more time::")
			txn, err := performFallbackUnlocking(nftContractAddress, receipentAddress, tokenId)
			if err == nil {
				return txn, err
			}
			time.Sleep(env.GAP_BTWN_RETRIES)
		}

		return txn, err
	} else {
		return txn, err
	}
}

func performLocking(nftContractAddress common.Address, tokenId *big.Int) (string, error) {
	client, err := ethclient.Dial(env.GOERLI_TESTNET_RPC_CLIENT)
	if err != nil {
		return "", err
	}

	// Convert private key hex to crypto.PrivateKey
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(env.PRIVATE_KEY_OF_ACCOUNT_THAT_DEPLOYED_LOCKER_CONTRACT, "0x"))
	if err != nil {
		return "", err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", err
	}

	// Create a new contract instance
	contract, err := abigens.NewLockerContract(common.HexToAddress(env.LOCKER_CONTRACT_ADDRESS), client)
	if err != nil {
		return "", err
	}

	// Estimate gas for the function cal

	auth.GasLimit = 3000000

	// Call the contract function
	txn, err := contract.FinalizeLock(
		auth,
		nftContractAddress,
		tokenId,
	)
	if err != nil {
		return "", err
	}

	return txn.Hash().Hex(), nil
}

func PerformLockingWithRetry(nftContractAddress common.Address, tokenId *big.Int, retries uint) (string, error) {
	txn, err := performLocking(nftContractAddress, tokenId)
	if err != nil {
		for i := 0; i < int(retries); i++ {
			fmt.Println("Retrying txn Locking:", retries, "more time::")
			txn, err := performLocking(nftContractAddress, tokenId)
			if err == nil {
				return txn, err
			}
			time.Sleep(env.GAP_BTWN_RETRIES)
		}

		return txn, err
	} else {
		return txn, err
	}
}

func revertBurn(nftContractAddress common.Address, sTokenId *big.Int) (string, error) {
	client, err := ethclient.Dial(env.EVM_XRPL_SIDECHAIN_RPC)
	if err != nil {
		return "", err
	}

	// Convert private key hex to crypto.PrivateKey
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(env.PRIVATE_KEY_OF_ACCOUNT_THAT_DEPLOYED_MINTER_CONTRACT, "0x"))
	if err != nil {
		return "", err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", err
	}

	// Create a new contract instance
	contract, err := abigens.NewMinterContract(common.HexToAddress(env.MINTER_CONTRACT_ADDRESS), client)
	if err != nil {
		return "", err
	}

	// Estimate gas for the function cal

	auth.GasLimit = 3000000

	// Call the contract function
	txn, err := contract.RevertBurn(
		auth,
		nftContractAddress,
		sTokenId,
	)
	if err != nil {
		return "", err
	}

	return txn.Hash().Hex(), nil
}

func RevertBurnWithRetries(nftContractAddress common.Address, sTokenId *big.Int, retries uint) (string, error) {
	txn, err := revertBurn(nftContractAddress, sTokenId)
	if err != nil {
		for i := 0; i < int(retries); i++ {
			fmt.Println("Retrying txn revert_burn:", retries, "more time::")
			txn, err := revertBurn(nftContractAddress, sTokenId)
			if err == nil {
				return txn, err
			}
			time.Sleep(env.GAP_BTWN_RETRIES)
		}

		return txn, err
	} else {
		return txn, err
	}
}

func finalizeBurn(nftContractAddress common.Address, sTokenId *big.Int) (string, error) {
	client, err := ethclient.Dial(env.EVM_XRPL_SIDECHAIN_RPC)
	if err != nil {
		return "", err
	}

	// Convert private key hex to crypto.PrivateKey
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(env.PRIVATE_KEY_OF_ACCOUNT_THAT_DEPLOYED_MINTER_CONTRACT, "0x"))
	if err != nil {
		return "", err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", err
	}

	// Create a new contract instance
	contract, err := abigens.NewMinterContract(common.HexToAddress(env.MINTER_CONTRACT_ADDRESS), client)
	if err != nil {
		return "", err
	}

	// Estimate gas for the function cal

	auth.GasLimit = 3000000

	// Call the contract function
	txn, err := contract.FinalizeBurn(
		auth,
		nftContractAddress,
		sTokenId,
	)
	if err != nil {
		return "", err
	}

	return txn.Hash().Hex(), nil

}

func finalizeBurnWithRetries(nftContractAddress common.Address, sTokenId *big.Int, retries uint) (string, error) {
	txn, err := finalizeBurn(nftContractAddress, sTokenId)
	if err != nil {
		for i := 0; i < int(retries); i++ {
			fmt.Println("Retrying txn Burn:", retries, "more time::")
			txn, err := finalizeBurn(nftContractAddress, sTokenId)
			if err == nil {
				return txn, err
			}
			time.Sleep(env.GAP_BTWN_RETRIES)
		}

		return txn, err
	} else {
		return txn, err
	}
}
