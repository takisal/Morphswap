package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"

	overallContractMask "morphswap/node/overallContractMask"
	postInteractionContract "morphswap/node/postInteractionContract"
	preInteraction "morphswap/node/preInteractionContract"
	"morphswap/node/setup"

	"github.com/btcsuite/btcd/btcutil"
	secondCfg "github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/go-sql-driver/mysql"
	bitcoindclient "github.com/joakimofv/go-bitcoindclient/v23"
	"github.com/libsv/go-bk/bec"
	"github.com/libsv/go-bk/bip32"
	"github.com/libsv/go-bk/chaincfg"
	"github.com/tyler-smith/go-bip39"
)

var personalRoot uint
var recommendedAddressAmount uint
var currentNodeIDs []uint8
var currentTXCount uint
var processedBTCOutbound uint
var mainPolygonClient *ethclient.Client
var relativePrivKeyArray []*bec.PrivateKey
var relativePubKeyArray []string
var validNodesIP []string
var ipToCNID map[string]uint8
var receiveWalletPubKey []*bec.PublicKey
var indexConverter map[string]uint8
var localNodeID uint8
var addressConverter map[string][]*bec.PublicKey
var validMultiSigs map[uint]string
var validMultiSigAddresses []btcutil.Address
var bitcoinRPC *rpcclient.Client
var receivedPSBTs map[string][]bitcoindclient.SignRawTransactionWithKeyReqPrevTxs
var defaultPolygonAccount string
var polygonAccountType common.Address
var polygonBTCOverallContract *overallContractMask.OverallContractMask
var preInteractionContract *preInteraction.PreInteraction
var overallContractCentral *overallContractMask.OverallContractMask
var postContract *postInteractionContract.PostInteractionContract
var bitcoinTXFee uint
var earlyReceipt map[string][]bitcoindclient.SignRawTransactionWithKeyReqPrevTxs
var holdingTable map[string]string
var alreadyProcessed map[string]bool
var multiSigDepositAddress string
var multiSigDepositAddressUtil btcutil.Address
var mySQLClient *sql.DB
var recWalletPrivateKey *bec.PrivateKey
var localPortNum uint
var cNIDToPublicKeys map[uint][]string
var listenGL map[string]bool
var openEndpoint bool
var publicKeyConverter map[string]string
var privateKeyConverter map[string]*bec.PrivateKey
var mneumonicString string
var receiveWalletPubKeyString string
var receiveWalletPubKeyStringSlice []string
var receiveWalletPubKeyBECSlice []*bec.PublicKey
var publicKey1 string
var publicKey2 string
var publicKey3 string
var privateKey1 *bec.PrivateKey
var privateKey2 *bec.PrivateKey
var privateKey3 *bec.PrivateKey
var receiveWalletPubKeySingle *bec.PublicKey
var extendedKey *bip32.ExtendedKey
var polygonPrivKey *ecdsa.PrivateKey
var redeemScriptCoverter map[string]string = make(map[string]string)
var bc *bitcoindclient.BitcoindClient
var SignalNumber uint
var performSetup bool
var configValues setup.ConfigValues

// hardened
var path string = "2147483648/0/0"

type rPubKeys []*bec.PublicKey

type PSBTFullStruct struct {
	PSBT        string                                               `json:"PSBT"`
	SPK         uint                                                 `json:"SPK"`
	SVRA        []bitcoindclient.SignRawTransactionWithKeyReqPrevTxs `json:"SVRA"`
	TxVirginHex string                                               `json:"TxVirginHex"`
}
type DataTransmission struct {
	AddressConverter map[string][]*bec.PublicKey `json:"AddressConverter"`
	PublicKeys       []*bec.PublicKey            `json:"PublicKeys"`
}
type ReceiptFullStruct struct {
	PortTEST     uint     `json:"PortTEST"`
	Pubykeyarray []string `json:"Pubykeyarray"`
	Supplemental bool     `json:"Supplemental"`
}

func init() {
	//initialize maps
	cNIDToPublicKeys = make(map[uint][]string)
	ipToCNID = make(map[string]uint8)
	indexConverter = make(map[string]uint8)
	addressConverter = make(map[string][]*bec.PublicKey)
	validMultiSigs = make(map[uint]string)
	receivedPSBTs = make(map[string][]bitcoindclient.SignRawTransactionWithKeyReqPrevTxs)
	earlyReceipt = make(map[string][]bitcoindclient.SignRawTransactionWithKeyReqPrevTxs)
	holdingTable = make(map[string]string)
	alreadyProcessed = make(map[string]bool)
	cNIDToPublicKeys = make(map[uint][]string)
	listenGL = make(map[string]bool)
	publicKeyConverter = make(map[string]string)
	privateKeyConverter = make(map[string]*bec.PrivateKey)
	performSetup = true
	localNodeID = uint8(configValues.NODE_ID)
	localPortNum = configValues.PORT
	privateKey, err := crypto.HexToECDSA(configValues.PRIV_KEY)
	if err != nil {
		log.Fatal(err)
	}
	polygonPrivKey = privateKey
	//generate multiSigReceiver
	personalRoot = configValues.BTC_WALLET_ROOT
	mneumonicString = configValues.BTC_MNEUMONIC
	seed := bip39.NewSeed(mneumonicString, "")
	extendedKey, _ = bip32.NewMaster(seed, &chaincfg.MainNet)
	child, _ := extendedKey.DeriveChildFromPath(path)
	privKey, _ := child.ECPrivKey()
	recWalletPrivateKey = privKey
	pubKey, _ := child.ECPubKey()
	encStr := hex.EncodeToString(pubKey.SerialiseCompressed()[:])
	receiveWalletPubKeyString = encStr
	receiveWalletPubKeySingle = pubKey

	//generate initial 3 keys
	generateInitial()

	openEndpoint = false
}
func main() {

	c := time.Tick(30 * time.Second)
	var tickCounter uint = 0
	go func() {
		for range c {

			tickCounter++
			if tickCounter%2 == 0 {
				monitorIncomingBitcoinTransactions()
				peripheralToCentralSwapmine()
				log.Println("Bitcoin checked")
				log.Println("Swapmine checked")
			} else {
				checkIncomingPolygonTransaction()

				log.Println("Polygon checked")
			}

		}
	}()
	var IAbool bool
	fmt.Print("Sync to Polygon? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	if input == "y" {
		performSetup = false
	}
	fmt.Print("Import addresses? (y/n): ")
	reader1 := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input1, err := reader1.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	input1 = strings.TrimSuffix(input1, "\n")
	if input == "n" {
		IAbool = false
	}
	connectToPolygon()
	connectToBitcoind()
	connectToMySQL()
	log.Println("Current TX Count: ", currentTXCount, "Processed BTC Outbound Transactions: ", processedBTCOutbound)
	//Doesn't import addresses on startup if the user doesnt want them
	time.AfterFunc(6*time.Second, func() {
		onStartup(performSetup, IAbool)
		if performSetup == false {
			subscribeToPolygon()
		}

	})

	uptimeTicker := time.NewTicker(90 * time.Second)
	http.HandleFunc("/psbt", psbtHandlingPOST)
	http.HandleFunc("/recvnodepubkeys", recvnodepubkeysPOST)
	http.HandleFunc("/reqpubkeys", reqpubkeysGET)
	http.HandleFunc("/fullpubkeylist", fullpubkeylistGET)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(localPortNum), nil))
	log.Println("HTTP handlers: online")
	for {
		select {
		case <-uptimeTicker.C:
			if len(validMultiSigs) > 0 {
				monitorIncomingBitcoinTransactions()
			}

		}
	}

}

