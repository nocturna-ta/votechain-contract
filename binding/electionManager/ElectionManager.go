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
	Bin: "0x608060405234801561000f575f5ffd5b5060405161247038038061247083398181016040528101906100319190610115565b815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050610153565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100e4826100bb565b9050919050565b6100f4816100da565b81146100fe575f5ffd5b50565b5f8151905061010f816100eb565b92915050565b5f5f6040838503121561012b5761012a6100b7565b5b5f61013885828601610101565b925050602061014985828601610101565b9150509250929050565b612310806101605f395ff3fe608060405234801561000f575f5ffd5b50600436106100a7575f3560e01c806395b799071161006f57806395b799071461016557806395f59fe1146101815780639f9070d51461019d578063a9d93bf0146101d0578063df3cf749146101ec578063e030cba01461021c576100a7565b80634771b573146100ab5780634a3ccc96146100db5780635001f3b51461010b5780635597f0831461012957806365910a4f14610147575b5f5ffd5b6100c560048036038101906100c091906115b2565b61024f565b6040516100d29190611706565b60405180910390f35b6100f560048036038101906100f091906115b2565b61041a565b6040516101029190611706565b60405180910390f35b610113610598565b60405161012091906117a0565b60405180910390f35b6101316105bc565b60405161013e91906118db565b60405180910390f35b61014f61075b565b60405161015c919061191b565b60405180910390f35b61017f600480360381019061017a9190611934565b610780565b005b61019b60048036038101906101969190611ada565b610ba6565b005b6101b760048036038101906101b29190611bb1565b610dd0565b6040516101c79493929190611c42565b60405180910390f35b6101ea60048036038101906101e59190611934565b610f23565b005b610206600480360381019061020191906115b2565b611191565b6040516102139190611706565b60405180910390f35b61023660048036038101906102319190611c93565b6113bc565b6040516102469493929190611c42565b60405180910390f35b610257611518565b5f6002848460405161026a929190611d08565b908152602001604051809103902090505f815f01805461028990611d4d565b9050036102c2576040517fcd9f14e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806040518060800160405290815f820180546102dd90611d4d565b80601f016020809104026020016040519081016040528092919081815260200182805461030990611d4d565b80156103545780601f1061032b57610100808354040283529160200191610354565b820191905f5260205f20905b81548152906001019060200180831161033757829003601f168201915b5050505050815260200160018201805461036d90611d4d565b80601f016020809104026020016040519081016040528092919081815260200182805461039990611d4d565b80156103e45780601f106103bb576101008083540402835291602001916103e4565b820191905f5260205f20905b8154815290600101906020018083116103c757829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff16151515158152505091505092915050565b610422611518565b60028383604051610434929190611d08565b90815260200160405180910390206040518060800160405290815f8201805461045c90611d4d565b80601f016020809104026020016040519081016040528092919081815260200182805461048890611d4d565b80156104d35780601f106104aa576101008083540402835291602001916104d3565b820191905f5260205f20905b8154815290600101906020018083116104b657829003601f168201915b505050505081526020016001820180546104ec90611d4d565b80601f016020809104026020016040519081016040528092919081815260200182805461051890611d4d565b80156105635780601f1061053a57610100808354040283529160200191610563565b820191905f5260205f20905b81548152906001019060200180831161054657829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff161515151581525050905092915050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606003805480602002602001604051908101604052809291908181526020015f905b82821015610752578382905f5260205f2090600402016040518060800160405290815f8201805461060f90611d4d565b80601f016020809104026020016040519081016040528092919081815260200182805461063b90611d4d565b80156106865780601f1061065d57610100808354040283529160200191610686565b820191905f5260205f20905b81548152906001019060200180831161066957829003601f168201915b5050505050815260200160018201805461069f90611d4d565b80601f01602080910402602001604051908101604052809291908181526020018280546106cb90611d4d565b80156107165780601f106106ed57610100808354040283529160200191610716565b820191905f5260205f20905b8154815290600101906020018083116106f957829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff161515151581525050815260200190600101906105df565b50505050905090565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fb4ab1646040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107e9573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061080d9190611db8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610871576040517f5c427cd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60028585604051610884929190611d08565b90815260200160405180910390206040518060800160405290815f820180546108ac90611d4d565b80601f01602080910402602001604051908101604052809291908181526020018280546108d890611d4d565b80156109235780601f106108fa57610100808354040283529160200191610923565b820191905f5260205f20905b81548152906001019060200180831161090657829003601f168201915b5050505050815260200160018201805461093c90611d4d565b80601f016020809104026020016040519081016040528092919081815260200182805461096890611d4d565b80156109b35780601f1061098a576101008083540402835291602001916109b3565b820191905f5260205f20905b81548152906001019060200180831161099657829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff16151515158152505090505f815f01515114610a1f576040517fe44dbe0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f604051806080016040528087878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020015f8152602001600115158152509050809150600381908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f019081610b0b9190611f7a565b506020820151816001019081610b219190611f7a565b50604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050508585604051610b5d929190611d08565b60405180910390207ffbb39612bc65a9d2c5c0137dc505f3a2fa0d4c0dac443a06fcf9df0c281c55b28585604051610b96929190612075565b60405180910390a2505050505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663408e27276040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c0f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c3391906120c1565b610c69576040517f5c427cd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60028686604051610c7c929190611d08565b908152602001604051809103902090505f815f018054610c9b90611d4d565b90501480610cb65750806003015f9054906101000a900460ff165b15610ced576040517fcd9f14e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8383604051610cfd92919061211a565b604051809103902081600101604051610d1691906121c4565b604051809103902014610d55576040517f373251d700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806002015f815480929190610d6990612207565b91905055508585604051610d7e929190611d08565b604051809103902082604051610d94919061227e565b60405180910390207f791f7d5f0b0d6e798f239ccca607156ff12f293d6709379301dc1a27e520618160405160405180910390a3505050505050565b60038181548110610ddf575f80fd5b905f5260205f2090600402015f91509050805f018054610dfe90611d4d565b80601f0160208091040260200160405190810160405280929190818152602001828054610e2a90611d4d565b8015610e755780601f10610e4c57610100808354040283529160200191610e75565b820191905f5260205f20905b815481529060010190602001808311610e5857829003601f168201915b505050505090806001018054610e8a90611d4d565b80601f0160208091040260200160405190810160405280929190818152602001828054610eb690611d4d565b8015610f015780601f10610ed857610100808354040283529160200191610f01565b820191905f5260205f20905b815481529060010190602001808311610ee457829003601f168201915b505050505090806002015490806003015f9054906101000a900460ff16905084565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fb4ab1646040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f8c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fb09190611db8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611014576040517f5c427cd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60028585604051611027929190611d08565b908152602001604051809103902090505f815f01805461104690611d4d565b90500361107f576040517fcd9f14e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b828260405161108f92919061211a565b6040518091039020816001016040516110a891906121c4565b6040518091039020146110e7576040517f373251d700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806003015f9054906101000a900460ff1615816003015f6101000a81548160ff0219169083151502179055508282604051611123929190611d08565b6040518091039020858560405161113b929190611d08565b60405180910390207fc4ed90e5395aebcc1e7e027dfc5cafaa02443f04824a16ab3ad2f6a52f8ca06d836003015f9054906101000a900460ff166040516111829190612294565b60405180910390a35050505050565b611199611518565b5f5f90505b6003805490508110156113835783836040516111bb92919061211a565b6040518091039020600382815481106111d7576111d66122ad565b5b905f5260205f2090600402016001016040516111f391906121c4565b6040518091039020036113765760038181548110611214576112136122ad565b5b905f5260205f2090600402016040518060800160405290815f8201805461123a90611d4d565b80601f016020809104026020016040519081016040528092919081815260200182805461126690611d4d565b80156112b15780601f10611288576101008083540402835291602001916112b1565b820191905f5260205f20905b81548152906001019060200180831161129457829003601f168201915b505050505081526020016001820180546112ca90611d4d565b80601f01602080910402602001604051908101604052809291908181526020018280546112f690611d4d565b80156113415780601f1061131857610100808354040283529160200191611341565b820191905f5260205f20905b81548152906001019060200180831161132457829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900460ff1615151515815250509150506113b6565b808060010191505061119e565b506040517fcd9f14e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b92915050565b6002818051602081018201805184825260208301602085012081835280955050505050505f91509050805f0180546113f390611d4d565b80601f016020809104026020016040519081016040528092919081815260200182805461141f90611d4d565b801561146a5780601f106114415761010080835404028352916020019161146a565b820191905f5260205f20905b81548152906001019060200180831161144d57829003601f168201915b50505050509080600101805461147f90611d4d565b80601f01602080910402602001604051908101604052809291908181526020018280546114ab90611d4d565b80156114f65780601f106114cd576101008083540402835291602001916114f6565b820191905f5260205f20905b8154815290600101906020018083116114d957829003601f168201915b505050505090806002015490806003015f9054906101000a900460ff16905084565b604051806080016040528060608152602001606081526020015f81526020015f151581525090565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261157257611571611551565b5b8235905067ffffffffffffffff81111561158f5761158e611555565b5b6020830191508360018202830111156115ab576115aa611559565b5b9250929050565b5f5f602083850312156115c8576115c7611549565b5b5f83013567ffffffffffffffff8111156115e5576115e461154d565b5b6115f18582860161155d565b92509250509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61163f826115fd565b6116498185611607565b9350611659818560208601611617565b61166281611625565b840191505092915050565b5f819050919050565b61167f8161166d565b82525050565b5f8115159050919050565b61169981611685565b82525050565b5f608083015f8301518482035f8601526116b98282611635565b915050602083015184820360208601526116d38282611635565b91505060408301516116e86040860182611676565b5060608301516116fb6060860182611690565b508091505092915050565b5f6020820190508181035f83015261171e818461169f565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f61176861176361175e84611726565b611745565b611726565b9050919050565b5f6117798261174e565b9050919050565b5f61178a8261176f565b9050919050565b61179a81611780565b82525050565b5f6020820190506117b35f830184611791565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f608083015f8301518482035f8601526117fc8282611635565b915050602083015184820360208601526118168282611635565b915050604083015161182b6040860182611676565b50606083015161183e6060860182611690565b508091505092915050565b5f61185483836117e2565b905092915050565b5f602082019050919050565b5f611872826117b9565b61187c81856117c3565b93508360208202850161188e856117d3565b805f5b858110156118c957848403895281516118aa8582611849565b94506118b58361185c565b925060208a01995050600181019050611891565b50829750879550505050505092915050565b5f6020820190508181035f8301526118f38184611868565b905092915050565b5f6119058261176f565b9050919050565b611915816118fb565b82525050565b5f60208201905061192e5f83018461190c565b92915050565b5f5f5f5f6040858703121561194c5761194b611549565b5b5f85013567ffffffffffffffff8111156119695761196861154d565b5b6119758782880161155d565b9450945050602085013567ffffffffffffffff8111156119985761199761154d565b5b6119a48782880161155d565b925092505092959194509250565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6119ec82611625565b810181811067ffffffffffffffff82111715611a0b57611a0a6119b6565b5b80604052505050565b5f611a1d611540565b9050611a2982826119e3565b919050565b5f67ffffffffffffffff821115611a4857611a476119b6565b5b611a5182611625565b9050602081019050919050565b828183375f83830152505050565b5f611a7e611a7984611a2e565b611a14565b905082815260208101848484011115611a9a57611a996119b2565b5b611aa5848285611a5e565b509392505050565b5f82601f830112611ac157611ac0611551565b5b8135611ad1848260208601611a6c565b91505092915050565b5f5f5f5f5f60608688031215611af357611af2611549565b5b5f86013567ffffffffffffffff811115611b1057611b0f61154d565b5b611b1c8882890161155d565b9550955050602086013567ffffffffffffffff811115611b3f57611b3e61154d565b5b611b4b8882890161155d565b9350935050604086013567ffffffffffffffff811115611b6e57611b6d61154d565b5b611b7a88828901611aad565b9150509295509295909350565b611b908161166d565b8114611b9a575f5ffd5b50565b5f81359050611bab81611b87565b92915050565b5f60208284031215611bc657611bc5611549565b5b5f611bd384828501611b9d565b91505092915050565b5f82825260208201905092915050565b5f611bf6826115fd565b611c008185611bdc565b9350611c10818560208601611617565b611c1981611625565b840191505092915050565b611c2d8161166d565b82525050565b611c3c81611685565b82525050565b5f6080820190508181035f830152611c5a8187611bec565b90508181036020830152611c6e8186611bec565b9050611c7d6040830185611c24565b611c8a6060830184611c33565b95945050505050565b5f60208284031215611ca857611ca7611549565b5b5f82013567ffffffffffffffff811115611cc557611cc461154d565b5b611cd184828501611aad565b91505092915050565b5f81905092915050565b5f611cef8385611cda565b9350611cfc838584611a5e565b82840190509392505050565b5f611d14828486611ce4565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611d6457607f821691505b602082108103611d7757611d76611d20565b5b50919050565b5f611d8782611726565b9050919050565b611d9781611d7d565b8114611da1575f5ffd5b50565b5f81519050611db281611d8e565b92915050565b5f60208284031215611dcd57611dcc611549565b5b5f611dda84828501611da4565b91505092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611e3f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611e04565b611e498683611e04565b95508019841693508086168417925050509392505050565b5f611e7b611e76611e718461166d565b611745565b61166d565b9050919050565b5f819050919050565b611e9483611e61565b611ea8611ea082611e82565b848454611e10565b825550505050565b5f5f905090565b611ebf611eb0565b611eca818484611e8b565b505050565b5b81811015611eed57611ee25f82611eb7565b600181019050611ed0565b5050565b601f821115611f3257611f0381611de3565b611f0c84611df5565b81016020851015611f1b578190505b611f2f611f2785611df5565b830182611ecf565b50505b505050565b5f82821c905092915050565b5f611f525f1984600802611f37565b1980831691505092915050565b5f611f6a8383611f43565b9150826002028217905092915050565b611f83826115fd565b67ffffffffffffffff811115611f9c57611f9b6119b6565b5b611fa68254611d4d565b611fb1828285611ef1565b5f60209050601f831160018114611fe2575f8415611fd0578287015190505b611fda8582611f5f565b865550612041565b601f198416611ff086611de3565b5f5b8281101561201757848901518255600182019150602085019450602081019050611ff2565b868310156120345784890151612030601f891682611f43565b8355505b6001600288020188555050505b505050505050565b5f6120548385611bdc565b9350612061838584611a5e565b61206a83611625565b840190509392505050565b5f6020820190508181035f83015261208e818486612049565b90509392505050565b6120a081611685565b81146120aa575f5ffd5b50565b5f815190506120bb81612097565b92915050565b5f602082840312156120d6576120d5611549565b5b5f6120e3848285016120ad565b91505092915050565b5f81905092915050565b5f61210183856120ec565b935061210e838584611a5e565b82840190509392505050565b5f6121268284866120f6565b91508190509392505050565b5f819050815f5260205f209050919050565b5f815461215081611d4d565b61215a81866120ec565b9450600182165f81146121745760018114612189576121bb565b60ff19831686528115158202860193506121bb565b61219285612132565b5f5b838110156121b357815481890152600182019150602081019050612194565b838801955050505b50505092915050565b5f6121cf8284612144565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6122118261166d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612243576122426121da565b5b600182019050919050565b5f612258826115fd565b6122628185611cda565b9350612272818560208601611617565b80840191505092915050565b5f612289828461224e565b915081905092915050565b5f6020820190506122a75f830184611c33565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea26469706673582212208345334f3d2050fd6cc30845012700672d5025b45d7712bfe208b176ebab0bf264736f6c634300081c0033",
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
