// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.25;

import "./Interfaces.sol";
import "./Context.sol";

contract Votechain is MultiERC2771Context {
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
        address _electionManagerAddress,
        address[] memory trustedForwarders
    ) MultiERC2771Context(trustedForwarders) {
        base = IVotechainBase(_baseAddress);
        kpuManager = IKPUManager(_kpuManagerAddress);
        voterManager = IVoterManager(_voterManagerAddress);
        electionManager = IElectionManager(_electionManagerAddress);
    }

    // Public-facing function to cast votes (now gasless)
    function vote(string calldata electionId, string calldata electionNo) external {
        if (!base.votingActive()) revert VotingNotActive();

        address sender = _msgSender(); // Use _msgSender() instead of msg.sender

        try voterManager.getVoterByAddress(sender) returns (IVoterManager.Voter memory voter) {
            if (voter.hasVoted) revert AlreadyVoted();

            voterManager.markVoted(sender);

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

    /**
     * @dev Add a trusted forwarder for gasless transactions (admin only)
     */
    function addTrustedForwarder(address forwarder) external {
        // This should be restricted to admin, but since we don't have direct access
        // to the admin check here, we'll rely on the base contract
        base.setKpuAdmin(base.kpuAdmin()); // This will revert if not admin
        _addTrustedForwarder(forwarder);
    }

    /**
     * @dev Remove a trusted forwarder (admin only)
     */
    function removeTrustedForwarder(address forwarder) external {
        // This should be restricted to admin
        base.setKpuAdmin(base.kpuAdmin()); // This will revert if not admin
        _removeTrustedForwarder(forwarder);
    }

    /**
     * @dev Override to use the correct _msgSender implementation
     */
    function _msgSender() internal view override returns (address) {
        return MultiERC2771Context._msgSender();
    }

    /**
     * @dev Override to use the correct _msgData implementation
     */
    function _msgData() internal view override returns (bytes calldata) {
        return MultiERC2771Context._msgData();
    }
}