func generateInitial() {
	//privKeys: privateKey1, privateKey2, privateKey3 (type privKey)
	//pubKeys: publicKey1, publicKey2, publicKey3 (type string)
	log.Println("Generating Initial Keys")
	seed := bip39.NewSeed(mneumonicString, "")
	extendedKey, _ := bip32.NewMaster(seed, &chaincfg.MainNet)
	var adjustedPath string = "2147483648/0/10" + fmt.Sprint(int(personalRoot))
	child, _ := extendedKey.DeriveChildFromPath(adjustedPath)
	privateKey1, _ = child.ECPrivKey()
	pubKey, _ := child.ECPubKey()
	publicKey1 = hex.EncodeToString(pubKey.SerialiseCompressed()[:])
	adjustedPath = "2147483648/0/10" + fmt.Sprint(int(personalRoot)+1)
	child, _ = extendedKey.DeriveChildFromPath(adjustedPath)
	privateKey2, _ = child.ECPrivKey()
	pubKey, _ = child.ECPubKey()
	publicKey2 = hex.EncodeToString(pubKey.SerialiseCompressed()[:])
	adjustedPath = "2147483648/0/10" + fmt.Sprint(int(personalRoot)+2)
	child, _ = extendedKey.DeriveChildFromPath(adjustedPath)
	privateKey3, _ = child.ECPrivKey()
	pubKey, _ = child.ECPubKey()
	publicKey3 = hex.EncodeToString(pubKey.SerialiseCompressed()[:])
}
func onStartup(importAddresses bool, importAddressesScope bool) {
	contract, err := preInteraction.NewPreInteraction(common.HexToAddress(configValues.CONTRACTPRE_ADDR), mainPolygonClient)
	if err != nil {
		log.Println("Failed to create preInteraction Contract")
	}
	big.NewInt(0)
	_inids, _ := contract.Getinids(&bind.CallOpts{Pending: false, From: polygonAccountType, BlockNumber: nil, Context: nil})
	preInteractionContract = contract
	currentNodeIDs = _inids
	multiSigQuery, _ := contract.Getmultisigamount(&bind.CallOpts{Pending: false, From: polygonAccountType, BlockNumber: nil, Context: nil})
	recommendedAddressAmount = uint((*multiSigQuery).Uint64())
	relativePrivKeyArray = []*bec.PrivateKey{privateKey1, privateKey2, privateKey3}
	relativePubKeyArray = []string{publicKey1, publicKey2, publicKey3}
	cNIDToPublicKeys[uint(localNodeID)] = []string{publicKey1, publicKey2, publicKey3}
	seed := bip39.NewSeed(mneumonicString, "")
	extendedKey, _ := bip32.NewMaster(seed, &chaincfg.MainNet)
	for i := 3; i < int(recommendedAddressAmount); i++ {
		//TODO bring back to hardened
		tPath := "0/0/10" + fmt.Sprint(int(personalRoot)+i)
		childDerived, _ := extendedKey.DeriveChildFromPath(tPath)
		tPrivKey, _ := childDerived.ECPrivKey()
		tPubKeyEC, _ := childDerived.ECPubKey()
		tPubKey := hex.EncodeToString(tPubKeyEC.SerialiseCompressed()[:])
		relativePrivKeyArray = append(relativePrivKeyArray, tPrivKey)
		relativePubKeyArray = append(relativePubKeyArray, tPubKey)
		cNIDToPublicKeys[uint(localNodeID)] = append(cNIDToPublicKeys[uint(localNodeID)], tPubKey)
	}
	var counter uint8 = 0
	var pubKeyHolder [][]*bec.PublicKey = make([][]*bec.PublicKey, recommendedAddressAmount+1, recommendedAddressAmount+1)
	time.Sleep(10 * time.Second)
	for j := 0; j < len(currentNodeIDs); j++ {
		returnedIP, err := contract.CnidIp(&bind.CallOpts{Pending: false, From: polygonAccountType, BlockNumber: nil, Context: nil}, currentNodeIDs[j])
		if err != nil {
			log.Println("Failed to fetch IPs from Contract")
		}
		validNodesIP = append(validNodesIP, returnedIP)
		ipToCNID[returnedIP] = currentNodeIDs[j]
		url := "http://" + returnedIP + "/reqpubkeys"
		method := "GET"

		payload := strings.NewReader(``)

		clientR := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("v3", "CG-")

		res, err := clientR.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		var parsedBody []string
		err = json.Unmarshal(body, &parsedBody)
		log.Println(len(parsedBody))
		for k := 0; k < len(parsedBody); k++ {
			if pubKeyHolder[k] == nil {
				pubKeyHolder[k] = []*bec.PublicKey{fromStringToPubKey(parsedBody[k])}
			} else {
				pubKeyHolder[k] = append(pubKeyHolder[k], fromStringToPubKey(parsedBody[k]))
			}
		}
		counter++
		if int(counter) == len(currentNodeIDs) {
			for n := 0; n < len(pubKeyHolder); n++ {
				sort.Sort(rPubKeys(pubKeyHolder[n]))
			}

			for d := 0; d < len(pubKeyHolder[0]); d++ {
				receiveWalletPubKeyBECSlice = append(receiveWalletPubKeyBECSlice, pubKeyHolder[0][d])
				receiveWalletPubKeyStringSlice = append(receiveWalletPubKeyStringSlice, hex.EncodeToString((pubKeyHolder[0][d].SerialiseCompressed()[:])))
			}
			var utilFormP []btcutil.Address
			utilFormP = []btcutil.Address{}
			for yP := 0; yP < len(pubKeyHolder[0]); yP++ {
				paramsP := &secondCfg.MainNetParams
				newAddrWP, _ := btcutil.NewAddressPubKey(pubKeyHolder[0][yP].SerialiseCompressed(), paramsP)
				utilFormP = append(utilFormP, newAddrWP)
			}
			newMultiSigP, _ := bitcoinRPC.CreateMultisig(2, utilFormP)
			addressConverter[newMultiSigP.Address] = pubKeyHolder[0]
			multiSigDepositAddress = newMultiSigP.Address
			redeemScriptCoverter[newMultiSigP.Address] = newMultiSigP.RedeemScript
			multiSigDepositAddressUtil, _ = btcutil.DecodeAddress(multiSigDepositAddress, &secondCfg.MainNetParams)
			log.Println(multiSigDepositAddress)
			for p := 1; p < len(pubKeyHolder); p++ {
				var utilForm []btcutil.Address
				utilForm = []btcutil.Address{}
				for y := 0; y < len(pubKeyHolder[p]); y++ {
					params := &secondCfg.MainNetParams
					newAddrW, _ := btcutil.NewAddressPubKey(pubKeyHolder[p][y].SerialiseCompressed(), params)
					utilForm = append(utilForm, newAddrW)
				}
				newMultiSig, err := bitcoinRPC.CreateMultisig(2, utilForm)
				if err != nil {
					fmt.Println("error:", err)
				}
				addressConverter[newMultiSig.Address] = make([]*bec.PublicKey, 0)
				fmt.Println(newMultiSig.Address, newMultiSig.RedeemScript)
				redeemScriptCoverter[newMultiSig.Address] = newMultiSig.RedeemScript
				addressConverter[newMultiSig.Address] = pubKeyHolder[p]
			}
		}
	}
	var importAndVerify func(indexCount uint8)
	importAndVerify = func(indexCount uint8) {
		p2shAddress, _ := contract.Btcrarray(&bind.CallOpts{Pending: false, From: polygonAccountType, BlockNumber: nil, Context: nil}, big.NewInt(int64(indexCount)))
		if err != nil {
			log.Println("Error retrieving BTC address for index", indexCount)
		}

		indexConverter[p2shAddress] = indexCount + 1
		privateKeyConverter[p2shAddress] = relativePrivKeyArray[indexCount+1]
		publicKeyConverter[p2shAddress] = relativePubKeyArray[indexCount+1]
		if importAddressesScope == true {
			errIA := bitcoinRPC.ImportAddressRescan(p2shAddress, "", false)
			if errIA != nil {
				log.Println("Error importing BTC P2SH address", errIA, p2shAddress)
			} else {
				log.Println("Successfully imported Address: ", p2shAddress)
			}
		}
		validMultiSigs[uint(indexCount)] = p2shAddress

		decodedAddress, _ := btcutil.DecodeAddress(p2shAddress, &secondCfg.MainNetParams)
		validMultiSigAddresses = append(validMultiSigAddresses, decodedAddress)
		if uint(indexCount) < recommendedAddressAmount-2 {
			time.AfterFunc(5*time.Second, func() {
				importAndVerify(indexCount + 1)
			})
		}
	}
	//TODO: change to specific choice input
	if importAddresses == true {
		importAndVerify(0)
	}

}

