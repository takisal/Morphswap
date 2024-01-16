// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity ^0.8.12;

import "./IERC20.sol";
import "./AssetPool.sol";
import "./MorphswapStorage.sol";
import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract BuyContract is ChainlinkClient, MorphswapStorage {
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
    ) public payable returns (bool) {
        if (referrer != address(0)) {
            require(
                oldUser[msg.sender] == false,
                "Can only set referral on first transaction"
            );
            require(oldUser[referrer], "Can only be referred by a user");
            //can only set referral on first transaction
            //To set it after first transaction, user must use the standalone setReferrer function
            if (referredToReferrer[msg.sender] == address(0)) {
                referredToReferrer[msg.sender] = referrer;
                referrerToReferred[referrer].push(msg.sender);
            }
        }

        PoolPair memory pooPairInfo = idToPair[pairID];
        StackTooDeepAvoider1 memory container;
        container.preTransferBalance = IERC20(pooPairInfo.thisChainAsset)
            .balanceOf(pooPairInfo.thisChainPool);
        container.pairID = pairID;
        container.chain2Wallet = chain2Wallet;
        container.secondPairID = secondPairID;
        container.tipAmount = tipAmount;
        uint _ICID = pooPairInfo.iCID;
        require(pooPairInfo.thisChainAsset != address(0));

        if (alternateFee) {
            require(
                tipAmount >=
                    (
                        multiChainHop
                            ? (eCIDToTipMultiplier[chain2] *
                                defaultTipAlternate *
                                3) / 2
                            : eCIDToTipMultiplier[idToPair[pairID].otherChain] *
                                defaultTipAlternate
                    ) ||
                    msg.sender == address(this),
                "Declared tip amount must be greater than default tip"
            );
            require(
                _alternateFeeActive,
                "Can only pay with alt fee once activated"
            );

            container.preTipAmount = _morphswapToken.balanceOf(
                iCIDToAltNCPA[_ICID]
            );
            require(
                _morphswapToken.transferFrom(
                    msg.sender,
                    iCIDToAltNCPA[_ICID],
                    tipAmount
                ),
                "Error transferring alternate fee tip"
            );
        } else {
            require(
                tipAmount >=
                    (
                        multiChainHop
                            ? (eCIDToTipMultiplier[chain2] * defaultTip * 3) / 2
                            : eCIDToTipMultiplier[idToPair[pairID].otherChain] *
                                defaultTip
                    ) ||
                    msg.sender == address(this),
                "Declared tip amount must be greater than default tip"
            );

            require(
                msg.value >= tipAmount || msg.sender == address(this),
                "Sent tip must be equal or greater than declared tip unless called self-referentially"
            );
            container.preTipAmount = mCPAArray[_ICID].balance;
            //CHANGED tipAmount to msg.value
            (bool tipResult, ) = mCPAArray[_ICID].call{value: msg.value}("");
            require(tipResult, "Error sending tip");
        }

        require(pooPairInfo.isValid, "Can only buy from a valid pool");

        txNumber++;
        iCIDToLastRTXNumber[_ICID]++;

        require(
            container.preTransferBalance > 0,
            "Can only buy from a pool with a balance greater than 0"
        );
        uint referralBonus = (saleAmount) / 10000;
        uint endSaleAmount = saleAmount -
            (referralBonus * _referralBonusMultiplier);
        if (referredToReferrer[msg.sender] == address(0)) {
            endSaleAmount =
                saleAmount -
                (referralBonus * (_referralBonusMultiplier * 2));
            require(
                IERC20(pooPairInfo.thisChainAsset).transferFrom(
                    msg.sender,
                    _admin,
                    referralBonus * (_referralBonusMultiplier * 2)
                ),
                "Error transferring tokens; make sure contract has allowance"
            );
        } else {
            require(
                IERC20(pooPairInfo.thisChainAsset).transferFrom(
                    msg.sender,
                    referredToReferrer[msg.sender],
                    referralBonus * _referralBonusMultiplier
                ),
                "Error transferring tokens; make sure contract has allowance"
            );
        }
        require(
            IERC20(pooPairInfo.thisChainAsset).transferFrom(
                msg.sender,
                pooPairInfo.thisChainPool,
                endSaleAmount
            ),
            "Error transferring tokens; make sure contract has allowance"
        );
        assert(
            uint64(
                ((endSaleAmount - (referralBonus * _fee)) * oneQuadrillion) /
                    (container.preTransferBalance +
                        (endSaleAmount - (referralBonus * _fee)))
            ) ==
                ((endSaleAmount - (referralBonus * _fee)) * oneQuadrillion) /
                    (container.preTransferBalance +
                        (endSaleAmount - (referralBonus * _fee)))
        );

        uint tipRatioRecord = (tipAmount * oneQuadrillion) /
            (container.preTipAmount + tipAmount);
        uint64 tipRatioSend = uint64(tipRatioRecord);
        assert(tipRatioSend == tipRatioRecord);

        if (multiChainHop) {
            //Multi-chain swaps cannot start on central chain
            require(
                !centralContract,
                "Cannot do multi-chain swap with the central chain as starting point"
            );
            //tip amount is divided by 3 for ease of computation when tx is recieved by mid-point chain
            //CHANGED (7/28) FROM 3rd arg being uint8(_icid) to uint8(eCIDToTipMultiplier[c2])
            gTXNumberToTXObject[txNumber - 1] = TXObject(
                10,
                internalChainID,
                uint8(eCIDToTipMultiplier[chain2]),
                container.pairID,
                container.chain2Wallet,
                secondPairID,
                address(0),
                address(0),
                uint64(
                    ((endSaleAmount - (referralBonus * _fee)) *
                        oneQuadrillion) /
                        (container.preTransferBalance +
                            (endSaleAmount - (referralBonus * _fee)))
                ),
                tipRatioSend / 3,
                iCIDToLastRTXNumber[_ICID] - 1,
                alternateFee
            );

            iCIDToRTXNumberToTXObject[_ICID][
                iCIDToLastRTXNumber[_ICID] - uint128(1)
            ] = gTXNumberToTXObject[txNumber - 1];
            rTXToBlockNumber[iCIDToLastRTXNumber[_ICID] - uint128(1)] = block
                .number;
            //tip amount is divided by 3 for ease of computation when tx is recieved by mid-point chain
            //event MultiChainBuy(uint indexed otherchainidnum, uint indexed pairidnum, uint32 internalfinalchainnum, uint finalpairnum, uint sa
            emit MultiChainBuy({
                otherChainIdNumber: idToPair[container.pairID].otherChain,
                pairIdNum: container.pairID,
                internalFinalChainNumber: uint32(container.secondPairID % 1000),
                finalPairNumber: container.secondPairID,
                saleAmount: endSaleAmount - (referralBonus * _fee),
                poolBalance: container.preTransferBalance,
                chain2Wallet: container.chain2Wallet,
                transactionNumber: txNumber - 1,
                tipAmountQuad: container.tipAmount,
                methodID: MethodIDs.Buy
            });
        } else {
            gTXNumberToTXObject[txNumber - 1] = TXObject(
                1,
                internalChainID,
                uint8(_ICID),
                container.pairID,
                chain2Wallet,
                0,
                address(0),
                address(0),
                uint64(
                    ((endSaleAmount - (referralBonus * _fee)) *
                        oneQuadrillion) /
                        (container.preTransferBalance +
                            (endSaleAmount - (referralBonus * _fee)))
                ),
                tipRatioSend,
                iCIDToLastRTXNumber[_ICID] - 1,
                alternateFee
            );

            uint128 cur_rtxnum = iCIDToLastRTXNumber[_ICID] - uint128(1);
            iCIDToRTXNumberToTXObject[_ICID][cur_rtxnum] = gTXNumberToTXObject[
                txNumber - 1
            ];
            rTXToBlockNumber[cur_rtxnum] = block.number;
            emit Buy({
                otherChainIdNum: idToPair[container.pairID].otherChain,
                pairIdNum: container.pairID,
                saleAmount: endSaleAmount - (referralBonus * _fee),
                poolBalance: container.preTransferBalance,
                chain2Wallet: container.chain2Wallet,
                transactionNumber: txNumber - 1,
                tipAmount: container.tipAmount,
                preTipAmount: container.preTipAmount,
                methodID: MethodIDs.Buy
            });
        }

        if (oldUser[msg.sender] == false) {
            oldUser[msg.sender] = true;
        }

        return true;
    }
}
