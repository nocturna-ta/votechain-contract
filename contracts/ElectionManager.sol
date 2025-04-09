// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.25;

import "./Interfaces.sol";

contract ElectionManager is IElectionManager {
    error InvalidElection();
    error ElectionAlreadyExists();
    error UnauthorizedCaller();

    IVotechainBase public base;
    IVoterManager public voterManager;

    mapping(string => Election) public _electionss;
    Election[] public electionAddressArray;

    event ElectionAdded(string indexed electionId, string electionName);
    event ElectionStatusChange(string indexed electionId, bool isActive);
    event VoteCasted(string indexed nik, string indexed candidateId);

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
        string calldata electionName,
        string calldata electionNo
    ) external override onlyKpuAdmin {
        if (bytes(_electionss[electionId].id).length != 0) revert ElectionAlreadyExists();

        Election memory newElection = Election({
            id: electionId,
            electionName: electionName,
            electionNo: electionNo,
            voteCount: 0,
            isActive: true
        });
        _electionss[electionId] = newElection;
        electionAddressArray.push(newElection);

        emit ElectionAdded(electionId, electionName);
    }

    function toggleElectionActive(string calldata electionId) external override onlyKpuAdmin {
        if (bytes(_electionss[electionId].id).length == 0) revert InvalidElection();
        _electionss[electionId].isActive = !_electionss[electionId].isActive;
        emit ElectionStatusChange(electionId, _electionss[electionId].isActive);
    }

    function vote(string calldata electionId, string memory voterNik) external override votingIsActive {

        if (bytes(_electionss[electionId].id).length == 0 || !_electionss[electionId].isActive)
            revert InvalidElection();

        _electionss[electionId].voteCount++;
        emit VoteCasted(voterNik, electionId);
    }

    function getElection(string calldata electionId) external view override returns (Election memory) {
        if (bytes(_electionss[electionId].id).length == 0) revert InvalidElection();
        return _electionss[electionId];
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
