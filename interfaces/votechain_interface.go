package interfaces

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/nocturna-ta/votechain-contract/binding"
	"math/big"
)

type IVotechainCaller interface {
	CandidateAddressesArray(opts *bind.CallOpts, arg0 *big.Int) (struct {
		Id          string
		Name        string
		CandidateNo string
		VoteCount   *big.Int
		IsActive    bool
	}, error)
	Candidates(opts *bind.CallOpts, arg0 string) (struct {
		Id          string
		Name        string
		CandidateNo string
		VoteCount   *big.Int
		IsActive    bool
	}, error)
	GetAllCandidates(opts *bind.CallOpts) ([]binding.VotechainCandidate, error)
	GetAllKPUBranches(opts *bind.CallOpts) ([]binding.VotechainKPUBranch, error)
	GetAllVoter(opts *bind.CallOpts) ([]binding.VotechainVoter, error)
	GetBranchByAddress(opts *bind.CallOpts, branchAddress common.Address) (binding.VotechainKPUBranch, error)
	GetCandidate(opts *bind.CallOpts, candidateId string) (binding.VotechainCandidate, error)
	GetCandidateByNo(opts *bind.CallOpts, candidateNo string) (binding.VotechainCandidate, error)
	GetVoterByAddress(opts *bind.CallOpts, voterAddress common.Address) (binding.VotechainVoter, error)
	GetVoterByNIK(opts *bind.CallOpts, nik string) (binding.VotechainVoter, error)
	GetVoterByRegion(opts *bind.CallOpts, region string) ([]binding.VotechainVoter, error)
	KpuAdmin(opts *bind.CallOpts) (common.Address, error)
	KpuBranchAddressesArray(opts *bind.CallOpts, arg0 *big.Int) (struct {
		Name          string
		BranchAddress common.Address
		IsActive      bool
		Region        string
	}, error)
	KpuBranches(opts *bind.CallOpts, arg0 common.Address) (struct {
		Name          string
		BranchAddress common.Address
		IsActive      bool
		Region        string
	}, error)
	VoterAddressesArray(opts *bind.CallOpts, arg0 *big.Int) (struct {
		Nik          string
		VoterAddress common.Address
		HasVoted     bool
		Region       string
		IsRegistered bool
	}, error)
	VoterNIKByAddress(opts *bind.CallOpts, arg0 common.Address) (string, error)
	Voters(opts *bind.CallOpts, arg0 string) (struct {
		Nik          string
		VoterAddress common.Address
		HasVoted     bool
		Region       string
		IsRegistered bool
	}, error)
	VotingActive(opts *bind.CallOpts) (bool, error)
}

// IVotechainTransactor defines the write methods of the Votechain contract
type IVotechainTransactor interface {
	AddCandidate(opts *bind.TransactOpts, candidateId string, name string, candidateNo string) (*types.Transaction, error)
	DeactivateKPUBranch(opts *bind.TransactOpts, branchAddress common.Address) (*types.Transaction, error)
	RegisterKPUBranch(opts *bind.TransactOpts, branchAddress common.Address, name string, region string) (*types.Transaction, error)
	RegisterVoter(opts *bind.TransactOpts, nik string, voterAddress common.Address) (*types.Transaction, error)
	SetKpuAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)
	SetVotingStatus(opts *bind.TransactOpts, status bool) (*types.Transaction, error)
	ToggleCandidateActive(opts *bind.TransactOpts, candidateId string) (*types.Transaction, error)
	Vote(opts *bind.TransactOpts, candidateId string) (*types.Transaction, error)
}

// IVotechainFilterer defines the event filtering methods of the Votechain contract
type IVotechainFilterer interface {
	FilterCandidateAdded(opts *bind.FilterOpts, candidateId []string) (*binding.VotechainCandidateAddedIterator, error)
	WatchCandidateAdded(opts *bind.WatchOpts, sink chan<- *binding.VotechainCandidateAdded, candidateId []string) (event.Subscription, error)
	ParseCandidateAdded(log types.Log) (*binding.VotechainCandidateAdded, error)
	FilterCandidateStatusChanged(opts *bind.FilterOpts, candidateId []string) (*binding.VotechainCandidateStatusChangedIterator, error)
	WatchCandidateStatusChanged(opts *bind.WatchOpts, sink chan<- *binding.VotechainCandidateStatusChanged, candidateId []string) (event.Subscription, error)
	ParseCandidateStatusChanged(log types.Log) (*binding.VotechainCandidateStatusChanged, error)
	FilterKPUBranchDeactivated(opts *bind.FilterOpts, branchAddress []common.Address) (*binding.VotechainKPUBranchDeactivatedIterator, error)
	WatchKPUBranchDeactivated(opts *bind.WatchOpts, sink chan<- *binding.VotechainKPUBranchDeactivated, branchAddress []common.Address) (event.Subscription, error)
	ParseKPUBranchDeactivated(log types.Log) (*binding.VotechainKPUBranchDeactivated, error)
	FilterKPUBranchRegistered(opts *bind.FilterOpts, branchAddress []common.Address) (*binding.VotechainKPUBranchRegisteredIterator, error)
	WatchKPUBranchRegistered(opts *bind.WatchOpts, sink chan<- *binding.VotechainKPUBranchRegistered, branchAddress []common.Address) (event.Subscription, error)
	ParseKPUBranchRegistered(log types.Log) (*binding.VotechainKPUBranchRegistered, error)
	FilterVoteCasted(opts *bind.FilterOpts, nik []string, candidateId []string) (*binding.VotechainVoteCastedIterator, error)
	WatchVoteCasted(opts *bind.WatchOpts, sink chan<- *binding.VotechainVoteCasted, nik []string, candidateId []string) (event.Subscription, error)
	ParseVoteCasted(log types.Log) (*binding.VotechainVoteCasted, error)
	FilterVoterRegistered(opts *bind.FilterOpts, nik []string, voterAddress []common.Address) (*binding.VotechainVoterRegisteredIterator, error)
	WatchVoterRegistered(opts *bind.WatchOpts, sink chan<- *binding.VotechainVoterRegistered, nik []string, voterAddress []common.Address) (event.Subscription, error)
	ParseVoterRegistered(log types.Log) (*binding.VotechainVoterRegistered, error)
	FilterVotingStatusChanged(opts *bind.FilterOpts) (*binding.VotechainVotingStatusChangedIterator, error)
	WatchVotingStatusChanged(opts *bind.WatchOpts, sink chan<- *binding.VotechainVotingStatusChanged) (event.Subscription, error)
	ParseVotingStatusChanged(log types.Log) (*binding.VotechainVotingStatusChanged, error)
}

// IVotechain combines all the interfaces
type IVotechain interface {
	IVotechainCaller
	IVotechainTransactor
	IVotechainFilterer
}
