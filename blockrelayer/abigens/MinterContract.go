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

// MinterNFTDetails is an auto generated low-level Go binding around an user-defined struct.
type MinterNFTDetails struct {
	NftId                         *big.Int
	DestinationAddress            common.Address
	SourceAddress                 common.Address
	ChainName                     string
	SourceTokenId                 *big.Int
	DestinationTokenId            *big.Int
	ContractAddress               common.Address
	ReceipentAddressInSourceChain common.Address
	Symbol                        string
	Status                        uint8
	TokenURI                      string
}

// MinterContractMetaData contains all meta data concerning the MinterContract contract.
var MinterContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sourceTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationTokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receipentAddressInSourceChain\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"enumMinter.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structMinter.NFTDetails\",\"name\":\"nftDetails\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventType\",\"type\":\"string\"}],\"name\":\"BurnInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sourceTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationTokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receipentAddressInSourceChain\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"enumMinter.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structMinter.NFTDetails\",\"name\":\"nftDetails\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventType\",\"type\":\"string\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sourceTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationTokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receipentAddressInSourceChain\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"enumMinter.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structMinter.NFTDetails\",\"name\":\"nftDetails\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventType\",\"type\":\"string\"}],\"name\":\"MintInitiate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sourceTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationTokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receipentAddressInSourceChain\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"enumMinter.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structMinter.NFTDetails\",\"name\":\"nftDetails\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventType\",\"type\":\"string\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sTokenId\",\"type\":\"uint256\"}],\"name\":\"finalizeBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sTokenId\",\"type\":\"uint256\"}],\"name\":\"revertBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sTokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"}],\"name\":\"sendBackNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sTokenId\",\"type\":\"uint256\"}],\"name\":\"startMintNFT\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sourceTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationTokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receipentAddressInSourceChain\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"enumMinter.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structMinter.NFTDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIdToNFTDetailsMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sourceTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationTokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receipentAddressInSourceChain\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"enumMinter.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MinterContractABI is the input ABI used to generate the binding from.
// Deprecated: Use MinterContractMetaData.ABI instead.
var MinterContractABI = MinterContractMetaData.ABI

// MinterContract is an auto generated Go binding around an Ethereum contract.
type MinterContract struct {
	MinterContractCaller     // Read-only binding to the contract
	MinterContractTransactor // Write-only binding to the contract
	MinterContractFilterer   // Log filterer for contract events
}

// MinterContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinterContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinterContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinterContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinterContractSession struct {
	Contract     *MinterContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinterContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinterContractCallerSession struct {
	Contract *MinterContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MinterContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinterContractTransactorSession struct {
	Contract     *MinterContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MinterContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinterContractRaw struct {
	Contract *MinterContract // Generic contract binding to access the raw methods on
}

// MinterContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinterContractCallerRaw struct {
	Contract *MinterContractCaller // Generic read-only contract binding to access the raw methods on
}

// MinterContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinterContractTransactorRaw struct {
	Contract *MinterContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinterContract creates a new instance of MinterContract, bound to a specific deployed contract.
func NewMinterContract(address common.Address, backend bind.ContractBackend) (*MinterContract, error) {
	contract, err := bindMinterContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinterContract{MinterContractCaller: MinterContractCaller{contract: contract}, MinterContractTransactor: MinterContractTransactor{contract: contract}, MinterContractFilterer: MinterContractFilterer{contract: contract}}, nil
}

// NewMinterContractCaller creates a new read-only instance of MinterContract, bound to a specific deployed contract.
func NewMinterContractCaller(address common.Address, caller bind.ContractCaller) (*MinterContractCaller, error) {
	contract, err := bindMinterContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinterContractCaller{contract: contract}, nil
}

// NewMinterContractTransactor creates a new write-only instance of MinterContract, bound to a specific deployed contract.
func NewMinterContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MinterContractTransactor, error) {
	contract, err := bindMinterContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinterContractTransactor{contract: contract}, nil
}

// NewMinterContractFilterer creates a new log filterer instance of MinterContract, bound to a specific deployed contract.
func NewMinterContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MinterContractFilterer, error) {
	contract, err := bindMinterContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinterContractFilterer{contract: contract}, nil
}

