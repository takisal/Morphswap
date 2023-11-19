// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity ^0.8.9;



contract PostInteractionContract {

     
    event InteractionNotification(address indexed recipAddress, uint indexed btcAmount);
    uint txid;
    struct exitmultisigObj {
        string recipientBtcAddr;
        uint satsAmount;
        bool procd;
    }
    mapping (address => string) addr_to_btcAddr;
    mapping (address => bool) addr_to_btcAddr_valid;
    address msbtc;
    
    address delegc;
    mapping(uint => exitmultisigObj) txlist;
    address admin;
    address wcont;
    address obtcc;
    mapping (address => bool) vpools;
    constructor(address _msbtc, address _delegc, address _wrappingcontract, address _obt){
        txid = 0;
        msbtc = _msbtc;
        
        delegc = _delegc;
        wcont = _wrappingcontract;
        admin = msg.sender;
        obtcc = _obt;
    }
    function addPools(address _np) public returns (bool) {
        require (msg.sender == obtcc);
        vpools[_np] = true;
        return true;
    }
    function changewc(address _nwc) public returns (bool){
        require(msg.sender == admin);
        wcont = _nwc;
        return true;
    }
    function changeobtcc(address _obtc) public returns (bool){
        require(msg.sender == admin);
        obtcc = _obtc;
        return true;
    }
    function changeadmin(address newadm) public returns (bool){
        require(msg.sender == admin);
        admin = newadm;
        return true;
    }
    function addToList(address _recip, uint _sats) public returns (bool) {
        require(msg.sender == admin || msg.sender == obtcc || msg.sender == wcont || vpools[msg.sender]);
     (bool success, ) = delegc.delegatecall(msg.data);
     _recip;
     _sats;
        return success;
    }
    function gettxs(uint _txid) public view returns (exitmultisigObj memory){
        return txlist[_txid];
    }
    function marktxscomplete(uint _txid) public returns (bool){
        (bool success, ) = delegc.delegatecall(msg.data);
         _txid;
        return success;
    }
     function populateShaTable(string calldata _btcAddress) public returns (address){
         (bool success, bytes memory cdata) = delegc.delegatecall(msg.data);
         if (success){
            _btcAddress;
         }
        return abi.decode(cdata, (address));
    }
    function specadmin(uint _txid) public returns (bool){
        (bool success, ) = delegc.delegatecall(msg.data);
        _txid;
        return success;
    }
    function setde(address _newdeleg) public returns (bool){
        delegc = _newdeleg;
        return true;
    }
     fallback() external payable {  }

    receive() external payable { }
}