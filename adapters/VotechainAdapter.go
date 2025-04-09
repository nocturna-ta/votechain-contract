package adapters

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nocturna-ta/votechain-contract/binding/votechain"
	"github.com/nocturna-ta/votechain-contract/interfaces"
)

type VotechainAdapter struct {
	votechain *votechain.Votechain
}

func NewVotechainAdapter(votechain *votechain.Votechain) interfaces.VotechainInterface {
	return &VotechainAdapter{
		votechain: votechain,
	}
}

func (v *VotechainAdapter) Vote(opts *bind.TransactOpts, electionId string) (*types.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VotechainAdapter) RegisterKPUProvinsi(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error) {
	return v.votechain.RegisterKPUProvinsi(opts, Address, name, region)
}

func (v *VotechainAdapter) RegisterKPUKota(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error) {
	return v.votechain.RegisterKPUKota(opts, Address, name, region)
}

func (v *VotechainAdapter) RegisterVoter(opts *bind.TransactOpts, nik string, voterAddress common.Address) (*types.Transaction, error) {
	return v.votechain.RegisterVoter(opts, nik, voterAddress)
}

func (v *VotechainAdapter) AddElection(opts *bind.TransactOpts, electionId string, electionName string, electionNo string) (*types.Transaction, error) {
	return v.votechain.AddElection(opts, electionId, electionName, electionNo)
}

func (v *VotechainAdapter) ToggleElectionActive(opts *bind.TransactOpts, electionId string) (*types.Transaction, error) {
	return v.votechain.ToggleElectionActive(opts, electionId)
}

func (v *VotechainAdapter) SetVotingStatus(opts *bind.TransactOpts, status bool) (*types.Transaction, error) {
	return v.votechain.SetVotingStatus(opts, status)
}

func (v *VotechainAdapter) SetKpuAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return v.votechain.SetKpuAdmin(opts, newAdmin)
}