func reSyncKeys(trustSelf bool) {
	var rskCounter uint = 0
	var holderPubKeysHashed []string
	var holderParsed []DataTransmission
	for i := 0; i < len(currentNodeIDs); i++ {
		if currentNodeIDs[i] != localNodeID {

			url := "http://" + validNodesIP[i] + "/fullpubkeylist"
			method := "GET"

			payload := strings.NewReader(``)

			clientR := &http.Client{}
			req, err := http.NewRequest(method, url, payload)

			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Add("v3", "CG-")

			res, err := clientR.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			var bodyParsed DataTransmission
			err = json.Unmarshal(body, &bodyParsed)
			bodyHash := sha256.Sum256(body)
			holderParsed = append(holderParsed, bodyParsed)
			holderPubKeysHashed = append(holderPubKeysHashed, hex.EncodeToString(bodyHash[:]))
			//Hash body
			rskCounter++
		}
		countTracker := make(map[string]uint)
		var countWinner uint
		var winnerIndex uint
		for i := 0; i < len(holderPubKeysHashed); i++ {
			if countTracker[holderPubKeysHashed[i]] == 0 {
				countTracker[holderPubKeysHashed[i]] = 1
			} else {
				countTracker[holderPubKeysHashed[i]]++
			}
			if countTracker[holderPubKeysHashed[i]] > countWinner {
				countWinner = countTracker[holderPubKeysHashed[i]]
				winnerIndex = uint(i)
			}
		}

		addressConverter = holderParsed[winnerIndex].AddressConverter
		receiveWalletPubKey = holderParsed[winnerIndex].PublicKeys

	}
}

