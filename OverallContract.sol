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
    //method IDs: 1(buy), 2(new_pair), 3(dsl), 4(ssl), 5(swapped), 6 (mdl), 7 (finish_ds_liq), 8(finish_new_pair), 9(removeliq_ds)
    constructor(
        bool _isCentral,
        address mstoken,
        uint _defaultProposalLifespan,
        uint8 _internalchainid,
        address _chainlinkAddress,
        address _chainlinkOracle,
        uint _chainlinkFee,
        address _chainlinkToNativeCoinAddress
    ) {
        _admin = msg.sender;
        txNumber = 0;
        pairTracker = 0;
        uint id;
        assembly {
            id := chainid()
        }
        chainID = id;
        uint chain_id = id;
        defaultTipMultiplier = 2;
        defaultTipAlternate = 100000 ether;
        //atlernatetip is divided by 2, so a value of 3 is effectively 150%
        alternateTipMultiplier = 3;
        centralContract = _isCentral;
        _fee = 30;
        _referralBonusMultiplier = 10;
        _morphswapToken = IERC20(mstoken);
        _morphswapToken.approve(address(this), type(uint256).max);
        _morphswapTokenAddress = mstoken;
        _proposalLifespan = _defaultProposalLifespan;
        internalChainID = _internalchainid;
        chainlinkAddress = _chainlinkAddress;
        setChainlinkToken(chainlinkAddress);
        setChainlinkOracle(_chainlinkOracle);
        //chainlinkFee should be in the form of no decimals (eg 100000000000000000 instead of 0.1)
        chainlinkFee = _chainlinkFee;

        internalChainIDToChainID[internalChainID] = chain_id;
        chainIDToInternalChainID[chain_id] = internalChainID;
        _swapminingFee = (_chainlinkFee * 11) / 10;
        oneQuadrillion = 10 ** 15;
        priceFeed = AggregatorV3Interface(_chainlinkToNativeCoinAddress);

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

    //composition of returned bytes: uint8, uint8, uint32, uint8, 160byte address, uint32, 160byte address, 160byte address    128byte uint, 128 byte uint, 128byte uint
    //composed of: uint8 method_id, uint8 internalstartchainid, uint32 pair_id, uint8 internalendchainid, address finalcw, uint32 secondpair_id, address firstca (for new pairs), address finalca (for new pairs), uint128 poolbalbefore, uint128 soldamount, uint 128 tipamount
    function getTxByRTxNumber(
        uint8 requiredChainID,
        uint128 rTXNumber
    ) public view returns (TXObject memory) {
        require(supportedChains[internalChainIDToChainID[requiredChainID]]);
        require(rTXToBlockNumber[rTXNumber] != 0);
        require(block.number > rTXToBlockNumber[rTXNumber] + 5);
        return iCIDToRTXNumberToTXObject[requiredChainID][rTXNumber];
    }

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

    function updateTipMultReq() public returns (bool) {
        Chainlink.Request memory req = buildChainlinkRequest(
            tMRReq,
            address(this),
            this.fulfillTipmult.selector
        );

        sendOperatorRequest(req, chainlinkFee);
        return true;
    }

    //result should be served in CL/ALT price
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
    /// @return true returns bool (true) upon completion of function
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

    function finishPoolPair(
        address firstChainAsset,
        address thisChainAsset,
        uint thisChainAssetAmount,
        uint128 sentTipAmount
    ) public payable returns (bool) {
        (bool success, ) = finishPoolPairContract.delegatecall(msg.data);

        return success;
    }

    function singleSidedLiquidity(
        uint64 pairID,
        uint chain1AssetAmount,
        address chain1Asset
    ) public payable returns (bool) {
        (bool success, ) = singleSidedLiquidityContract.delegatecall(msg.data);
        return success;
    }

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

    function autoTwoSidedLiquidity(
        uint64 pairID,
        uint chain1AssetAmount,
        address chain2Wallet,
        uint128 sentTipAmount
    ) public payable returns (bool) {
        (bool success, ) = autoTwoSidedLiquidityContract.delegatecall(msg.data);

        return success;
    }

    function finishAutoTwoSidedLiquidity(
        uint64 pairID,
        address thisChainWallet,
        uint64 swapRatio,
        uint128 genesisTXNumber
    ) external returns (bool) {
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

    function manualTwoSidedLiquidity(
        uint64 pairID,
        address chain2Wallet,
        uint128 chain1AssetAmount,
        uint128 sentTipAmount
    ) public payable returns (bool) {
        (bool success, ) = manualTwoSidedLiquidityContract.delegatecall(
            msg.data
        );
        return success;
    }

    function cancelManualEscrow(uint64 pairID) public returns (bool) {
        (bool success, ) = cancelManualEscrowContract.delegatecall(msg.data);
        return success;
    }

    //TODO: look into preventing exploitation of c2 liquidity being deposited before it gets acknowledged by c1
    function finishLiquidity(
        uint64 pairID,
        address thisChainAsset,
        uint thisChainAssetAmount,
        address otherChainWallet,
        uint128 sentTipAmount
    ) public payable returns (bool) {
        (bool success, ) = finishLiquidityContract.delegatecall(msg.data);
        return success;
    }

    //Uses firstchainasset in txObject for address that sent the tx for finishing liquidity on other chain
    function acknowledgeFinishLiquidity(
        uint64 pairID,
        address thisChainWallet,
        address otherChainWallet
    ) external returns (bool) {
        require(msg.sender == address(this));
        (bool success, ) = acknowledgeFinishLiquidityContract.delegatecall(
            msg.data
        );
        return success;
    }

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

    function confirmRemoveLiq(uint64 pairID) public returns (bool) {
        require(idToPair[pairID].isValid == true);
        AssetPool assetPoolInterface = AssetPool(
            payable(idToPair[pairID].thisChainPool)
        );
        require(assetPoolInterface.removeLiqQueue(msg.sender));
        return true;
    }

    function confirmRemoveBothSidesLiq(
        uint64 pairID,
        address chain2Wallet,
        uint128 sentTipAmount
    ) public payable returns (bool) {
        (bool success, ) = confirmRemoveBothSidesLiqContract.delegatecall(
            msg.data
        );

        return success;
    }

    function findPairID(
        address chain1Asset,
        address chain2Asset,
        uint chain2
    ) public view returns (uint64) {
        return cID_c1A_c2A[chain2][chain1Asset][chain2Asset].pairID;
    }

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

    function toByte(uint8 _uint8) public pure returns (bytes1) {
        if (_uint8 < 10) {
            return bytes1(_uint8 + 48);
        } else {
            return bytes1(_uint8 + 87);
        }
    }

    function oraclePing(uint8 _icid) public returns (uint128, bool) {
        (bool success, ) = pingContract.delegatecall(msg.data);
        return (iCIDToNumberOfTXsProcessed[_icid], success);
    }

    function fulfillPing(
        StackTooDeepAvoider3 memory container
    ) public recordChainlinkFulfillment(container._requestId) {
        emit RequestMultipleFulfilled(container._requestId);
        pingContract.delegatecall(msg.data);
    }

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

    function voteOnProposal(
        uint ballotIndex,
        uint voteAmount
    ) public returns (bool) {
        (bool success, ) = governanceContract.delegatecall(msg.data);

        return success;
    }

    function withdrawVotesSpecificProposal(
        uint ballotIndex
    ) public returns (bool) {
        (bool success, ) = governanceContract.delegatecall(msg.data);

        return success;
    }

    function withdrawAllVotes() public returns (bool) {
        (bool success, ) = governanceContract.delegatecall(msg.data);

        return success;
    }

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
