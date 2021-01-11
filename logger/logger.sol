//SPDX-License-Identifier: MIT-License
pragma solidity ^0.8.0;

/**
 * @title Logger
 * @dev Logger & retrieve value in a variable
 */
contract Logger {
    
    // event for logging
    event Log(address indexed sender, string data);

    /**
     * @dev dataLog create a log for the sender
     * @param data log data to be emitted
     */
    function dataLog(string memory data) public {
        emit Log(msg.sender, data);
    }

}