// ==============MONITOR INCOMING TXS====================
func monitorIncomingBitcoinTransactions() {
	if len(validMultiSigAddresses) > 0 {
		uTXOs, err := bitcoinRPC.ListUnspentMinMaxAddresses(1, 999, validMultiSigAddresses)
		if err != nil {
			log.Println("Error listing UTXOs")
		}

		for i := 0; i < len(uTXOs); i++ {

			SatoshiAmountUTXO := uint(uTXOs[i].Amount * 100000000)
			if SatoshiAmountUTXO > 50000 {
				//TODO estimate
				gasFee := 6000
				rScript := redeemScriptCoverter[uTXOs[i].Address]
				log.Println(uTXOs[i].Address, rScript)
				crinp := bitcoindclient.CreateRawTransactionReqInputs{TxID: uTXOs[i].TxID, Vout: float64(uTXOs[i].Vout)}
				outpmap := make(map[string]float64)
				outpmap[multiSigDepositAddress] = uTXOs[i].Amount - float64(float64(gasFee)/100000000)
				crout := bitcoindclient.CreateRawTransactionReqOutputs{A: outpmap}
				thr := bitcoindclient.CreateRawTransactionReq{Inputs: []bitcoindclient.CreateRawTransactionReqInputs{crinp}, Outputs: []bitcoindclient.CreateRawTransactionReqOutputs{crout}, LockTime: float64(0), Replaceable: false}
				transactionResult, erry := bc.CreateRawTransaction(context.Background(), thr)
				if erry != nil {
					// Handle err
					log.Println("creatingrawtx error:", erry)
				} else {
					log.Println("Success creating")

				}

				bPK1 := privateKeyConverter[uTXOs[i].Address]

				serializedPrivateKey := bPK1.Serialise()
				privateKeyFormatted := secp256k1.PrivKeyFromBytes(serializedPrivateKey)
				currentWIF, _ := btcutil.NewWIF(privateKeyFormatted, &secondCfg.MainNetParams, true)
				kRawTransaction := bitcoindclient.SignRawTransactionWithKeyReqPrevTxs{TxID: uTXOs[i].TxID, Vout: float64(uTXOs[i].Vout), ScriptPubkey: uTXOs[i].ScriptPubKey, RedeemScript: rScript}
				var svvr []bitcoindclient.SignRawTransactionWithKeyReqPrevTxs = []bitcoindclient.SignRawTransactionWithKeyReqPrevTxs{kRawTransaction}
				sReq := bitcoindclient.SignRawTransactionWithKeyReq{HexString: transactionResult.Hex, Privkeys: []string{currentWIF.String()}, PrevTxs: svvr}
				signedRawTransaction, err := bc.SignRawTransactionWithKey(context.Background(), sReq)

				if err != nil {
					// Handle err
					log.Println("signingrawtx error:", err)
				} else {
					log.Println("Success signing")
				}

				receivedPSBTs[transactionResult.Hex] = svvr

				nonceCurrent, err := mainPolygonClient.PendingNonceAt(context.Background(), polygonAccountType)
				if err != nil {
					log.Println("Error retrieving nonce")
				} else {

					gasPrice, err := mainPolygonClient.SuggestGasPrice(context.Background())
					if err != nil {
						log.Println("Error retrieving gas price")
					}

					auth, _ := bind.NewKeyedTransactorWithChainID(polygonPrivKey, big.NewInt(137))
					auth.Nonce = big.NewInt(int64(nonceCurrent))
					auth.Value = big.NewInt(0)      // in wei
					auth.GasLimit = uint64(1000000) // in units
					auth.GasPrice = gasPrice
					submittedTX, err := preInteractionContract.SubmitConsensus(auth, uTXOs[0].Address, uTXOs[0].Address, new(big.Int).SetUint64(uint64(SatoshiAmountUTXO-bitcoinTXFee)))
					if err != nil {
						log.Println("Error submitting to Polygon", err)
					} else {
						log.Println("TX hash for polygon: ", submittedTX.Hash().String())
					}
					if reflect.DeepEqual(earlyReceipt[transactionResult.Hex], svvr) {
						psbt := holdingTable[transactionResult.Hex]
						bPK := privateKeyConverter[uTXOs[i].Address]
						serializedPrivateKey2 := bPK.Serialise()
						privateKeyFormatted2 := secp256k1.PrivKeyFromBytes(serializedPrivateKey2)
						currentWIF2, _ := btcutil.NewWIF(privateKeyFormatted2, &secondCfg.MainNetParams, true)
						sReq2 := bitcoindclient.SignRawTransactionWithKeyReq{HexString: psbt, Privkeys: []string{currentWIF2.String()}, PrevTxs: svvr}
						signedRawTransaction2, err := bc.SignRawTransactionWithKey(context.Background(), sReq2)

						if alreadyProcessed[transactionResult.Hex] == false && err == nil {

							alreadyProcessed[transactionResult.Hex] = true

							hrtt := 0.0
							var maxFee *float64 = &hrtt

							txHexToSend := bitcoindclient.SendRawTransactionReq{HexString: signedRawTransaction2.Hex, MaxFeeRate: maxFee}
							rawSent, err := bc.SendRawTransaction(context.Background(), txHexToSend)
							if err != nil {
								log.Println("Error broadcasting raw BTC transaction", err)
							} else {
								log.Println("Successfully broadcast raw BTC transaction: ", rawSent.Hex)
							}
						}
						time.AfterFunc(30*time.Second, func() {
							nonceCurrent, nErr := mainPolygonClient.PendingNonceAt(context.Background(), polygonAccountType)
							if nErr != nil {
								log.Println(nErr)
							} else {

								gasPrice, err := mainPolygonClient.SuggestGasPrice(context.Background())
								if err != nil {
									log.Println("Error retrieving Polygon Block Number")
								} else {
									auth, _ := bind.NewKeyedTransactorWithChainID(polygonPrivKey, big.NewInt(137))
									auth.Nonce = big.NewInt(int64(nonceCurrent))
									auth.Value = big.NewInt(0)      // in wei
									auth.GasLimit = uint64(1000000) // in units
									auth.GasPrice = gasPrice
									polygonBTCOverallContract.OraclePing(auth, uint8(7))

								}
							}

						})

					} else {
						dataToSend := PSBTFullStruct{signedRawTransaction.Hex, uint(indexConverter[uTXOs[0].Address]), svvr, transactionResult.Hex}
						for i := 0; i < len(validNodesIP); i++ {
							if ipToCNID[validNodesIP[i]] != localNodeID {
								jsonBytes, _ := json.Marshal(dataToSend)
								res, err := http.Post("http://"+validNodesIP[i]+"/psbt", "application/json",
									bytes.NewBuffer(jsonBytes))
								if err != nil {
									fmt.Println(err)
									return
								}

								_, err = io.ReadAll(res.Body)
								if err != nil {
									fmt.Println(err)
									return
								}
							}

						}

					}

				}
			}
		}
	}
}

