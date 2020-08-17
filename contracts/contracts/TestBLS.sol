pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

import "./Pairing.sol";
import "./BLS.sol";

/**
 * @title BLSTest
 * @dev Testing contract for the BLS library.
 */
contract TestBLS {
    function verify(
        Pairing.G2Point calldata _publicKey,
        bytes calldata _message,
        Pairing.G1Point calldata _signature
    ) external {
        require(BLS.verify(_publicKey, _message, _signature));
    }

    function testPairing2(
        Pairing.G1Point calldata _g1point1,
        Pairing.G2Point calldata _g2point1,
        Pairing.G1Point calldata _g1point2,
        Pairing.G2Point calldata _g2point2
    ) public returns (bool) {
        return Pairing.pairing2(_g1point1, _g2point1, _g1point2, _g2point2);
    }
}
