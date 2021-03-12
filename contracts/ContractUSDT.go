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

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_A\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_B\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_judge\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_Server\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_Content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_feePercentLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"}],\"name\":\"Intervene\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"A\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"B\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"content\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePercentLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"}],\"name\":\"intervene\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"judge\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"server\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(address)
func (_Contract *ContractCaller) A(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(address)
func (_Contract *ContractSession) A() (common.Address, error) {
	return _Contract.Contract.A(&_Contract.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(address)
func (_Contract *ContractCallerSession) A() (common.Address, error) {
	return _Contract.Contract.A(&_Contract.CallOpts)
}

// B is a free data retrieval call binding the contract method 0x32e7c5bf.
//
// Solidity: function B() view returns(address)
func (_Contract *ContractCaller) B(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "B")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// B is a free data retrieval call binding the contract method 0x32e7c5bf.
//
// Solidity: function B() view returns(address)
func (_Contract *ContractSession) B() (common.Address, error) {
	return _Contract.Contract.B(&_Contract.CallOpts)
}

// B is a free data retrieval call binding the contract method 0x32e7c5bf.
//
// Solidity: function B() view returns(address)
func (_Contract *ContractCallerSession) B() (common.Address, error) {
	return _Contract.Contract.B(&_Contract.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Contract *ContractCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Contract *ContractSession) Balance() (*big.Int, error) {
	return _Contract.Contract.Balance(&_Contract.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Contract *ContractCallerSession) Balance() (*big.Int, error) {
	return _Contract.Contract.Balance(&_Contract.CallOpts)
}

// Content is a free data retrieval call binding the contract method 0x8a4d5a67.
//
// Solidity: function content() view returns(string)
func (_Contract *ContractCaller) Content(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "content")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Content is a free data retrieval call binding the contract method 0x8a4d5a67.
//
// Solidity: function content() view returns(string)
func (_Contract *ContractSession) Content() (string, error) {
	return _Contract.Contract.Content(&_Contract.CallOpts)
}

// Content is a free data retrieval call binding the contract method 0x8a4d5a67.
//
// Solidity: function content() view returns(string)
func (_Contract *ContractCallerSession) Content() (string, error) {
	return _Contract.Contract.Content(&_Contract.CallOpts)
}

// FeePercentLimit is a free data retrieval call binding the contract method 0x06a3097c.
//
// Solidity: function feePercentLimit() view returns(uint256)
func (_Contract *ContractCaller) FeePercentLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "feePercentLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePercentLimit is a free data retrieval call binding the contract method 0x06a3097c.
//
// Solidity: function feePercentLimit() view returns(uint256)
func (_Contract *ContractSession) FeePercentLimit() (*big.Int, error) {
	return _Contract.Contract.FeePercentLimit(&_Contract.CallOpts)
}

// FeePercentLimit is a free data retrieval call binding the contract method 0x06a3097c.
//
// Solidity: function feePercentLimit() view returns(uint256)
func (_Contract *ContractCallerSession) FeePercentLimit() (*big.Int, error) {
	return _Contract.Contract.FeePercentLimit(&_Contract.CallOpts)
}

// Judge is a free data retrieval call binding the contract method 0x573255f4.
//
// Solidity: function judge() view returns(address)
func (_Contract *ContractCaller) Judge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "judge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Judge is a free data retrieval call binding the contract method 0x573255f4.
//
// Solidity: function judge() view returns(address)
func (_Contract *ContractSession) Judge() (common.Address, error) {
	return _Contract.Contract.Judge(&_Contract.CallOpts)
}

// Judge is a free data retrieval call binding the contract method 0x573255f4.
//
// Solidity: function judge() view returns(address)
func (_Contract *ContractCallerSession) Judge() (common.Address, error) {
	return _Contract.Contract.Judge(&_Contract.CallOpts)
}

// Server is a free data retrieval call binding the contract method 0xfd922a42.
//
// Solidity: function server() view returns(address)
func (_Contract *ContractCaller) Server(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "server")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Server is a free data retrieval call binding the contract method 0xfd922a42.
//
// Solidity: function server() view returns(address)
func (_Contract *ContractSession) Server() (common.Address, error) {
	return _Contract.Contract.Server(&_Contract.CallOpts)
}

// Server is a free data retrieval call binding the contract method 0xfd922a42.
//
// Solidity: function server() view returns(address)
func (_Contract *ContractCallerSession) Server() (common.Address, error) {
	return _Contract.Contract.Server(&_Contract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Contract *ContractCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Contract *ContractSession) Token() (common.Address, error) {
	return _Contract.Contract.Token(&_Contract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Contract *ContractCallerSession) Token() (common.Address, error) {
	return _Contract.Contract.Token(&_Contract.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Contract *ContractTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Contract *ContractSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Contract *ContractTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, amount)
}

// Intervene is a paid mutator transaction binding the contract method 0xc5a90628.
//
// Solidity: function intervene(uint256 amountToA, uint256 amountToB, uint256 feePercent) returns()
func (_Contract *ContractTransactor) Intervene(opts *bind.TransactOpts, amountToA *big.Int, amountToB *big.Int, feePercent *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "intervene", amountToA, amountToB, feePercent)
}

// Intervene is a paid mutator transaction binding the contract method 0xc5a90628.
//
// Solidity: function intervene(uint256 amountToA, uint256 amountToB, uint256 feePercent) returns()
func (_Contract *ContractSession) Intervene(amountToA *big.Int, amountToB *big.Int, feePercent *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Intervene(&_Contract.TransactOpts, amountToA, amountToB, feePercent)
}

// Intervene is a paid mutator transaction binding the contract method 0xc5a90628.
//
// Solidity: function intervene(uint256 amountToA, uint256 amountToB, uint256 feePercent) returns()
func (_Contract *ContractTransactorSession) Intervene(amountToA *big.Int, amountToB *big.Int, feePercent *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Intervene(&_Contract.TransactOpts, amountToA, amountToB, feePercent)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa41fe49f.
//
// Solidity: function withdraw(uint256 amountToA, uint256 amountToB, uint256 feePercent) returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, amountToA *big.Int, amountToB *big.Int, feePercent *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw", amountToA, amountToB, feePercent)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa41fe49f.
//
// Solidity: function withdraw(uint256 amountToA, uint256 amountToB, uint256 feePercent) returns()
func (_Contract *ContractSession) Withdraw(amountToA *big.Int, amountToB *big.Int, feePercent *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, amountToA, amountToB, feePercent)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa41fe49f.
//
// Solidity: function withdraw(uint256 amountToA, uint256 amountToB, uint256 feePercent) returns()
func (_Contract *ContractTransactorSession) Withdraw(amountToA *big.Int, amountToB *big.Int, feePercent *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, amountToA, amountToB, feePercent)
}

// ContractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Contract contract.
type ContractDepositIterator struct {
	Event *ContractDeposit // Event containing the contract specifics and raw log

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
func (it *ContractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDeposit)
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
		it.Event = new(ContractDeposit)
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
func (it *ContractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDeposit represents a Deposit event raised by the Contract contract.
type ContractDeposit struct {
	Time   *big.Int
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xeaa18152488ce5959073c9c79c88ca90b3d96c00de1f118cfaad664c3dab06b9.
//
// Solidity: event Deposit(uint256 time, address from, uint256 amount)
func (_Contract *ContractFilterer) FilterDeposit(opts *bind.FilterOpts) (*ContractDepositIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &ContractDepositIterator{contract: _Contract.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xeaa18152488ce5959073c9c79c88ca90b3d96c00de1f118cfaad664c3dab06b9.
//
// Solidity: event Deposit(uint256 time, address from, uint256 amount)
func (_Contract *ContractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ContractDeposit) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDeposit)
				if err := _Contract.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xeaa18152488ce5959073c9c79c88ca90b3d96c00de1f118cfaad664c3dab06b9.
//
// Solidity: event Deposit(uint256 time, address from, uint256 amount)
func (_Contract *ContractFilterer) ParseDeposit(log types.Log) (*ContractDeposit, error) {
	event := new(ContractDeposit)
	if err := _Contract.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInterveneIterator is returned from FilterIntervene and is used to iterate over the raw logs and unpacked data for Intervene events raised by the Contract contract.
type ContractInterveneIterator struct {
	Event *ContractIntervene // Event containing the contract specifics and raw log

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
func (it *ContractInterveneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIntervene)
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
		it.Event = new(ContractIntervene)
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
func (it *ContractInterveneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInterveneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIntervene represents a Intervene event raised by the Contract contract.
type ContractIntervene struct {
	Time       *big.Int
	AmountToA  *big.Int
	AmountToB  *big.Int
	FeePercent *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterIntervene is a free log retrieval operation binding the contract event 0xb9e9e6f3028f9a28cd3222d51b668eb386c96682724c9d61ef77394784200076.
//
// Solidity: event Intervene(uint256 time, uint256 amountToA, uint256 amountToB, uint256 feePercent)
func (_Contract *ContractFilterer) FilterIntervene(opts *bind.FilterOpts) (*ContractInterveneIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Intervene")
	if err != nil {
		return nil, err
	}
	return &ContractInterveneIterator{contract: _Contract.contract, event: "Intervene", logs: logs, sub: sub}, nil
}

// WatchIntervene is a free log subscription operation binding the contract event 0xb9e9e6f3028f9a28cd3222d51b668eb386c96682724c9d61ef77394784200076.
//
// Solidity: event Intervene(uint256 time, uint256 amountToA, uint256 amountToB, uint256 feePercent)
func (_Contract *ContractFilterer) WatchIntervene(opts *bind.WatchOpts, sink chan<- *ContractIntervene) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Intervene")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIntervene)
				if err := _Contract.contract.UnpackLog(event, "Intervene", log); err != nil {
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

// ParseIntervene is a log parse operation binding the contract event 0xb9e9e6f3028f9a28cd3222d51b668eb386c96682724c9d61ef77394784200076.
//
// Solidity: event Intervene(uint256 time, uint256 amountToA, uint256 amountToB, uint256 feePercent)
func (_Contract *ContractFilterer) ParseIntervene(log types.Log) (*ContractIntervene, error) {
	event := new(ContractIntervene)
	if err := _Contract.contract.UnpackLog(event, "Intervene", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Contract contract.
type ContractWithdrawIterator struct {
	Event *ContractWithdraw // Event containing the contract specifics and raw log

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
func (it *ContractWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWithdraw)
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
		it.Event = new(ContractWithdraw)
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
func (it *ContractWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWithdraw represents a Withdraw event raised by the Contract contract.
type ContractWithdraw struct {
	Time       *big.Int
	AmountToA  *big.Int
	AmountToB  *big.Int
	FeePercent *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xb088a38065873c9d2da855cf4a2e4aa012b28e45ead53c5b9edbcf71614cfb23.
//
// Solidity: event Withdraw(uint256 time, uint256 amountToA, uint256 amountToB, uint256 feePercent)
func (_Contract *ContractFilterer) FilterWithdraw(opts *bind.FilterOpts) (*ContractWithdrawIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &ContractWithdrawIterator{contract: _Contract.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xb088a38065873c9d2da855cf4a2e4aa012b28e45ead53c5b9edbcf71614cfb23.
//
// Solidity: event Withdraw(uint256 time, uint256 amountToA, uint256 amountToB, uint256 feePercent)
func (_Contract *ContractFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ContractWithdraw) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWithdraw)
				if err := _Contract.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xb088a38065873c9d2da855cf4a2e4aa012b28e45ead53c5b9edbcf71614cfb23.
//
// Solidity: event Withdraw(uint256 time, uint256 amountToA, uint256 amountToB, uint256 feePercent)
func (_Contract *ContractFilterer) ParseWithdraw(log types.Log) (*ContractWithdraw, error) {
	event := new(ContractWithdraw)
	if err := _Contract.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
