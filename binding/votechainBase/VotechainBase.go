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
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"trustedForwarders\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"OnlyKpuAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotingNotActive\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"TrustedForwarderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"TrustedForwarderRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"trusted\",\"type\":\"bool\"}],\"name\":\"TrustedForwarderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"VotingStatusChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"addTrustedForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTrustedForwarders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kpuAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"removeTrustedForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setKpuAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setVotingStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161116d38038061116d83398181016040528101906100319190610477565b805f5f90505b815181101561007657610069828281518110610056576100556104be565b5b60200260200101516100cb60201b60201c565b8080600101915050610037565b505061008661021360201b60201c565b60025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506104eb565b5f5f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff166102105760015f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550600181908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f3ef8564460ada92419608d823c014975d98f8104d7d1e68c222967ac6814cd9760405160405180910390a25b50565b5f61022261022760201b60201c565b905090565b5f6102373361026460201b60201c565b1561024b57601436033560601c9050610260565b6102596102b560201b60201c565b9050610261565b5b90565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f33905090565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610317826102d1565b810181811067ffffffffffffffff82111715610336576103356102e1565b5b80604052505050565b5f6103486102bc565b9050610354828261030e565b919050565b5f67ffffffffffffffff821115610373576103726102e1565b5b602082029050602081019050919050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6103b182610388565b9050919050565b6103c1816103a7565b81146103cb575f5ffd5b50565b5f815190506103dc816103b8565b92915050565b5f6103f46103ef84610359565b61033f565b9050808382526020820190506020840283018581111561041757610416610384565b5b835b81811015610440578061042c88826103ce565b845260208401935050602081019050610419565b5050509392505050565b5f82601f83011261045e5761045d6102cd565b5b815161046e8482602086016103e2565b91505092915050565b5f6020828403121561048c5761048b6102c5565b5b5f82015167ffffffffffffffff8111156104a9576104a86102c9565b5b6104b58482850161044a565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b610c75806104f85f395ff3fe608060405234801561000f575f5ffd5b5060043610610086575f3560e01c80637478c9fe116100595780637478c9fe146101105780639df86dc11461012c578063fb4ab16414610148578063ff7196c81461016657610086565b806326d006681461008a5780633751d89c146100a6578063408e2727146100c2578063572b6c05146100e0575b5f5ffd5b6100a4600480360381019061009f91906109ca565b610184565b005b6100c060048036038101906100bb91906109ca565b61021d565b005b6100ca6102b6565b6040516100d79190610a0f565b60405180910390f35b6100fa60048036038101906100f591906109ca565b6102c9565b6040516101079190610a0f565b60405180910390f35b61012a60048036038101906101259190610a52565b61031a565b005b610146600480360381019061014191906109ca565b6103fb565b005b6101506104cb565b60405161015d9190610a8c565b60405180910390f35b61016e6104f0565b60405161017b9190610b5c565b60405180910390f35b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166101c461057b565b73ffffffffffffffffffffffffffffffffffffffff1614610211576040517f8b8958dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61021a81610589565b50565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661025d61057b565b73ffffffffffffffffffffffffffffffffffffffff16146102aa576040517f8b8958dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102b3816106d1565b50565b600260149054906101000a900460ff1681565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661035a61057b565b73ffffffffffffffffffffffffffffffffffffffff16146103a7576040517f8b8958dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600260146101000a81548160ff0219169083151502179055507f9069a1a16ace751e8690f383e12f87b01e8488ba387e626810bd113fef0417f9816040516103f09190610a0f565b60405180910390a150565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661043b61057b565b73ffffffffffffffffffffffffffffffffffffffff1614610488576040517f8b8958dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060600180548060200260200160405190810160405280929190818152602001828054801561057157602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610528575b5050505050905090565b5f610584610934565b905090565b5f5f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff166106ce5760015f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550600181908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f3ef8564460ada92419608d823c014975d98f8104d7d1e68c222967ac6814cd9760405160405180910390a25b50565b5f5f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610931575f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f5f90505b6001805490508110156108ec578173ffffffffffffffffffffffffffffffffffffffff16600182815481106107af576107ae610b7c565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036108df5760018080805490506108049190610bdf565b8154811061081557610814610b7c565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166001828154811061085157610850610b7c565b5b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018054806108a8576108a7610c12565b5b600190038181905f5260205f20015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590556108ec565b8080600101915050610777565b508073ffffffffffffffffffffffffffffffffffffffff167fd2d636efcad0cea42e170256f4c5d8b1cd81e47b855557edaf44014e6cc4cee660405160405180910390a25b50565b5f61093e336102c9565b1561095257601436033560601c9050610961565b61095a610965565b9050610962565b5b90565b5f33905090565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61099982610970565b9050919050565b6109a98161098f565b81146109b3575f5ffd5b50565b5f813590506109c4816109a0565b92915050565b5f602082840312156109df576109de61096c565b5b5f6109ec848285016109b6565b91505092915050565b5f8115159050919050565b610a09816109f5565b82525050565b5f602082019050610a225f830184610a00565b92915050565b610a31816109f5565b8114610a3b575f5ffd5b50565b5f81359050610a4c81610a28565b92915050565b5f60208284031215610a6757610a6661096c565b5b5f610a7484828501610a3e565b91505092915050565b610a868161098f565b82525050565b5f602082019050610a9f5f830184610a7d565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610ad78161098f565b82525050565b5f610ae88383610ace565b60208301905092915050565b5f602082019050919050565b5f610b0a82610aa5565b610b148185610aaf565b9350610b1f83610abf565b805f5b83811015610b4f578151610b368882610add565b9750610b4183610af4565b925050600181019050610b22565b5085935050505092915050565b5f6020820190508181035f830152610b748184610b00565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f819050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610be982610ba9565b9150610bf483610ba9565b9250828203905081811115610c0c57610c0b610bb2565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea26469706673582212204dcf427298a26ca1612f466d354faee3a3efc4837d22f1207da393181358572164736f6c634300081c0033",
}

// VotechainBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use VotechainBaseMetaData.ABI instead.
var VotechainBaseABI = VotechainBaseMetaData.ABI

// VotechainBaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VotechainBaseMetaData.Bin instead.
var VotechainBaseBin = VotechainBaseMetaData.Bin

// DeployVotechainBase deploys a new Ethereum contract, binding an instance of VotechainBase to it.
func DeployVotechainBase(auth *bind.TransactOpts, backend bind.ContractBackend, trustedForwarders []common.Address) (common.Address, *types.Transaction, *VotechainBase, error) {
	parsed, err := VotechainBaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VotechainBaseBin), backend, trustedForwarders)
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

// GetTrustedForwarders is a free data retrieval call binding the contract method 0xff7196c8.
//
// Solidity: function getTrustedForwarders() view returns(address[])
func (_VotechainBase *VotechainBaseCaller) GetTrustedForwarders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _VotechainBase.contract.Call(opts, &out, "getTrustedForwarders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTrustedForwarders is a free data retrieval call binding the contract method 0xff7196c8.
//
// Solidity: function getTrustedForwarders() view returns(address[])
func (_VotechainBase *VotechainBaseSession) GetTrustedForwarders() ([]common.Address, error) {
	return _VotechainBase.Contract.GetTrustedForwarders(&_VotechainBase.CallOpts)
}

// GetTrustedForwarders is a free data retrieval call binding the contract method 0xff7196c8.
//
// Solidity: function getTrustedForwarders() view returns(address[])
func (_VotechainBase *VotechainBaseCallerSession) GetTrustedForwarders() ([]common.Address, error) {
	return _VotechainBase.Contract.GetTrustedForwarders(&_VotechainBase.CallOpts)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_VotechainBase *VotechainBaseCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _VotechainBase.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_VotechainBase *VotechainBaseSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _VotechainBase.Contract.IsTrustedForwarder(&_VotechainBase.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_VotechainBase *VotechainBaseCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _VotechainBase.Contract.IsTrustedForwarder(&_VotechainBase.CallOpts, forwarder)
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

// AddTrustedForwarder is a paid mutator transaction binding the contract method 0x26d00668.
//
// Solidity: function addTrustedForwarder(address forwarder) returns()
func (_VotechainBase *VotechainBaseTransactor) AddTrustedForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _VotechainBase.contract.Transact(opts, "addTrustedForwarder", forwarder)
}

// AddTrustedForwarder is a paid mutator transaction binding the contract method 0x26d00668.
//
// Solidity: function addTrustedForwarder(address forwarder) returns()
func (_VotechainBase *VotechainBaseSession) AddTrustedForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _VotechainBase.Contract.AddTrustedForwarder(&_VotechainBase.TransactOpts, forwarder)
}

// AddTrustedForwarder is a paid mutator transaction binding the contract method 0x26d00668.
//
// Solidity: function addTrustedForwarder(address forwarder) returns()
func (_VotechainBase *VotechainBaseTransactorSession) AddTrustedForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _VotechainBase.Contract.AddTrustedForwarder(&_VotechainBase.TransactOpts, forwarder)
}

