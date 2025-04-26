// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.25;

import "./Interfaces.sol";

contract ElectionManager is IElectionManager {
    error InvalidElection();
    error ElectionAlreadyExists();
    error UnauthorizedCaller();
    error ElectionNumberMismatch();

    IVotechainBase public base;
    IVoterManager public voterManager;

    mapping(string => Election) public _electionss;
    Election[] public electionAddressArray;

    event ElectionAdded(string indexed electionId, string electionNo);
    event ElectionStatusChange(string indexed electionId, string indexed electionNo, bool isActive);
    event VoteCasted(string indexed nik, string indexed electionId);

    constructor(address _baseAddress, address _voterManagerAddress) {
        base = IVotechainBase(_baseAddress);
        voterManager = IVoterManager(_voterManagerAddress);
    }

    function elections(string calldata electionId) external view returns (Election memory) {
        return _electionss[electionId];
    }


    modifier onlyKpuAdmin() {
        if (msg.sender != base.kpuAdmin()) revert UnauthorizedCaller();
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

        Election memory newElection = Election({
            id: electionId,
            electionNo: electionNo,
            voteCount: 0,
            isActive: true
        });
        election = newElection;
        electionAddressArray.push(newElection);

        emit ElectionAdded(electionId, electionNo);
    }

    function toggleElectionActive(string calldata electionId, string calldata electionNo) external override onlyKpuAdmin {
        Election storage election = _electionss[electionId];
        if (bytes(election.id).length == 0) revert InvalidElection();
        if (keccak256(bytes(election.electionNo)) != keccak256(bytes(electionNo))) revert ElectionNumberMismatch();

        election.isActive = !election.isActive;
        emit ElectionStatusChange(electionId, electionNo, election.isActive);
    }

    function vote(string calldata electionId, string calldata electionNo, string memory voterNik) external override votingIsActive {
        Election storage election = _electionss[electionId];
        if (bytes(election.id).length == 0 || election.isActive)
            revert InvalidElection();

        if (keccak256(bytes(election.electionNo)) != keccak256(bytes(electionNo))) revert ElectionNumberMismatch();

        election.voteCount++;
        emit VoteCasted(voterNik, electionId);
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
}
