// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eiptwostore

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
)

// EiptwostoreMetaData contains all meta data concerning the Eiptwostore contract.
var EiptwostoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200154638038062001546833981810160405281019062000037919062000193565b81600390805190602001906200004f92919062000071565b5080600490805190602001906200006892919062000071565b50505062000337565b8280546200007f90620002a3565b90600052602060002090601f016020900481019282620000a35760008555620000ef565b82601f10620000be57805160ff1916838001178555620000ef565b82800160010185558215620000ef579182015b82811115620000ee578251825591602001919060010190620000d1565b5b509050620000fe919062000102565b5090565b5b808211156200011d57600081600090555060010162000103565b5090565b60006200013862000132846200023a565b62000206565b9050828152602081018484840111156200015157600080fd5b6200015e8482856200026d565b509392505050565b600082601f8301126200017857600080fd5b81516200018a84826020860162000121565b91505092915050565b60008060408385031215620001a757600080fd5b600083015167ffffffffffffffff811115620001c257600080fd5b620001d08582860162000166565b925050602083015167ffffffffffffffff811115620001ee57600080fd5b620001fc8582860162000166565b9150509250929050565b6000604051905081810181811067ffffffffffffffff8211171562000230576200022f62000308565b5b8060405250919050565b600067ffffffffffffffff82111562000258576200025762000308565b5b601f19601f8301169050602081019050919050565b60005b838110156200028d57808201518184015260208101905062000270565b838111156200029d576000848401525b50505050565b60006002820490506001821680620002bc57607f821691505b60208210811415620002d357620002d2620002d9565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6111ff80620003476000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461016857806370a082311461019857806395d89b41146101c8578063a457c2d7146101e6578063a9059cbb14610216578063dd62ed3e14610246576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100fc57806323b872dd1461011a578063313ce5671461014a575b600080fd5b6100b6610276565b6040516100c39190610ec8565b60405180910390f35b6100e660048036038101906100e19190610b67565b610308565b6040516100f39190610ead565b60405180910390f35b61010461032b565b6040516101119190610fca565b60405180910390f35b610134600480360381019061012f9190610b18565b610335565b6040516101419190610ead565b60405180910390f35b610152610364565b60405161015f9190610fe5565b60405180910390f35b610182600480360381019061017d9190610b67565b61036d565b60405161018f9190610ead565b60405180910390f35b6101b260048036038101906101ad9190610ab3565b6103a4565b6040516101bf9190610fca565b60405180910390f35b6101d06103ec565b6040516101dd9190610ec8565b60405180910390f35b61020060048036038101906101fb9190610b67565b61047e565b60405161020d9190610ead565b60405180910390f35b610230600480360381019061022b9190610b67565b6104f5565b60405161023d9190610ead565b60405180910390f35b610260600480360381019061025b9190610adc565b610518565b60405161026d9190610fca565b60405180910390f35b606060038054610285906110fa565b80601f01602080910402602001604051908101604052809291908181526020018280546102b1906110fa565b80156102fe5780601f106102d3576101008083540402835291602001916102fe565b820191906000526020600020905b8154815290600101906020018083116102e157829003601f168201915b5050505050905090565b60008061031361059f565b90506103208185856105a7565b600191505092915050565b6000600254905090565b60008061034061059f565b905061034d858285610772565b6103588585856107fe565b60019150509392505050565b60006012905090565b60008061037861059f565b905061039981858561038a8589610518565b610394919061101c565b6105a7565b600191505092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600480546103fb906110fa565b80601f0160208091040260200160405190810160405280929190818152602001828054610427906110fa565b80156104745780601f1061044957610100808354040283529160200191610474565b820191906000526020600020905b81548152906001019060200180831161045757829003601f168201915b5050505050905090565b60008061048961059f565b905060006104978286610518565b9050838110156104dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d390610faa565b60405180910390fd5b6104e982868684036105a7565b60019250505092915050565b60008061050061059f565b905061050d8185856107fe565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610617576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060e90610f8a565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610687576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067e90610f0a565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516107659190610fca565b60405180910390a3505050565b600061077e8484610518565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146107f857818110156107ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e190610f2a565b60405180910390fd5b6107f784848484036105a7565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561086e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086590610f6a565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156108de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d590610eea565b60405180910390fd5b6108e9838383610a7f565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508181101561096f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096690610f4a565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610a02919061101c565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a669190610fca565b60405180910390a3610a79848484610a84565b50505050565b505050565b505050565b600081359050610a988161119b565b92915050565b600081359050610aad816111b2565b92915050565b600060208284031215610ac557600080fd5b6000610ad384828501610a89565b91505092915050565b60008060408385031215610aef57600080fd5b6000610afd85828601610a89565b9250506020610b0e85828601610a89565b9150509250929050565b600080600060608486031215610b2d57600080fd5b6000610b3b86828701610a89565b9350506020610b4c86828701610a89565b9250506040610b5d86828701610a9e565b9150509250925092565b60008060408385031215610b7a57600080fd5b6000610b8885828601610a89565b9250506020610b9985828601610a9e565b9150509250929050565b610bac81611084565b82525050565b6000610bbd82611000565b610bc7818561100b565b9350610bd78185602086016110c7565b610be08161118a565b840191505092915050565b6000610bf860238361100b565b91507f45524332303a207472616e7366657220746f20746865207a65726f206164647260008301527f65737300000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610c5e60228361100b565b91507f45524332303a20617070726f766520746f20746865207a65726f20616464726560008301527f73730000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610cc4601d8361100b565b91507f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006000830152602082019050919050565b6000610d0460268361100b565b91507f45524332303a207472616e7366657220616d6f756e742065786365656473206260008301527f616c616e636500000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610d6a60258361100b565b91507f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008301527f64726573730000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610dd060248361100b565b91507f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008301527f72657373000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610e3660258361100b565b91507f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008301527f207a65726f0000000000000000000000000000000000000000000000000000006020830152604082019050919050565b610e98816110b0565b82525050565b610ea7816110ba565b82525050565b6000602082019050610ec26000830184610ba3565b92915050565b60006020820190508181036000830152610ee28184610bb2565b905092915050565b60006020820190508181036000830152610f0381610beb565b9050919050565b60006020820190508181036000830152610f2381610c51565b9050919050565b60006020820190508181036000830152610f4381610cb7565b9050919050565b60006020820190508181036000830152610f6381610cf7565b9050919050565b60006020820190508181036000830152610f8381610d5d565b9050919050565b60006020820190508181036000830152610fa381610dc3565b9050919050565b60006020820190508181036000830152610fc381610e29565b9050919050565b6000602082019050610fdf6000830184610e8f565b92915050565b6000602082019050610ffa6000830184610e9e565b92915050565b600081519050919050565b600082825260208201905092915050565b6000611027826110b0565b9150611032836110b0565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156110675761106661112c565b5b828201905092915050565b600061107d82611090565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b838110156110e55780820151818401526020810190506110ca565b838111156110f4576000848401525b50505050565b6000600282049050600182168061111257607f821691505b602082108114156111265761112561115b565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000601f19601f8301169050919050565b6111a481611072565b81146111af57600080fd5b50565b6111bb816110b0565b81146111c657600080fd5b5056fea2646970667358221220714429594a8488cb22c59cc605b28c98d0f25ce3e36e054118f5135b9927577b64736f6c63430008000033",
}

