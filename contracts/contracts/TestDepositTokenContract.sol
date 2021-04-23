// SPDX-License-Identifier: MIT

pragma solidity =0.8.4;
pragma experimental ABIEncoderV2;

import {
    ERC777
} from "OpenZeppelin/openzeppelin-contracts@3.3.0/contracts/token/ERC777/ERC777.sol";

contract TestDepositTokenContract is ERC777 {
    constructor() ERC777("SDT", "SDT", new address[](0)) {
        _mint(msg.sender, 1000000, "", "");
    }
}
