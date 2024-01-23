// SPDX-License-Identifier: MIT

// Morphswap

pragma solidity ^0.8.12;

import "./IERC20.sol";
import "./AssetPool.sol";
import "./MorphswapStorage.sol";
import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract BuyWithNativeCoinContract is ChainlinkClient, MorphswapStorage {
    using Chainlink for Chainlink.Request;

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
        if (referrerBool) {
            require(oldUser[msg.sender] == false);
            require(oldUser[referrer]);
            //can only set referral on first transaction
            //To set it after first transaction, user must use the standalone setReferrer function
            if (referredToReferrer[msg.sender] == address(0)) {
                referredToReferrer[msg.sender] = referrer;
                referrerToReferred[referrer].push(msg.sender);
            }
        }
        require(idToPair[pairID].isValid);
        require(idToPair[pairID].thisChainAsset == address(0));
        uint _ICID = idToPair[pairID].iCID;
        uint postTipValue;
        uint preTipAmount;
        bool tipResult;
        if (alternateFee) {
            postTipValue = msg.value;
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
            require(_alternateFeeActive);
            preTipAmount = _morphswapToken.balanceOf(iCIDToAltNCPA[_ICID]);
            tipResult = _morphswapToken.transferFrom(
                msg.sender,
                iCIDToAltNCPA[_ICID],
                tipAmount
            );
        } else {
            postTipValue = msg.value - tipAmount;
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

            require(msg.value >= tipAmount || msg.sender == address(this));

            preTipAmount = mCPAArray[_ICID].balance;
            (tipResult, ) = mCPAArray[_ICID].call{value: tipAmount}("");
        }
        StackTooDeepAvoider1 memory container;
        txNumber++;
        iCIDToLastRTXNumber[_ICID]++;
        container.pairID = pairID;
        container.tipAmount = tipAmount;
        container.chain2Wallet = chain2Wallet;
        container.secondPairID = secondPairID;
        container.multiChainHop = multiChainHop;
        container.alternateFee = alternateFee;
        container.preTipAmount = preTipAmount;
        container._ICID = _ICID;

        require(tipResult);
        uint preTransferBalance = idToPair[container.pairID]
            .thisChainPool
            .balance;
        require(preTransferBalance > 0);
        uint oneTenThousandth = (postTipValue) / 10000;
        uint endSaleAmount;
        if (referredToReferrer[msg.sender] == address(0)) {
            endSaleAmount =
                postTipValue -
                (oneTenThousandth * (_referralBonusMultiplier * 2));
            (bool refbonusresult, ) = _admin.call{
                value: oneTenThousandth * (_referralBonusMultiplier * 2)
            }("");
            require(refbonusresult);
        } else {
            endSaleAmount =
                postTipValue -
                (oneTenThousandth * _referralBonusMultiplier);
            (bool refbonusresult, ) = referredToReferrer[msg.sender].call{
                value: oneTenThousandth * _referralBonusMultiplier
            }("");
            require(refbonusresult);
        }

        (bool swapResult, ) = idToPair[container.pairID].thisChainPool.call{
            value: endSaleAmount
        }("");

        require(swapResult);

        uint ratioRecord = ((endSaleAmount - (oneTenThousandth * _fee)) *
            oneQuadrillion) /
            (preTransferBalance + (endSaleAmount - (oneTenThousandth * _fee)));
        uint64 ratioSend = uint64(ratioRecord);
        assert(ratioSend == ratioRecord);
        uint64 tipRatioSend = uint64(
            (container.tipAmount * oneQuadrillion) /
                (preTipAmount + container.tipAmount)
        );
        require(
            tipRatioSend ==
                (container.tipAmount * oneQuadrillion) /
                    (preTipAmount + container.tipAmount)
        );

        if (container.multiChainHop) {
            require(!centralContract);
            //tip amount is divided by 3 for ease of computation when tx is recieved by mid-point chain
            //CHANGED (7/28) FROM 3rd arg being uint8(_ICID) to uint8(eCIDToTipMultiplier[c2])
            gTXNumberToTXObject[txNumber - 1] = TXObject(
                10,
                internalChainID,
                uint8(eCIDToTipMultiplier[chain2]),
                container.pairID,
                container.chain2Wallet,
                container.secondPairID,
                address(0),
                address(0),
                ratioSend,
                tipRatioSend / 3,
                iCIDToLastRTXNumber[container._ICID] - 1,
                container.alternateFee
            );

            iCIDToRTXNumberToTXObject[_ICID][
                iCIDToLastRTXNumber[container._ICID] - uint128(1)
            ] = gTXNumberToTXObject[txNumber - 1];
            rTXToBlockNumber[iCIDToLastRTXNumber[_ICID] - uint128(1)] = block
                .number;
            //tip amount is divided by 3 for ease of computation when tx is recieved by mid-point chain
            //event MultiChainBuy(uint indexed otherchainidnum, uint indexed pairidnum, uint32 internalfinalchainnum, uint finalpairnum, uint saleAmount, uint poolBalance, address c2w, uint indexed txnumber, uint tipAmount, uint pretipAmount, uint methodid);
            emit MultiChainBuy(
                idToPair[container.pairID].otherChain,
                container.pairID,
                container.secondPairID,
                container.secondPairID % 1000,
                endSaleAmount - (oneTenThousandth * _fee),
                preTransferBalance,
                container.chain2Wallet,
                txNumber - 1,
                tipRatioSend / 3,
                1
            );
        } else {
            gTXNumberToTXObject[txNumber - 1] = TXObject(
                1,
                internalChainID,
                uint8(_ICID),
                container.pairID,
                container.chain2Wallet,
                0,
                address(0),
                address(0),
                ratioSend,
                tipRatioSend,
                iCIDToLastRTXNumber[container._ICID] - 1,
                container.alternateFee
            );

            iCIDToRTXNumberToTXObject[container._ICID][
                iCIDToLastRTXNumber[container._ICID] - uint128(1)
            ] = gTXNumberToTXObject[txNumber - 1];
            rTXToBlockNumber[iCIDToLastRTXNumber[_ICID] - uint128(1)] = block
                .number;
            emit Buy(
                idToPair[container.pairID].otherChain,
                container.pairID,
                endSaleAmount - (oneTenThousandth * _fee),
                preTransferBalance + (oneTenThousandth * _fee),
                container.chain2Wallet,
                txNumber - 1,
                container.tipAmount,
                container.preTipAmount,
                1
            );
        }

        if (oldUser[msg.sender] == false) {
            oldUser[msg.sender] = true;
        }

        return true;
    }
}
