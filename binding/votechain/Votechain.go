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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_kpuManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voterManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_electionManagerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyVoted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidElection\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VoterNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotingNotActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"}],\"name\":\"addElection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"contractIVotechainBase\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"electionManager\",\"outputs\":[{\"internalType\":\"contractIElectionManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kpuManager\",\"outputs\":[{\"internalType\":\"contractIKPUManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"registerKPUKota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"registerKPUProvinsi\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"}],\"name\":\"registerVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setKpuAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setVotingStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"}],\"name\":\"toggleElectionActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"updateKPUKota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"updateKPUProvinsi\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterManager\",\"outputs\":[{\"internalType\":\"contractIVoterManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161165d38038061165d83398181016040528101906100319190610197565b835f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050506101fb565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101668261013d565b9050919050565b6101768161015c565b8114610180575f5ffd5b50565b5f815190506101918161016d565b92915050565b5f5f5f5f608085870312156101af576101ae610139565b5b5f6101bc87828801610183565b94505060206101cd87828801610183565b93505060406101de87828801610183565b92505060606101ef87828801610183565b91505092959194509250565b611455806102085f395ff3fe608060405234801561000f575f5ffd5b50600436106100e8575f3560e01c80637478c9fe1161008a5780639df86dc1116100645780639df86dc11461020c578063db5dfdd714610228578063fc36e15b14610244578063fd1c6e6114610260576100e8565b80637478c9fe146101b65780637da32949146101d257806392423815146101f0576100e8565b80634a075de2116100c65780634a075de2146101425780634da1b97a1461015e5780635001f3b51461017a57806365910a4f14610198576100e8565b806302819b99146100ec5780632db33a2f1461010857806334d1d2f214610124575b5f5ffd5b61010660048036038101906101019190610bf5565b61027c565b005b610122600480360381019061011d9190610bf5565b610312565b005b61012c6103a8565b6040516101399190610ce1565b60405180910390f35b61015c60048036038101906101579190610cfa565b6103cd565b005b61017860048036038101906101739190610bf5565b61045d565b005b6101826104f3565b60405161018f9190610d77565b60405180910390f35b6101a0610517565b6040516101ad9190610db0565b60405180910390f35b6101d060048036038101906101cb9190610dfe565b61053c565b005b6101da6105c5565b6040516101e79190610e49565b60405180910390f35b61020a60048036038101906102059190610bf5565b6105ea565b005b61022660048036038101906102219190610e62565b610680565b005b610242600480360381019061023d9190610e8d565b610709565b005b61025e60048036038101906102599190610e8d565b610796565b005b61027a60048036038101906102759190610ed8565b610a90565b005b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166302819b9986868686866040518663ffffffff1660e01b81526004016102de959493929190610ff1565b5f604051808303815f87803b1580156102f5575f5ffd5b505af1158015610307573d5f5f3e3d5ffd5b505050505050505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632db33a2f86868686866040518663ffffffff1660e01b8152600401610374959493929190610ff1565b5f604051808303815f87803b15801561038b575f5ffd5b505af115801561039d573d5f5f3e3d5ffd5b505050505050505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634a075de28484846040518463ffffffff1660e01b815260040161042b93929190611038565b5f604051808303815f87803b158015610442575f5ffd5b505af1158015610454573d5f5f3e3d5ffd5b50505050505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634da1b97a86868686866040518663ffffffff1660e01b81526004016104bf959493929190610ff1565b5f604051808303815f87803b1580156104d6575f5ffd5b505af11580156104e8573d5f5f3e3d5ffd5b505050505050505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637478c9fe826040518263ffffffff1660e01b81526004016105959190611077565b5f604051808303815f87803b1580156105ac575f5ffd5b505af11580156105be573d5f5f3e3d5ffd5b5050505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639242381586868686866040518663ffffffff1660e01b815260040161064c959493929190610ff1565b5f604051808303815f87803b158015610663575f5ffd5b505af1158015610675573d5f5f3e3d5ffd5b505050505050505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639df86dc1826040518263ffffffff1660e01b81526004016106d99190611090565b5f604051808303815f87803b1580156106f0575f5ffd5b505af1158015610702573d5f5f3e3d5ffd5b5050505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663db5dfdd783836040518363ffffffff1660e01b81526004016107659291906110a9565b5f604051808303815f87803b15801561077c575f5ffd5b505af115801561078e573d5f5f3e3d5ffd5b505050505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663408e27276040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107ff573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061082391906110df565b610859576040517f9b8cc47500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634bdd7585336040518263ffffffff1660e01b81526004016108b39190611090565b5f60405180830381865afa9250505080156108f057506040513d5f823e3d601f19601f820116820180604052508101906108ed919061130f565b60015b610926576040517f6f08c58700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806040015115610962576040517f7c9a1cf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a8fd0cfa336040518263ffffffff1660e01b81526004016109bc9190611090565b6020604051808303815f875af11580156109d8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109fc91906110df565b5060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e8d5940d8484845f01516040518463ffffffff1660e01b8152600401610a5e93929190611398565b5f604051808303815f87803b158015610a75575f5ffd5b505af1158015610a87573d5f5f3e3d5ffd5b50505050505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fd1c6e618787878787876040518763ffffffff1660e01b8152600401610af4969594939291906113cf565b5f604051808303815f87803b158015610b0b575f5ffd5b505af1158015610b1d573d5f5f3e3d5ffd5b50505050505050505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b6382610b3a565b9050919050565b610b7381610b59565b8114610b7d575f5ffd5b50565b5f81359050610b8e81610b6a565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112610bb557610bb4610b94565b5b8235905067ffffffffffffffff811115610bd257610bd1610b98565b5b602083019150836001820283011115610bee57610bed610b9c565b5b9250929050565b5f5f5f5f5f60608688031215610c0e57610c0d610b32565b5b5f610c1b88828901610b80565b955050602086013567ffffffffffffffff811115610c3c57610c3b610b36565b5b610c4888828901610ba0565b9450945050604086013567ffffffffffffffff811115610c6b57610c6a610b36565b5b610c7788828901610ba0565b92509250509295509295909350565b5f819050919050565b5f610ca9610ca4610c9f84610b3a565b610c86565b610b3a565b9050919050565b5f610cba82610c8f565b9050919050565b5f610ccb82610cb0565b9050919050565b610cdb81610cc1565b82525050565b5f602082019050610cf45f830184610cd2565b92915050565b5f5f5f60408486031215610d1157610d10610b32565b5b5f84013567ffffffffffffffff811115610d2e57610d2d610b36565b5b610d3a86828701610ba0565b93509350506020610d4d86828701610b80565b9150509250925092565b5f610d6182610cb0565b9050919050565b610d7181610d57565b82525050565b5f602082019050610d8a5f830184610d68565b92915050565b5f610d9a82610cb0565b9050919050565b610daa81610d90565b82525050565b5f602082019050610dc35f830184610da1565b92915050565b5f8115159050919050565b610ddd81610dc9565b8114610de7575f5ffd5b50565b5f81359050610df881610dd4565b92915050565b5f60208284031215610e1357610e12610b32565b5b5f610e2084828501610dea565b91505092915050565b5f610e3382610cb0565b9050919050565b610e4381610e29565b82525050565b5f602082019050610e5c5f830184610e3a565b92915050565b5f60208284031215610e7757610e76610b32565b5b5f610e8484828501610b80565b91505092915050565b5f5f60208385031215610ea357610ea2610b32565b5b5f83013567ffffffffffffffff811115610ec057610ebf610b36565b5b610ecc85828601610ba0565b92509250509250929050565b5f5f5f5f5f5f60608789031215610ef257610ef1610b32565b5b5f87013567ffffffffffffffff811115610f0f57610f0e610b36565b5b610f1b89828a01610ba0565b9650965050602087013567ffffffffffffffff811115610f3e57610f3d610b36565b5b610f4a89828a01610ba0565b9450945050604087013567ffffffffffffffff811115610f6d57610f6c610b36565b5b610f7989828a01610ba0565b92509250509295509295509295565b610f9181610b59565b82525050565b5f82825260208201905092915050565b828183375f83830152505050565b5f601f19601f8301169050919050565b5f610fd08385610f97565b9350610fdd838584610fa7565b610fe683610fb5565b840190509392505050565b5f6060820190506110045f830188610f88565b8181036020830152611017818688610fc5565b9050818103604083015261102c818486610fc5565b90509695505050505050565b5f6040820190508181035f830152611051818587610fc5565b90506110606020830184610f88565b949350505050565b61107181610dc9565b82525050565b5f60208201905061108a5f830184611068565b92915050565b5f6020820190506110a35f830184610f88565b92915050565b5f6020820190508181035f8301526110c2818486610fc5565b90509392505050565b5f815190506110d981610dd4565b92915050565b5f602082840312156110f4576110f3610b32565b5b5f611101848285016110cb565b91505092915050565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61114482610fb5565b810181811067ffffffffffffffff821117156111635761116261110e565b5b80604052505050565b5f611175610b29565b9050611181828261113b565b919050565b5f5ffd5b5f5ffd5b5f67ffffffffffffffff8211156111a8576111a761110e565b5b6111b182610fb5565b9050602081019050919050565b8281835e5f83830152505050565b5f6111de6111d98461118e565b61116c565b9050828152602081018484840111156111fa576111f961118a565b5b6112058482856111be565b509392505050565b5f82601f83011261122157611220610b94565b5b81516112318482602086016111cc565b91505092915050565b5f8151905061124881610b6a565b92915050565b5f60a082840312156112635761126261110a565b5b61126d60a061116c565b90505f82015167ffffffffffffffff81111561128c5761128b611186565b5b6112988482850161120d565b5f8301525060206112ab8482850161123a565b60208301525060406112bf848285016110cb565b604083015250606082015167ffffffffffffffff8111156112e3576112e2611186565b5b6112ef8482850161120d565b6060830152506080611303848285016110cb565b60808301525092915050565b5f6020828403121561132457611323610b32565b5b5f82015167ffffffffffffffff81111561134157611340610b36565b5b61134d8482850161124e565b91505092915050565b5f81519050919050565b5f61136a82611356565b6113748185610f97565b93506113848185602086016111be565b61138d81610fb5565b840191505092915050565b5f6040820190508181035f8301526113b1818587610fc5565b905081810360208301526113c58184611360565b9050949350505050565b5f6060820190508181035f8301526113e881888a610fc5565b905081810360208301526113fd818688610fc5565b90508181036040830152611412818486610fc5565b905097965050505050505056fea26469706673582212201f59cf3c4dbf581c05f14b0344f97308686addb0c94e57a94a4e80868992299964736f6c634300081c0033",
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

