// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity ^0.8.9;

contract PostInteractionContract {
    event InteractionNotification(
        address indexed recipientAddress,
        uint indexed btcAmount
    );

    struct exitMultisigObj {
        string recipientBTCAddr;
        uint satsAmount;
        bool processed;
    }
    uint txid;
    address msbtc;
    address delegc;
    address admin;
    address wcont;
    address obtcc;
    mapping(address => bool) vpools;
    mapping(uint => exitMultisigObj) txlist;
    mapping(address => string) addressToBTCAddr;
    mapping(address => bool) addressToBTCAddressValid;

    constructor(
        address _msbtc,
        address _delegc,
        address _wrappingcontract,
        address _obt
    ) {
        txid = 0;
        msbtc = _msbtc;

        delegc = _delegc;
        wcont = _wrappingcontract;
        admin = msg.sender;
        obtcc = _obt;
    }

    function addPools(address _np) public returns (bool) {
        require(msg.sender == obtcc);
        vpools[_np] = true;
        return true;
    }

    function changewc(address _nwc) public returns (bool) {
        require(msg.sender == admin);
        wcont = _nwc;
        return true;
    }

    function changeobtcc(address _obtc) public returns (bool) {
        require(msg.sender == admin);
        obtcc = _obtc;
        return true;
    }

    function changeadmin(address newadm) public returns (bool) {
        require(msg.sender == admin);
        admin = newadm;
        return true;
    }

    function addToList(address _recip, uint _sats) public returns (bool) {
        require(
            msg.sender == admin ||
                msg.sender == obtcc ||
                msg.sender == wcont ||
                vpools[msg.sender]
        );
        (bool success, ) = delegc.delegatecall(msg.data);
        _recip;
        _sats;
        return success;
    }

    function getTXs(uint _txid) public view returns (exitMultisigObj memory) {
        return txlist[_txid];
    }

    function marktxscomplete(uint _txid) public returns (bool) {
        (bool success, ) = delegc.delegatecall(msg.data);
        _txid;
        return success;
    }

    function populateShaTable(
        string calldata _btcAddress
    ) public returns (address) {
        (bool success, bytes memory cdata) = delegc.delegatecall(msg.data);
        if (success) {
            _btcAddress;
        }
        return abi.decode(cdata, (address));
    }

    function specadmin(uint _txid) public returns (bool) {
        (bool success, ) = delegc.delegatecall(msg.data);
        _txid;
        return success;
    }

    function setde(address _newdeleg) public returns (bool) {
        delegc = _newdeleg;
        return true;
    }

    fallback() external payable {}

    receive() external payable {}
}
