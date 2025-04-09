package adapters

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nocturna-ta/votechain-contract/binding/votechainBase"
	"github.com/nocturna-ta/votechain-contract/interfaces"
)

type VotechainBaseAdapter struct {
	votechainBase *votechainBase.VotechainBase
}

func NewVotechainBaseAdapter(votechainBase *votechainBase.VotechainBase) interfaces.VotechainBaseInterface {
	return &VotechainBaseAdapter{
		votechainBase: votechainBase,
	}
}

func (v *VotechainBaseAdapter) KpuAdmin(opts *bind.CallOpts) (common.Address, error) {
	return v.KpuAdmin(opts)
}

func (v *VotechainBaseAdapter) VotingActive(opts *bind.CallOpts) (bool, error) {
	return v.VotingActive(opts)
}

func (v *VotechainBaseAdapter) SetKpuAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return v.SetKpuAdmin(opts, newAdmin)
}

func (v *VotechainBaseAdapter) SetVotingStatus(opts *bind.TransactOpts, status bool) (*types.Transaction, error) {
	return v.SetVotingStatus(opts, status)
}
