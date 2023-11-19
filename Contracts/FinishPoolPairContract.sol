// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity ^0.8.12;

import "./IERC20.sol";
import "./AssetPool.sol";
import "./MorphswapStorage.sol";
import "/home/eric/MS_Audit/node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "/home/eric/MS_Audit/node_modules/@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract FinishPoolPairContract is ChainlinkClient, MorphswapStorage {
    using Chainlink for Chainlink.Request;

    function finishPoolPair(
        address firstChainAsset,
        address thisChainAsset,
        uint thisChainAssetAmount,
        uint128 tipAmount
    ) public payable returns (bool) {
        require(msg.value >= tipAmount);

        require(thisChainAssetAmount > 0);
        txNumber++;

        TXObject memory poolGenesisTX = tCW_C1A_C2A_TXObject[msg.sender][
            firstChainAsset
        ][thisChainAsset];
        require(poolGenesisTX.pairID != 0);
        uint _ICID = poolGenesisTX.internalStartChainID;
        require(
            tipAmount >=
                defaultTip *
                    eCIDToTipMultiplier[internalChainIDToChainID[_ICID]]
        );
        uint preTipAmount = mCPAArray[_ICID].balance;
        (bool tipResult, ) = mCPAArray[_ICID].call{value: tipAmount}("");
        require(tipResult);
        iCIDToLastRTXNumber[_ICID]++;

        require(idToPair[poolGenesisTX.pairID].isValid != true);
        AssetPool _tcapInterface = new AssetPool(
            thisChainAsset,
            poolGenesisTX.pairID,
            false
        );
        if (thisChainAsset == address(0)) {
            thisChainAssetAmount = msg.value - tipAmount;
            (bool sendresult, ) = address(_tcapInterface).call{
                value: thisChainAssetAmount
            }("");
            require(sendresult);
        } else {
            require(
                IERC20(thisChainAsset).transferFrom(
                    msg.sender,
                    address(_tcapInterface),
                    thisChainAssetAmount
                )
            );
        }
        (bool sent, , ) = _tcapInterface.addLiquidity(
            msg.sender,
            thisChainAssetAmount
        );
        require(sent);
        require(
            cID_c1A_c2A[
                internalChainIDToChainID[poolGenesisTX.internalStartChainID]
            ][thisChainAsset][firstChainAsset].isValid != true
        );
        require(
            uint64((tipAmount * oneQuadrillion) / (preTipAmount + tipAmount)) ==
                (tipAmount * oneQuadrillion) / (preTipAmount + tipAmount)
        );
        gTXNumberToTXObject[txNumber - 1] = TXObject(
            8,
            internalChainID,
            uint8(_ICID),
            poolGenesisTX.pairID,
            address(0),
            0,
            firstChainAsset,
            thisChainAsset,
            0,
            uint64((tipAmount * oneQuadrillion) / (preTipAmount + tipAmount)),
            iCIDToLastRTXNumber[_ICID] - 1,
            false
        );
        uint128 currentRTXNumber = iCIDToLastRTXNumber[_ICID] - uint128(1);
        iCIDToRTXNumberToTXObject[_ICID][
            currentRTXNumber
        ] = gTXNumberToTXObject[txNumber - 1];
        idToPair[poolGenesisTX.pairID] = PoolPair(
            thisChainAsset,
            address(_tcapInterface),
            internalChainIDToChainID[poolGenesisTX.internalStartChainID],
            poolGenesisTX.internalStartChainID,
            firstChainAsset,
            poolGenesisTX.pairID,
            true
        );
        cID_c1A_c2A[
            internalChainIDToChainID[poolGenesisTX.internalStartChainID]
        ][thisChainAsset][firstChainAsset] = idToPair[poolGenesisTX.pairID];
        rTXToBlockNumber[currentRTXNumber] = block.number;
        if (oldUser[msg.sender] == false) {
            oldUser[msg.sender] = true;
        }

        emit FinishedNewPair(
            poolGenesisTX.pairID,
            thisChainAsset,
            txNumber - 1,
            firstChainAsset,
            _ICID,
            tipAmount,
            preTipAmount,
            uint256(MethodIDs.FinishNewPair)
        );

        return true;
    }
}