// ======================CHECK INCOMING POLYGON TRANSACTIONS====================
func checkIncomingPolygonTransaction() {
	txStruct, _ := postContract.Gettxs(&bind.CallOpts{Pending: false, From: polygonAccountType, BlockNumber: nil, Context: nil}, new(big.Int).SetUint64(uint64(processedBTCOutbound)))
	if len(txStruct.RecipientBtcAddr) > 0 && txStruct.Procd == false {
		decodedAddress := txStruct.RecipientBtcAddr
		var amountToSend uint

		amountToSend = uint(txStruct.SatsAmount.Uint64())

		var uTXO []bitcoindclient.CreateRawTransactionReqInputs
		var uTXOSatAmount uint = 0
		gasFee := uint(6000) //TODO: ESTIMATE GAS
		uTXOList, err2 := bitcoinRPC.ListUnspentMinMaxAddresses(1, 999, []btcutil.Address{multiSigDepositAddressUtil})
		if err2 != nil {
			log.Println(err2)

			processedBTCOutbound++
			_, qErr := mySQLClient.Query(`UPDATE `+configValues.DB_NAME+` SET cval = ? WHERE name = 'pbtco'`, []uint{processedBTCOutbound})
			if qErr != nil {
				log.Println(qErr)
			} else {
				log.Printf("Updated PBTCO: %d", processedBTCOutbound)
			}
			var svvr []bitcoindclient.SignRawTransactionWithKeyReqPrevTxs = []bitcoindclient.SignRawTransactionWithKeyReqPrevTxs{}

			for i := 0; (i < len(uTXOList)) && (uTXOSatAmount < amountToSend+gasFee); i++ {
				specificUTXO := bitcoindclient.CreateRawTransactionReqInputs{TxID: uTXOList[i].TxID, Vout: float64(uTXOList[i].Vout)}
				uTXO = append(uTXO, specificUTXO)
				uTXOSatAmount += uint(uTXOList[i].Amount * 100000000)

				krt := bitcoindclient.SignRawTransactionWithKeyReqPrevTxs{TxID: uTXOList[i].TxID, Vout: float64(uTXOList[i].Vout), ScriptPubkey: uTXOList[i].ScriptPubKey, RedeemScript: redeemScriptCoverter[multiSigDepositAddress]}
				svvr = append(svvr, krt)
			}

			SatoshiAmountUTXO := txStruct.SatsAmount.Uint64()
			if SatoshiAmountUTXO > 50000 {
				//TODO estimate
				gasFee := 6000
				rScript := redeemScriptCoverter[multiSigDepositAddress]

				log.Println("rscript", rScript)
				outpmap := make(map[string]float64)
				outpmap[decodedAddress] = float64(float64(int(SatoshiAmountUTXO)-gasFee) / 100000000)
				outpmap[multiSigDepositAddress] = float64(float64((uTXOSatAmount-uint(SatoshiAmountUTXO))-uint(gasFee)) / 100000000)
				crout := bitcoindclient.CreateRawTransactionReqOutputs{A: outpmap}
				thr := bitcoindclient.CreateRawTransactionReq{Inputs: uTXO, Outputs: []bitcoindclient.CreateRawTransactionReqOutputs{crout}, LockTime: float64(0), Replaceable: false}
				transactionResult, erry := bc.CreateRawTransaction(context.Background(), thr)
				if erry != nil {
					// Handle err
					log.Println("creatingrawtx error:", erry)
				} else {
					log.Println("Success creating")
					log.Println(transactionResult.Hex)
				}

				serializedPrivateKey := recWalletPrivateKey.Serialise()
				privateKeyFormatted := secp256k1.PrivKeyFromBytes(serializedPrivateKey)
				currentWIF, _ := btcutil.NewWIF(privateKeyFormatted, &secondCfg.MainNetParams, true)
				signRequirement := bitcoindclient.SignRawTransactionWithKeyReq{HexString: transactionResult.Hex, Privkeys: []string{currentWIF.String()}, PrevTxs: svvr}
				signedRawTransaction, _ := bc.SignRawTransactionWithKey(context.Background(), signRequirement)

				receivedPSBTs[transactionResult.Hex] = svvr

				if reflect.DeepEqual(earlyReceipt[transactionResult.Hex], svvr) {
					nFSBT := holdingTable[transactionResult.Hex]

					serializedPrivateKey2 := recWalletPrivateKey.Serialise()
					privateKeyFormatted2 := secp256k1.PrivKeyFromBytes(serializedPrivateKey2)
					currentWIF2, _ := btcutil.NewWIF(privateKeyFormatted2, &secondCfg.MainNetParams, true)
					signRequirement2 := bitcoindclient.SignRawTransactionWithKeyReq{HexString: nFSBT, Privkeys: []string{currentWIF2.String()}, PrevTxs: svvr}
					signedRawTransaction2, err := bc.SignRawTransactionWithKey(context.Background(), signRequirement2)

					if alreadyProcessed[transactionResult.Hex] == false && err == nil {
						alreadyProcessed[transactionResult.Hex] = true
						hrtt := 0.1
						var maxFee *float64 = &hrtt
						txHexToSend := bitcoindclient.SendRawTransactionReq{HexString: signedRawTransaction2.Hex, MaxFeeRate: maxFee}
						rawSent, err := bc.SendRawTransaction(context.Background(), txHexToSend)
						if err != nil {
							log.Println("Error broadcasting raw BTC transaction")
						} else {
							log.Println("Successfully broadcast raw BTC transaction: ", rawSent.Hex)
						}
					}

				} else {
					//if not already recieved, pass it to other nodes
					tempStruct := PSBTFullStruct{signedRawTransaction.Hex, uint(999), svvr, transactionResult.Hex}
					for i := 0; i < len(validNodesIP); i++ {
						if i != int(localNodeID) {
							url := "http://" + validNodesIP[i] + "/psbt"
							method := "POST"

							jsonBytes, _ := json.Marshal(tempStruct)
							payload := bytes.NewReader(jsonBytes)

							clientR := &http.Client{}
							req, err := http.NewRequest(method, url, payload)

							if err != nil {
								fmt.Println(err)
								return
							}
							req.Header.Add("v3", "CG-")

							res, err := clientR.Do(req)
							if err != nil {
								fmt.Println(err)
								return
							}
							defer res.Body.Close()

							_, err = io.ReadAll(res.Body)
							if err != nil {
								fmt.Println(err)
								return
							}
						}

					}
				}
			}

		}

	}
}

