// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package votechain

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

// VotechainMetaData contains all meta data concerning the Votechain contract.
var VotechainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_kpuManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voterManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_electionManagerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyVoted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidElection\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VoterNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotingNotActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"}],\"name\":\"addElection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"contractIVotechainBase\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"electionManager\",\"outputs\":[{\"internalType\":\"contractIElectionManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kpuManager\",\"outputs\":[{\"internalType\":\"contractIKPUManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"registerKPUKota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"registerKPUProvinsi\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"}],\"name\":\"registerVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setKpuAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setVotingStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"}],\"name\":\"toggleElectionActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterManager\",\"outputs\":[{\"internalType\":\"contractIVoterManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040516114c33803806114c383398181016040528101906100319190610197565b835f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050506101fb565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101668261013d565b9050919050565b6101768161015c565b8114610180575f5ffd5b50565b5f815190506101918161016d565b92915050565b5f5f5f5f608085870312156101af576101ae610139565b5b5f6101bc87828801610183565b94505060206101cd87828801610183565b93505060406101de87828801610183565b92505060606101ef87828801610183565b91505092959194509250565b6112bb806102085f395ff3fe608060405234801561000f575f5ffd5b50600436106100b2575f3560e01c80637da329491161006f5780637da329491461016457806392423815146101825780639df86dc11461019e578063db5dfdd7146101ba578063fc36e15b146101d6578063fd1c6e61146101f2576100b2565b80632db33a2f146100b657806334d1d2f2146100d25780634a075de2146100f05780635001f3b51461010c57806365910a4f1461012a5780637478c9fe14610148575b5f5ffd5b6100d060048036038101906100cb9190610a5b565b61020e565b005b6100da6102a4565b6040516100e79190610b47565b60405180910390f35b61010a60048036038101906101059190610b60565b6102c9565b005b610114610359565b6040516101219190610bdd565b60405180910390f35b61013261037d565b60405161013f9190610c16565b60405180910390f35b610162600480360381019061015d9190610c64565b6103a2565b005b61016c61042b565b6040516101799190610caf565b60405180910390f35b61019c60048036038101906101979190610a5b565b610450565b005b6101b860048036038101906101b39190610cc8565b6104e6565b005b6101d460048036038101906101cf9190610cf3565b61056f565b005b6101f060048036038101906101eb9190610cf3565b6105fc565b005b61020c60048036038101906102079190610d3e565b6108f6565b005b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632db33a2f86868686866040518663ffffffff1660e01b8152600401610270959493929190610e57565b5f604051808303815f87803b158015610287575f5ffd5b505af1158015610299573d5f5f3e3d5ffd5b505050505050505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634a075de28484846040518463ffffffff1660e01b815260040161032793929190610e9e565b5f604051808303815f87803b15801561033e575f5ffd5b505af1158015610350573d5f5f3e3d5ffd5b50505050505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637478c9fe826040518263ffffffff1660e01b81526004016103fb9190610edd565b5f604051808303815f87803b158015610412575f5ffd5b505af1158015610424573d5f5f3e3d5ffd5b5050505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639242381586868686866040518663ffffffff1660e01b81526004016104b2959493929190610e57565b5f604051808303815f87803b1580156104c9575f5ffd5b505af11580156104db573d5f5f3e3d5ffd5b505050505050505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639df86dc1826040518263ffffffff1660e01b815260040161053f9190610ef6565b5f604051808303815f87803b158015610556575f5ffd5b505af1158015610568573d5f5f3e3d5ffd5b5050505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663db5dfdd783836040518363ffffffff1660e01b81526004016105cb929190610f0f565b5f604051808303815f87803b1580156105e2575f5ffd5b505af11580156105f4573d5f5f3e3d5ffd5b505050505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663408e27276040518163ffffffff1660e01b8152600401602060405180830381865afa158015610665573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106899190610f45565b6106bf576040517f9b8cc47500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634bdd7585336040518263ffffffff1660e01b81526004016107199190610ef6565b5f60405180830381865afa92505050801561075657506040513d5f823e3d601f19601f820116820180604052508101906107539190611175565b60015b61078c576040517f6f08c58700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060400151156107c8576040517f7c9a1cf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a8fd0cfa336040518263ffffffff1660e01b81526004016108229190610ef6565b6020604051808303815f875af115801561083e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108629190610f45565b5060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e8d5940d8484845f01516040518463ffffffff1660e01b81526004016108c4939291906111fe565b5f604051808303815f87803b1580156108db575f5ffd5b505af11580156108ed573d5f5f3e3d5ffd5b50505050505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fd1c6e618787878787876040518763ffffffff1660e01b815260040161095a96959493929190611235565b5f604051808303815f87803b158015610971575f5ffd5b505af1158015610983573d5f5f3e3d5ffd5b50505050505050505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6109c9826109a0565b9050919050565b6109d9816109bf565b81146109e3575f5ffd5b50565b5f813590506109f4816109d0565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112610a1b57610a1a6109fa565b5b8235905067ffffffffffffffff811115610a3857610a376109fe565b5b602083019150836001820283011115610a5457610a53610a02565b5b9250929050565b5f5f5f5f5f60608688031215610a7457610a73610998565b5b5f610a81888289016109e6565b955050602086013567ffffffffffffffff811115610aa257610aa161099c565b5b610aae88828901610a06565b9450945050604086013567ffffffffffffffff811115610ad157610ad061099c565b5b610add88828901610a06565b92509250509295509295909350565b5f819050919050565b5f610b0f610b0a610b05846109a0565b610aec565b6109a0565b9050919050565b5f610b2082610af5565b9050919050565b5f610b3182610b16565b9050919050565b610b4181610b27565b82525050565b5f602082019050610b5a5f830184610b38565b92915050565b5f5f5f60408486031215610b7757610b76610998565b5b5f84013567ffffffffffffffff811115610b9457610b9361099c565b5b610ba086828701610a06565b93509350506020610bb3868287016109e6565b9150509250925092565b5f610bc782610b16565b9050919050565b610bd781610bbd565b82525050565b5f602082019050610bf05f830184610bce565b92915050565b5f610c0082610b16565b9050919050565b610c1081610bf6565b82525050565b5f602082019050610c295f830184610c07565b92915050565b5f8115159050919050565b610c4381610c2f565b8114610c4d575f5ffd5b50565b5f81359050610c5e81610c3a565b92915050565b5f60208284031215610c7957610c78610998565b5b5f610c8684828501610c50565b91505092915050565b5f610c9982610b16565b9050919050565b610ca981610c8f565b82525050565b5f602082019050610cc25f830184610ca0565b92915050565b5f60208284031215610cdd57610cdc610998565b5b5f610cea848285016109e6565b91505092915050565b5f5f60208385031215610d0957610d08610998565b5b5f83013567ffffffffffffffff811115610d2657610d2561099c565b5b610d3285828601610a06565b92509250509250929050565b5f5f5f5f5f5f60608789031215610d5857610d57610998565b5b5f87013567ffffffffffffffff811115610d7557610d7461099c565b5b610d8189828a01610a06565b9650965050602087013567ffffffffffffffff811115610da457610da361099c565b5b610db089828a01610a06565b9450945050604087013567ffffffffffffffff811115610dd357610dd261099c565b5b610ddf89828a01610a06565b92509250509295509295509295565b610df7816109bf565b82525050565b5f82825260208201905092915050565b828183375f83830152505050565b5f601f19601f8301169050919050565b5f610e368385610dfd565b9350610e43838584610e0d565b610e4c83610e1b565b840190509392505050565b5f606082019050610e6a5f830188610dee565b8181036020830152610e7d818688610e2b565b90508181036040830152610e92818486610e2b565b90509695505050505050565b5f6040820190508181035f830152610eb7818587610e2b565b9050610ec66020830184610dee565b949350505050565b610ed781610c2f565b82525050565b5f602082019050610ef05f830184610ece565b92915050565b5f602082019050610f095f830184610dee565b92915050565b5f6020820190508181035f830152610f28818486610e2b565b90509392505050565b5f81519050610f3f81610c3a565b92915050565b5f60208284031215610f5a57610f59610998565b5b5f610f6784828501610f31565b91505092915050565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610faa82610e1b565b810181811067ffffffffffffffff82111715610fc957610fc8610f74565b5b80604052505050565b5f610fdb61098f565b9050610fe78282610fa1565b919050565b5f5ffd5b5f5ffd5b5f67ffffffffffffffff82111561100e5761100d610f74565b5b61101782610e1b565b9050602081019050919050565b8281835e5f83830152505050565b5f61104461103f84610ff4565b610fd2565b9050828152602081018484840111156110605761105f610ff0565b5b61106b848285611024565b509392505050565b5f82601f830112611087576110866109fa565b5b8151611097848260208601611032565b91505092915050565b5f815190506110ae816109d0565b92915050565b5f60a082840312156110c9576110c8610f70565b5b6110d360a0610fd2565b90505f82015167ffffffffffffffff8111156110f2576110f1610fec565b5b6110fe84828501611073565b5f830152506020611111848285016110a0565b602083015250604061112584828501610f31565b604083015250606082015167ffffffffffffffff81111561114957611148610fec565b5b61115584828501611073565b606083015250608061116984828501610f31565b60808301525092915050565b5f6020828403121561118a57611189610998565b5b5f82015167ffffffffffffffff8111156111a7576111a661099c565b5b6111b3848285016110b4565b91505092915050565b5f81519050919050565b5f6111d0826111bc565b6111da8185610dfd565b93506111ea818560208601611024565b6111f381610e1b565b840191505092915050565b5f6040820190508181035f830152611217818587610e2b565b9050818103602083015261122b81846111c6565b9050949350505050565b5f6060820190508181035f83015261124e81888a610e2b565b90508181036020830152611263818688610e2b565b90508181036040830152611278818486610e2b565b905097965050505050505056fea26469706673582212202b10b680125697620af6239789df4f52cb6b7af1310b2697119488024bc1a88d64736f6c634300081c0033",
}

