package adapters

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nocturna-ta/votechain-contract/binding/electionManager"
	"github.com/nocturna-ta/votechain-contract/interfaces"
)

type ElectionManagerAdapter struct {
	manager *electionManager.ElectionManager
}

func NewElectionManagerAdapter(manager *electionManager.ElectionManager) interfaces.ElectionManagerInterface {
	return &ElectionManagerAdapter{
		manager: manager,
	}
}

func (e *ElectionManagerAdapter) GetElectionByNo(opts *bind.CallOpts, electionNo string) (electionManager.IElectionManagerElection, error) {
	return e.manager.GetElectionByNo(opts, electionNo)
}

func (e *ElectionManagerAdapter) Elections(opts *bind.CallOpts, electionId string) (electionManager.IElectionManagerElection, error) {
	return e.manager.Elections(opts, electionId)
}

func (e *ElectionManagerAdapter) GetElection(opts *bind.CallOpts, electionId string) (electionManager.IElectionManagerElection, error) {
	return e.manager.GetElection(opts, electionId)
}

func (e *ElectionManagerAdapter) GetAllElection(opts *bind.CallOpts) ([]electionManager.IElectionManagerElection, error) {
	return e.manager.GetAllElection(opts)
}

func (e *ElectionManagerAdapter) AddElection(opts *bind.TransactOpts, electionId string, electionName string, electionNo string) (*types.Transaction, error) {
	return e.manager.AddElection(opts, electionId, electionName, electionNo)
}

func (e *ElectionManagerAdapter) ToggleElectionActive(opts *bind.TransactOpts, electionId string) (*types.Transaction, error) {
	return e.manager.ToggleElectionActive(opts, electionId)
}

func (e *ElectionManagerAdapter) Vote(opts *bind.TransactOpts, electionId string, voterNik string) (*types.Transaction, error) {
	return e.manager.Vote(opts, electionId, voterNik)
}
