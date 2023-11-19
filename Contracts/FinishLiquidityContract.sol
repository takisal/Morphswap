// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity ^0.8.12;

import "./IERC20.sol";
import "./AssetPool.sol";
import "./MorphswapStorage.sol";
import "/home/eric/MS_Audit/node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "/home/eric/MS_Audit/node_modules/@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract FinishLiquidityContract is ChainlinkClient, MorphswapStorage {
    using Chainlink for Chainlink.Request;

    function finishLiquidity(
        uint64 pairID,
        address thisChainAsset,
        uint thisChainAssetAmount,
        address otherChainWallet,
        uint128 tipAmount
    ) public payable returns (bool) {
        require(thisChainAssetAmount > 0, "Amount cannot be zero");
        require(
            tipAmount >=
                defaultTip * eCIDToTipMultiplier[idToPair[pairID].otherChain],
            "Declared tip must be equal or more than default tip"
        );
        require(
            msg.value >= tipAmount,
            "Message value must be equal to or more than declared tip"
        );
        require(
            pairIDWaitingForLiqFromTCWallet[pairID][msg.sender],
            "Must be waiting for liquidity for pair from sender"
        );

        StackTooDeepAvoider2 memory container;
        container.pairID = pairID;
        container.otherChainWallet = otherChainWallet;
        container.thisChainPool = idToPair[container.pairID].thisChainPool;
        container.otherChain = idToPair[container.pairID].otherChain;
        container._ICID = idToPair[container.pairID].iCID;
        require(
            thisChainAsset == idToPair[container.pairID].thisChainAsset,
            "declared asset must equal this chain asset"
        );
        iCIDToLastRTXNumber[container._ICID]++;
        txNumber++;
        pairIDWaitingForLiqFromTCWallet[container.pairID][msg.sender] = false;
        container.totalValue = msg.value;
        uint preTipAmount = mCPAArray[container._ICID].balance;
        (bool tipResult, ) = mCPAArray[container._ICID].call{value: tipAmount}(
            ""
        );
        require(tipResult);
        if (thisChainAsset == address(0)) {
            (bool sentResult, ) = idToPair[container.pairID].thisChainPool.call{
                value: container.totalValue - tipAmount
            }("");
            require(sentResult);
            thisChainAssetAmount = container.totalValue - tipAmount;
        } else {
            require(
                IERC20(thisChainAsset).transferFrom(
                    msg.sender,
                    container.thisChainPool,
                    thisChainAssetAmount
                )
            );
        }
        (bool sent, uint addedLP, uint oldLPAmount) = AssetPool(
            payable(container.thisChainPool)
        ).addLiquidity(msg.sender, thisChainAssetAmount);
        require(sent);
        uint ratioRecord = (addedLP * oneQuadrillion) / oldLPAmount;
        container.ratioSend = uint64(ratioRecord);
        require(container.ratioSend == ratioRecord);
        uint tipRatioRecord = (tipAmount * oneQuadrillion) /
            (preTipAmount + tipAmount);
        container.tipRatioSend = uint64(tipRatioRecord);
        require(container.tipRatioSend == tipRatioRecord);

        gTXNumberToTXObject[txNumber - 1] = TXObject(
            7,
            internalChainID,
            uint8(container._ICID),
            container.pairID,
            container.otherChainWallet,
            0,
            msg.sender,
            idToPair[container.pairID].otherChainAsset,
            container.ratioSend,
            container.tipRatioSend,
            iCIDToLastRTXNumber[container._ICID] - 1,
            false
        );
        container.currentRTXNumber =
            iCIDToLastRTXNumber[container._ICID] -
            uint128(1);
        iCIDToRTXNumberToTXObject[container._ICID][
            container.currentRTXNumber
        ] = gTXNumberToTXObject[txNumber - 1];
        rTXToBlockNumber[container.currentRTXNumber] = block.number;
        if (oldUser[msg.sender] == false) {
            oldUser[msg.sender] = true;
        }
        emit FinishedLiq(
            container.otherChain,
            container.pairID,
            container.thisChainPool,
            msg.sender,
            container.otherChainWallet,
            txNumber - 1,
            block.number,
            container.tipRatioSend,
            addedLP,
            oldLPAmount,
            uint256(MethodIDs.FinishDoubleLiquidity)
        );
        return true;
    }
}
