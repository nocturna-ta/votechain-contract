// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.25;

import "./Interfaces.sol";

contract Votechain {
    error VoterNotRegistered();
    error AlreadyVoted();
    error InvalidElection();
    error VotingNotActive();

    IVotechainBase public base;
    IKPUManager public kpuManager;
    IVoterManager public voterManager;
    IElectionManager public electionManager;

    constructor(
        address _baseAddress,
        address _kpuManagerAddress,
        address _voterManagerAddress,
        address _electionManagerAddress
    ) {
        base = IVotechainBase(_baseAddress);
        kpuManager = IKPUManager(_kpuManagerAddress);
        voterManager = IVoterManager(_voterManagerAddress);
        electionManager = IElectionManager(_electionManagerAddress);
    }

    // Public-facing function to cast votes
    function vote(string calldata electionId, string calldata electionNo) external {
        if (!base.votingActive()) revert VotingNotActive();

        try voterManager.getVoterByAddress(msg.sender) returns (IVoterManager.Voter memory voter) {
            if (voter.hasVoted) revert AlreadyVoted();

            voterManager.markVoted(msg.sender);

            electionManager.vote(electionId, electionNo, voter.nik);
        } catch {
            revert VoterNotRegistered();
        }
    }

    function registerKPUProvinsi(
        address Address,
        string calldata name,
        string calldata region
    ) external {
        kpuManager.registerKPUProvinsi(Address, name, region);
    }

    function registerKPUKota(
        address Address,
        string calldata name,
        string calldata region
    ) external {
        kpuManager.registerKPUKota(Address, name, region);
    }

    function registerVoter(
        string calldata nik,
        address voterAddress
    ) external {
        voterManager.registerVoter(nik, voterAddress);
    }

    function addElection(
        string calldata electionId,

        string calldata electionNo
    ) external {
        electionManager.addElection(electionId, electionNo);
    }

    function toggleElectionActive(string calldata electionId, string calldata electionNo) external {
        electionManager.toggleElectionActive(electionId, electionNo);
    }

    function setVotingStatus(bool status) external {
        base.setVotingStatus(status);
    }

    function setKpuAdmin(address newAdmin) external {
        base.setKpuAdmin(newAdmin);
    }

    function updateKPUProvinsi(
        address Address,
        string calldata name,
        string calldata region
    ) external {
        kpuManager.updateKPUProvinsi(Address, name, region);
    }

    function updateKPUKota(
        address Address,
        string calldata name,
        string calldata region
    ) external {
        kpuManager.updateKPUKota(Address, name, region);
    }

}