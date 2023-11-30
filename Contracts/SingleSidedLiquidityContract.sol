// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity ^0.8.12;

import "./IERC20.sol";
import "./AssetPool.sol";
import "./MorphswapStorage.sol";
import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract SingleSidedLiquidityContract is ChainlinkClient, MorphswapStorage {
    using Chainlink for Chainlink.Request;

    function singleSidedLiquidity(
        uint64 pairID,
        uint chain1AssetAmount,
        address chain1Asset
    ) public payable returns (bool) {
        //something is wrong with liquidity providing maybe? idk
        require(idToPair[pairID].isValid);
        require(idToPair[pairID].thisChainAsset == chain1Asset);
        if (chain1Asset == address(0)) {
            chain1AssetAmount = msg.value;
            address payable tempad = payable(idToPair[pairID].thisChainPool);
            (bool sentresult, ) = tempad.call{value: msg.value}("");
            require(sentresult);
        } else {
            require(
                IERC20(chain1Asset).transferFrom(
                    msg.sender,
                    idToPair[pairID].thisChainPool,
                    chain1AssetAmount
                )
            );
        }

        (bool sent, uint addedLP, uint oldLPTs) = AssetPool(
            payable(idToPair[pairID].thisChainPool)
        ).addLiquidity(msg.sender, chain1AssetAmount);
        require(sent);
        if (oldUser[msg.sender] == false) {
            oldUser[msg.sender] = true;
        }

        emit SingleLiq(
            pairID,
            chain1Asset,
            msg.sender,
            chain1AssetAmount,
            addedLP,
            oldLPTs,
            block.number,
            uint256(MethodIDs.SingleSidedLiquidity)
        );
        return true;
    }
}