func psbtHandlingPOST(rw http.ResponseWriter, req *http.Request) {
	var decodedBody PSBTFullStruct
	body, err := io.ReadAll(req.Body)
	err = json.Unmarshal(body, &decodedBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	var SPK uint = decodedBody.SPK
	psbtHex := decodedBody.PSBT
	txvh := decodedBody.TxVirginHex
	if err != nil {
		panic(err)
	}
	sVRASlice := decodedBody.SVRA
	respBytes, _ := hex.DecodeString("success")
	rw.Write(respBytes)
	log.Println(reflect.DeepEqual(receivedPSBTs[txvh], decodedBody.SVRA))

	log.Println("Virgin Hex: ", txvh, "PSBT Hex: ", psbtHex, receivedPSBTs[txvh], decodedBody.SVRA, SPK)
	if reflect.DeepEqual(receivedPSBTs[txvh], decodedBody.SVRA) {
		//make sure result is good
		lpp := recWalletPrivateKey.Serialise()
		if SPK == 999 {
			lpp = relativePrivKeyArray[SPK].Serialise()
		}
		kj := secp256k1.PrivKeyFromBytes(lpp)
		frrq, _ := btcutil.NewWIF(kj, &secondCfg.MainNetParams, true)
		jreq := bitcoindclient.SignRawTransactionWithKeyReq{HexString: psbtHex, Privkeys: []string{frrq.String()}, PrevTxs: sVRASlice}
		refy, err := bc.SignRawTransactionWithKey(context.Background(), jreq)
		if err != nil {
			log.Println(err)
		}
		if alreadyProcessed[txvh] == false && err == nil {
			log.Println(986)
			alreadyProcessed[txvh] = true

			hrtt := 0.1
			var maxFee *float64 = &hrtt

			txHexToSend := bitcoindclient.SendRawTransactionReq{HexString: refy.Hex, MaxFeeRate: maxFee}
			rawSent, err := bc.SendRawTransaction(context.Background(), txHexToSend)
			if err != nil {
				log.Println("Error broadcasting raw BTC transaction")
			} else {
				log.Println("Successfully broadcast raw BTC transaction: ", rawSent.Hex)
			}
		}
		time.AfterFunc(30*time.Second, func() {
			nonceCurrent, nErr := mainPolygonClient.PendingNonceAt(context.Background(), polygonAccountType)
			if nErr != nil {
				log.Println(nErr)
			} else {

				gasPrice, err := mainPolygonClient.SuggestGasPrice(context.Background())
				if err != nil {
					log.Println("Error retrieving Polygon Block Number")
				} else {
					auth, _ := bind.NewKeyedTransactorWithChainID(polygonPrivKey, big.NewInt(137))
					auth.Nonce = big.NewInt(int64(nonceCurrent))
					auth.Value = big.NewInt(0)      // in wei
					auth.GasLimit = uint64(1000000) // in units
					auth.GasPrice = gasPrice
					submittedTXop, errop := polygonBTCOverallContract.OraclePing(auth, uint8(7))
					if errop != nil {
						log.Println("Error making Oracle Ping")
					} else {
						log.Println("Successfully swapmined: ", submittedTXop.Hash())
					}

				}
			}

		})
	} else {
		earlyReceipt[txvh] = sVRASlice
		holdingTable[txvh] = psbtHex
	}

}

//==================================SHARE PUBLICKEYS AND DEAL WITH MULTISIGS BEING SUBMITTED TO POLYGON==================

//subscribe to Polygon
//make recvnodepubkeys valid endpoint temporarily

func subscribeToPolygon() {
	sentSigBool, err := preInteractionContract.SignalSent(&bind.CallOpts{Pending: false, From: polygonAccountType, BlockNumber: nil, Context: nil}, big.NewInt(int64(SignalNumber)))
	if err != nil {
		log.Println(err)
	}
	var inputsK uint
	inputsK = uint(sentSigBool)
	if inputsK == 2 || inputsK == 3 {
		SignalNumber++
		openEndpoint = true
		time.AfterFunc(300*time.Second, func() {
			openEndpoint = false
		})
		cNIDToPublicKeys = make(map[uint][]string)
		listenGL = make(map[string]bool)
		relativePrivKeyArray = make([]*bec.PrivateKey, 0)
		cNIDToPublicKeys[uint(localNodeID)] = make([]string, 0)

		cNIDsResult, errCNIDs := preInteractionContract.Getinids(&bind.CallOpts{Pending: false, From: polygonAccountType, BlockNumber: nil, Context: nil})
		if errCNIDs != nil {
			log.Println("Error calling getinids() on preInteractionContract")
		}

		currentNodeIDs = cNIDsResult
		//get amount of recaddresses to have
		msqResult, msqErr := preInteractionContract.Getmultisigamount(&bind.CallOpts{Pending: false, From: polygonAccountType, BlockNumber: nil, Context: nil})
		if msqErr != nil {
			log.Println("Error calling getmultisigamount() on preInteractionContract")
		}
		recommendedAddressAmount = uint(msqResult.Uint64())

		relativePrivKeyArray = []*bec.PrivateKey{recWalletPrivateKey, privateKey1, privateKey2, privateKey3}
		//generate the number of pubkeys
		relativePubKeyArray = []string{receiveWalletPubKeyString, publicKey1, publicKey2, publicKey3}
		cNIDToPublicKeys[uint(localNodeID)] = []string{receiveWalletPubKeyString, publicKey1, publicKey2, publicKey3}
		for i := 3; uint(i) < recommendedAddressAmount; i++ {
			var pathY string = "2147483648/0/10" + fmt.Sprint(uint(personalRoot)+uint(i))
			child, _ := extendedKey.DeriveChildFromPath(pathY)
			privKey, _ := child.ECPrivKey()
			pubKey, _ := child.ECPubKey()
			pubKeyStr := hex.EncodeToString((pubKey.SerialiseCompressed()[:]))
			relativePrivKeyArray = append(relativePrivKeyArray, privKey)
			relativePubKeyArray = append(relativePubKeyArray, pubKeyStr)
			cNIDToPublicKeys[uint(localNodeID)] = append(cNIDToPublicKeys[uint(localNodeID)], pubKeyStr)
		}
		validNodesIP = make([]string, 0)

		//query (ip+port*)s from smart contract *if necessary
		for u := 0; u < len(currentNodeIDs); u++ {
			retIP, err := preInteractionContract.CnidIp(&bind.CallOpts{Pending: false, From: polygonAccountType, BlockNumber: nil, Context: nil}, currentNodeIDs[u])
			if err != nil {
				log.Printf("Error retrieving IP for Node %d", currentNodeIDs[u])
			}
			validNodesIP = append(validNodesIP, retIP)
			ipToCNID[retIP] = currentNodeIDs[u]
		}

		method := "POST"
		var payload *bytes.Buffer
		for p := 0; p < len(relativePubKeyArray); p++ {
			log.Println(p, relativePubKeyArray[p])
		}

		clientR := &http.Client{}
		//then send node's pubkeys to all nodes
		time.Sleep(5 * time.Second)
		for j := 0; j < len(validNodesIP); j++ {
			if ipToCNID[validNodesIP[j]] != localNodeID {
				if inputsK == 3 {

					jsonBytes, _ := json.Marshal(ReceiptFullStruct{localPortNum, relativePubKeyArray[(len(relativePubKeyArray) - 4):], true})
					payload = bytes.NewBuffer(jsonBytes)

				} else {
					jsonBytes, _ := json.Marshal(ReceiptFullStruct{localPortNum, relativePubKeyArray, false})
					payload = bytes.NewBuffer(jsonBytes)
				}
				url := "http://" + validNodesIP[j] + "/recvnodepubkeys"
				req, err := http.NewRequest(method, url, payload)
				req.Close = true
				if err != nil {
					fmt.Println(err)
					return
				}
				req.Header.Add("Content-Type", "application/json")

				res, err := clientR.Do(req)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer res.Body.Close()

				time.Sleep(5 * time.Second)
			}
		}
		_, qErr := mySQLClient.Query(`UPDATE `+configValues.DB_NAME+` SET cval = ? WHERE name = 'signum'`, SignalNumber)
		if qErr != nil {
			log.Println(qErr)
		} else {
			log.Printf("Updated signum %d", SignalNumber)
		}
	}
}

func reqpubkeysGET(rw http.ResponseWriter, req *http.Request) {
	//send back array of strings
	var dataSendRPK []string
	if len(relativePubKeyArray) == int(recommendedAddressAmount) {
		dataSendRPK = append(dataSendRPK, receiveWalletPubKeyString)
	}
	for i := 0; i < len(relativePubKeyArray); i++ {
		dataSendRPK = append(dataSendRPK, relativePubKeyArray[i])
	}
	jsonBytes, _ := json.Marshal(dataSendRPK)
	log.Println("dataSPK length: ", len(dataSendRPK))
	rw.Write(jsonBytes)
}
func fullpubkeylistGET(rw http.ResponseWriter, req *http.Request) {
	var dataSendFPK DataTransmission
	dataSendFPK = DataTransmission{addressConverter, receiveWalletPubKey}
	jsonBytes, _ := json.Marshal(dataSendFPK)
	rw.Write(jsonBytes)
}
func recvnodepubkeysPOST(rw http.ResponseWriter, req *http.Request) {
	if openEndpoint == true {
		tbSent, _ := hex.DecodeString("success")
		rw.Write(tbSent)
		var decodedBody ReceiptFullStruct
		body, err := io.ReadAll(req.Body)
		err = json.Unmarshal(body, &decodedBody)
		var portTEST uint = decodedBody.PortTEST
		var pubKeyArray []string = decodedBody.Pubykeyarray
		log.Println(len(pubKeyArray), "length of recieved pubkeyarray")
		if err != nil {
			log.Println(decodedBody)
			panic(err)
		}
		userRemoteIP := req.RemoteAddr
		if userRemoteIP == "" {
			log.Println("Invalid IP")
			return
		}
		r, _ := regexp.Compile(":")
		kIndex := r.FindStringIndex(userRemoteIP)
		if kIndex == nil {
			log.Println("Invalid IP")
			return
		}
		userRemoteIP = userRemoteIP[:kIndex[0]]

		if listenGL[userRemoteIP+":"+fmt.Sprint(portTEST)] == false {
			listenGL[userRemoteIP+":"+fmt.Sprint(portTEST)] = true
			//TODO change this to bool and not placeholder (9)
			if ipToCNID[userRemoteIP+":"+fmt.Sprint(portTEST)] != 9 {
				cNIDToPublicKeys[uint(ipToCNID[userRemoteIP+":"+fmt.Sprint(portTEST)])] = pubKeyArray
			}
			var checkFlag bool = true
			var lowestNum uint = 0
			for i := 0; i < len(currentNodeIDs); i++ {

				if cNIDToPublicKeys[uint(currentNodeIDs[i])] == nil || len(cNIDToPublicKeys[uint(currentNodeIDs[i])]) < 3 {
					checkFlag = false
				} else {
					if lowestNum == 0 {
						lowestNum = uint(len(cNIDToPublicKeys[uint(currentNodeIDs[i])]))
					} else if uint(len(cNIDToPublicKeys[uint(currentNodeIDs[i])])) < lowestNum {
						lowestNum = uint(len(cNIDToPublicKeys[uint(currentNodeIDs[i])]))
					}
				}
			}

			if checkFlag == true && lowestNum != 0 {
				for k := 0; uint(k) < lowestNum; k++ {

					publicKeysNew := make([]*bec.PublicKey, 0)
					for j := 0; j < len(currentNodeIDs); j++ {
						decPubKey := fromStringToPubKey(cNIDToPublicKeys[uint(j)][k])
						publicKeysNew = append(publicKeysNew, decPubKey)
					}
					sort.Sort(rPubKeys(publicKeysNew))
					var utilFormP []btcutil.Address
					utilFormP = []btcutil.Address{}
					for yP := 0; yP < len(publicKeysNew); yP++ {
						paramsP := &secondCfg.MainNetParams
						newAddrWP, _ := btcutil.NewAddressPubKey(publicKeysNew[yP].SerialiseCompressed(), paramsP)
						utilFormP = append(utilFormP, newAddrWP)
					}

					newMultiSigP, _ := bitcoinRPC.CreateMultisig(2, utilFormP)
					multisigaddrRecNew := newMultiSigP.Address
					log.Println(k, multisigaddrRecNew)
					if k == 0 && decodedBody.Supplemental == false {
						//TODO: ONLY DO THIS FOR ADDING ADDITIONAL ADDRESSES
						//TODO insert into a mysql table and populate at msaddepo at every startup with it
						addressConverter[newMultiSigP.Address] = publicKeysNew
						multiSigDepositAddress = multisigaddrRecNew
						multiSigDepositAddressUtil, _ = btcutil.DecodeAddress(multiSigDepositAddress, &secondCfg.MainNetParams)
						redeemScriptCoverter[newMultiSigP.Address] = newMultiSigP.RedeemScript
					} else {
						indexConverter[multisigaddrRecNew] = uint8(k)
						privateKeyConverter[multisigaddrRecNew] = relativePrivKeyArray[k]
						publicKeyConverter[multisigaddrRecNew] = relativePubKeyArray[k]
						addressRecNew, _ := btcutil.DecodeAddress(multisigaddrRecNew, &secondCfg.MainNetParams)
						validMultiSigAddresses = append(validMultiSigAddresses, addressRecNew)
						validMultiSigs[uint(len(validMultiSigAddresses)-1)] = multisigaddrRecNew
						redeemScriptCoverter[newMultiSigP.Address] = newMultiSigP.RedeemScript
					}
					importResult := bitcoinRPC.ImportAddressRescan(multisigaddrRecNew, "", false)
					if importResult == nil {
						//log.Println("Successfully imported new MultiSig P2SH address")
					}
				}
				var populateBTCAddressesOnChain func(addressIndex uint)

				populateBTCAddressesOnChain = func(addressIndex uint) {

					if addressIndex < uint(len(validMultiSigs)) {
						nonce, err := mainPolygonClient.PendingNonceAt(context.Background(), polygonAccountType)
						if err != nil {
							log.Println("Error retrieving nonce")
						} else {

							gasPrice, err := mainPolygonClient.SuggestGasPrice(context.Background())
							if err != nil {
								log.Println("Error retrieving gas price")
							} else {
								blockNumberPolygon, err := mainPolygonClient.BlockNumber(context.Background())
								if err != nil {
									log.Println("Error retrieving Polygon Block Number")
								} else {
									//TODO send Transaction

									auth, _ := bind.NewKeyedTransactorWithChainID(polygonPrivKey, big.NewInt(137))
									auth.Nonce = big.NewInt(int64(nonce))
									auth.Value = big.NewInt(0)      // in wei
									auth.GasLimit = uint64(1000000) // in units
									auth.GasPrice = gasPrice
									createdTX, err := preInteractionContract.PopulateRecBTCaddress(auth, validMultiSigs[addressIndex], new(big.Int).SetUint64(blockNumberPolygon-(200000+(blockNumberPolygon%43200))))
									if err != nil {
										log.Println("Error creating populate transaction: ", err)
									} else {
										log.Println(createdTX.Hash())
									}
								}
							}
						}
						time.AfterFunc(10*time.Second, func() {
							populateBTCAddressesOnChain(addressIndex + 1)
						})
					}

				}
				populateBTCAddressesOnChain(0)
			}
		}

	} else {
		tbSent, _ := hex.DecodeString("Failure")
		rw.Write(tbSent)
	}
}

func peripheralToCentralSwapmine() {
	txStructPrime, _ := overallContractCentral.GetTxByRTxNumber(&bind.CallOpts{Pending: false, From: polygonAccountType, BlockNumber: nil, Context: nil}, uint8(7), big.NewInt(int64(currentTXCount)))
	if txStructPrime.PairId != 0 {
		currentTXCount++
		_, qErr := mySQLClient.Query(`UPDATE `+configValues.DB_NAME+` SET cval = ? WHERE name = 'crtx'`, currentTXCount)
		if qErr != nil {
			log.Println(qErr)
		} else {
			log.Printf("Updated CRTX %d", currentTXCount)
		}

		nonceCurrent, nErr := mainPolygonClient.PendingNonceAt(context.Background(), polygonAccountType)
		if nErr != nil {
			log.Println(nErr)
		} else {
			gasPrice, err := mainPolygonClient.SuggestGasPrice(context.Background())
			if err != nil {
				log.Println("Error retrieving Polygon Block Number")
			} else {
				auth, _ := bind.NewKeyedTransactorWithChainID(polygonPrivKey, big.NewInt(137))
				auth.Nonce = big.NewInt(int64(nonceCurrent))
				auth.Value = big.NewInt(0)      // in wei
				auth.GasLimit = uint64(1000000) // in units
				auth.GasPrice = gasPrice
				polygonBTCOverallContract.OraclePing(auth, uint8(0))

			}
		}

	}
}
