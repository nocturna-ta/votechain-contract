package adapters

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nocturna-ta/votechain-contract/binding/kpuManager"
	"github.com/nocturna-ta/votechain-contract/interfaces"
)

type KpuManagerAdapter struct {
	manager *kpuManager.KpuManager
}

func NewKpuManagerAdapter(manager *kpuManager.KpuManager) interfaces.KpuManagerInterface {
	return &KpuManagerAdapter{
		manager: manager,
	}
}

func (k *KpuManagerAdapter) IsKPUKota(opts *bind.CallOpts, kpuAddress common.Address) (bool, error) {
	return k.manager.IsKPUKota(opts, kpuAddress)
}

func (k *KpuManagerAdapter) IsKPUProvinsi(opts *bind.CallOpts, kpuAddress common.Address) (bool, error) {
	return k.manager.IsKPUProvinsi(opts, kpuAddress)
}

func (k *KpuManagerAdapter) GetKpuKotaRegion(opts *bind.CallOpts, kpuAddress common.Address) (string, error) {
	return k.manager.GetKpuKotaRegion(opts, kpuAddress)
}

func (k *KpuManagerAdapter) RegisterKPUKota(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error) {
	return k.manager.RegisterKPUKota(opts, Address, name, region)
}

func (k *KpuManagerAdapter) RegisterKPUProvinsi(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error) {
	return k.manager.RegisterKPUProvinsi(opts, Address, name, region)
}
