// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.25;

import "./Interfaces.sol";
import "./Context.sol";

contract VoterManager is IVoterManager, MultiERC2771Context {
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
    event VoterMarkedAsVoted(address indexed voterAddress, string indexed nik);

    constructor(
        address _baseAddress,
        address _kpuManagerAddress,
        address[] memory trustedForwarders
    ) MultiERC2771Context(trustedForwarders) {
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
        if (!kpuManager.isKPUKota(_msgSender())) revert UnauthorizedKPU();
        _;
    }

    function registerVoter(
        string calldata nik,
        address voterAddress
    ) external override onlyKpuKota {
        if (_voterss[nik].isRegistered) revert VoterAlreadyRegistered();
        if (bytes(_voterNIKByAddresses[voterAddress]).length != 0) revert AddressAlreadyRegistered();

        string memory region = kpuManager.getKpuKotaRegion(_msgSender());

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

        for (uint i = 0; i < voterAddressesArray.length; i++) {
            if (voterAddressesArray[i].voterAddress == voterAddress) {
                voterAddressesArray[i].hasVoted = true;
                break;
            }
        }
        emit VoterMarkedAsVoted(voterAddress, nik);
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

    /**
     * @dev Add a trusted forwarder for gasless transactions
     */
    function addTrustedForwarder(address forwarder) external {
        // Only KPU admin should be able to do this
        if (_msgSender() != base.kpuAdmin()) revert UnauthorizedKPU();
        _addTrustedForwarder(forwarder);
    }

    /**
     * @dev Remove a trusted forwarder
     */
    function removeTrustedForwarder(address forwarder) external {
        // Only KPU admin should be able to do this
        if (_msgSender() != base.kpuAdmin()) revert UnauthorizedKPU();
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