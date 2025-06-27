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

func (k *KpuManagerAdapter) GetAllKPUKota(opts *bind.CallOpts) (*[]kpuManager.IKPUManagerKPUKota, error) {
	kpu, err := k.manager.GetAllKPUKota(opts)
	if err != nil {
		return nil, err
	}
	return &kpu, nil
}

func (k *KpuManagerAdapter) GetAllKPUProvinsi(opts *bind.CallOpts) (*[]kpuManager.IKPUManagerKPUProvinsi, error) {
	kpu, err := k.manager.GetAllKPUProvinsi(opts)
	if err != nil {
		return nil, err
	}
	return &kpu, nil
}

func (k *KpuManagerAdapter) GetKpuKotaByAddress(opts *bind.CallOpts, Address common.Address) (*kpuManager.IKPUManagerKPUKota, error) {
	kpu, err := k.manager.GetKpuKotaByAddress(opts, Address)
	if err != nil {
		return nil, err
	}
	return &kpu, nil
}

func (k *KpuManagerAdapter) GetKpuProvinsiByAddress(opts *bind.CallOpts, Address common.Address) (*kpuManager.IKPUManagerKPUProvinsi, error) {
	kpu, err := k.manager.GetKpuProvinsiByAddress(opts, Address)
	if err != nil {
		return nil, err
	}
	return &kpu, nil
}

func (k *KpuManagerAdapter) KpuKota(opts *bind.CallOpts, addr common.Address) (*kpuManager.IKPUManagerKPUKota, error) {
	kpu, err := k.manager.KpuKota(opts, addr)
	if err != nil {
		return nil, err
	}

	return &kpu, nil
}

func (k *KpuManagerAdapter) KpuProvinsi(opts *bind.CallOpts, addr common.Address) (*kpuManager.IKPUManagerKPUProvinsi, error) {
	kpu, err := k.manager.KpuProvinsi(opts, addr)
	if err != nil {
		return nil, err
	}

	return &kpu, nil
}

func (k *KpuManagerAdapter) UpdateKPUProvinsi(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error) {
	return k.manager.UpdateKPUProvinsi(opts, Address, name, region)
}

func (k *KpuManagerAdapter) UpdateKPUKota(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error) {
	return k.manager.UpdateKPUKota(opts, Address, name, region)
}
