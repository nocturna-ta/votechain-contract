// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.25;

import "./Interfaces.sol";
import "./Context.sol";

contract VotechainBase is IVotechainBase, MultiERC2771Context {
    error OnlyKpuAdmin();
    error VotingNotActive();

    address public override kpuAdmin;
    bool public override votingActive;

    event VotingStatusChanged(bool isActive);
    event TrustedForwarderUpdated(address indexed forwarder, bool trusted);

    modifier onlyKpuAdmin() {
        if (_msgSender() != kpuAdmin) revert OnlyKpuAdmin();
        _;
    }

    modifier votingIsActive() {
        if (!votingActive) revert VotingNotActive();
        _;
    }

    constructor(address[] memory trustedForwarders) MultiERC2771Context(trustedForwarders) {
        kpuAdmin = _msgSender();
    }

    function setVotingStatus(bool status) external override onlyKpuAdmin {
        votingActive = status;
        emit VotingStatusChanged(status);
    }

    function setKpuAdmin(address newAdmin) external override onlyKpuAdmin {
        kpuAdmin = newAdmin;
    }

    /**
     * @dev Add a trusted forwarder for gasless transactions
     * @param forwarder Address of the forwarder to trust
     */
    function addTrustedForwarder(address forwarder) external onlyKpuAdmin {
        _addTrustedForwarder(forwarder);
    }

    /**
     * @dev Remove a trusted forwarder
     * @param forwarder Address of the forwarder to remove
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