// UpdateKPUKota is a paid mutator transaction binding the contract method 0x4da1b97a.
//
// Solidity: function updateKPUKota(address Address, string name, string region) returns()
func (_Votechain *VotechainTransactor) UpdateKPUKota(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "updateKPUKota", Address, name, region)
}

// UpdateKPUKota is a paid mutator transaction binding the contract method 0x4da1b97a.
//
// Solidity: function updateKPUKota(address Address, string name, string region) returns()
func (_Votechain *VotechainSession) UpdateKPUKota(Address common.Address, name string, region string) (*types.Transaction, error) {
	return _Votechain.Contract.UpdateKPUKota(&_Votechain.TransactOpts, Address, name, region)
}

// UpdateKPUKota is a paid mutator transaction binding the contract method 0x4da1b97a.
//
// Solidity: function updateKPUKota(address Address, string name, string region) returns()
func (_Votechain *VotechainTransactorSession) UpdateKPUKota(Address common.Address, name string, region string) (*types.Transaction, error) {
	return _Votechain.Contract.UpdateKPUKota(&_Votechain.TransactOpts, Address, name, region)
}

// UpdateKPUProvinsi is a paid mutator transaction binding the contract method 0x02819b99.
//
// Solidity: function updateKPUProvinsi(address Address, string name, string region) returns()
func (_Votechain *VotechainTransactor) UpdateKPUProvinsi(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "updateKPUProvinsi", Address, name, region)
}

// UpdateKPUProvinsi is a paid mutator transaction binding the contract method 0x02819b99.
//
// Solidity: function updateKPUProvinsi(address Address, string name, string region) returns()
func (_Votechain *VotechainSession) UpdateKPUProvinsi(Address common.Address, name string, region string) (*types.Transaction, error) {
	return _Votechain.Contract.UpdateKPUProvinsi(&_Votechain.TransactOpts, Address, name, region)
}

// UpdateKPUProvinsi is a paid mutator transaction binding the contract method 0x02819b99.
//
// Solidity: function updateKPUProvinsi(address Address, string name, string region) returns()
func (_Votechain *VotechainTransactorSession) UpdateKPUProvinsi(Address common.Address, name string, region string) (*types.Transaction, error) {
	return _Votechain.Contract.UpdateKPUProvinsi(&_Votechain.TransactOpts, Address, name, region)
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
