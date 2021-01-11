// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package logger

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

// LoggerABI is the input ABI used to generate the binding from.
const LoggerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"Log\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"dataLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// LoggerBin is the compiled bytecode used for deploying new contracts.
var LoggerBin = "0x608060405234801561001057600080fd5b506102d6806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806353f2b23714610030575b600080fd5b61004a60048036038101906100459190610105565b61004c565b005b3373ffffffffffffffffffffffffffffffffffffffff167f0738f4da267a110d810e6e89fc59e46be6de0c37b1d5cd559b267dc3688e74e082604051610092919061017f565b60405180910390a250565b60006100b06100ab846101d2565b6101a1565b9050828152602081018484840111156100c857600080fd5b6100d384828561021e565b509392505050565b600082601f8301126100ec57600080fd5b81356100fc84826020860161009d565b91505092915050565b60006020828403121561011757600080fd5b600082013567ffffffffffffffff81111561013157600080fd5b61013d848285016100db565b91505092915050565b600061015182610202565b61015b818561020d565b935061016b81856020860161022d565b6101748161028f565b840191505092915050565b600060208201905081810360008301526101998184610146565b905092915050565b6000604051905081810181811067ffffffffffffffff821117156101c8576101c7610260565b5b8060405250919050565b600067ffffffffffffffff8211156101ed576101ec610260565b5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b82818337600083830152505050565b60005b8381101561024b578082015181840152602081019050610230565b8381111561025a576000848401525b50505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f830116905091905056fea264697066735822122051d1d54c6e6dd1de8e09ed351a131f6182aa186ab4eb606a5b259e33925c696464736f6c63430008000033"

// DeployLogger deploys a new Ethereum contract, binding an instance of Logger to it.
func DeployLogger(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Logger, error) {
	parsed, err := abi.JSON(strings.NewReader(LoggerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LoggerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Logger{LoggerCaller: LoggerCaller{contract: contract}, LoggerTransactor: LoggerTransactor{contract: contract}, LoggerFilterer: LoggerFilterer{contract: contract}}, nil
}

// Logger is an auto generated Go binding around an Ethereum contract.
type Logger struct {
	LoggerCaller     // Read-only binding to the contract
	LoggerTransactor // Write-only binding to the contract
	LoggerFilterer   // Log filterer for contract events
}

// LoggerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoggerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoggerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoggerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoggerSession struct {
	Contract     *Logger           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoggerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoggerCallerSession struct {
	Contract *LoggerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LoggerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoggerTransactorSession struct {
	Contract     *LoggerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoggerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoggerRaw struct {
	Contract *Logger // Generic contract binding to access the raw methods on
}

// LoggerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoggerCallerRaw struct {
	Contract *LoggerCaller // Generic read-only contract binding to access the raw methods on
}

// LoggerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoggerTransactorRaw struct {
	Contract *LoggerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLogger creates a new instance of Logger, bound to a specific deployed contract.
func NewLogger(address common.Address, backend bind.ContractBackend) (*Logger, error) {
	contract, err := bindLogger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Logger{LoggerCaller: LoggerCaller{contract: contract}, LoggerTransactor: LoggerTransactor{contract: contract}, LoggerFilterer: LoggerFilterer{contract: contract}}, nil
}

// NewLoggerCaller creates a new read-only instance of Logger, bound to a specific deployed contract.
func NewLoggerCaller(address common.Address, caller bind.ContractCaller) (*LoggerCaller, error) {
	contract, err := bindLogger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoggerCaller{contract: contract}, nil
}

// NewLoggerTransactor creates a new write-only instance of Logger, bound to a specific deployed contract.
func NewLoggerTransactor(address common.Address, transactor bind.ContractTransactor) (*LoggerTransactor, error) {
	contract, err := bindLogger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoggerTransactor{contract: contract}, nil
}

// NewLoggerFilterer creates a new log filterer instance of Logger, bound to a specific deployed contract.
func NewLoggerFilterer(address common.Address, filterer bind.ContractFilterer) (*LoggerFilterer, error) {
	contract, err := bindLogger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoggerFilterer{contract: contract}, nil
}

// bindLogger binds a generic wrapper to an already deployed contract.
func bindLogger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LoggerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logger *LoggerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Logger.Contract.LoggerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logger *LoggerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logger.Contract.LoggerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logger *LoggerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Logger.Contract.LoggerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logger *LoggerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Logger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logger *LoggerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logger *LoggerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Logger.Contract.contract.Transact(opts, method, params...)
}

// DataLog is a paid mutator transaction binding the contract method 0x53f2b237.
//
// Solidity: function dataLog(string data) returns()
func (_Logger *LoggerTransactor) DataLog(opts *bind.TransactOpts, data string) (*types.Transaction, error) {
	return _Logger.contract.Transact(opts, "dataLog", data)
}

// DataLog is a paid mutator transaction binding the contract method 0x53f2b237.
//
// Solidity: function dataLog(string data) returns()
func (_Logger *LoggerSession) DataLog(data string) (*types.Transaction, error) {
	return _Logger.Contract.DataLog(&_Logger.TransactOpts, data)
}

// DataLog is a paid mutator transaction binding the contract method 0x53f2b237.
//
// Solidity: function dataLog(string data) returns()
func (_Logger *LoggerTransactorSession) DataLog(data string) (*types.Transaction, error) {
	return _Logger.Contract.DataLog(&_Logger.TransactOpts, data)
}

// LoggerLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the Logger contract.
type LoggerLogIterator struct {
	Event *LoggerLog // Event containing the contract specifics and raw log

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
func (it *LoggerLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerLog)
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
		it.Event = new(LoggerLog)
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
func (it *LoggerLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerLog represents a Log event raised by the Logger contract.
type LoggerLog struct {
	Sender common.Address
	Data   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x0738f4da267a110d810e6e89fc59e46be6de0c37b1d5cd559b267dc3688e74e0.
//
// Solidity: event Log(address indexed sender, string data)
func (_Logger *LoggerFilterer) FilterLog(opts *bind.FilterOpts, sender []common.Address) (*LoggerLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Logger.contract.FilterLogs(opts, "Log", senderRule)
	if err != nil {
		return nil, err
	}
	return &LoggerLogIterator{contract: _Logger.contract, event: "Log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x0738f4da267a110d810e6e89fc59e46be6de0c37b1d5cd559b267dc3688e74e0.
//
// Solidity: event Log(address indexed sender, string data)
func (_Logger *LoggerFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *LoggerLog, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Logger.contract.WatchLogs(opts, "Log", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerLog)
				if err := _Logger.contract.UnpackLog(event, "Log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0x0738f4da267a110d810e6e89fc59e46be6de0c37b1d5cd559b267dc3688e74e0.
//
// Solidity: event Log(address indexed sender, string data)
func (_Logger *LoggerFilterer) ParseLog(log types.Log) (*LoggerLog, error) {
	event := new(LoggerLog)
	if err := _Logger.contract.UnpackLog(event, "Log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
