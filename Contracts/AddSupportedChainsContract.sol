// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity ^0.8.12;

import "./IERC20.sol";
import "./AssetPool.sol";
import "./MorphswapStorage.sol";
import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract AddSupportedChainsContract is ChainlinkClient, MorphswapStorage {
    using Chainlink for Chainlink.Request;

    function addSupportedChains(
        uint _schain,
        string memory jobid,
        address otherchainmorphswap
    ) public returns (bool) {
        require(msg.sender == _admin);
        chainIdToInternalchainid[_schain] = uint8(supportedChainsList.length);
        internalChainIdToChainId[supportedChainsList.length] = _schain;
        require(supportedChainsList.length < 255);
        iCIDToJID[supportedChainsList.length] = bytes32(bytes(jobid));

        supportedChains[_schain] = true;
        if (centralContract != true) {
            pairTracker += 2;
        }

        uint noncentralcid = centralContract
            ? chainIdToInternalchainid[_schain]
            : chainIdToInternalchainid[chainID];
        uint32 nativecoinpoolid = uint32(1000 + noncentralcid);
        mCPAArray.push(
            address(new AssetPool(address(0), nativecoinpoolid, true))
        );
        idToPair[nativecoinpoolid] = PoolPair({
            thisChainAsset: address(0),
            thisChainPool: mCPAArray[supportedChainsList.length],
            otherChain: _schain,
            iCID: uint8(supportedChainsList.length),
            otherChainAsset: address(0),
            pairID: nativecoinpoolid,
            isValid: true
        });
        iCIDTomCPAArray[supportedChainsList.length] = mCPAArray[
            supportedChainsList.length
        ];
        cID_c1A_c2A[_schain][address(0)][address(0)] = idToPair[
            nativecoinpoolid
        ];

        uint32 alttippoolid = uint32(2000 + noncentralcid);
        alternateTipArray.push(
            address(new AssetPool(_morphswapTokenAddress, alttippoolid, true))
        );
        idToPair[alttippoolid] = PoolPair({
            thisChainAsset: _morphswapTokenAddress,
            thisChainPool: alternateTipArray[supportedChainsList.length],
            otherChain: _schain,
            iCID: uint8(supportedChainsList.length),
            otherChainAsset: otherchainmorphswap,
            pairID: alttippoolid,
            isValid: true
        });
        iCIDToAltNcpa[supportedChainsList.length] = alternateTipArray[
            supportedChainsList.length
        ];
        cID_c1A_c2A[_schain][_morphswapTokenAddress][
            otherchainmorphswap
        ] = idToPair[alttippoolid];
        supportedChainsList.push(_schain);
        return true;
    }
}
