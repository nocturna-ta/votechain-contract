package adapters

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nocturna-ta/votechain-contract/binding/voterManager"
	"github.com/nocturna-ta/votechain-contract/interfaces"
)

type VoterManagerAdapter struct {
	manager *voterManager.VoterManager
}

func NewVoterManagerAdapter(manager *voterManager.VoterManager) interfaces.VoterManagerInterface {
	return &VoterManagerAdapter{
		manager: manager,
	}
}

func (v *VoterManagerAdapter) GetVoterByAddress(opts *bind.CallOpts, voterAddress common.Address) (*voterManager.IVoterManagerVoter, error) {
	voter, err := v.manager.GetVoterByAddress(opts, voterAddress)
	if err != nil {
		return nil, err
	}

	return &voter, nil
}

func (v *VoterManagerAdapter) GetVoterNikByAddress(opts *bind.CallOpts, voterAddress common.Address) (string, error) {
	return v.manager.GetVoterNikByAddress(opts, voterAddress)
}

func (v *VoterManagerAdapter) RegisterVoter(opts *bind.TransactOpts, nik string, voterAddress common.Address) (*types.Transaction, error) {
	return v.manager.RegisterVoter(opts, nik, voterAddress)
}

func (v *VoterManagerAdapter) MarkVoted(opts *bind.TransactOpts, voterAddress common.Address) (*types.Transaction, error) {
	return v.manager.MarkVoted(opts, voterAddress)
}

func (v *VoterManagerAdapter) KpuManager(opts *bind.CallOpts) (common.Address, error) {
	return v.manager.KpuManager(opts)
}

func (v *VoterManagerAdapter) VoterNIKByAddress(opts *bind.CallOpts, voter common.Address) (string, error) {
	return v.manager.VoterNIKByAddress(opts, voter)
}

func (v *VoterManagerAdapter) VoterNIKByAddresses(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	return v.manager.VoterNIKByAddresses(opts, arg0)
}

func (v *VoterManagerAdapter) Voters(opts *bind.CallOpts, nik string) (*voterManager.IVoterManagerVoter, error) {
	voter, err := v.manager.Voters(opts, nik)
	if err != nil {
		return nil, err
	}

	return &voter, nil
}

func (v *VoterManagerAdapter) GetAllVoter(opts *bind.CallOpts) (*[]voterManager.IVoterManagerVoter, error) {
	voters, err := v.manager.GetAllVoter(opts)
	if err != nil {
		return nil, err
	}

	return &voters, nil
}

func (v *VoterManagerAdapter) GetVoterByNIK(opts *bind.CallOpts, nik string) (*voterManager.IVoterManagerVoter, error) {
	voter, err := v.manager.GetVoterByNIK(opts, nik)
	if err != nil {
		return nil, err
	}

	return &voter, nil
}

func (v *VoterManagerAdapter) GetVoterByRegion(opts *bind.CallOpts, region string) (*[]voterManager.IVoterManagerVoter, error) {
	voters, err := v.manager.GetVoterByRegion(opts, region)
	if err != nil {
		return nil, err
	}
	return &voters, nil
}
