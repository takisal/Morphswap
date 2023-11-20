// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity ^0.8.12;

import "./IERC20.sol";

contract PostInteractionContract {
    event InteractionNotification(
        address indexed recipAddress,
        uint indexed btcAmount
    );

    struct ExitMultisigObject {
        string recipientBtcAddr;
        uint satsAmount;
        bool procd;
    }
    uint txID;
    address msBTC;
    address delegateCaller;
    address admin;

    IERC20 msBTCInterface;

    mapping(address => string) addressToBTCAddress;
    mapping(uint => exitmultisigObj) txList;
    mapping(address => bool) vPools;

    constructor(address _msBTC) {
        txID = 0;
        msBTC = _msBTC;
        msBTCInterface = IERC20(_msBTC);
    }

    function addToList(address _recip, uint _sats) public returns (bool) {
        //V2 make exitmultisigObj take address instead of string
        //allow conversion of address to string afterwards by invoking a function
        txlist[txID] = ExitMultisigObject(
            addressToBTCAddress[_recip],
            _sats,
            false
        );
        txID++;

        return true;
    }

    function getTXs(
        uint _txID
    ) public view returns (ExitMultisigObject memory) {
        return txlist[_txID];
    }

    function markTXsComplete(uint _txID) public returns (bool) {
        txlist[_txID].procd = true;
        return true;
    }

    function populateShaTable(
        string calldata _btcAddress
    ) public returns (address) {
        bytes20 hashBTC = bytes20((sha256(bytes(_btcAddress))));
        address decodedAddress = address(hashBTC);
        addressToBTCAddress[decodedAddress] = _btcAddress;
        return decodedAddress;
    }

    function specadmin(uint _txID) public returns (bool) {
        (bool success, bytes memory cdata) = delegateCaller.delegatecall(
            msg.data
        );
        return success;
    }

    fallback() external payable {}

    receive() external payable {}
}
