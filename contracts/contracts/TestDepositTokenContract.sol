pragma solidity >=0.7.0 <0.8.0;
pragma experimental ABIEncoderV2;

import "OpenZeppelin/openzeppelin-contracts@3.3.0/contracts/token/ERC777/ERC777.sol";

contract TestDepositTokenContract is ERC777 {
    constructor() ERC777("SDT", "SDT", new address[](0)) {
        _mint(msg.sender, 1000000, "", "");
    }
}
