// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.25;

interface IVotechainBase {
    function kpuAdmin() external view returns (address);
    function votingActive() external view returns (bool);
    function setVotingStatus(bool status) external;
    function setKpuAdmin(address newAdmin) external;
}

interface IKPUManager {
    struct KPUProvinsi {
        string name;
        address Address;
        bool isActive;
        string region;
    }

    struct KPUKota {
        string name;
        address Address;
        bool isActive;
        string region;
    }

    function kpuProvinsi(address) external view returns (KPUProvinsi memory);
    function kpuKota(address) external view returns (KPUKota memory);
    function registerKPUProvinsi(address Address, string calldata name, string calldata region) external;
    function registerKPUKota(address Address, string calldata name, string calldata region) external;
    function deactivateKPUProvinsi(address Address) external;
    function deactivateKPUKota(address Address) external;
    function isKPUKota(address kpuAddress) external view returns (bool);
    function isKPUProvinsi(address kpuAddress) external view returns (bool);
    function getKpuKotaRegion(address kpuAddress) external view returns (string memory);
    function getAllKPUProvinsi() external view returns (KPUProvinsi[] memory);
    function getAllKPUKota() external view returns (KPUKota[] memory);
    function getKpuProvinsiByAddress(address Address) external view returns (KPUProvinsi memory);
    function getKpuKotaByAddress(address Address) external view returns (KPUKota memory);
    function updateKPUProvinsi(address Address, string calldata name, string calldata region) external;
    function updateKPUKota(address Address, string calldata name, string calldata region) external;
}

interface IVoterManager {
    struct Voter {
        string nik;
        address voterAddress;
        bool hasVoted;
        string region;
        bool isRegistered;
    }

    function registerVoter(string calldata nik, address voterAddress) external;
    function markVoted(address voterAddress) external returns (bool);
    function getVoterByAddress(address voterAddress) external view returns (Voter memory);
    function getVoterNikByAddress(address voterAddress) external view returns (string memory);
    function getAllVoter() external view returns (Voter[] memory);
    function getVoterByNIK(string calldata nik) external view returns (Voter memory);
    function getVoterByRegion(string calldata region) external view returns (Voter[] memory);
}

interface IElectionManager {
    struct Election {
        string id;
        string electionNo;
        uint256 voteCount;
        bool isActive;
    }

    function addElection(string calldata electionId, string calldata electionNo) external;
    function toggleElectionActive(string calldata electionId, string calldata electionNo) external;
    function vote(string calldata electionId, string calldata electionNo, string memory voterNik) external;
    function getElection(string calldata electionId) external view returns (Election memory);
    function getAllElection() external view returns (Election[] memory);
    function getElectionByNo(string calldata electionNo) external view returns (Election memory);
}
