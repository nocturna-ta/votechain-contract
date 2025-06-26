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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_kpuManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voterManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_electionManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"trustedForwarders\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyVoted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidElection\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VoterNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotingNotActive\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"TrustedForwarderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"TrustedForwarderRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"}],\"name\":\"addElection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"addTrustedForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"contractIVotechainBase\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"electionManager\",\"outputs\":[{\"internalType\":\"contractIElectionManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTrustedForwarders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kpuManager\",\"outputs\":[{\"internalType\":\"contractIKPUManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"registerKPUKota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"registerKPUProvinsi\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"}],\"name\":\"registerVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"removeTrustedForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setKpuAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setVotingStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"}],\"name\":\"toggleElectionActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"updateKPUKota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"updateKPUProvinsi\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterManager\",\"outputs\":[{\"internalType\":\"contractIVoterManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040516122a23803806122a283398181016040528101906100319190610485565b805f5f90505b81518110156100765761006982828151811061005657610055610518565b5b602002602001015161018260201b60201c565b8080600101915050610037565b50508460025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508360035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050610545565b5f5f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff166102c75760015f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550600181908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f3ef8564460ada92419608d823c014975d98f8104d7d1e68c222967ac6814cd9760405160405180910390a25b50565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610304826102db565b9050919050565b610314816102fa565b811461031e575f5ffd5b50565b5f8151905061032f8161030b565b92915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61037f82610339565b810181811067ffffffffffffffff8211171561039e5761039d610349565b5b80604052505050565b5f6103b06102ca565b90506103bc8282610376565b919050565b5f67ffffffffffffffff8211156103db576103da610349565b5b602082029050602081019050919050565b5f5ffd5b5f6104026103fd846103c1565b6103a7565b90508083825260208201905060208402830185811115610425576104246103ec565b5b835b8181101561044e578061043a8882610321565b845260208401935050602081019050610427565b5050509392505050565b5f82601f83011261046c5761046b610335565b5b815161047c8482602086016103f0565b91505092915050565b5f5f5f5f5f60a0868803121561049e5761049d6102d3565b5b5f6104ab88828901610321565b95505060206104bc88828901610321565b94505060406104cd88828901610321565b93505060606104de88828901610321565b925050608086015167ffffffffffffffff8111156104ff576104fe6102d7565b5b61050b88828901610458565b9150509295509295909350565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b611d50806105525f395ff3fe608060405234801561000f575f5ffd5b5060043610610114575f3560e01c806365910a4f116100a057806395b799071161006f57806395b79907146102a05780639df86dc1146102bc578063a9d93bf0146102d8578063e8d5940d146102f4578063ff7196c81461031057610114565b806365910a4f1461022c5780637478c9fe1461024a5780637da3294914610266578063924238151461028457610114565b80633751d89c116100e75780633751d89c1461018a5780634a075de2146101a65780634da1b97a146101c25780635001f3b5146101de578063572b6c05146101fc57610114565b806302819b991461011857806326d00668146101345780632db33a2f1461015057806334d1d2f21461016c575b5f5ffd5b610132600480360381019061012d91906113ca565b61032e565b005b61014e6004803603810190610149919061145b565b6103c4565b005b61016a600480360381019061016591906113ca565b6104e4565b005b61017461057a565b60405161018191906114e1565b60405180910390f35b6101a4600480360381019061019f919061145b565b61059f565b005b6101c060048036038101906101bb91906114fa565b6106bf565b005b6101dc60048036038101906101d791906113ca565b61074f565b005b6101e66107e5565b6040516101f39190611577565b60405180910390f35b6102166004803603810190610211919061145b565b61080a565b60405161022391906115aa565b60405180910390f35b61023461085b565b60405161024191906115e3565b60405180910390f35b610264600480360381019061025f9190611626565b610880565b005b61026e61090a565b60405161027b9190611671565b60405180910390f35b61029e600480360381019061029991906113ca565b61092f565b005b6102ba60048036038101906102b5919061168a565b6109c5565b005b6102d660048036038101906102d1919061145b565b610a58565b005b6102f260048036038101906102ed919061168a565b610ae2565b005b61030e6004803603810190610309919061168a565b610b75565b005b610318610e82565b60405161032591906117bf565b60405180910390f35b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166302819b9986868686866040518663ffffffff1660e01b8152600401610390959493929190611848565b5f604051808303815f87803b1580156103a7575f5ffd5b505af11580156103b9573d5f5f3e3d5ffd5b505050505050505050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639df86dc160025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fb4ab1646040518163ffffffff1660e01b8152600401602060405180830381865afa15801561046b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061048f91906118a3565b6040518263ffffffff1660e01b81526004016104ab91906118ce565b5f604051808303815f87803b1580156104c2575f5ffd5b505af11580156104d4573d5f5f3e3d5ffd5b505050506104e181610f0d565b50565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632db33a2f86868686866040518663ffffffff1660e01b8152600401610546959493929190611848565b5f604051808303815f87803b15801561055d575f5ffd5b505af115801561056f573d5f5f3e3d5ffd5b505050505050505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639df86dc160025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fb4ab1646040518163ffffffff1660e01b8152600401602060405180830381865afa158015610646573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061066a91906118a3565b6040518263ffffffff1660e01b815260040161068691906118ce565b5f604051808303815f87803b15801561069d575f5ffd5b505af11580156106af573d5f5f3e3d5ffd5b505050506106bc81611055565b50565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634a075de28484846040518463ffffffff1660e01b815260040161071d939291906118e7565b5f604051808303815f87803b158015610734575f5ffd5b505af1158015610746573d5f5f3e3d5ffd5b50505050505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634da1b97a86868686866040518663ffffffff1660e01b81526004016107b1959493929190611848565b5f604051808303815f87803b1580156107c8575f5ffd5b505af11580156107da573d5f5f3e3d5ffd5b505050505050505050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637478c9fe826040518263ffffffff1660e01b81526004016108da91906115aa565b5f604051808303815f87803b1580156108f1575f5ffd5b505af1158015610903573d5f5f3e3d5ffd5b5050505050565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639242381586868686866040518663ffffffff1660e01b8152600401610991959493929190611848565b5f604051808303815f87803b1580156109a8575f5ffd5b505af11580156109ba573d5f5f3e3d5ffd5b505050505050505050565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166395b79907858585856040518563ffffffff1660e01b8152600401610a259493929190611917565b5f604051808303815f87803b158015610a3c575f5ffd5b505af1158015610a4e573d5f5f3e3d5ffd5b5050505050505050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639df86dc1826040518263ffffffff1660e01b8152600401610ab291906118ce565b5f604051808303815f87803b158015610ac9575f5ffd5b505af1158015610adb573d5f5f3e3d5ffd5b5050505050565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9d93bf0858585856040518563ffffffff1660e01b8152600401610b429493929190611917565b5f604051808303815f87803b158015610b59575f5ffd5b505af1158015610b6b573d5f5f3e3d5ffd5b5050505050505050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663408e27276040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bdf573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c039190611964565b610c39576040517f9b8cc47500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610c426112b8565b905060045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634bdd7585826040518263ffffffff1660e01b8152600401610c9e91906118ce565b5f60405180830381865afa925050508015610cdb57506040513d5f823e3d601f19601f82011682018060405250810190610cd89190611b80565b60015b610d11576040517f6f08c58700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806040015115610d4d576040517f7c9a1cf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a8fd0cfa836040518263ffffffff1660e01b8152600401610da791906118ce565b6020604051808303815f875af1158015610dc3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610de79190611964565b5060055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166395f59fe187878787865f01516040518663ffffffff1660e01b8152600401610e4d959493929190611c09565b5f604051808303815f87803b158015610e64575f5ffd5b505af1158015610e76573d5f5f3e3d5ffd5b50505050505050505050565b60606001805480602002602001604051908101604052809291908181526020018280548015610f0357602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610eba575b5050505050905090565b5f5f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff166110525760015f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550600181908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f3ef8564460ada92419608d823c014975d98f8104d7d1e68c222967ac6814cd9760405160405180910390a25b50565b5f5f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16156112b5575f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f5f90505b600180549050811015611270578173ffffffffffffffffffffffffffffffffffffffff166001828154811061113357611132611c57565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036112635760018080805490506111889190611cba565b8154811061119957611198611c57565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600182815481106111d5576111d4611c57565b5b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600180548061122c5761122b611ced565b5b600190038181905f5260205f20015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055611270565b80806001019150506110fb565b508073ffffffffffffffffffffffffffffffffffffffff167fd2d636efcad0cea42e170256f4c5d8b1cd81e47b855557edaf44014e6cc4cee660405160405180910390a25b50565b5f6112c16112c6565b905090565b5f6112d03361080a565b156112e457601436033560601c90506112f3565b6112ec6112f7565b90506112f4565b5b90565b5f33905090565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6113388261130f565b9050919050565b6113488161132e565b8114611352575f5ffd5b50565b5f813590506113638161133f565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261138a57611389611369565b5b8235905067ffffffffffffffff8111156113a7576113a661136d565b5b6020830191508360018202830111156113c3576113c2611371565b5b9250929050565b5f5f5f5f5f606086880312156113e3576113e2611307565b5b5f6113f088828901611355565b955050602086013567ffffffffffffffff8111156114115761141061130b565b5b61141d88828901611375565b9450945050604086013567ffffffffffffffff8111156114405761143f61130b565b5b61144c88828901611375565b92509250509295509295909350565b5f602082840312156114705761146f611307565b5b5f61147d84828501611355565b91505092915050565b5f819050919050565b5f6114a96114a461149f8461130f565b611486565b61130f565b9050919050565b5f6114ba8261148f565b9050919050565b5f6114cb826114b0565b9050919050565b6114db816114c1565b82525050565b5f6020820190506114f45f8301846114d2565b92915050565b5f5f5f6040848603121561151157611510611307565b5b5f84013567ffffffffffffffff81111561152e5761152d61130b565b5b61153a86828701611375565b9350935050602061154d86828701611355565b9150509250925092565b5f611561826114b0565b9050919050565b61157181611557565b82525050565b5f60208201905061158a5f830184611568565b92915050565b5f8115159050919050565b6115a481611590565b82525050565b5f6020820190506115bd5f83018461159b565b92915050565b5f6115cd826114b0565b9050919050565b6115dd816115c3565b82525050565b5f6020820190506115f65f8301846115d4565b92915050565b61160581611590565b811461160f575f5ffd5b50565b5f81359050611620816115fc565b92915050565b5f6020828403121561163b5761163a611307565b5b5f61164884828501611612565b91505092915050565b5f61165b826114b0565b9050919050565b61166b81611651565b82525050565b5f6020820190506116845f830184611662565b92915050565b5f5f5f5f604085870312156116a2576116a1611307565b5b5f85013567ffffffffffffffff8111156116bf576116be61130b565b5b6116cb87828801611375565b9450945050602085013567ffffffffffffffff8111156116ee576116ed61130b565b5b6116fa87828801611375565b925092505092959194509250565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61173a8161132e565b82525050565b5f61174b8383611731565b60208301905092915050565b5f602082019050919050565b5f61176d82611708565b6117778185611712565b935061178283611722565b805f5b838110156117b25781516117998882611740565b97506117a483611757565b925050600181019050611785565b5085935050505092915050565b5f6020820190508181035f8301526117d78184611763565b905092915050565b6117e88161132e565b82525050565b5f82825260208201905092915050565b828183375f83830152505050565b5f601f19601f8301169050919050565b5f61182783856117ee565b93506118348385846117fe565b61183d8361180c565b840190509392505050565b5f60608201905061185b5f8301886117df565b818103602083015261186e81868861181c565b9050818103604083015261188381848661181c565b90509695505050505050565b5f8151905061189d8161133f565b92915050565b5f602082840312156118b8576118b7611307565b5b5f6118c58482850161188f565b91505092915050565b5f6020820190506118e15f8301846117df565b92915050565b5f6040820190508181035f83015261190081858761181c565b905061190f60208301846117df565b949350505050565b5f6040820190508181035f83015261193081868861181c565b9050818103602083015261194581848661181c565b905095945050505050565b5f8151905061195e816115fc565b92915050565b5f6020828403121561197957611978611307565b5b5f61198684828501611950565b91505092915050565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6119c98261180c565b810181811067ffffffffffffffff821117156119e8576119e7611993565b5b80604052505050565b5f6119fa6112fe565b9050611a0682826119c0565b919050565b5f5ffd5b5f5ffd5b5f67ffffffffffffffff821115611a2d57611a2c611993565b5b611a368261180c565b9050602081019050919050565b8281835e5f83830152505050565b5f611a63611a5e84611a13565b6119f1565b905082815260208101848484011115611a7f57611a7e611a0f565b5b611a8a848285611a43565b509392505050565b5f82601f830112611aa657611aa5611369565b5b8151611ab6848260208601611a51565b91505092915050565b5f60a08284031215611ad457611ad361198f565b5b611ade60a06119f1565b90505f82015167ffffffffffffffff811115611afd57611afc611a0b565b5b611b0984828501611a92565b5f830152506020611b1c8482850161188f565b6020830152506040611b3084828501611950565b604083015250606082015167ffffffffffffffff811115611b5457611b53611a0b565b5b611b6084828501611a92565b6060830152506080611b7484828501611950565b60808301525092915050565b5f60208284031215611b9557611b94611307565b5b5f82015167ffffffffffffffff811115611bb257611bb161130b565b5b611bbe84828501611abf565b91505092915050565b5f81519050919050565b5f611bdb82611bc7565b611be581856117ee565b9350611bf5818560208601611a43565b611bfe8161180c565b840191505092915050565b5f6060820190508181035f830152611c2281878961181c565b90508181036020830152611c3781858761181c565b90508181036040830152611c4b8184611bd1565b90509695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f819050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611cc482611c84565b9150611ccf83611c84565b9250828203905081811115611ce757611ce6611c8d565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea26469706673582212202976c44bb54dbd3ee5b5cc1ba0931b776f7edd6bc610c1f62f07cfc305e71a7d64736f6c634300081c0033",
}

