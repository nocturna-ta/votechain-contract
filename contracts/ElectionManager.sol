// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.25;

import "./Interfaces.sol";
import "./Context.sol";

contract ElectionManager is IElectionManager, MultiERC2771Context {
    error InvalidElection();
    error ElectionAlreadyExists();
    error UnauthorizedCaller();
    error ElectionNumberMismatch();
    error VoterNotRegistered();
    error VoterAlreadyVoted();
    error NIKMismatch();
    error ElectionNotActive();

    IVotechainBase public base;
    IVoterManager public voterManager;

    mapping(string => Election) public _electionss;
    Election[] public electionAddressArray;

    event ElectionAdded(string indexed electionId, string electionNo);
    event ElectionStatusChange(string indexed electionId, string indexed electionNo, bool isActive);
    event VoteCasted(string indexed nik, string indexed electionId, string indexed electionNo);

    constructor(
        address _baseAddress,
        address _voterManagerAddress,
        address[] memory trustedForwarders
    ) MultiERC2771Context(trustedForwarders) {
        base = IVotechainBase(_baseAddress);
        voterManager = IVoterManager(_voterManagerAddress);
    }

    function elections(string calldata electionId) external view returns (Election memory) {
        return _electionss[electionId];
    }

    modifier onlyKpuAdmin() {
        if (_msgSender() != base.kpuAdmin()) revert UnauthorizedCaller();
        _;
    }

    modifier votingIsActive() {
        if (!base.votingActive()) revert UnauthorizedCaller();
        _;
    }

    function addElection(
        string calldata electionId,
        string calldata electionNo
    ) external override onlyKpuAdmin {

        Election memory election = _electionss[electionId];
        if (bytes(election.id).length != 0) revert ElectionAlreadyExists();

        for (uint i = 0; i < electionAddressArray.length; i++) {
            if (keccak256(bytes(electionAddressArray[i].electionNo)) == keccak256(bytes(electionNo))) {
                revert ElectionAlreadyExists();
            }
        }

        Election memory newElection = Election({
            id: electionId,
            electionNo: electionNo,
            voteCount: 0,
            isActive: true
        });
        election = newElection;
        _electionss[electionId] = newElection;
        electionAddressArray.push(newElection);

        emit ElectionAdded(electionId, electionNo);
    }

    function toggleElectionActive(string calldata electionId, string calldata electionNo) external override onlyKpuAdmin {
        Election storage election = _electionss[electionId];
        if (bytes(election.id).length == 0) revert InvalidElection();
        if (keccak256(bytes(election.electionNo)) != keccak256(bytes(electionNo))) revert ElectionNumberMismatch();

        election.isActive = !election.isActive;
        for (uint i = 0; i < electionAddressArray.length; i++) {
            if (keccak256(bytes(electionAddressArray[i].id)) == keccak256(bytes(electionId))) {
                electionAddressArray[i].isActive = election.isActive;
                break;
            }
        }
        emit ElectionStatusChange(electionId, electionNo, election.isActive);
    }

    function vote(string calldata electionId, string calldata electionNo, string calldata voterNik) external override votingIsActive {
        Election storage election = _electionss[electionId];
        if (bytes(election.id).length == 0) revert InvalidElection();
        if (!election.isActive) revert ElectionNotActive();

        if (keccak256(bytes(election.electionNo)) != keccak256(bytes(electionNo))) revert ElectionNumberMismatch();

        // Use _msgSender() instead of msg.sender for gasless transactions
        address sender = _msgSender();
        IVoterManager.Voter memory voter = voterManager.getVoterByAddress(sender);
        if (voter.hasVoted) revert VoterAlreadyVoted();

        if (keccak256(bytes(voter.nik)) != keccak256(bytes(voterNik))) revert NIKMismatch();

        // Update vote count in storage
        election.voteCount += 1;

        // Update vote count in array
        for (uint i = 0; i < electionAddressArray.length; i++) {
            if (keccak256(bytes(electionAddressArray[i].electionNo)) == keccak256(bytes(electionNo))) {
                electionAddressArray[i].voteCount += 1;
                break;
            }
        }

        emit VoteCasted(voterNik, electionId, electionNo);
    }

    function getElection(string calldata electionId) external view override returns (Election memory) {
        Election storage election = _electionss[electionId];
        if (bytes(election.id).length == 0) revert InvalidElection();
        return election;
    }

    function getAllElection() external view returns (Election[] memory) {
        return electionAddressArray;
    }

    function getElectionByNo(string calldata electionNo) external view returns (Election memory) {
        for (uint i = 0; i < electionAddressArray.length; i++) {
            if (keccak256(bytes(electionAddressArray[i].electionNo)) == keccak256(bytes(electionNo))) {
                return electionAddressArray[i];
            }
        }
        revert InvalidElection();
    }

    /**
     * @dev Add a trusted forwarder for gasless transactions
     */
    function addTrustedForwarder(address forwarder) external onlyKpuAdmin {
        _addTrustedForwarder(forwarder);
    }

    /**
     * @dev Remove a trusted forwarder
     */
    function removeTrustedForwarder(address forwarder) external onlyKpuAdmin {
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