// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.25;

contract Votechain {
    error OnlyKpuAdmin();
    error OnlyKpuBranch();
    error VotingNotActive();
    error BranchAlreadyRegistered();
    error BranchNotActive();
    error VoterAlreadyRegistered();
    error AddressAlreadyRegistered();
    error VoterNotRegistered();
    error AlreadyVoted();
    error InvalidCandidate();
    error UnauthorizedVoter();
    error CandidateAlreadyExists();

    struct KPUBranch {
        string name;
        address branchAddress;
        bool isActive;
        string region;
    }

    struct Voter {
        string nik;
        address voterAddress;
        bool hasVoted;
        string region;
        bool isRegistered;
    }

    struct Candidate {
        string id; // UUID as a string
        string name;
        uint candidateNo;
        uint256 voteCount;
        bool isActive;
    }

    address public kpuAdmin;
    bool public votingActive;

    mapping(string => Candidate) public candidates; // Mapping from UUID to Candidate
    mapping(address => KPUBranch) public kpuBranches;
    mapping(address => string) public voterNIKByAddress;
    mapping(string => Voter) public voters;

    KPUBranch[] public kpuBranchAddressesArray;
    Voter[] public voterAddressesArray;
    Candidate[] public candidateAddressesArray;

    event KPUBranchRegistered(address indexed branchAddress, string name, string region);
    event KPUBranchDeactivated(address indexed branchAddress);
    event VoterRegistered(string indexed nik, address indexed voterAddress, string region);
    event VoteCasted(string indexed nik, string indexed candidateId);
    event CandidateAdded(string indexed candidateId, string name);
    event CandidateStatusChanged(string indexed candidateId, bool isActive);
    event VotingStatusChanged(bool isActive);

    modifier onlyKpuAdmin() {
        if (msg.sender != kpuAdmin) revert OnlyKpuAdmin();
        _;
    }

    modifier onlyKpuBranch() {
        if (!kpuBranches[msg.sender].isActive) revert OnlyKpuBranch();
        _;
    }

    modifier votingIsActive() {
        if (!votingActive) revert VotingNotActive();
        _;
    }

    constructor() {
        kpuAdmin = msg.sender;
    }

    function registerKPUBranch(
        address branchAddress,
        string calldata name,
        string calldata region
    ) external onlyKpuAdmin {
        if (kpuBranches[branchAddress].isActive) revert BranchAlreadyRegistered();

        KPUBranch memory newBranch = KPUBranch(name, branchAddress, true, region);
        kpuBranches[branchAddress] = newBranch;
        kpuBranchAddressesArray.push(newBranch);

        emit KPUBranchRegistered(branchAddress, name, region);
    }

    function deactivateKPUBranch(address branchAddress) external onlyKpuAdmin {
        if (!kpuBranches[branchAddress].isActive) revert BranchNotActive();
        kpuBranches[branchAddress].isActive = false;
        emit KPUBranchDeactivated(branchAddress);
    }

    function registerVoter(
        string calldata nik,
        address voterAddress
    ) external onlyKpuBranch {
        if (voters[nik].isRegistered) revert VoterAlreadyRegistered();
        if (bytes(voterNIKByAddress[voterAddress]).length != 0) revert AddressAlreadyRegistered();
        Voter memory newVoter = Voter({
            nik: nik,
            voterAddress: voterAddress,
            hasVoted: false,
            region: kpuBranches[msg.sender].region,
            isRegistered: true
        });

        voters[nik] = newVoter;
        voterNIKByAddress[voterAddress] = nik;
        voterAddressesArray.push(newVoter);

        emit VoterRegistered(nik, voterAddress, kpuBranches[msg.sender].region);
    }

    function vote(string calldata candidateId) external votingIsActive {
        string memory nik = voterNIKByAddress[msg.sender];
        if (bytes(nik).length == 0) revert VoterNotRegistered();

        Voter storage voter = voters[nik];
        if (voter.hasVoted) revert AlreadyVoted();
        if (bytes(candidates[candidateId].id).length == 0 || !candidates[candidateId].isActive)
            revert InvalidCandidate();

        voter.hasVoted = true;
        candidates[candidateId].voteCount++;

        emit VoteCasted(nik, candidateId);
    }

    function addCandidate(string calldata candidateId, string calldata name, uint candidateNo) external onlyKpuAdmin {
        if (bytes(candidates[candidateId].id).length != 0) revert CandidateAlreadyExists();

        Candidate memory newCandidate = Candidate({
            id: candidateId,
            name: name,
            candidateNo: candidateNo,
            voteCount: 0,
            isActive: true
        });
        candidates[candidateId] = newCandidate;
        candidateAddressesArray.push(newCandidate);

        emit CandidateAdded(candidateId, name);
    }

    function toggleCandidateActive(string calldata candidateId) external onlyKpuAdmin {
        if (bytes(candidates[candidateId].id).length == 0) revert InvalidCandidate();
        candidates[candidateId].isActive = !candidates[candidateId].isActive;
        emit CandidateStatusChanged(candidateId, candidates[candidateId].isActive);
    }

    function setVotingStatus(bool status) external onlyKpuAdmin {
        votingActive = status;
        emit VotingStatusChanged(status);
    }

    function setKpuAdmin(address newAdmin) external onlyKpuAdmin {
        kpuAdmin = newAdmin;
    }

    function getAllKPUBranches() external view returns (KPUBranch[] memory) {
        return kpuBranchAddressesArray;
    }

    function getBranchByAddress(address branchAddress) external view returns (KPUBranch memory) {
        if (!kpuBranches[branchAddress].isActive) revert BranchNotActive();
        return kpuBranches[branchAddress];
    }

    function getAllVoter() external view returns (Voter[] memory) {
        return voterAddressesArray;
    }

    function getVoterByNIK(string calldata nik) external view returns (Voter memory) {
        return voters[nik];
    }

    function getVoterByAddress(address voterAddress) external view returns (Voter memory) {
        string memory nik = voterNIKByAddress[voterAddress];
        if (bytes(nik).length == 0) revert VoterNotRegistered();
        return voters[nik];
    }

    function getVoterByRegion(string calldata region) external view returns (Voter[] memory) {
        uint256 count;
        for (uint256 i = 0; i < voterAddressesArray.length; i++) {
            if (keccak256(bytes(voterAddressesArray[i].region)) == keccak256(bytes(region))) {
                count++;
            }
        }

        Voter[] memory regionVoters = new Voter[](count);
        uint256 index;
        for (uint256 i = 0; i < voterAddressesArray.length; i++) {
            if (keccak256(bytes(voterAddressesArray[i].region)) == keccak256(bytes(region))) {
                regionVoters[index] = voterAddressesArray[i];
                index++;
            }
        }
        return regionVoters;
    }

    function getCandidate(string calldata candidateId) external view returns (Candidate memory) {
        if (bytes(candidates[candidateId].id).length == 0) revert InvalidCandidate();
        return candidates[candidateId];
    }

    function getAllCandidates() external view returns (Candidate[] memory) {
        return candidateAddressesArray;
    }

    function getCandidateByNo(uint candidateNo) external view returns (Candidate memory) {
        for (uint i = 0; i < candidateAddressesArray.length; i++) {
            if (candidateAddressesArray[i].candidateNo == candidateNo) {
                return candidateAddressesArray[i];
            }
        }
        revert InvalidCandidate();
    }

}