// VotechainABI is the input ABI used to generate the binding from.
// Deprecated: Use VotechainMetaData.ABI instead.
var VotechainABI = VotechainMetaData.ABI

// VotechainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VotechainMetaData.Bin instead.
var VotechainBin = VotechainMetaData.Bin

// DeployVotechain deploys a new Ethereum contract, binding an instance of Votechain to it.
func DeployVotechain(auth *bind.TransactOpts, backend bind.ContractBackend, _baseAddress common.Address, _kpuManagerAddress common.Address, _voterManagerAddress common.Address, _electionManagerAddress common.Address, trustedForwarders []common.Address) (common.Address, *types.Transaction, *Votechain, error) {
	parsed, err := VotechainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VotechainBin), backend, _baseAddress, _kpuManagerAddress, _voterManagerAddress, _electionManagerAddress, trustedForwarders)
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

// GetTrustedForwarders is a free data retrieval call binding the contract method 0xff7196c8.
//
// Solidity: function getTrustedForwarders() view returns(address[])
func (_Votechain *VotechainCaller) GetTrustedForwarders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "getTrustedForwarders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTrustedForwarders is a free data retrieval call binding the contract method 0xff7196c8.
//
// Solidity: function getTrustedForwarders() view returns(address[])
func (_Votechain *VotechainSession) GetTrustedForwarders() ([]common.Address, error) {
	return _Votechain.Contract.GetTrustedForwarders(&_Votechain.CallOpts)
}

