// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.25;

import "./Interfaces.sol";
import "./VotechainBase.sol";

contract KPUManager is IKPUManager{

    error OnlyKpuProvinsi();
    error OnlyKpuKota();
    error KPUProvinsiAlreadyRegistered();
    error KPUProvinsiNotActive();
    error KPUKotaAlreadyRegistered();
    error KPUKotaNotActive();
    error KPUNotFound();

    IVotechainBase public base;

    mapping(address => KPUProvinsi) internal _kpuProvinsi;
    mapping(address => KPUKota) internal _kpuKota;

    KPUProvinsi[] public kpuProvinsiAddressesArray;
    KPUKota[] public kpuKotaAddressesArray;

    event KPUProvinsiRegistered(address indexed Address, string name, string region);
    event KPUKotaRegistered(address indexed Address, string name, string region);
    event KPUKotaDeactivated(address indexed Address);
    event KPUProvinsiDeactivated(address indexed Address);
    event KPUProvinsiUpdated(address indexed Address, string name, string region);
    event KPUKotaUpdated(address indexed Address, string name, string region);

    modifier onlyKpuAdmin() {
        if (msg.sender != base.kpuAdmin()) revert OnlyKpuProvinsi();
        _;
    }

    modifier onlyKpuProvinsi() {
        if (!_kpuProvinsi[msg.sender].isActive) revert OnlyKpuProvinsi();
        _;
    }

    modifier onlyKpuKota(){
        if (!_kpuKota[msg.sender].isActive) revert OnlyKpuKota();
        _;
    }

    constructor(address _baseAddress) {
        base = IVotechainBase(_baseAddress);
    }

    function kpuProvinsi(address addr) external view override returns (KPUProvinsi memory) {
        return _kpuProvinsi[addr];
    }

    function kpuKota(address addr) external view override returns (KPUKota memory) {
        return _kpuKota[addr];
    }

    function registerKPUProvinsi(
        address Address,
        string calldata name,
        string calldata region
    ) external override onlyKpuAdmin {
        if (_kpuProvinsi[Address].isActive) revert KPUProvinsiAlreadyRegistered();

        KPUProvinsi memory newKPU = KPUProvinsi(name, Address, true, region);
        _kpuProvinsi[Address] = newKPU;
        kpuProvinsiAddressesArray.push(newKPU);

        emit KPUProvinsiRegistered(Address, name, region);
    }

    function registerKPUKota(
        address Address,
        string calldata name,
        string calldata region
    ) external override onlyKpuProvinsi {
        if (_kpuKota[Address].isActive) revert KPUKotaAlreadyRegistered();

        KPUKota memory newKPU = KPUKota(name, Address, true, region);
        _kpuKota[Address] = newKPU;
        kpuKotaAddressesArray.push(newKPU);

        emit KPUKotaRegistered(Address, name, region);
    }

    function deactivateKPUProvinsi(address Address) external override onlyKpuAdmin {
        if (!_kpuProvinsi[Address].isActive) revert KPUProvinsiNotActive();
        _kpuProvinsi[Address].isActive = false;
        emit KPUProvinsiDeactivated(Address);
    }

    function deactivateKPUKota(address Address) external override onlyKpuAdmin {
        if (!_kpuKota[Address].isActive) revert KPUKotaNotActive();
        _kpuKota[Address].isActive = false;
        emit KPUKotaDeactivated(Address);
    }

    // Helper functions for other contracts
    function isKPUKota(address kpuAddress) external view override returns (bool) {
        return _kpuKota[kpuAddress].isActive;
    }

    function isKPUProvinsi(address kpuAddress) external view override returns (bool) {
        return _kpuProvinsi[kpuAddress].isActive;
    }

    function getKpuKotaRegion(address kpuAddress) external view override returns (string memory) {
        return _kpuKota[kpuAddress].region;
    }

    function getAllKPUProvinsi() external view returns (KPUProvinsi[] memory) {
        return kpuProvinsiAddressesArray;
    }

    function getAllKPUKota() external view returns (KPUKota[] memory) {
        return kpuKotaAddressesArray;
    }

    function getKpuProvinsiByAddress(address Address) external view returns (KPUProvinsi memory) {
        if (!_kpuProvinsi[Address].isActive) revert KPUProvinsiNotActive();
        return _kpuProvinsi[Address];
    }

    function getKpuKotaByAddress(address Address) external view returns (KPUKota memory) {
        if (!_kpuKota[Address].isActive) revert KPUKotaNotActive();
        return _kpuKota[Address];
    }

    function updateKPUProvinsi(
        address Address,
        string calldata name,
        string calldata region
    ) external override onlyKpuProvinsi {
        if (!_kpuProvinsi[Address].isActive) revert KPUProvinsiNotActive();

        _kpuProvinsi[Address].name = name;
        _kpuProvinsi[Address].region = region;

        for (uint i = 0; i < kpuProvinsiAddressesArray.length; i++) {
            if (kpuProvinsiAddressesArray[i].Address == Address) {
                kpuProvinsiAddressesArray[i].name = name;
                kpuProvinsiAddressesArray[i].region = region;
                break;
            }
        }

        emit KPUProvinsiUpdated(Address, name, region);
    }

    function updateKPUKota(
        address Address,
        string calldata name,
        string calldata region
    ) external override onlyKpuKota {
        if (!_kpuKota[Address].isActive) revert KPUKotaNotActive();

        _kpuKota[Address].name = name;
        _kpuKota[Address].region = region;

        // Update in the array too
        for (uint i = 0; i < kpuKotaAddressesArray.length; i++) {
            if (kpuKotaAddressesArray[i].Address == Address) {
                kpuKotaAddressesArray[i].name = name;
                kpuKotaAddressesArray[i].region = region;
                break;
            }
        }

        emit KPUKotaUpdated(Address, name, region);
    }
}
