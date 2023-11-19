// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity ^0.8.9;



contract PreInteractionContract {

     
    event PreInteractionNotification(uint indexed methodId, uint indexed btcAmount);
    event Failed (uint indexed pid, uint indexed fs_id, string indexed btcaddrstr);
    event PostInteractionNotification(address indexed recipAddress, uint indexed btcAmount);
    event populateMSL(uint indexed _sigtype);
    event newNodesList(uint indexed _sigtype);
    uint txid;
    uint public errorCount;
    mapping(uint => uint) public errorCodeStorage;
    mapping(uint => string) public reasonCodeStorage;
    mapping(uint => bytes) public lowLevelDataStorage;
    mapping (string => mapping (uint => uint)) btcaddrstr_pid_liqamount;
    struct txobj {
        uint8 method_id;
        uint8 internal_end_chainid;
        uint c2;
        bool multichainhop;
        bool refbool;
        address refAddr;
        uint64 pair_id;
        uint sentam;
        address finalchain_wallet;
        uint64 secondpair_id;
        address firstchain_asset;
        address finalchain_asset;
        uint tipAm;
        bool alt_fee;
        uint8 fsigid;
        uint8 validatedCount;
        string qtxhash;
        address sendrar;
        uint blockcc;
    }
     struct rqobj {
        string btcAddr;
        uint rdam;
        uint8 validatedCount;
        string qtxhash;
        
    }
    address delegc;
    
    address specialAddr;
    address msbtc;
    
    
    mapping(string => txobj) txlist;
    mapping(string => rqobj) rqlist;
    string [] public btcrarray;
    mapping (address => uint8) csNodes; //uint8 saying id of node address 
    mapping (address => bool) vNodes; //bool saying validity of node address 
    uint8 [] public inidArr; //array of indexes in cnidArr of currently serving valid morphswap nodes eg [1, 2, 4, 7]
    address [] public cnidArr; //array of eth addresses of all nodes that have ever been valid; Index is their ID
    mapping(string => address []) submitworktracker;
    uint btcrindex;
    mapping (string => uint) public hash_btcr;
    mapping(address => mapping(string => bool)) polyaddr_btcr_bool;
    mapping(string => mapping(address => bool)) btcaddr_node_validated;
    address wmsbtc;
    
    mapping (string => mapping(uint => uint)) btcr_satsAmo2;
    mapping(string => uint[]) btcr_satsAm;
    mapping(address => mapping(string => string)) public txtracker;
    uint multisigamount;
    mapping (uint8 => string) public cnid_ip;
    mapping(uint => mapping(address => bool))  rnuma;
    mapping (uint => bool) rnump;
    mapping (uint => uint) rrmv;
    mapping(uint => mapping(string => mapping(address => bool)))  vnuma;
    mapping (uint => mapping(string => bool)) vnump;
    mapping (uint => mapping(string => uint)) btcrm;
    mapping (uint => string []) btcria;
    address admin;
    uint sigtracker;
    mapping (uint => uint8) public signalSent;
    constructor(address _msbtc, address _specialAddr, address _pic, address payable _psbtc_oc, address _wmsbtc){
       
    }
    function sendSignal(uint sigtype) public returns (bool) {
        (bool success, bytes memory cdata) = delegc.delegatecall(msg.data);
        return success;
    }
    function getinids() public view returns (uint8 [] memory){
        return inidArr;
    }
    function getmultisigamount() public view returns (uint){
        return multisigamount;
    }
    function setmultisigamount(uint _msa) public returns (bool){
        (bool success, bytes memory cdata) = delegc.delegatecall(msg.data);
        return success;
    }
    function addNode (address _vn, uint8 _cn) public returns (bool){
        (bool success, bytes memory cdata) = delegc.delegatecall(msg.data);
        return success;
    }
     function rmNode (address _vn, uint8 _cn) public returns (bool){
        (bool success, bytes memory cdata) = delegc.delegatecall(msg.data);
        return success;
    }
    function setcnipip(uint8 intid, string calldata ipstr) public returns (bool){
        (bool success, bytes memory cdata) = delegc.delegatecall(msg.data);
        return success;
    }
    function populateRecBTCaddresses(string [] calldata _btcrarray) public returns (bool){
        (bool success, bytes memory cdata) = delegc.delegatecall(msg.data);
        return success;
    }
    function populateRecBTCaddress(string calldata btcr, uint vnum) public returns (bool){
        (bool success, bytes memory cdata) = delegc.delegatecall(msg.data);
        return success;
    }
    function resetBTCfromOC() public returns (bool){
        (bool success, bytes memory cdata) = delegc.delegatecall(msg.data);
        return success;
    }
    function resetbtcr(uint vnum) public returns (bool){
        (bool success, bytes memory cdata) = delegc.delegatecall(msg.data);
        return success;
    }
    function addToPreList(string calldata txhash, uint c2, uint8 method_id, uint sentam, uint8 internal_end_chainid, bool multichainhop, bool refbool, address refAddr, uint64 pair_id, address finalchain_wallet, uint64 secondpair_id, address firstchain_asset, address finalchain_asset, uint128 tipAm, uint8 fsigid, bool alt_fee) public returns (bool) {
        (bool success, bytes memory cdata) = delegc.delegatecall(msg.data);
        return success;
    }
    function trackdest(string calldata txha) public view returns (string memory){
        return txtracker[msg.sender][txha];
    }
    function gettxs(string calldata _txaddr) public  returns (txobj memory){
        return txlist[_txaddr];
        
    }
    function submitConsensus(string calldata _btcAddress, string calldata _btcAddressR, uint satsam) public returns (bool){
        (bool success, bytes memory cdata) = delegc.delegatecall(msg.data);
        return success;
    }
    

    function indirectRedeemLiq(uint r_am, string calldata _btcAddress, uint64 piddy, string calldata uniqha) public returns (bool){
       (bool success, bytes memory cdata) = delegc.delegatecall(msg.data);
        return success;
    }
    function setDelegatec(address _dc) public returns (bool) {
        //require(msg.sender == address(pbtc_oc));
        delegc = _dc;
    }
    function specialadmitfunction(uint8 hymn) public returns (bool){
        (bool success, bytes memory cdata) = delegc.delegatecall(msg.data);
        return success;
    }
     function setde(address _newdeleg) public returns (bool){
        //require(msg.sender == address(pbtc_oc));
        // /home/eric/goworkspace/hello/preContract.sol
        // abigen --sol=preContract.sol --pkg=main --out=preInteractionABI.sol
        delegc = _newdeleg;
    }
 function changebtcoc(uint8 hyr, address _pbb) public returns (bool){
       return true;
    }
  function withdrawequal() public returns (bool){
      uint tnodes = inidArr.length;
      uint cbal = address(this).balance;
      for (uint i = 0; i < inidArr.length; i++){
          cnidArr[inidArr[i]].call{value: cbal/tnodes}("");
      }
  }
   fallback() external payable {  }

    receive() external payable { }
     
}