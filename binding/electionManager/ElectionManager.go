// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package electionManager

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

// IElectionManagerElection is an auto generated low-level Go binding around an user-defined struct.
type IElectionManagerElection struct {
	Id           string
	ElectionName string
	ElectionNo   string
	VoteCount    *big.Int
	IsActive     bool
}

// ElectionManagerMetaData contains all meta data concerning the ElectionManager contract.
var ElectionManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voterManagerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ElectionAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidElection\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedCaller\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"electionName\",\"type\":\"string\"}],\"name\":\"ElectionAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"ElectionStatusChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"candidateId\",\"type\":\"string\"}],\"name\":\"VoteCasted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"_electionss\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"}],\"name\":\"addElection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"contractIVotechainBase\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"electionAddressArray\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"}],\"name\":\"elections\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structIElectionManager.Election\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllElection\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structIElectionManager.Election[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"}],\"name\":\"getElection\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structIElectionManager.Election\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"}],\"name\":\"getElectionByNo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structIElectionManager.Election\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"}],\"name\":\"toggleElectionActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"voterNik\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterManager\",\"outputs\":[{\"internalType\":\"contractIVoterManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161277b38038061277b83398181016040528101906100319190610115565b815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050610153565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100e4826100bb565b9050919050565b6100f4816100da565b81146100fe575f5ffd5b50565b5f8151905061010f816100eb565b92915050565b5f5f6040838503121561012b5761012a6100b7565b5b5f61013885828601610101565b925050602061014985828601610101565b9150509250929050565b61261b806101605f395ff3fe608060405234801561000f575f5ffd5b50600436106100a7575f3560e01c80639f9070d51161006f5780639f9070d514610165578063db5dfdd714610199578063df3cf749146101b5578063e030cba0146101e5578063e8d5940d14610219578063fd1c6e6114610235576100a7565b80634771b573146100ab5780634a3ccc96146100db5780635001f3b51461010b5780635597f0831461012957806365910a4f14610147575b5f5ffd5b6100c560048036038101906100c09190611876565b610251565b6040516100d291906119e4565b60405180910390f35b6100f560048036038101906100f09190611876565b6104c6565b60405161010291906119e4565b60405180910390f35b6101136106d4565b6040516101209190611a7e565b60405180910390f35b6101316106f8565b60405161013e9190611bd3565b60405180910390f35b61014f610927565b60405161015c9190611c13565b60405180910390f35b61017f600480360381019061017a9190611c56565b61094c565b604051610190959493929190611ce7565b60405180910390f35b6101b360048036038101906101ae9190611876565b610b2b565b005b6101cf60048036038101906101ca9190611876565b610d6f565b6040516101dc91906119e4565b60405180910390f35b6101ff60048036038101906101fa9190611e75565b61102a565b604051610210959493929190611ce7565b60405180910390f35b610233600480360381019061022e9190611ebc565b611212565b005b61024f600480360381019061024a9190611f35565b61140c565b005b6102596117d5565b5f6002848460405161026c929190612013565b90815260200160405180910390205f01805461028790612058565b9050036102c0576040517fcd9f14e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600283836040516102d2929190612013565b90815260200160405180910390206040518060a00160405290815f820180546102fa90612058565b80601f016020809104026020016040519081016040528092919081815260200182805461032690612058565b80156103715780601f1061034857610100808354040283529160200191610371565b820191905f5260205f20905b81548152906001019060200180831161035457829003601f168201915b5050505050815260200160018201805461038a90612058565b80601f01602080910402602001604051908101604052809291908181526020018280546103b690612058565b80156104015780601f106103d857610100808354040283529160200191610401565b820191905f5260205f20905b8154815290600101906020018083116103e457829003601f168201915b5050505050815260200160028201805461041a90612058565b80601f016020809104026020016040519081016040528092919081815260200182805461044690612058565b80156104915780601f1061046857610100808354040283529160200191610491565b820191905f5260205f20905b81548152906001019060200180831161047457829003601f168201915b5050505050815260200160038201548152602001600482015f9054906101000a900460ff161515151581525050905092915050565b6104ce6117d5565b600283836040516104e0929190612013565b90815260200160405180910390206040518060a00160405290815f8201805461050890612058565b80601f016020809104026020016040519081016040528092919081815260200182805461053490612058565b801561057f5780601f106105565761010080835404028352916020019161057f565b820191905f5260205f20905b81548152906001019060200180831161056257829003601f168201915b5050505050815260200160018201805461059890612058565b80601f01602080910402602001604051908101604052809291908181526020018280546105c490612058565b801561060f5780601f106105e65761010080835404028352916020019161060f565b820191905f5260205f20905b8154815290600101906020018083116105f257829003601f168201915b5050505050815260200160028201805461062890612058565b80601f016020809104026020016040519081016040528092919081815260200182805461065490612058565b801561069f5780601f106106765761010080835404028352916020019161069f565b820191905f5260205f20905b81548152906001019060200180831161068257829003601f168201915b5050505050815260200160038201548152602001600482015f9054906101000a900460ff161515151581525050905092915050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606003805480602002602001604051908101604052809291908181526020015f905b8282101561091e578382905f5260205f2090600502016040518060a00160405290815f8201805461074b90612058565b80601f016020809104026020016040519081016040528092919081815260200182805461077790612058565b80156107c25780601f10610799576101008083540402835291602001916107c2565b820191905f5260205f20905b8154815290600101906020018083116107a557829003601f168201915b505050505081526020016001820180546107db90612058565b80601f016020809104026020016040519081016040528092919081815260200182805461080790612058565b80156108525780601f1061082957610100808354040283529160200191610852565b820191905f5260205f20905b81548152906001019060200180831161083557829003601f168201915b5050505050815260200160028201805461086b90612058565b80601f016020809104026020016040519081016040528092919081815260200182805461089790612058565b80156108e25780601f106108b9576101008083540402835291602001916108e2565b820191905f5260205f20905b8154815290600101906020018083116108c557829003601f168201915b5050505050815260200160038201548152602001600482015f9054906101000a900460ff1615151515815250508152602001906001019061071b565b50505050905090565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6003818154811061095b575f80fd5b905f5260205f2090600502015f91509050805f01805461097a90612058565b80601f01602080910402602001604051908101604052809291908181526020018280546109a690612058565b80156109f15780601f106109c8576101008083540402835291602001916109f1565b820191905f5260205f20905b8154815290600101906020018083116109d457829003601f168201915b505050505090806001018054610a0690612058565b80601f0160208091040260200160405190810160405280929190818152602001828054610a3290612058565b8015610a7d5780601f10610a5457610100808354040283529160200191610a7d565b820191905f5260205f20905b815481529060010190602001808311610a6057829003601f168201915b505050505090806002018054610a9290612058565b80601f0160208091040260200160405190810160405280929190818152602001828054610abe90612058565b8015610b095780601f10610ae057610100808354040283529160200191610b09565b820191905f5260205f20905b815481529060010190602001808311610aec57829003601f168201915b505050505090806003015490806004015f9054906101000a900460ff16905085565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fb4ab1646040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b94573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bb891906120c3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c1c576040517f5c427cd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60028383604051610c2f929190612013565b90815260200160405180910390205f018054610c4a90612058565b905003610c83576040517fcd9f14e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028282604051610c95929190612013565b90815260200160405180910390206004015f9054906101000a900460ff161560028383604051610cc6929190612013565b90815260200160405180910390206004015f6101000a81548160ff0219169083151502179055508181604051610cfd929190612013565b60405180910390207f4cdbc6bfe3fe60973ba81af65167b9aeabf33c024d02873823ddbddb210b68e660028484604051610d38929190612013565b90815260200160405180910390206004015f9054906101000a900460ff16604051610d6391906120ee565b60405180910390a25050565b610d776117d5565b5f5f90505b600380549050811015610ff1578383604051610d99929190612135565b604051809103902060038281548110610db557610db461214d565b5b905f5260205f209060050201600201604051610dd1919061220c565b604051809103902003610fe45760038181548110610df257610df161214d565b5b905f5260205f2090600502016040518060a00160405290815f82018054610e1890612058565b80601f0160208091040260200160405190810160405280929190818152602001828054610e4490612058565b8015610e8f5780601f10610e6657610100808354040283529160200191610e8f565b820191905f5260205f20905b815481529060010190602001808311610e7257829003601f168201915b50505050508152602001600182018054610ea890612058565b80601f0160208091040260200160405190810160405280929190818152602001828054610ed490612058565b8015610f1f5780601f10610ef657610100808354040283529160200191610f1f565b820191905f5260205f20905b815481529060010190602001808311610f0257829003601f168201915b50505050508152602001600282018054610f3890612058565b80601f0160208091040260200160405190810160405280929190818152602001828054610f6490612058565b8015610faf5780601f10610f8657610100808354040283529160200191610faf565b820191905f5260205f20905b815481529060010190602001808311610f9257829003601f168201915b5050505050815260200160038201548152602001600482015f9054906101000a900460ff161515151581525050915050611024565b8080600101915050610d7c565b506040517fcd9f14e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b92915050565b6002818051602081018201805184825260208301602085012081835280955050505050505f91509050805f01805461106190612058565b80601f016020809104026020016040519081016040528092919081815260200182805461108d90612058565b80156110d85780601f106110af576101008083540402835291602001916110d8565b820191905f5260205f20905b8154815290600101906020018083116110bb57829003601f168201915b5050505050908060010180546110ed90612058565b80601f016020809104026020016040519081016040528092919081815260200182805461111990612058565b80156111645780601f1061113b57610100808354040283529160200191611164565b820191905f5260205f20905b81548152906001019060200180831161114757829003601f168201915b50505050509080600201805461117990612058565b80601f01602080910402602001604051908101604052809291908181526020018280546111a590612058565b80156111f05780601f106111c7576101008083540402835291602001916111f0565b820191905f5260205f20905b8154815290600101906020018083116111d357829003601f168201915b505050505090806003015490806004015f9054906101000a900460ff16905085565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663408e27276040518163ffffffff1660e01b8152600401602060405180830381865afa15801561127b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061129f919061224c565b6112d5576040517f5c427cd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f600284846040516112e8929190612013565b90815260200160405180910390205f01805461130390612058565b9050148061133e57506002838360405161131e929190612013565b90815260200160405180910390206004015f9054906101000a900460ff16155b15611375576040517fcd9f14e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028383604051611387929190612013565b90815260200160405180910390206003015f8154809291906113a8906122a4565b919050555082826040516113bd929190612013565b6040518091039020816040516113d3919061231b565b60405180910390207f791f7d5f0b0d6e798f239ccca607156ff12f293d6709379301dc1a27e520618160405160405180910390a3505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fb4ab1646040518163ffffffff1660e01b8152600401602060405180830381865afa158015611475573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061149991906120c3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146114fd576040517f5c427cd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60028787604051611510929190612013565b90815260200160405180910390205f01805461152b90612058565b905014611564576040517fe44dbe0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6040518060a0016040528088888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200184848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020015f8152602001600115158152509050806002888860405161166d929190612013565b90815260200160405180910390205f820151815f01908161168e91906124c8565b5060208201518160010190816116a491906124c8565b5060408201518160020190816116ba91906124c8565b50606082015181600301556080820151816004015f6101000a81548160ff021916908315150217905550905050600381908060018154018082558091505060019003905f5260205f2090600502015f909190919091505f820151815f01908161172391906124c8565b50602082015181600101908161173991906124c8565b50604082015181600201908161174f91906124c8565b50606082015181600301556080820151816004015f6101000a81548160ff0219169083151502179055505050868660405161178b929190612013565b60405180910390207ffbb39612bc65a9d2c5c0137dc505f3a2fa0d4c0dac443a06fcf9df0c281c55b286866040516117c49291906125c3565b60405180910390a250505050505050565b6040518060a001604052806060815260200160608152602001606081526020015f81526020015f151581525090565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261183657611835611815565b5b8235905067ffffffffffffffff81111561185357611852611819565b5b60208301915083600182028301111561186f5761186e61181d565b5b9250929050565b5f5f6020838503121561188c5761188b61180d565b5b5f83013567ffffffffffffffff8111156118a9576118a8611811565b5b6118b585828601611821565b92509250509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f611903826118c1565b61190d81856118cb565b935061191d8185602086016118db565b611926816118e9565b840191505092915050565b5f819050919050565b61194381611931565b82525050565b5f8115159050919050565b61195d81611949565b82525050565b5f60a083015f8301518482035f86015261197d82826118f9565b9150506020830151848203602086015261199782826118f9565b915050604083015184820360408601526119b182826118f9565b91505060608301516119c6606086018261193a565b5060808301516119d96080860182611954565b508091505092915050565b5f6020820190508181035f8301526119fc8184611963565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f611a46611a41611a3c84611a04565b611a23565b611a04565b9050919050565b5f611a5782611a2c565b9050919050565b5f611a6882611a4d565b9050919050565b611a7881611a5e565b82525050565b5f602082019050611a915f830184611a6f565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f60a083015f8301518482035f860152611ada82826118f9565b91505060208301518482036020860152611af482826118f9565b91505060408301518482036040860152611b0e82826118f9565b9150506060830151611b23606086018261193a565b506080830151611b366080860182611954565b508091505092915050565b5f611b4c8383611ac0565b905092915050565b5f602082019050919050565b5f611b6a82611a97565b611b748185611aa1565b935083602082028501611b8685611ab1565b805f5b85811015611bc15784840389528151611ba28582611b41565b9450611bad83611b54565b925060208a01995050600181019050611b89565b50829750879550505050505092915050565b5f6020820190508181035f830152611beb8184611b60565b905092915050565b5f611bfd82611a4d565b9050919050565b611c0d81611bf3565b82525050565b5f602082019050611c265f830184611c04565b92915050565b611c3581611931565b8114611c3f575f5ffd5b50565b5f81359050611c5081611c2c565b92915050565b5f60208284031215611c6b57611c6a61180d565b5b5f611c7884828501611c42565b91505092915050565b5f82825260208201905092915050565b5f611c9b826118c1565b611ca58185611c81565b9350611cb58185602086016118db565b611cbe816118e9565b840191505092915050565b611cd281611931565b82525050565b611ce181611949565b82525050565b5f60a0820190508181035f830152611cff8188611c91565b90508181036020830152611d138187611c91565b90508181036040830152611d278186611c91565b9050611d366060830185611cc9565b611d436080830184611cd8565b9695505050505050565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611d87826118e9565b810181811067ffffffffffffffff82111715611da657611da5611d51565b5b80604052505050565b5f611db8611804565b9050611dc48282611d7e565b919050565b5f67ffffffffffffffff821115611de357611de2611d51565b5b611dec826118e9565b9050602081019050919050565b828183375f83830152505050565b5f611e19611e1484611dc9565b611daf565b905082815260208101848484011115611e3557611e34611d4d565b5b611e40848285611df9565b509392505050565b5f82601f830112611e5c57611e5b611815565b5b8135611e6c848260208601611e07565b91505092915050565b5f60208284031215611e8a57611e8961180d565b5b5f82013567ffffffffffffffff811115611ea757611ea6611811565b5b611eb384828501611e48565b91505092915050565b5f5f5f60408486031215611ed357611ed261180d565b5b5f84013567ffffffffffffffff811115611ef057611eef611811565b5b611efc86828701611821565b9350935050602084013567ffffffffffffffff811115611f1f57611f1e611811565b5b611f2b86828701611e48565b9150509250925092565b5f5f5f5f5f5f60608789031215611f4f57611f4e61180d565b5b5f87013567ffffffffffffffff811115611f6c57611f6b611811565b5b611f7889828a01611821565b9650965050602087013567ffffffffffffffff811115611f9b57611f9a611811565b5b611fa789828a01611821565b9450945050604087013567ffffffffffffffff811115611fca57611fc9611811565b5b611fd689828a01611821565b92509250509295509295509295565b5f81905092915050565b5f611ffa8385611fe5565b9350612007838584611df9565b82840190509392505050565b5f61201f828486611fef565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061206f57607f821691505b6020821081036120825761208161202b565b5b50919050565b5f61209282611a04565b9050919050565b6120a281612088565b81146120ac575f5ffd5b50565b5f815190506120bd81612099565b92915050565b5f602082840312156120d8576120d761180d565b5b5f6120e5848285016120af565b91505092915050565b5f6020820190506121015f830184611cd8565b92915050565b5f81905092915050565b5f61211c8385612107565b9350612129838584611df9565b82840190509392505050565b5f612141828486612111565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f819050815f5260205f209050919050565b5f815461219881612058565b6121a28186612107565b9450600182165f81146121bc57600181146121d157612203565b60ff1983168652811515820286019350612203565b6121da8561217a565b5f5b838110156121fb578154818901526001820191506020810190506121dc565b838801955050505b50505092915050565b5f612217828461218c565b915081905092915050565b61222b81611949565b8114612235575f5ffd5b50565b5f8151905061224681612222565b92915050565b5f602082840312156122615761226061180d565b5b5f61226e84828501612238565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6122ae82611931565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036122e0576122df612277565b5b600182019050919050565b5f6122f5826118c1565b6122ff8185611fe5565b935061230f8185602086016118db565b80840191505092915050565b5f61232682846122eb565b915081905092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261238d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82612352565b6123978683612352565b95508019841693508086168417925050509392505050565b5f6123c96123c46123bf84611931565b611a23565b611931565b9050919050565b5f819050919050565b6123e2836123af565b6123f66123ee826123d0565b84845461235e565b825550505050565b5f5f905090565b61240d6123fe565b6124188184846123d9565b505050565b5b8181101561243b576124305f82612405565b60018101905061241e565b5050565b601f8211156124805761245181612331565b61245a84612343565b81016020851015612469578190505b61247d61247585612343565b83018261241d565b50505b505050565b5f82821c905092915050565b5f6124a05f1984600802612485565b1980831691505092915050565b5f6124b88383612491565b9150826002028217905092915050565b6124d1826118c1565b67ffffffffffffffff8111156124ea576124e9611d51565b5b6124f48254612058565b6124ff82828561243f565b5f60209050601f831160018114612530575f841561251e578287015190505b61252885826124ad565b86555061258f565b601f19841661253e86612331565b5f5b8281101561256557848901518255600182019150602085019450602081019050612540565b86831015612582578489015161257e601f891682612491565b8355505b6001600288020188555050505b505050505050565b5f6125a28385611c81565b93506125af838584611df9565b6125b8836118e9565b840190509392505050565b5f6020820190508181035f8301526125dc818486612597565b9050939250505056fea2646970667358221220c1da3e07d6c914872f7feb634a084edf5ad798817b0bd3721ad2657dcd174e1b64736f6c634300081c0033",
}

// ElectionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ElectionManagerMetaData.ABI instead.
var ElectionManagerABI = ElectionManagerMetaData.ABI

// ElectionManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ElectionManagerMetaData.Bin instead.
var ElectionManagerBin = ElectionManagerMetaData.Bin

// DeployElectionManager deploys a new Ethereum contract, binding an instance of ElectionManager to it.
func DeployElectionManager(auth *bind.TransactOpts, backend bind.ContractBackend, _baseAddress common.Address, _voterManagerAddress common.Address) (common.Address, *types.Transaction, *ElectionManager, error) {
	parsed, err := ElectionManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ElectionManagerBin), backend, _baseAddress, _voterManagerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ElectionManager{ElectionManagerCaller: ElectionManagerCaller{contract: contract}, ElectionManagerTransactor: ElectionManagerTransactor{contract: contract}, ElectionManagerFilterer: ElectionManagerFilterer{contract: contract}}, nil
}

// ElectionManager is an auto generated Go binding around an Ethereum contract.
type ElectionManager struct {
	ElectionManagerCaller     // Read-only binding to the contract
	ElectionManagerTransactor // Write-only binding to the contract
	ElectionManagerFilterer   // Log filterer for contract events
}

// ElectionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ElectionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ElectionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ElectionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ElectionManagerSession struct {
	Contract     *ElectionManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ElectionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ElectionManagerCallerSession struct {
	Contract *ElectionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ElectionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ElectionManagerTransactorSession struct {
	Contract     *ElectionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ElectionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ElectionManagerRaw struct {
	Contract *ElectionManager // Generic contract binding to access the raw methods on
}

// ElectionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ElectionManagerCallerRaw struct {
	Contract *ElectionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ElectionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ElectionManagerTransactorRaw struct {
	Contract *ElectionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElectionManager creates a new instance of ElectionManager, bound to a specific deployed contract.
func NewElectionManager(address common.Address, backend bind.ContractBackend) (*ElectionManager, error) {
	contract, err := bindElectionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ElectionManager{ElectionManagerCaller: ElectionManagerCaller{contract: contract}, ElectionManagerTransactor: ElectionManagerTransactor{contract: contract}, ElectionManagerFilterer: ElectionManagerFilterer{contract: contract}}, nil
}

// NewElectionManagerCaller creates a new read-only instance of ElectionManager, bound to a specific deployed contract.
func NewElectionManagerCaller(address common.Address, caller bind.ContractCaller) (*ElectionManagerCaller, error) {
	contract, err := bindElectionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ElectionManagerCaller{contract: contract}, nil
}

// NewElectionManagerTransactor creates a new write-only instance of ElectionManager, bound to a specific deployed contract.
func NewElectionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ElectionManagerTransactor, error) {
	contract, err := bindElectionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ElectionManagerTransactor{contract: contract}, nil
}

// NewElectionManagerFilterer creates a new log filterer instance of ElectionManager, bound to a specific deployed contract.
func NewElectionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ElectionManagerFilterer, error) {
	contract, err := bindElectionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ElectionManagerFilterer{contract: contract}, nil
}

// bindElectionManager binds a generic wrapper to an already deployed contract.
func bindElectionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ElectionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElectionManager *ElectionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ElectionManager.Contract.ElectionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElectionManager *ElectionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionManager.Contract.ElectionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElectionManager *ElectionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElectionManager.Contract.ElectionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElectionManager *ElectionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ElectionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElectionManager *ElectionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElectionManager *ElectionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElectionManager.Contract.contract.Transact(opts, method, params...)
}

