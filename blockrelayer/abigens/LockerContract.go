// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigens

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LockerLockedNFTDetails is an auto generated low-level Go binding around an user-defined struct.
type LockerLockedNFTDetails struct {
	TokenId                     *big.Int
	NftContractAddress          common.Address
	Owner                       common.Address
	ReceipentInDestinationChain common.Address
	DestinationChainName        string
	LockState                   uint8
	Symbol                      string
	TokenURI                    string
}

// LockerContractMetaData contains all meta data concerning the LockerContract contract.
var LockerContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receipentInDestinationChain\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"destinationChainName\",\"type\":\"string\"},{\"internalType\":\"enumLocker.LockState\",\"name\":\"lockState\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structLocker.LockedNFTDetails\",\"name\":\"_lockedNFTDetails\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"name\":\"LockInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receipentInDestinationChain\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"destinationChainName\",\"type\":\"string\"},{\"internalType\":\"enumLocker.LockState\",\"name\":\"lockState\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structLocker.LockedNFTDetails\",\"name\":\"_lockedNFTDetails\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"name\":\"LockedNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receipentInDestinationChain\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"destinationChainName\",\"type\":\"string\"},{\"internalType\":\"enumLocker.LockState\",\"name\":\"lockState\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structLocker.LockedNFTDetails\",\"name\":\"_lockedNFTDetails\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"name\":\"UnlockedNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receipentInDestinationChain\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"destinationChainName\",\"type\":\"string\"},{\"internalType\":\"enumLocker.LockState\",\"name\":\"lockState\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structLocker.LockedNFTDetails\",\"name\":\"_lockedNFTDetails\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"name\":\"WithdrawnNFT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receipent\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"fallBackUnlocking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"finalizeLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receipent\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"finalizeUnlocking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"returnAllMyContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"}],\"name\":\"returnAllMyTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receipentAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"destinationChainName\",\"type\":\"string\"}],\"name\":\"sendNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIdToNFTMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receipentInDestinationChain\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"destinationChainName\",\"type\":\"string\"},{\"internalType\":\"enumLocker.LockState\",\"name\":\"lockState\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LockerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use LockerContractMetaData.ABI instead.
var LockerContractABI = LockerContractMetaData.ABI

// LockerContract is an auto generated Go binding around an Ethereum contract.
type LockerContract struct {
	LockerContractCaller     // Read-only binding to the contract
	LockerContractTransactor // Write-only binding to the contract
	LockerContractFilterer   // Log filterer for contract events
}

// LockerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockerContractSession struct {
	Contract     *LockerContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockerContractCallerSession struct {
	Contract *LockerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LockerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockerContractTransactorSession struct {
	Contract     *LockerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LockerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockerContractRaw struct {
	Contract *LockerContract // Generic contract binding to access the raw methods on
}

// LockerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockerContractCallerRaw struct {
	Contract *LockerContractCaller // Generic read-only contract binding to access the raw methods on
}

// LockerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockerContractTransactorRaw struct {
	Contract *LockerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLockerContract creates a new instance of LockerContract, bound to a specific deployed contract.
func NewLockerContract(address common.Address, backend bind.ContractBackend) (*LockerContract, error) {
	contract, err := bindLockerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LockerContract{LockerContractCaller: LockerContractCaller{contract: contract}, LockerContractTransactor: LockerContractTransactor{contract: contract}, LockerContractFilterer: LockerContractFilterer{contract: contract}}, nil
}

// NewLockerContractCaller creates a new read-only instance of LockerContract, bound to a specific deployed contract.
func NewLockerContractCaller(address common.Address, caller bind.ContractCaller) (*LockerContractCaller, error) {
	contract, err := bindLockerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockerContractCaller{contract: contract}, nil
}

// NewLockerContractTransactor creates a new write-only instance of LockerContract, bound to a specific deployed contract.
func NewLockerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*LockerContractTransactor, error) {
	contract, err := bindLockerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockerContractTransactor{contract: contract}, nil
}

// NewLockerContractFilterer creates a new log filterer instance of LockerContract, bound to a specific deployed contract.
func NewLockerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*LockerContractFilterer, error) {
	contract, err := bindLockerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockerContractFilterer{contract: contract}, nil
}