// EiptwostoreABI is the input ABI used to generate the binding from.
// Deprecated: Use EiptwostoreMetaData.ABI instead.
var EiptwostoreABI = EiptwostoreMetaData.ABI

// EiptwostoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EiptwostoreMetaData.Bin instead.
var EiptwostoreBin = EiptwostoreMetaData.Bin

// DeployEiptwostore deploys a new Ethereum contract, binding an instance of Eiptwostore to it.
func DeployEiptwostore(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *Eiptwostore, error) {
	parsed, err := EiptwostoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EiptwostoreBin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Eiptwostore{EiptwostoreCaller: EiptwostoreCaller{contract: contract}, EiptwostoreTransactor: EiptwostoreTransactor{contract: contract}, EiptwostoreFilterer: EiptwostoreFilterer{contract: contract}}, nil
}

// Eiptwostore is an auto generated Go binding around an Ethereum contract.
type Eiptwostore struct {
	EiptwostoreCaller     // Read-only binding to the contract
	EiptwostoreTransactor // Write-only binding to the contract
	EiptwostoreFilterer   // Log filterer for contract events
}

// EiptwostoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type EiptwostoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EiptwostoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EiptwostoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EiptwostoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EiptwostoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EiptwostoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EiptwostoreSession struct {
	Contract     *Eiptwostore      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EiptwostoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EiptwostoreCallerSession struct {
	Contract *EiptwostoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EiptwostoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EiptwostoreTransactorSession struct {
	Contract     *EiptwostoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EiptwostoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type EiptwostoreRaw struct {
	Contract *Eiptwostore // Generic contract binding to access the raw methods on
}

// EiptwostoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EiptwostoreCallerRaw struct {
	Contract *EiptwostoreCaller // Generic read-only contract binding to access the raw methods on
}

// EiptwostoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EiptwostoreTransactorRaw struct {
	Contract *EiptwostoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEiptwostore creates a new instance of Eiptwostore, bound to a specific deployed contract.
func NewEiptwostore(address common.Address, backend bind.ContractBackend) (*Eiptwostore, error) {
	contract, err := bindEiptwostore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eiptwostore{EiptwostoreCaller: EiptwostoreCaller{contract: contract}, EiptwostoreTransactor: EiptwostoreTransactor{contract: contract}, EiptwostoreFilterer: EiptwostoreFilterer{contract: contract}}, nil
}

// NewEiptwostoreCaller creates a new read-only instance of Eiptwostore, bound to a specific deployed contract.
func NewEiptwostoreCaller(address common.Address, caller bind.ContractCaller) (*EiptwostoreCaller, error) {
	contract, err := bindEiptwostore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EiptwostoreCaller{contract: contract}, nil
}

// NewEiptwostoreTransactor creates a new write-only instance of Eiptwostore, bound to a specific deployed contract.
func NewEiptwostoreTransactor(address common.Address, transactor bind.ContractTransactor) (*EiptwostoreTransactor, error) {
	contract, err := bindEiptwostore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EiptwostoreTransactor{contract: contract}, nil
}

// NewEiptwostoreFilterer creates a new log filterer instance of Eiptwostore, bound to a specific deployed contract.
func NewEiptwostoreFilterer(address common.Address, filterer bind.ContractFilterer) (*EiptwostoreFilterer, error) {
	contract, err := bindEiptwostore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EiptwostoreFilterer{contract: contract}, nil
}

// bindEiptwostore binds a generic wrapper to an already deployed contract.
func bindEiptwostore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EiptwostoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eiptwostore *EiptwostoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eiptwostore.Contract.EiptwostoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eiptwostore *EiptwostoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eiptwostore.Contract.EiptwostoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eiptwostore *EiptwostoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eiptwostore.Contract.EiptwostoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eiptwostore *EiptwostoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eiptwostore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eiptwostore *EiptwostoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eiptwostore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eiptwostore *EiptwostoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eiptwostore.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Eiptwostore *EiptwostoreCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Eiptwostore.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Eiptwostore *EiptwostoreSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Eiptwostore.Contract.Allowance(&_Eiptwostore.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Eiptwostore *EiptwostoreCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Eiptwostore.Contract.Allowance(&_Eiptwostore.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Eiptwostore *EiptwostoreCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Eiptwostore.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Eiptwostore *EiptwostoreSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Eiptwostore.Contract.BalanceOf(&_Eiptwostore.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Eiptwostore *EiptwostoreCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Eiptwostore.Contract.BalanceOf(&_Eiptwostore.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Eiptwostore *EiptwostoreCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Eiptwostore.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Eiptwostore *EiptwostoreSession) Decimals() (uint8, error) {
	return _Eiptwostore.Contract.Decimals(&_Eiptwostore.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Eiptwostore *EiptwostoreCallerSession) Decimals() (uint8, error) {
	return _Eiptwostore.Contract.Decimals(&_Eiptwostore.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Eiptwostore *EiptwostoreCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Eiptwostore.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Eiptwostore *EiptwostoreSession) Name() (string, error) {
	return _Eiptwostore.Contract.Name(&_Eiptwostore.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Eiptwostore *EiptwostoreCallerSession) Name() (string, error) {
	return _Eiptwostore.Contract.Name(&_Eiptwostore.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Eiptwostore *EiptwostoreCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Eiptwostore.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Eiptwostore *EiptwostoreSession) Symbol() (string, error) {
	return _Eiptwostore.Contract.Symbol(&_Eiptwostore.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Eiptwostore *EiptwostoreCallerSession) Symbol() (string, error) {
	return _Eiptwostore.Contract.Symbol(&_Eiptwostore.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Eiptwostore *EiptwostoreCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Eiptwostore.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Eiptwostore *EiptwostoreSession) TotalSupply() (*big.Int, error) {
	return _Eiptwostore.Contract.TotalSupply(&_Eiptwostore.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Eiptwostore *EiptwostoreCallerSession) TotalSupply() (*big.Int, error) {
	return _Eiptwostore.Contract.TotalSupply(&_Eiptwostore.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Eiptwostore *EiptwostoreTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Eiptwostore.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Eiptwostore *EiptwostoreSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Eiptwostore.Contract.Approve(&_Eiptwostore.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Eiptwostore *EiptwostoreTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Eiptwostore.Contract.Approve(&_Eiptwostore.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Eiptwostore *EiptwostoreTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Eiptwostore.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Eiptwostore *EiptwostoreSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Eiptwostore.Contract.DecreaseAllowance(&_Eiptwostore.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Eiptwostore *EiptwostoreTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Eiptwostore.Contract.DecreaseAllowance(&_Eiptwostore.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Eiptwostore *EiptwostoreTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Eiptwostore.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Eiptwostore *EiptwostoreSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Eiptwostore.Contract.IncreaseAllowance(&_Eiptwostore.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Eiptwostore *EiptwostoreTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Eiptwostore.Contract.IncreaseAllowance(&_Eiptwostore.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Eiptwostore *EiptwostoreTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Eiptwostore.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Eiptwostore *EiptwostoreSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Eiptwostore.Contract.Transfer(&_Eiptwostore.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Eiptwostore *EiptwostoreTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Eiptwostore.Contract.Transfer(&_Eiptwostore.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Eiptwostore *EiptwostoreTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Eiptwostore.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Eiptwostore *EiptwostoreSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Eiptwostore.Contract.TransferFrom(&_Eiptwostore.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Eiptwostore *EiptwostoreTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Eiptwostore.Contract.TransferFrom(&_Eiptwostore.TransactOpts, from, to, amount)
}

// EiptwostoreApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Eiptwostore contract.
type EiptwostoreApprovalIterator struct {
	Event *EiptwostoreApproval // Event containing the contract specifics and raw log

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
func (it *EiptwostoreApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EiptwostoreApproval)
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
		it.Event = new(EiptwostoreApproval)
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
func (it *EiptwostoreApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EiptwostoreApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EiptwostoreApproval represents a Approval event raised by the Eiptwostore contract.
type EiptwostoreApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Eiptwostore *EiptwostoreFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EiptwostoreApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Eiptwostore.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EiptwostoreApprovalIterator{contract: _Eiptwostore.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Eiptwostore *EiptwostoreFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EiptwostoreApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Eiptwostore.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EiptwostoreApproval)
				if err := _Eiptwostore.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Eiptwostore *EiptwostoreFilterer) ParseApproval(log types.Log) (*EiptwostoreApproval, error) {
	event := new(EiptwostoreApproval)
	if err := _Eiptwostore.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EiptwostoreTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Eiptwostore contract.
type EiptwostoreTransferIterator struct {
	Event *EiptwostoreTransfer // Event containing the contract specifics and raw log

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
func (it *EiptwostoreTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EiptwostoreTransfer)
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
		it.Event = new(EiptwostoreTransfer)
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
func (it *EiptwostoreTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EiptwostoreTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EiptwostoreTransfer represents a Transfer event raised by the Eiptwostore contract.
type EiptwostoreTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Eiptwostore *EiptwostoreFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EiptwostoreTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Eiptwostore.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EiptwostoreTransferIterator{contract: _Eiptwostore.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Eiptwostore *EiptwostoreFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EiptwostoreTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Eiptwostore.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EiptwostoreTransfer)
				if err := _Eiptwostore.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Eiptwostore *EiptwostoreFilterer) ParseTransfer(log types.Log) (*EiptwostoreTransfer, error) {
	event := new(EiptwostoreTransfer)
	if err := _Eiptwostore.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
