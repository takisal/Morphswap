// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity ^0.8.12;

import "./IERC20.sol";
import "./AssetPool.sol";
import "./MorphswapStorage.sol";
import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract TestingContract is ChainlinkClient, MorphswapStorage {
    using Chainlink for Chainlink.Request;

    function changeContractAddress(
        uint contractCode,
        address contractAddress
    ) public {
        if (contractCode == 1) {
            buyContract = contractAddress;
        } else if (contractCode == 2) {
            buyWithNativeCoinContract = contractAddress;
        } else if (contractCode == 3) {
            deployNewPoolPairContract = contractAddress;
        } else if (contractCode == 4) {
            finishPoolPairContract = contractAddress;
        } else if (contractCode == 5) {
            autoTwoSidedLiquidityContract = contractAddress;
        } else if (contractCode == 6) {
            manualTwoSidedLiquidityContract = contractAddress;
        } else if (contractCode == 7) {
            finishLiquidityContract = contractAddress;
        } else if (contractCode == 8) {
            confirmRemoveBothSidesLiqContract = contractAddress;
        } else if (contractCode == 9) {
            addSupportedChainsContract = contractAddress;
        } else if (contractCode == 10) {
            acknowledgeFinishLiquidityContract = contractAddress;
        } else if (contractCode == 11) {
            governanceContract = contractAddress;
        } else if (contractCode == 12) {
            singleSidedLiquidityContract = contractAddress;
        } else if (contractCode == 13) {
            contractAddressncelManualEscrowContract = contractAddress;
        } else if (contractCode == 14) {
            pingContract = contractAddress;
        }
    }
}