// bindLockerContract binds a generic wrapper to an already deployed contract.
func bindLockerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LockerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockerContract *LockerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LockerContract.Contract.LockerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockerContract *LockerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockerContract.Contract.LockerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockerContract *LockerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockerContract.Contract.LockerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockerContract *LockerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LockerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockerContract *LockerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockerContract *LockerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockerContract.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LockerContract *LockerContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockerContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LockerContract *LockerContractSession) Owner() (common.Address, error) {
	return _LockerContract.Contract.Owner(&_LockerContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LockerContract *LockerContractCallerSession) Owner() (common.Address, error) {
	return _LockerContract.Contract.Owner(&_LockerContract.CallOpts)
}

// ReturnAllMyContracts is a free data retrieval call binding the contract method 0x8b4fc180.
//
// Solidity: function returnAllMyContracts() view returns(address[])
func (_LockerContract *LockerContractCaller) ReturnAllMyContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LockerContract.contract.Call(opts, &out, "returnAllMyContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ReturnAllMyContracts is a free data retrieval call binding the contract method 0x8b4fc180.
//
// Solidity: function returnAllMyContracts() view returns(address[])
func (_LockerContract *LockerContractSession) ReturnAllMyContracts() ([]common.Address, error) {
	return _LockerContract.Contract.ReturnAllMyContracts(&_LockerContract.CallOpts)
}

// ReturnAllMyContracts is a free data retrieval call binding the contract method 0x8b4fc180.
//
// Solidity: function returnAllMyContracts() view returns(address[])
func (_LockerContract *LockerContractCallerSession) ReturnAllMyContracts() ([]common.Address, error) {
	return _LockerContract.Contract.ReturnAllMyContracts(&_LockerContract.CallOpts)
}

// ReturnAllMyTokens is a free data retrieval call binding the contract method 0x6de0009f.
//
// Solidity: function returnAllMyTokens(address nftContractAddress) view returns(uint256[])
func (_LockerContract *LockerContractCaller) ReturnAllMyTokens(opts *bind.CallOpts, nftContractAddress common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _LockerContract.contract.Call(opts, &out, "returnAllMyTokens", nftContractAddress)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ReturnAllMyTokens is a free data retrieval call binding the contract method 0x6de0009f.
//
// Solidity: function returnAllMyTokens(address nftContractAddress) view returns(uint256[])
func (_LockerContract *LockerContractSession) ReturnAllMyTokens(nftContractAddress common.Address) ([]*big.Int, error) {
	return _LockerContract.Contract.ReturnAllMyTokens(&_LockerContract.CallOpts, nftContractAddress)
}

// ReturnAllMyTokens is a free data retrieval call binding the contract method 0x6de0009f.
//
// Solidity: function returnAllMyTokens(address nftContractAddress) view returns(uint256[])
func (_LockerContract *LockerContractCallerSession) ReturnAllMyTokens(nftContractAddress common.Address) ([]*big.Int, error) {
	return _LockerContract.Contract.ReturnAllMyTokens(&_LockerContract.CallOpts, nftContractAddress)
}

// TokenIdToNFTMap is a free data retrieval call binding the contract method 0x9ba222cc.
//
// Solidity: function tokenIdToNFTMap(address , uint256 ) view returns(uint256 tokenId, address nftContractAddress, address owner, address receipentInDestinationChain, string destinationChainName, uint8 lockState, string symbol, string tokenURI)
func (_LockerContract *LockerContractCaller) TokenIdToNFTMap(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	TokenId                     *big.Int
	NftContractAddress          common.Address
	Owner                       common.Address
	ReceipentInDestinationChain common.Address
	DestinationChainName        string
	LockState                   uint8
	Symbol                      string
	TokenURI                    string
}, error) {
	var out []interface{}
	err := _LockerContract.contract.Call(opts, &out, "tokenIdToNFTMap", arg0, arg1)

	outstruct := new(struct {
		TokenId                     *big.Int
		NftContractAddress          common.Address
		Owner                       common.Address
		ReceipentInDestinationChain common.Address
		DestinationChainName        string
		LockState                   uint8
		Symbol                      string
		TokenURI                    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NftContractAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ReceipentInDestinationChain = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.DestinationChainName = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.LockState = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Symbol = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.TokenURI = *abi.ConvertType(out[7], new(string)).(*string)

	return *outstruct, err

}

// TokenIdToNFTMap is a free data retrieval call binding the contract method 0x9ba222cc.
//
// Solidity: function tokenIdToNFTMap(address , uint256 ) view returns(uint256 tokenId, address nftContractAddress, address owner, address receipentInDestinationChain, string destinationChainName, uint8 lockState, string symbol, string tokenURI)
func (_LockerContract *LockerContractSession) TokenIdToNFTMap(arg0 common.Address, arg1 *big.Int) (struct {
	TokenId                     *big.Int
	NftContractAddress          common.Address
	Owner                       common.Address
	ReceipentInDestinationChain common.Address
	DestinationChainName        string
	LockState                   uint8
	Symbol                      string
	TokenURI                    string
}, error) {
	return _LockerContract.Contract.TokenIdToNFTMap(&_LockerContract.CallOpts, arg0, arg1)
}

// TokenIdToNFTMap is a free data retrieval call binding the contract method 0x9ba222cc.
//
// Solidity: function tokenIdToNFTMap(address , uint256 ) view returns(uint256 tokenId, address nftContractAddress, address owner, address receipentInDestinationChain, string destinationChainName, uint8 lockState, string symbol, string tokenURI)
func (_LockerContract *LockerContractCallerSession) TokenIdToNFTMap(arg0 common.Address, arg1 *big.Int) (struct {
	TokenId                     *big.Int
	NftContractAddress          common.Address
	Owner                       common.Address
	ReceipentInDestinationChain common.Address
	DestinationChainName        string
	LockState                   uint8
	Symbol                      string
	TokenURI                    string
}, error) {
	return _LockerContract.Contract.TokenIdToNFTMap(&_LockerContract.CallOpts, arg0, arg1)
}

// FallBackUnlocking is a paid mutator transaction binding the contract method 0xf74f2045.
//
// Solidity: function fallBackUnlocking(address nftContractAddress, address receipent, uint256 tokenId) returns()
func (_LockerContract *LockerContractTransactor) FallBackUnlocking(opts *bind.TransactOpts, nftContractAddress common.Address, receipent common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LockerContract.contract.Transact(opts, "fallBackUnlocking", nftContractAddress, receipent, tokenId)
}

// FallBackUnlocking is a paid mutator transaction binding the contract method 0xf74f2045.
//
// Solidity: function fallBackUnlocking(address nftContractAddress, address receipent, uint256 tokenId) returns()
func (_LockerContract *LockerContractSession) FallBackUnlocking(nftContractAddress common.Address, receipent common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LockerContract.Contract.FallBackUnlocking(&_LockerContract.TransactOpts, nftContractAddress, receipent, tokenId)
}

// FallBackUnlocking is a paid mutator transaction binding the contract method 0xf74f2045.
//
// Solidity: function fallBackUnlocking(address nftContractAddress, address receipent, uint256 tokenId) returns()
func (_LockerContract *LockerContractTransactorSession) FallBackUnlocking(nftContractAddress common.Address, receipent common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LockerContract.Contract.FallBackUnlocking(&_LockerContract.TransactOpts, nftContractAddress, receipent, tokenId)
}

// FinalizeLock is a paid mutator transaction binding the contract method 0x0db9bb74.
//
// Solidity: function finalizeLock(address nftContractAddress, uint256 tokenId) returns()
func (_LockerContract *LockerContractTransactor) FinalizeLock(opts *bind.TransactOpts, nftContractAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LockerContract.contract.Transact(opts, "finalizeLock", nftContractAddress, tokenId)
}

// FinalizeLock is a paid mutator transaction binding the contract method 0x0db9bb74.
//
// Solidity: function finalizeLock(address nftContractAddress, uint256 tokenId) returns()
func (_LockerContract *LockerContractSession) FinalizeLock(nftContractAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LockerContract.Contract.FinalizeLock(&_LockerContract.TransactOpts, nftContractAddress, tokenId)
}

// FinalizeLock is a paid mutator transaction binding the contract method 0x0db9bb74.
//
// Solidity: function finalizeLock(address nftContractAddress, uint256 tokenId) returns()
func (_LockerContract *LockerContractTransactorSession) FinalizeLock(nftContractAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LockerContract.Contract.FinalizeLock(&_LockerContract.TransactOpts, nftContractAddress, tokenId)
}

// FinalizeUnlocking is a paid mutator transaction binding the contract method 0x3703c51c.
//
// Solidity: function finalizeUnlocking(address nftContractAddress, address receipent, uint256 tokenId) returns()
func (_LockerContract *LockerContractTransactor) FinalizeUnlocking(opts *bind.TransactOpts, nftContractAddress common.Address, receipent common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LockerContract.contract.Transact(opts, "finalizeUnlocking", nftContractAddress, receipent, tokenId)
}

// FinalizeUnlocking is a paid mutator transaction binding the contract method 0x3703c51c.
//
// Solidity: function finalizeUnlocking(address nftContractAddress, address receipent, uint256 tokenId) returns()
func (_LockerContract *LockerContractSession) FinalizeUnlocking(nftContractAddress common.Address, receipent common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LockerContract.Contract.FinalizeUnlocking(&_LockerContract.TransactOpts, nftContractAddress, receipent, tokenId)
}

// FinalizeUnlocking is a paid mutator transaction binding the contract method 0x3703c51c.
//
// Solidity: function finalizeUnlocking(address nftContractAddress, address receipent, uint256 tokenId) returns()
func (_LockerContract *LockerContractTransactorSession) FinalizeUnlocking(nftContractAddress common.Address, receipent common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LockerContract.Contract.FinalizeUnlocking(&_LockerContract.TransactOpts, nftContractAddress, receipent, tokenId)
}

// SendNFT is a paid mutator transaction binding the contract method 0x0699adf8.
//
// Solidity: function sendNFT(uint256 tokenId, address nftContractAddress, address receipentAddress, string destinationChainName) returns()
func (_LockerContract *LockerContractTransactor) SendNFT(opts *bind.TransactOpts, tokenId *big.Int, nftContractAddress common.Address, receipentAddress common.Address, destinationChainName string) (*types.Transaction, error) {
	return _LockerContract.contract.Transact(opts, "sendNFT", tokenId, nftContractAddress, receipentAddress, destinationChainName)
}

// SendNFT is a paid mutator transaction binding the contract method 0x0699adf8.
//
// Solidity: function sendNFT(uint256 tokenId, address nftContractAddress, address receipentAddress, string destinationChainName) returns()
func (_LockerContract *LockerContractSession) SendNFT(tokenId *big.Int, nftContractAddress common.Address, receipentAddress common.Address, destinationChainName string) (*types.Transaction, error) {
	return _LockerContract.Contract.SendNFT(&_LockerContract.TransactOpts, tokenId, nftContractAddress, receipentAddress, destinationChainName)
}

// SendNFT is a paid mutator transaction binding the contract method 0x0699adf8.
//
// Solidity: function sendNFT(uint256 tokenId, address nftContractAddress, address receipentAddress, string destinationChainName) returns()
func (_LockerContract *LockerContractTransactorSession) SendNFT(tokenId *big.Int, nftContractAddress common.Address, receipentAddress common.Address, destinationChainName string) (*types.Transaction, error) {
	return _LockerContract.Contract.SendNFT(&_LockerContract.TransactOpts, tokenId, nftContractAddress, receipentAddress, destinationChainName)
}

// WithdrawNFT is a paid mutator transaction binding the contract method 0x6088e93a.
//
// Solidity: function withdrawNFT(address nftContractAddress, uint256 tokenId) returns()
func (_LockerContract *LockerContractTransactor) WithdrawNFT(opts *bind.TransactOpts, nftContractAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LockerContract.contract.Transact(opts, "withdrawNFT", nftContractAddress, tokenId)
}

// WithdrawNFT is a paid mutator transaction binding the contract method 0x6088e93a.
//
// Solidity: function withdrawNFT(address nftContractAddress, uint256 tokenId) returns()
func (_LockerContract *LockerContractSession) WithdrawNFT(nftContractAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LockerContract.Contract.WithdrawNFT(&_LockerContract.TransactOpts, nftContractAddress, tokenId)
}

// WithdrawNFT is a paid mutator transaction binding the contract method 0x6088e93a.
//
// Solidity: function withdrawNFT(address nftContractAddress, uint256 tokenId) returns()
func (_LockerContract *LockerContractTransactorSession) WithdrawNFT(nftContractAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LockerContract.Contract.WithdrawNFT(&_LockerContract.TransactOpts, nftContractAddress, tokenId)
}

// LockerContractLockInitiatedIterator is returned from FilterLockInitiated and is used to iterate over the raw logs and unpacked data for LockInitiated events raised by the LockerContract contract.
type LockerContractLockInitiatedIterator struct {
	Event *LockerContractLockInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockerContractLockInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockerContractLockInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockerContractLockInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockerContractLockInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockerContractLockInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockerContractLockInitiated represents a LockInitiated event raised by the LockerContract contract.
type LockerContractLockInitiated struct {
	Timestamp        *big.Int
	LockedNFTDetails LockerLockedNFTDetails
	Status           string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLockInitiated is a free log retrieval operation binding the contract event 0x567754a4afc5e2b1dc0c007728368f8b25cbc441c08de8af45848d490c90c331.
//
// Solidity: event LockInitiated(uint256 indexed _timestamp, (uint256,address,address,address,string,uint8,string,string) _lockedNFTDetails, string status)
func (_LockerContract *LockerContractFilterer) FilterLockInitiated(opts *bind.FilterOpts, _timestamp []*big.Int) (*LockerContractLockInitiatedIterator, error) {

	var _timestampRule []interface{}
	for _, _timestampItem := range _timestamp {
		_timestampRule = append(_timestampRule, _timestampItem)
	}

	logs, sub, err := _LockerContract.contract.FilterLogs(opts, "LockInitiated", _timestampRule)
	if err != nil {
		return nil, err
	}
	return &LockerContractLockInitiatedIterator{contract: _LockerContract.contract, event: "LockInitiated", logs: logs, sub: sub}, nil
}

// WatchLockInitiated is a free log subscription operation binding the contract event 0x567754a4afc5e2b1dc0c007728368f8b25cbc441c08de8af45848d490c90c331.
//
// Solidity: event LockInitiated(uint256 indexed _timestamp, (uint256,address,address,address,string,uint8,string,string) _lockedNFTDetails, string status)
func (_LockerContract *LockerContractFilterer) WatchLockInitiated(opts *bind.WatchOpts, sink chan<- *LockerContractLockInitiated, _timestamp []*big.Int) (event.Subscription, error) {

	var _timestampRule []interface{}
	for _, _timestampItem := range _timestamp {
		_timestampRule = append(_timestampRule, _timestampItem)
	}

	logs, sub, err := _LockerContract.contract.WatchLogs(opts, "LockInitiated", _timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockerContractLockInitiated)
				if err := _LockerContract.contract.UnpackLog(event, "LockInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLockInitiated is a log parse operation binding the contract event 0x567754a4afc5e2b1dc0c007728368f8b25cbc441c08de8af45848d490c90c331.
//
// Solidity: event LockInitiated(uint256 indexed _timestamp, (uint256,address,address,address,string,uint8,string,string) _lockedNFTDetails, string status)
func (_LockerContract *LockerContractFilterer) ParseLockInitiated(log types.Log) (*LockerContractLockInitiated, error) {
	event := new(LockerContractLockInitiated)
	if err := _LockerContract.contract.UnpackLog(event, "LockInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockerContractLockedNFTIterator is returned from FilterLockedNFT and is used to iterate over the raw logs and unpacked data for LockedNFT events raised by the LockerContract contract.
type LockerContractLockedNFTIterator struct {
	Event *LockerContractLockedNFT // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockerContractLockedNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockerContractLockedNFT)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockerContractLockedNFT)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockerContractLockedNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockerContractLockedNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockerContractLockedNFT represents a LockedNFT event raised by the LockerContract contract.
type LockerContractLockedNFT struct {
	Timestamp        *big.Int
	LockedNFTDetails LockerLockedNFTDetails
	Status           string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLockedNFT is a free log retrieval operation binding the contract event 0x55b16d1d4c18d7391766529364f6b1d4bfedb753dbfc84d35d8fe439730a082a.
//
// Solidity: event LockedNFT(uint256 indexed _timestamp, (uint256,address,address,address,string,uint8,string,string) _lockedNFTDetails, string status)
func (_LockerContract *LockerContractFilterer) FilterLockedNFT(opts *bind.FilterOpts, _timestamp []*big.Int) (*LockerContractLockedNFTIterator, error) {

	var _timestampRule []interface{}
	for _, _timestampItem := range _timestamp {
		_timestampRule = append(_timestampRule, _timestampItem)
	}

	logs, sub, err := _LockerContract.contract.FilterLogs(opts, "LockedNFT", _timestampRule)
	if err != nil {
		return nil, err
	}
	return &LockerContractLockedNFTIterator{contract: _LockerContract.contract, event: "LockedNFT", logs: logs, sub: sub}, nil
}

// WatchLockedNFT is a free log subscription operation binding the contract event 0x55b16d1d4c18d7391766529364f6b1d4bfedb753dbfc84d35d8fe439730a082a.
//
// Solidity: event LockedNFT(uint256 indexed _timestamp, (uint256,address,address,address,string,uint8,string,string) _lockedNFTDetails, string status)
func (_LockerContract *LockerContractFilterer) WatchLockedNFT(opts *bind.WatchOpts, sink chan<- *LockerContractLockedNFT, _timestamp []*big.Int) (event.Subscription, error) {

	var _timestampRule []interface{}
	for _, _timestampItem := range _timestamp {
		_timestampRule = append(_timestampRule, _timestampItem)
	}

	logs, sub, err := _LockerContract.contract.WatchLogs(opts, "LockedNFT", _timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockerContractLockedNFT)
				if err := _LockerContract.contract.UnpackLog(event, "LockedNFT", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLockedNFT is a log parse operation binding the contract event 0x55b16d1d4c18d7391766529364f6b1d4bfedb753dbfc84d35d8fe439730a082a.
//
// Solidity: event LockedNFT(uint256 indexed _timestamp, (uint256,address,address,address,string,uint8,string,string) _lockedNFTDetails, string status)
func (_LockerContract *LockerContractFilterer) ParseLockedNFT(log types.Log) (*LockerContractLockedNFT, error) {
	event := new(LockerContractLockedNFT)
	if err := _LockerContract.contract.UnpackLog(event, "LockedNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockerContractUnlockedNFTIterator is returned from FilterUnlockedNFT and is used to iterate over the raw logs and unpacked data for UnlockedNFT events raised by the LockerContract contract.
type LockerContractUnlockedNFTIterator struct {
	Event *LockerContractUnlockedNFT // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockerContractUnlockedNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockerContractUnlockedNFT)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockerContractUnlockedNFT)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockerContractUnlockedNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockerContractUnlockedNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockerContractUnlockedNFT represents a UnlockedNFT event raised by the LockerContract contract.
type LockerContractUnlockedNFT struct {
	Timestamp        *big.Int
	LockedNFTDetails LockerLockedNFTDetails
	Status           string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUnlockedNFT is a free log retrieval operation binding the contract event 0xd3f5baff4596e3d4a2409b45d4c0bd700920bc145804fe5d7f839f7aa33eb6f7.
//
// Solidity: event UnlockedNFT(uint256 indexed _timestamp, (uint256,address,address,address,string,uint8,string,string) _lockedNFTDetails, string status)
func (_LockerContract *LockerContractFilterer) FilterUnlockedNFT(opts *bind.FilterOpts, _timestamp []*big.Int) (*LockerContractUnlockedNFTIterator, error) {

	var _timestampRule []interface{}
	for _, _timestampItem := range _timestamp {
		_timestampRule = append(_timestampRule, _timestampItem)
	}

	logs, sub, err := _LockerContract.contract.FilterLogs(opts, "UnlockedNFT", _timestampRule)
	if err != nil {
		return nil, err
	}
	return &LockerContractUnlockedNFTIterator{contract: _LockerContract.contract, event: "UnlockedNFT", logs: logs, sub: sub}, nil
}

// WatchUnlockedNFT is a free log subscription operation binding the contract event 0xd3f5baff4596e3d4a2409b45d4c0bd700920bc145804fe5d7f839f7aa33eb6f7.
//
// Solidity: event UnlockedNFT(uint256 indexed _timestamp, (uint256,address,address,address,string,uint8,string,string) _lockedNFTDetails, string status)
func (_LockerContract *LockerContractFilterer) WatchUnlockedNFT(opts *bind.WatchOpts, sink chan<- *LockerContractUnlockedNFT, _timestamp []*big.Int) (event.Subscription, error) {

	var _timestampRule []interface{}
	for _, _timestampItem := range _timestamp {
		_timestampRule = append(_timestampRule, _timestampItem)
	}

	logs, sub, err := _LockerContract.contract.WatchLogs(opts, "UnlockedNFT", _timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockerContractUnlockedNFT)
				if err := _LockerContract.contract.UnpackLog(event, "UnlockedNFT", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnlockedNFT is a log parse operation binding the contract event 0xd3f5baff4596e3d4a2409b45d4c0bd700920bc145804fe5d7f839f7aa33eb6f7.
//
// Solidity: event UnlockedNFT(uint256 indexed _timestamp, (uint256,address,address,address,string,uint8,string,string) _lockedNFTDetails, string status)
func (_LockerContract *LockerContractFilterer) ParseUnlockedNFT(log types.Log) (*LockerContractUnlockedNFT, error) {
	event := new(LockerContractUnlockedNFT)
	if err := _LockerContract.contract.UnpackLog(event, "UnlockedNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockerContractWithdrawnNFTIterator is returned from FilterWithdrawnNFT and is used to iterate over the raw logs and unpacked data for WithdrawnNFT events raised by the LockerContract contract.
type LockerContractWithdrawnNFTIterator struct {
	Event *LockerContractWithdrawnNFT // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockerContractWithdrawnNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockerContractWithdrawnNFT)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockerContractWithdrawnNFT)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockerContractWithdrawnNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockerContractWithdrawnNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockerContractWithdrawnNFT represents a WithdrawnNFT event raised by the LockerContract contract.
type LockerContractWithdrawnNFT struct {
	Timestamp        *big.Int
	LockedNFTDetails LockerLockedNFTDetails
	Status           string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnNFT is a free log retrieval operation binding the contract event 0x15efd0a750b37ae3ed76ada1c07d105132cc8fcdf626928b820ffd460ce70131.
//
// Solidity: event WithdrawnNFT(uint256 indexed _timestamp, (uint256,address,address,address,string,uint8,string,string) _lockedNFTDetails, string status)
func (_LockerContract *LockerContractFilterer) FilterWithdrawnNFT(opts *bind.FilterOpts, _timestamp []*big.Int) (*LockerContractWithdrawnNFTIterator, error) {

	var _timestampRule []interface{}
	for _, _timestampItem := range _timestamp {
		_timestampRule = append(_timestampRule, _timestampItem)
	}

	logs, sub, err := _LockerContract.contract.FilterLogs(opts, "WithdrawnNFT", _timestampRule)
	if err != nil {
		return nil, err
	}
	return &LockerContractWithdrawnNFTIterator{contract: _LockerContract.contract, event: "WithdrawnNFT", logs: logs, sub: sub}, nil
}

// WatchWithdrawnNFT is a free log subscription operation binding the contract event 0x15efd0a750b37ae3ed76ada1c07d105132cc8fcdf626928b820ffd460ce70131.
//
// Solidity: event WithdrawnNFT(uint256 indexed _timestamp, (uint256,address,address,address,string,uint8,string,string) _lockedNFTDetails, string status)
func (_LockerContract *LockerContractFilterer) WatchWithdrawnNFT(opts *bind.WatchOpts, sink chan<- *LockerContractWithdrawnNFT, _timestamp []*big.Int) (event.Subscription, error) {

	var _timestampRule []interface{}
	for _, _timestampItem := range _timestamp {
		_timestampRule = append(_timestampRule, _timestampItem)
	}

	logs, sub, err := _LockerContract.contract.WatchLogs(opts, "WithdrawnNFT", _timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockerContractWithdrawnNFT)
				if err := _LockerContract.contract.UnpackLog(event, "WithdrawnNFT", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawnNFT is a log parse operation binding the contract event 0x15efd0a750b37ae3ed76ada1c07d105132cc8fcdf626928b820ffd460ce70131.
//
// Solidity: event WithdrawnNFT(uint256 indexed _timestamp, (uint256,address,address,address,string,uint8,string,string) _lockedNFTDetails, string status)
func (_LockerContract *LockerContractFilterer) ParseWithdrawnNFT(log types.Log) (*LockerContractWithdrawnNFT, error) {
	event := new(LockerContractWithdrawnNFT)
	if err := _LockerContract.contract.UnpackLog(event, "WithdrawnNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