// GetTrustedForwarders is a free data retrieval call binding the contract method 0xff7196c8.
//
// Solidity: function getTrustedForwarders() view returns(address[])
func (_Votechain *VotechainCallerSession) GetTrustedForwarders() ([]common.Address, error) {
	return _Votechain.Contract.GetTrustedForwarders(&_Votechain.CallOpts)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Votechain *VotechainCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Votechain *VotechainSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Votechain.Contract.IsTrustedForwarder(&_Votechain.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Votechain *VotechainCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Votechain.Contract.IsTrustedForwarder(&_Votechain.CallOpts, forwarder)
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

// AddElection is a paid mutator transaction binding the contract method 0x95b79907.
//
// Solidity: function addElection(string electionId, string electionNo) returns()
func (_Votechain *VotechainTransactor) AddElection(opts *bind.TransactOpts, electionId string, electionNo string) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "addElection", electionId, electionNo)
}

// AddElection is a paid mutator transaction binding the contract method 0x95b79907.
//
// Solidity: function addElection(string electionId, string electionNo) returns()
func (_Votechain *VotechainSession) AddElection(electionId string, electionNo string) (*types.Transaction, error) {
	return _Votechain.Contract.AddElection(&_Votechain.TransactOpts, electionId, electionNo)
}

// AddElection is a paid mutator transaction binding the contract method 0x95b79907.
//
// Solidity: function addElection(string electionId, string electionNo) returns()
func (_Votechain *VotechainTransactorSession) AddElection(electionId string, electionNo string) (*types.Transaction, error) {
	return _Votechain.Contract.AddElection(&_Votechain.TransactOpts, electionId, electionNo)
}

