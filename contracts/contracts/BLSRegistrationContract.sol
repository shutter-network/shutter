// SPDX-License-Identifier: MIT

pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

import "./BLS.sol";
import "./Pairing.sol";

contract BLSRegistrationContract {
    // Solidity does not allow mappings with value type Pairing.G2Point as it stores the
    // coefficients as arrays. Therefore, we use the following flattened struct.
    struct G2PointFlat {
        uint256 x0;
        uint256 x1;
        uint256 y0;
        uint256 y1;
    }

    event PublicKeyRegistered(
        address indexed sender,
        Pairing.G2Point publicKey
    );

    mapping(address => G2PointFlat) public publicKeys;

    function registerPublicKey(
        Pairing.G2Point calldata _publicKey,
        uint256 _blockNumber,
        bytes32 _blockHash,
        Pairing.G1Point calldata _signature
    ) external {
        require(
            _publicKey.x[0] != 0 ||
                _publicKey.x[1] != 0 ||
                _publicKey.y[0] != 0 ||
                _publicKey.y[1] != 0
        );

        G2PointFlat storage _existingPublicKey = publicKeys[msg.sender];
        require(
            _existingPublicKey.x0 == 0 &&
                _existingPublicKey.x1 == 0 &&
                _existingPublicKey.y0 == 0 &&
                _existingPublicKey.y1 == 0
        );

        require(blockhash(_blockNumber) != bytes32(0));
        require(_blockHash == blockhash(_blockNumber));

        require(
            BLS.verify(_publicKey, abi.encodePacked(_blockHash), _signature)
        );

        publicKeys[msg.sender] = G2PointFlat(
            _publicKey.x[0],
            _publicKey.x[1],
            _publicKey.y[0],
            _publicKey.y[1]
        );
        emit PublicKeyRegistered(msg.sender, _publicKey);
    }

    function getPublicKey(address _owner)
        public
        view
        returns (Pairing.G2Point memory)
    {
        G2PointFlat memory _publicKeyFlat = publicKeys[_owner];
        Pairing.G2Point memory _publicKey;
        _publicKey.x[0] = _publicKeyFlat.x0;
        _publicKey.x[1] = _publicKeyFlat.x1;
        _publicKey.y[0] = _publicKeyFlat.y0;
        _publicKey.y[1] = _publicKeyFlat.y1;
        return _publicKey;
    }
}
