package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sap200/env"
	t "github.com/sap200/types"
	"github.com/sap200/utils"
)

var wg = &sync.WaitGroup{}

func StartListener() {
	lockerContractAbiString, err := utils.JSONFileToString("./contract_abis/locker_abi.json")
	if err != nil {
		panic(err)
	}

	minterContractAbiString, err := utils.JSONFileToString("./contract_abis/minter_abi.json")
	if err != nil {
		panic(err)
	}

	lockerContractAbi, err := abi.JSON(strings.NewReader(lockerContractAbiString))
	if err != nil {
		panic(err)
	}
	minterContractAbi, err := abi.JSON(strings.NewReader(minterContractAbiString))
	if err != nil {
		panic(err)
	}

	goerliClient, err := ethclient.DialContext(context.Background(), env.GOERLI_TESTNET_RPC)
	if err != nil {
		panic(err)
	}

	evmxrplClient, err := ethclient.DialContext(context.Background(), env.EVM_XRPL_SIDECHAIN_RPC)
	if err != nil {
		panic(err)
	}

	lockerContractAddress := common.HexToAddress(env.LOCKER_CONTRACT_ADDRESS)
	minterContractAddress := common.HexToAddress(env.MINTER_CONTRACT_ADDRESS)

	blockNumber, err := getCurrentBlockNumber(goerliClient)
	if err != nil {
		panic(err)
	}

	filterQueryGoerli := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockNumber.Int64() - 60), // Replace with the starting block number for Contract 1
		ToBlock:   nil,                                  // Replace with the ending block number (nil for latest)
		Addresses: []common.Address{
			lockerContractAddress,
		},
	}

	ch1 := make(chan types.Log)
	subLocker, err := goerliClient.SubscribeFilterLogs(context.Background(), filterQueryGoerli, ch1)
	if err != nil {
		panic(err)
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		handleEventForLocker(subLocker, ch1, lockerContractAbi, "Locker")
	}()

	go func() {
		defer wg.Done()
		handleEventForMinter(evmxrplClient, minterContractAddress, minterContractAbi, "Minter")
	}()

	fmt.Println("Started sepolia->evmxrpl and sepolia<-evmxrpl relayer ::")

	wg.Wait()

}

func handleEventForLocker(subLocker ethereum.Subscription, ch <-chan types.Log, parsedABI abi.ABI, contractName string) {
	for {
		select {
		case err := <-subLocker.Err():
			fmt.Println("Error from Locker subscriber: ", err)
		case eventLog := <-ch:
			// Decode the event data into a struct
			iface, err := parsedABI.Unpack("LockInitiated", eventLog.Data)
			var data t.LockerNFTDetails
			inputBytes, err := utils.InterfaceToBytes(iface[0])
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = json.Unmarshal(inputBytes, &data)
			if err != nil {
				fmt.Println(err)
				continue
			}
			eventName := iface[1].(string)
			if eventName == t.LOCK_INITIATED {
				fmt.Println("Event received from Locker Contract: ", iface)
				fmt.Println("Debug Log: ", data)
				txHash, err := PerformMintWithRetry(data.Owner, data.ReceipentInDestinationChain, data.TokenURI, data.Symbol, data.NftContractAddress, data.TokenId, env.RETRIES)
				if err != nil {
					fmt.Println("Error while sending TXN to Minter contract: ", err)
					// revert the txn in locker contract
					tx1, err1 := PerformFallbackUnlockingWithRetry(data.NftContractAddress, data.Owner, data.TokenId, env.RETRIES)
					if err1 != nil {
						fmt.Println("Error while sending TXN to Locker contract", err1)
					} else {
						fmt.Println("Txn sent to locker contract for unlocking nft with hash: ", tx1)
					}
				} else {
					fmt.Println("Txn success with hash: ", txHash)
				}

			} else if eventName == "UNLOCKED" {
				fmt.Println("Finalizing Burn")
				tx1, err1 := finalizeBurnWithRetries(common.HexToAddress(data.NftContractAddress), big.NewInt(int64(data.TokenId)), env.RETRIES)
				if err1 != nil {
					fmt.Println("Error wile finalizing burn: ", err1)
				} else {
					fmt.Println("Txn sent for final burn with hash: ", tx1)
				}

			} else {
				fmt.Printf("NO_ACTION:: Event received from %s: %+v\n", contractName, iface)
			}
		}
	}

}

func handleEventForMinter(client *ethclient.Client, minterContractAddress common.Address, parsedABI abi.ABI, contractName string) {
	for {
		// Get the latest block number
		blockNumber, err := client.BlockNumber(context.Background())
		if err != nil {
			fmt.Println("Unable to get current block")
			continue
		}
		filterQueryEvmxrpl := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(blockNumber - env.BLOCK_OFFSET)),
			ToBlock:   nil,
			Addresses: []common.Address{
				minterContractAddress,
			},
		}

		logs, err := client.FilterLogs(context.Background(), filterQueryEvmxrpl)
		if err != nil {
			fmt.Println("Error in handleEventForMinter", err)
			continue
		}

		for _, log := range logs {
			iface, err := parsedABI.Unpack("Minted", log.Data)
			if err != nil {
				fmt.Println(err)
				continue
			}

			logIdentity := utils.HashFunc(strconv.Itoa(int(log.Index)) + string(log.TxHash.Bytes()) + string(log.BlockHash.Bytes()))

			ok, _ := utils.DoesKeyExists(logIdentity)
			fmt.Println("Found", logIdentity)
			if !ok {
				fmt.Println(iface)
				fmt.Println("Not Found", logIdentity)
				nftContractAddress := iface[0].(common.Address)
				sTokenId := iface[1].(*big.Int)
				action := iface[4].(string)
				if action == "MINTED" {
					// proceed to finalize locking
					txnHash, err1 := PerformLockingWithRetry(nftContractAddress, sTokenId, env.RETRIES)
					if err1 != nil {
						fmt.Println("Error while sending TXN to Locker contract", err1)
					} else {
						fmt.Println("Txn sent to locker contract for locking nft with hash: ", txnHash)
					}
				} else if action == "BURN_INITIATED" {
					fmt.Println(nftContractAddress)
					fmt.Println(sTokenId)
					inputBytes, err := utils.InterfaceToBytes(iface[3])
					if err != nil {
						fmt.Println(err)
						continue
					}
					var data t.MinterNFTDetails
					err = json.Unmarshal(inputBytes, &data)
					if err != nil {
						fmt.Println(err)
						continue
					}

					txnHash, err := PerformUnlockingWithRetry(data.ContractAddress, data.ReceipentAddressInSourceChain, data.SourceTokenId, env.RETRIES)
					if err != nil {
						// TODO: make a function in smart contract so user can withdraw NFT
						// TODO: Here update the state to WITHDRAWABLE in case transfer fails
						tx1, err1 := RevertBurnWithRetries(common.HexToAddress(data.ContractAddress), big.NewInt(int64(data.SourceTokenId)), env.RETRIES)
						if err1 != nil {
							fmt.Println("Unable to revert", err1)
						} else {
							fmt.Println("Revert Txn Success: ", tx1)
						}
					} else {
						fmt.Println("Txn sent to locker contract for unlocking nft with hash: ", txnHash)
					}

				}

				utils.AddToDB(logIdentity, env.LOG_DB_VALUE_OF_KEYS)
			}

		}

		// POLL Frequency
		time.Sleep(env.POLLING_SECONDS * time.Second)
	}
}

func getCurrentBlockNumber(client *ethclient.Client) (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	return header.Number, nil
}
