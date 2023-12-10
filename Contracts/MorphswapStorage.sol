// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity 0.8.12;

import "./IERC20.sol";
import "./AssetPool.sol";
import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract MorphswapStorage {
    //Events
    event Buy(
        uint indexed otherChainIdNum,
        uint indexed pairIdNum,
        uint saleAmount,
        uint poolBalance,
        address chain2Wallet,
        uint indexed transactionNumber,
        uint tipAmount,
        uint preTipAmount,
        uint methodID
    );
    event MultiChainBuy(
        uint indexed otherChainIdNumber,
        uint indexed pairIdNum,
        uint64 internalFinalChainNumber,
        uint finalPairNumber,
        uint saleAmount,
        uint poolBalance,
        address chain2Wallet,
        uint indexed transactionNumber,
        uint tipAmountQuad,
        uint methodID
    );
    event NewPair(
        uint indexed otherChainIdNumber,
        uint indexed pairIdNumber,
        address chain2Asset,
        address chain2Wallet,
        uint transactionNumber,
        uint indexed newPairTransaction,
        uint methodID
    );
    event FinishedNewPair(
        uint indexed pairId,
        address tca,
        uint _txno,
        address oca,
        uint oc,
        uint tipAmount,
        uint pretipAmount,
        uint methodId
    );
    event AcknowledgedFinishedPair(
        uint indexed pairId,
        uint8 oc,
        address tca,
        address oca
    );
    event SingleLiq(
        uint indexed pairId,
        address indexed c1a,
        address liqProvider,
        uint c1aAdded,
        uint newLpTokens,
        uint prevLpSupply,
        uint indexed blockNo,
        uint methodId
    );
    event FinishedSwap(
        uint indexed pairId,
        uint indexed successfullyFinalizedTx,
        address tcw,
        uint methodId
    );
    event AutoDoubleLiq(
        uint indexed otherChainIdNum,
        uint pairIdNum,
        uint liqAmount,
        uint prevTotalLiq,
        address c2w,
        uint indexed txNumber,
        uint tipAmount,
        uint pretipAmount,
        uint methodId
    );
    event FinishAutoLiq(
        uint indexed pairId,
        uint liqAmount,
        uint prevTotalLiq,
        uint genesisTXNumber,
        address swapminer
    );
    event ManualDoubleLiq(
        uint indexed otherChainIdNum,
        uint pairIdNum,
        address c2a,
        address c2w,
        uint blockNo,
        uint indexed txNumber,
        uint methodId
    );
    event FinishedLiq(
        uint indexed otherChainIdNum,
        uint pairIdNum,
        address tcaPoolAddress,
        address tcw,
        address ocw,
        uint indexed txNumber,
        uint blockNo,
        uint64 tipRatio,
        uint liqProvided,
        uint prevLiqTotal,
        uint methodid
    );
    event CompletedEscrow(
        uint indexed pairId,
        uint liqAmount,
        uint prevTotalLiq
    );
    event removeBothSidesLiq(
        uint indexed pairId,
        address c2w,
        uint sentLp,
        uint totalLp,
        uint txNumber,
        uint blockNo,
        uint methodId
    );
    event RequestFulfilled(bytes32 indexed requestId, bytes indexed data);
    event RequestMultipleFulfilled(bytes32 indexed requestId);
    event Failed(uint8 indexed icId, uint indexed rtxNumber);
    event CancelledEscrow(uint indexed pairid, address tcw);
    //Structs
    struct PoolPair {
        address thisChainAsset;
        address thisChainPool;
        uint otherChain;
        uint8 iCID;
        address otherChainAsset;
        uint64 pairID;
        bool isValid;
    }

    struct EscrowLog {
        uint escrowBalance;
        address escrowKey;
    }
    struct Proposal {
        uint proposalType;
        uint newValue;
        uint votes;
        uint validUntil;
    }
    struct TXObject {
        uint8 methodID;
        uint8 internalStartChainID;
        uint8 internalEndChainID;
        uint64 pairID;
        address finalchainWallet;
        uint64 secondpairID;
        address firstchainAsset;
        address finalchainAsset;
        uint64 quadrillionRatio;
        uint64 quadrillionTipRatio;
        uint128 rtxNumber;
        bool alternateFee;
    }
    //-Compiler structs
    struct StackTooDeepAvoider1 {
        uint64 pairID;
        uint preTransferBalance;
        uint preTipAmount;
        uint tipAmount;
        address chain2Wallet;
        uint64 secondPairID;
        uint _ICID;
        bool alternateFee;
        bool multiChainHop;
        uint chain1AssetAmount;
        address chain1Asset;
        uint chain2;
        address chain2Asset;
        uint128 rTXNumber;
        uint64 convertedPairID;
    }
    struct StackTooDeepAvoider2 {
        uint64 pairID;
        address otherChainWallet;
        address thisChainPool;
        uint otherChain;
        uint _ICID;
        uint totalValue;
        uint128 sentTipam;
        uint64 tipRatioSend;
        uint128 currentRTXNumber;
        uint64 ratioSend;
    }
    struct StackTooDeepAvoider3 {
        bytes32 _requestId;
        uint8 methodID;
        uint8 internalStartChainID;
        uint8 internalEndChainID;
        uint64 pairID;
        address finalChainWallet;
        uint64 secondPairID;
        address firstChainAsset;
        address finalChainAsset;
        uint64 sentRatio;
        uint64 tipRatio;
        uint128 rTXNumber;
        bool paidWithAlt;
        bytes20 swapminer;
    }

    //Configuration variables
    //-Interfaces
    IERC20 _morphswapToken;
    //-Booleans
    bool _alternateFeeActive;
    bool _alternatePriceFeed;
    bool _alternateJobID;
    //-Numbers
    //--Unsigned integers
    //---256-bit unsigned integers
    uint256 _fee;
    uint256 _referralBonusMultiplier;
    uint256 _proposalLifespan;
    uint256 _swapminingFee;
    //-Addresses
    address _admin;
    address _oracle;
    address _morphswapTokenAddress;

    //Standard state variables
    AggregatorV3Interface internal priceFeed;
    AggregatorV3Interface internal priceFeedAlternate;

    //-Numbers
    //--Unsigned integers
    //---64-bit unsigned integers
    uint64 pairTracker;
    //---128-bit unsigned integers
    uint128 defaultTipMultiplier;
    //---256-bit unsigned integers
    uint256 txNumber;
    uint256 alternateTipMultiplier;
    uint256 chainlinkPrice;
    uint256 chainlinkFee;
    //-Bytes
    //--32-bit bytes
    bytes32 jidAlt;
    bytes32 tMRReq;
    //-Booleans
    bool centralContract;
    //-Addresses
    address chainlinkAddress;

    //Delegate contract addresses
    address testingContract;
    address buyContract;
    address buyWithNativeCoinContract;
    address deployNewPoolPairContract;
    address finishPoolPairContract;
    address autoTwoSidedLiquidityContract;
    address manualTwoSidedLiquidityContract;
    address finishLiquidityContract;
    address confirmRemoveBothSidesLiqContract;
    address addSupportedChainsContract;
    address acknowledgeFinishLiquidityContract;
    address governanceContract;
    address singleSidedLiquidityContract;
    address cancelManualEscrowContract;
    address pingContract;

    //Mappings
    mapping(uint => uint) public eCIDToTipMultiplier;
    mapping(address => address[]) public referrerToReferred;
    mapping(uint => address) iCIDToMCPAArray;
    mapping(uint8 => mapping(uint128 => bool)) txProcessed;
    mapping(uint8 => mapping(uint128 => uint)) swapToBeDone;
    mapping(uint => string) icidToRpc;
    mapping(uint => string) public reasonCodeStorage;
    mapping(uint => uint) public errorCodeStorage;
    mapping(uint => bytes) public lowLevelDataStorage;
    mapping(uint => bytes32) iCIDToJID;
    mapping(address => uint[]) addressToProposalsVotedOn;
    mapping(address => mapping(uint => uint)) addressToBallotToVotes;
    mapping(address => uint) delegatedTokensToAddress;
    mapping(address => bool) public oldUser;
    mapping(address => address) public referredToReferrer;
    mapping(uint => mapping(address => bool)) pairIDWaitingForLiqFromTCWallet;
    mapping(uint => bool) supportedChains;
    mapping(uint => mapping(address => EscrowLog)) pID_c1wEscrowlog;
    mapping(uint => mapping(address => mapping(address => PoolPair))) cID_c1A_c2A;
    mapping(uint64 => PoolPair) public idToPair;
    mapping(uint => uint8) chainIDToInternalChainID;
    mapping(uint => uint) internalChainIDToChainID;
    mapping(uint => TXObject) gTXNumberToTXObject;
    mapping(uint => mapping(uint128 => TXObject)) iCIDToRTXNumberToTXObject;
    mapping(uint => uint128) iCIDToLastRTXNumber;
    mapping(uint => uint128) iCIDToNumberOfTXsProcessed;
    mapping(uint => mapping(address => address)) greenlitICIDToAddressMap;
    mapping(address => mapping(address => mapping(address => TXObject))) tCW_C1A_C2A_TXObject;
    mapping(uint128 => uint) rTXToBlockNumber;
    mapping(uint => address) iCIDToAltNCPA;

    //Enums
    enum MethodIDs {
        Empty,
        Buy,
        NewPair,
        DoubleLiquidity,
        SingleSidedLiquidity,
        Swapped,
        ManualDoubleLiquidity,
        FinishDoubleLiquidity,
        FinishNewPair,
        RemoveDoubleLiquidity
    }
    //Arrays
    //-Address arrays
    address[] alternateTipArray;
    address[] mCPAArray;
    //-256-bit unsigned integer arrays
    uint[] supportedChainsList;
    uint[] chainlinkFeeArray;
    //-Struct arrays
    Proposal[] public _ballot;

    //Constants
    //-64-bit unsigned integers
    uint64 constant oneQuadrillion = 1000000000000000;

    //Public variables
    uint8 public internalChainID;
    uint128 public defaultTip;
    uint256 public defaultTipAlternate;
    uint256 public errorCount;
    uint256 public chainID;
}