// RemoveTrustedForwarder is a paid mutator transaction binding the contract method 0x3751d89c.
//
// Solidity: function removeTrustedForwarder(address forwarder) returns()
func (_VotechainBase *VotechainBaseTransactor) RemoveTrustedForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _VotechainBase.contract.Transact(opts, "removeTrustedForwarder", forwarder)
}

// RemoveTrustedForwarder is a paid mutator transaction binding the contract method 0x3751d89c.
//
// Solidity: function removeTrustedForwarder(address forwarder) returns()
func (_VotechainBase *VotechainBaseSession) RemoveTrustedForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _VotechainBase.Contract.RemoveTrustedForwarder(&_VotechainBase.TransactOpts, forwarder)
}

// RemoveTrustedForwarder is a paid mutator transaction binding the contract method 0x3751d89c.
//
// Solidity: function removeTrustedForwarder(address forwarder) returns()
func (_VotechainBase *VotechainBaseTransactorSession) RemoveTrustedForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _VotechainBase.Contract.RemoveTrustedForwarder(&_VotechainBase.TransactOpts, forwarder)
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

// VotechainBaseTrustedForwarderAddedIterator is returned from FilterTrustedForwarderAdded and is used to iterate over the raw logs and unpacked data for TrustedForwarderAdded events raised by the VotechainBase contract.
type VotechainBaseTrustedForwarderAddedIterator struct {
	Event *VotechainBaseTrustedForwarderAdded // Event containing the contract specifics and raw log

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
func (it *VotechainBaseTrustedForwarderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotechainBaseTrustedForwarderAdded)
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
		it.Event = new(VotechainBaseTrustedForwarderAdded)
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
func (it *VotechainBaseTrustedForwarderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotechainBaseTrustedForwarderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotechainBaseTrustedForwarderAdded represents a TrustedForwarderAdded event raised by the VotechainBase contract.
type VotechainBaseTrustedForwarderAdded struct {
	Forwarder common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTrustedForwarderAdded is a free log retrieval operation binding the contract event 0x3ef8564460ada92419608d823c014975d98f8104d7d1e68c222967ac6814cd97.
//
// Solidity: event TrustedForwarderAdded(address indexed forwarder)
func (_VotechainBase *VotechainBaseFilterer) FilterTrustedForwarderAdded(opts *bind.FilterOpts, forwarder []common.Address) (*VotechainBaseTrustedForwarderAddedIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _VotechainBase.contract.FilterLogs(opts, "TrustedForwarderAdded", forwarderRule)
	if err != nil {
		return nil, err
	}
	return &VotechainBaseTrustedForwarderAddedIterator{contract: _VotechainBase.contract, event: "TrustedForwarderAdded", logs: logs, sub: sub}, nil
}

// WatchTrustedForwarderAdded is a free log subscription operation binding the contract event 0x3ef8564460ada92419608d823c014975d98f8104d7d1e68c222967ac6814cd97.
//
// Solidity: event TrustedForwarderAdded(address indexed forwarder)
func (_VotechainBase *VotechainBaseFilterer) WatchTrustedForwarderAdded(opts *bind.WatchOpts, sink chan<- *VotechainBaseTrustedForwarderAdded, forwarder []common.Address) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _VotechainBase.contract.WatchLogs(opts, "TrustedForwarderAdded", forwarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotechainBaseTrustedForwarderAdded)
				if err := _VotechainBase.contract.UnpackLog(event, "TrustedForwarderAdded", log); err != nil {
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

// ParseTrustedForwarderAdded is a log parse operation binding the contract event 0x3ef8564460ada92419608d823c014975d98f8104d7d1e68c222967ac6814cd97.
//
// Solidity: event TrustedForwarderAdded(address indexed forwarder)
func (_VotechainBase *VotechainBaseFilterer) ParseTrustedForwarderAdded(log types.Log) (*VotechainBaseTrustedForwarderAdded, error) {
	event := new(VotechainBaseTrustedForwarderAdded)
	if err := _VotechainBase.contract.UnpackLog(event, "TrustedForwarderAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotechainBaseTrustedForwarderRemovedIterator is returned from FilterTrustedForwarderRemoved and is used to iterate over the raw logs and unpacked data for TrustedForwarderRemoved events raised by the VotechainBase contract.
type VotechainBaseTrustedForwarderRemovedIterator struct {
	Event *VotechainBaseTrustedForwarderRemoved // Event containing the contract specifics and raw log

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
func (it *VotechainBaseTrustedForwarderRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotechainBaseTrustedForwarderRemoved)
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
		it.Event = new(VotechainBaseTrustedForwarderRemoved)
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
func (it *VotechainBaseTrustedForwarderRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotechainBaseTrustedForwarderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotechainBaseTrustedForwarderRemoved represents a TrustedForwarderRemoved event raised by the VotechainBase contract.
type VotechainBaseTrustedForwarderRemoved struct {
	Forwarder common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTrustedForwarderRemoved is a free log retrieval operation binding the contract event 0xd2d636efcad0cea42e170256f4c5d8b1cd81e47b855557edaf44014e6cc4cee6.
//
// Solidity: event TrustedForwarderRemoved(address indexed forwarder)
func (_VotechainBase *VotechainBaseFilterer) FilterTrustedForwarderRemoved(opts *bind.FilterOpts, forwarder []common.Address) (*VotechainBaseTrustedForwarderRemovedIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _VotechainBase.contract.FilterLogs(opts, "TrustedForwarderRemoved", forwarderRule)
	if err != nil {
		return nil, err
	}
	return &VotechainBaseTrustedForwarderRemovedIterator{contract: _VotechainBase.contract, event: "TrustedForwarderRemoved", logs: logs, sub: sub}, nil
}

// WatchTrustedForwarderRemoved is a free log subscription operation binding the contract event 0xd2d636efcad0cea42e170256f4c5d8b1cd81e47b855557edaf44014e6cc4cee6.
//
// Solidity: event TrustedForwarderRemoved(address indexed forwarder)
func (_VotechainBase *VotechainBaseFilterer) WatchTrustedForwarderRemoved(opts *bind.WatchOpts, sink chan<- *VotechainBaseTrustedForwarderRemoved, forwarder []common.Address) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _VotechainBase.contract.WatchLogs(opts, "TrustedForwarderRemoved", forwarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotechainBaseTrustedForwarderRemoved)
				if err := _VotechainBase.contract.UnpackLog(event, "TrustedForwarderRemoved", log); err != nil {
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

// ParseTrustedForwarderRemoved is a log parse operation binding the contract event 0xd2d636efcad0cea42e170256f4c5d8b1cd81e47b855557edaf44014e6cc4cee6.
//
// Solidity: event TrustedForwarderRemoved(address indexed forwarder)
func (_VotechainBase *VotechainBaseFilterer) ParseTrustedForwarderRemoved(log types.Log) (*VotechainBaseTrustedForwarderRemoved, error) {
	event := new(VotechainBaseTrustedForwarderRemoved)
	if err := _VotechainBase.contract.UnpackLog(event, "TrustedForwarderRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotechainBaseTrustedForwarderUpdatedIterator is returned from FilterTrustedForwarderUpdated and is used to iterate over the raw logs and unpacked data for TrustedForwarderUpdated events raised by the VotechainBase contract.
type VotechainBaseTrustedForwarderUpdatedIterator struct {
	Event *VotechainBaseTrustedForwarderUpdated // Event containing the contract specifics and raw log

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
func (it *VotechainBaseTrustedForwarderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotechainBaseTrustedForwarderUpdated)
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
		it.Event = new(VotechainBaseTrustedForwarderUpdated)
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
func (it *VotechainBaseTrustedForwarderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotechainBaseTrustedForwarderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotechainBaseTrustedForwarderUpdated represents a TrustedForwarderUpdated event raised by the VotechainBase contract.
type VotechainBaseTrustedForwarderUpdated struct {
	Forwarder common.Address
	Trusted   bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTrustedForwarderUpdated is a free log retrieval operation binding the contract event 0xbee55516e29d3969d3cb8eb01351eb3c52d06f9e2435bd5a8bfe3647e185df92.
//
// Solidity: event TrustedForwarderUpdated(address indexed forwarder, bool trusted)
func (_VotechainBase *VotechainBaseFilterer) FilterTrustedForwarderUpdated(opts *bind.FilterOpts, forwarder []common.Address) (*VotechainBaseTrustedForwarderUpdatedIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _VotechainBase.contract.FilterLogs(opts, "TrustedForwarderUpdated", forwarderRule)
	if err != nil {
		return nil, err
	}
	return &VotechainBaseTrustedForwarderUpdatedIterator{contract: _VotechainBase.contract, event: "TrustedForwarderUpdated", logs: logs, sub: sub}, nil
}

// WatchTrustedForwarderUpdated is a free log subscription operation binding the contract event 0xbee55516e29d3969d3cb8eb01351eb3c52d06f9e2435bd5a8bfe3647e185df92.
//
// Solidity: event TrustedForwarderUpdated(address indexed forwarder, bool trusted)
func (_VotechainBase *VotechainBaseFilterer) WatchTrustedForwarderUpdated(opts *bind.WatchOpts, sink chan<- *VotechainBaseTrustedForwarderUpdated, forwarder []common.Address) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _VotechainBase.contract.WatchLogs(opts, "TrustedForwarderUpdated", forwarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotechainBaseTrustedForwarderUpdated)
				if err := _VotechainBase.contract.UnpackLog(event, "TrustedForwarderUpdated", log); err != nil {
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

// ParseTrustedForwarderUpdated is a log parse operation binding the contract event 0xbee55516e29d3969d3cb8eb01351eb3c52d06f9e2435bd5a8bfe3647e185df92.
//
// Solidity: event TrustedForwarderUpdated(address indexed forwarder, bool trusted)
func (_VotechainBase *VotechainBaseFilterer) ParseTrustedForwarderUpdated(log types.Log) (*VotechainBaseTrustedForwarderUpdated, error) {
	event := new(VotechainBaseTrustedForwarderUpdated)
	if err := _VotechainBase.contract.UnpackLog(event, "TrustedForwarderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