// VotechainABI is the input ABI used to generate the binding from.
// Deprecated: Use VotechainMetaData.ABI instead.
var VotechainABI = VotechainMetaData.ABI

// VotechainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VotechainMetaData.Bin instead.
var VotechainBin = VotechainMetaData.Bin

// DeployVotechain deploys a new Ethereum contract, binding an instance of Votechain to it.
func DeployVotechain(auth *bind.TransactOpts, backend bind.ContractBackend, _baseAddress common.Address, _kpuManagerAddress common.Address, _voterManagerAddress common.Address, _electionManagerAddress common.Address) (common.Address, *types.Transaction, *Votechain, error) {
	parsed, err := VotechainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VotechainBin), backend, _baseAddress, _kpuManagerAddress, _voterManagerAddress, _electionManagerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Votechain{VotechainCaller: VotechainCaller{contract: contract}, VotechainTransactor: VotechainTransactor{contract: contract}, VotechainFilterer: VotechainFilterer{contract: contract}}, nil
}

// Votechain is an auto generated Go binding around an Ethereum contract.
type Votechain struct {
	VotechainCaller     // Read-only binding to the contract
	VotechainTransactor // Write-only binding to the contract
	VotechainFilterer   // Log filterer for contract events
}

// VotechainCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotechainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotechainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotechainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotechainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotechainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotechainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotechainSession struct {
	Contract     *Votechain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotechainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotechainCallerSession struct {
	Contract *VotechainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VotechainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotechainTransactorSession struct {
	Contract     *VotechainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VotechainRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotechainRaw struct {
	Contract *Votechain // Generic contract binding to access the raw methods on
}

// VotechainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotechainCallerRaw struct {
	Contract *VotechainCaller // Generic read-only contract binding to access the raw methods on
}

// VotechainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotechainTransactorRaw struct {
	Contract *VotechainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotechain creates a new instance of Votechain, bound to a specific deployed contract.
func NewVotechain(address common.Address, backend bind.ContractBackend) (*Votechain, error) {
	contract, err := bindVotechain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Votechain{VotechainCaller: VotechainCaller{contract: contract}, VotechainTransactor: VotechainTransactor{contract: contract}, VotechainFilterer: VotechainFilterer{contract: contract}}, nil
}

// NewVotechainCaller creates a new read-only instance of Votechain, bound to a specific deployed contract.
func NewVotechainCaller(address common.Address, caller bind.ContractCaller) (*VotechainCaller, error) {
	contract, err := bindVotechain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotechainCaller{contract: contract}, nil
}

// NewVotechainTransactor creates a new write-only instance of Votechain, bound to a specific deployed contract.
func NewVotechainTransactor(address common.Address, transactor bind.ContractTransactor) (*VotechainTransactor, error) {
	contract, err := bindVotechain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotechainTransactor{contract: contract}, nil
}

// NewVotechainFilterer creates a new log filterer instance of Votechain, bound to a specific deployed contract.
func NewVotechainFilterer(address common.Address, filterer bind.ContractFilterer) (*VotechainFilterer, error) {
	contract, err := bindVotechain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotechainFilterer{contract: contract}, nil
}

// bindVotechain binds a generic wrapper to an already deployed contract.
func bindVotechain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotechainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Votechain *VotechainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Votechain.Contract.VotechainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Votechain *VotechainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Votechain.Contract.VotechainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Votechain *VotechainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Votechain.Contract.VotechainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Votechain *VotechainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Votechain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Votechain *VotechainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Votechain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Votechain *VotechainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Votechain.Contract.contract.Transact(opts, method, params...)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_Votechain *VotechainCaller) Base(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "base")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_Votechain *VotechainSession) Base() (common.Address, error) {
	return _Votechain.Contract.Base(&_Votechain.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_Votechain *VotechainCallerSession) Base() (common.Address, error) {
	return _Votechain.Contract.Base(&_Votechain.CallOpts)
}

// ElectionManager is a free data retrieval call binding the contract method 0x7da32949.
//
// Solidity: function electionManager() view returns(address)
func (_Votechain *VotechainCaller) ElectionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "electionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ElectionManager is a free data retrieval call binding the contract method 0x7da32949.
//
// Solidity: function electionManager() view returns(address)
func (_Votechain *VotechainSession) ElectionManager() (common.Address, error) {
	return _Votechain.Contract.ElectionManager(&_Votechain.CallOpts)
}

// ElectionManager is a free data retrieval call binding the contract method 0x7da32949.
//
// Solidity: function electionManager() view returns(address)
func (_Votechain *VotechainCallerSession) ElectionManager() (common.Address, error) {
	return _Votechain.Contract.ElectionManager(&_Votechain.CallOpts)
}

// KpuManager is a free data retrieval call binding the contract method 0x34d1d2f2.
//
// Solidity: function kpuManager() view returns(address)
func (_Votechain *VotechainCaller) KpuManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "kpuManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KpuManager is a free data retrieval call binding the contract method 0x34d1d2f2.
//
// Solidity: function kpuManager() view returns(address)
func (_Votechain *VotechainSession) KpuManager() (common.Address, error) {
	return _Votechain.Contract.KpuManager(&_Votechain.CallOpts)
}

// KpuManager is a free data retrieval call binding the contract method 0x34d1d2f2.
//
// Solidity: function kpuManager() view returns(address)
func (_Votechain *VotechainCallerSession) KpuManager() (common.Address, error) {
	return _Votechain.Contract.KpuManager(&_Votechain.CallOpts)
}

// VoterManager is a free data retrieval call binding the contract method 0x65910a4f.
//
// Solidity: function voterManager() view returns(address)
func (_Votechain *VotechainCaller) VoterManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "voterManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoterManager is a free data retrieval call binding the contract method 0x65910a4f.
//
// Solidity: function voterManager() view returns(address)
func (_Votechain *VotechainSession) VoterManager() (common.Address, error) {
	return _Votechain.Contract.VoterManager(&_Votechain.CallOpts)
}

// VoterManager is a free data retrieval call binding the contract method 0x65910a4f.
//
// Solidity: function voterManager() view returns(address)
func (_Votechain *VotechainCallerSession) VoterManager() (common.Address, error) {
	return _Votechain.Contract.VoterManager(&_Votechain.CallOpts)
}

// AddElection is a paid mutator transaction binding the contract method 0xfd1c6e61.
//
// Solidity: function addElection(string electionId, string electionName, string electionNo) returns()
func (_Votechain *VotechainTransactor) AddElection(opts *bind.TransactOpts, electionId string, electionName string, electionNo string) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "addElection", electionId, electionName, electionNo)
}

