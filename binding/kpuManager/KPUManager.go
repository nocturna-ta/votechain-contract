// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kpuManager

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

// IKPUManagerKPUKota is an auto generated low-level Go binding around an user-defined struct.
type IKPUManagerKPUKota struct {
	Name     string
	Address  common.Address
	IsActive bool
	Region   string
}

// IKPUManagerKPUProvinsi is an auto generated low-level Go binding around an user-defined struct.
type IKPUManagerKPUProvinsi struct {
	Name     string
	Address  common.Address
	IsActive bool
	Region   string
}

// KpuManagerMetaData contains all meta data concerning the KpuManager contract.
var KpuManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"KPUKotaAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KPUKotaNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KPUProvinsiAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KPUProvinsiNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyKpuKota\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyKpuProvinsi\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"}],\"name\":\"KPUKotaDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"KPUKotaRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"}],\"name\":\"KPUProvinsiDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"KPUProvinsiRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"contractIVotechainBase\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"}],\"name\":\"deactivateKPUKota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"}],\"name\":\"deactivateKPUProvinsi\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllKPUKota\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"internalType\":\"structIKPUManager.KPUKota[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllKPUProvinsi\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"internalType\":\"structIKPUManager.KPUProvinsi[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"}],\"name\":\"getKpuKotaByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"internalType\":\"structIKPUManager.KPUKota\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kpuAddress\",\"type\":\"address\"}],\"name\":\"getKpuKotaRegion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"}],\"name\":\"getKpuProvinsiByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"internalType\":\"structIKPUManager.KPUProvinsi\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kpuAddress\",\"type\":\"address\"}],\"name\":\"isKPUKota\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kpuAddress\",\"type\":\"address\"}],\"name\":\"isKPUProvinsi\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"kpuKota\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"internalType\":\"structIKPUManager.KPUKota\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kpuKotaAddressesArray\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"kpuProvinsi\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"internalType\":\"structIKPUManager.KPUProvinsi\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kpuProvinsiAddressesArray\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"registerKPUKota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"registerKPUProvinsi\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051612d45380380612d45833981810160405281019061003191906100d4565b805f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a38261007a565b9050919050565b6100b381610099565b81146100bd575f5ffd5b50565b5f815190506100ce816100aa565b92915050565b5f602082840312156100e9576100e8610076565b5b5f6100f6848285016100c0565b91505092915050565b612c398061010c5f395ff3fe608060405234801561000f575f5ffd5b50600436106100fe575f3560e01c806374c8038011610095578063b7a3418e11610064578063b7a3418e146102e8578063cea4146914610306578063df3aae0b14610322578063f9cdd55614610352576100fe565b806374c803801461023c57806375a4352f1461026c578063924238151461029c5780639424ab90146102b8576100fe565b80635001f3b5116100d15780635001f3b51461019d57806356d35b51146101bb578063631a4438146101d95780636fee7d581461020c576100fe565b8063147d559b146101025780632db33a2f14610132578063356b04801461014e5780634bcb1b0f1461016a575b5f5ffd5b61011c60048036038101906101179190612157565b610382565b6040516101299190612282565b60405180910390f35b61014c60048036038101906101479190612303565b610568565b005b61016860048036038101906101639190612157565b610989565b005b610184600480360381019061017f91906123c7565b610b9c565b6040516101949493929190612458565b60405180910390f35b6101a5610d0f565b6040516101b29190612504565b60405180910390f35b6101c3610d33565b6040516101d0919061263f565b60405180910390f35b6101f360048036038101906101ee91906123c7565b610f1e565b6040516102039493929190612458565b60405180910390f35b61022660048036038101906102219190612157565b611091565b6040516102339190612282565b60405180910390f35b61025660048036038101906102519190612157565b6112fb565b604051610263919061265f565b60405180910390f35b61028660048036038101906102819190612157565b6113cb565b604051610293919061267f565b60405180910390f35b6102b660048036038101906102b19190612303565b611421565b005b6102d260048036038101906102cd9190612157565b6117d5565b6040516102df91906126ff565b60405180910390f35b6102f0611a3f565b6040516102fd9190612841565b60405180910390f35b610320600480360381019061031b9190612157565b611c2a565b005b61033c60048036038101906103379190612157565b611e3d565b604051610349919061267f565b60405180910390f35b61036c60048036038101906103679190612157565b611e93565b60405161037991906126ff565b60405180910390f35b61038a612079565b60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060800160405290815f820180546103e19061288e565b80601f016020809104026020016040519081016040528092919081815260200182805461040d9061288e565b80156104585780601f1061042f57610100808354040283529160200191610458565b820191905f5260205f20905b81548152906001019060200180831161043b57829003601f168201915b50505050508152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900460ff161515151581526020016002820180546104e19061288e565b80601f016020809104026020016040519081016040528092919081815260200182805461050d9061288e565b80156105585780601f1061052f57610100808354040283529160200191610558565b820191905f5260205f20905b81548152906001019060200180831161053b57829003601f168201915b5050505050815250509050919050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fb4ab1646040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105d1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105f591906128d2565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610659576040517ff75c676f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060010160149054906101000a900460ff16156106de576040517f6fed5b9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f604051806080016040528086868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018773ffffffffffffffffffffffffffffffffffffffff16815260200160011515815260200184848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081525090508060015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f0190816107f29190612ac1565b506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548160ff021916908315150217905550606082015181600201908161086e9190612ac1565b50905050600381908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f0190816108ae9190612ac1565b506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548160ff021916908315150217905550606082015181600201908161092a9190612ac1565b5050508573ffffffffffffffffffffffffffffffffffffffff167fe253ed8255fe5723a22412bf5804ee9034911e408cdbfe8547438f23ef231035868686866040516109799493929190612bca565b60405180910390a2505050505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fb4ab1646040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109f2573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a1691906128d2565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a7a576040517ff75c676f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60015f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060010160149054906101000a900460ff16610afe576040517f27b434e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060010160146101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167f9821f051eb616938dcfef57fac817566f3954f1ed18aa52b2916921cb2bacc6960405160405180910390a250565b60048181548110610bab575f80fd5b905f5260205f2090600302015f91509050805f018054610bca9061288e565b80601f0160208091040260200160405190810160405280929190818152602001828054610bf69061288e565b8015610c415780601f10610c1857610100808354040283529160200191610c41565b820191905f5260205f20905b815481529060010190602001808311610c2457829003601f168201915b505050505090806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160149054906101000a900460ff1690806002018054610c8e9061288e565b80601f0160208091040260200160405190810160405280929190818152602001828054610cba9061288e565b8015610d055780601f10610cdc57610100808354040283529160200191610d05565b820191905f5260205f20905b815481529060010190602001808311610ce857829003601f168201915b5050505050905084565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606003805480602002602001604051908101604052809291908181526020015f905b82821015610f15578382905f5260205f2090600302016040518060800160405290815f82018054610d869061288e565b80601f0160208091040260200160405190810160405280929190818152602001828054610db29061288e565b8015610dfd5780601f10610dd457610100808354040283529160200191610dfd565b820191905f5260205f20905b815481529060010190602001808311610de057829003601f168201915b50505050508152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900460ff16151515158152602001600282018054610e869061288e565b80601f0160208091040260200160405190810160405280929190818152602001828054610eb29061288e565b8015610efd5780601f10610ed457610100808354040283529160200191610efd565b820191905f5260205f20905b815481529060010190602001808311610ee057829003601f168201915b50505050508152505081526020019060010190610d56565b50505050905090565b60038181548110610f2d575f80fd5b905f5260205f2090600302015f91509050805f018054610f4c9061288e565b80601f0160208091040260200160405190810160405280929190818152602001828054610f789061288e565b8015610fc35780601f10610f9a57610100808354040283529160200191610fc3565b820191905f5260205f20905b815481529060010190602001808311610fa657829003601f168201915b505050505090806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160149054906101000a900460ff16908060020180546110109061288e565b80601f016020809104026020016040519081016040528092919081815260200182805461103c9061288e565b80156110875780601f1061105e57610100808354040283529160200191611087565b820191905f5260205f20905b81548152906001019060200180831161106a57829003601f168201915b5050505050905084565b611099612079565b60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060010160149054906101000a900460ff1661111d576040517f8789d78900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060800160405290815f820180546111749061288e565b80601f01602080910402602001604051908101604052809291908181526020018280546111a09061288e565b80156111eb5780601f106111c2576101008083540402835291602001916111eb565b820191905f5260205f20905b8154815290600101906020018083116111ce57829003601f168201915b50505050508152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900460ff161515151581526020016002820180546112749061288e565b80601f01602080910402602001604051908101604052809291908181526020018280546112a09061288e565b80156112eb5780601f106112c2576101008083540402835291602001916112eb565b820191905f5260205f20905b8154815290600101906020018083116112ce57829003601f168201915b5050505050815250509050919050565b606060025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060020180546113489061288e565b80601f01602080910402602001604051908101604052809291908181526020018280546113749061288e565b80156113bf5780601f10611396576101008083540402835291602001916113bf565b820191905f5260205f20905b8154815290600101906020018083116113a257829003601f168201915b50505050509050919050565b5f60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060010160149054906101000a900460ff169050919050565b60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060010160149054906101000a900460ff166114a5576040517ff75c676f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060010160149054906101000a900460ff161561152a576040517f4f576a3200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f604051806080016040528086868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018773ffffffffffffffffffffffffffffffffffffffff16815260200160011515815260200184848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081525090508060025f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f01908161163e9190612ac1565b506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548160ff02191690831515021790555060608201518160020190816116ba9190612ac1565b50905050600481908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f0190816116fa9190612ac1565b506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548160ff02191690831515021790555060608201518160020190816117769190612ac1565b5050508573ffffffffffffffffffffffffffffffffffffffff167f9c6f19aa0cf3f6a0d4fdbd9269f4a794835bf59c3c7561af88f025740ef3cb77868686866040516117c59493929190612bca565b60405180910390a2505050505050565b6117dd6120b7565b60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060010160149054906101000a900460ff16611861576040517f27b434e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060800160405290815f820180546118b89061288e565b80601f01602080910402602001604051908101604052809291908181526020018280546118e49061288e565b801561192f5780601f106119065761010080835404028352916020019161192f565b820191905f5260205f20905b81548152906001019060200180831161191257829003601f168201915b50505050508152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900460ff161515151581526020016002820180546119b89061288e565b80601f01602080910402602001604051908101604052809291908181526020018280546119e49061288e565b8015611a2f5780601f10611a0657610100808354040283529160200191611a2f565b820191905f5260205f20905b815481529060010190602001808311611a1257829003601f168201915b5050505050815250509050919050565b60606004805480602002602001604051908101604052809291908181526020015f905b82821015611c21578382905f5260205f2090600302016040518060800160405290815f82018054611a929061288e565b80601f0160208091040260200160405190810160405280929190818152602001828054611abe9061288e565b8015611b095780601f10611ae057610100808354040283529160200191611b09565b820191905f5260205f20905b815481529060010190602001808311611aec57829003601f168201915b50505050508152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900460ff16151515158152602001600282018054611b929061288e565b80601f0160208091040260200160405190810160405280929190818152602001828054611bbe9061288e565b8015611c095780601f10611be057610100808354040283529160200191611c09565b820191905f5260205f20905b815481529060010190602001808311611bec57829003601f168201915b50505050508152505081526020019060010190611a62565b50505050905090565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fb4ab1646040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c93573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611cb791906128d2565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611d1b576040517ff75c676f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060010160149054906101000a900460ff16611d9f576040517f8789d78900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060010160146101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167f4287f31e2e9e49d47750bbe8bcdc9848644c0566e597ce529738d63b27b7b99260405160405180910390a250565b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060010160149054906101000a900460ff169050919050565b611e9b6120b7565b60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060800160405290815f82018054611ef29061288e565b80601f0160208091040260200160405190810160405280929190818152602001828054611f1e9061288e565b8015611f695780601f10611f4057610100808354040283529160200191611f69565b820191905f5260205f20905b815481529060010190602001808311611f4c57829003601f168201915b50505050508152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900460ff16151515158152602001600282018054611ff29061288e565b80601f016020809104026020016040519081016040528092919081815260200182805461201e9061288e565b80156120695780601f1061204057610100808354040283529160200191612069565b820191905f5260205f20905b81548152906001019060200180831161204c57829003601f168201915b5050505050815250509050919050565b6040518060800160405280606081526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f15158152602001606081525090565b6040518060800160405280606081526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f15158152602001606081525090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f612126826120fd565b9050919050565b6121368161211c565b8114612140575f5ffd5b50565b5f813590506121518161212d565b92915050565b5f6020828403121561216c5761216b6120f5565b5b5f61217984828501612143565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6121c482612182565b6121ce818561218c565b93506121de81856020860161219c565b6121e7816121aa565b840191505092915050565b6121fb8161211c565b82525050565b5f8115159050919050565b61221581612201565b82525050565b5f608083015f8301518482035f86015261223582826121ba565b915050602083015161224a60208601826121f2565b50604083015161225d604086018261220c565b506060830151848203606086015261227582826121ba565b9150508091505092915050565b5f6020820190508181035f83015261229a818461221b565b905092915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126122c3576122c26122a2565b5b8235905067ffffffffffffffff8111156122e0576122df6122a6565b5b6020830191508360018202830111156122fc576122fb6122aa565b5b9250929050565b5f5f5f5f5f6060868803121561231c5761231b6120f5565b5b5f61232988828901612143565b955050602086013567ffffffffffffffff81111561234a576123496120f9565b5b612356888289016122ae565b9450945050604086013567ffffffffffffffff811115612379576123786120f9565b5b612385888289016122ae565b92509250509295509295909350565b5f819050919050565b6123a681612394565b81146123b0575f5ffd5b50565b5f813590506123c18161239d565b92915050565b5f602082840312156123dc576123db6120f5565b5b5f6123e9848285016123b3565b91505092915050565b5f82825260208201905092915050565b5f61240c82612182565b61241681856123f2565b935061242681856020860161219c565b61242f816121aa565b840191505092915050565b6124438161211c565b82525050565b61245281612201565b82525050565b5f6080820190508181035f8301526124708187612402565b905061247f602083018661243a565b61248c6040830185612449565b818103606083015261249e8184612402565b905095945050505050565b5f819050919050565b5f6124cc6124c76124c2846120fd565b6124a9565b6120fd565b9050919050565b5f6124dd826124b2565b9050919050565b5f6124ee826124d3565b9050919050565b6124fe816124e4565b82525050565b5f6020820190506125175f8301846124f5565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f608083015f8301518482035f86015261256082826121ba565b915050602083015161257560208601826121f2565b506040830151612588604086018261220c565b50606083015184820360608601526125a082826121ba565b9150508091505092915050565b5f6125b88383612546565b905092915050565b5f602082019050919050565b5f6125d68261251d565b6125e08185612527565b9350836020820285016125f285612537565b805f5b8581101561262d578484038952815161260e85826125ad565b9450612619836125c0565b925060208a019950506001810190506125f5565b50829750879550505050505092915050565b5f6020820190508181035f83015261265781846125cc565b905092915050565b5f6020820190508181035f8301526126778184612402565b905092915050565b5f6020820190506126925f830184612449565b92915050565b5f608083015f8301518482035f8601526126b282826121ba565b91505060208301516126c760208601826121f2565b5060408301516126da604086018261220c565b50606083015184820360608601526126f282826121ba565b9150508091505092915050565b5f6020820190508181035f8301526127178184612698565b905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f608083015f8301518482035f86015261276282826121ba565b915050602083015161277760208601826121f2565b50604083015161278a604086018261220c565b50606083015184820360608601526127a282826121ba565b9150508091505092915050565b5f6127ba8383612748565b905092915050565b5f602082019050919050565b5f6127d88261271f565b6127e28185612729565b9350836020820285016127f485612739565b805f5b8581101561282f578484038952815161281085826127af565b945061281b836127c2565b925060208a019950506001810190506127f7565b50829750879550505050505092915050565b5f6020820190508181035f83015261285981846127ce565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806128a557607f821691505b6020821081036128b8576128b7612861565b5b50919050565b5f815190506128cc8161212d565b92915050565b5f602082840312156128e7576128e66120f5565b5b5f6128f4848285016128be565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026129867fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261294b565b612990868361294b565b95508019841693508086168417925050509392505050565b5f6129c26129bd6129b884612394565b6124a9565b612394565b9050919050565b5f819050919050565b6129db836129a8565b6129ef6129e7826129c9565b848454612957565b825550505050565b5f5f905090565b612a066129f7565b612a118184846129d2565b505050565b5b81811015612a3457612a295f826129fe565b600181019050612a17565b5050565b601f821115612a7957612a4a8161292a565b612a538461293c565b81016020851015612a62578190505b612a76612a6e8561293c565b830182612a16565b50505b505050565b5f82821c905092915050565b5f612a995f1984600802612a7e565b1980831691505092915050565b5f612ab18383612a8a565b9150826002028217905092915050565b612aca82612182565b67ffffffffffffffff811115612ae357612ae26128fd565b5b612aed825461288e565b612af8828285612a38565b5f60209050601f831160018114612b29575f8415612b17578287015190505b612b218582612aa6565b865550612b88565b601f198416612b378661292a565b5f5b82811015612b5e57848901518255600182019150602085019450602081019050612b39565b86831015612b7b5784890151612b77601f891682612a8a565b8355505b6001600288020188555050505b505050505050565b828183375f83830152505050565b5f612ba983856123f2565b9350612bb6838584612b90565b612bbf836121aa565b840190509392505050565b5f6040820190508181035f830152612be3818688612b9e565b90508181036020830152612bf8818486612b9e565b90509594505050505056fea264697066735822122053cd89f873c52b53b457b76591571621d7d170801a3be77652825f41b0341d1264736f6c634300081c0033",
}

// KpuManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use KpuManagerMetaData.ABI instead.
var KpuManagerABI = KpuManagerMetaData.ABI

// KpuManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KpuManagerMetaData.Bin instead.
var KpuManagerBin = KpuManagerMetaData.Bin

// DeployKpuManager deploys a new Ethereum contract, binding an instance of KpuManager to it.
func DeployKpuManager(auth *bind.TransactOpts, backend bind.ContractBackend, _baseAddress common.Address) (common.Address, *types.Transaction, *KpuManager, error) {
	parsed, err := KpuManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KpuManagerBin), backend, _baseAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KpuManager{KpuManagerCaller: KpuManagerCaller{contract: contract}, KpuManagerTransactor: KpuManagerTransactor{contract: contract}, KpuManagerFilterer: KpuManagerFilterer{contract: contract}}, nil
}

// KpuManager is an auto generated Go binding around an Ethereum contract.
type KpuManager struct {
	KpuManagerCaller     // Read-only binding to the contract
	KpuManagerTransactor // Write-only binding to the contract
	KpuManagerFilterer   // Log filterer for contract events
}

// KpuManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type KpuManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KpuManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KpuManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KpuManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KpuManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KpuManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KpuManagerSession struct {
	Contract     *KpuManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KpuManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KpuManagerCallerSession struct {
	Contract *KpuManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// KpuManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KpuManagerTransactorSession struct {
	Contract     *KpuManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// KpuManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type KpuManagerRaw struct {
	Contract *KpuManager // Generic contract binding to access the raw methods on
}

// KpuManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KpuManagerCallerRaw struct {
	Contract *KpuManagerCaller // Generic read-only contract binding to access the raw methods on
}

// KpuManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KpuManagerTransactorRaw struct {
	Contract *KpuManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKpuManager creates a new instance of KpuManager, bound to a specific deployed contract.
func NewKpuManager(address common.Address, backend bind.ContractBackend) (*KpuManager, error) {
	contract, err := bindKpuManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KpuManager{KpuManagerCaller: KpuManagerCaller{contract: contract}, KpuManagerTransactor: KpuManagerTransactor{contract: contract}, KpuManagerFilterer: KpuManagerFilterer{contract: contract}}, nil
}

// NewKpuManagerCaller creates a new read-only instance of KpuManager, bound to a specific deployed contract.
func NewKpuManagerCaller(address common.Address, caller bind.ContractCaller) (*KpuManagerCaller, error) {
	contract, err := bindKpuManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KpuManagerCaller{contract: contract}, nil
}

// NewKpuManagerTransactor creates a new write-only instance of KpuManager, bound to a specific deployed contract.
func NewKpuManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*KpuManagerTransactor, error) {
	contract, err := bindKpuManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KpuManagerTransactor{contract: contract}, nil
}

// NewKpuManagerFilterer creates a new log filterer instance of KpuManager, bound to a specific deployed contract.
func NewKpuManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*KpuManagerFilterer, error) {
	contract, err := bindKpuManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KpuManagerFilterer{contract: contract}, nil
}

