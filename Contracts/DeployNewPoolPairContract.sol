// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity ^0.8.12;

import "./IERC20.sol";
import "./AssetPool.sol";
import "./MorphswapStorage.sol";
import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract DeployNewPoolPairContract is ChainlinkClient, MorphswapStorage {
    using Chainlink for Chainlink.Request;

    function deployNewPoolPair(
        uint chain1AssetAmount,
        address chain1Asset,
        uint chain2,
        address chain2Asset,
        address chain2Wallet,
        uint128 tipAmount
    ) public payable returns (address, uint) {
        require(tipAmount >= defaultTip * eCIDToTipMultiplier[chain2]);
        require(msg.value >= tipAmount);
        require(supportedChains[chain2]);
        require(cID_c1A_c2A[chain2][chain1Asset][chain2Asset].isValid != true);

        StackTooDeepAvoider1 memory container;
        container.chain2Wallet = chain2Wallet;
        container.chain1Asset = chain1Asset;
        container.chain2Asset = chain2Asset;
        container.chain2 = chain2;
        container.chain1AssetAmount = chain1AssetAmount;
        require(
            tCW_C1A_C2A_TXObject[chain2Wallet][chain1Asset][chain2Asset]
                .alternateFee == false
        );
        tCW_C1A_C2A_TXObject[chain2Wallet][chain1Asset][chain2Asset]
            .alternateFee = true;
        uint64 _ICID = chainIDToInternalChainID[chain2];

        txNumber++;
        iCIDToLastRTXNumber[_ICID]++;
        container._ICID = _ICID;
        container.rTXNumber = iCIDToLastRTXNumber[_ICID];
        uint postFundAmount = msg.value - tipAmount;
        uint preTipAmount = mCPAArray[_ICID].balance;
        (bool tipResult, ) = mCPAArray[_ICID].call{value: tipAmount}("");
        require(tipResult);
        pairTracker++;
        container.convertedPairID =
            (pairTracker * uint64(1000)) +
            internalChainID;
        AssetPool _chain1PAInterface = new AssetPool(
            chain1Asset,
            container.convertedPairID,
            false
        );
        address chain1PoolAddress = address(_chain1PAInterface);

        require(chain1PoolAddress != address(0));
        IERC20 _c1aInterface = IERC20(chain1Asset);

        uint tipRatioRecord = (tipAmount * oneQuadrillion) /
            (preTipAmount + tipAmount);
        uint64 tipRatioSend = uint64(tipRatioRecord);
        require(tipRatioSend == tipRatioRecord);

        if (container.chain1Asset != address(0)) {
            require(container.chain1AssetAmount > 0);
            require(
                _c1aInterface.transferFrom(
                    msg.sender,
                    chain1PoolAddress,
                    container.chain1AssetAmount
                )
            );
            (bool sent, , ) = _chain1PAInterface.addLiquidity(
                msg.sender,
                container.chain1AssetAmount
            );
            require(sent);
            gTXNumberToTXObject[txNumber - 1] = TXObject(
                2,
                internalChainID,
                uint8(_ICID),
                container.convertedPairID,
                container.chain2Wallet,
                0,
                container.chain1Asset,
                container.chain2Asset,
                0,
                tipRatioSend,
                iCIDToLastRTXNumber[_ICID] - 1,
                false
            );
        } else {
            require(postFundAmount > 1000000000000000);
            (bool fundresult, ) = chain1PoolAddress.call{value: postFundAmount}(
                ""
            );
            require(fundresult);
            (bool sent, , ) = _chain1PAInterface.addLiquidity(
                msg.sender,
                postFundAmount
            );
            require(sent);
            gTXNumberToTXObject[txNumber - 1] = TXObject(
                2,
                internalChainID,
                uint8(_ICID),
                container.convertedPairID,
                container.chain2Wallet,
                0,
                container.chain1Asset,
                container.chain2Asset,
                0,
                tipRatioSend,
                container.rTXNumber - 1,
                false
            );
        }

        uint128 currentRTXNumber = iCIDToLastRTXNumber[_ICID] - uint128(1);
        iCIDToRTXNumberToTXObject[_ICID][
            currentRTXNumber
        ] = gTXNumberToTXObject[txNumber - 1];
        rTXToBlockNumber[currentRTXNumber] = block.number;
        //Marks off as false until completed on second chain
        idToPair[container.convertedPairID] = PoolPair(
            container.chain1Asset,
            chain1PoolAddress,
            container.chain2,
            chainIDToInternalChainID[container.chain2],
            container.chain2Asset,
            container.convertedPairID,
            false
        );

        if (oldUser[msg.sender] == false) {
            oldUser[msg.sender] = true;
        }

        emit NewPair(
            container.chain2,
            container.convertedPairID,
            container.chain2Asset,
            container.chain2Wallet,
            txNumber - 1,
            pairTracker,
            uint256(MethodIDs.NewPair)
        );

        return (chain1PoolAddress, container.convertedPairID);
    }
}