// AddElection is a paid mutator transaction binding the contract method 0xfd1c6e61.
//
// Solidity: function addElection(string electionId, string electionName, string electionNo) returns()
func (_Votechain *VotechainSession) AddElection(electionId string, electionName string, electionNo string) (*types.Transaction, error) {
	return _Votechain.Contract.AddElection(&_Votechain.TransactOpts, electionId, electionName, electionNo)
}

// AddElection is a paid mutator transaction binding the contract method 0xfd1c6e61.
//
// Solidity: function addElection(string electionId, string electionName, string electionNo) returns()
func (_Votechain *VotechainTransactorSession) AddElection(electionId string, electionName string, electionNo string) (*types.Transaction, error) {
	return _Votechain.Contract.AddElection(&_Votechain.TransactOpts, electionId, electionName, electionNo)
}

// RegisterKPUKota is a paid mutator transaction binding the contract method 0x92423815.
//
// Solidity: function registerKPUKota(address Address, string name, string region) returns()
func (_Votechain *VotechainTransactor) RegisterKPUKota(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "registerKPUKota", Address, name, region)
}

// RegisterKPUKota is a paid mutator transaction binding the contract method 0x92423815.
//
// Solidity: function registerKPUKota(address Address, string name, string region) returns()
func (_Votechain *VotechainSession) RegisterKPUKota(Address common.Address, name string, region string) (*types.Transaction, error) {
	return _Votechain.Contract.RegisterKPUKota(&_Votechain.TransactOpts, Address, name, region)
}

