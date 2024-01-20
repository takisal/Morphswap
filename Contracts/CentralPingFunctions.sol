// SPDX-License-Identifier: MIT

// Morphswap
pragma solidity ^0.8.12;

import "./IERC20.sol";
import "./AssetPool.sol";
import "./MorphswapStorage.sol";
import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract CentralPingFunctions is ChainlinkClient, MorphswapStorage {
    using Chainlink for Chainlink.Request;

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
    ) public payable virtual returns (bool) {
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
    ) public payable virtual returns (bool) {
        (bool success, ) = buyWithNativeCoinContract.delegatecall(msg.data);

        return success;
    }

    function swapFinish(
        uint64 pairID,
        address thisChainWallet,
        uint64 swapRatio,
        uint128 transactionNumber
    ) external virtual returns (bool) {
        return true;
    }

    function acknowledgeFinishLiquidity(
        uint64 pairID,
        address thisChainWallet,
        address otherChainWallet
    ) external virtual returns (bool) {
        require(msg.sender == address(this));
        (bool success, ) = acknowledgeFinishLiquidityContract.delegatecall(
            msg.data
        );

        return success;
    }

    function finishAutoTwoSidedLiquidity(
        uint64 pairID,
        address thisChainWallet,
        uint64 swapRatio,
        uint128 genesisTXNumber
    ) external returns (bool) {
        require(msg.sender == address(this));
        uint checkswaptobedone = swapToBeDone[idToPair[pairID].iCID][
            genesisTXNumber
        ];
        swapToBeDone[idToPair[pairID].iCID][genesisTXNumber]++;
        require(checkswaptobedone == 1);

        uint thisChainAssetAmount = idToPair[pairID].thisChainAsset ==
            address(0)
            ? (idToPair[pairID].thisChainPool.balance * swapRatio) /
                oneQuadrillion
            : ((IERC20(idToPair[pairID].thisChainAsset).balanceOf(
                idToPair[pairID].thisChainPool
            ) * swapRatio) / oneQuadrillion);

        require(thisChainAssetAmount > 0);
        (bool sent, uint addedLP, uint oldLPTS) = AssetPool(
            payable(idToPair[pairID].thisChainPool)
        ).addLiquidity(thisChainWallet, thisChainAssetAmount);
        require(sent);
        emit FinishAutoLiq(
            pairID,
            addedLP,
            oldLPTS,
            genesisTXNumber,
            tx.origin
        );
        return true;
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
}
