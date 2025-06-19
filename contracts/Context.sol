// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.25;

/**
 * @dev Provides information about the current execution context, including the
 * sender of the transaction and its data. While these are generally available
 * via msg.sender and msg.data, they should not be accessed in such a direct
 * manner, since when dealing with meta-transactions the account sending and
 * paying for execution may not be the actual sender (as far as an application
 * is concerned).
 *
 * This contract is only required for intermediate, library-like contracts.
 */
abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }
}

/**
 * @dev Context variant with ERC2771 support.
 */
abstract contract ERC2771Context is Context {
    /// @custom:oz-upgrades-unsafe-allow state-variable-immutable
    address private immutable _trustedForwarder;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor(address trustedForwarder) {
        _trustedForwarder = trustedForwarder;
    }

    function isTrustedForwarder(address forwarder) public view virtual returns (bool) {
        return forwarder == _trustedForwarder;
    }

    function _msgSender() internal view virtual override returns (address sender) {
        if (isTrustedForwarder(msg.sender)) {
            // The assembly code is more direct than the Solidity version using `abi.decode`.
            /// @solidity memory-safe-assembly
            assembly {
                sender := shr(96, calldataload(sub(calldatasize(), 20)))
            }
        } else {
            return super._msgSender();
        }
    }

    function _msgData() internal view virtual override returns (bytes calldata) {
        if (isTrustedForwarder(msg.sender)) {
            return msg.data[:msg.data.length - 20];
        } else {
            return super._msgData();
        }
    }
}

/**
 * @dev Multi-forwarder version for flexibility
 */
abstract contract MultiERC2771Context is Context {
    mapping(address => bool) private _trustedForwarders;
    address[] private _trustedForwardersList;

    event TrustedForwarderAdded(address indexed forwarder);
    event TrustedForwarderRemoved(address indexed forwarder);

    constructor(address[] memory trustedForwarders) {
        for (uint256 i = 0; i < trustedForwarders.length; i++) {
            _addTrustedForwarder(trustedForwarders[i]);
        }
    }

    function isTrustedForwarder(address forwarder) public view virtual returns (bool) {
        return _trustedForwarders[forwarder];
    }

    function getTrustedForwarders() public view returns (address[] memory) {
        return _trustedForwardersList;
    }

    function _addTrustedForwarder(address forwarder) internal {
        if (!_trustedForwarders[forwarder]) {
            _trustedForwarders[forwarder] = true;
            _trustedForwardersList.push(forwarder);
            emit TrustedForwarderAdded(forwarder);
        }
    }

    function _removeTrustedForwarder(address forwarder) internal {
        if (_trustedForwarders[forwarder]) {
            _trustedForwarders[forwarder] = false;

            // Remove from array
            for (uint256 i = 0; i < _trustedForwardersList.length; i++) {
                if (_trustedForwardersList[i] == forwarder) {
                    _trustedForwardersList[i] = _trustedForwardersList[_trustedForwardersList.length - 1];
                    _trustedForwardersList.pop();
                    break;
                }
            }
            emit TrustedForwarderRemoved(forwarder);
        }
    }

    function _msgSender() internal view virtual override returns (address sender) {
        if (isTrustedForwarder(msg.sender)) {
            // The assembly code is more direct than the Solidity version using `abi.decode`.
            /// @solidity memory-safe-assembly
            assembly {
                sender := shr(96, calldataload(sub(calldatasize(), 20)))
            }
        } else {
            return super._msgSender();
        }
    }

    function _msgData() internal view virtual override returns (bytes calldata) {
        if (isTrustedForwarder(msg.sender)) {
            return msg.data[:msg.data.length - 20];
        } else {
            return super._msgData();
        }
    }
}