// AddTrustedForwarder is a paid mutator transaction binding the contract method 0x26d00668.
//
// Solidity: function addTrustedForwarder(address forwarder) returns()
func (_Votechain *VotechainTransactor) AddTrustedForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "addTrustedForwarder", forwarder)
}

// AddTrustedForwarder is a paid mutator transaction binding the contract method 0x26d00668.
//
// Solidity: function addTrustedForwarder(address forwarder) returns()
func (_Votechain *VotechainSession) AddTrustedForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _Votechain.Contract.AddTrustedForwarder(&_Votechain.TransactOpts, forwarder)
}

// AddTrustedForwarder is a paid mutator transaction binding the contract method 0x26d00668.
//
// Solidity: function addTrustedForwarder(address forwarder) returns()
func (_Votechain *VotechainTransactorSession) AddTrustedForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _Votechain.Contract.AddTrustedForwarder(&_Votechain.TransactOpts, forwarder)
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

// RemoveTrustedForwarder is a paid mutator transaction binding the contract method 0x3751d89c.
//
// Solidity: function removeTrustedForwarder(address forwarder) returns()
func (_Votechain *VotechainTransactor) RemoveTrustedForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "removeTrustedForwarder", forwarder)
}

// RemoveTrustedForwarder is a paid mutator transaction binding the contract method 0x3751d89c.
//
// Solidity: function removeTrustedForwarder(address forwarder) returns()
func (_Votechain *VotechainSession) RemoveTrustedForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _Votechain.Contract.RemoveTrustedForwarder(&_Votechain.TransactOpts, forwarder)
}

