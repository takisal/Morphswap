// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity ^0.8.12;

import "./IERC20.sol";
import "./AssetPool.sol";
import "./MorphswapStorage.sol";
import "./CentralPingFunctions.sol";
import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract PingContract is
    ChainlinkClient,
    MorphswapStorage,
    CentralPingFunctions
{
    using Chainlink for Chainlink.Request;

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

    function oraclePing(uint8 _ICID) public returns (uint128, bool) {
        require(
            IERC20(_chainlinkAddress).transferFrom(
                msg.sender,
                address(this),
                _swapminingFee * eCIDToTipMultiplier[chainID]
            )
        );
        Chainlink.Request memory request = buildChainlinkRequest(
            iCIDToJID[_ICID],
            address(this),
            this.fulfillPing.selector
        );
        //Job definition parses the hextxdata (the uint for the current number of tx processed) into the proper function call for getTxByrTXNumberber(uint8, uin128)
        request.addUint("hextxdata", iCIDToNumberOfTXsProcessed[_ICID]);
        request.add(
            "swapminer",
            string.concat("0x", bytes20ToString(bytes20(msg.sender)))
        );

        sendOperatorRequest(
            request,
            chainlinkFee * eCIDToTipMultiplier[chainID]
        );
        return (iCIDToNumberOfTXsProcessed[_ICID], true);
    }

    function fulfillPing(StackTooDeepAvoider3 memory container) public {
        require(container.pairID != 0);
        require(
            container.rTXNumber ==
                iCIDToNumberOfTXsProcessed[container.internalStartChainID]
        );
        require(
            txProcessed[container.internalStartChainID][container.rTXNumber] !=
                true
        );
        txProcessed[container.internalStartChainID][container.rTXNumber] = true;

        uint128 numICID = iCIDToNumberOfTXsProcessed[
            container.internalStartChainID
        ];
        iCIDToNumberOfTXsProcessed[container.internalStartChainID]++;

        if (container.paidWithAlt) {
            AssetPool(payable(iCIDToAltNCPA[container.internalStartChainID]))
                .sendTip(container.tipRatio, address(container.swapminer));
        } else {
            AssetPool(payable(mCPAArray[container.internalStartChainID]))
                .sendTip(container.tipRatio, address(container.swapminer));
        }
        if (container.methodID == 1) {
            swapToBeDone[container.internalStartChainID][container.rTXNumber]++;

            try
                this.swapFinish(
                    container.pairID,
                    container.finalChainWallet,
                    container.sentRatio,
                    uint128(numICID)
                )
            returns (bool result) {
                //return res;
            } catch Error(string memory reason) {
                reasonCodeStorage[errorCount] = reason;
                errorCount++;
                emit Failed(
                    container.internalStartChainID,
                    container.rTXNumber
                );
                //return false;
            } catch Panic(uint errorCode) {
                errorCodeStorage[errorCount] = errorCode;
                errorCount++;
                emit Failed(
                    container.internalStartChainID,
                    container.rTXNumber
                );
                //return false;
            } catch (bytes memory lowLevelData) {
                errorCount++;
                lowLevelDataStorage[errorCount] = lowLevelData;
                emit Failed(
                    container.internalStartChainID,
                    container.rTXNumber
                );
                //return false;
            }
        } else if (container.methodID == 2) {
            if (
                tCW_C1A_C2A_TXObject[container.finalChainWallet][
                    container.firstChainAsset
                ][container.finalChainAsset].methodID == 0
            ) {
                tCW_C1A_C2A_TXObject[container.finalChainWallet][
                    container.firstChainAsset
                ][container.finalChainAsset] = TXObject(
                    container.methodID,
                    container.internalStartChainID,
                    container.internalEndChainID,
                    container.pairID,
                    container.finalChainWallet,
                    container.secondPairID,
                    container.firstChainAsset,
                    container.finalChainAsset,
                    container.sentRatio,
                    container.tipRatio,
                    container.rTXNumber,
                    container.paidWithAlt
                );
            }
        } else if (container.methodID == 3) {
            swapToBeDone[container.internalStartChainID][container.rTXNumber]++;
            try
                this.finishAutoTwoSidedLiquidity(
                    container.pairID,
                    container.finalChainWallet,
                    container.sentRatio,
                    numICID
                )
            returns (bool result) {
                //return res;
            } catch Error(string memory reason) {
                reasonCodeStorage[errorCount] = reason;
                errorCount++;
                emit Failed(
                    container.internalStartChainID,
                    container.rTXNumber
                );
                //return false;
            } catch Panic(uint errorCode) {
                errorCodeStorage[errorCount] = errorCode;
                errorCount++;
                emit Failed(
                    container.internalStartChainID,
                    container.rTXNumber
                );
                //return false;
            } catch (bytes memory lowLevelData) {
                errorCount++;
                lowLevelDataStorage[errorCount] = lowLevelData;
                emit Failed(
                    container.internalStartChainID,
                    container.rTXNumber
                );
                //return false;
            }
        } else if (container.methodID == 6) {
            pairIDWaitingForLiqFromTCWallet[container.pairID][
                container.finalChainWallet
            ] = true;
        } else if (container.methodID == 7) {
            try
                this.acknowledgeFinishLiquidity(
                    container.pairID,
                    container.finalChainWallet,
                    container.firstChainAsset
                )
            returns (bool result) {
                //return res;
            } catch Error(string memory reason) {
                reasonCodeStorage[errorCount] = reason;
                errorCount++;
                emit Failed(
                    container.internalStartChainID,
                    container.rTXNumber
                );
                //return false;
            } catch Panic(uint errorCode) {
                errorCodeStorage[errorCount] = errorCode;
                errorCount++;
                emit Failed(
                    container.internalStartChainID,
                    container.rTXNumber
                );
                //return false;
            } catch (bytes memory lowLevelData) {
                errorCount++;
                lowLevelDataStorage[errorCount] = lowLevelData;
                emit Failed(
                    container.internalStartChainID,
                    container.rTXNumber
                );
                //return false;
            }
        } else if (container.methodID == 8) {
            try this.markNewPoolPairComplete(container.pairID) returns (
                bool result
            ) {
                //return res;
            } catch Error(string memory reason) {
                reasonCodeStorage[errorCount] = reason;
                errorCount++;
                emit Failed(
                    container.internalStartChainID,
                    container.rTXNumber
                );
                //return false;
            } catch Panic(uint errorCode) {
                errorCodeStorage[errorCount] = errorCode;
                errorCount++;
                emit Failed(
                    container.internalStartChainID,
                    container.rTXNumber
                );
                //return false;
            } catch (bytes memory lowLevelData) {
                errorCount++;
                lowLevelDataStorage[errorCount] = lowLevelData;
                emit Failed(
                    container.internalStartChainID,
                    container.rTXNumber
                );
                //return false;
            }
        } else if (
            container.methodID == 9 &&
            greenlitICIDToAddressMap[container.internalStartChainID][
                container.firstChainAsset
            ] ==
            container.finalChainWallet
        ) {
            try
                AssetPool(payable(idToPair[container.pairID].thisChainPool))
                    .removeLiqBypassQueue(
                        container.finalChainWallet,
                        container.sentRatio
                    )
            returns (bool result) {
                // return res;
            } catch Error(string memory reason) {
                reasonCodeStorage[errorCount] = reason;
                errorCount++;
                emit Failed(
                    container.internalStartChainID,
                    container.rTXNumber
                );
                //  return false;
            } catch Panic(uint errorCode) {
                errorCodeStorage[errorCount] = errorCode;
                errorCount++;
                emit Failed(
                    container.internalStartChainID,
                    container.rTXNumber
                );
                //  return false;
            } catch (bytes memory lowLevelData) {
                errorCount++;
                lowLevelDataStorage[errorCount] = lowLevelData;
                emit Failed(
                    container.internalStartChainID,
                    container.rTXNumber
                );
                // return false;
            }
        } else if (
            container.methodID == 10 &&
            eCIDToTipMultiplier[idToPair[container.secondPairID].otherChain] ==
            container.internalEndChainID
        ) {
            uint128 defaulttipmult = uint128(
                eCIDToTipMultiplier[idToPair[container.secondPairID].otherChain]
            );
            bool fudge = container.paidWithAlt;
            uint rt = (uint(container.tipRatio) * 2 * uint(oneQuadrillion)) /
                (uint(oneQuadrillion) + (uint(container.tipRatio) * 3));
            //send the bought tca to a buy function call
            swapToBeDone[container.internalStartChainID][container.rTXNumber]++;
            if (idToPair[container.pairID].thisChainAsset == address(0)) {
                uint previousBalance_nc = address(this).balance;

                //TODO: re-entry attacks
                try
                    this.swapFinish(
                        container.pairID,
                        address(this),
                        container.sentRatio,
                        numICID
                    )
                returns (bool result) {
                    //  return res;
                } catch Error(string memory reason) {
                    reasonCodeStorage[errorCount] = reason;
                    errorCount++;
                    emit Failed(
                        container.internalStartChainID,
                        container.rTXNumber
                    );
                    //  return false;
                } catch Panic(uint errorCode) {
                    errorCodeStorage[errorCount] = errorCode;
                    errorCount++;
                    emit Failed(
                        container.internalStartChainID,
                        container.rTXNumber
                    );
                    //   return false;
                } catch (bytes memory lowLevelData) {
                    errorCount++;
                    lowLevelDataStorage[errorCount] = lowLevelData;
                    emit Failed(
                        container.internalStartChainID,
                        container.rTXNumber
                    );
                    //   return false;
                }
                uint preTipBalanceNC = address(this).balance;
                //uint curbal_nc = address(this).balance;
                if (container.paidWithAlt) {
                    try
                        AssetPool(
                            payable(
                                iCIDToAltNCPA[container.internalStartChainID]
                            )
                        ).sendTip(rt, address(this))
                    returns (bool result) {
                        //    return res;
                    } catch Error(string memory reason) {
                        reasonCodeStorage[errorCount] = reason;
                        errorCount++;
                        emit Failed(
                            container.internalStartChainID,
                            container.rTXNumber
                        );
                        //   return false;
                    } catch Panic(uint errorCode) {
                        errorCodeStorage[errorCount] = errorCode;
                        errorCount++;
                        emit Failed(
                            container.internalStartChainID,
                            container.rTXNumber
                        );
                        //   return false;
                    } catch (bytes memory lowLevelData) {
                        errorCount++;
                        lowLevelDataStorage[errorCount] = lowLevelData;
                        emit Failed(
                            container.internalStartChainID,
                            container.rTXNumber
                        );
                        //    return false;
                    }
                } else {
                    try
                        AssetPool(
                            payable(mCPAArray[container.internalStartChainID])
                        ).sendTip(rt, address(this))
                    returns (bool result) {
                        //    return res;
                    } catch Error(string memory reason) {
                        reasonCodeStorage[errorCount] = reason;
                        errorCount++;
                        emit Failed(
                            container.internalStartChainID,
                            container.rTXNumber
                        );
                        //    return false;
                    } catch Panic(uint errorCode) {
                        errorCodeStorage[errorCount] = errorCode;
                        errorCount++;
                        emit Failed(
                            container.internalStartChainID,
                            container.rTXNumber
                        );
                        //     return false;
                    } catch (bytes memory lowLevelData) {
                        errorCount++;
                        lowLevelDataStorage[errorCount] = lowLevelData;
                        emit Failed(
                            container.internalStartChainID,
                            container.rTXNumber
                        );
                        //    return false;
                    }
                }

                uint curbal_nc = address(this).balance;

                //DONE: change defaultTip being passed in here
                try
                    this.buyWithNativeCoin{
                        value: curbal_nc - previousBalance_nc
                    }(
                        container.secondPairID,
                        container.finalChainWallet,
                        (defaultTip * defaulttipmult) >
                            (curbal_nc - preTipBalanceNC)
                            ? uint128(curbal_nc - preTipBalanceNC)
                            : (defaultTip * defaulttipmult),
                        false,
                        0,
                        false,
                        address(0),
                        fudge,
                        0
                    )
                returns (bool result) {} catch (bytes memory lowLevelData) {
                    errorCount++;
                    lowLevelDataStorage[errorCount] = lowLevelData;
                    emit Failed(
                        container.internalStartChainID,
                        container.rTXNumber
                    );
                    // return false;
                }
            } else {
                IERC20 c1aInterface = IERC20(
                    idToPair[container.pairID].thisChainAsset
                );
                uint previousBalance = c1aInterface.balanceOf(address(this));
                c1aInterface.approve(address(this), type(uint256).max);

                try
                    this.swapFinish(
                        container.pairID,
                        address(this),
                        container.sentRatio,
                        numICID
                    )
                returns (bool result) {
                    //   return res;
                } catch Error(string memory reason) {
                    reasonCodeStorage[errorCount] = reason;
                    errorCount++;
                    emit Failed(
                        container.internalStartChainID,
                        container.rTXNumber
                    );
                    //   return false;
                } catch Panic(uint errorCode) {
                    errorCodeStorage[errorCount] = errorCode;
                    errorCount++;
                    emit Failed(
                        container.internalStartChainID,
                        container.rTXNumber
                    );
                    //   return false;
                } catch (bytes memory lowLevelData) {
                    errorCount++;
                    lowLevelDataStorage[errorCount] = lowLevelData;
                    emit Failed(
                        container.internalStartChainID,
                        container.rTXNumber
                    );
                    //    return false;
                }
                uint currentBalance = c1aInterface.balanceOf(address(this));
                uint128 preTipBalance;
                uint128 postTipBalance;
                //uint rt = (tipratio*2)/(one_quadrillion+(tipratio*3));
                if (container.paidWithAlt) {
                    preTipBalance = uint128(
                        _morphswapToken.balanceOf(address(this))
                    );

                    try
                        AssetPool(
                            payable(
                                iCIDToAltNCPA[container.internalStartChainID]
                            )
                        ).sendTip(rt, address(this))
                    returns (bool result) {} catch (bytes memory lowLevelData) {
                        errorCount++;
                        lowLevelDataStorage[errorCount] = lowLevelData;
                        emit Failed(
                            container.internalStartChainID,
                            container.rTXNumber
                        );
                        //   return false;
                    }
                    postTipBalance = uint128(
                        _morphswapToken.balanceOf(address(this))
                    );
                    try
                        this.buy{value: 0}(
                            container.secondPairID,
                            currentBalance - previousBalance,
                            container.finalChainWallet,
                            defaultTip * defaulttipmult >
                                postTipBalance - preTipBalance
                                ? postTipBalance - preTipBalance
                                : defaultTip * defaulttipmult,
                            false,
                            0,
                            address(0),
                            true,
                            0
                        )
                    returns (bool res) {
                        // return res;
                    } catch Error(string memory reason) {
                        reasonCodeStorage[errorCount] = reason;
                        errorCount++;
                        //  return false;
                    } catch Panic(uint errorCode) {
                        errorCodeStorage[errorCount] = errorCode;
                        errorCount++;
                        //  return false;
                    } catch (bytes memory lowLevelData) {
                        errorCount++;
                        lowLevelDataStorage[errorCount] = lowLevelData;
                        emit Failed(
                            container.internalStartChainID,
                            container.rTXNumber
                        );
                        // return false;
                    }
                } else {
                    preTipBalance = uint128(address(this).balance);
                    //CHANGED 7/28 from postTipBalance
                    assert(preTipBalance == address(this).balance);
                    try
                        AssetPool(
                            payable(mCPAArray[container.internalStartChainID])
                        ).sendTip(rt, address(this))
                    returns (bool res) {} catch (bytes memory lowLevelData) {
                        errorCount++;
                        lowLevelDataStorage[errorCount] = lowLevelData;
                        emit Failed(
                            container.internalStartChainID,
                            container.rTXNumber
                        );
                        // return false;
                    }
                    postTipBalance = uint128(address(this).balance);

                    assert(postTipBalance == address(this).balance);
                    try
                        this.buy{value: postTipBalance - preTipBalance}(
                            container.secondPairID,
                            currentBalance - previousBalance,
                            container.finalChainWallet,
                            defaultTip * defaulttipmult >
                                postTipBalance - preTipBalance
                                ? postTipBalance - preTipBalance
                                : defaultTip * defaulttipmult,
                            false,
                            0,
                            address(0),
                            false,
                            0
                        )
                    returns (bool res) {
                        // return res;
                    } catch Error(string memory reason) {
                        reasonCodeStorage[errorCount] = reason;
                        errorCount++;
                        //  return false;
                    } catch Panic(uint errorCode) {
                        errorCodeStorage[errorCount] = errorCode;
                        errorCount++;
                        //  return false;
                    } catch (bytes memory lowLevelData) {
                        errorCount++;
                        lowLevelDataStorage[errorCount] = lowLevelData;
                        emit Failed(
                            container.internalStartChainID,
                            container.rTXNumber
                        );
                        // return false;
                    }
                }
            }
        }
    }
}
