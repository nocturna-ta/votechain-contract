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
	Id         string
	ElectionNo string
	VoteCount  *big.Int
	IsActive   bool
}

// ElectionManagerMetaData contains all meta data concerning the ElectionManager contract.
var ElectionManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voterManagerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ElectionAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ElectionNumberMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidElection\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedCaller\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"}],\"name\":\"ElectionAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"ElectionStatusChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"}],\"name\":\"VoteCasted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"_electionss\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"}],\"name\":\"addElection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"contractIVotechainBase\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"electionAddressArray\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"}],\"name\":\"elections\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structIElectionManager.Election\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllElection\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structIElectionManager.Election[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"}],\"name\":\"getElection\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structIElectionManager.Election\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"}],\"name\":\"getElectionByNo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structIElectionManager.Election\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"}],\"name\":\"toggleElectionActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"electionId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"electionNo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"voterNik\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterManager\",\"outputs\":[{\"internalType\":\"contractIVoterManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161258e38038061258e83398181016040528101906100319190610115565b815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050610153565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100e4826100bb565b9050919050565b6100f4816100da565b81146100fe575f5ffd5b50565b5f8151905061010f816100eb565b92915050565b5f5f6040838503121561012b5761012a6100b7565b5b5f61013885828601610101565b925050602061014985828601610101565b9150509250929050565b61242e806101605f395ff3fe608060405234801561000f575f5ffd5b50600436106100a7575f3560e01c806395b799071161006f57806395b799071461016557806395f59fe1146101815780639f9070d51461019d578063a9d93bf0146101d0578063df3cf749146101ec578063e030cba01461021c576100a7565b80634771b573146100ab5780634a3ccc96146100db5780635001f3b51461010b5780635597f0831461012957806365910a4f14610147575b5f5ffd5b6100c560048036038101906100c091906116d0565b61024f565b6040516100d29190611824565b60405180910390f35b6100f560048036038101906100f091906116d0565b61041a565b6040516101029190611824565b60405180910390f35b610113610598565b60405161012091906118be565b60405180910390f35b6101316105bc565b60405161013e91906119f9565b60405180910390f35b61014f61075b565b60405161015c9190611a39565b60405180910390f35b61017f600480360381019061017a9190611a52565b610780565b005b61019b60048036038101906101969190611bf8565b610cc4565b005b6101b760048036038101906101b29190611ccf565b610eee565b6040516101c79493929190611d60565b60405180910390f35b6101ea60048036038101906101e59190611a52565b611041565b005b610206600480360381019061020191906116d0565b6112af565b6040516102139190611824565b60405180910390f35b61023660048036038101906102319190611db1565b6114da565b6040516102469493929190611d60565b60405180910390f35b610257611636565b5f6002848460405161026a929190611e26565b908152602001604051809103902090505f815f01805461028990611e6b565b9050036102c2576040517fcd9f14e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806040518060800160405290815f820180546102dd90611e6b565b80601f016020809104026020016040519081016040528092919081815260200182805461030990611e6b565b80156103545780601f1061032b57610100808354040283529160200191610354565b820191905f5260205f20905b81548152906001019060200180831161033757829003601f168201915b5050505050815260200160018201805461036d90611e6b565b80601f016020809104026020016040519081016040528092919081815260200182805461039990611e6b565b80156103e45780601f106103bb576101008083540402835291602001916103e4565b820191905f5260205f20905b8154815290600101906020018083116103c757829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff16151515158152505091505092915050565b610422611636565b60028383604051610434929190611e26565b90815260200160405180910390206040518060800160405290815f8201805461045c90611e6b565b80601f016020809104026020016040519081016040528092919081815260200182805461048890611e6b565b80156104d35780601f106104aa576101008083540402835291602001916104d3565b820191905f5260205f20905b8154815290600101906020018083116104b657829003601f168201915b505050505081526020016001820180546104ec90611e6b565b80601f016020809104026020016040519081016040528092919081815260200182805461051890611e6b565b80156105635780601f1061053a57610100808354040283529160200191610563565b820191905f5260205f20905b81548152906001019060200180831161054657829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff161515151581525050905092915050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606003805480602002602001604051908101604052809291908181526020015f905b82821015610752578382905f5260205f2090600402016040518060800160405290815f8201805461060f90611e6b565b80601f016020809104026020016040519081016040528092919081815260200182805461063b90611e6b565b80156106865780601f1061065d57610100808354040283529160200191610686565b820191905f5260205f20905b81548152906001019060200180831161066957829003601f168201915b5050505050815260200160018201805461069f90611e6b565b80601f01602080910402602001604051908101604052809291908181526020018280546106cb90611e6b565b80156107165780601f106106ed57610100808354040283529160200191610716565b820191905f5260205f20905b8154815290600101906020018083116106f957829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff161515151581525050815260200190600101906105df565b50505050905090565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fb4ab1646040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107e9573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061080d9190611ed6565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610871576040517f5c427cd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60028585604051610884929190611e26565b90815260200160405180910390206040518060800160405290815f820180546108ac90611e6b565b80601f01602080910402602001604051908101604052809291908181526020018280546108d890611e6b565b80156109235780601f106108fa57610100808354040283529160200191610923565b820191905f5260205f20905b81548152906001019060200180831161090657829003601f168201915b5050505050815260200160018201805461093c90611e6b565b80601f016020809104026020016040519081016040528092919081815260200182805461096890611e6b565b80156109b35780601f1061098a576101008083540402835291602001916109b3565b820191905f5260205f20905b81548152906001019060200180831161099657829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff16151515158152505090505f815f01515114610a1f576040517fe44dbe0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f90505b600380549050811015610ac5578383604051610a41929190611f2f565b604051809103902060038281548110610a5d57610a5c611f47565b5b905f5260205f209060040201600101604051610a799190612006565b604051809103902003610ab8576040517fe44dbe0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8080600101915050610a24565b505f604051806080016040528087878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020015f81526020016001151581525090508091508060028787604051610b89929190611e26565b90815260200160405180910390205f820151815f019081610baa91906121b3565b506020820151816001019081610bc091906121b3565b50604082015181600201556060820151816003015f6101000a81548160ff021916908315150217905550905050600381908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f019081610c2991906121b3565b506020820151816001019081610c3f91906121b3565b50604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050508585604051610c7b929190611e26565b60405180910390207ffbb39612bc65a9d2c5c0137dc505f3a2fa0d4c0dac443a06fcf9df0c281c55b28585604051610cb49291906122ae565b60405180910390a2505050505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663408e27276040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d2d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d5191906122fa565b610d87576040517f5c427cd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60028686604051610d9a929190611e26565b908152602001604051809103902090505f815f018054610db990611e6b565b90501480610dd45750806003015f9054906101000a900460ff165b15610e0b576040517fcd9f14e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8383604051610e1b929190611f2f565b604051809103902081600101604051610e349190612006565b604051809103902014610e73576040517f373251d700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806002015f815480929190610e8790612352565b91905055508585604051610e9c929190611e26565b604051809103902082604051610eb291906123c9565b60405180910390207f791f7d5f0b0d6e798f239ccca607156ff12f293d6709379301dc1a27e520618160405160405180910390a3505050505050565b60038181548110610efd575f80fd5b905f5260205f2090600402015f91509050805f018054610f1c90611e6b565b80601f0160208091040260200160405190810160405280929190818152602001828054610f4890611e6b565b8015610f935780601f10610f6a57610100808354040283529160200191610f93565b820191905f5260205f20905b815481529060010190602001808311610f7657829003601f168201915b505050505090806001018054610fa890611e6b565b80601f0160208091040260200160405190810160405280929190818152602001828054610fd490611e6b565b801561101f5780601f10610ff65761010080835404028352916020019161101f565b820191905f5260205f20905b81548152906001019060200180831161100257829003601f168201915b505050505090806002015490806003015f9054906101000a900460ff16905084565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fb4ab1646040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110aa573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110ce9190611ed6565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611132576040517f5c427cd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60028585604051611145929190611e26565b908152602001604051809103902090505f815f01805461116490611e6b565b90500361119d576040517fcd9f14e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82826040516111ad929190611f2f565b6040518091039020816001016040516111c69190612006565b604051809103902014611205576040517f373251d700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806003015f9054906101000a900460ff1615816003015f6101000a81548160ff0219169083151502179055508282604051611241929190611e26565b60405180910390208585604051611259929190611e26565b60405180910390207fc4ed90e5395aebcc1e7e027dfc5cafaa02443f04824a16ab3ad2f6a52f8ca06d836003015f9054906101000a900460ff166040516112a091906123df565b60405180910390a35050505050565b6112b7611636565b5f5f90505b6003805490508110156114a15783836040516112d9929190611f2f565b6040518091039020600382815481106112f5576112f4611f47565b5b905f5260205f2090600402016001016040516113119190612006565b604051809103902003611494576003818154811061133257611331611f47565b5b905f5260205f2090600402016040518060800160405290815f8201805461135890611e6b565b80601f016020809104026020016040519081016040528092919081815260200182805461138490611e6b565b80156113cf5780601f106113a6576101008083540402835291602001916113cf565b820191905f5260205f20905b8154815290600101906020018083116113b257829003601f168201915b505050505081526020016001820180546113e890611e6b565b80601f016020809104026020016040519081016040528092919081815260200182805461141490611e6b565b801561145f5780601f106114365761010080835404028352916020019161145f565b820191905f5260205f20905b81548152906001019060200180831161144257829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff1615151515815250509150506114d4565b80806001019150506112bc565b506040517fcd9f14e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b92915050565b6002818051602081018201805184825260208301602085012081835280955050505050505f91509050805f01805461151190611e6b565b80601f016020809104026020016040519081016040528092919081815260200182805461153d90611e6b565b80156115885780601f1061155f57610100808354040283529160200191611588565b820191905f5260205f20905b81548152906001019060200180831161156b57829003601f168201915b50505050509080600101805461159d90611e6b565b80601f01602080910402602001604051908101604052809291908181526020018280546115c990611e6b565b80156116145780601f106115eb57610100808354040283529160200191611614565b820191905f5260205f20905b8154815290600101906020018083116115f757829003601f168201915b505050505090806002015490806003015f9054906101000a900460ff16905084565b604051806080016040528060608152602001606081526020015f81526020015f151581525090565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126116905761168f61166f565b5b8235905067ffffffffffffffff8111156116ad576116ac611673565b5b6020830191508360018202830111156116c9576116c8611677565b5b9250929050565b5f5f602083850312156116e6576116e5611667565b5b5f83013567ffffffffffffffff8111156117035761170261166b565b5b61170f8582860161167b565b92509250509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61175d8261171b565b6117678185611725565b9350611777818560208601611735565b61178081611743565b840191505092915050565b5f819050919050565b61179d8161178b565b82525050565b5f8115159050919050565b6117b7816117a3565b82525050565b5f608083015f8301518482035f8601526117d78282611753565b915050602083015184820360208601526117f18282611753565b91505060408301516118066040860182611794565b50606083015161181960608601826117ae565b508091505092915050565b5f6020820190508181035f83015261183c81846117bd565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f61188661188161187c84611844565b611863565b611844565b9050919050565b5f6118978261186c565b9050919050565b5f6118a88261188d565b9050919050565b6118b88161189e565b82525050565b5f6020820190506118d15f8301846118af565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f608083015f8301518482035f86015261191a8282611753565b915050602083015184820360208601526119348282611753565b91505060408301516119496040860182611794565b50606083015161195c60608601826117ae565b508091505092915050565b5f6119728383611900565b905092915050565b5f602082019050919050565b5f611990826118d7565b61199a81856118e1565b9350836020820285016119ac856118f1565b805f5b858110156119e757848403895281516119c88582611967565b94506119d38361197a565b925060208a019950506001810190506119af565b50829750879550505050505092915050565b5f6020820190508181035f830152611a118184611986565b905092915050565b5f611a238261188d565b9050919050565b611a3381611a19565b82525050565b5f602082019050611a4c5f830184611a2a565b92915050565b5f5f5f5f60408587031215611a6a57611a69611667565b5b5f85013567ffffffffffffffff811115611a8757611a8661166b565b5b611a938782880161167b565b9450945050602085013567ffffffffffffffff811115611ab657611ab561166b565b5b611ac28782880161167b565b925092505092959194509250565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611b0a82611743565b810181811067ffffffffffffffff82111715611b2957611b28611ad4565b5b80604052505050565b5f611b3b61165e565b9050611b478282611b01565b919050565b5f67ffffffffffffffff821115611b6657611b65611ad4565b5b611b6f82611743565b9050602081019050919050565b828183375f83830152505050565b5f611b9c611b9784611b4c565b611b32565b905082815260208101848484011115611bb857611bb7611ad0565b5b611bc3848285611b7c565b509392505050565b5f82601f830112611bdf57611bde61166f565b5b8135611bef848260208601611b8a565b91505092915050565b5f5f5f5f5f60608688031215611c1157611c10611667565b5b5f86013567ffffffffffffffff811115611c2e57611c2d61166b565b5b611c3a8882890161167b565b9550955050602086013567ffffffffffffffff811115611c5d57611c5c61166b565b5b611c698882890161167b565b9350935050604086013567ffffffffffffffff811115611c8c57611c8b61166b565b5b611c9888828901611bcb565b9150509295509295909350565b611cae8161178b565b8114611cb8575f5ffd5b50565b5f81359050611cc981611ca5565b92915050565b5f60208284031215611ce457611ce3611667565b5b5f611cf184828501611cbb565b91505092915050565b5f82825260208201905092915050565b5f611d148261171b565b611d1e8185611cfa565b9350611d2e818560208601611735565b611d3781611743565b840191505092915050565b611d4b8161178b565b82525050565b611d5a816117a3565b82525050565b5f6080820190508181035f830152611d788187611d0a565b90508181036020830152611d8c8186611d0a565b9050611d9b6040830185611d42565b611da86060830184611d51565b95945050505050565b5f60208284031215611dc657611dc5611667565b5b5f82013567ffffffffffffffff811115611de357611de261166b565b5b611def84828501611bcb565b91505092915050565b5f81905092915050565b5f611e0d8385611df8565b9350611e1a838584611b7c565b82840190509392505050565b5f611e32828486611e02565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611e8257607f821691505b602082108103611e9557611e94611e3e565b5b50919050565b5f611ea582611844565b9050919050565b611eb581611e9b565b8114611ebf575f5ffd5b50565b5f81519050611ed081611eac565b92915050565b5f60208284031215611eeb57611eea611667565b5b5f611ef884828501611ec2565b91505092915050565b5f81905092915050565b5f611f168385611f01565b9350611f23838584611b7c565b82840190509392505050565b5f611f3b828486611f0b565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f819050815f5260205f209050919050565b5f8154611f9281611e6b565b611f9c8186611f01565b9450600182165f8114611fb65760018114611fcb57611ffd565b60ff1983168652811515820286019350611ffd565b611fd485611f74565b5f5b83811015611ff557815481890152600182019150602081019050611fd6565b838801955050505b50505092915050565b5f6120118284611f86565b915081905092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026120787fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261203d565b612082868361203d565b95508019841693508086168417925050509392505050565b5f6120b46120af6120aa8461178b565b611863565b61178b565b9050919050565b5f819050919050565b6120cd8361209a565b6120e16120d9826120bb565b848454612049565b825550505050565b5f5f905090565b6120f86120e9565b6121038184846120c4565b505050565b5b818110156121265761211b5f826120f0565b600181019050612109565b5050565b601f82111561216b5761213c8161201c565b6121458461202e565b81016020851015612154578190505b6121686121608561202e565b830182612108565b50505b505050565b5f82821c905092915050565b5f61218b5f1984600802612170565b1980831691505092915050565b5f6121a3838361217c565b9150826002028217905092915050565b6121bc8261171b565b67ffffffffffffffff8111156121d5576121d4611ad4565b5b6121df8254611e6b565b6121ea82828561212a565b5f60209050601f83116001811461221b575f8415612209578287015190505b6122138582612198565b86555061227a565b601f1984166122298661201c565b5f5b828110156122505784890151825560018201915060208501945060208101905061222b565b8683101561226d5784890151612269601f89168261217c565b8355505b6001600288020188555050505b505050505050565b5f61228d8385611cfa565b935061229a838584611b7c565b6122a383611743565b840190509392505050565b5f6020820190508181035f8301526122c7818486612282565b90509392505050565b6122d9816117a3565b81146122e3575f5ffd5b50565b5f815190506122f4816122d0565b92915050565b5f6020828403121561230f5761230e611667565b5b5f61231c848285016122e6565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61235c8261178b565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361238e5761238d612325565b5b600182019050919050565b5f6123a38261171b565b6123ad8185611df8565b93506123bd818560208601611735565b80840191505092915050565b5f6123d48284612399565b915081905092915050565b5f6020820190506123f25f830184611d51565b9291505056fea2646970667358221220713598eed25495573ce2c4eb17adac00f9ff2e96dbc4f53320a8fe079ae340f464736f6c634300081c0033",
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
// Solidity: function _electionss(string ) view returns(string id, string electionNo, uint256 voteCount, bool isActive)
func (_ElectionManager *ElectionManagerCaller) Electionss(opts *bind.CallOpts, arg0 string) (struct {
	Id         string
	ElectionNo string
	VoteCount  *big.Int
	IsActive   bool
}, error) {
	var out []interface{}
	err := _ElectionManager.contract.Call(opts, &out, "_electionss", arg0)

	outstruct := new(struct {
		Id         string
		ElectionNo string
		VoteCount  *big.Int
		IsActive   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ElectionNo = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Electionss is a free data retrieval call binding the contract method 0xe030cba0.
//
// Solidity: function _electionss(string ) view returns(string id, string electionNo, uint256 voteCount, bool isActive)
func (_ElectionManager *ElectionManagerSession) Electionss(arg0 string) (struct {
	Id         string
	ElectionNo string
	VoteCount  *big.Int
	IsActive   bool
}, error) {
	return _ElectionManager.Contract.Electionss(&_ElectionManager.CallOpts, arg0)
}

// Electionss is a free data retrieval call binding the contract method 0xe030cba0.
//
// Solidity: function _electionss(string ) view returns(string id, string electionNo, uint256 voteCount, bool isActive)
func (_ElectionManager *ElectionManagerCallerSession) Electionss(arg0 string) (struct {
	Id         string
	ElectionNo string
	VoteCount  *big.Int
	IsActive   bool
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
// Solidity: function electionAddressArray(uint256 ) view returns(string id, string electionNo, uint256 voteCount, bool isActive)
func (_ElectionManager *ElectionManagerCaller) ElectionAddressArray(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id         string
	ElectionNo string
	VoteCount  *big.Int
	IsActive   bool
}, error) {
	var out []interface{}
	err := _ElectionManager.contract.Call(opts, &out, "electionAddressArray", arg0)

	outstruct := new(struct {
		Id         string
		ElectionNo string
		VoteCount  *big.Int
		IsActive   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ElectionNo = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// ElectionAddressArray is a free data retrieval call binding the contract method 0x9f9070d5.
//
// Solidity: function electionAddressArray(uint256 ) view returns(string id, string electionNo, uint256 voteCount, bool isActive)
func (_ElectionManager *ElectionManagerSession) ElectionAddressArray(arg0 *big.Int) (struct {
	Id         string
	ElectionNo string
	VoteCount  *big.Int
	IsActive   bool
}, error) {
	return _ElectionManager.Contract.ElectionAddressArray(&_ElectionManager.CallOpts, arg0)
}

// ElectionAddressArray is a free data retrieval call binding the contract method 0x9f9070d5.
//
// Solidity: function electionAddressArray(uint256 ) view returns(string id, string electionNo, uint256 voteCount, bool isActive)
func (_ElectionManager *ElectionManagerCallerSession) ElectionAddressArray(arg0 *big.Int) (struct {
	Id         string
	ElectionNo string
	VoteCount  *big.Int
	IsActive   bool
}, error) {
	return _ElectionManager.Contract.ElectionAddressArray(&_ElectionManager.CallOpts, arg0)
}

// Elections is a free data retrieval call binding the contract method 0x4a3ccc96.
//
// Solidity: function elections(string electionId) view returns((string,string,uint256,bool))
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
// Solidity: function elections(string electionId) view returns((string,string,uint256,bool))
func (_ElectionManager *ElectionManagerSession) Elections(electionId string) (IElectionManagerElection, error) {
	return _ElectionManager.Contract.Elections(&_ElectionManager.CallOpts, electionId)
}

// Elections is a free data retrieval call binding the contract method 0x4a3ccc96.
//
// Solidity: function elections(string electionId) view returns((string,string,uint256,bool))
func (_ElectionManager *ElectionManagerCallerSession) Elections(electionId string) (IElectionManagerElection, error) {
	return _ElectionManager.Contract.Elections(&_ElectionManager.CallOpts, electionId)
}

// GetAllElection is a free data retrieval call binding the contract method 0x5597f083.
//
// Solidity: function getAllElection() view returns((string,string,uint256,bool)[])
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
// Solidity: function getAllElection() view returns((string,string,uint256,bool)[])
func (_ElectionManager *ElectionManagerSession) GetAllElection() ([]IElectionManagerElection, error) {
	return _ElectionManager.Contract.GetAllElection(&_ElectionManager.CallOpts)
}

// GetAllElection is a free data retrieval call binding the contract method 0x5597f083.
//
// Solidity: function getAllElection() view returns((string,string,uint256,bool)[])
func (_ElectionManager *ElectionManagerCallerSession) GetAllElection() ([]IElectionManagerElection, error) {
	return _ElectionManager.Contract.GetAllElection(&_ElectionManager.CallOpts)
}

// GetElection is a free data retrieval call binding the contract method 0x4771b573.
//
// Solidity: function getElection(string electionId) view returns((string,string,uint256,bool))
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
// Solidity: function getElection(string electionId) view returns((string,string,uint256,bool))
func (_ElectionManager *ElectionManagerSession) GetElection(electionId string) (IElectionManagerElection, error) {
	return _ElectionManager.Contract.GetElection(&_ElectionManager.CallOpts, electionId)
}

// GetElection is a free data retrieval call binding the contract method 0x4771b573.
//
// Solidity: function getElection(string electionId) view returns((string,string,uint256,bool))
func (_ElectionManager *ElectionManagerCallerSession) GetElection(electionId string) (IElectionManagerElection, error) {
	return _ElectionManager.Contract.GetElection(&_ElectionManager.CallOpts, electionId)
}

// GetElectionByNo is a free data retrieval call binding the contract method 0xdf3cf749.
//
// Solidity: function getElectionByNo(string electionNo) view returns((string,string,uint256,bool))
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
// Solidity: function getElectionByNo(string electionNo) view returns((string,string,uint256,bool))
func (_ElectionManager *ElectionManagerSession) GetElectionByNo(electionNo string) (IElectionManagerElection, error) {
	return _ElectionManager.Contract.GetElectionByNo(&_ElectionManager.CallOpts, electionNo)
}

// GetElectionByNo is a free data retrieval call binding the contract method 0xdf3cf749.
//
// Solidity: function getElectionByNo(string electionNo) view returns((string,string,uint256,bool))
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

// AddElection is a paid mutator transaction binding the contract method 0x95b79907.
//
// Solidity: function addElection(string electionId, string electionNo) returns()
func (_ElectionManager *ElectionManagerTransactor) AddElection(opts *bind.TransactOpts, electionId string, electionNo string) (*types.Transaction, error) {
	return _ElectionManager.contract.Transact(opts, "addElection", electionId, electionNo)
}

// AddElection is a paid mutator transaction binding the contract method 0x95b79907.
//
// Solidity: function addElection(string electionId, string electionNo) returns()
func (_ElectionManager *ElectionManagerSession) AddElection(electionId string, electionNo string) (*types.Transaction, error) {
	return _ElectionManager.Contract.AddElection(&_ElectionManager.TransactOpts, electionId, electionNo)
}

// AddElection is a paid mutator transaction binding the contract method 0x95b79907.
//
// Solidity: function addElection(string electionId, string electionNo) returns()
func (_ElectionManager *ElectionManagerTransactorSession) AddElection(electionId string, electionNo string) (*types.Transaction, error) {
	return _ElectionManager.Contract.AddElection(&_ElectionManager.TransactOpts, electionId, electionNo)
}

// ToggleElectionActive is a paid mutator transaction binding the contract method 0xa9d93bf0.
//
// Solidity: function toggleElectionActive(string electionId, string electionNo) returns()
func (_ElectionManager *ElectionManagerTransactor) ToggleElectionActive(opts *bind.TransactOpts, electionId string, electionNo string) (*types.Transaction, error) {
	return _ElectionManager.contract.Transact(opts, "toggleElectionActive", electionId, electionNo)
}

// ToggleElectionActive is a paid mutator transaction binding the contract method 0xa9d93bf0.
//
// Solidity: function toggleElectionActive(string electionId, string electionNo) returns()
func (_ElectionManager *ElectionManagerSession) ToggleElectionActive(electionId string, electionNo string) (*types.Transaction, error) {
	return _ElectionManager.Contract.ToggleElectionActive(&_ElectionManager.TransactOpts, electionId, electionNo)
}

// ToggleElectionActive is a paid mutator transaction binding the contract method 0xa9d93bf0.
//
// Solidity: function toggleElectionActive(string electionId, string electionNo) returns()
func (_ElectionManager *ElectionManagerTransactorSession) ToggleElectionActive(electionId string, electionNo string) (*types.Transaction, error) {
	return _ElectionManager.Contract.ToggleElectionActive(&_ElectionManager.TransactOpts, electionId, electionNo)
}

// Vote is a paid mutator transaction binding the contract method 0x95f59fe1.
//
// Solidity: function vote(string electionId, string electionNo, string voterNik) returns()
func (_ElectionManager *ElectionManagerTransactor) Vote(opts *bind.TransactOpts, electionId string, electionNo string, voterNik string) (*types.Transaction, error) {
	return _ElectionManager.contract.Transact(opts, "vote", electionId, electionNo, voterNik)
}

// Vote is a paid mutator transaction binding the contract method 0x95f59fe1.
//
// Solidity: function vote(string electionId, string electionNo, string voterNik) returns()
func (_ElectionManager *ElectionManagerSession) Vote(electionId string, electionNo string, voterNik string) (*types.Transaction, error) {
	return _ElectionManager.Contract.Vote(&_ElectionManager.TransactOpts, electionId, electionNo, voterNik)
}

// Vote is a paid mutator transaction binding the contract method 0x95f59fe1.
//
// Solidity: function vote(string electionId, string electionNo, string voterNik) returns()
func (_ElectionManager *ElectionManagerTransactorSession) Vote(electionId string, electionNo string, voterNik string) (*types.Transaction, error) {
	return _ElectionManager.Contract.Vote(&_ElectionManager.TransactOpts, electionId, electionNo, voterNik)
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
	ElectionId common.Hash
	ElectionNo string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterElectionAdded is a free log retrieval operation binding the contract event 0xfbb39612bc65a9d2c5c0137dc505f3a2fa0d4c0dac443a06fcf9df0c281c55b2.
//
// Solidity: event ElectionAdded(string indexed electionId, string electionNo)
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
// Solidity: event ElectionAdded(string indexed electionId, string electionNo)
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
// Solidity: event ElectionAdded(string indexed electionId, string electionNo)
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
	ElectionNo common.Hash
	IsActive   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterElectionStatusChange is a free log retrieval operation binding the contract event 0xc4ed90e5395aebcc1e7e027dfc5cafaa02443f04824a16ab3ad2f6a52f8ca06d.
//
// Solidity: event ElectionStatusChange(string indexed electionId, string indexed electionNo, bool isActive)
func (_ElectionManager *ElectionManagerFilterer) FilterElectionStatusChange(opts *bind.FilterOpts, electionId []string, electionNo []string) (*ElectionManagerElectionStatusChangeIterator, error) {

	var electionIdRule []interface{}
	for _, electionIdItem := range electionId {
		electionIdRule = append(electionIdRule, electionIdItem)
	}
	var electionNoRule []interface{}
	for _, electionNoItem := range electionNo {
		electionNoRule = append(electionNoRule, electionNoItem)
	}

	logs, sub, err := _ElectionManager.contract.FilterLogs(opts, "ElectionStatusChange", electionIdRule, electionNoRule)
	if err != nil {
		return nil, err
	}
	return &ElectionManagerElectionStatusChangeIterator{contract: _ElectionManager.contract, event: "ElectionStatusChange", logs: logs, sub: sub}, nil
}

// WatchElectionStatusChange is a free log subscription operation binding the contract event 0xc4ed90e5395aebcc1e7e027dfc5cafaa02443f04824a16ab3ad2f6a52f8ca06d.
//
// Solidity: event ElectionStatusChange(string indexed electionId, string indexed electionNo, bool isActive)
func (_ElectionManager *ElectionManagerFilterer) WatchElectionStatusChange(opts *bind.WatchOpts, sink chan<- *ElectionManagerElectionStatusChange, electionId []string, electionNo []string) (event.Subscription, error) {

	var electionIdRule []interface{}
	for _, electionIdItem := range electionId {
		electionIdRule = append(electionIdRule, electionIdItem)
	}
	var electionNoRule []interface{}
	for _, electionNoItem := range electionNo {
		electionNoRule = append(electionNoRule, electionNoItem)
	}

	logs, sub, err := _ElectionManager.contract.WatchLogs(opts, "ElectionStatusChange", electionIdRule, electionNoRule)
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

// ParseElectionStatusChange is a log parse operation binding the contract event 0xc4ed90e5395aebcc1e7e027dfc5cafaa02443f04824a16ab3ad2f6a52f8ca06d.
//
// Solidity: event ElectionStatusChange(string indexed electionId, string indexed electionNo, bool isActive)
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
	Nik        common.Hash
	ElectionId common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCasted is a free log retrieval operation binding the contract event 0x791f7d5f0b0d6e798f239ccca607156ff12f293d6709379301dc1a27e5206181.
//
// Solidity: event VoteCasted(string indexed nik, string indexed electionId)
func (_ElectionManager *ElectionManagerFilterer) FilterVoteCasted(opts *bind.FilterOpts, nik []string, electionId []string) (*ElectionManagerVoteCastedIterator, error) {

	var nikRule []interface{}
	for _, nikItem := range nik {
		nikRule = append(nikRule, nikItem)
	}
	var electionIdRule []interface{}
	for _, electionIdItem := range electionId {
		electionIdRule = append(electionIdRule, electionIdItem)
	}

	logs, sub, err := _ElectionManager.contract.FilterLogs(opts, "VoteCasted", nikRule, electionIdRule)
	if err != nil {
		return nil, err
	}
	return &ElectionManagerVoteCastedIterator{contract: _ElectionManager.contract, event: "VoteCasted", logs: logs, sub: sub}, nil
}

// WatchVoteCasted is a free log subscription operation binding the contract event 0x791f7d5f0b0d6e798f239ccca607156ff12f293d6709379301dc1a27e5206181.
//
// Solidity: event VoteCasted(string indexed nik, string indexed electionId)
func (_ElectionManager *ElectionManagerFilterer) WatchVoteCasted(opts *bind.WatchOpts, sink chan<- *ElectionManagerVoteCasted, nik []string, electionId []string) (event.Subscription, error) {

	var nikRule []interface{}
	for _, nikItem := range nik {
		nikRule = append(nikRule, nikItem)
	}
	var electionIdRule []interface{}
	for _, electionIdItem := range electionId {
		electionIdRule = append(electionIdRule, electionIdItem)
	}

	logs, sub, err := _ElectionManager.contract.WatchLogs(opts, "VoteCasted", nikRule, electionIdRule)
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
// Solidity: event VoteCasted(string indexed nik, string indexed electionId)
func (_ElectionManager *ElectionManagerFilterer) ParseVoteCasted(log types.Log) (*ElectionManagerVoteCasted, error) {
	event := new(ElectionManagerVoteCasted)
	if err := _ElectionManager.contract.UnpackLog(event, "VoteCasted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