// bindKpuManager binds a generic wrapper to an already deployed contract.
func bindKpuManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KpuManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KpuManager *KpuManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KpuManager.Contract.KpuManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KpuManager *KpuManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KpuManager.Contract.KpuManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KpuManager *KpuManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KpuManager.Contract.KpuManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KpuManager *KpuManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KpuManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KpuManager *KpuManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KpuManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KpuManager *KpuManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KpuManager.Contract.contract.Transact(opts, method, params...)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_KpuManager *KpuManagerCaller) Base(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KpuManager.contract.Call(opts, &out, "base")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_KpuManager *KpuManagerSession) Base() (common.Address, error) {
	return _KpuManager.Contract.Base(&_KpuManager.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_KpuManager *KpuManagerCallerSession) Base() (common.Address, error) {
	return _KpuManager.Contract.Base(&_KpuManager.CallOpts)
}

// GetAllKPUKota is a free data retrieval call binding the contract method 0xb7a3418e.
//
// Solidity: function getAllKPUKota() view returns((string,address,bool,string)[])
func (_KpuManager *KpuManagerCaller) GetAllKPUKota(opts *bind.CallOpts) ([]IKPUManagerKPUKota, error) {
	var out []interface{}
	err := _KpuManager.contract.Call(opts, &out, "getAllKPUKota")

	if err != nil {
		return *new([]IKPUManagerKPUKota), err
	}

	out0 := *abi.ConvertType(out[0], new([]IKPUManagerKPUKota)).(*[]IKPUManagerKPUKota)

	return out0, err

}

// GetAllKPUKota is a free data retrieval call binding the contract method 0xb7a3418e.
//
// Solidity: function getAllKPUKota() view returns((string,address,bool,string)[])
func (_KpuManager *KpuManagerSession) GetAllKPUKota() ([]IKPUManagerKPUKota, error) {
	return _KpuManager.Contract.GetAllKPUKota(&_KpuManager.CallOpts)
}

// GetAllKPUKota is a free data retrieval call binding the contract method 0xb7a3418e.
//
// Solidity: function getAllKPUKota() view returns((string,address,bool,string)[])
func (_KpuManager *KpuManagerCallerSession) GetAllKPUKota() ([]IKPUManagerKPUKota, error) {
	return _KpuManager.Contract.GetAllKPUKota(&_KpuManager.CallOpts)
}

// GetAllKPUProvinsi is a free data retrieval call binding the contract method 0x56d35b51.
//
// Solidity: function getAllKPUProvinsi() view returns((string,address,bool,string)[])
func (_KpuManager *KpuManagerCaller) GetAllKPUProvinsi(opts *bind.CallOpts) ([]IKPUManagerKPUProvinsi, error) {
	var out []interface{}
	err := _KpuManager.contract.Call(opts, &out, "getAllKPUProvinsi")

	if err != nil {
		return *new([]IKPUManagerKPUProvinsi), err
	}

	out0 := *abi.ConvertType(out[0], new([]IKPUManagerKPUProvinsi)).(*[]IKPUManagerKPUProvinsi)

	return out0, err

}

// GetAllKPUProvinsi is a free data retrieval call binding the contract method 0x56d35b51.
//
// Solidity: function getAllKPUProvinsi() view returns((string,address,bool,string)[])
func (_KpuManager *KpuManagerSession) GetAllKPUProvinsi() ([]IKPUManagerKPUProvinsi, error) {
	return _KpuManager.Contract.GetAllKPUProvinsi(&_KpuManager.CallOpts)
}

// GetAllKPUProvinsi is a free data retrieval call binding the contract method 0x56d35b51.
//
// Solidity: function getAllKPUProvinsi() view returns((string,address,bool,string)[])
func (_KpuManager *KpuManagerCallerSession) GetAllKPUProvinsi() ([]IKPUManagerKPUProvinsi, error) {
	return _KpuManager.Contract.GetAllKPUProvinsi(&_KpuManager.CallOpts)
}

// GetKpuKotaByAddress is a free data retrieval call binding the contract method 0x6fee7d58.
//
// Solidity: function getKpuKotaByAddress(address Address) view returns((string,address,bool,string))
func (_KpuManager *KpuManagerCaller) GetKpuKotaByAddress(opts *bind.CallOpts, Address common.Address) (IKPUManagerKPUKota, error) {
	var out []interface{}
	err := _KpuManager.contract.Call(opts, &out, "getKpuKotaByAddress", Address)

	if err != nil {
		return *new(IKPUManagerKPUKota), err
	}

	out0 := *abi.ConvertType(out[0], new(IKPUManagerKPUKota)).(*IKPUManagerKPUKota)

	return out0, err

}

// GetKpuKotaByAddress is a free data retrieval call binding the contract method 0x6fee7d58.
//
// Solidity: function getKpuKotaByAddress(address Address) view returns((string,address,bool,string))
func (_KpuManager *KpuManagerSession) GetKpuKotaByAddress(Address common.Address) (IKPUManagerKPUKota, error) {
	return _KpuManager.Contract.GetKpuKotaByAddress(&_KpuManager.CallOpts, Address)
}

// GetKpuKotaByAddress is a free data retrieval call binding the contract method 0x6fee7d58.
//
// Solidity: function getKpuKotaByAddress(address Address) view returns((string,address,bool,string))
func (_KpuManager *KpuManagerCallerSession) GetKpuKotaByAddress(Address common.Address) (IKPUManagerKPUKota, error) {
	return _KpuManager.Contract.GetKpuKotaByAddress(&_KpuManager.CallOpts, Address)
}

// GetKpuKotaRegion is a free data retrieval call binding the contract method 0x74c80380.
//
// Solidity: function getKpuKotaRegion(address kpuAddress) view returns(string)
func (_KpuManager *KpuManagerCaller) GetKpuKotaRegion(opts *bind.CallOpts, kpuAddress common.Address) (string, error) {
	var out []interface{}
	err := _KpuManager.contract.Call(opts, &out, "getKpuKotaRegion", kpuAddress)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetKpuKotaRegion is a free data retrieval call binding the contract method 0x74c80380.
//
// Solidity: function getKpuKotaRegion(address kpuAddress) view returns(string)
func (_KpuManager *KpuManagerSession) GetKpuKotaRegion(kpuAddress common.Address) (string, error) {
	return _KpuManager.Contract.GetKpuKotaRegion(&_KpuManager.CallOpts, kpuAddress)
}

// GetKpuKotaRegion is a free data retrieval call binding the contract method 0x74c80380.
//
// Solidity: function getKpuKotaRegion(address kpuAddress) view returns(string)
func (_KpuManager *KpuManagerCallerSession) GetKpuKotaRegion(kpuAddress common.Address) (string, error) {
	return _KpuManager.Contract.GetKpuKotaRegion(&_KpuManager.CallOpts, kpuAddress)
}

// GetKpuProvinsiByAddress is a free data retrieval call binding the contract method 0x9424ab90.
//
// Solidity: function getKpuProvinsiByAddress(address Address) view returns((string,address,bool,string))
func (_KpuManager *KpuManagerCaller) GetKpuProvinsiByAddress(opts *bind.CallOpts, Address common.Address) (IKPUManagerKPUProvinsi, error) {
	var out []interface{}
	err := _KpuManager.contract.Call(opts, &out, "getKpuProvinsiByAddress", Address)

	if err != nil {
		return *new(IKPUManagerKPUProvinsi), err
	}

	out0 := *abi.ConvertType(out[0], new(IKPUManagerKPUProvinsi)).(*IKPUManagerKPUProvinsi)

	return out0, err

}

// GetKpuProvinsiByAddress is a free data retrieval call binding the contract method 0x9424ab90.
//
// Solidity: function getKpuProvinsiByAddress(address Address) view returns((string,address,bool,string))
func (_KpuManager *KpuManagerSession) GetKpuProvinsiByAddress(Address common.Address) (IKPUManagerKPUProvinsi, error) {
	return _KpuManager.Contract.GetKpuProvinsiByAddress(&_KpuManager.CallOpts, Address)
}

// GetKpuProvinsiByAddress is a free data retrieval call binding the contract method 0x9424ab90.
//
// Solidity: function getKpuProvinsiByAddress(address Address) view returns((string,address,bool,string))
func (_KpuManager *KpuManagerCallerSession) GetKpuProvinsiByAddress(Address common.Address) (IKPUManagerKPUProvinsi, error) {
	return _KpuManager.Contract.GetKpuProvinsiByAddress(&_KpuManager.CallOpts, Address)
}

// IsKPUKota is a free data retrieval call binding the contract method 0x75a4352f.
//
// Solidity: function isKPUKota(address kpuAddress) view returns(bool)
func (_KpuManager *KpuManagerCaller) IsKPUKota(opts *bind.CallOpts, kpuAddress common.Address) (bool, error) {
	var out []interface{}
	err := _KpuManager.contract.Call(opts, &out, "isKPUKota", kpuAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKPUKota is a free data retrieval call binding the contract method 0x75a4352f.
//
// Solidity: function isKPUKota(address kpuAddress) view returns(bool)
func (_KpuManager *KpuManagerSession) IsKPUKota(kpuAddress common.Address) (bool, error) {
	return _KpuManager.Contract.IsKPUKota(&_KpuManager.CallOpts, kpuAddress)
}

// IsKPUKota is a free data retrieval call binding the contract method 0x75a4352f.
//
// Solidity: function isKPUKota(address kpuAddress) view returns(bool)
func (_KpuManager *KpuManagerCallerSession) IsKPUKota(kpuAddress common.Address) (bool, error) {
	return _KpuManager.Contract.IsKPUKota(&_KpuManager.CallOpts, kpuAddress)
}

// IsKPUProvinsi is a free data retrieval call binding the contract method 0xdf3aae0b.
//
// Solidity: function isKPUProvinsi(address kpuAddress) view returns(bool)
func (_KpuManager *KpuManagerCaller) IsKPUProvinsi(opts *bind.CallOpts, kpuAddress common.Address) (bool, error) {
	var out []interface{}
	err := _KpuManager.contract.Call(opts, &out, "isKPUProvinsi", kpuAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKPUProvinsi is a free data retrieval call binding the contract method 0xdf3aae0b.
//
// Solidity: function isKPUProvinsi(address kpuAddress) view returns(bool)
func (_KpuManager *KpuManagerSession) IsKPUProvinsi(kpuAddress common.Address) (bool, error) {
	return _KpuManager.Contract.IsKPUProvinsi(&_KpuManager.CallOpts, kpuAddress)
}

// IsKPUProvinsi is a free data retrieval call binding the contract method 0xdf3aae0b.
//
// Solidity: function isKPUProvinsi(address kpuAddress) view returns(bool)
func (_KpuManager *KpuManagerCallerSession) IsKPUProvinsi(kpuAddress common.Address) (bool, error) {
	return _KpuManager.Contract.IsKPUProvinsi(&_KpuManager.CallOpts, kpuAddress)
}

// KpuKota is a free data retrieval call binding the contract method 0x147d559b.
//
// Solidity: function kpuKota(address addr) view returns((string,address,bool,string))
func (_KpuManager *KpuManagerCaller) KpuKota(opts *bind.CallOpts, addr common.Address) (IKPUManagerKPUKota, error) {
	var out []interface{}
	err := _KpuManager.contract.Call(opts, &out, "kpuKota", addr)

	if err != nil {
		return *new(IKPUManagerKPUKota), err
	}

	out0 := *abi.ConvertType(out[0], new(IKPUManagerKPUKota)).(*IKPUManagerKPUKota)

	return out0, err

}

// KpuKota is a free data retrieval call binding the contract method 0x147d559b.
//
// Solidity: function kpuKota(address addr) view returns((string,address,bool,string))
func (_KpuManager *KpuManagerSession) KpuKota(addr common.Address) (IKPUManagerKPUKota, error) {
	return _KpuManager.Contract.KpuKota(&_KpuManager.CallOpts, addr)
}

// KpuKota is a free data retrieval call binding the contract method 0x147d559b.
//
// Solidity: function kpuKota(address addr) view returns((string,address,bool,string))
func (_KpuManager *KpuManagerCallerSession) KpuKota(addr common.Address) (IKPUManagerKPUKota, error) {
	return _KpuManager.Contract.KpuKota(&_KpuManager.CallOpts, addr)
}

// KpuKotaAddressesArray is a free data retrieval call binding the contract method 0x4bcb1b0f.
//
// Solidity: function kpuKotaAddressesArray(uint256 ) view returns(string name, address Address, bool isActive, string region)
func (_KpuManager *KpuManagerCaller) KpuKotaAddressesArray(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name     string
	Address  common.Address
	IsActive bool
	Region   string
}, error) {
	var out []interface{}
	err := _KpuManager.contract.Call(opts, &out, "kpuKotaAddressesArray", arg0)

	outstruct := new(struct {
		Name     string
		Address  common.Address
		IsActive bool
		Region   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Address = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.IsActive = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Region = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// KpuKotaAddressesArray is a free data retrieval call binding the contract method 0x4bcb1b0f.
//
// Solidity: function kpuKotaAddressesArray(uint256 ) view returns(string name, address Address, bool isActive, string region)
func (_KpuManager *KpuManagerSession) KpuKotaAddressesArray(arg0 *big.Int) (struct {
	Name     string
	Address  common.Address
	IsActive bool
	Region   string
}, error) {
	return _KpuManager.Contract.KpuKotaAddressesArray(&_KpuManager.CallOpts, arg0)
}

// KpuKotaAddressesArray is a free data retrieval call binding the contract method 0x4bcb1b0f.
//
// Solidity: function kpuKotaAddressesArray(uint256 ) view returns(string name, address Address, bool isActive, string region)
func (_KpuManager *KpuManagerCallerSession) KpuKotaAddressesArray(arg0 *big.Int) (struct {
	Name     string
	Address  common.Address
	IsActive bool
	Region   string
}, error) {
	return _KpuManager.Contract.KpuKotaAddressesArray(&_KpuManager.CallOpts, arg0)
}

// KpuProvinsi is a free data retrieval call binding the contract method 0xf9cdd556.
//
// Solidity: function kpuProvinsi(address addr) view returns((string,address,bool,string))
func (_KpuManager *KpuManagerCaller) KpuProvinsi(opts *bind.CallOpts, addr common.Address) (IKPUManagerKPUProvinsi, error) {
	var out []interface{}
	err := _KpuManager.contract.Call(opts, &out, "kpuProvinsi", addr)

	if err != nil {
		return *new(IKPUManagerKPUProvinsi), err
	}

	out0 := *abi.ConvertType(out[0], new(IKPUManagerKPUProvinsi)).(*IKPUManagerKPUProvinsi)

	return out0, err

}

// KpuProvinsi is a free data retrieval call binding the contract method 0xf9cdd556.
//
// Solidity: function kpuProvinsi(address addr) view returns((string,address,bool,string))
func (_KpuManager *KpuManagerSession) KpuProvinsi(addr common.Address) (IKPUManagerKPUProvinsi, error) {
	return _KpuManager.Contract.KpuProvinsi(&_KpuManager.CallOpts, addr)
}

// KpuProvinsi is a free data retrieval call binding the contract method 0xf9cdd556.
//
// Solidity: function kpuProvinsi(address addr) view returns((string,address,bool,string))
func (_KpuManager *KpuManagerCallerSession) KpuProvinsi(addr common.Address) (IKPUManagerKPUProvinsi, error) {
	return _KpuManager.Contract.KpuProvinsi(&_KpuManager.CallOpts, addr)
}

// KpuProvinsiAddressesArray is a free data retrieval call binding the contract method 0x631a4438.
//
// Solidity: function kpuProvinsiAddressesArray(uint256 ) view returns(string name, address Address, bool isActive, string region)
func (_KpuManager *KpuManagerCaller) KpuProvinsiAddressesArray(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name     string
	Address  common.Address
	IsActive bool
	Region   string
}, error) {
	var out []interface{}
	err := _KpuManager.contract.Call(opts, &out, "kpuProvinsiAddressesArray", arg0)

	outstruct := new(struct {
		Name     string
		Address  common.Address
		IsActive bool
		Region   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Address = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.IsActive = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Region = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// KpuProvinsiAddressesArray is a free data retrieval call binding the contract method 0x631a4438.
//
// Solidity: function kpuProvinsiAddressesArray(uint256 ) view returns(string name, address Address, bool isActive, string region)
func (_KpuManager *KpuManagerSession) KpuProvinsiAddressesArray(arg0 *big.Int) (struct {
	Name     string
	Address  common.Address
	IsActive bool
	Region   string
}, error) {
	return _KpuManager.Contract.KpuProvinsiAddressesArray(&_KpuManager.CallOpts, arg0)
}

// KpuProvinsiAddressesArray is a free data retrieval call binding the contract method 0x631a4438.
//
// Solidity: function kpuProvinsiAddressesArray(uint256 ) view returns(string name, address Address, bool isActive, string region)
func (_KpuManager *KpuManagerCallerSession) KpuProvinsiAddressesArray(arg0 *big.Int) (struct {
	Name     string
	Address  common.Address
	IsActive bool
	Region   string
}, error) {
	return _KpuManager.Contract.KpuProvinsiAddressesArray(&_KpuManager.CallOpts, arg0)
}

// DeactivateKPUKota is a paid mutator transaction binding the contract method 0xcea41469.
//
// Solidity: function deactivateKPUKota(address Address) returns()
func (_KpuManager *KpuManagerTransactor) DeactivateKPUKota(opts *bind.TransactOpts, Address common.Address) (*types.Transaction, error) {
	return _KpuManager.contract.Transact(opts, "deactivateKPUKota", Address)
}

// DeactivateKPUKota is a paid mutator transaction binding the contract method 0xcea41469.
//
// Solidity: function deactivateKPUKota(address Address) returns()
func (_KpuManager *KpuManagerSession) DeactivateKPUKota(Address common.Address) (*types.Transaction, error) {
	return _KpuManager.Contract.DeactivateKPUKota(&_KpuManager.TransactOpts, Address)
}

// DeactivateKPUKota is a paid mutator transaction binding the contract method 0xcea41469.
//
// Solidity: function deactivateKPUKota(address Address) returns()
func (_KpuManager *KpuManagerTransactorSession) DeactivateKPUKota(Address common.Address) (*types.Transaction, error) {
	return _KpuManager.Contract.DeactivateKPUKota(&_KpuManager.TransactOpts, Address)
}

// DeactivateKPUProvinsi is a paid mutator transaction binding the contract method 0x356b0480.
//
// Solidity: function deactivateKPUProvinsi(address Address) returns()
func (_KpuManager *KpuManagerTransactor) DeactivateKPUProvinsi(opts *bind.TransactOpts, Address common.Address) (*types.Transaction, error) {
	return _KpuManager.contract.Transact(opts, "deactivateKPUProvinsi", Address)
}

// DeactivateKPUProvinsi is a paid mutator transaction binding the contract method 0x356b0480.
//
// Solidity: function deactivateKPUProvinsi(address Address) returns()
func (_KpuManager *KpuManagerSession) DeactivateKPUProvinsi(Address common.Address) (*types.Transaction, error) {
	return _KpuManager.Contract.DeactivateKPUProvinsi(&_KpuManager.TransactOpts, Address)
}

// DeactivateKPUProvinsi is a paid mutator transaction binding the contract method 0x356b0480.
//
// Solidity: function deactivateKPUProvinsi(address Address) returns()
func (_KpuManager *KpuManagerTransactorSession) DeactivateKPUProvinsi(Address common.Address) (*types.Transaction, error) {
	return _KpuManager.Contract.DeactivateKPUProvinsi(&_KpuManager.TransactOpts, Address)
}

// RegisterKPUKota is a paid mutator transaction binding the contract method 0x92423815.
//
// Solidity: function registerKPUKota(address Address, string name, string region) returns()
func (_KpuManager *KpuManagerTransactor) RegisterKPUKota(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error) {
	return _KpuManager.contract.Transact(opts, "registerKPUKota", Address, name, region)
}

// RegisterKPUKota is a paid mutator transaction binding the contract method 0x92423815.
//
// Solidity: function registerKPUKota(address Address, string name, string region) returns()
func (_KpuManager *KpuManagerSession) RegisterKPUKota(Address common.Address, name string, region string) (*types.Transaction, error) {
	return _KpuManager.Contract.RegisterKPUKota(&_KpuManager.TransactOpts, Address, name, region)
}

// RegisterKPUKota is a paid mutator transaction binding the contract method 0x92423815.
//
// Solidity: function registerKPUKota(address Address, string name, string region) returns()
func (_KpuManager *KpuManagerTransactorSession) RegisterKPUKota(Address common.Address, name string, region string) (*types.Transaction, error) {
	return _KpuManager.Contract.RegisterKPUKota(&_KpuManager.TransactOpts, Address, name, region)
}

// RegisterKPUProvinsi is a paid mutator transaction binding the contract method 0x2db33a2f.
//
// Solidity: function registerKPUProvinsi(address Address, string name, string region) returns()
func (_KpuManager *KpuManagerTransactor) RegisterKPUProvinsi(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error) {
	return _KpuManager.contract.Transact(opts, "registerKPUProvinsi", Address, name, region)
}

// RegisterKPUProvinsi is a paid mutator transaction binding the contract method 0x2db33a2f.
//
// Solidity: function registerKPUProvinsi(address Address, string name, string region) returns()
func (_KpuManager *KpuManagerSession) RegisterKPUProvinsi(Address common.Address, name string, region string) (*types.Transaction, error) {
	return _KpuManager.Contract.RegisterKPUProvinsi(&_KpuManager.TransactOpts, Address, name, region)
}

// RegisterKPUProvinsi is a paid mutator transaction binding the contract method 0x2db33a2f.
//
// Solidity: function registerKPUProvinsi(address Address, string name, string region) returns()
func (_KpuManager *KpuManagerTransactorSession) RegisterKPUProvinsi(Address common.Address, name string, region string) (*types.Transaction, error) {
	return _KpuManager.Contract.RegisterKPUProvinsi(&_KpuManager.TransactOpts, Address, name, region)
}

// KpuManagerKPUKotaDeactivatedIterator is returned from FilterKPUKotaDeactivated and is used to iterate over the raw logs and unpacked data for KPUKotaDeactivated events raised by the KpuManager contract.
type KpuManagerKPUKotaDeactivatedIterator struct {
	Event *KpuManagerKPUKotaDeactivated // Event containing the contract specifics and raw log

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
func (it *KpuManagerKPUKotaDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KpuManagerKPUKotaDeactivated)
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
		it.Event = new(KpuManagerKPUKotaDeactivated)
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
func (it *KpuManagerKPUKotaDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KpuManagerKPUKotaDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KpuManagerKPUKotaDeactivated represents a KPUKotaDeactivated event raised by the KpuManager contract.
type KpuManagerKPUKotaDeactivated struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKPUKotaDeactivated is a free log retrieval operation binding the contract event 0x4287f31e2e9e49d47750bbe8bcdc9848644c0566e597ce529738d63b27b7b992.
//
// Solidity: event KPUKotaDeactivated(address indexed Address)
func (_KpuManager *KpuManagerFilterer) FilterKPUKotaDeactivated(opts *bind.FilterOpts, Address []common.Address) (*KpuManagerKPUKotaDeactivatedIterator, error) {

	var AddressRule []interface{}
	for _, AddressItem := range Address {
		AddressRule = append(AddressRule, AddressItem)
	}

	logs, sub, err := _KpuManager.contract.FilterLogs(opts, "KPUKotaDeactivated", AddressRule)
	if err != nil {
		return nil, err
	}
	return &KpuManagerKPUKotaDeactivatedIterator{contract: _KpuManager.contract, event: "KPUKotaDeactivated", logs: logs, sub: sub}, nil
}

// WatchKPUKotaDeactivated is a free log subscription operation binding the contract event 0x4287f31e2e9e49d47750bbe8bcdc9848644c0566e597ce529738d63b27b7b992.
//
// Solidity: event KPUKotaDeactivated(address indexed Address)
func (_KpuManager *KpuManagerFilterer) WatchKPUKotaDeactivated(opts *bind.WatchOpts, sink chan<- *KpuManagerKPUKotaDeactivated, Address []common.Address) (event.Subscription, error) {

	var AddressRule []interface{}
	for _, AddressItem := range Address {
		AddressRule = append(AddressRule, AddressItem)
	}

	logs, sub, err := _KpuManager.contract.WatchLogs(opts, "KPUKotaDeactivated", AddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KpuManagerKPUKotaDeactivated)
				if err := _KpuManager.contract.UnpackLog(event, "KPUKotaDeactivated", log); err != nil {
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

// ParseKPUKotaDeactivated is a log parse operation binding the contract event 0x4287f31e2e9e49d47750bbe8bcdc9848644c0566e597ce529738d63b27b7b992.
//
// Solidity: event KPUKotaDeactivated(address indexed Address)
func (_KpuManager *KpuManagerFilterer) ParseKPUKotaDeactivated(log types.Log) (*KpuManagerKPUKotaDeactivated, error) {
	event := new(KpuManagerKPUKotaDeactivated)
	if err := _KpuManager.contract.UnpackLog(event, "KPUKotaDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KpuManagerKPUKotaRegisteredIterator is returned from FilterKPUKotaRegistered and is used to iterate over the raw logs and unpacked data for KPUKotaRegistered events raised by the KpuManager contract.
type KpuManagerKPUKotaRegisteredIterator struct {
	Event *KpuManagerKPUKotaRegistered // Event containing the contract specifics and raw log

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
func (it *KpuManagerKPUKotaRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KpuManagerKPUKotaRegistered)
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
		it.Event = new(KpuManagerKPUKotaRegistered)
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
func (it *KpuManagerKPUKotaRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KpuManagerKPUKotaRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KpuManagerKPUKotaRegistered represents a KPUKotaRegistered event raised by the KpuManager contract.
type KpuManagerKPUKotaRegistered struct {
	Address common.Address
	Name    string
	Region  string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKPUKotaRegistered is a free log retrieval operation binding the contract event 0x9c6f19aa0cf3f6a0d4fdbd9269f4a794835bf59c3c7561af88f025740ef3cb77.
//
// Solidity: event KPUKotaRegistered(address indexed Address, string name, string region)
func (_KpuManager *KpuManagerFilterer) FilterKPUKotaRegistered(opts *bind.FilterOpts, Address []common.Address) (*KpuManagerKPUKotaRegisteredIterator, error) {

	var AddressRule []interface{}
	for _, AddressItem := range Address {
		AddressRule = append(AddressRule, AddressItem)
	}

	logs, sub, err := _KpuManager.contract.FilterLogs(opts, "KPUKotaRegistered", AddressRule)
	if err != nil {
		return nil, err
	}
	return &KpuManagerKPUKotaRegisteredIterator{contract: _KpuManager.contract, event: "KPUKotaRegistered", logs: logs, sub: sub}, nil
}

// WatchKPUKotaRegistered is a free log subscription operation binding the contract event 0x9c6f19aa0cf3f6a0d4fdbd9269f4a794835bf59c3c7561af88f025740ef3cb77.
//
// Solidity: event KPUKotaRegistered(address indexed Address, string name, string region)
func (_KpuManager *KpuManagerFilterer) WatchKPUKotaRegistered(opts *bind.WatchOpts, sink chan<- *KpuManagerKPUKotaRegistered, Address []common.Address) (event.Subscription, error) {

	var AddressRule []interface{}
	for _, AddressItem := range Address {
		AddressRule = append(AddressRule, AddressItem)
	}

	logs, sub, err := _KpuManager.contract.WatchLogs(opts, "KPUKotaRegistered", AddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KpuManagerKPUKotaRegistered)
				if err := _KpuManager.contract.UnpackLog(event, "KPUKotaRegistered", log); err != nil {
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

// ParseKPUKotaRegistered is a log parse operation binding the contract event 0x9c6f19aa0cf3f6a0d4fdbd9269f4a794835bf59c3c7561af88f025740ef3cb77.
//
// Solidity: event KPUKotaRegistered(address indexed Address, string name, string region)
func (_KpuManager *KpuManagerFilterer) ParseKPUKotaRegistered(log types.Log) (*KpuManagerKPUKotaRegistered, error) {
	event := new(KpuManagerKPUKotaRegistered)
	if err := _KpuManager.contract.UnpackLog(event, "KPUKotaRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KpuManagerKPUProvinsiDeactivatedIterator is returned from FilterKPUProvinsiDeactivated and is used to iterate over the raw logs and unpacked data for KPUProvinsiDeactivated events raised by the KpuManager contract.
type KpuManagerKPUProvinsiDeactivatedIterator struct {
	Event *KpuManagerKPUProvinsiDeactivated // Event containing the contract specifics and raw log

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
func (it *KpuManagerKPUProvinsiDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KpuManagerKPUProvinsiDeactivated)
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
		it.Event = new(KpuManagerKPUProvinsiDeactivated)
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
func (it *KpuManagerKPUProvinsiDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KpuManagerKPUProvinsiDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KpuManagerKPUProvinsiDeactivated represents a KPUProvinsiDeactivated event raised by the KpuManager contract.
type KpuManagerKPUProvinsiDeactivated struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKPUProvinsiDeactivated is a free log retrieval operation binding the contract event 0x9821f051eb616938dcfef57fac817566f3954f1ed18aa52b2916921cb2bacc69.
//
// Solidity: event KPUProvinsiDeactivated(address indexed Address)
func (_KpuManager *KpuManagerFilterer) FilterKPUProvinsiDeactivated(opts *bind.FilterOpts, Address []common.Address) (*KpuManagerKPUProvinsiDeactivatedIterator, error) {

	var AddressRule []interface{}
	for _, AddressItem := range Address {
		AddressRule = append(AddressRule, AddressItem)
	}

	logs, sub, err := _KpuManager.contract.FilterLogs(opts, "KPUProvinsiDeactivated", AddressRule)
	if err != nil {
		return nil, err
	}
	return &KpuManagerKPUProvinsiDeactivatedIterator{contract: _KpuManager.contract, event: "KPUProvinsiDeactivated", logs: logs, sub: sub}, nil
}

// WatchKPUProvinsiDeactivated is a free log subscription operation binding the contract event 0x9821f051eb616938dcfef57fac817566f3954f1ed18aa52b2916921cb2bacc69.
//
// Solidity: event KPUProvinsiDeactivated(address indexed Address)
func (_KpuManager *KpuManagerFilterer) WatchKPUProvinsiDeactivated(opts *bind.WatchOpts, sink chan<- *KpuManagerKPUProvinsiDeactivated, Address []common.Address) (event.Subscription, error) {

	var AddressRule []interface{}
	for _, AddressItem := range Address {
		AddressRule = append(AddressRule, AddressItem)
	}

	logs, sub, err := _KpuManager.contract.WatchLogs(opts, "KPUProvinsiDeactivated", AddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KpuManagerKPUProvinsiDeactivated)
				if err := _KpuManager.contract.UnpackLog(event, "KPUProvinsiDeactivated", log); err != nil {
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

// ParseKPUProvinsiDeactivated is a log parse operation binding the contract event 0x9821f051eb616938dcfef57fac817566f3954f1ed18aa52b2916921cb2bacc69.
//
// Solidity: event KPUProvinsiDeactivated(address indexed Address)
func (_KpuManager *KpuManagerFilterer) ParseKPUProvinsiDeactivated(log types.Log) (*KpuManagerKPUProvinsiDeactivated, error) {
	event := new(KpuManagerKPUProvinsiDeactivated)
	if err := _KpuManager.contract.UnpackLog(event, "KPUProvinsiDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KpuManagerKPUProvinsiRegisteredIterator is returned from FilterKPUProvinsiRegistered and is used to iterate over the raw logs and unpacked data for KPUProvinsiRegistered events raised by the KpuManager contract.
type KpuManagerKPUProvinsiRegisteredIterator struct {
	Event *KpuManagerKPUProvinsiRegistered // Event containing the contract specifics and raw log

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
func (it *KpuManagerKPUProvinsiRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KpuManagerKPUProvinsiRegistered)
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
		it.Event = new(KpuManagerKPUProvinsiRegistered)
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
func (it *KpuManagerKPUProvinsiRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KpuManagerKPUProvinsiRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KpuManagerKPUProvinsiRegistered represents a KPUProvinsiRegistered event raised by the KpuManager contract.
type KpuManagerKPUProvinsiRegistered struct {
	Address common.Address
	Name    string
	Region  string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKPUProvinsiRegistered is a free log retrieval operation binding the contract event 0xe253ed8255fe5723a22412bf5804ee9034911e408cdbfe8547438f23ef231035.
//
// Solidity: event KPUProvinsiRegistered(address indexed Address, string name, string region)
func (_KpuManager *KpuManagerFilterer) FilterKPUProvinsiRegistered(opts *bind.FilterOpts, Address []common.Address) (*KpuManagerKPUProvinsiRegisteredIterator, error) {

	var AddressRule []interface{}
	for _, AddressItem := range Address {
		AddressRule = append(AddressRule, AddressItem)
	}

	logs, sub, err := _KpuManager.contract.FilterLogs(opts, "KPUProvinsiRegistered", AddressRule)
	if err != nil {
		return nil, err
	}
	return &KpuManagerKPUProvinsiRegisteredIterator{contract: _KpuManager.contract, event: "KPUProvinsiRegistered", logs: logs, sub: sub}, nil
}

// WatchKPUProvinsiRegistered is a free log subscription operation binding the contract event 0xe253ed8255fe5723a22412bf5804ee9034911e408cdbfe8547438f23ef231035.
//
// Solidity: event KPUProvinsiRegistered(address indexed Address, string name, string region)
func (_KpuManager *KpuManagerFilterer) WatchKPUProvinsiRegistered(opts *bind.WatchOpts, sink chan<- *KpuManagerKPUProvinsiRegistered, Address []common.Address) (event.Subscription, error) {

	var AddressRule []interface{}
	for _, AddressItem := range Address {
		AddressRule = append(AddressRule, AddressItem)
	}

	logs, sub, err := _KpuManager.contract.WatchLogs(opts, "KPUProvinsiRegistered", AddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KpuManagerKPUProvinsiRegistered)
				if err := _KpuManager.contract.UnpackLog(event, "KPUProvinsiRegistered", log); err != nil {
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

// ParseKPUProvinsiRegistered is a log parse operation binding the contract event 0xe253ed8255fe5723a22412bf5804ee9034911e408cdbfe8547438f23ef231035.
//
// Solidity: event KPUProvinsiRegistered(address indexed Address, string name, string region)
func (_KpuManager *KpuManagerFilterer) ParseKPUProvinsiRegistered(log types.Log) (*KpuManagerKPUProvinsiRegistered, error) {
	event := new(KpuManagerKPUProvinsiRegistered)
	if err := _KpuManager.contract.UnpackLog(event, "KPUProvinsiRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
