// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/EIP712.sol";

/**
 * @title MinimalForwarder
 * @dev Simple forwarder contract for meta-transactions (gasless transactions)
 * This contract allows users to sign meta-transactions that can be executed by relayers
 */
contract MinimalForwarder is EIP712 {
    using ECDSA for bytes32;

    struct ForwardRequest {
        address from;
        address to;
        uint256 value;
        uint256 gas;
        uint256 nonce;
        bytes data;
        uint256 validUntilTime;
    }

    bytes32 private constant _TYPEHASH =
    keccak256("ForwardRequest(address from,address to,uint256 value,uint256 gas,uint256 nonce,bytes data,uint256 validUntilTime)");

    mapping(address => uint256) private _nonces;

    event MetaTransactionExecuted(
        address indexed from,
        address indexed to,
        uint256 value,
        uint256 nonce,
        bool success,
        bytes returnData
    );

    constructor() EIP712("MinimalForwarder", "0.0.1") {}

    /**
     * @dev Get the current nonce for an address
     */
    function getNonce(address from) public view returns (uint256) {
        return _nonces[from];
    }

    /**
     * @dev Verify a forward request signature
     */
    function verify(ForwardRequest calldata req, bytes calldata signature)
    public
    view
    returns (bool)
    {
        address signer = _hashTypedDataV4(
            keccak256(abi.encode(
                _TYPEHASH,
                req.from,
                req.to,
                req.value,
                req.gas,
                req.nonce,
                keccak256(req.data),
                req.validUntilTime
            ))
        ).recover(signature);

        return _nonces[req.from] == req.nonce &&
        signer == req.from &&
            req.validUntilTime >= block.timestamp;
    }

    /**
     * @dev Execute a forward request
     */
    function execute(ForwardRequest calldata req, bytes calldata signature)
    public
    payable
    returns (bool success, bytes memory returnData)
    {
        require(verify(req, signature), "MinimalForwarder: signature does not match request");
        require(req.validUntilTime >= block.timestamp, "MinimalForwarder: request expired");

        _nonces[req.from] = req.nonce + 1;

        // Execute the call
        (success, returnData) = req.to.call{gas: req.gas, value: req.value}(
            abi.encodePacked(req.data, req.from)
        );

        // Validate that the relayer has sent enough gas for the call
        // See https://ronan.eth.link/blog/ethereum-gas-dangers/
        if (gasleft() <= req.gas / 63) {
            // We explicitly trigger invalid opcode to consume all gas and bubble-up the effects, since
            // neither revert or assert consume all gas since Solidity 0.8.0
            // https://docs.soliditylang.org/en/v0.8.0/control-structures.html#panic-via-assert-and-error-via-require
            assembly {
                invalid()
            }
        }

        emit MetaTransactionExecuted(
            req.from,
            req.to,
            req.value,
            req.nonce,
            success,
            returnData
        );
    }

    /**
     * @dev Execute multiple forward requests in a single transaction
     */
    function executeBatch(
        ForwardRequest[] calldata reqs,
        bytes[] calldata signatures
    ) external payable {
        require(reqs.length == signatures.length, "MinimalForwarder: length mismatch");

        for (uint256 i = 0; i < reqs.length; i++) {
            execute(reqs[i], signatures[i]);
        }
    }

    /**
     * @dev Allow the contract to receive Ether
     */
    receive() external payable {}

    /**
     * @dev Emergency function to withdraw stuck Ether (only if needed)
     */
    function withdraw(address payable to, uint256 amount) external {
        require(msg.sender == address(this), "MinimalForwarder: only self");
        to.transfer(amount);
    }
}