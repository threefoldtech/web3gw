// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tft

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

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// TftMetaData contains all meta data concerning the Tft contract.
var TftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numberOfSignatures\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSignatures\",\"type\":\"uint256\"}],\"name\":\"InsufficientSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"AddedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"removedOwner\",\"type\":\"address\"}],\"name\":\"RemovedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"blockchain_address\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetSignaturesRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_txid\",\"type\":\"string\"}],\"name\":\"isMintID\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"is_owner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"mintTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owners_list\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_toRemove\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newSigners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"signaturesRequired\",\"type\":\"uint256\"}],\"name\":\"setSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"blockchain_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TftABI is the input ABI used to generate the binding from.
// Deprecated: Use TftMetaData.ABI instead.
var TftABI = TftMetaData.ABI

// Tft is an auto generated Go binding around an Ethereum contract.
type Tft struct {
	TftCaller     // Read-only binding to the contract
	TftTransactor // Write-only binding to the contract
	TftFilterer   // Log filterer for contract events
}

// TftCaller is an auto generated read-only Go binding around an Ethereum contract.
type TftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TftSession struct {
	Contract     *Tft              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TftCallerSession struct {
	Contract *TftCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TftTransactorSession struct {
	Contract     *TftTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TftRaw is an auto generated low-level Go binding around an Ethereum contract.
type TftRaw struct {
	Contract *Tft // Generic contract binding to access the raw methods on
}

// TftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TftCallerRaw struct {
	Contract *TftCaller // Generic read-only contract binding to access the raw methods on
}

// TftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TftTransactorRaw struct {
	Contract *TftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTft creates a new instance of Tft, bound to a specific deployed contract.
func NewTft(address common.Address, backend bind.ContractBackend) (*Tft, error) {
	contract, err := bindTft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tft{TftCaller: TftCaller{contract: contract}, TftTransactor: TftTransactor{contract: contract}, TftFilterer: TftFilterer{contract: contract}}, nil
}

// NewTftCaller creates a new read-only instance of Tft, bound to a specific deployed contract.
func NewTftCaller(address common.Address, caller bind.ContractCaller) (*TftCaller, error) {
	contract, err := bindTft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TftCaller{contract: contract}, nil
}

// NewTftTransactor creates a new write-only instance of Tft, bound to a specific deployed contract.
func NewTftTransactor(address common.Address, transactor bind.ContractTransactor) (*TftTransactor, error) {
	contract, err := bindTft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TftTransactor{contract: contract}, nil
}

// NewTftFilterer creates a new log filterer instance of Tft, bound to a specific deployed contract.
func NewTftFilterer(address common.Address, filterer bind.ContractFilterer) (*TftFilterer, error) {
	contract, err := bindTft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TftFilterer{contract: contract}, nil
}

// bindTft binds a generic wrapper to an already deployed contract.
func bindTft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tft *TftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tft.Contract.TftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tft *TftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tft.Contract.TftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tft *TftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tft.Contract.TftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tft *TftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tft *TftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tft *TftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tft.Contract.contract.Transact(opts, method, params...)
}

// GetSignaturesRequired is a free data retrieval call binding the contract method 0xbc3962b5.
//
// Solidity: function GetSignaturesRequired() view returns(uint256)
func (_Tft *TftCaller) GetSignaturesRequired(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tft.contract.Call(opts, &out, "GetSignaturesRequired")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSignaturesRequired is a free data retrieval call binding the contract method 0xbc3962b5.
//
// Solidity: function GetSignaturesRequired() view returns(uint256)
func (_Tft *TftSession) GetSignaturesRequired() (*big.Int, error) {
	return _Tft.Contract.GetSignaturesRequired(&_Tft.CallOpts)
}

// GetSignaturesRequired is a free data retrieval call binding the contract method 0xbc3962b5.
//
// Solidity: function GetSignaturesRequired() view returns(uint256)
func (_Tft *TftCallerSession) GetSignaturesRequired() (*big.Int, error) {
	return _Tft.Contract.GetSignaturesRequired(&_Tft.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_Tft *TftCaller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tft.contract.Call(opts, &out, "allowance", tokenOwner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_Tft *TftSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Tft.Contract.Allowance(&_Tft.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_Tft *TftCallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Tft.Contract.Allowance(&_Tft.CallOpts, tokenOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_Tft *TftCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tft.contract.Call(opts, &out, "balanceOf", tokenOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_Tft *TftSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Tft.Contract.BalanceOf(&_Tft.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_Tft *TftCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Tft.Contract.BalanceOf(&_Tft.CallOpts, tokenOwner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tft *TftCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tft.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tft *TftSession) Decimals() (uint8, error) {
	return _Tft.Contract.Decimals(&_Tft.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tft *TftCallerSession) Decimals() (uint8, error) {
	return _Tft.Contract.Decimals(&_Tft.CallOpts)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_Tft *TftCaller) GetSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Tft.contract.Call(opts, &out, "getSigners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_Tft *TftSession) GetSigners() ([]common.Address, error) {
	return _Tft.Contract.GetSigners(&_Tft.CallOpts)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_Tft *TftCallerSession) GetSigners() ([]common.Address, error) {
	return _Tft.Contract.GetSigners(&_Tft.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Tft *TftCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tft.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Tft *TftSession) Implementation() (common.Address, error) {
	return _Tft.Contract.Implementation(&_Tft.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Tft *TftCallerSession) Implementation() (common.Address, error) {
	return _Tft.Contract.Implementation(&_Tft.CallOpts)
}

// IsMintID is a free data retrieval call binding the contract method 0xdd6ad77e.
//
// Solidity: function isMintID(string _txid) view returns(bool)
func (_Tft *TftCaller) IsMintID(opts *bind.CallOpts, _txid string) (bool, error) {
	var out []interface{}
	err := _Tft.contract.Call(opts, &out, "isMintID", _txid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMintID is a free data retrieval call binding the contract method 0xdd6ad77e.
//
// Solidity: function isMintID(string _txid) view returns(bool)
func (_Tft *TftSession) IsMintID(_txid string) (bool, error) {
	return _Tft.Contract.IsMintID(&_Tft.CallOpts, _txid)
}

// IsMintID is a free data retrieval call binding the contract method 0xdd6ad77e.
//
// Solidity: function isMintID(string _txid) view returns(bool)
func (_Tft *TftCallerSession) IsMintID(_txid string) (bool, error) {
	return _Tft.Contract.IsMintID(&_Tft.CallOpts, _txid)
}

// IsOwner is a free data retrieval call binding the contract method 0x0776076f.
//
// Solidity: function is_owner(address owner) view returns(bool)
func (_Tft *TftCaller) IsOwner(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _Tft.contract.Call(opts, &out, "is_owner", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x0776076f.
//
// Solidity: function is_owner(address owner) view returns(bool)
func (_Tft *TftSession) IsOwner(owner common.Address) (bool, error) {
	return _Tft.Contract.IsOwner(&_Tft.CallOpts, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x0776076f.
//
// Solidity: function is_owner(address owner) view returns(bool)
func (_Tft *TftCallerSession) IsOwner(owner common.Address) (bool, error) {
	return _Tft.Contract.IsOwner(&_Tft.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tft *TftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tft *TftSession) Name() (string, error) {
	return _Tft.Contract.Name(&_Tft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tft *TftCallerSession) Name() (string, error) {
	return _Tft.Contract.Name(&_Tft.CallOpts)
}

// OwnersList is a free data retrieval call binding the contract method 0xb41a88c0.
//
// Solidity: function owners_list() view returns(address[])
func (_Tft *TftCaller) OwnersList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Tft.contract.Call(opts, &out, "owners_list")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// OwnersList is a free data retrieval call binding the contract method 0xb41a88c0.
//
// Solidity: function owners_list() view returns(address[])
func (_Tft *TftSession) OwnersList() ([]common.Address, error) {
	return _Tft.Contract.OwnersList(&_Tft.CallOpts)
}

// OwnersList is a free data retrieval call binding the contract method 0xb41a88c0.
//
// Solidity: function owners_list() view returns(address[])
func (_Tft *TftCallerSession) OwnersList() ([]common.Address, error) {
	return _Tft.Contract.OwnersList(&_Tft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tft *TftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tft *TftSession) Symbol() (string, error) {
	return _Tft.Contract.Symbol(&_Tft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tft *TftCallerSession) Symbol() (string, error) {
	return _Tft.Contract.Symbol(&_Tft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tft *TftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tft.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tft *TftSession) TotalSupply() (*big.Int, error) {
	return _Tft.Contract.TotalSupply(&_Tft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tft *TftCallerSession) TotalSupply() (*big.Int, error) {
	return _Tft.Contract.TotalSupply(&_Tft.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Tft *TftCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tft.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Tft *TftSession) Version() (string, error) {
	return _Tft.Contract.Version(&_Tft.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Tft *TftCallerSession) Version() (string, error) {
	return _Tft.Contract.Version(&_Tft.CallOpts)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address _newOwner) returns()
func (_Tft *TftTransactor) AddOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Tft.contract.Transact(opts, "addOwner", _newOwner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address _newOwner) returns()
func (_Tft *TftSession) AddOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Tft.Contract.AddOwner(&_Tft.TransactOpts, _newOwner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address _newOwner) returns()
func (_Tft *TftTransactorSession) AddOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Tft.Contract.AddOwner(&_Tft.TransactOpts, _newOwner)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_Tft *TftTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Tft.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_Tft *TftSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Tft.Contract.Approve(&_Tft.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_Tft *TftTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Tft.Contract.Approve(&_Tft.TransactOpts, spender, tokens)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf5d124de.
//
// Solidity: function mintTokens(address receiver, uint256 tokens, string txid, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Tft *TftTransactor) MintTokens(opts *bind.TransactOpts, receiver common.Address, tokens *big.Int, txid string, _signatures []Signature) (*types.Transaction, error) {
	return _Tft.contract.Transact(opts, "mintTokens", receiver, tokens, txid, _signatures)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf5d124de.
//
// Solidity: function mintTokens(address receiver, uint256 tokens, string txid, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Tft *TftSession) MintTokens(receiver common.Address, tokens *big.Int, txid string, _signatures []Signature) (*types.Transaction, error) {
	return _Tft.Contract.MintTokens(&_Tft.TransactOpts, receiver, tokens, txid, _signatures)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf5d124de.
//
// Solidity: function mintTokens(address receiver, uint256 tokens, string txid, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Tft *TftTransactorSession) MintTokens(receiver common.Address, tokens *big.Int, txid string, _signatures []Signature) (*types.Transaction, error) {
	return _Tft.Contract.MintTokens(&_Tft.TransactOpts, receiver, tokens, txid, _signatures)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address _toRemove) returns()
func (_Tft *TftTransactor) RemoveOwner(opts *bind.TransactOpts, _toRemove common.Address) (*types.Transaction, error) {
	return _Tft.contract.Transact(opts, "removeOwner", _toRemove)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address _toRemove) returns()
func (_Tft *TftSession) RemoveOwner(_toRemove common.Address) (*types.Transaction, error) {
	return _Tft.Contract.RemoveOwner(&_Tft.TransactOpts, _toRemove)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address _toRemove) returns()
func (_Tft *TftTransactorSession) RemoveOwner(_toRemove common.Address) (*types.Transaction, error) {
	return _Tft.Contract.RemoveOwner(&_Tft.TransactOpts, _toRemove)
}

// SetSigners is a paid mutator transaction binding the contract method 0xb4368904.
//
// Solidity: function setSigners(address[] newSigners, uint256 signaturesRequired) returns()
func (_Tft *TftTransactor) SetSigners(opts *bind.TransactOpts, newSigners []common.Address, signaturesRequired *big.Int) (*types.Transaction, error) {
	return _Tft.contract.Transact(opts, "setSigners", newSigners, signaturesRequired)
}

// SetSigners is a paid mutator transaction binding the contract method 0xb4368904.
//
// Solidity: function setSigners(address[] newSigners, uint256 signaturesRequired) returns()
func (_Tft *TftSession) SetSigners(newSigners []common.Address, signaturesRequired *big.Int) (*types.Transaction, error) {
	return _Tft.Contract.SetSigners(&_Tft.TransactOpts, newSigners, signaturesRequired)
}

// SetSigners is a paid mutator transaction binding the contract method 0xb4368904.
//
// Solidity: function setSigners(address[] newSigners, uint256 signaturesRequired) returns()
func (_Tft *TftTransactorSession) SetSigners(newSigners []common.Address, signaturesRequired *big.Int) (*types.Transaction, error) {
	return _Tft.Contract.SetSigners(&_Tft.TransactOpts, newSigners, signaturesRequired)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_Tft *TftTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Tft.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_Tft *TftSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Tft.Contract.Transfer(&_Tft.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_Tft *TftTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Tft.Contract.Transfer(&_Tft.TransactOpts, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Tft *TftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Tft.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Tft *TftSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Tft.Contract.TransferFrom(&_Tft.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Tft *TftTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Tft.Contract.TransferFrom(&_Tft.TransactOpts, from, to, tokens)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x5a8b1a9f.
//
// Solidity: function upgradeTo(string _version, address _implementation) returns()
func (_Tft *TftTransactor) UpgradeTo(opts *bind.TransactOpts, _version string, _implementation common.Address) (*types.Transaction, error) {
	return _Tft.contract.Transact(opts, "upgradeTo", _version, _implementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x5a8b1a9f.
//
// Solidity: function upgradeTo(string _version, address _implementation) returns()
func (_Tft *TftSession) UpgradeTo(_version string, _implementation common.Address) (*types.Transaction, error) {
	return _Tft.Contract.UpgradeTo(&_Tft.TransactOpts, _version, _implementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x5a8b1a9f.
//
// Solidity: function upgradeTo(string _version, address _implementation) returns()
func (_Tft *TftTransactorSession) UpgradeTo(_version string, _implementation common.Address) (*types.Transaction, error) {
	return _Tft.Contract.UpgradeTo(&_Tft.TransactOpts, _version, _implementation)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9a493e75.
//
// Solidity: function withdraw(uint256 tokens, string blockchain_address, string network) returns(bool success)
func (_Tft *TftTransactor) Withdraw(opts *bind.TransactOpts, tokens *big.Int, blockchain_address string, network string) (*types.Transaction, error) {
	return _Tft.contract.Transact(opts, "withdraw", tokens, blockchain_address, network)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9a493e75.
//
// Solidity: function withdraw(uint256 tokens, string blockchain_address, string network) returns(bool success)
func (_Tft *TftSession) Withdraw(tokens *big.Int, blockchain_address string, network string) (*types.Transaction, error) {
	return _Tft.Contract.Withdraw(&_Tft.TransactOpts, tokens, blockchain_address, network)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9a493e75.
//
// Solidity: function withdraw(uint256 tokens, string blockchain_address, string network) returns(bool success)
func (_Tft *TftTransactorSession) Withdraw(tokens *big.Int, blockchain_address string, network string) (*types.Transaction, error) {
	return _Tft.Contract.Withdraw(&_Tft.TransactOpts, tokens, blockchain_address, network)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Tft *TftTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tft.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Tft *TftSession) Receive() (*types.Transaction, error) {
	return _Tft.Contract.Receive(&_Tft.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Tft *TftTransactorSession) Receive() (*types.Transaction, error) {
	return _Tft.Contract.Receive(&_Tft.TransactOpts)
}

// TftAddedOwnerIterator is returned from FilterAddedOwner and is used to iterate over the raw logs and unpacked data for AddedOwner events raised by the Tft contract.
type TftAddedOwnerIterator struct {
	Event *TftAddedOwner // Event containing the contract specifics and raw log

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
func (it *TftAddedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TftAddedOwner)
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
		it.Event = new(TftAddedOwner)
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
func (it *TftAddedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TftAddedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TftAddedOwner represents a AddedOwner event raised by the Tft contract.
type TftAddedOwner struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddedOwner is a free log retrieval operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address newOwner)
func (_Tft *TftFilterer) FilterAddedOwner(opts *bind.FilterOpts) (*TftAddedOwnerIterator, error) {

	logs, sub, err := _Tft.contract.FilterLogs(opts, "AddedOwner")
	if err != nil {
		return nil, err
	}
	return &TftAddedOwnerIterator{contract: _Tft.contract, event: "AddedOwner", logs: logs, sub: sub}, nil
}

// WatchAddedOwner is a free log subscription operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address newOwner)
func (_Tft *TftFilterer) WatchAddedOwner(opts *bind.WatchOpts, sink chan<- *TftAddedOwner) (event.Subscription, error) {

	logs, sub, err := _Tft.contract.WatchLogs(opts, "AddedOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TftAddedOwner)
				if err := _Tft.contract.UnpackLog(event, "AddedOwner", log); err != nil {
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

// ParseAddedOwner is a log parse operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address newOwner)
func (_Tft *TftFilterer) ParseAddedOwner(log types.Log) (*TftAddedOwner, error) {
	event := new(TftAddedOwner)
	if err := _Tft.contract.UnpackLog(event, "AddedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Tft contract.
type TftApprovalIterator struct {
	Event *TftApproval // Event containing the contract specifics and raw log

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
func (it *TftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TftApproval)
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
		it.Event = new(TftApproval)
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
func (it *TftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TftApproval represents a Approval event raised by the Tft contract.
type TftApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Tft *TftFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*TftApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tft.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TftApprovalIterator{contract: _Tft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Tft *TftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TftApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tft.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TftApproval)
				if err := _Tft.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Tft *TftFilterer) ParseApproval(log types.Log) (*TftApproval, error) {
	event := new(TftApproval)
	if err := _Tft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TftMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Tft contract.
type TftMintIterator struct {
	Event *TftMint // Event containing the contract specifics and raw log

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
func (it *TftMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TftMint)
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
		it.Event = new(TftMint)
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
func (it *TftMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TftMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TftMint represents a Mint event raised by the Tft contract.
type TftMint struct {
	Receiver common.Address
	Tokens   *big.Int
	Txid     common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x85a66b9141978db9980f7e0ce3b468cebf4f7999f32b23091c5c03e798b1ba7a.
//
// Solidity: event Mint(address indexed receiver, uint256 tokens, string indexed txid)
func (_Tft *TftFilterer) FilterMint(opts *bind.FilterOpts, receiver []common.Address, txid []string) (*TftMintIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	var txidRule []interface{}
	for _, txidItem := range txid {
		txidRule = append(txidRule, txidItem)
	}

	logs, sub, err := _Tft.contract.FilterLogs(opts, "Mint", receiverRule, txidRule)
	if err != nil {
		return nil, err
	}
	return &TftMintIterator{contract: _Tft.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x85a66b9141978db9980f7e0ce3b468cebf4f7999f32b23091c5c03e798b1ba7a.
//
// Solidity: event Mint(address indexed receiver, uint256 tokens, string indexed txid)
func (_Tft *TftFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *TftMint, receiver []common.Address, txid []string) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	var txidRule []interface{}
	for _, txidItem := range txid {
		txidRule = append(txidRule, txidItem)
	}

	logs, sub, err := _Tft.contract.WatchLogs(opts, "Mint", receiverRule, txidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TftMint)
				if err := _Tft.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x85a66b9141978db9980f7e0ce3b468cebf4f7999f32b23091c5c03e798b1ba7a.
//
// Solidity: event Mint(address indexed receiver, uint256 tokens, string indexed txid)
func (_Tft *TftFilterer) ParseMint(log types.Log) (*TftMint, error) {
	event := new(TftMint)
	if err := _Tft.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TftRemovedOwnerIterator is returned from FilterRemovedOwner and is used to iterate over the raw logs and unpacked data for RemovedOwner events raised by the Tft contract.
type TftRemovedOwnerIterator struct {
	Event *TftRemovedOwner // Event containing the contract specifics and raw log

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
func (it *TftRemovedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TftRemovedOwner)
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
		it.Event = new(TftRemovedOwner)
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
func (it *TftRemovedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TftRemovedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TftRemovedOwner represents a RemovedOwner event raised by the Tft contract.
type TftRemovedOwner struct {
	RemovedOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemovedOwner is a free log retrieval operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address removedOwner)
func (_Tft *TftFilterer) FilterRemovedOwner(opts *bind.FilterOpts) (*TftRemovedOwnerIterator, error) {

	logs, sub, err := _Tft.contract.FilterLogs(opts, "RemovedOwner")
	if err != nil {
		return nil, err
	}
	return &TftRemovedOwnerIterator{contract: _Tft.contract, event: "RemovedOwner", logs: logs, sub: sub}, nil
}

// WatchRemovedOwner is a free log subscription operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address removedOwner)
func (_Tft *TftFilterer) WatchRemovedOwner(opts *bind.WatchOpts, sink chan<- *TftRemovedOwner) (event.Subscription, error) {

	logs, sub, err := _Tft.contract.WatchLogs(opts, "RemovedOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TftRemovedOwner)
				if err := _Tft.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
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

// ParseRemovedOwner is a log parse operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address removedOwner)
func (_Tft *TftFilterer) ParseRemovedOwner(log types.Log) (*TftRemovedOwner, error) {
	event := new(TftRemovedOwner)
	if err := _Tft.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Tft contract.
type TftTransferIterator struct {
	Event *TftTransfer // Event containing the contract specifics and raw log

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
func (it *TftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TftTransfer)
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
		it.Event = new(TftTransfer)
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
func (it *TftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TftTransfer represents a Transfer event raised by the Tft contract.
type TftTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Tft *TftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TftTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tft.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TftTransferIterator{contract: _Tft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Tft *TftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TftTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tft.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TftTransfer)
				if err := _Tft.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Tft *TftFilterer) ParseTransfer(log types.Log) (*TftTransfer, error) {
	event := new(TftTransfer)
	if err := _Tft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TftUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Tft contract.
type TftUpgradedIterator struct {
	Event *TftUpgraded // Event containing the contract specifics and raw log

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
func (it *TftUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TftUpgraded)
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
		it.Event = new(TftUpgraded)
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
func (it *TftUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TftUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TftUpgraded represents a Upgraded event raised by the Tft contract.
type TftUpgraded struct {
	Version        common.Hash
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0x8e05e0e35ff592971ca8b477d4285a33a61ded208d644042667b78693a472f5e.
//
// Solidity: event Upgraded(string indexed version, address indexed implementation)
func (_Tft *TftFilterer) FilterUpgraded(opts *bind.FilterOpts, version []string, implementation []common.Address) (*TftUpgradedIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Tft.contract.FilterLogs(opts, "Upgraded", versionRule, implementationRule)
	if err != nil {
		return nil, err
	}
	return &TftUpgradedIterator{contract: _Tft.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0x8e05e0e35ff592971ca8b477d4285a33a61ded208d644042667b78693a472f5e.
//
// Solidity: event Upgraded(string indexed version, address indexed implementation)
func (_Tft *TftFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TftUpgraded, version []string, implementation []common.Address) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Tft.contract.WatchLogs(opts, "Upgraded", versionRule, implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TftUpgraded)
				if err := _Tft.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0x8e05e0e35ff592971ca8b477d4285a33a61ded208d644042667b78693a472f5e.
//
// Solidity: event Upgraded(string indexed version, address indexed implementation)
func (_Tft *TftFilterer) ParseUpgraded(log types.Log) (*TftUpgraded, error) {
	event := new(TftUpgraded)
	if err := _Tft.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TftWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Tft contract.
type TftWithdrawIterator struct {
	Event *TftWithdraw // Event containing the contract specifics and raw log

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
func (it *TftWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TftWithdraw)
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
		it.Event = new(TftWithdraw)
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
func (it *TftWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TftWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TftWithdraw represents a Withdraw event raised by the Tft contract.
type TftWithdraw struct {
	Receiver          common.Address
	Tokens            *big.Int
	BlockchainAddress string
	Network           string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xbf4bee5506452a156854c54e249d6b04b0cd83287ba208202be81a4f87a55739.
//
// Solidity: event Withdraw(address indexed receiver, uint256 tokens, string blockchain_address, string network)
func (_Tft *TftFilterer) FilterWithdraw(opts *bind.FilterOpts, receiver []common.Address) (*TftWithdrawIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Tft.contract.FilterLogs(opts, "Withdraw", receiverRule)
	if err != nil {
		return nil, err
	}
	return &TftWithdrawIterator{contract: _Tft.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xbf4bee5506452a156854c54e249d6b04b0cd83287ba208202be81a4f87a55739.
//
// Solidity: event Withdraw(address indexed receiver, uint256 tokens, string blockchain_address, string network)
func (_Tft *TftFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *TftWithdraw, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Tft.contract.WatchLogs(opts, "Withdraw", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TftWithdraw)
				if err := _Tft.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xbf4bee5506452a156854c54e249d6b04b0cd83287ba208202be81a4f87a55739.
//
// Solidity: event Withdraw(address indexed receiver, uint256 tokens, string blockchain_address, string network)
func (_Tft *TftFilterer) ParseWithdraw(log types.Log) (*TftWithdraw, error) {
	event := new(TftWithdraw)
	if err := _Tft.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
