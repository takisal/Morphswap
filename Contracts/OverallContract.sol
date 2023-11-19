// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity 0.8.12;

import "./IERC20.sol";
import "./AssetPool.sol";
import "./MorphswapStorage.sol";
import "/home/eric/MS_Audit/node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "/home/eric/MS_Audit/node_modules/@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract OverallContract is ChainlinkClient, MorphswapStorage {
    using Chainlink for Chainlink.Request;

    //Pair IDs are canon across all chains. They are suffixed by the internal chain id (eg 115002 for pair 115 on chain internal id 2)
    //Method IDs: 1(Buy), 2(New Pair), 3(Begin DSL), 4(SSL), 5(Swapped), 6 (MDL), 7 (Finish DSL), 8(Finish New Pair), 9(Remove DSL)
    constructor(
        bool _isCentral,
        address _msToken,
        uint _defaultProposalLifespan,
        uint8 _internalchainid,
        address _chainlinkAddress,
        address _chainlinkOracle,
        uint _chainlinkFee,
        address _chainlinkToNativeCoinAddress
    ) {
        uint id;
        assembly {
            id := chainid()
        }
        chainID = id;
        uint chain_id = id;

        _fee = 30;
        _referralBonusMultiplier = 10;
        _admin = msg.sender;
        _morphswapToken = IERC20(_msToken);
        _morphswapToken.approve(address(this), type(uint256).max);
        _morphswapTokenAddress = _msToken;
        _proposalLifespan = _defaultProposalLifespan;

        txNumber = 0;
        pairTracker = 0;
        oneQuadrillion = 10 ** 15;
        defaultTipMultiplier = 2;
        defaultTipAlternate = 100000 ether;

        centralContract = _isCentral;
        internalChainID = _internalchainid;
        chainlinkAddress = _chainlinkAddress;
        setChainlinkToken(chainlinkAddress);
        setChainlinkOracle(_chainlinkOracle);
        internalChainIDToChainID[internalChainID] = chain_id;
        chainIDToInternalChainID[chain_id] = internalChainID;

        priceFeed = AggregatorV3Interface(_chainlinkToNativeCoinAddress);
        //chainlinkFee should be in the form of no decimals (eg 100000000000000000 instead of 0.1)
        chainlinkFee = _chainlinkFee;
        _swapminingFee = (_chainlinkFee * 11) / 10;

        //atlernatetip is divided by 2, so a value of 3 is effectively 150%
        alternateTipMultiplier = 3;

        //Avoids the errors that would result from trying to add own chain to supported chains list (only relevant for central chain)
        if (centralContract) {
            chainIDToInternalChainID[chain_id] = 0;
            internalChainIDToChainID[0] = chain_id;
            iCIDToJID[0] = bytes32(bytes("NOJID"));
            supportedChainsList.push(chain_id);
            mCPAArray.push(address(0));
            iCIDToMCPAArray[0] = address(0);
            alternateTipArray.push(address(0));
        }
    }

    function initializeDekegates(
        address _testingContract,
        address _buyContract,
        address _buyWithNativeCoinContract,
        address _acknowledgeFinishLiquidityContract,
        address _addSupportedChainsContract,
        address _autoTwoSidedLiquidityContract,
        address _confirmRemoveBothSidesLiqContract,
        address _deployNewPoolPairContract,
        address _finishLiquidityContract,
        address _finishPoolPairContract,
        address _manualTwoSidedLiquidityContract,
        address _governanceContract,
        address _pingContract,
        address _singleSidedLiquidityContract,
        address _cancelManualEscrowContract
    ) public {
        testingContract = _testingContract;
        buyContract = _buyContract;
        buyWithNativeCoinContract = _buyWithNativeCoinContract;
        acknowledgeFinishLiquidityContract = _acknowledgeFinishLiquidityContract;
        addSupportedChainsContract = _addSupportedChainsContract;
        autoTwoSidedLiquidityContract = _autoTwoSidedLiquidityContract;
        confirmRemoveBothSidesLiqContract = _confirmRemoveBothSidesLiqContract;
        deployNewPoolPairContract = _deployNewPoolPairContract;
        finishLiquidityContract = _finishLiquidityContract;
        finishPoolPairContract = _finishPoolPairContract;
        manualTwoSidedLiquidityContract = _manualTwoSidedLiquidityContract;
        governanceContract = _governanceContract;
        pingContract = _pingContract;
        singleSidedLiquidityContract = _singleSidedLiquidityContract;
        cancelManualEscrowContract = _cancelManualEscrowContract;
    }

    //================================================================================
    //
    //
    //Administrative functions
    //
    //
    //================================================================================
    function changeContractAddress(uint cn, address ca) public {
        require(msg.sender == _admin);
        testingContract.delegatecall(msg.data);
    }

    function greenLightAddress(uint8 _icid, address gla) public returns (bool) {
        greenlitICIDToAddressMap[_icid][gla] = msg.sender;
    }

    function getPoolAmount(uint64 pID) public view returns (uint) {
        if (
            idToPair[pID].thisChainAsset == address(0) && idToPair[pID].isValid
        ) {
            return idToPair[pID].thisChainPool.balance;
        } else {
            return (
                IERC20(idToPair[pID].thisChainAsset).balanceOf(
                    idToPair[pID].thisChainPool
                )
            );
        }
    }

    function withdrawAsset(
        bool isCoin,
        address ercAddress,
        uint amountToWithdraw
    ) public returns (bool) {
        require(msg.sender == _admin);

        if (isCoin) {
            (bool sent, ) = msg.sender.call{value: address(this).balance}("");
            require(sent);
        } else {
            IERC20 ercToken = IERC20(ercAddress);
            require(
                ercToken.transfer(msg.sender, amountToWithdraw),
                "Failed to send asset"
            );
        }
        return true;
    }

    function changeJIDalt(string memory altjid) public returns (bool) {
        require(msg.sender == _admin);
        jidAlt = bytes32(bytes(altjid));
        return true;
    }

    function changeTMRReq(string memory tMR) public returns (bool) {
        require(msg.sender == _admin);
        tMRReq = bytes32(bytes(tMR));
        return true;
    }

    function settipmults(uint eCID, uint tipMultiplier) public returns (bool) {
        require(msg.sender == _admin);
        eCIDToTipMultiplier[eCID] = tipMultiplier;
        return true;
    }

    function changeAltMult(uint _altmult) public returns (bool) {
        require(msg.sender == _admin);
        alternateTipMultiplier = _altmult;
        return true;
    }

    function addAltNCPools(
        uint _ICID,
        address altNCPoolAddress
    ) public returns (bool) {
        require(msg.sender == _admin);
        iCIDToAltNCPA[_ICID] = altNCPoolAddress;
        return true;
    }

    function changeCLtoken(address chainlinkAddress) public returns (bool) {
        require(msg.sender == _admin);
        setChainlinkToken(chainlinkAddress);
        return true;
    }

    function changeCLtoNC(address cLToNCAddress) public returns (bool) {
        require(msg.sender == _admin);
        priceFeed = AggregatorV3Interface(cLToNCAddress);
        return true;
    }

    function changeAlternatetoNC(address altToNCAddress) public returns (bool) {
        require(msg.sender == _admin);
        priceFeedAlternate = AggregatorV3Interface(altToNCAddress);
        return true;
    }

    function changeSupportedChainCLfee(
        uint _ICID,
        uint fee
    ) public returns (bool) {
        require(msg.sender == _admin);
        chainlinkFeeArray[_ICID] = fee;
        return true;
    }

    function changeOracle(address oracleAddress) public returns (bool) {
        require(msg.sender == _admin);
        setChainlinkOracle(oracleAddress);
        return true;
    }

    function changeCLfee(uint newFee) public returns (bool) {
        require(msg.sender == _admin);
        chainlinkFee = newFee;
        _swapminingFee = (newFee * 11) / 10;
        return true;
    }

    function changeSMfee(uint newFee) public returns (bool) {
        require(msg.sender == _admin);
        _swapminingFee = newFee;
        chainlinkFee = (newFee * 10) / 11;
        return true;
    }

    function setAdmin(address newAdmin) public returns (bool) {
        require(msg.sender == _admin);
        _admin = newAdmin;
        return true;
    }

    function setDefaultTipMultiplier(
        uint128 newTipMultiplier
    ) public returns (bool) {
        require(msg.sender == _admin);
        defaultTipMultiplier = newTipMultiplier;
        return true;
    }

    function setOracleAddress(address newOracle) public returns (bool) {
        require(msg.sender == _admin);
        _oracle = newOracle;
        return true;
    }

    function addSupportedChains(
        uint supportedChain,
        string memory jobID,
        address otherChainMorphswap
    ) public returns (bool) {
        (bool success, ) = addSupportedChainsContract.delegatecall(msg.data);

        return success;
    }

    function setVotingTime(uint newProposalLifespan) public returns (bool) {
        require(msg.sender == _admin);
        _proposalLifespan = newProposalLifespan;
        return true;
    }

    function EMERGENCYsetlastIndexforChain(
        uint8 _ICID,
        uint128 newProcessNumber
    ) public returns (bool) {
        require(msg.sender == _admin);
        iCIDToLastRTXNumber[_ICID] = newProcessNumber;
        return true;
    }

    function EMERGENCYsetlastProcessedforChain(
        uint8 _ICID,
        uint128 newProcessNumber
    ) public returns (bool) {
        require(msg.sender == _admin);
        iCIDToNumberOfTXsProcessed[_ICID] = newProcessNumber;
        return true;
    }

    function setReferrer(address _referrer) public returns (bool) {
        require(referredToReferrer[msg.sender] == address(0));
        require(oldUser[_referrer]);
        require(false);
        referredToReferrer[msg.sender] = _referrer;
        referrerToReferred[_referrer].push(msg.sender);
        return true;
    }

    function changeJobId(
        uint8 _ICID,
        string memory jobID
    ) public returns (bool) {
        require(msg.sender == _admin);
        iCIDToJID[_ICID] = bytes32(bytes(jobID));
        return true;
    }

    function updateTipMultReq() public returns (bool) {
        Chainlink.Request memory req = buildChainlinkRequest(
            tMRReq,
            address(this),
            this.fulfillTipmult.selector
        );

        sendOperatorRequest(req, chainlinkFee);
        return true;
    }

    /// @notice Queries a meta-transaction by it's RTX number
    /// @dev Queries a meta-transaction using a chain ID and an RTX number. Requires a transaction to be on-chain for 6 blocks or more to avoid complications resulting from block reorganizations
    /// @param requiredChainID the ID of the destination chain
    /// @param rTXNumber the RTX number of the transaction
    /// @return TXObject returns the transaction info in the form of a TXObject struct
    function getTxByRTxNumber(
        uint8 requiredChainID,
        uint128 rTXNumber
    ) public view returns (TXObject memory) {
        require(supportedChains[internalChainIDToChainID[requiredChainID]]);
        require(rTXToBlockNumber[rTXNumber] != 0);
        require(block.number > rTXToBlockNumber[rTXNumber] + 5);
        return iCIDToRTXNumberToTXObject[requiredChainID][rTXNumber];
    }

    /// @notice Updates the required amount for swapminer tips
    /// @dev Queries the chainlink network for price  of link in terms of native coin, so that the function can then calculate and set the required amount to be used in the calculations done at genesis of a meta-transaction
    /// @return true returns true upon successful completion of function
    function updateTipDefault() public returns (bool) {
        int chainlinkInt;
        (
            ,
            /*uint80 roundID*/ chainlinkInt,
            /*uint startedAt*/
            /*uint timeStamp*/
            /*uint80 answeredInRound*/
            ,
            ,

        ) = priceFeed.latestRoundData();
        //wil return 18 decimals
        require(chainlinkInt > 0);
        chainlinkPrice = uint(chainlinkInt);
        defaultTip =
            uint128(
                chainlinkPrice /
                    (1000000000000000000 / ((chainlinkFee * 11) / 10))
            ) *
            defaultTipMultiplier;
        return true;
    }

    /// @notice Updates the required amount for swapminer tips in alternate currency
    /// @dev Queries the chainlink network for price  of link in terms of the set alternate currency (likely Morphswap's governance token), so that the function can then calculate and set the required amount to be used in the calculations done at genesis of a meta-transaction. Result should be served in CL/ALT price
    /// @return true returns true upon successful completion of function
    function updateTipDefaultAlternate() public returns (bool) {
        require(_alternatePriceFeed);
        int chainlinkInt;
        (
            ,
            /*uint80 roundID*/ chainlinkInt,
            /*uint startedAt*/
            /*uint timeStamp*/
            /*uint80 answeredInRound*/
            ,
            ,

        ) = priceFeedAlternate.latestRoundData();
        //wil return 18 decimals
        require(chainlinkInt > 0);
        chainlinkPrice = uint(chainlinkInt);
        defaultTipAlternate =
            uint128(
                chainlinkPrice /
                    (1000000000000000000 / ((chainlinkFee * 11) / 10))
            ) *
            alternateTipMultiplier;
        return true;
    }

    /// @notice Called by chainlink operator node to fulfill the update request for the tip multiplier
    /// @dev Should only be called be chainlink nodes; updates another chain's set minimum tip muliplier
    /// @param _requestId The request ID of the fulfillment
    /// @param eCID external chain ID
    /// @param _tipMultiplier the tip multiplier for the queried chain
    /// Emits a {RequestFulfilled} event
    function fulfillTipmult(
        bytes32 _requestId,
        uint eCID,
        uint _tipMultiplier
    ) public recordChainlinkFulfillment(_requestId) {
        emit RequestFulfilled(
            _requestId,
            abi.encodePacked(bytes32(_tipMultiplier))
        );
        if (eCID == chainID) {
            defaultTipMultiplier = uint128(_tipMultiplier);
        }
        eCIDToTipMultiplier[eCID] = _tipMultiplier;
    }

    //========================================================================================
    //
    //
    //Primary functions
    //
    //
    //========================================================================================

    /// @notice Initiates a cross-chain swap from a token
    /// @dev Any swap from a token to either a token or a coin should be done by calling this function
    /// @param pairID the pair ID number
    /// @param saleAmount the amount of token being swapped from
    /// @param chain2Wallet the address of the recipient of the swap, on the destination chain
    /// @param tipAmount the amount of the native coin to be swapped, with the proceeds going to the swapminer
    /// @param multiChainHop whether or not this swap goes through more than two chains
    /// @param secondPairID the ID of the second pair, if the swap goes through more than two chains
    /// @param referrer the referrer of the user. Only valid for a wallet's first meta-transaction. A value of 0x00.. is treated as a null value.
    /// @param alternateFee whether or not the swapminer tip will be paid in the alternate tip currency
    /// @param chain2 the final chain of the swap
    /// @return bool returns status of the delegatecall
    /// Emits a {Buy} event
    function buy(
        uint64 pairID,
        uint saleAmount,
        address chain2Wallet,
        uint128 tipAmount,
        bool multiChainHop,
        uint64 secondPairID,
        address referrer,
        bool alternateFee,
        uint chain2
    ) public payable returns (bool) {
        (bool success, ) = buyContract.delegatecall(msg.data);

        return success;
    }

    /// @notice Initiates a cross-chain swap from a coin
    /// @dev Any swap from a coin to either a token or a coin should be done by calling this function. The desired amount to be swapped should be passed in as msg.value
    /// @param pairID the pair ID number
    /// @param chain2Wallet the address of the recipient of the swap, on the destination chain
    /// @param tipAmount the amount of the native coin to be swapped, with the proceeds going to the swapminer
    /// @param multiChainHop whether or not this swap goes through more than two chains
    /// @param secondPairID the ID of the second pair, if the swap goes through more than two chains
    /// @param referrer the referrer of the user. Only valid for a wallet's first meta-transaction. A value of 0x00.. is treated as a null value.
    /// @param alternateFee whether or not the swapminer tip will be paid in the alternate tip currency
    /// @param chain2 the final chain of the swap
    /// @return bool returns status of the delegatecall
    /// Emits a {Buy} event
    function buyWithNativeCoin(
        uint64 pairID,
        address chain2Wallet,
        uint128 tipAmount,
        bool multiChainHop,
        uint32 secondPairID,
        bool referrerBool,
        address referrer,
        bool alternateFee,
        uint chain2
    ) public payable returns (bool) {
        (bool success, ) = buyWithNativeCoinContract.delegatecall(msg.data);

        return success;
    }

    /// @notice Begins the process of creating a pool. Can only be used on the central chain.
    /// @dev creates data to be sent to the peripheral chain in the pool, to be processed by the {finishPoolPair} function, called by the user on the peripheral chain of the pool
    /// @param chain1AssetAmount the amount of this chain's asset the pool will start with, provided by user who calls this function
    /// @param chain1Asset the address of the central chain's asset that comprises half of the pool
    /// @param chain2 the address of the central chain's asset that comprises half of the pool
    /// @param chain2Asset the address of the peripheral chain's asset that comprises half of the pool
    /// @param chain2Wallet the address of the account that will be calling the {finishPoolPair} function to finish the pool's creation
    /// @param tipAmount the amount of the native coin to be swapped, with the proceeds going to the swapminer
    /// @return address returns address of the asset pool on this chain
    /// @return uint returns pair ID of the pool
    /// Emits a {NewPair} event
    function deployNewPoolPair(
        uint chain1AssetAmount,
        address chain1Asset,
        uint chain2,
        address chain2Asset,
        address chain2Wallet,
        uint128 tipAmount
    ) public payable returns (address, uint) {
        (bool success, bytes memory data) = deployNewPoolPairContract
            .delegatecall(msg.data);
        return abi.decode(data, (address, uint));
    }

    /// @notice Marks a pool as valid upon recieving event from other chains
    /// @dev handles events from other chains emitted from function {finishPoolPair} upon completion of a new pool pair. Marks the pool as valid in its struct
    /// @param pID the pair ID number
    /// @return bool returns true upon completion of function
    /// Emits a {AcknowledgedFinishedPair} event

    function markNewPoolPairComplete(uint64 pID) external returns (bool) {
        require(msg.sender == address(this));
        require(idToPair[pID].isValid != true);
        require(
            cID_c1A_c2A[idToPair[pID].otherChain][idToPair[pID].thisChainAsset][
                idToPair[pID].otherChainAsset
            ].isValid != true
        );
        idToPair[pID].isValid = true;
        cID_c1A_c2A[idToPair[pID].otherChain][idToPair[pID].thisChainAsset][
            idToPair[pID].otherChainAsset
        ] = idToPair[pID];

        emit AcknowledgedFinishedPair(
            pID,
            idToPair[pID].iCID,
            idToPair[pID].thisChainAsset,
            idToPair[pID].otherChainAsset
        );
        return true;
    }

    /// @notice Finishes the process of creating a pool. Can only be used on the peripheral chain.
    /// @dev acts upon data received by the peripheral chain in the pool, to the be processed on the central chain by the {markNewPoolPairComplete} function
    /// @param firstChainAsset the address of the central chain's asset that comprises half of the pool
    /// @param thisChainAsset the address of the central chain's asset that comprises half of the pool
    /// @param thisChainAssetAmount the amount of this chain's asset the pool will start with, provided by user who calls this function
    /// @param tipAmount the amount of the native coin to be swapped, with the proceeds going to the swapminer
    /// @return bool returns status of the delegatecall
    /// Emits a {NewPair} event
    function finishPoolPair(
        address firstChainAsset,
        address thisChainAsset,
        uint thisChainAssetAmount,
        uint128 tipAmount
    ) public payable returns (bool) {
        (bool success, ) = finishPoolPairContract.delegatecall(msg.data);

        return success;
    }

    /// @notice Provides liquidity to a pool and credits the provider
    /// @dev Provides liquidity to the pool on behalf of the address that invokes the function, and sends LP (Liquidity Provider) tokens to that user, which can be redeemed for an amount of liquidity proportional to the amount of total total LP tokens.
    /// @param pairID the pair ID number
    /// @param chain1AssetAmount the amount of this chain's asset the user is providing to the pool as liquidity
    /// @param chain1Asset the address of the central chain's asset that comprises half of the pool
    /// @return bool returns status of the delegatecall
    /// Emits a {SingleLiq} event
    function singleSidedLiquidity(
        uint64 pairID,
        uint chain1AssetAmount,
        address chain1Asset
    ) public payable returns (bool) {
        (bool success, ) = singleSidedLiquidityContract.delegatecall(msg.data);
        return success;
    }

    /// @notice Finishes a cross-chain swap
    /// @dev Invoked by the {fulfillPing} function to handle the destination chain's duties with regards to a cross-chain swap. Instructs the appropriate asset pool to send tokens to the approriate address.
    /// @param pairID the pair ID number
    /// @param thisChainWallet the address of the recipient of the swap
    /// @param swapRatio the ratio of the origin chain's swapped assets to the pool size
    /// @param transactionNumber the number of the meta-transaction
    /// @return bool returns status of the delegatecall
    /// Emits a {FinishedSwap} event
    function swapFinish(
        uint64 pairID,
        address thisChainWallet,
        uint64 swapRatio,
        uint128 transactionNumber
    ) external returns (bool) {
        require(msg.sender == address(this));
        uint checkswapToBeDone = swapToBeDone[idToPair[pairID].iCID][
            transactionNumber
        ];
        swapToBeDone[idToPair[pairID].iCID][transactionNumber]++;
        require(checkswapToBeDone == 1);
        AssetPool chain1AssetPool = AssetPool(
            payable(idToPair[pairID].thisChainPool)
        );
        require(chain1AssetPool.sendToUser(swapRatio, thisChainWallet));
        emit FinishedSwap(pairID, transactionNumber, thisChainWallet, 5);
        return true;
    }

    /// @notice Automatically provides two sides of liquidity
    /// @dev Initiates a series of meta-transactions that effectively supplies liquidity to both sides of a pair, creding the specified wallet on both chains
    /// @param pairID the pair ID number
    /// @param chain1AssetAmount the amount of this chain's asset for this pair being supplied to this side of the pair
    /// @param chain2Wallet the address of the recipient of the LP tokens on the destination chain
    /// @param tipAmount the amount of the native coin to be swapped, with the proceeds going to the swapminer
    /// @return bool returns status of the delegatecall
    /// Emits a {AutoDoubleLiq} event
    function autoTwoSidedLiquidity(
        uint64 pairID,
        uint chain1AssetAmount,
        address chain2Wallet,
        uint128 tipAmount
    ) public payable returns (bool) {
        (bool success, ) = autoTwoSidedLiquidityContract.delegatecall(msg.data);

        return success;
    }

    /// @notice Called to finish automatic double-sided liquidity
    /// @dev Invoked by internally in the handling of a received meta-transaction initiated by the {autoTwoSidedLiquidity} function
    /// @param pairID the pair ID number
    /// @param thisChainWallet the address of the recipient of the swap
    /// @param swapRatio the ratio of the origin chain's swapped assets to the pool size
    /// @param genesisTXNumber the number of the meta-transaction in GTX form rather than RTX
    /// @return bool returns true upon successful completion of function
    /// Emits a {FinishAutoLiq} event
    function finishAutoTwoSidedLiquidity(
        uint64 pairID,
        address thisChainWallet,
        uint64 swapRatio,
        uint128 genesisTXNumber
    ) internal returns (bool) {
        require(msg.sender == address(this));
        uint checkSwapToBeDone = swapToBeDone[idToPair[pairID].iCID][
            genesisTXNumber
        ];
        swapToBeDone[idToPair[pairID].iCID][genesisTXNumber]++;
        require(checkSwapToBeDone == 1);

        uint thisChainAssetAmount = idToPair[pairID].thisChainAsset ==
            address(0)
            ? (idToPair[pairID].thisChainPool.balance * swapRatio) /
                oneQuadrillion
            : ((IERC20(idToPair[pairID].thisChainAsset).balanceOf(
                idToPair[pairID].thisChainPool
            ) * swapRatio) / oneQuadrillion);

        require(thisChainAssetAmount > 0);
        (bool sent, uint addedLP, uint oldLPTs) = AssetPool(
            payable(idToPair[pairID].thisChainPool)
        ).addLiquidity(thisChainWallet, thisChainAssetAmount);
        require(sent);
        emit FinishAutoLiq(
            pairID,
            addedLP,
            oldLPTs,
            genesisTXNumber,
            tx.origin
        );
        return true;
    }

    /// @notice Adds the first side manually of double-sided liquidity
    /// @dev Queues liquidity to be added upon release from escrow
    /// @param pairID the pair ID number
    /// @param chain2Wallet the address of the account providing liquidity on the second chain
    /// @param chain1AssetAmount the amount of this chain's asset constituent of this pair to be provided as liquidity to the pool
    /// @param tipAmount the amount of the native coin to be swapped, with the proceeds going to the swapminer
    /// @return bool returns status of the delegatecall
    /// Emits a {ManualDoubleLiq} event
    function manualTwoSidedLiquidity(
        uint64 pairID,
        address chain2Wallet,
        uint128 chain1AssetAmount,
        uint128 tipAmount
    ) public payable returns (bool) {
        (bool success, ) = manualTwoSidedLiquidityContract.delegatecall(
            msg.data
        );
        return success;
    }

    /// @notice Removes assets from escrow, cancelling manual double-sided liquidity
    /// @dev Called by user to remove funds from escrow if they decide not to add funds to second side of the pair
    /// @param pairID the pair ID number
    /// @return bool returns status of the delegatecall
    /// Emits a {CancelledEscrow} event
    function cancelManualEscrow(uint64 pairID) public returns (bool) {
        (bool success, ) = cancelManualEscrowContract.delegatecall(msg.data);
        return success;
    }

    /// @notice Finishes adding double-sided liquidity
    /// @dev Invoked by a user after they have already passed in liquidity on the first chain in the pair using the {manualTwoSidedLiquidity} function
    /// @param pairID the pair ID number
    /// @param thisChainAsset the asset that makes up this chain's side of the pair
    /// @param thisChainAssetAmount the amount of asset being provided to this chain's side of the pair to finish the manual double-sided liquidity provision process
    /// @param otherChainWallet the address of the recipient of the swap, on the destination chain
    /// @param tipAmount the amount of the native coin to be swapped, with the proceeds going to the swapminer
    /// @return bool returns status of the delegatecall
    /// Emits a {FinishedLiq} event
    function finishLiquidity(
        uint64 pairID,
        address thisChainAsset,
        uint thisChainAssetAmount,
        address otherChainWallet,
        uint128 tipAmount
    ) public payable returns (bool) {
        (bool success, ) = finishLiquidityContract.delegatecall(msg.data);
        return success;
    }

    /// @notice Finishes double-sided liquidity process
    /// @dev Should only ever be called by the fulfillPing function when handling a meta-transaction. Results from a call to {finishLiquidity} on other chain
    /// @param pairID the pair ID number
    /// @param thisChainWallet the address that originally initiated the double-sided liquidity process
    /// @param otherChainWallet the address that provided liquidity on the pair's other chain
    /// @return bool returns status of the delegatecall
    function acknowledgeFinishLiquidity(
        uint64 pairID,
        address thisChainWallet,
        address otherChainWallet
    ) external returns (bool) {
        require(msg.sender == address(this));
        //Uses firstchainasset in txObject for address that sent the tx for finishing liquidity on other chain
        (bool success, ) = acknowledgeFinishLiquidityContract.delegatecall(
            msg.data
        );
        return success;
    }

    /// @notice Removes the specified amount of liquidity from the specified pool
    /// @dev Invoked by user to remove liquidity without having to access the pool directly
    /// @param pairID the pair ID number
    /// @param redeemAmount the quantity of LP tokens to be redeemed
    /// @return bool returns true upon successful call to the asset pool
    function removeLiqByProxy(
        uint64 pairID,
        uint redeemAmount
    ) public returns (bool) {
        require(idToPair[pairID].isValid == true);
        require(redeemAmount > 0);
        AssetPool assetPoolInterface = AssetPool(
            payable(idToPair[pairID].thisChainPool)
        );
        require(assetPoolInterface.balanceOf(msg.sender) >= redeemAmount);
        require(
            assetPoolInterface.removeLiqAddToQueue(redeemAmount, msg.sender)
        );
        return true;
    }

    /// @notice Confirms the user's liquidity removal request
    /// @dev Must be invoked after a set amount of blocks to finalize removal of liquidity. This delay is to prevent manipulating pool balances upon incoming transactions
    /// @param pairID the pair ID number
    /// @return bool returns status of the delegatecall
    function confirmRemoveLiq(uint64 pairID) public returns (bool) {
        require(idToPair[pairID].isValid == true);
        AssetPool assetPoolInterface = AssetPool(
            payable(idToPair[pairID].thisChainPool)
        );
        require(assetPoolInterface.removeLiqQueue(msg.sender));
        return true;
    }

    /// @notice Confirms the user's liquidity removal request from both sides of a pool . Allows user to remove liquidity from both sides of a pool without having to interact with both chains directly. Account on other chain must have previously greenlit removal from the invoking address.
    /// @dev Must be invoked after a set amount of blocks to finalize removal of liquidity. This delay is to prevent manipulating pool balances upon incoming transactions. The delay is only based on the local chain.
    /// @param pairID the pair ID number
    /// @param chain2Wallet the address that the other side's liquidity is credited to
    /// @param tipAmount the amount of the native coin to be swapped, with the proceeds going to the swapminer
    /// @return bool returns status of the delegatecall
    function confirmRemoveBothSidesLiq(
        uint64 pairID,
        address chain2Wallet,
        uint128 the
    ) public payable returns (bool) {
        (bool success, ) = confirmRemoveBothSidesLiqContract.delegatecall(
            msg.data
        );

        return success;
    }

    /// @notice Queries the pair ID given two assets and a chain
    /// @dev Invoked by the {fulfillPing} function to handle the destination chain's duties with regards to a cross-chain swap. Instructs the appropriate asset pool to send tokens to the approriate address.
    /// @param chain1Asset the address of the asset making up this chain's side of the pair
    /// @param chain2Asset the address of the asset making up the other chain's side of the pair
    /// @param chain2 the other chain in the pair
    /// @return uint64 the pair ID (is 0 if pair does not exist)
    function findPairID(
        address chain1Asset,
        address chain2Asset,
        uint chain2
    ) public view returns (uint64) {
        return cID_c1A_c2A[chain2][chain1Asset][chain2Asset].pairID;
    }

    /// @notice Converts a series of bytes to a hexadecimal string representation
    /// @dev Returns the string representation of a bytes20 value
    /// @param _bytes20 the data to represent as a string
    /// @return string returns hexadecimal representation of the bytes20 argument
    /// Emits a {FinishedSwap} event
    function bytes20ToString(
        bytes20 _bytes20
    ) public pure returns (string memory) {
        uint8 i = 0;
        bytes memory bytesArray = new bytes((_bytes20.length * 2));
        for (i = 0; i < bytesArray.length - 1; i++) {
            uint8 _f = uint8(_bytes20[i / 2] & 0x0f);
            uint8 _l = uint8(_bytes20[i / 2] >> 4);

            bytesArray[i] = toByte(_l);
            i = i + 1;
            bytesArray[i] = toByte(_f);
        }

        return string(bytesArray);
    }

    /// @notice Converts a byte from uint form to bytes1
    /// @dev Converts a uint8 to bytes1
    /// @param _uint8 the data to convert to a byte, in the form of a uint8
    /// @return bytes1 the resulting byte
    function toByte(uint8 _uint8) public pure returns (bytes1) {
        if (_uint8 < 10) {
            return bytes1(_uint8 + 48);
        } else {
            return bytes1(_uint8 + 87);
        }
    }

    /// @notice Initiates a swapmine attempt
    /// @dev Tells a chainlink node to query a specific chain for information and data about the next transaction to be fulfilled on this chain
    /// @param _ICID the internal chain ID of the chain to query
    /// @return uint128 returns RTX number of transaction being queried
    /// @return bool returns status of the delegatecall
    function oraclePing(uint8 _ICID) public returns (uint128, bool) {
        (bool success, ) = pingContract.delegatecall(msg.data);
        return (iCIDToNumberOfTXsProcessed[_ICID], success);
    }

    /// @notice Handles all incoming meta-transactions
    /// @dev Should only ever be invoked by chainlink nodes when serving a response from an {oraclePing} function call.
    /// @param container the data structure of the meta-transaction, along with request/transaction metadata

    function fulfillPing(
        StackTooDeepAvoider3 memory container
    ) public recordChainlinkFulfillment(container._requestId) {
        emit RequestMultipleFulfilled(container._requestId);
        pingContract.delegatecall(msg.data);
    }

    /// @notice Creates a new proposal. Requires over 2% of total governance token supply.
    /// @dev Called to create a new proposal. Will revert if user stakes less than 2% of total governance token supply; this is to prevent spamming.
    /// @param proposalType the type of the proposal
    /// @param newRate the new value for the variable maintained by the proposal type
    /// @param startingWeight the initial amount of governance tokens to stake in favor of the proposal being made
    /// @return uint returns status of the delegatecall
    function addProposal(
        uint proposalType,
        uint newRate,
        uint startingWeight
    ) public returns (uint) {
        (bool success, bytes memory data) = governanceContract.delegatecall(
            msg.data
        );

        return 1;
    }

    //TODO
    /// @notice Votes on a proposal, either for or against
    /// @dev Invoked by a user to stake a specified amount of governance tokens on a specified proposal
    /// @param ballotIndex the ballot index number to vote on
    /// @param voteAmount the amount of governance tokens to stake for or against the ballot
    /// @return bool returns status of the delegatecall
    function voteOnProposal(
        uint ballotIndex,
        uint voteAmount
    ) public returns (bool) {
        (bool success, ) = governanceContract.delegatecall(msg.data);

        return success;
    }

    /// @notice Withdraws votes from a specific proposal
    /// @dev Withdraws votes from the proposal specified by user, returning the governance tokens back to their account
    /// @param ballotIndex the index of the ballot
    /// @return bool returns status of the delegatecall
    function withdrawVotesSpecificProposal(
        uint ballotIndex
    ) public returns (bool) {
        (bool success, ) = governanceContract.delegatecall(msg.data);

        return success;
    }

    /// @notice Withdraws all of the calling user's votes from all proposals
    /// @dev Iterates though all ballots the user has voted on and withdraws all of the governance token's they've staked from each proposal
    /// @return bool returns status of the delegatecall
    function withdrawAllVotes() public returns (bool) {
        (bool success, ) = governanceContract.delegatecall(msg.data);

        return success;
    }

    /// @notice Queries all active proposals
    /// @dev Iterates over all proposals, adding only active ones to the array to return to caller
    /// @return Proposal[] list of currently active (non-expired) proposals
    /// Emits a {FinishedSwap} event
    function activeProposals()
        public
        view
        returns (Proposal[] memory currentBallot)
    {
        require(_ballot.length > 0);
        uint j = 0;
        for (uint i = 0; i < _ballot.length; i++) {
            if (_ballot[i].validUntil < block.number) {
                currentBallot[j] = (_ballot[i]);
                j++;
            }
        }
        return currentBallot;
    }

    fallback() external payable {}

    receive() external payable {}
}
