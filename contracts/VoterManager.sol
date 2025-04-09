// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.25;

import "./Interfaces.sol";

contract VoterManager is IVoterManager {
    error VoterAlreadyRegistered();
    error AddressAlreadyRegistered();
    error VoterNotRegistered();
    error AlreadyVoted();
    error UnauthorizedKPU();

    IVotechainBase public base;
    IKPUManager public kpuManager;

    mapping(address => string) public _voterNIKByAddresses;
    mapping(string => Voter) public _voterss;
    Voter[] public voterAddressesArray;

    event VoterRegistered(string indexed nik, address indexed voterAddress, string region);

    constructor(address _baseAddress, address _kpuManagerAddress) {
        base = IVotechainBase(_baseAddress);
        kpuManager = IKPUManager(_kpuManagerAddress);
    }

    function voterNIKByAddress(address voter) external view returns (string memory) {
        return _voterNIKByAddresses[voter];
    }

    function voters(string calldata nik) external view returns (Voter memory) {
        return _voterss[nik];
    }

    modifier onlyKpuKota() {
        if (!kpuManager.isKPUKota(msg.sender)) revert UnauthorizedKPU();
        _;
    }

    function registerVoter(
        string calldata nik,
        address voterAddress
    ) external override onlyKpuKota {
        if (_voterss[nik].isRegistered) revert VoterAlreadyRegistered();
        if (bytes(_voterNIKByAddresses[voterAddress]).length != 0) revert AddressAlreadyRegistered();

        string memory region = kpuManager.getKpuKotaRegion(msg.sender);

        Voter memory newVoter = Voter({
            nik: nik,
            voterAddress: voterAddress,
            hasVoted: false,
            region: region,
            isRegistered: true
        });

        _voterss[nik] = newVoter;
        _voterNIKByAddresses[voterAddress] = nik;
        voterAddressesArray.push(newVoter);

        emit VoterRegistered(nik, voterAddress, region);
    }

    function markVoted(address voterAddress) external override returns (bool) {
        string memory nik = _voterNIKByAddresses[voterAddress];
        if (bytes(nik).length == 0) revert VoterNotRegistered();

        Voter storage voter = _voterss[nik];
        if (voter.hasVoted) revert AlreadyVoted();

        voter.hasVoted = true;
        return true;
    }

    function getVoterNikByAddress(address voterAddress) external view override returns (string memory) {
        string memory nik = _voterNIKByAddresses[voterAddress];
        if (bytes(nik).length == 0) revert VoterNotRegistered();
        return nik;
    }

    function getVoterByAddress(address voterAddress) external view override returns (Voter memory) {
        string memory nik = _voterNIKByAddresses[voterAddress];
        if (bytes(nik).length == 0) revert VoterNotRegistered();
        return _voterss[nik];
    }

    function getAllVoter() external view returns (Voter[] memory) {
        return voterAddressesArray;
    }

    function getVoterByNIK(string calldata nik) external view returns (Voter memory) {
        if (!_voterss[nik].isRegistered) revert VoterNotRegistered();
        return _voterss[nik];
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
}
