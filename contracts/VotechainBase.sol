// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.25;

import "./Interfaces.sol";

contract VotechainBase is IVotechainBase {
    error OnlyKpuAdmin();
    error VotingNotActive();

    address public override kpuAdmin;
    bool public override votingActive;

    event VotingStatusChanged(bool isActive);

    modifier onlyKpuAdmin() {
        if (msg.sender != kpuAdmin) revert OnlyKpuAdmin();
        _;
    }

    modifier votingIsActive() {
        if (!votingActive) revert VotingNotActive();
        _;
    }

    constructor(){
        kpuAdmin = msg.sender;
    }

    function setVotingStatus(bool status) external override onlyKpuAdmin {
        votingActive = status;
        emit VotingStatusChanged(status);
    }

    function setKpuAdmin(address newAdmin) external override onlyKpuAdmin {
        kpuAdmin = newAdmin;
    }




}
