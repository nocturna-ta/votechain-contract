// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package votechainBase

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

// VotechainBaseMetaData contains all meta data concerning the VotechainBase contract.
var VotechainBaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"OnlyKpuAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotingNotActive\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"VotingStatusChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"kpuAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setKpuAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setVotingStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506104068061005b5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063408e27271461004e5780637478c9fe1461006c5780639df86dc114610088578063fb4ab164146100a4575b5f5ffd5b6100566100c2565b60405161006391906102b1565b60405180910390f35b610086600480360381019061008191906102f8565b6100d4565b005b6100a2600480360381019061009d919061037d565b6101ac565b005b6100ac610273565b6040516100b991906103b7565b60405180910390f35b5f60149054906101000a900460ff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610159576040517f8b8958dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f60146101000a81548160ff0219169083151502179055507f9069a1a16ace751e8690f383e12f87b01e8488ba387e626810bd113fef0417f9816040516101a191906102b1565b60405180910390a150565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610231576040517f8b8958dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8115159050919050565b6102ab81610297565b82525050565b5f6020820190506102c45f8301846102a2565b92915050565b5f5ffd5b6102d781610297565b81146102e1575f5ffd5b50565b5f813590506102f2816102ce565b92915050565b5f6020828403121561030d5761030c6102ca565b5b5f61031a848285016102e4565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61034c82610323565b9050919050565b61035c81610342565b8114610366575f5ffd5b50565b5f8135905061037781610353565b92915050565b5f60208284031215610392576103916102ca565b5b5f61039f84828501610369565b91505092915050565b6103b181610342565b82525050565b5f6020820190506103ca5f8301846103a8565b9291505056fea26469706673582212200821fbf30f108caca03e130399f225bbec89ff5c39e13ee7f28b29b369751b4964736f6c634300081c0033",
}

// VotechainBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use VotechainBaseMetaData.ABI instead.
var VotechainBaseABI = VotechainBaseMetaData.ABI

// VotechainBaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VotechainBaseMetaData.Bin instead.
var VotechainBaseBin = VotechainBaseMetaData.Bin

// DeployVotechainBase deploys a new Ethereum contract, binding an instance of VotechainBase to it.
func DeployVotechainBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VotechainBase, error) {
	parsed, err := VotechainBaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VotechainBaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VotechainBase{VotechainBaseCaller: VotechainBaseCaller{contract: contract}, VotechainBaseTransactor: VotechainBaseTransactor{contract: contract}, VotechainBaseFilterer: VotechainBaseFilterer{contract: contract}}, nil
}

// VotechainBase is an auto generated Go binding around an Ethereum contract.
type VotechainBase struct {
	VotechainBaseCaller     // Read-only binding to the contract
	VotechainBaseTransactor // Write-only binding to the contract
	VotechainBaseFilterer   // Log filterer for contract events
}

// VotechainBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotechainBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotechainBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotechainBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotechainBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotechainBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotechainBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotechainBaseSession struct {
	Contract     *VotechainBase    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotechainBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotechainBaseCallerSession struct {
	Contract *VotechainBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VotechainBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotechainBaseTransactorSession struct {
	Contract     *VotechainBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VotechainBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotechainBaseRaw struct {
	Contract *VotechainBase // Generic contract binding to access the raw methods on
}

// VotechainBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotechainBaseCallerRaw struct {
	Contract *VotechainBaseCaller // Generic read-only contract binding to access the raw methods on
}

// VotechainBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotechainBaseTransactorRaw struct {
	Contract *VotechainBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotechainBase creates a new instance of VotechainBase, bound to a specific deployed contract.
func NewVotechainBase(address common.Address, backend bind.ContractBackend) (*VotechainBase, error) {
	contract, err := bindVotechainBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VotechainBase{VotechainBaseCaller: VotechainBaseCaller{contract: contract}, VotechainBaseTransactor: VotechainBaseTransactor{contract: contract}, VotechainBaseFilterer: VotechainBaseFilterer{contract: contract}}, nil
}

// NewVotechainBaseCaller creates a new read-only instance of VotechainBase, bound to a specific deployed contract.
func NewVotechainBaseCaller(address common.Address, caller bind.ContractCaller) (*VotechainBaseCaller, error) {
	contract, err := bindVotechainBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotechainBaseCaller{contract: contract}, nil
}

// NewVotechainBaseTransactor creates a new write-only instance of VotechainBase, bound to a specific deployed contract.
func NewVotechainBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*VotechainBaseTransactor, error) {
	contract, err := bindVotechainBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotechainBaseTransactor{contract: contract}, nil
}