// RegisterKPUKota is a paid mutator transaction binding the contract method 0x92423815.
//
// Solidity: function registerKPUKota(address Address, string name, string region) returns()
func (_Votechain *VotechainTransactorSession) RegisterKPUKota(Address common.Address, name string, region string) (*types.Transaction, error) {
	return _Votechain.Contract.RegisterKPUKota(&_Votechain.TransactOpts, Address, name, region)
}

// RegisterKPUProvinsi is a paid mutator transaction binding the contract method 0x2db33a2f.
//
// Solidity: function registerKPUProvinsi(address Address, string name, string region) returns()
func (_Votechain *VotechainTransactor) RegisterKPUProvinsi(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "registerKPUProvinsi", Address, name, region)
}

// RegisterKPUProvinsi is a paid mutator transaction binding the contract method 0x2db33a2f.
//
// Solidity: function registerKPUProvinsi(address Address, string name, string region) returns()
func (_Votechain *VotechainSession) RegisterKPUProvinsi(Address common.Address, name string, region string) (*types.Transaction, error) {
	return _Votechain.Contract.RegisterKPUProvinsi(&_Votechain.TransactOpts, Address, name, region)
}

// RegisterKPUProvinsi is a paid mutator transaction binding the contract method 0x2db33a2f.
//
// Solidity: function registerKPUProvinsi(address Address, string name, string region) returns()
func (_Votechain *VotechainTransactorSession) RegisterKPUProvinsi(Address common.Address, name string, region string) (*types.Transaction, error) {
	return _Votechain.Contract.RegisterKPUProvinsi(&_Votechain.TransactOpts, Address, name, region)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x4a075de2.
//
// Solidity: function registerVoter(string nik, address voterAddress) returns()
func (_Votechain *VotechainTransactor) RegisterVoter(opts *bind.TransactOpts, nik string, voterAddress common.Address) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "registerVoter", nik, voterAddress)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x4a075de2.
//
// Solidity: function registerVoter(string nik, address voterAddress) returns()
func (_Votechain *VotechainSession) RegisterVoter(nik string, voterAddress common.Address) (*types.Transaction, error) {
	return _Votechain.Contract.RegisterVoter(&_Votechain.TransactOpts, nik, voterAddress)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x4a075de2.
//
// Solidity: function registerVoter(string nik, address voterAddress) returns()
func (_Votechain *VotechainTransactorSession) RegisterVoter(nik string, voterAddress common.Address) (*types.Transaction, error) {
	return _Votechain.Contract.RegisterVoter(&_Votechain.TransactOpts, nik, voterAddress)
}

// SetKpuAdmin is a paid mutator transaction binding the contract method 0x9df86dc1.
//
// Solidity: function setKpuAdmin(address newAdmin) returns()
func (_Votechain *VotechainTransactor) SetKpuAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "setKpuAdmin", newAdmin)
}

