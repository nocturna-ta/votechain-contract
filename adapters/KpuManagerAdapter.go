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

func (k *KpuManagerAdapter) DeactivateKPUKota(opts *bind.TransactOpts, Address common.Address) (*types.Transaction, error) {
	return k.manager.DeactivateKPUKota(opts, Address)
}

func (k *KpuManagerAdapter) DeactivateKPUProvinsi(opts *bind.TransactOpts, Address common.Address) (*types.Transaction, error) {
	return k.manager.DeactivateKPUProvinsi(opts, Address)
}

func (k *KpuManagerAdapter) GetAllKPUKota(opts *bind.CallOpts) ([]kpuManager.IKPUManagerKPUKota, error) {
	return k.manager.GetAllKPUKota(opts)
}

func (k *KpuManagerAdapter) GetAllKPUProvinsi(opts *bind.CallOpts) ([]kpuManager.IKPUManagerKPUProvinsi, error) {
	return k.manager.GetAllKPUProvinsi(opts)
}

func (k *KpuManagerAdapter) GetKpuKotaByAddress(opts *bind.CallOpts, Address common.Address) (kpuManager.IKPUManagerKPUKota, error) {
	return k.manager.GetKpuKotaByAddress(opts, Address)
}

func (k *KpuManagerAdapter) GetKpuProvinsiByAddress(opts *bind.CallOpts, Address common.Address) (kpuManager.IKPUManagerKPUProvinsi, error) {
	return k.manager.GetKpuProvinsiByAddress(opts, Address)
}

func (k *KpuManagerAdapter) KpuKota(opts *bind.CallOpts, addr common.Address) (kpuManager.IKPUManagerKPUKota, error) {
	return k.manager.KpuKota(opts, addr)
}

func (k *KpuManagerAdapter) KpuProvinsi(opts *bind.CallOpts, addr common.Address) (kpuManager.IKPUManagerKPUProvinsi, error) {
	return k.manager.KpuProvinsi(opts, addr)
}

func (k *KpuManagerAdapter) UpdateKPUProvinsi(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error) {
	return k.manager.UpdateKPUProvinsi(opts, Address, name, region)
}

func (k *KpuManagerAdapter) UpdateKPUKota(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error) {
	return k.manager.UpdateKPUKota(opts, Address, name, region)
}