// NewVotechainBaseFilterer creates a new log filterer instance of VotechainBase, bound to a specific deployed contract.
func NewVotechainBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*VotechainBaseFilterer, error) {
	contract, err := bindVotechainBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotechainBaseFilterer{contract: contract}, nil
}

// bindVotechainBase binds a generic wrapper to an already deployed contract.
func bindVotechainBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotechainBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotechainBase *VotechainBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotechainBase.Contract.VotechainBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotechainBase *VotechainBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotechainBase.Contract.VotechainBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotechainBase *VotechainBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotechainBase.Contract.VotechainBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotechainBase *VotechainBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotechainBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotechainBase *VotechainBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotechainBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotechainBase *VotechainBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotechainBase.Contract.contract.Transact(opts, method, params...)
}

// KpuAdmin is a free data retrieval call binding the contract method 0xfb4ab164.
//
// Solidity: function kpuAdmin() view returns(address)
func (_VotechainBase *VotechainBaseCaller) KpuAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotechainBase.contract.Call(opts, &out, "kpuAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KpuAdmin is a free data retrieval call binding the contract method 0xfb4ab164.
//
// Solidity: function kpuAdmin() view returns(address)
func (_VotechainBase *VotechainBaseSession) KpuAdmin() (common.Address, error) {
	return _VotechainBase.Contract.KpuAdmin(&_VotechainBase.CallOpts)
}

// KpuAdmin is a free data retrieval call binding the contract method 0xfb4ab164.
//
// Solidity: function kpuAdmin() view returns(address)
func (_VotechainBase *VotechainBaseCallerSession) KpuAdmin() (common.Address, error) {
	return _VotechainBase.Contract.KpuAdmin(&_VotechainBase.CallOpts)
}

// VotingActive is a free data retrieval call binding the contract method 0x408e2727.
//
// Solidity: function votingActive() view returns(bool)
func (_VotechainBase *VotechainBaseCaller) VotingActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VotechainBase.contract.Call(opts, &out, "votingActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VotingActive is a free data retrieval call binding the contract method 0x408e2727.
//
// Solidity: function votingActive() view returns(bool)
func (_VotechainBase *VotechainBaseSession) VotingActive() (bool, error) {
	return _VotechainBase.Contract.VotingActive(&_VotechainBase.CallOpts)
}

// VotingActive is a free data retrieval call binding the contract method 0x408e2727.
//
// Solidity: function votingActive() view returns(bool)
func (_VotechainBase *VotechainBaseCallerSession) VotingActive() (bool, error) {
	return _VotechainBase.Contract.VotingActive(&_VotechainBase.CallOpts)
}

// SetKpuAdmin is a paid mutator transaction binding the contract method 0x9df86dc1.
//
// Solidity: function setKpuAdmin(address newAdmin) returns()
func (_VotechainBase *VotechainBaseTransactor) SetKpuAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _VotechainBase.contract.Transact(opts, "setKpuAdmin", newAdmin)
}

// SetKpuAdmin is a paid mutator transaction binding the contract method 0x9df86dc1.
//
// Solidity: function setKpuAdmin(address newAdmin) returns()
func (_VotechainBase *VotechainBaseSession) SetKpuAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _VotechainBase.Contract.SetKpuAdmin(&_VotechainBase.TransactOpts, newAdmin)
}

// SetKpuAdmin is a paid mutator transaction binding the contract method 0x9df86dc1.
//
// Solidity: function setKpuAdmin(address newAdmin) returns()
func (_VotechainBase *VotechainBaseTransactorSession) SetKpuAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _VotechainBase.Contract.SetKpuAdmin(&_VotechainBase.TransactOpts, newAdmin)
}