// RemoveTrustedForwarder is a paid mutator transaction binding the contract method 0x3751d89c.
//
// Solidity: function removeTrustedForwarder(address forwarder) returns()
func (_Votechain *VotechainTransactorSession) RemoveTrustedForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _Votechain.Contract.RemoveTrustedForwarder(&_Votechain.TransactOpts, forwarder)
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

// ToggleElectionActive is a paid mutator transaction binding the contract method 0xa9d93bf0.
//
// Solidity: function toggleElectionActive(string electionId, string electionNo) returns()
func (_Votechain *VotechainTransactor) ToggleElectionActive(opts *bind.TransactOpts, electionId string, electionNo string) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "toggleElectionActive", electionId, electionNo)
}

// ToggleElectionActive is a paid mutator transaction binding the contract method 0xa9d93bf0.
//
// Solidity: function toggleElectionActive(string electionId, string electionNo) returns()
func (_Votechain *VotechainSession) ToggleElectionActive(electionId string, electionNo string) (*types.Transaction, error) {
	return _Votechain.Contract.ToggleElectionActive(&_Votechain.TransactOpts, electionId, electionNo)
}

// ToggleElectionActive is a paid mutator transaction binding the contract method 0xa9d93bf0.
//
// Solidity: function toggleElectionActive(string electionId, string electionNo) returns()
func (_Votechain *VotechainTransactorSession) ToggleElectionActive(electionId string, electionNo string) (*types.Transaction, error) {
	return _Votechain.Contract.ToggleElectionActive(&_Votechain.TransactOpts, electionId, electionNo)
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

// Vote is a paid mutator transaction binding the contract method 0xe8d5940d.
//
// Solidity: function vote(string electionId, string electionNo) returns()
func (_Votechain *VotechainTransactor) Vote(opts *bind.TransactOpts, electionId string, electionNo string) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "vote", electionId, electionNo)
}

// Vote is a paid mutator transaction binding the contract method 0xe8d5940d.
//
// Solidity: function vote(string electionId, string electionNo) returns()
func (_Votechain *VotechainSession) Vote(electionId string, electionNo string) (*types.Transaction, error) {
	return _Votechain.Contract.Vote(&_Votechain.TransactOpts, electionId, electionNo)
}

// Vote is a paid mutator transaction binding the contract method 0xe8d5940d.
//
// Solidity: function vote(string electionId, string electionNo) returns()
func (_Votechain *VotechainTransactorSession) Vote(electionId string, electionNo string) (*types.Transaction, error) {
	return _Votechain.Contract.Vote(&_Votechain.TransactOpts, electionId, electionNo)
}

