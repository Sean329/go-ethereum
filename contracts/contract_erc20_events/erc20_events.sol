// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract ERC20EVENTS {
    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}