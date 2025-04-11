package interfaces

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nocturna-ta/votechain-contract/binding/electionManager"
	"github.com/nocturna-ta/votechain-contract/binding/kpuManager"
	"github.com/nocturna-ta/votechain-contract/binding/voterManager"
)

type ElectionManagerInterface interface {
	Elections(opts *bind.CallOpts, electionId string) (electionManager.IElectionManagerElection, error)
	GetElection(opts *bind.CallOpts, electionId string) (electionManager.IElectionManagerElection, error)
	GetAllElection(opts *bind.CallOpts) ([]electionManager.IElectionManagerElection, error)
	AddElection(opts *bind.TransactOpts, electionId string, electionName string, electionNo string) (*types.Transaction, error)
	ToggleElectionActive(opts *bind.TransactOpts, electionId string) (*types.Transaction, error)
	Vote(opts *bind.TransactOpts, electionId string, voterNik string) (*types.Transaction, error)
	GetElectionByNo(opts *bind.CallOpts, electionNo string) (electionManager.IElectionManagerElection, error)
}

type VotechainBaseInterface interface {
	KpuAdmin(opts *bind.CallOpts) (common.Address, error)
	VotingActive(opts *bind.CallOpts) (bool, error)
	SetKpuAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)
	SetVotingStatus(opts *bind.TransactOpts, status bool) (*types.Transaction, error)
}

type KpuManagerInterface interface {
	IsKPUKota(opts *bind.CallOpts, kpuAddress common.Address) (bool, error)
	IsKPUProvinsi(opts *bind.CallOpts, kpuAddress common.Address) (bool, error)
	GetKpuKotaRegion(opts *bind.CallOpts, kpuAddress common.Address) (string, error)
	RegisterKPUKota(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error)
	RegisterKPUProvinsi(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error)
	DeactivateKPUKota(opts *bind.TransactOpts, Address common.Address) (*types.Transaction, error)
	DeactivateKPUProvinsi(opts *bind.TransactOpts, Address common.Address) (*types.Transaction, error)
	GetAllKPUKota(opts *bind.CallOpts) ([]kpuManager.IKPUManagerKPUKota, error)
	GetAllKPUProvinsi(opts *bind.CallOpts) ([]kpuManager.IKPUManagerKPUProvinsi, error)
	GetKpuKotaByAddress(opts *bind.CallOpts, Address common.Address) (kpuManager.IKPUManagerKPUKota, error)
	GetKpuProvinsiByAddress(opts *bind.CallOpts, Address common.Address) (kpuManager.IKPUManagerKPUProvinsi, error)
	KpuKota(opts *bind.CallOpts, addr common.Address) (kpuManager.IKPUManagerKPUKota, error)
	KpuProvinsi(opts *bind.CallOpts, addr common.Address) (kpuManager.IKPUManagerKPUProvinsi, error)
}

type VoterManagerInterface interface {
	GetVoterByAddress(opts *bind.CallOpts, voterAddress common.Address) (voterManager.IVoterManagerVoter, error)
	GetVoterNikByAddress(opts *bind.CallOpts, voterAddress common.Address) (string, error)
	RegisterVoter(opts *bind.TransactOpts, nik string, voterAddress common.Address) (*types.Transaction, error)
	MarkVoted(opts *bind.TransactOpts, voterAddress common.Address) (*types.Transaction, error)
	KpuManager(opts *bind.CallOpts) (common.Address, error)
	VoterNIKByAddress(opts *bind.CallOpts, voter common.Address) (string, error)
	VoterNIKByAddresses(opts *bind.CallOpts, arg0 common.Address) (string, error)
	Voters(opts *bind.CallOpts, nik string) (voterManager.IVoterManagerVoter, error)
	GetAllVoter(opts *bind.CallOpts) ([]voterManager.IVoterManagerVoter, error)
	GetVoterByNIK(opts *bind.CallOpts, nik string) (voterManager.IVoterManagerVoter, error)
	GetVoterByRegion(opts *bind.CallOpts, region string) ([]voterManager.IVoterManagerVoter, error)
}

type VotechainInterface interface {
	Vote(opts *bind.TransactOpts, electionId string) (*types.Transaction, error)
	RegisterKPUProvinsi(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error)
	RegisterKPUKota(opts *bind.TransactOpts, Address common.Address, name string, region string) (*types.Transaction, error)
	RegisterVoter(opts *bind.TransactOpts, nik string, voterAddress common.Address) (*types.Transaction, error)
	AddElection(opts *bind.TransactOpts, electionId string, electionName string, electionNo string) (*types.Transaction, error)
	ToggleElectionActive(opts *bind.TransactOpts, electionId string) (*types.Transaction, error)
	SetVotingStatus(opts *bind.TransactOpts, status bool) (*types.Transaction, error)
	SetKpuAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)
}
