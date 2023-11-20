// SPDX-License-Identifier: MIT
// Morphswap
pragma solidity ^0.8.12;

import "./IERC20.sol";
import "./postInteractionContract.sol";
import "./preInteractionContractStorage.sol";
import "./OverallContract.sol";
import "./AssetPool.sol";

contract PreInteractionContract is PreInteractionContractStorage {
    constructor(
        address _msBTC,
        address _specialAddr,
        address _postInteractionContract,
        address payable _psbtc_oc,
        address _wrappedMSBTC
    ) {
        txid = 0;
        admin = msg.sender;
        msBTC = _msBTC;
        msBTCInterface = IERC20(_msBTC);
        specialAddr = _specialAddr;
        postInteractionContract = PostInteractionContract(
            payable(_postInteractionContract)
        );
        cBTCOverallContract = OverallContract(_psbtc_oc);
        wrappedMSBTC = _wrappedMSBTC;
        wrappedMSBTCInterface = IERC20(_wmsbtc);
        requiredSignatures = 20;
    }

    function sendSignal(uint signalType) public returns (bool) {
        require(
            msg.sender == admin || msg.sender == address(cBTCOverallContract)
        );
        //signalType 2 means fill up multisgaddresslist
        signalTracker++;
        if (signalType == 2) {
            signalSent[signalTracker] = 2;
            emit populateMSL(signalType);
        }
        //signalType 3 means new nodes list
        else if (signalType == 3) {
            signalSent[signalTracker] = 3;
            emit newNodesList(signalType);
        }

        return true;
    }

    function getINIDs() public view returns (uint8[] memory) {
        return iNIDArray;
    }

    function getrequiredSignatures() public view returns (uint) {
        return requiredSignatures;
    }

    function setRequiredSignatures(uint _multiSigAmount) public returns (bool) {
        require(
            msg.sender == address(cBTCOverallContract) || msg.sender == admin
        );
        requiredSignatures = _multiSigAmount;
        return true;
    }

    function addNode(address _vNode) public returns (bool) {
        require(msg.sender == address(cBTCOverallContract));
        vNodes[_vNode] = true;
        cNIDArray.push(_vNode);
        iNIDArray.push(uint8(cNIDArray.length));
        return true;
    }

    function rmNode(address _vNode, uint8 _cNode) public returns (bool) {
        require(msg.sender == address(cBTCOverallContract));
        vNodes[_vNode] = false;
        //uint8 [] memory localAr;
        uint8[] memory localAr = new uint8[](iNIDArray.length - 1);
        bool trackerBool = false;
        for (uint i = 0; i < iNIDArray.length; i++) {
            if (iNIDArray[i] != _cNode) {
                localAr[i - (trackerBool ? 1 : 0)] = (uint8(iNIDArray[i]));
            } else {
                trackerBool = true;
            }
        }
        iNIDArray = localAr;
        return true;
    }

    function setCNIPip(
        uint8 intID,
        string calldata ipString
    ) public returns (bool) {
        require(cNIDArray[intID] == msg.sender);
        cNIDToIP[intID] = ipString;
        return true;
    }

    function populateRecBTCaddresses(
        string[] calldata _btcRArray
    ) public returns (bool) {
        require(vNodes[msg.sender] == true);
        for (uint i = 0; i < _btcRArray.length; i++) {
            btcRArray.push(_btcRArray[i]);
        }
        return true;
    }

    function populateRecBTCaddress(
        string calldata btcR,
        uint vNumber
    ) public returns (bool) {
        require(vnum < block.number);
        require(vNodes[msg.sender] == true);
        require(vNumberA[vnum][btcR][msg.sender] == false);
        require(vNumberP[vnum][btcR] == false);
        vNumberA[vnum][btcR][msg.sender] = true;

        btcRMToCount[vNumber][btcR]++;
        btcria[vNumber].push(btcR);
        if (
            btcRMToCount[vNumber][btcR] ==
            (
                iNIDArray.length % 2 == 0
                    ? iNIDArray.length
                    : iNIDArray.length + 1
            ) /
                2
        ) {
            btcRArray.push(btcR);
            if (!vNumberP[vNumber][btcR]) {
                vNumberP[vNumber][btcR] = true;
            }
        }

        return true;
    }

    function resetBTCfromOC() public returns (bool) {
        require(
            msg.sender == admin || msg.sender == address(cBTCOverallContract)
        );
        delete btcRArray;
    }

    function resetBTCR(uint vNumber) public returns (bool) {
        require(vNumber < block.number);
        require(vNodes[msg.sender] == true);
        require(rNumberA[vNumber][msg.sender] == false);
        require(rNumberP[vNumber] == false);
        rNumberA[vNumber][msg.sender] = true;
        rRemovalCount[vNumber]++;
        if (
            rRemovalCount[vNumber] ==
            (
                iNIDArray.length % 2 == 0
                    ? iNIDArray.length
                    : iNIDArray.length + 1
            ) /
                2
        ) {
            delete btcRArray;
            if (!rNumberP[vNumber]) {
                rNumberP[vNumber] = true;
            }
        }

        return true;
    }

    function addToPreList(
        string calldata txHash,
        uint chain2,
        uint8 methodID,
        uint sentAmount,
        uint8 internalEndChainID,
        bool multiChainHop,
        bool referralBool,
        address referralAddress,
        uint64 pairID,
        address finalChainWallet,
        uint64 secondPairID,
        address firstChainAsset,
        address finalChainAsset,
        uint128 tipAmount,
        uint8 fsignatureID,
        bool alternateFee
    ) public returns (bool) {
        string memory btcR = btcRArray[btcRIndex];
        uint incrementCount = 0;
        btcRIndex++;
        incrementCount++;
        if (btcRIndex >= btcRArray.length) {
            btcRIndex = 0;
        }
        while (
            txlist[btcR].blockcc > block.number - 5000 &&
            incrementCount <= btcRArray.length
        ) {
            if (btcRIndex >= btcRArray.length) {
                btcRIndex = 0;
            }
            btcR = btcRArray[btcRIndex];
            btcRIndex++;
            incrementCount++;
        }
        require(incrementCount <= btcRArray.length);
        txtracker[msg.sender][txHash] = btcR;
        require(tipAmount > 15000);
        hadhToBTCR[btcR] = 0;
        require(hadhToBTCR[btcR] == 0);

        //for SSL, finalchainwallet used as wallet to send liq tokens to if refbool is true (refbool will be for whether they want to manage liq thru their polygon wallet)
        txlist[btcr] = TXOBject(
            methodID,
            internalEndChainID,
            chain2,
            multiChainHop,
            referralBool,
            referralAddress,
            pairID,
            sentAmount,
            finalChainWallet,
            secondPairID,
            firstChainAsset,
            finalChainAsset,
            tipAmount,
            alternateFee,
            fsignatureID,
            0,
            txHash,
            msg.sender,
            block.number
        );
        //require(postInteractionContract.populateShaTable(btcsender) != address(0));
        hadhToBTCR[btcr] = 1;
        if (btcr_satsAm[btcr].length > 0) {
            delete btcr_satsAm[btcr];
        }
        emit PreInteractionNotification(0, 0);
        return true;
    }

    function trackdest(
        string calldata _txHash
    ) public view returns (string memory) {
        return txTracker[msg.sender][_txHash];
    }

    function gettxs(
        string calldata _txAddress
    ) public view returns (TXObject memory) {
        return txList[_txAddress];
    }

    function submitConsensus(
        string calldata _btcAddress,
        string calldata _btcAddressR,
        uint _satsAmount /*V2*/
    ) public returns (bool) {
        //_btcAddress for where to allocate liq to. Should be user's address
        // use hashToBTCR for txid, require(hadhToBTCR[txid] == false); make hadhToBTCR[txid] = true when marked off,
        require(txList[_btcAddressR].blockcc > block.number - 4999);
        require(vNodes[msg.sender] == true);
        require(txList[_btcAddressR].fsignatureID != 0);
        //string memory teststring = _btcAddress[2:7];
        bool alreadySubmitted = false;
        for (uint8 k = 0; k < submittedWorkTracker[_btcAddressR].length; k++) {
            if (submittedWorkTracker[_btcAddressR][k] == msg.sender) {
                alreadySubmitted = true;
            }
        }
        require(alreadySubmitted == false);
        btcr_satsAm[_btcAddressR].push(_satsAmount);
        submittedWorkTracker[_btcAddressR].push(msg.sender);
        txList[_btcAddressR].validatedCount++;
        btcr_satsAmo2[_btcAddressR][_satsAmount]++; // V2
        if (
            (txList[_btcAddressR].validatedCount == 2 ||
                txList[_btcAddressR].validatedCount == 5 ||
                txList[_btcAddressR].validatedCount == 8) &&
            btcr_satsAmo2[_btcAddressR][_satsAmount] >= 2 /*V2*/
        ) {
            TXObject memory holderObj = txList[_btcAddressR];
            //submit to OverallContract

            if (
                txList[_btcAddressR].fsigid == 1 &&
                (_satsAmount - holderObj.tipAm) > 0
            ) {
                //V3 change holderObj.sentAm to _satsAmount
                try
                    cBTCOverallContract.buy{value: 0}(
                        holderObj.pairID,
                        _satsAmount - holderObj.tipAmount,
                        holderObj.finalChainWallet,
                        uint128(holderObj.tipAmount),
                        holderObj.multiChainHop,
                        holderObj.secondPairID,
                        holderObj.referralAddress,
                        holderObj.alternateFee,
                        holderObj.chain2
                    )
                returns (bool res) {
                    // return res;
                } catch Error(string memory reason) {
                    reasonCodeStorage[errorCount] = reason;
                    errorCount++;
                    //  return false;
                } catch Panic(uint ec) {
                    errorCodeStorage[errorCount] = ec;
                    errorCount++;
                    //  return false;
                } catch (bytes memory lowLevelData) {
                    errorCount++;
                    lowLevelDataStorage[errorCount] = lowLevelData;
                    emit Failed(holderObj.pair_id, 1, _btcAddressR);
                    // return false;
                }
            } else if (
                txList[_btcAddressR].fsignatureID == 2 &&
                (_satsAmount - holderObj.tipAmount) > 0
            ) {
                try
                    cBTCOverallContract.deployNewPoolPair{value: 0}(
                        _satsAmount - holderObj.tipAm,
                        msBTC,
                        holderObj.chain2,
                        holderObj.finalChainAsset,
                        holderObj.finalChainWallet,
                        uint128(holderObj.tipAmount)
                    )
                returns (address res, uint pidInteger) {
                    // use refAddr as polygon address to manage liq from
                    IERC20 assetPoolToken = IERC20(res);
                    uint postDNPBalance = assetPoolToken.balanceOf(
                        address(this)
                    );
                    if (holderObj.refbool == true) {
                        assetPoolToken.transfer(
                            holderObj.refAddr,
                            postDNPBalance
                        );
                    } else {
                        btcAddressStringToPIDToLiquidityAmount[_btcAddress][
                            holderObj.pair_id
                        ] = postDNPBalance;
                    }
                } catch Error(string memory reason) {
                    reasonCodeStorage[errorCount] = reason;
                    errorCount++;
                    //  return false;
                } catch Panic(uint ec) {
                    errorCodeStorage[errorCount] = ec;
                    errorCount++;
                    //  return false;
                } catch (bytes memory lowLevelData) {
                    errorCount++;
                    lowLevelDataStorage[errorCount] = lowLevelData;
                    emit Failed(holderObj.pair_id, 2, _btcAddressR);
                    // return false;
                }
                // function singleSidedLiquidity(uint64 pairID, uint c1a_amount, address c1a) public payable returns (bool){
            } else if (txlist[_btcAddressR].fsignatureID == 3) {
                (, address tcp, , , , , ) = cBTCOverallContract.idToPair(
                    holderObj.pairID
                );
                IERC20 aptoken = IERC20(tcp);
                uint preSSLBalance = aptoken.balanceOf(address(this));

                try
                    cBTCOverallContract.singleSidedLiquidity{value: 0}(
                        holderObj.pair_id,
                        _satsAmount,
                        holderObj.firstchain_asset
                    )
                returns (bool res) {
                    uint postSSLBalance = aptoken.balanceOf(address(this));
                    if (holderObj.refbool == true) {
                        //using finalchain wallet as the wallet to receieve the LP tokens, and refbool == true means polygon-wallet managed, refbool == false means btc pubkey managed
                        aptoken.transfer(
                            holderObj.finalchain_wallet,
                            postSSLBalance - preSSLBalance
                        );
                    } else {
                        btcaddrstr_pid_liqamount[_btcAddress][
                            holderObj.pair_id
                        ] = postSSLBalance - preSSLBalance;
                    }
                } catch Error(string memory reason) {
                    reasonCodeStorage[errorCount] = reason;
                    errorCount++;
                    //  return false;
                } catch Panic(uint ec) {
                    errorCodeStorage[errorCount] = ec;
                    errorCount++;
                    //  return false;
                } catch (bytes memory lowLevelData) {
                    errorCount++;
                    lowLevelDataStorage[errorCount] = lowLevelData;
                    emit Failed(holderObj.pair_id, 3, _btcAddressR);
                    // return false;
                }
            } else if (txlist[_btcAddressR].fsigid == 4) {
                wrappedMSBTCInterface.transfer(
                    txlist[_btcAddressR].refAddr,
                    _satsAmount
                );
            }

            delete submittedWorkTracker[_btcAddressR];
            hadhToBTCR[_btcAddressR] = 0;
            for (uint i = 0; i < btcr_satsAm[_btcAddressR].length; i++) {
                delete btcr_satsAmo2[_btcAddressR][
                    btcr_satsAm[_btcAddressR][i]
                ];
            }
            delete btcr_satsAm[_btcAddressR];

            txlist[_btcAddressR].blockcc = block.number - 4000;
        }
        return true;
    }

    function indirectRedeemLiq(
        uint redeemAmount,
        string calldata _btcAddress,
        uint64 _PID,
        string memory uniqueHash
    ) public returns (bool) {
        require(vNodes[msg.sender] = true);
        require(redeemAmount <= btcaddrstr_pid_liqamount[_btcAddress][_PID]);
        rqList[uniqueHash].validatedCount++;
        if (rqList[uniqueHash].validatedCount == 2) {
            (, address tcp, , , , , ) = cBTCOverallContract.idToPair(_PID);
            AssetPool aptoken = AssetPool(payable(tcp));

            require(
                aptoken.removeLiqQueueSolo(
                    address(this),
                    _btcAddress,
                    redeemAmount
                )
            );
        }
        return true;
    }

    fallback() external payable {}

    receive() external payable {}
}
