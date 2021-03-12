// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SeverABI is the input ABI used to generate the binding from.
const SeverABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_server\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guarantor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Intervene\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guarantor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guarantor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bail\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractsNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_A\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_B\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_judge\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_Content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_feePercentLimit\",\"type\":\"uint256\"}],\"name\":\"createContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"guarantor\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_guarantor\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"intervene\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAdr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"}],\"name\":\"interveneContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"relations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"relationsNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"server\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Sever is an auto generated Go binding around an Ethereum contract.
type Sever struct {
	SeverCaller     // Read-only binding to the contract
	SeverTransactor // Write-only binding to the contract
	SeverFilterer   // Log filterer for contract events
}

// SeverCaller is an auto generated read-only Go binding around an Ethereum contract.
type SeverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SeverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SeverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SeverSession struct {
	Contract     *Sever            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SeverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SeverCallerSession struct {
	Contract *SeverCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SeverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SeverTransactorSession struct {
	Contract     *SeverTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SeverRaw is an auto generated low-level Go binding around an Ethereum contract.
type SeverRaw struct {
	Contract *Sever // Generic contract binding to access the raw methods on
}

// SeverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SeverCallerRaw struct {
	Contract *SeverCaller // Generic read-only contract binding to access the raw methods on
}

// SeverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SeverTransactorRaw struct {
	Contract *SeverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSever creates a new instance of Sever, bound to a specific deployed contract.
func NewSever(address common.Address, backend bind.ContractBackend) (*Sever, error) {
	contract, err := bindSever(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sever{SeverCaller: SeverCaller{contract: contract}, SeverTransactor: SeverTransactor{contract: contract}, SeverFilterer: SeverFilterer{contract: contract}}, nil
}

// NewSeverCaller creates a new read-only instance of Sever, bound to a specific deployed contract.
func NewSeverCaller(address common.Address, caller bind.ContractCaller) (*SeverCaller, error) {
	contract, err := bindSever(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SeverCaller{contract: contract}, nil
}

// NewSeverTransactor creates a new write-only instance of Sever, bound to a specific deployed contract.
func NewSeverTransactor(address common.Address, transactor bind.ContractTransactor) (*SeverTransactor, error) {
	contract, err := bindSever(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SeverTransactor{contract: contract}, nil
}

// NewSeverFilterer creates a new log filterer instance of Sever, bound to a specific deployed contract.
func NewSeverFilterer(address common.Address, filterer bind.ContractFilterer) (*SeverFilterer, error) {
	contract, err := bindSever(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SeverFilterer{contract: contract}, nil
}

// bindSever binds a generic wrapper to an already deployed contract.
func bindSever(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SeverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sever *SeverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sever.Contract.SeverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sever *SeverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sever.Contract.SeverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sever *SeverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sever.Contract.SeverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sever *SeverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sever.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sever *SeverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sever.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sever *SeverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sever.Contract.contract.Transact(opts, method, params...)
}

// Bail is a free data retrieval call binding the contract method 0x4dc639ea.
//
// Solidity: function bail(address ) view returns(uint256)
func (_Sever *SeverCaller) Bail(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sever.contract.Call(opts, &out, "bail", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bail is a free data retrieval call binding the contract method 0x4dc639ea.
//
// Solidity: function bail(address ) view returns(uint256)
func (_Sever *SeverSession) Bail(arg0 common.Address) (*big.Int, error) {
	return _Sever.Contract.Bail(&_Sever.CallOpts, arg0)
}

// Bail is a free data retrieval call binding the contract method 0x4dc639ea.
//
// Solidity: function bail(address ) view returns(uint256)
func (_Sever *SeverCallerSession) Bail(arg0 common.Address) (*big.Int, error) {
	return _Sever.Contract.Bail(&_Sever.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0x474da79a.
//
// Solidity: function contracts(uint256 ) view returns(address)
func (_Sever *SeverCaller) Contracts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Sever.contract.Call(opts, &out, "contracts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Contracts is a free data retrieval call binding the contract method 0x474da79a.
//
// Solidity: function contracts(uint256 ) view returns(address)
func (_Sever *SeverSession) Contracts(arg0 *big.Int) (common.Address, error) {
	return _Sever.Contract.Contracts(&_Sever.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0x474da79a.
//
// Solidity: function contracts(uint256 ) view returns(address)
func (_Sever *SeverCallerSession) Contracts(arg0 *big.Int) (common.Address, error) {
	return _Sever.Contract.Contracts(&_Sever.CallOpts, arg0)
}

// ContractsNum is a free data retrieval call binding the contract method 0x778d1b4a.
//
// Solidity: function contractsNum() view returns(uint256)
func (_Sever *SeverCaller) ContractsNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sever.contract.Call(opts, &out, "contractsNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractsNum is a free data retrieval call binding the contract method 0x778d1b4a.
//
// Solidity: function contractsNum() view returns(uint256)
func (_Sever *SeverSession) ContractsNum() (*big.Int, error) {
	return _Sever.Contract.ContractsNum(&_Sever.CallOpts)
}

// ContractsNum is a free data retrieval call binding the contract method 0x778d1b4a.
//
// Solidity: function contractsNum() view returns(uint256)
func (_Sever *SeverCallerSession) ContractsNum() (*big.Int, error) {
	return _Sever.Contract.ContractsNum(&_Sever.CallOpts)
}

// Guarantor is a free data retrieval call binding the contract method 0xcfe9bb98.
//
// Solidity: function guarantor(uint256 ) view returns(address)
func (_Sever *SeverCaller) Guarantor(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Sever.contract.Call(opts, &out, "guarantor", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guarantor is a free data retrieval call binding the contract method 0xcfe9bb98.
//
// Solidity: function guarantor(uint256 ) view returns(address)
func (_Sever *SeverSession) Guarantor(arg0 *big.Int) (common.Address, error) {
	return _Sever.Contract.Guarantor(&_Sever.CallOpts, arg0)
}

// Guarantor is a free data retrieval call binding the contract method 0xcfe9bb98.
//
// Solidity: function guarantor(uint256 ) view returns(address)
func (_Sever *SeverCallerSession) Guarantor(arg0 *big.Int) (common.Address, error) {
	return _Sever.Contract.Guarantor(&_Sever.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x01984892.
//
// Solidity: function name(address ) view returns(string)
func (_Sever *SeverCaller) Name(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Sever.contract.Call(opts, &out, "name", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x01984892.
//
// Solidity: function name(address ) view returns(string)
func (_Sever *SeverSession) Name(arg0 common.Address) (string, error) {
	return _Sever.Contract.Name(&_Sever.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x01984892.
//
// Solidity: function name(address ) view returns(string)
func (_Sever *SeverCallerSession) Name(arg0 common.Address) (string, error) {
	return _Sever.Contract.Name(&_Sever.CallOpts, arg0)
}

// Relations is a free data retrieval call binding the contract method 0xbcafddd3.
//
// Solidity: function relations(address , uint256 ) view returns(uint256)
func (_Sever *SeverCaller) Relations(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Sever.contract.Call(opts, &out, "relations", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Relations is a free data retrieval call binding the contract method 0xbcafddd3.
//
// Solidity: function relations(address , uint256 ) view returns(uint256)
func (_Sever *SeverSession) Relations(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Sever.Contract.Relations(&_Sever.CallOpts, arg0, arg1)
}

// Relations is a free data retrieval call binding the contract method 0xbcafddd3.
//
// Solidity: function relations(address , uint256 ) view returns(uint256)
func (_Sever *SeverCallerSession) Relations(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Sever.Contract.Relations(&_Sever.CallOpts, arg0, arg1)
}

// RelationsNum is a free data retrieval call binding the contract method 0x4b828a58.
//
// Solidity: function relationsNum(address ) view returns(uint256)
func (_Sever *SeverCaller) RelationsNum(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sever.contract.Call(opts, &out, "relationsNum", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelationsNum is a free data retrieval call binding the contract method 0x4b828a58.
//
// Solidity: function relationsNum(address ) view returns(uint256)
func (_Sever *SeverSession) RelationsNum(arg0 common.Address) (*big.Int, error) {
	return _Sever.Contract.RelationsNum(&_Sever.CallOpts, arg0)
}

// RelationsNum is a free data retrieval call binding the contract method 0x4b828a58.
//
// Solidity: function relationsNum(address ) view returns(uint256)
func (_Sever *SeverCallerSession) RelationsNum(arg0 common.Address) (*big.Int, error) {
	return _Sever.Contract.RelationsNum(&_Sever.CallOpts, arg0)
}

// Server is a free data retrieval call binding the contract method 0xfd922a42.
//
// Solidity: function server() view returns(address)
func (_Sever *SeverCaller) Server(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sever.contract.Call(opts, &out, "server")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Server is a free data retrieval call binding the contract method 0xfd922a42.
//
// Solidity: function server() view returns(address)
func (_Sever *SeverSession) Server() (common.Address, error) {
	return _Sever.Contract.Server(&_Sever.CallOpts)
}

// Server is a free data retrieval call binding the contract method 0xfd922a42.
//
// Solidity: function server() view returns(address)
func (_Sever *SeverCallerSession) Server() (common.Address, error) {
	return _Sever.Contract.Server(&_Sever.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Sever *SeverCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sever.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Sever *SeverSession) Token() (common.Address, error) {
	return _Sever.Contract.Token(&_Sever.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Sever *SeverCallerSession) Token() (common.Address, error) {
	return _Sever.Contract.Token(&_Sever.CallOpts)
}

// CreateContract is a paid mutator transaction binding the contract method 0xd18c4a39.
//
// Solidity: function createContract(address _A, address _B, address _judge, string _Content, uint256 _feePercentLimit) returns(uint256)
func (_Sever *SeverTransactor) CreateContract(opts *bind.TransactOpts, _A common.Address, _B common.Address, _judge common.Address, _Content string, _feePercentLimit *big.Int) (*types.Transaction, error) {
	return _Sever.contract.Transact(opts, "createContract", _A, _B, _judge, _Content, _feePercentLimit)
}

// CreateContract is a paid mutator transaction binding the contract method 0xd18c4a39.
//
// Solidity: function createContract(address _A, address _B, address _judge, string _Content, uint256 _feePercentLimit) returns(uint256)
func (_Sever *SeverSession) CreateContract(_A common.Address, _B common.Address, _judge common.Address, _Content string, _feePercentLimit *big.Int) (*types.Transaction, error) {
	return _Sever.Contract.CreateContract(&_Sever.TransactOpts, _A, _B, _judge, _Content, _feePercentLimit)
}

// CreateContract is a paid mutator transaction binding the contract method 0xd18c4a39.
//
// Solidity: function createContract(address _A, address _B, address _judge, string _Content, uint256 _feePercentLimit) returns(uint256)
func (_Sever *SeverTransactorSession) CreateContract(_A common.Address, _B common.Address, _judge common.Address, _Content string, _feePercentLimit *big.Int) (*types.Transaction, error) {
	return _Sever.Contract.CreateContract(&_Sever.TransactOpts, _A, _B, _judge, _Content, _feePercentLimit)
}

// Intervene is a paid mutator transaction binding the contract method 0x1d69a284.
//
// Solidity: function intervene(address _guarantor, address _to, uint256 amount) returns()
func (_Sever *SeverTransactor) Intervene(opts *bind.TransactOpts, _guarantor common.Address, _to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sever.contract.Transact(opts, "intervene", _guarantor, _to, amount)
}

// Intervene is a paid mutator transaction binding the contract method 0x1d69a284.
//
// Solidity: function intervene(address _guarantor, address _to, uint256 amount) returns()
func (_Sever *SeverSession) Intervene(_guarantor common.Address, _to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sever.Contract.Intervene(&_Sever.TransactOpts, _guarantor, _to, amount)
}

// Intervene is a paid mutator transaction binding the contract method 0x1d69a284.
//
// Solidity: function intervene(address _guarantor, address _to, uint256 amount) returns()
func (_Sever *SeverTransactorSession) Intervene(_guarantor common.Address, _to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sever.Contract.Intervene(&_Sever.TransactOpts, _guarantor, _to, amount)
}

// InterveneContract is a paid mutator transaction binding the contract method 0xd0508d45.
//
// Solidity: function interveneContract(address contractAdr, uint256 amountToA, uint256 amountToB, uint256 feePercent) returns()
func (_Sever *SeverTransactor) InterveneContract(opts *bind.TransactOpts, contractAdr common.Address, amountToA *big.Int, amountToB *big.Int, feePercent *big.Int) (*types.Transaction, error) {
	return _Sever.contract.Transact(opts, "interveneContract", contractAdr, amountToA, amountToB, feePercent)
}

// InterveneContract is a paid mutator transaction binding the contract method 0xd0508d45.
//
// Solidity: function interveneContract(address contractAdr, uint256 amountToA, uint256 amountToB, uint256 feePercent) returns()
func (_Sever *SeverSession) InterveneContract(contractAdr common.Address, amountToA *big.Int, amountToB *big.Int, feePercent *big.Int) (*types.Transaction, error) {
	return _Sever.Contract.InterveneContract(&_Sever.TransactOpts, contractAdr, amountToA, amountToB, feePercent)
}

// InterveneContract is a paid mutator transaction binding the contract method 0xd0508d45.
//
// Solidity: function interveneContract(address contractAdr, uint256 amountToA, uint256 amountToB, uint256 feePercent) returns()
func (_Sever *SeverTransactorSession) InterveneContract(contractAdr common.Address, amountToA *big.Int, amountToB *big.Int, feePercent *big.Int) (*types.Transaction, error) {
	return _Sever.Contract.InterveneContract(&_Sever.TransactOpts, contractAdr, amountToA, amountToB, feePercent)
}

// Register is a paid mutator transaction binding the contract method 0xa00fd3c8.
//
// Solidity: function register(uint256 amount, string _name) returns()
func (_Sever *SeverTransactor) Register(opts *bind.TransactOpts, amount *big.Int, _name string) (*types.Transaction, error) {
	return _Sever.contract.Transact(opts, "register", amount, _name)
}

// Register is a paid mutator transaction binding the contract method 0xa00fd3c8.
//
// Solidity: function register(uint256 amount, string _name) returns()
func (_Sever *SeverSession) Register(amount *big.Int, _name string) (*types.Transaction, error) {
	return _Sever.Contract.Register(&_Sever.TransactOpts, amount, _name)
}

// Register is a paid mutator transaction binding the contract method 0xa00fd3c8.
//
// Solidity: function register(uint256 amount, string _name) returns()
func (_Sever *SeverTransactorSession) Register(amount *big.Int, _name string) (*types.Transaction, error) {
	return _Sever.Contract.Register(&_Sever.TransactOpts, amount, _name)
}

// SeverInterveneIterator is returned from FilterIntervene and is used to iterate over the raw logs and unpacked data for Intervene events raised by the Sever contract.
type SeverInterveneIterator struct {
	Event *SeverIntervene // Event containing the contract specifics and raw log

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
func (it *SeverInterveneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeverIntervene)
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
		it.Event = new(SeverIntervene)
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
func (it *SeverInterveneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeverInterveneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeverIntervene represents a Intervene event raised by the Sever contract.
type SeverIntervene struct {
	Guarantor common.Address
	To        common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIntervene is a free log retrieval operation binding the contract event 0x741071bb26c76c5bad155a306936205baeae4e80b8dc1277a5ade177c12291a4.
//
// Solidity: event Intervene(address guarantor, address _to, uint256 amount)
func (_Sever *SeverFilterer) FilterIntervene(opts *bind.FilterOpts) (*SeverInterveneIterator, error) {

	logs, sub, err := _Sever.contract.FilterLogs(opts, "Intervene")
	if err != nil {
		return nil, err
	}
	return &SeverInterveneIterator{contract: _Sever.contract, event: "Intervene", logs: logs, sub: sub}, nil
}

// WatchIntervene is a free log subscription operation binding the contract event 0x741071bb26c76c5bad155a306936205baeae4e80b8dc1277a5ade177c12291a4.
//
// Solidity: event Intervene(address guarantor, address _to, uint256 amount)
func (_Sever *SeverFilterer) WatchIntervene(opts *bind.WatchOpts, sink chan<- *SeverIntervene) (event.Subscription, error) {

	logs, sub, err := _Sever.contract.WatchLogs(opts, "Intervene")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeverIntervene)
				if err := _Sever.contract.UnpackLog(event, "Intervene", log); err != nil {
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

// ParseIntervene is a log parse operation binding the contract event 0x741071bb26c76c5bad155a306936205baeae4e80b8dc1277a5ade177c12291a4.
//
// Solidity: event Intervene(address guarantor, address _to, uint256 amount)
func (_Sever *SeverFilterer) ParseIntervene(log types.Log) (*SeverIntervene, error) {
	event := new(SeverIntervene)
	if err := _Sever.contract.UnpackLog(event, "Intervene", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeverRegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the Sever contract.
type SeverRegisterIterator struct {
	Event *SeverRegister // Event containing the contract specifics and raw log

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
func (it *SeverRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeverRegister)
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
		it.Event = new(SeverRegister)
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
func (it *SeverRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeverRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeverRegister represents a Register event raised by the Sever contract.
type SeverRegister struct {
	Guarantor common.Address
	Amount    *big.Int
	Name      string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0xfd7cce4ad4056ce11f491575da172369a8731f4cd81ef1eedb704c7e304f19cc.
//
// Solidity: event Register(address guarantor, uint256 amount, string name)
func (_Sever *SeverFilterer) FilterRegister(opts *bind.FilterOpts) (*SeverRegisterIterator, error) {

	logs, sub, err := _Sever.contract.FilterLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return &SeverRegisterIterator{contract: _Sever.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0xfd7cce4ad4056ce11f491575da172369a8731f4cd81ef1eedb704c7e304f19cc.
//
// Solidity: event Register(address guarantor, uint256 amount, string name)
func (_Sever *SeverFilterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *SeverRegister) (event.Subscription, error) {

	logs, sub, err := _Sever.contract.WatchLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeverRegister)
				if err := _Sever.contract.UnpackLog(event, "Register", log); err != nil {
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

// ParseRegister is a log parse operation binding the contract event 0xfd7cce4ad4056ce11f491575da172369a8731f4cd81ef1eedb704c7e304f19cc.
//
// Solidity: event Register(address guarantor, uint256 amount, string name)
func (_Sever *SeverFilterer) ParseRegister(log types.Log) (*SeverRegister, error) {
	event := new(SeverRegister)
	if err := _Sever.contract.UnpackLog(event, "Register", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeverWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Sever contract.
type SeverWithdrawIterator struct {
	Event *SeverWithdraw // Event containing the contract specifics and raw log

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
func (it *SeverWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeverWithdraw)
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
		it.Event = new(SeverWithdraw)
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
func (it *SeverWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeverWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeverWithdraw represents a Withdraw event raised by the Sever contract.
type SeverWithdraw struct {
	Guarantor common.Address
	To        common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address guarantor, address _to, uint256 amount)
func (_Sever *SeverFilterer) FilterWithdraw(opts *bind.FilterOpts) (*SeverWithdrawIterator, error) {

	logs, sub, err := _Sever.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &SeverWithdrawIterator{contract: _Sever.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address guarantor, address _to, uint256 amount)
func (_Sever *SeverFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SeverWithdraw) (event.Subscription, error) {

	logs, sub, err := _Sever.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeverWithdraw)
				if err := _Sever.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address guarantor, address _to, uint256 amount)
func (_Sever *SeverFilterer) ParseWithdraw(log types.Log) (*SeverWithdraw, error) {
	event := new(SeverWithdraw)
	if err := _Sever.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