// SetKpuAdmin is a paid mutator transaction binding the contract method 0x9df86dc1.
//
// Solidity: function setKpuAdmin(address newAdmin) returns()
func (_Votechain *VotechainSession) SetKpuAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Votechain.Contract.SetKpuAdmin(&_Votechain.TransactOpts, newAdmin)
}

// SetKpuAdmin is a paid mutator transaction binding the contract method 0x9df86dc1.
//
// Solidity: function setKpuAdmin(address newAdmin) returns()
func (_Votechain *VotechainTransactorSession) SetKpuAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Votechain.Contract.SetKpuAdmin(&_Votechain.TransactOpts, newAdmin)
}

// SetVotingStatus is a paid mutator transaction binding the contract method 0x7478c9fe.
//
// Solidity: function setVotingStatus(bool status) returns()
func (_Votechain *VotechainTransactor) SetVotingStatus(opts *bind.TransactOpts, status bool) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "setVotingStatus", status)
}

// SetVotingStatus is a paid mutator transaction binding the contract method 0x7478c9fe.
//
// Solidity: function setVotingStatus(bool status) returns()
func (_Votechain *VotechainSession) SetVotingStatus(status bool) (*types.Transaction, error) {
	return _Votechain.Contract.SetVotingStatus(&_Votechain.TransactOpts, status)
}