// bindMinterContract binds a generic wrapper to an already deployed contract.
func bindMinterContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MinterContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterContract *MinterContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinterContract.Contract.MinterContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterContract *MinterContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterContract.Contract.MinterContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterContract *MinterContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterContract.Contract.MinterContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterContract *MinterContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinterContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterContract *MinterContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterContract *MinterContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterContract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MinterContract *MinterContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MinterContract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MinterContract *MinterContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MinterContract.Contract.BalanceOf(&_MinterContract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MinterContract *MinterContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MinterContract.Contract.BalanceOf(&_MinterContract.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_MinterContract *MinterContractCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MinterContract.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_MinterContract *MinterContractSession) BaseURI() (string, error) {
	return _MinterContract.Contract.BaseURI(&_MinterContract.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_MinterContract *MinterContractCallerSession) BaseURI() (string, error) {
	return _MinterContract.Contract.BaseURI(&_MinterContract.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MinterContract *MinterContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MinterContract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MinterContract *MinterContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MinterContract.Contract.GetApproved(&_MinterContract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MinterContract *MinterContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MinterContract.Contract.GetApproved(&_MinterContract.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MinterContract *MinterContractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MinterContract.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MinterContract *MinterContractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MinterContract.Contract.IsApprovedForAll(&_MinterContract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MinterContract *MinterContractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MinterContract.Contract.IsApprovedForAll(&_MinterContract.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MinterContract *MinterContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MinterContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MinterContract *MinterContractSession) Name() (string, error) {
	return _MinterContract.Contract.Name(&_MinterContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MinterContract *MinterContractCallerSession) Name() (string, error) {
	return _MinterContract.Contract.Name(&_MinterContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MinterContract *MinterContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinterContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MinterContract *MinterContractSession) Owner() (common.Address, error) {
	return _MinterContract.Contract.Owner(&_MinterContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MinterContract *MinterContractCallerSession) Owner() (common.Address, error) {
	return _MinterContract.Contract.Owner(&_MinterContract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MinterContract *MinterContractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MinterContract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MinterContract *MinterContractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MinterContract.Contract.OwnerOf(&_MinterContract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MinterContract *MinterContractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MinterContract.Contract.OwnerOf(&_MinterContract.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MinterContract *MinterContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MinterContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MinterContract *MinterContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MinterContract.Contract.SupportsInterface(&_MinterContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MinterContract *MinterContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MinterContract.Contract.SupportsInterface(&_MinterContract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MinterContract *MinterContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MinterContract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MinterContract *MinterContractSession) Symbol() (string, error) {
	return _MinterContract.Contract.Symbol(&_MinterContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MinterContract *MinterContractCallerSession) Symbol() (string, error) {
	return _MinterContract.Contract.Symbol(&_MinterContract.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_MinterContract *MinterContractCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MinterContract.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_MinterContract *MinterContractSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _MinterContract.Contract.TokenByIndex(&_MinterContract.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_MinterContract *MinterContractCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _MinterContract.Contract.TokenByIndex(&_MinterContract.CallOpts, index)
}

// TokenIdToNFTDetailsMapping is a free data retrieval call binding the contract method 0x17135bb1.
//
// Solidity: function tokenIdToNFTDetailsMapping(address , uint256 ) view returns(uint256 nftId, address destinationAddress, address sourceAddress, string chainName, uint256 sourceTokenId, uint256 destinationTokenId, address contractAddress, address receipentAddressInSourceChain, string symbol, uint8 status, string tokenURI)
func (_MinterContract *MinterContractCaller) TokenIdToNFTDetailsMapping(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	NftId                         *big.Int
	DestinationAddress            common.Address
	SourceAddress                 common.Address
	ChainName                     string
	SourceTokenId                 *big.Int
	DestinationTokenId            *big.Int
	ContractAddress               common.Address
	ReceipentAddressInSourceChain common.Address
	Symbol                        string
	Status                        uint8
	TokenURI                      string
}, error) {
	var out []interface{}
	err := _MinterContract.contract.Call(opts, &out, "tokenIdToNFTDetailsMapping", arg0, arg1)

	outstruct := new(struct {
		NftId                         *big.Int
		DestinationAddress            common.Address
		SourceAddress                 common.Address
		ChainName                     string
		SourceTokenId                 *big.Int
		DestinationTokenId            *big.Int
		ContractAddress               common.Address
		ReceipentAddressInSourceChain common.Address
		Symbol                        string
		Status                        uint8
		TokenURI                      string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NftId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DestinationAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.SourceAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ChainName = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.SourceTokenId = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.DestinationTokenId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ContractAddress = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.ReceipentAddressInSourceChain = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.Symbol = *abi.ConvertType(out[8], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[9], new(uint8)).(*uint8)
	outstruct.TokenURI = *abi.ConvertType(out[10], new(string)).(*string)

	return *outstruct, err

}

// TokenIdToNFTDetailsMapping is a free data retrieval call binding the contract method 0x17135bb1.
//
// Solidity: function tokenIdToNFTDetailsMapping(address , uint256 ) view returns(uint256 nftId, address destinationAddress, address sourceAddress, string chainName, uint256 sourceTokenId, uint256 destinationTokenId, address contractAddress, address receipentAddressInSourceChain, string symbol, uint8 status, string tokenURI)
func (_MinterContract *MinterContractSession) TokenIdToNFTDetailsMapping(arg0 common.Address, arg1 *big.Int) (struct {
	NftId                         *big.Int
	DestinationAddress            common.Address
	SourceAddress                 common.Address
	ChainName                     string
	SourceTokenId                 *big.Int
	DestinationTokenId            *big.Int
	ContractAddress               common.Address
	ReceipentAddressInSourceChain common.Address
	Symbol                        string
	Status                        uint8
	TokenURI                      string
}, error) {
	return _MinterContract.Contract.TokenIdToNFTDetailsMapping(&_MinterContract.CallOpts, arg0, arg1)
}

// TokenIdToNFTDetailsMapping is a free data retrieval call binding the contract method 0x17135bb1.
//
// Solidity: function tokenIdToNFTDetailsMapping(address , uint256 ) view returns(uint256 nftId, address destinationAddress, address sourceAddress, string chainName, uint256 sourceTokenId, uint256 destinationTokenId, address contractAddress, address receipentAddressInSourceChain, string symbol, uint8 status, string tokenURI)
func (_MinterContract *MinterContractCallerSession) TokenIdToNFTDetailsMapping(arg0 common.Address, arg1 *big.Int) (struct {
	NftId                         *big.Int
	DestinationAddress            common.Address
	SourceAddress                 common.Address
	ChainName                     string
	SourceTokenId                 *big.Int
	DestinationTokenId            *big.Int
	ContractAddress               common.Address
	ReceipentAddressInSourceChain common.Address
	Symbol                        string
	Status                        uint8
	TokenURI                      string
}, error) {
	return _MinterContract.Contract.TokenIdToNFTDetailsMapping(&_MinterContract.CallOpts, arg0, arg1)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_MinterContract *MinterContractCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MinterContract.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_MinterContract *MinterContractSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _MinterContract.Contract.TokenOfOwnerByIndex(&_MinterContract.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_MinterContract *MinterContractCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _MinterContract.Contract.TokenOfOwnerByIndex(&_MinterContract.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MinterContract *MinterContractCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _MinterContract.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MinterContract *MinterContractSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MinterContract.Contract.TokenURI(&_MinterContract.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MinterContract *MinterContractCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MinterContract.Contract.TokenURI(&_MinterContract.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MinterContract *MinterContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinterContract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MinterContract *MinterContractSession) TotalSupply() (*big.Int, error) {
	return _MinterContract.Contract.TotalSupply(&_MinterContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MinterContract *MinterContractCallerSession) TotalSupply() (*big.Int, error) {
	return _MinterContract.Contract.TotalSupply(&_MinterContract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MinterContract *MinterContractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MinterContract *MinterContractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.Contract.Approve(&_MinterContract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MinterContract *MinterContractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.Contract.Approve(&_MinterContract.TransactOpts, to, tokenId)
}

// FinalizeBurn is a paid mutator transaction binding the contract method 0xa8fd0092.
//
// Solidity: function finalizeBurn(address nftContractAddress, uint256 sTokenId) returns()
func (_MinterContract *MinterContractTransactor) FinalizeBurn(opts *bind.TransactOpts, nftContractAddress common.Address, sTokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.contract.Transact(opts, "finalizeBurn", nftContractAddress, sTokenId)
}

// FinalizeBurn is a paid mutator transaction binding the contract method 0xa8fd0092.
//
// Solidity: function finalizeBurn(address nftContractAddress, uint256 sTokenId) returns()
func (_MinterContract *MinterContractSession) FinalizeBurn(nftContractAddress common.Address, sTokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.Contract.FinalizeBurn(&_MinterContract.TransactOpts, nftContractAddress, sTokenId)
}

// FinalizeBurn is a paid mutator transaction binding the contract method 0xa8fd0092.
//
// Solidity: function finalizeBurn(address nftContractAddress, uint256 sTokenId) returns()
func (_MinterContract *MinterContractTransactorSession) FinalizeBurn(nftContractAddress common.Address, sTokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.Contract.FinalizeBurn(&_MinterContract.TransactOpts, nftContractAddress, sTokenId)
}

// RevertBurn is a paid mutator transaction binding the contract method 0xb2d02d39.
//
// Solidity: function revertBurn(address nftContractAddress, uint256 sTokenId) returns()
func (_MinterContract *MinterContractTransactor) RevertBurn(opts *bind.TransactOpts, nftContractAddress common.Address, sTokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.contract.Transact(opts, "revertBurn", nftContractAddress, sTokenId)
}

// RevertBurn is a paid mutator transaction binding the contract method 0xb2d02d39.
//
// Solidity: function revertBurn(address nftContractAddress, uint256 sTokenId) returns()
func (_MinterContract *MinterContractSession) RevertBurn(nftContractAddress common.Address, sTokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.Contract.RevertBurn(&_MinterContract.TransactOpts, nftContractAddress, sTokenId)
}

// RevertBurn is a paid mutator transaction binding the contract method 0xb2d02d39.
//
// Solidity: function revertBurn(address nftContractAddress, uint256 sTokenId) returns()
func (_MinterContract *MinterContractTransactorSession) RevertBurn(nftContractAddress common.Address, sTokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.Contract.RevertBurn(&_MinterContract.TransactOpts, nftContractAddress, sTokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MinterContract *MinterContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MinterContract *MinterContractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.Contract.SafeTransferFrom(&_MinterContract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MinterContract *MinterContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.Contract.SafeTransferFrom(&_MinterContract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_MinterContract *MinterContractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MinterContract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_MinterContract *MinterContractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MinterContract.Contract.SafeTransferFrom0(&_MinterContract.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_MinterContract *MinterContractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MinterContract.Contract.SafeTransferFrom0(&_MinterContract.TransactOpts, from, to, tokenId, _data)
}

// SendBackNFT is a paid mutator transaction binding the contract method 0x9105d56d.
//
// Solidity: function sendBackNFT(address nftContractAddress, uint256 sTokenId, address receiverAddress) returns()
func (_MinterContract *MinterContractTransactor) SendBackNFT(opts *bind.TransactOpts, nftContractAddress common.Address, sTokenId *big.Int, receiverAddress common.Address) (*types.Transaction, error) {
	return _MinterContract.contract.Transact(opts, "sendBackNFT", nftContractAddress, sTokenId, receiverAddress)
}

// SendBackNFT is a paid mutator transaction binding the contract method 0x9105d56d.
//
// Solidity: function sendBackNFT(address nftContractAddress, uint256 sTokenId, address receiverAddress) returns()
func (_MinterContract *MinterContractSession) SendBackNFT(nftContractAddress common.Address, sTokenId *big.Int, receiverAddress common.Address) (*types.Transaction, error) {
	return _MinterContract.Contract.SendBackNFT(&_MinterContract.TransactOpts, nftContractAddress, sTokenId, receiverAddress)
}

// SendBackNFT is a paid mutator transaction binding the contract method 0x9105d56d.
//
// Solidity: function sendBackNFT(address nftContractAddress, uint256 sTokenId, address receiverAddress) returns()
func (_MinterContract *MinterContractTransactorSession) SendBackNFT(nftContractAddress common.Address, sTokenId *big.Int, receiverAddress common.Address) (*types.Transaction, error) {
	return _MinterContract.Contract.SendBackNFT(&_MinterContract.TransactOpts, nftContractAddress, sTokenId, receiverAddress)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MinterContract *MinterContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _MinterContract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MinterContract *MinterContractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MinterContract.Contract.SetApprovalForAll(&_MinterContract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MinterContract *MinterContractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MinterContract.Contract.SetApprovalForAll(&_MinterContract.TransactOpts, operator, approved)
}

// StartMintNFT is a paid mutator transaction binding the contract method 0xe35d21aa.
//
// Solidity: function startMintNFT(address sourceAddress, address destinationAddress, string chainName, string tokenURI, string symbol, address nftContractAddress, uint256 sTokenId) returns((uint256,address,address,string,uint256,uint256,address,address,string,uint8,string))
func (_MinterContract *MinterContractTransactor) StartMintNFT(opts *bind.TransactOpts, sourceAddress common.Address, destinationAddress common.Address, chainName string, tokenURI string, symbol string, nftContractAddress common.Address, sTokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.contract.Transact(opts, "startMintNFT", sourceAddress, destinationAddress, chainName, tokenURI, symbol, nftContractAddress, sTokenId)
}

// StartMintNFT is a paid mutator transaction binding the contract method 0xe35d21aa.
//
// Solidity: function startMintNFT(address sourceAddress, address destinationAddress, string chainName, string tokenURI, string symbol, address nftContractAddress, uint256 sTokenId) returns((uint256,address,address,string,uint256,uint256,address,address,string,uint8,string))
func (_MinterContract *MinterContractSession) StartMintNFT(sourceAddress common.Address, destinationAddress common.Address, chainName string, tokenURI string, symbol string, nftContractAddress common.Address, sTokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.Contract.StartMintNFT(&_MinterContract.TransactOpts, sourceAddress, destinationAddress, chainName, tokenURI, symbol, nftContractAddress, sTokenId)
}

// StartMintNFT is a paid mutator transaction binding the contract method 0xe35d21aa.
//
// Solidity: function startMintNFT(address sourceAddress, address destinationAddress, string chainName, string tokenURI, string symbol, address nftContractAddress, uint256 sTokenId) returns((uint256,address,address,string,uint256,uint256,address,address,string,uint8,string))
func (_MinterContract *MinterContractTransactorSession) StartMintNFT(sourceAddress common.Address, destinationAddress common.Address, chainName string, tokenURI string, symbol string, nftContractAddress common.Address, sTokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.Contract.StartMintNFT(&_MinterContract.TransactOpts, sourceAddress, destinationAddress, chainName, tokenURI, symbol, nftContractAddress, sTokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MinterContract *MinterContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MinterContract *MinterContractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.Contract.TransferFrom(&_MinterContract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MinterContract *MinterContractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MinterContract.Contract.TransferFrom(&_MinterContract.TransactOpts, from, to, tokenId)
}

// MinterContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MinterContract contract.
type MinterContractApprovalIterator struct {
	Event *MinterContractApproval // Event containing the contract specifics and raw log

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
func (it *MinterContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterContractApproval)
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
		it.Event = new(MinterContractApproval)
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
func (it *MinterContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterContractApproval represents a Approval event raised by the MinterContract contract.
type MinterContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MinterContract *MinterContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MinterContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MinterContract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MinterContractApprovalIterator{contract: _MinterContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MinterContract *MinterContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MinterContractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MinterContract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterContractApproval)
				if err := _MinterContract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MinterContract *MinterContractFilterer) ParseApproval(log types.Log) (*MinterContractApproval, error) {
	event := new(MinterContractApproval)
	if err := _MinterContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MinterContract contract.
type MinterContractApprovalForAllIterator struct {
	Event *MinterContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MinterContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterContractApprovalForAll)
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
		it.Event = new(MinterContractApprovalForAll)
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
func (it *MinterContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterContractApprovalForAll represents a ApprovalForAll event raised by the MinterContract contract.
type MinterContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MinterContract *MinterContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MinterContractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MinterContract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MinterContractApprovalForAllIterator{contract: _MinterContract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MinterContract *MinterContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MinterContractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MinterContract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterContractApprovalForAll)
				if err := _MinterContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MinterContract *MinterContractFilterer) ParseApprovalForAll(log types.Log) (*MinterContractApprovalForAll, error) {
	event := new(MinterContractApprovalForAll)
	if err := _MinterContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterContractBurnInitiatedIterator is returned from FilterBurnInitiated and is used to iterate over the raw logs and unpacked data for BurnInitiated events raised by the MinterContract contract.
type MinterContractBurnInitiatedIterator struct {
	Event *MinterContractBurnInitiated // Event containing the contract specifics and raw log

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
func (it *MinterContractBurnInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterContractBurnInitiated)
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
		it.Event = new(MinterContractBurnInitiated)
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
func (it *MinterContractBurnInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterContractBurnInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterContractBurnInitiated represents a BurnInitiated event raised by the MinterContract contract.
type MinterContractBurnInitiated struct {
	Date         *big.Int
	ContractAddr common.Address
	STokenId     *big.Int
	NftId        *big.Int
	NftDetails   MinterNFTDetails
	EventType    string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBurnInitiated is a free log retrieval operation binding the contract event 0x49ab92467173ef42c440da13d0811d07b0cdb4dd4a5fd1e41f7ef99d19af0c73.
//
// Solidity: event BurnInitiated(uint256 indexed date, address contractAddr, uint256 sTokenId, uint256 nftId, (uint256,address,address,string,uint256,uint256,address,address,string,uint8,string) nftDetails, string eventType)
func (_MinterContract *MinterContractFilterer) FilterBurnInitiated(opts *bind.FilterOpts, date []*big.Int) (*MinterContractBurnInitiatedIterator, error) {

	var dateRule []interface{}
	for _, dateItem := range date {
		dateRule = append(dateRule, dateItem)
	}

	logs, sub, err := _MinterContract.contract.FilterLogs(opts, "BurnInitiated", dateRule)
	if err != nil {
		return nil, err
	}
	return &MinterContractBurnInitiatedIterator{contract: _MinterContract.contract, event: "BurnInitiated", logs: logs, sub: sub}, nil
}

// WatchBurnInitiated is a free log subscription operation binding the contract event 0x49ab92467173ef42c440da13d0811d07b0cdb4dd4a5fd1e41f7ef99d19af0c73.
//
// Solidity: event BurnInitiated(uint256 indexed date, address contractAddr, uint256 sTokenId, uint256 nftId, (uint256,address,address,string,uint256,uint256,address,address,string,uint8,string) nftDetails, string eventType)
func (_MinterContract *MinterContractFilterer) WatchBurnInitiated(opts *bind.WatchOpts, sink chan<- *MinterContractBurnInitiated, date []*big.Int) (event.Subscription, error) {

	var dateRule []interface{}
	for _, dateItem := range date {
		dateRule = append(dateRule, dateItem)
	}

	logs, sub, err := _MinterContract.contract.WatchLogs(opts, "BurnInitiated", dateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterContractBurnInitiated)
				if err := _MinterContract.contract.UnpackLog(event, "BurnInitiated", log); err != nil {
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

// ParseBurnInitiated is a log parse operation binding the contract event 0x49ab92467173ef42c440da13d0811d07b0cdb4dd4a5fd1e41f7ef99d19af0c73.
//
// Solidity: event BurnInitiated(uint256 indexed date, address contractAddr, uint256 sTokenId, uint256 nftId, (uint256,address,address,string,uint256,uint256,address,address,string,uint8,string) nftDetails, string eventType)
func (_MinterContract *MinterContractFilterer) ParseBurnInitiated(log types.Log) (*MinterContractBurnInitiated, error) {
	event := new(MinterContractBurnInitiated)
	if err := _MinterContract.contract.UnpackLog(event, "BurnInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterContractBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the MinterContract contract.
type MinterContractBurnedIterator struct {
	Event *MinterContractBurned // Event containing the contract specifics and raw log

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
func (it *MinterContractBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterContractBurned)
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
		it.Event = new(MinterContractBurned)
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
func (it *MinterContractBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterContractBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterContractBurned represents a Burned event raised by the MinterContract contract.
type MinterContractBurned struct {
	Date         *big.Int
	ContractAddr common.Address
	STokenId     *big.Int
	NftId        *big.Int
	NftDetails   MinterNFTDetails
	EventType    string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0xf9b52931c593264eb546fca11f94dfed22b5a39cdd206f3296aa110bd915e35d.
//
// Solidity: event Burned(uint256 indexed date, address contractAddr, uint256 sTokenId, uint256 nftId, (uint256,address,address,string,uint256,uint256,address,address,string,uint8,string) nftDetails, string eventType)
func (_MinterContract *MinterContractFilterer) FilterBurned(opts *bind.FilterOpts, date []*big.Int) (*MinterContractBurnedIterator, error) {

	var dateRule []interface{}
	for _, dateItem := range date {
		dateRule = append(dateRule, dateItem)
	}

	logs, sub, err := _MinterContract.contract.FilterLogs(opts, "Burned", dateRule)
	if err != nil {
		return nil, err
	}
	return &MinterContractBurnedIterator{contract: _MinterContract.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0xf9b52931c593264eb546fca11f94dfed22b5a39cdd206f3296aa110bd915e35d.
//
// Solidity: event Burned(uint256 indexed date, address contractAddr, uint256 sTokenId, uint256 nftId, (uint256,address,address,string,uint256,uint256,address,address,string,uint8,string) nftDetails, string eventType)
func (_MinterContract *MinterContractFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *MinterContractBurned, date []*big.Int) (event.Subscription, error) {

	var dateRule []interface{}
	for _, dateItem := range date {
		dateRule = append(dateRule, dateItem)
	}

	logs, sub, err := _MinterContract.contract.WatchLogs(opts, "Burned", dateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterContractBurned)
				if err := _MinterContract.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0xf9b52931c593264eb546fca11f94dfed22b5a39cdd206f3296aa110bd915e35d.
//
// Solidity: event Burned(uint256 indexed date, address contractAddr, uint256 sTokenId, uint256 nftId, (uint256,address,address,string,uint256,uint256,address,address,string,uint8,string) nftDetails, string eventType)
func (_MinterContract *MinterContractFilterer) ParseBurned(log types.Log) (*MinterContractBurned, error) {
	event := new(MinterContractBurned)
	if err := _MinterContract.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterContractMintInitiateIterator is returned from FilterMintInitiate and is used to iterate over the raw logs and unpacked data for MintInitiate events raised by the MinterContract contract.
type MinterContractMintInitiateIterator struct {
	Event *MinterContractMintInitiate // Event containing the contract specifics and raw log

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
func (it *MinterContractMintInitiateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterContractMintInitiate)
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
		it.Event = new(MinterContractMintInitiate)
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
func (it *MinterContractMintInitiateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterContractMintInitiateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterContractMintInitiate represents a MintInitiate event raised by the MinterContract contract.
type MinterContractMintInitiate struct {
	Date         *big.Int
	ContractAddr common.Address
	STokenId     *big.Int
	NftId        *big.Int
	NftDetails   MinterNFTDetails
	EventType    string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintInitiate is a free log retrieval operation binding the contract event 0x8e6dfa8bf81b5a84bbc858fd12cb421ba141075b33357affc3dc24a20ab2e452.
//
// Solidity: event MintInitiate(uint256 indexed date, address contractAddr, uint256 sTokenId, uint256 nftId, (uint256,address,address,string,uint256,uint256,address,address,string,uint8,string) nftDetails, string eventType)
func (_MinterContract *MinterContractFilterer) FilterMintInitiate(opts *bind.FilterOpts, date []*big.Int) (*MinterContractMintInitiateIterator, error) {

	var dateRule []interface{}
	for _, dateItem := range date {
		dateRule = append(dateRule, dateItem)
	}

	logs, sub, err := _MinterContract.contract.FilterLogs(opts, "MintInitiate", dateRule)
	if err != nil {
		return nil, err
	}
	return &MinterContractMintInitiateIterator{contract: _MinterContract.contract, event: "MintInitiate", logs: logs, sub: sub}, nil
}

// WatchMintInitiate is a free log subscription operation binding the contract event 0x8e6dfa8bf81b5a84bbc858fd12cb421ba141075b33357affc3dc24a20ab2e452.
//
// Solidity: event MintInitiate(uint256 indexed date, address contractAddr, uint256 sTokenId, uint256 nftId, (uint256,address,address,string,uint256,uint256,address,address,string,uint8,string) nftDetails, string eventType)
func (_MinterContract *MinterContractFilterer) WatchMintInitiate(opts *bind.WatchOpts, sink chan<- *MinterContractMintInitiate, date []*big.Int) (event.Subscription, error) {

	var dateRule []interface{}
	for _, dateItem := range date {
		dateRule = append(dateRule, dateItem)
	}

	logs, sub, err := _MinterContract.contract.WatchLogs(opts, "MintInitiate", dateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterContractMintInitiate)
				if err := _MinterContract.contract.UnpackLog(event, "MintInitiate", log); err != nil {
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

// ParseMintInitiate is a log parse operation binding the contract event 0x8e6dfa8bf81b5a84bbc858fd12cb421ba141075b33357affc3dc24a20ab2e452.
//
// Solidity: event MintInitiate(uint256 indexed date, address contractAddr, uint256 sTokenId, uint256 nftId, (uint256,address,address,string,uint256,uint256,address,address,string,uint8,string) nftDetails, string eventType)
func (_MinterContract *MinterContractFilterer) ParseMintInitiate(log types.Log) (*MinterContractMintInitiate, error) {
	event := new(MinterContractMintInitiate)
	if err := _MinterContract.contract.UnpackLog(event, "MintInitiate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterContractMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the MinterContract contract.
type MinterContractMintedIterator struct {
	Event *MinterContractMinted // Event containing the contract specifics and raw log

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
func (it *MinterContractMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterContractMinted)
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
		it.Event = new(MinterContractMinted)
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
func (it *MinterContractMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterContractMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterContractMinted represents a Minted event raised by the MinterContract contract.
type MinterContractMinted struct {
	Date         *big.Int
	ContractAddr common.Address
	STokenId     *big.Int
	NftId        *big.Int
	NftDetails   MinterNFTDetails
	EventType    string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x5c8e25a864ce27695a18afdfa87a17f90d368c021a59e56ae17d9cb02e7d718d.
//
// Solidity: event Minted(uint256 indexed date, address contractAddr, uint256 sTokenId, uint256 nftId, (uint256,address,address,string,uint256,uint256,address,address,string,uint8,string) nftDetails, string eventType)
func (_MinterContract *MinterContractFilterer) FilterMinted(opts *bind.FilterOpts, date []*big.Int) (*MinterContractMintedIterator, error) {

	var dateRule []interface{}
	for _, dateItem := range date {
		dateRule = append(dateRule, dateItem)
	}

	logs, sub, err := _MinterContract.contract.FilterLogs(opts, "Minted", dateRule)
	if err != nil {
		return nil, err
	}
	return &MinterContractMintedIterator{contract: _MinterContract.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x5c8e25a864ce27695a18afdfa87a17f90d368c021a59e56ae17d9cb02e7d718d.
//
// Solidity: event Minted(uint256 indexed date, address contractAddr, uint256 sTokenId, uint256 nftId, (uint256,address,address,string,uint256,uint256,address,address,string,uint8,string) nftDetails, string eventType)
func (_MinterContract *MinterContractFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *MinterContractMinted, date []*big.Int) (event.Subscription, error) {

	var dateRule []interface{}
	for _, dateItem := range date {
		dateRule = append(dateRule, dateItem)
	}

	logs, sub, err := _MinterContract.contract.WatchLogs(opts, "Minted", dateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterContractMinted)
				if err := _MinterContract.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x5c8e25a864ce27695a18afdfa87a17f90d368c021a59e56ae17d9cb02e7d718d.
//
// Solidity: event Minted(uint256 indexed date, address contractAddr, uint256 sTokenId, uint256 nftId, (uint256,address,address,string,uint256,uint256,address,address,string,uint8,string) nftDetails, string eventType)
func (_MinterContract *MinterContractFilterer) ParseMinted(log types.Log) (*MinterContractMinted, error) {
	event := new(MinterContractMinted)
	if err := _MinterContract.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MinterContract contract.
type MinterContractTransferIterator struct {
	Event *MinterContractTransfer // Event containing the contract specifics and raw log

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
func (it *MinterContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterContractTransfer)
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
		it.Event = new(MinterContractTransfer)
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
func (it *MinterContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterContractTransfer represents a Transfer event raised by the MinterContract contract.
type MinterContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MinterContract *MinterContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MinterContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MinterContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MinterContractTransferIterator{contract: _MinterContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MinterContract *MinterContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MinterContractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MinterContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterContractTransfer)
				if err := _MinterContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MinterContract *MinterContractFilterer) ParseTransfer(log types.Log) (*MinterContractTransfer, error) {
	event := new(MinterContractTransfer)
	if err := _MinterContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