// SetVotingStatus is a paid mutator transaction binding the contract method 0x7478c9fe.
//
// Solidity: function setVotingStatus(bool status) returns()
func (_VotechainBase *VotechainBaseTransactor) SetVotingStatus(opts *bind.TransactOpts, status bool) (*types.Transaction, error) {
	return _VotechainBase.contract.Transact(opts, "setVotingStatus", status)
}

// SetVotingStatus is a paid mutator transaction binding the contract method 0x7478c9fe.
//
// Solidity: function setVotingStatus(bool status) returns()
func (_VotechainBase *VotechainBaseSession) SetVotingStatus(status bool) (*types.Transaction, error) {
	return _VotechainBase.Contract.SetVotingStatus(&_VotechainBase.TransactOpts, status)
}

// SetVotingStatus is a paid mutator transaction binding the contract method 0x7478c9fe.
//
// Solidity: function setVotingStatus(bool status) returns()
func (_VotechainBase *VotechainBaseTransactorSession) SetVotingStatus(status bool) (*types.Transaction, error) {
	return _VotechainBase.Contract.SetVotingStatus(&_VotechainBase.TransactOpts, status)
}

// VotechainBaseVotingStatusChangedIterator is returned from FilterVotingStatusChanged and is used to iterate over the raw logs and unpacked data for VotingStatusChanged events raised by the VotechainBase contract.
type VotechainBaseVotingStatusChangedIterator struct {
	Event *VotechainBaseVotingStatusChanged // Event containing the contract specifics and raw log

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
func (it *VotechainBaseVotingStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotechainBaseVotingStatusChanged)
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
		it.Event = new(VotechainBaseVotingStatusChanged)
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
func (it *VotechainBaseVotingStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotechainBaseVotingStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotechainBaseVotingStatusChanged represents a VotingStatusChanged event raised by the VotechainBase contract.
type VotechainBaseVotingStatusChanged struct {
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVotingStatusChanged is a free log retrieval operation binding the contract event 0x9069a1a16ace751e8690f383e12f87b01e8488ba387e626810bd113fef0417f9.
//
// Solidity: event VotingStatusChanged(bool isActive)
func (_VotechainBase *VotechainBaseFilterer) FilterVotingStatusChanged(opts *bind.FilterOpts) (*VotechainBaseVotingStatusChangedIterator, error) {

	logs, sub, err := _VotechainBase.contract.FilterLogs(opts, "VotingStatusChanged")
	if err != nil {
		return nil, err
	}
	return &VotechainBaseVotingStatusChangedIterator{contract: _VotechainBase.contract, event: "VotingStatusChanged", logs: logs, sub: sub}, nil
}

// WatchVotingStatusChanged is a free log subscription operation binding the contract event 0x9069a1a16ace751e8690f383e12f87b01e8488ba387e626810bd113fef0417f9.
//
// Solidity: event VotingStatusChanged(bool isActive)
func (_VotechainBase *VotechainBaseFilterer) WatchVotingStatusChanged(opts *bind.WatchOpts, sink chan<- *VotechainBaseVotingStatusChanged) (event.Subscription, error) {

	logs, sub, err := _VotechainBase.contract.WatchLogs(opts, "VotingStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotechainBaseVotingStatusChanged)
				if err := _VotechainBase.contract.UnpackLog(event, "VotingStatusChanged", log); err != nil {
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

// ParseVotingStatusChanged is a log parse operation binding the contract event 0x9069a1a16ace751e8690f383e12f87b01e8488ba387e626810bd113fef0417f9.
//
// Solidity: event VotingStatusChanged(bool isActive)
func (_VotechainBase *VotechainBaseFilterer) ParseVotingStatusChanged(log types.Log) (*VotechainBaseVotingStatusChanged, error) {
	event := new(VotechainBaseVotingStatusChanged)
	if err := _VotechainBase.contract.UnpackLog(event, "VotingStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
