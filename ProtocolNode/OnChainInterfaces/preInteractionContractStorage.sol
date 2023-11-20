// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity 0.8.12;

import "./IERC20.sol";
import "./postInteractionContract.sol";
import "./OverallContract.sol";
import "./AssetPool.sol";

contract PreInteractionContractStorage {
    event PreInteractionNotification(
        uint indexed methodId,
        uint indexed btcAmount
    );
    event Failed(
        uint indexed pID,
        uint indexed fsID,
        string indexed btcAddressString
    );
    event PostInteractionNotification(
        address indexed recipientAddress,
        uint indexed btcAmount
    );
    event populateMSL(uint indexed _signalType);
    event newNodesList(uint indexed _signalType);

    mapping(uint => uint) public errorCodeStorage;
    mapping(uint => string) public reasonCodeStorage;
    mapping(uint => bytes) public lowLevelDataStorage;
    mapping(string => mapping(uint => uint)) btcAddressStringToPIDToLiquidityAmount;
    struct TXObject {
        uint8 methodID;
        uint8 internalEndChainid;
        uint chain2;
        bool multiChainHop;
        bool referralBool;
        address refAddress;
        uint64 pairID;
        uint sentAmount;
        address finalChainWallet;
        uint64 secondPairID;
        address firstChainAsset;
        address finalChainAsset;
        uint tipAmount;
        bool alternateFee;
        uint8 fsignatureID;
        uint8 validatedCount;
        string qTXHash;
        address sendRAR;
        uint blockDetected;
    }
    struct RQObject {
        string btcAddress;
        uint rDAM;
        uint8 validatedCount;
        string qtxhash;
    }
    uint txID;
    uint public errorCount;
    address delegateCallAddress;
    uint btcRIndex;
    address wrappedMSBTC;
    uint requiredSignatures;
    address admin;
    uint signalTracker;
    address specialAddress;
    address msBTC;

    IERC20 msBTCInterface;
    IERC20 wrappedMSBTCInterface;
    OverallContract cBTCOverallContract;
    PostInteractionContract postInteractionContract;

    mapping(address => uint8) csNodes;
    mapping(address => bool) vNodes;
    mapping(string => txobj) txList;
    mapping(string => rqobj) rqList;
    mapping(string => address[]) submittedWorkTracker;
    mapping(string => uint) public hadhToBTCR;
    mapping(address => mapping(string => bool)) centralAddressToBTCRToBool;
    mapping(string => mapping(address => bool)) btcAddressToNodeToValidated;
    mapping(string => uint[]) btcrToSatsAmountArray;
    mapping(string => mapping(uint => uint)) btcrToIndexToSatsAmount;
    mapping(address => mapping(string => string)) public txTracker;
    mapping(uint8 => string) public cNIDToIP;
    mapping(uint => mapping(address => bool)) rNumberA;
    mapping(uint => bool) rNumberP;
    mapping(uint => uint) rRemovalCount;
    mapping(uint => mapping(string => mapping(address => bool))) vNumberA;
    mapping(uint => mapping(string => bool)) vNumberP;
    mapping(uint => mapping(string => uint)) btcRMToCount;
    mapping(uint => string[]) btcToRIA;
    mapping(uint => uint8) public signalSent;

    uint8[] public iNIDArray;
    address[] public cNIDArray;
    string[] public btcRArray;
}