// SetVotingStatus is a paid mutator transaction binding the contract method 0x7478c9fe.
//
// Solidity: function setVotingStatus(bool status) returns()
func (_Votechain *VotechainTransactorSession) SetVotingStatus(status bool) (*types.Transaction, error) {
	return _Votechain.Contract.SetVotingStatus(&_Votechain.TransactOpts, status)
}

// ToggleElectionActive is a paid mutator transaction binding the contract method 0xdb5dfdd7.
//
// Solidity: function toggleElectionActive(string electionId) returns()
func (_Votechain *VotechainTransactor) ToggleElectionActive(opts *bind.TransactOpts, electionId string) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "toggleElectionActive", electionId)
}

// ToggleElectionActive is a paid mutator transaction binding the contract method 0xdb5dfdd7.
//
// Solidity: function toggleElectionActive(string electionId) returns()
func (_Votechain *VotechainSession) ToggleElectionActive(electionId string) (*types.Transaction, error) {
	return _Votechain.Contract.ToggleElectionActive(&_Votechain.TransactOpts, electionId)
}

// ToggleElectionActive is a paid mutator transaction binding the contract method 0xdb5dfdd7.
//
// Solidity: function toggleElectionActive(string electionId) returns()
func (_Votechain *VotechainTransactorSession) ToggleElectionActive(electionId string) (*types.Transaction, error) {
	return _Votechain.Contract.ToggleElectionActive(&_Votechain.TransactOpts, electionId)
}

// Vote is a paid mutator transaction binding the contract method 0xfc36e15b.
//
// Solidity: function vote(string electionId) returns()
func (_Votechain *VotechainTransactor) Vote(opts *bind.TransactOpts, electionId string) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "vote", electionId)
}

// Vote is a paid mutator transaction binding the contract method 0xfc36e15b.
//
// Solidity: function vote(string electionId) returns()
func (_Votechain *VotechainSession) Vote(electionId string) (*types.Transaction, error) {
	return _Votechain.Contract.Vote(&_Votechain.TransactOpts, electionId)
}

// Vote is a paid mutator transaction binding the contract method 0xfc36e15b.
//
// Solidity: function vote(string electionId) returns()
func (_Votechain *VotechainTransactorSession) Vote(electionId string) (*types.Transaction, error) {
	return _Votechain.Contract.Vote(&_Votechain.TransactOpts, electionId)
}
