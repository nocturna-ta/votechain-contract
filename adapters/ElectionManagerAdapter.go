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

func (e *ElectionManagerAdapter) GetElectionByNo(opts *bind.CallOpts, electionNo string) (*electionManager.IElectionManagerElection, error) {
	election, err := e.manager.GetElectionByNo(opts, electionNo)
	if err != nil {
		return nil, err
	}
	return &election, nil
}

func (e *ElectionManagerAdapter) Elections(opts *bind.CallOpts, electionId string) (*electionManager.IElectionManagerElection, error) {
	election, err := e.manager.Elections(opts, electionId)
	if err != nil {
		return nil, err
	}
	return &election, nil
}

func (e *ElectionManagerAdapter) GetElection(opts *bind.CallOpts, electionId string) (*electionManager.IElectionManagerElection, error) {
	election, err := e.manager.GetElection(opts, electionId)
	if err != nil {
		return nil, err
	}
	return &election, nil
}

func (e *ElectionManagerAdapter) GetAllElection(opts *bind.CallOpts) (*[]electionManager.IElectionManagerElection, error) {
	elections, err := e.manager.GetAllElection(opts)
	if err != nil {
		return nil, err
	}
	return &elections, nil
}

func (e *ElectionManagerAdapter) AddElection(opts *bind.TransactOpts, electionId string, electionNo string) (*types.Transaction, error) {
	return e.manager.AddElection(opts, electionId, electionNo)
}

func (e *ElectionManagerAdapter) ToggleElectionActive(opts *bind.TransactOpts, electionId string, electionNo string) (*types.Transaction, error) {
	return e.manager.ToggleElectionActive(opts, electionId, electionNo)
}

func (e *ElectionManagerAdapter) Vote(opts *bind.TransactOpts, electionId string, electionNo string, voterNik string) (*types.Transaction, error) {
	return e.manager.Vote(opts, electionId, electionNo, voterNik)
}