// Electionss is a free data retrieval call binding the contract method 0xe030cba0.
//
// Solidity: function _electionss(string ) view returns(string id, string electionName, string electionNo, uint256 voteCount, bool isActive)
func (_ElectionManager *ElectionManagerCaller) Electionss(opts *bind.CallOpts, arg0 string) (struct {
	Id           string
	ElectionName string
	ElectionNo   string
	VoteCount    *big.Int
	IsActive     bool
}, error) {
	var out []interface{}
	err := _ElectionManager.contract.Call(opts, &out, "_electionss", arg0)

	outstruct := new(struct {
		Id           string
		ElectionName string
		ElectionNo   string
		VoteCount    *big.Int
		IsActive     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ElectionName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ElectionNo = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Electionss is a free data retrieval call binding the contract method 0xe030cba0.
//
// Solidity: function _electionss(string ) view returns(string id, string electionName, string electionNo, uint256 voteCount, bool isActive)
func (_ElectionManager *ElectionManagerSession) Electionss(arg0 string) (struct {
	Id           string
	ElectionName string
	ElectionNo   string
	VoteCount    *big.Int
	IsActive     bool
}, error) {
	return _ElectionManager.Contract.Electionss(&_ElectionManager.CallOpts, arg0)
}

// Electionss is a free data retrieval call binding the contract method 0xe030cba0.
//
// Solidity: function _electionss(string ) view returns(string id, string electionName, string electionNo, uint256 voteCount, bool isActive)
func (_ElectionManager *ElectionManagerCallerSession) Electionss(arg0 string) (struct {
	Id           string
	ElectionName string
	ElectionNo   string
	VoteCount    *big.Int
	IsActive     bool
}, error) {
	return _ElectionManager.Contract.Electionss(&_ElectionManager.CallOpts, arg0)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_ElectionManager *ElectionManagerCaller) Base(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ElectionManager.contract.Call(opts, &out, "base")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_ElectionManager *ElectionManagerSession) Base() (common.Address, error) {
	return _ElectionManager.Contract.Base(&_ElectionManager.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_ElectionManager *ElectionManagerCallerSession) Base() (common.Address, error) {
	return _ElectionManager.Contract.Base(&_ElectionManager.CallOpts)
}

// ElectionAddressArray is a free data retrieval call binding the contract method 0x9f9070d5.
//
// Solidity: function electionAddressArray(uint256 ) view returns(string id, string electionName, string electionNo, uint256 voteCount, bool isActive)
func (_ElectionManager *ElectionManagerCaller) ElectionAddressArray(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id           string
	ElectionName string
	ElectionNo   string
	VoteCount    *big.Int
	IsActive     bool
}, error) {
	var out []interface{}
	err := _ElectionManager.contract.Call(opts, &out, "electionAddressArray", arg0)

	outstruct := new(struct {
		Id           string
		ElectionName string
		ElectionNo   string
		VoteCount    *big.Int
		IsActive     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ElectionName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ElectionNo = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// ElectionAddressArray is a free data retrieval call binding the contract method 0x9f9070d5.
//
// Solidity: function electionAddressArray(uint256 ) view returns(string id, string electionName, string electionNo, uint256 voteCount, bool isActive)
func (_ElectionManager *ElectionManagerSession) ElectionAddressArray(arg0 *big.Int) (struct {
	Id           string
	ElectionName string
	ElectionNo   string
	VoteCount    *big.Int
	IsActive     bool
}, error) {
	return _ElectionManager.Contract.ElectionAddressArray(&_ElectionManager.CallOpts, arg0)
}

// ElectionAddressArray is a free data retrieval call binding the contract method 0x9f9070d5.
//
// Solidity: function electionAddressArray(uint256 ) view returns(string id, string electionName, string electionNo, uint256 voteCount, bool isActive)
func (_ElectionManager *ElectionManagerCallerSession) ElectionAddressArray(arg0 *big.Int) (struct {
	Id           string
	ElectionName string
	ElectionNo   string
	VoteCount    *big.Int
	IsActive     bool
}, error) {
	return _ElectionManager.Contract.ElectionAddressArray(&_ElectionManager.CallOpts, arg0)
}

// Elections is a free data retrieval call binding the contract method 0x4a3ccc96.
//
// Solidity: function elections(string electionId) view returns((string,string,string,uint256,bool))
func (_ElectionManager *ElectionManagerCaller) Elections(opts *bind.CallOpts, electionId string) (IElectionManagerElection, error) {
	var out []interface{}
	err := _ElectionManager.contract.Call(opts, &out, "elections", electionId)

	if err != nil {
		return *new(IElectionManagerElection), err
	}

	out0 := *abi.ConvertType(out[0], new(IElectionManagerElection)).(*IElectionManagerElection)

	return out0, err

}

// Elections is a free data retrieval call binding the contract method 0x4a3ccc96.
//
// Solidity: function elections(string electionId) view returns((string,string,string,uint256,bool))
func (_ElectionManager *ElectionManagerSession) Elections(electionId string) (IElectionManagerElection, error) {
	return _ElectionManager.Contract.Elections(&_ElectionManager.CallOpts, electionId)
}

// Elections is a free data retrieval call binding the contract method 0x4a3ccc96.
//
// Solidity: function elections(string electionId) view returns((string,string,string,uint256,bool))
func (_ElectionManager *ElectionManagerCallerSession) Elections(electionId string) (IElectionManagerElection, error) {
	return _ElectionManager.Contract.Elections(&_ElectionManager.CallOpts, electionId)
}

// GetAllElection is a free data retrieval call binding the contract method 0x5597f083.
//
// Solidity: function getAllElection() view returns((string,string,string,uint256,bool)[])
func (_ElectionManager *ElectionManagerCaller) GetAllElection(opts *bind.CallOpts) ([]IElectionManagerElection, error) {
	var out []interface{}
	err := _ElectionManager.contract.Call(opts, &out, "getAllElection")

	if err != nil {
		return *new([]IElectionManagerElection), err
	}

	out0 := *abi.ConvertType(out[0], new([]IElectionManagerElection)).(*[]IElectionManagerElection)

	return out0, err

}

// GetAllElection is a free data retrieval call binding the contract method 0x5597f083.
//
// Solidity: function getAllElection() view returns((string,string,string,uint256,bool)[])
func (_ElectionManager *ElectionManagerSession) GetAllElection() ([]IElectionManagerElection, error) {
	return _ElectionManager.Contract.GetAllElection(&_ElectionManager.CallOpts)
}

// GetAllElection is a free data retrieval call binding the contract method 0x5597f083.
//
// Solidity: function getAllElection() view returns((string,string,string,uint256,bool)[])
func (_ElectionManager *ElectionManagerCallerSession) GetAllElection() ([]IElectionManagerElection, error) {
	return _ElectionManager.Contract.GetAllElection(&_ElectionManager.CallOpts)
}

// GetElection is a free data retrieval call binding the contract method 0x4771b573.
//
// Solidity: function getElection(string electionId) view returns((string,string,string,uint256,bool))
func (_ElectionManager *ElectionManagerCaller) GetElection(opts *bind.CallOpts, electionId string) (IElectionManagerElection, error) {
	var out []interface{}
	err := _ElectionManager.contract.Call(opts, &out, "getElection", electionId)

	if err != nil {
		return *new(IElectionManagerElection), err
	}

	out0 := *abi.ConvertType(out[0], new(IElectionManagerElection)).(*IElectionManagerElection)

	return out0, err

}

// GetElection is a free data retrieval call binding the contract method 0x4771b573.
//
// Solidity: function getElection(string electionId) view returns((string,string,string,uint256,bool))
func (_ElectionManager *ElectionManagerSession) GetElection(electionId string) (IElectionManagerElection, error) {
	return _ElectionManager.Contract.GetElection(&_ElectionManager.CallOpts, electionId)
}

// GetElection is a free data retrieval call binding the contract method 0x4771b573.
//
// Solidity: function getElection(string electionId) view returns((string,string,string,uint256,bool))
func (_ElectionManager *ElectionManagerCallerSession) GetElection(electionId string) (IElectionManagerElection, error) {
	return _ElectionManager.Contract.GetElection(&_ElectionManager.CallOpts, electionId)
}

// GetElectionByNo is a free data retrieval call binding the contract method 0xdf3cf749.
//
// Solidity: function getElectionByNo(string electionNo) view returns((string,string,string,uint256,bool))
func (_ElectionManager *ElectionManagerCaller) GetElectionByNo(opts *bind.CallOpts, electionNo string) (IElectionManagerElection, error) {
	var out []interface{}
	err := _ElectionManager.contract.Call(opts, &out, "getElectionByNo", electionNo)

	if err != nil {
		return *new(IElectionManagerElection), err
	}

	out0 := *abi.ConvertType(out[0], new(IElectionManagerElection)).(*IElectionManagerElection)

	return out0, err

}

// GetElectionByNo is a free data retrieval call binding the contract method 0xdf3cf749.
//
// Solidity: function getElectionByNo(string electionNo) view returns((string,string,string,uint256,bool))
func (_ElectionManager *ElectionManagerSession) GetElectionByNo(electionNo string) (IElectionManagerElection, error) {
	return _ElectionManager.Contract.GetElectionByNo(&_ElectionManager.CallOpts, electionNo)
}

// GetElectionByNo is a free data retrieval call binding the contract method 0xdf3cf749.
//
// Solidity: function getElectionByNo(string electionNo) view returns((string,string,string,uint256,bool))
func (_ElectionManager *ElectionManagerCallerSession) GetElectionByNo(electionNo string) (IElectionManagerElection, error) {
	return _ElectionManager.Contract.GetElectionByNo(&_ElectionManager.CallOpts, electionNo)
}

// VoterManager is a free data retrieval call binding the contract method 0x65910a4f.
//
// Solidity: function voterManager() view returns(address)
func (_ElectionManager *ElectionManagerCaller) VoterManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ElectionManager.contract.Call(opts, &out, "voterManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoterManager is a free data retrieval call binding the contract method 0x65910a4f.
//
// Solidity: function voterManager() view returns(address)
func (_ElectionManager *ElectionManagerSession) VoterManager() (common.Address, error) {
	return _ElectionManager.Contract.VoterManager(&_ElectionManager.CallOpts)
}

// VoterManager is a free data retrieval call binding the contract method 0x65910a4f.
//
// Solidity: function voterManager() view returns(address)
func (_ElectionManager *ElectionManagerCallerSession) VoterManager() (common.Address, error) {
	return _ElectionManager.Contract.VoterManager(&_ElectionManager.CallOpts)
}

// AddElection is a paid mutator transaction binding the contract method 0xfd1c6e61.
//
// Solidity: function addElection(string electionId, string electionName, string electionNo) returns()
func (_ElectionManager *ElectionManagerTransactor) AddElection(opts *bind.TransactOpts, electionId string, electionName string, electionNo string) (*types.Transaction, error) {
	return _ElectionManager.contract.Transact(opts, "addElection", electionId, electionName, electionNo)
}

// AddElection is a paid mutator transaction binding the contract method 0xfd1c6e61.
//
// Solidity: function addElection(string electionId, string electionName, string electionNo) returns()
func (_ElectionManager *ElectionManagerSession) AddElection(electionId string, electionName string, electionNo string) (*types.Transaction, error) {
	return _ElectionManager.Contract.AddElection(&_ElectionManager.TransactOpts, electionId, electionName, electionNo)
}

// AddElection is a paid mutator transaction binding the contract method 0xfd1c6e61.
//
// Solidity: function addElection(string electionId, string electionName, string electionNo) returns()
func (_ElectionManager *ElectionManagerTransactorSession) AddElection(electionId string, electionName string, electionNo string) (*types.Transaction, error) {
	return _ElectionManager.Contract.AddElection(&_ElectionManager.TransactOpts, electionId, electionName, electionNo)
}

// ToggleElectionActive is a paid mutator transaction binding the contract method 0xdb5dfdd7.
//
// Solidity: function toggleElectionActive(string electionId) returns()
func (_ElectionManager *ElectionManagerTransactor) ToggleElectionActive(opts *bind.TransactOpts, electionId string) (*types.Transaction, error) {
	return _ElectionManager.contract.Transact(opts, "toggleElectionActive", electionId)
}

// ToggleElectionActive is a paid mutator transaction binding the contract method 0xdb5dfdd7.
//
// Solidity: function toggleElectionActive(string electionId) returns()
func (_ElectionManager *ElectionManagerSession) ToggleElectionActive(electionId string) (*types.Transaction, error) {
	return _ElectionManager.Contract.ToggleElectionActive(&_ElectionManager.TransactOpts, electionId)
}

// ToggleElectionActive is a paid mutator transaction binding the contract method 0xdb5dfdd7.
//
// Solidity: function toggleElectionActive(string electionId) returns()
func (_ElectionManager *ElectionManagerTransactorSession) ToggleElectionActive(electionId string) (*types.Transaction, error) {
	return _ElectionManager.Contract.ToggleElectionActive(&_ElectionManager.TransactOpts, electionId)
}

// Vote is a paid mutator transaction binding the contract method 0xe8d5940d.
//
// Solidity: function vote(string electionId, string voterNik) returns()
func (_ElectionManager *ElectionManagerTransactor) Vote(opts *bind.TransactOpts, electionId string, voterNik string) (*types.Transaction, error) {
	return _ElectionManager.contract.Transact(opts, "vote", electionId, voterNik)
}

// Vote is a paid mutator transaction binding the contract method 0xe8d5940d.
//
// Solidity: function vote(string electionId, string voterNik) returns()
func (_ElectionManager *ElectionManagerSession) Vote(electionId string, voterNik string) (*types.Transaction, error) {
	return _ElectionManager.Contract.Vote(&_ElectionManager.TransactOpts, electionId, voterNik)
}

// Vote is a paid mutator transaction binding the contract method 0xe8d5940d.
//
// Solidity: function vote(string electionId, string voterNik) returns()
func (_ElectionManager *ElectionManagerTransactorSession) Vote(electionId string, voterNik string) (*types.Transaction, error) {
	return _ElectionManager.Contract.Vote(&_ElectionManager.TransactOpts, electionId, voterNik)
}

// ElectionManagerElectionAddedIterator is returned from FilterElectionAdded and is used to iterate over the raw logs and unpacked data for ElectionAdded events raised by the ElectionManager contract.
type ElectionManagerElectionAddedIterator struct {
	Event *ElectionManagerElectionAdded // Event containing the contract specifics and raw log

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
func (it *ElectionManagerElectionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionManagerElectionAdded)
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
		it.Event = new(ElectionManagerElectionAdded)
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
func (it *ElectionManagerElectionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionManagerElectionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionManagerElectionAdded represents a ElectionAdded event raised by the ElectionManager contract.
type ElectionManagerElectionAdded struct {
	ElectionId   common.Hash
	ElectionName string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterElectionAdded is a free log retrieval operation binding the contract event 0xfbb39612bc65a9d2c5c0137dc505f3a2fa0d4c0dac443a06fcf9df0c281c55b2.
//
// Solidity: event ElectionAdded(string indexed electionId, string electionName)
func (_ElectionManager *ElectionManagerFilterer) FilterElectionAdded(opts *bind.FilterOpts, electionId []string) (*ElectionManagerElectionAddedIterator, error) {

	var electionIdRule []interface{}
	for _, electionIdItem := range electionId {
		electionIdRule = append(electionIdRule, electionIdItem)
	}

	logs, sub, err := _ElectionManager.contract.FilterLogs(opts, "ElectionAdded", electionIdRule)
	if err != nil {
		return nil, err
	}
	return &ElectionManagerElectionAddedIterator{contract: _ElectionManager.contract, event: "ElectionAdded", logs: logs, sub: sub}, nil
}

// WatchElectionAdded is a free log subscription operation binding the contract event 0xfbb39612bc65a9d2c5c0137dc505f3a2fa0d4c0dac443a06fcf9df0c281c55b2.
//
// Solidity: event ElectionAdded(string indexed electionId, string electionName)
func (_ElectionManager *ElectionManagerFilterer) WatchElectionAdded(opts *bind.WatchOpts, sink chan<- *ElectionManagerElectionAdded, electionId []string) (event.Subscription, error) {

	var electionIdRule []interface{}
	for _, electionIdItem := range electionId {
		electionIdRule = append(electionIdRule, electionIdItem)
	}

	logs, sub, err := _ElectionManager.contract.WatchLogs(opts, "ElectionAdded", electionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionManagerElectionAdded)
				if err := _ElectionManager.contract.UnpackLog(event, "ElectionAdded", log); err != nil {
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

// ParseElectionAdded is a log parse operation binding the contract event 0xfbb39612bc65a9d2c5c0137dc505f3a2fa0d4c0dac443a06fcf9df0c281c55b2.
//
// Solidity: event ElectionAdded(string indexed electionId, string electionName)
func (_ElectionManager *ElectionManagerFilterer) ParseElectionAdded(log types.Log) (*ElectionManagerElectionAdded, error) {
	event := new(ElectionManagerElectionAdded)
	if err := _ElectionManager.contract.UnpackLog(event, "ElectionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionManagerElectionStatusChangeIterator is returned from FilterElectionStatusChange and is used to iterate over the raw logs and unpacked data for ElectionStatusChange events raised by the ElectionManager contract.
type ElectionManagerElectionStatusChangeIterator struct {
	Event *ElectionManagerElectionStatusChange // Event containing the contract specifics and raw log

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
func (it *ElectionManagerElectionStatusChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionManagerElectionStatusChange)
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
		it.Event = new(ElectionManagerElectionStatusChange)
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
func (it *ElectionManagerElectionStatusChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionManagerElectionStatusChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionManagerElectionStatusChange represents a ElectionStatusChange event raised by the ElectionManager contract.
type ElectionManagerElectionStatusChange struct {
	ElectionId common.Hash
	IsActive   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterElectionStatusChange is a free log retrieval operation binding the contract event 0x4cdbc6bfe3fe60973ba81af65167b9aeabf33c024d02873823ddbddb210b68e6.
//
// Solidity: event ElectionStatusChange(string indexed electionId, bool isActive)
func (_ElectionManager *ElectionManagerFilterer) FilterElectionStatusChange(opts *bind.FilterOpts, electionId []string) (*ElectionManagerElectionStatusChangeIterator, error) {

	var electionIdRule []interface{}
	for _, electionIdItem := range electionId {
		electionIdRule = append(electionIdRule, electionIdItem)
	}

	logs, sub, err := _ElectionManager.contract.FilterLogs(opts, "ElectionStatusChange", electionIdRule)
	if err != nil {
		return nil, err
	}
	return &ElectionManagerElectionStatusChangeIterator{contract: _ElectionManager.contract, event: "ElectionStatusChange", logs: logs, sub: sub}, nil
}

// WatchElectionStatusChange is a free log subscription operation binding the contract event 0x4cdbc6bfe3fe60973ba81af65167b9aeabf33c024d02873823ddbddb210b68e6.
//
// Solidity: event ElectionStatusChange(string indexed electionId, bool isActive)
func (_ElectionManager *ElectionManagerFilterer) WatchElectionStatusChange(opts *bind.WatchOpts, sink chan<- *ElectionManagerElectionStatusChange, electionId []string) (event.Subscription, error) {

	var electionIdRule []interface{}
	for _, electionIdItem := range electionId {
		electionIdRule = append(electionIdRule, electionIdItem)
	}

	logs, sub, err := _ElectionManager.contract.WatchLogs(opts, "ElectionStatusChange", electionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionManagerElectionStatusChange)
				if err := _ElectionManager.contract.UnpackLog(event, "ElectionStatusChange", log); err != nil {
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

// ParseElectionStatusChange is a log parse operation binding the contract event 0x4cdbc6bfe3fe60973ba81af65167b9aeabf33c024d02873823ddbddb210b68e6.
//
// Solidity: event ElectionStatusChange(string indexed electionId, bool isActive)
func (_ElectionManager *ElectionManagerFilterer) ParseElectionStatusChange(log types.Log) (*ElectionManagerElectionStatusChange, error) {
	event := new(ElectionManagerElectionStatusChange)
	if err := _ElectionManager.contract.UnpackLog(event, "ElectionStatusChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionManagerVoteCastedIterator is returned from FilterVoteCasted and is used to iterate over the raw logs and unpacked data for VoteCasted events raised by the ElectionManager contract.
type ElectionManagerVoteCastedIterator struct {
	Event *ElectionManagerVoteCasted // Event containing the contract specifics and raw log

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
func (it *ElectionManagerVoteCastedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionManagerVoteCasted)
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
		it.Event = new(ElectionManagerVoteCasted)
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
func (it *ElectionManagerVoteCastedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionManagerVoteCastedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionManagerVoteCasted represents a VoteCasted event raised by the ElectionManager contract.
type ElectionManagerVoteCasted struct {
	Nik         common.Hash
	CandidateId common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVoteCasted is a free log retrieval operation binding the contract event 0x791f7d5f0b0d6e798f239ccca607156ff12f293d6709379301dc1a27e5206181.
//
// Solidity: event VoteCasted(string indexed nik, string indexed candidateId)
func (_ElectionManager *ElectionManagerFilterer) FilterVoteCasted(opts *bind.FilterOpts, nik []string, candidateId []string) (*ElectionManagerVoteCastedIterator, error) {

	var nikRule []interface{}
	for _, nikItem := range nik {
		nikRule = append(nikRule, nikItem)
	}
	var candidateIdRule []interface{}
	for _, candidateIdItem := range candidateId {
		candidateIdRule = append(candidateIdRule, candidateIdItem)
	}

	logs, sub, err := _ElectionManager.contract.FilterLogs(opts, "VoteCasted", nikRule, candidateIdRule)
	if err != nil {
		return nil, err
	}
	return &ElectionManagerVoteCastedIterator{contract: _ElectionManager.contract, event: "VoteCasted", logs: logs, sub: sub}, nil
}

// WatchVoteCasted is a free log subscription operation binding the contract event 0x791f7d5f0b0d6e798f239ccca607156ff12f293d6709379301dc1a27e5206181.
//
// Solidity: event VoteCasted(string indexed nik, string indexed candidateId)
func (_ElectionManager *ElectionManagerFilterer) WatchVoteCasted(opts *bind.WatchOpts, sink chan<- *ElectionManagerVoteCasted, nik []string, candidateId []string) (event.Subscription, error) {

	var nikRule []interface{}
	for _, nikItem := range nik {
		nikRule = append(nikRule, nikItem)
	}
	var candidateIdRule []interface{}
	for _, candidateIdItem := range candidateId {
		candidateIdRule = append(candidateIdRule, candidateIdItem)
	}

	logs, sub, err := _ElectionManager.contract.WatchLogs(opts, "VoteCasted", nikRule, candidateIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionManagerVoteCasted)
				if err := _ElectionManager.contract.UnpackLog(event, "VoteCasted", log); err != nil {
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

// ParseVoteCasted is a log parse operation binding the contract event 0x791f7d5f0b0d6e798f239ccca607156ff12f293d6709379301dc1a27e5206181.
//
// Solidity: event VoteCasted(string indexed nik, string indexed candidateId)
func (_ElectionManager *ElectionManagerFilterer) ParseVoteCasted(log types.Log) (*ElectionManagerVoteCasted, error) {
	event := new(ElectionManagerVoteCasted)
	if err := _ElectionManager.contract.UnpackLog(event, "VoteCasted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