// VotechainTrustedForwarderAddedIterator is returned from FilterTrustedForwarderAdded and is used to iterate over the raw logs and unpacked data for TrustedForwarderAdded events raised by the Votechain contract.
type VotechainTrustedForwarderAddedIterator struct {
	Event *VotechainTrustedForwarderAdded // Event containing the contract specifics and raw log

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
func (it *VotechainTrustedForwarderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotechainTrustedForwarderAdded)
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
		it.Event = new(VotechainTrustedForwarderAdded)
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
func (it *VotechainTrustedForwarderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotechainTrustedForwarderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotechainTrustedForwarderAdded represents a TrustedForwarderAdded event raised by the Votechain contract.
type VotechainTrustedForwarderAdded struct {
	Forwarder common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTrustedForwarderAdded is a free log retrieval operation binding the contract event 0x3ef8564460ada92419608d823c014975d98f8104d7d1e68c222967ac6814cd97.
//
// Solidity: event TrustedForwarderAdded(address indexed forwarder)
func (_Votechain *VotechainFilterer) FilterTrustedForwarderAdded(opts *bind.FilterOpts, forwarder []common.Address) (*VotechainTrustedForwarderAddedIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _Votechain.contract.FilterLogs(opts, "TrustedForwarderAdded", forwarderRule)
	if err != nil {
		return nil, err
	}
	return &VotechainTrustedForwarderAddedIterator{contract: _Votechain.contract, event: "TrustedForwarderAdded", logs: logs, sub: sub}, nil
}

// WatchTrustedForwarderAdded is a free log subscription operation binding the contract event 0x3ef8564460ada92419608d823c014975d98f8104d7d1e68c222967ac6814cd97.
//
// Solidity: event TrustedForwarderAdded(address indexed forwarder)
func (_Votechain *VotechainFilterer) WatchTrustedForwarderAdded(opts *bind.WatchOpts, sink chan<- *VotechainTrustedForwarderAdded, forwarder []common.Address) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _Votechain.contract.WatchLogs(opts, "TrustedForwarderAdded", forwarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotechainTrustedForwarderAdded)
				if err := _Votechain.contract.UnpackLog(event, "TrustedForwarderAdded", log); err != nil {
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
func (_Votechain *VotechainFilterer) ParseTrustedForwarderAdded(log types.Log) (*VotechainTrustedForwarderAdded, error) {
	event := new(VotechainTrustedForwarderAdded)
	if err := _Votechain.contract.UnpackLog(event, "TrustedForwarderAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotechainTrustedForwarderRemovedIterator is returned from FilterTrustedForwarderRemoved and is used to iterate over the raw logs and unpacked data for TrustedForwarderRemoved events raised by the Votechain contract.
type VotechainTrustedForwarderRemovedIterator struct {
	Event *VotechainTrustedForwarderRemoved // Event containing the contract specifics and raw log

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
func (it *VotechainTrustedForwarderRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotechainTrustedForwarderRemoved)
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
		it.Event = new(VotechainTrustedForwarderRemoved)
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
func (it *VotechainTrustedForwarderRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotechainTrustedForwarderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotechainTrustedForwarderRemoved represents a TrustedForwarderRemoved event raised by the Votechain contract.
type VotechainTrustedForwarderRemoved struct {
	Forwarder common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTrustedForwarderRemoved is a free log retrieval operation binding the contract event 0xd2d636efcad0cea42e170256f4c5d8b1cd81e47b855557edaf44014e6cc4cee6.
//
// Solidity: event TrustedForwarderRemoved(address indexed forwarder)
func (_Votechain *VotechainFilterer) FilterTrustedForwarderRemoved(opts *bind.FilterOpts, forwarder []common.Address) (*VotechainTrustedForwarderRemovedIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _Votechain.contract.FilterLogs(opts, "TrustedForwarderRemoved", forwarderRule)
	if err != nil {
		return nil, err
	}
	return &VotechainTrustedForwarderRemovedIterator{contract: _Votechain.contract, event: "TrustedForwarderRemoved", logs: logs, sub: sub}, nil
}

// WatchTrustedForwarderRemoved is a free log subscription operation binding the contract event 0xd2d636efcad0cea42e170256f4c5d8b1cd81e47b855557edaf44014e6cc4cee6.
//
// Solidity: event TrustedForwarderRemoved(address indexed forwarder)
func (_Votechain *VotechainFilterer) WatchTrustedForwarderRemoved(opts *bind.WatchOpts, sink chan<- *VotechainTrustedForwarderRemoved, forwarder []common.Address) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _Votechain.contract.WatchLogs(opts, "TrustedForwarderRemoved", forwarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotechainTrustedForwarderRemoved)
				if err := _Votechain.contract.UnpackLog(event, "TrustedForwarderRemoved", log); err != nil {
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
func (_Votechain *VotechainFilterer) ParseTrustedForwarderRemoved(log types.Log) (*VotechainTrustedForwarderRemoved, error) {
	event := new(VotechainTrustedForwarderRemoved)
	if err := _Votechain.contract.UnpackLog(event, "TrustedForwarderRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
