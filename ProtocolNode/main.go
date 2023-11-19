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
	"github.com/btcsuite/btcd/txscript"
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
	"github.com/spf13/viper"
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
var ip_cNID map[string]uint8
var receiveWalletPubKey []*bec.PublicKey
var indexConverter map[string]uint8
var localNodeID uint8
var addressConverter map[string][]*bec.PublicKey
var validMultiSigs map[uint]string
var validMultiSigAddrs []btcutil.Address
var bitcoinRPC *rpcclient.Client
var recPSBTs map[string][]bitcoindclient.SignRawTransactionWithKeyReqPrevTxs
var defaultPolygonAccount string
var polygonAccountType common.Address
var polygonBTCOverallContract *overallContractMask.OverallContractMask
var preInteractionContract *preInteraction.PreInteraction
var overallContractCentral *overallContractMask.OverallContractMask
var postContract *postInteractionContract.PostInteractionContract
var bitcoinTxFee uint
var earlyRec map[string][]bitcoindclient.SignRawTransactionWithKeyReqPrevTxs
var holdingTable map[string]string
var alreadyProcessed map[string]bool
var multiSigDepositAddress string
var multiSigDepositAddressUtil btcutil.Address
var mySQLClient *sql.DB
var recWalletPrivateKey *bec.PrivateKey
var localPortNum uint
var cNID_pubKeys map[uint][]string
var listenGL map[string]bool
var openEndpoint bool
var pubKeyConverter map[string]string
var privKeyConverter map[string]*bec.PrivateKey
var mneumonicStr string
var receiveWalletPubKeyString string
var allRecWalletPubKeyStr []string
var allRecWalletPubKeyBEC []*bec.PublicKey
var pubKey1 string
var pubKey2 string
var pubKey3 string
var privateKey1 *bec.PrivateKey
var privateKey2 *bec.PrivateKey
var privateKey3 *bec.PrivateKey
var receiveWalletPubKeySingle *bec.PublicKey
var extKey *bip32.ExtendedKey
var polygonPrivKey *ecdsa.PrivateKey
var redeemScriptCoverter map[string]string = make(map[string]string)
var bc *bitcoindclient.BitcoindClient
var SigNumber uint
var SUBool bool
var configValues setup.ConfigValues

// hardened
var path string = "2147483648/0/0"

type rPubKeys []*bec.PublicKey

type psbtFullStruct struct {
	Psbt        string                                               `json:"Psbt"`
	Spk         uint                                                 `json:"Spk"`
	Svvra       []bitcoindclient.SignRawTransactionWithKeyReqPrevTxs `json:"Svvra"`
	TxVirginHex string                                               `json:"TxVirginHex"`
}
type dataTransmiion struct {
	Addrconv   map[string][]*bec.PublicKey `json:"Addrconv"`
	PublicKeys []*bec.PublicKey            `json:"PublicKeys"`
}
type recvFullStruct struct {
	PortTEST     uint     `json:"PortTEST"`
	Pubykeyarray []string `json:"Pubykeyarray"`
	Supplemental bool     `json:"Supplemental"`
}

type filettaRow struct {
	col_name string
	col_val  int
}

func (s rPubKeys) Len() int {
	return len(s)
}
func (s rPubKeys) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s rPubKeys) Less(i, j int) bool {
	//return strings.ToUpper(hex.EncodeToString((s[i].SerialiseCompressed()[:]))) < strings.ToUpper(hex.EncodeToString((s[j].SerialiseCompressed()[:])))
	v := new(big.Int).SetBytes(s[i].SerialiseCompressed()[:])
	d := new(big.Int).SetBytes(s[j].SerialiseCompressed()[:])
	return (*v).Cmp(d) < 0
}
func getRedeemHashFromKeyArray(pkArr []*bec.PublicKey, nSigs uint, mSigs uint) []byte {

	builder := txscript.NewScriptBuilder()
	switch nSigs {
	case 2:
		builder.AddOp(txscript.OP_2)
	case 3:
		builder.AddOp(txscript.OP_3)
	case 4:
		builder.AddOp(txscript.OP_4)
	case 5:
		builder.AddOp(txscript.OP_5)
	case 6:
		builder.AddOp(txscript.OP_6)
	case 7:
		builder.AddOp(txscript.OP_7)
	}

	for k := 0; k < int(mSigs); k++ {
		builder.AddData(pkArr[k].SerialiseCompressed())
	}

	switch mSigs {
	case 2:
		builder.AddOp(txscript.OP_2)
	case 3:
		builder.AddOp(txscript.OP_3)
	case 4:
		builder.AddOp(txscript.OP_4)
	case 5:
		builder.AddOp(txscript.OP_5)
	case 6:
		builder.AddOp(txscript.OP_6)
	case 7:
		builder.AddOp(txscript.OP_7)
	}

	builder.AddOp(txscript.OP_CHECKMULTISIG)

	redeemScript, err := builder.Script()
	if err != nil {
		rVal, _ := hex.DecodeString("Error")
		return rVal
	}

	redeemHash := btcutil.Hash160(redeemScript)
	return redeemHash
}

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	localPortNum = viper.GetUint("EXTERNAL_PORT")
	if localPortNum == 0 {
		fmt.Print("It looks like there is no config file. Would you like to set up the config file now (using the CLI)? (Y/N):")
		reader := bufio.NewReader(os.Stdin)
		// ReadString will block until the delimiter is entered
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			panic("Invalid Entry")
		}
		// remove the delimeter from the string
		input = strings.TrimSuffix(input, "\n")
		if input != "y" && input != "Y" {
			os.Exit(3)
		}
		//begin "set up wizard"
		configValues = setup.SetupWalkthrough()
		log.Println("Successfully wrote and saved new config file")
	} else {
		setup.ReadConfig()
	}
	path += fmt.Sprint(configValues.BTC_WALLET_ROOT)
}
func init() {

	//initialize maps
	cNID_pubKeys = make(map[uint][]string)
	ip_cNID = make(map[string]uint8)
	indexConverter = make(map[string]uint8)
	addressConverter = make(map[string][]*bec.PublicKey)
	validMultiSigs = make(map[uint]string)
	recPSBTs = make(map[string][]bitcoindclient.SignRawTransactionWithKeyReqPrevTxs)
	earlyRec = make(map[string][]bitcoindclient.SignRawTransactionWithKeyReqPrevTxs)
	holdingTable = make(map[string]string)
	alreadyProcessed = make(map[string]bool)
	cNID_pubKeys = make(map[uint][]string)
	listenGL = make(map[string]bool)
	pubKeyConverter = make(map[string]string)
	privKeyConverter = make(map[string]*bec.PrivateKey)
	SUBool = true
	localNodeID = uint8(configValues.NODE_ID)
	localPortNum = configValues.PORT
	log.Println(localNodeID)
	privateKey, err := crypto.HexToECDSA(configValues.PRIV_KEY)
	if err != nil {
		log.Fatal(err)
	}
	polygonPrivKey = privateKey
	//generate multiSigReceiver
	personalRoot = configValues.BTC_WALLET_ROOT
	mneumonicStr = configValues.BTC_MNEUMONIC
	seed := bip39.NewSeed(mneumonicStr, "")
	extKey, _ = bip32.NewMaster(seed, &chaincfg.MainNet)
	child, _ := extKey.DeriveChildFromPath(path)
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
		SUBool = false
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
		onStartup(SUBool, IAbool)
		if SUBool == false {
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

func connectToBitcoind() {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         configValues.BTCNODE_HOST,
		User:         configValues.BTCNODE_USER,
		Pass:         configValues.BTCNODE_PASS,
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	bitcoinRPC = client

	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bitcoin Block count: %d", blockCount)
	bcScope, err := bitcoindclient.New(bitcoindclient.Config{
		RpcAddress:  configValues.BTCNODE_HOST,
		RpcUser:     configValues.BTCNODE_USER,
		RpcPassword: configValues.BTCNODE_PASS,
	})
	bc = bcScope
}

func connectToPolygon() {
	client, err := ethclient.Dial(configValues.HTTP_URL)
	mainPolygonClient = client
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, errGas := client.SuggestGasPrice(context.Background())
	if errGas != nil {
		log.Fatal(errGas)
	}
	log.Println("Succesful connection. Gas price: ", gasPrice)
	oBTCContract, err := overallContractMask.NewOverallContractMask(common.HexToAddress(configValues.CONTRACT_ADDR), mainPolygonClient)
	if err != nil {
		log.Fatal(err)
	}
	polygonBTCOverallContract = oBTCContract
	oContract, err := overallContractMask.NewOverallContractMask(common.HexToAddress(configValues.CONTRACTBTC_ADDR), mainPolygonClient)
	if err != nil {
		log.Fatal(err)
	}
	overallContractCentral = oContract
	postIntContract, err := postInteractionContract.NewPostInteractionContract(common.HexToAddress(configValues.CONTRACTPOST_ADDR), mainPolygonClient)
	if err != nil {
		log.Fatal(err)
	}
	postContract = postIntContract
	//TODO have user enter this
	polyAcctStr := configValues.ADDR
	polygonAccountType = common.HexToAddress(polyAcctStr)

	defaultPolygonAccount = polyAcctStr
}

func connectToMySQL() {
	con, errC := sql.Open("mysql", configValues.DB_URI)
	if errC != nil {
		log.Fatal(errC)
	}

	rows, err := con.Query("select * from " + configValues.DB_NAME)
	//error handling
	if err != nil {
		log.Println(err)
	}
	mySQLClient = con
	var ida string
	var idb uint
	for rows.Next() {
		err = rows.Scan(&ida, &idb)

		//error handling
		if err != nil {
			log.Println(err)
		}
		if ida == "crtx" {
			currentTXCount = idb

		} else if ida == "pbtco" {
			processedBTCOutbound = idb
		} else if ida == "signum" {
			SigNumber = idb
		}
	}

}

func fromStringToPubKey(passedStr string) *bec.PublicKey {
	koblitzCurveA := bec.S256()
	decodedBytes, _ := hex.DecodeString(passedStr)
	pubKeyFinal, _ := bec.ParsePubKey(decodedBytes, koblitzCurveA)
	return pubKeyFinal
}
func fromStrToAddr(submittedAddr string) btcutil.Address {
	temp, _ := btcutil.DecodeAddress(submittedAddr, &secondCfg.MainNetParams)
	return temp
}

func generateInitial() {
	//privKeys: privateKey1, privateKey2, privateKey3 (type privKey)
	//pubKeys: pubKey1, pubKey2, pubKey3 (type string)
	log.Println("Generating Initial Keys")
	seed := bip39.NewSeed(mneumonicStr, "")
	extKey, _ := bip32.NewMaster(seed, &chaincfg.MainNet)
	var adjPath string = "2147483648/0/10" + fmt.Sprint(int(personalRoot))
	child, _ := extKey.DeriveChildFromPath(adjPath)
	privateKey1, _ = child.ECPrivKey()
	pubKey, _ := child.ECPubKey()
	pubKey1 = hex.EncodeToString(pubKey.SerialiseCompressed()[:])
	adjPath = "2147483648/0/10" + fmt.Sprint(int(personalRoot)+1)
	child, _ = extKey.DeriveChildFromPath(adjPath)
	privateKey2, _ = child.ECPrivKey()
	pubKey, _ = child.ECPubKey()
	pubKey2 = hex.EncodeToString(pubKey.SerialiseCompressed()[:])
	adjPath = "2147483648/0/10" + fmt.Sprint(int(personalRoot)+2)
	child, _ = extKey.DeriveChildFromPath(adjPath)
	privateKey3, _ = child.ECPrivKey()
	pubKey, _ = child.ECPubKey()
	pubKey3 = hex.EncodeToString(pubKey.SerialiseCompressed()[:])
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
	relativePubKeyArray = []string{pubKey1, pubKey2, pubKey3}
	cNID_pubKeys[uint(localNodeID)] = []string{pubKey1, pubKey2, pubKey3}
	seed := bip39.NewSeed(mneumonicStr, "")
	extKey, _ := bip32.NewMaster(seed, &chaincfg.MainNet)
	for i := 3; i < int(recommendedAddressAmount); i++ {
		//TODO bring back to hardened
		tPath := "0/0/10" + fmt.Sprint(int(personalRoot)+i)
		childDerived, _ := extKey.DeriveChildFromPath(tPath)
		tPrivKey, _ := childDerived.ECPrivKey()
		tPubKeyEC, _ := childDerived.ECPubKey()
		tPubKey := hex.EncodeToString(tPubKeyEC.SerialiseCompressed()[:])
		relativePrivKeyArray = append(relativePrivKeyArray, tPrivKey)
		relativePubKeyArray = append(relativePubKeyArray, tPubKey)
		cNID_pubKeys[uint(localNodeID)] = append(cNID_pubKeys[uint(localNodeID)], tPubKey)
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
		ip_cNID[returnedIP] = currentNodeIDs[j]
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
				allRecWalletPubKeyBEC = append(allRecWalletPubKeyBEC, pubKeyHolder[0][d])
				allRecWalletPubKeyStr = append(allRecWalletPubKeyStr, hex.EncodeToString((pubKeyHolder[0][d].SerialiseCompressed()[:])))
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
		privKeyConverter[p2shAddress] = relativePrivKeyArray[indexCount+1]
		pubKeyConverter[p2shAddress] = relativePubKeyArray[indexCount+1]
		if importAddressesScope == true {
			errIA := bitcoinRPC.ImportAddressRescan(p2shAddress, "", false)
			if errIA != nil {
				log.Println("Error importing BTC P2SH address", errIA, p2shAddress)
			} else {
				log.Println("Successfully imported Address: ", p2shAddress)
			}
		}
		validMultiSigs[uint(indexCount)] = p2shAddress

		decAddress, _ := btcutil.DecodeAddress(p2shAddress, &secondCfg.MainNetParams)
		validMultiSigAddrs = append(validMultiSigAddrs, decAddress)
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
	var holderParsed []dataTransmiion
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

			var bodyParsed dataTransmiion
			err = json.Unmarshal(body, &bodyParsed)
			bodyHash := sha256.Sum256(body)
			holderParsed = append(holderParsed, bodyParsed)
			holderPubKeysHashed = append(holderPubKeysHashed, hex.EncodeToString(bodyHash[:]))
			//Hash body
			rskCounter++
		}
		cTracker := make(map[string]uint)
		var cWinner uint
		var cIndex uint
		for i := 0; i < len(holderPubKeysHashed); i++ {
			if cTracker[holderPubKeysHashed[i]] == 0 {
				cTracker[holderPubKeysHashed[i]] = 1
			} else {
				cTracker[holderPubKeysHashed[i]]++
			}
			if cTracker[holderPubKeysHashed[i]] > cWinner {
				cWinner = cTracker[holderPubKeysHashed[i]]
				cIndex = uint(i)
			}
		}

		addressConverter = holderParsed[cIndex].Addrconv
		receiveWalletPubKey = holderParsed[cIndex].PublicKeys

	}
}

// ==============MONITOR INCOMING TXS====================
func monitorIncomingBitcoinTransactions() {
	if len(validMultiSigAddrs) > 0 {
		uTXOs, err := bitcoinRPC.ListUnspentMinMaxAddresses(1, 999, validMultiSigAddrs)
		if err != nil {
			log.Println("Error listing UTXOs")
		}

		for i := 0; i < len(uTXOs); i++ {

			uTXO_satoshiAmount := uint(uTXOs[i].Amount * 100000000)
			if uTXO_satoshiAmount > 50000 {
				//TODO estimate
				gasFee := 6000
				rScript := redeemScriptCoverter[uTXOs[i].Address]
				log.Println(uTXOs[i].Address, rScript)
				crinp := bitcoindclient.CreateRawTransactionReqInputs{TxID: uTXOs[i].TxID, Vout: float64(uTXOs[i].Vout)}
				outpmap := make(map[string]float64)
				outpmap[multiSigDepositAddress] = uTXOs[i].Amount - float64(float64(gasFee)/100000000)
				crout := bitcoindclient.CreateRawTransactionReqOutputs{A: outpmap}
				thr := bitcoindclient.CreateRawTransactionReq{Inputs: []bitcoindclient.CreateRawTransactionReqInputs{crinp}, Outputs: []bitcoindclient.CreateRawTransactionReqOutputs{crout}, LockTime: float64(0), Replaceable: false}
				txres, erry := bc.CreateRawTransaction(context.Background(), thr)
				if erry != nil {
					// Handle err
					log.Println("creatingrawtx error:", erry)
				} else {
					log.Println("Success creating")

				}

				bPK1 := privKeyConverter[uTXOs[i].Address]

				lpp := bPK1.Serialise()
				kj := secp256k1.PrivKeyFromBytes(lpp)
				frrq, _ := btcutil.NewWIF(kj, &secondCfg.MainNetParams, true)
				krt := bitcoindclient.SignRawTransactionWithKeyReqPrevTxs{TxID: uTXOs[i].TxID, Vout: float64(uTXOs[i].Vout), ScriptPubkey: uTXOs[i].ScriptPubKey, RedeemScript: rScript}
				var svvr []bitcoindclient.SignRawTransactionWithKeyReqPrevTxs = []bitcoindclient.SignRawTransactionWithKeyReqPrevTxs{krt}
				jreq := bitcoindclient.SignRawTransactionWithKeyReq{HexString: txres.Hex, Privkeys: []string{frrq.String()}, PrevTxs: svvr}
				refy, err := bc.SignRawTransactionWithKey(context.Background(), jreq)

				if err != nil {
					// Handle err
					log.Println("signingrawtx error:", err)
				} else {
					log.Println("Success signing")
				}

				recPSBTs[txres.Hex] = svvr

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
					submTX, err := preInteractionContract.SubmitConsensus(auth, uTXOs[0].Address, uTXOs[0].Address, new(big.Int).SetUint64(uint64(uTXO_satoshiAmount-bitcoinTxFee)))
					if err != nil {
						log.Println("Error submitting to Polygon", err)
					} else {
						log.Println("TX hash for polygon: ", submTX.Hash().String())
					}
					if reflect.DeepEqual(earlyRec[txres.Hex], svvr) {
						psbt := holdingTable[txres.Hex]
						bPK := privKeyConverter[uTXOs[i].Address]
						lpp1 := bPK.Serialise()
						kj1 := secp256k1.PrivKeyFromBytes(lpp1)
						frrq, _ := btcutil.NewWIF(kj1, &secondCfg.MainNetParams, true)
						jreq2 := bitcoindclient.SignRawTransactionWithKeyReq{HexString: psbt, Privkeys: []string{frrq.String()}, PrevTxs: svvr}
						refy2, err := bc.SignRawTransactionWithKey(context.Background(), jreq2)

						if alreadyProcessed[txres.Hex] == false && err == nil {

							alreadyProcessed[txres.Hex] = true

							hrtt := 0.0
							var maxFee *float64 = &hrtt

							txHexToSend := bitcoindclient.SendRawTransactionReq{HexString: refy2.Hex, MaxFeeRate: maxFee}
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
						tempStruct := psbtFullStruct{refy.Hex, uint(indexConverter[uTXOs[0].Address]), svvr, txres.Hex}
						for i := 0; i < len(validNodesIP); i++ {
							if ip_cNID[validNodesIP[i]] != localNodeID {
								jsonBytes, _ := json.Marshal(tempStruct)
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
		var amtToSend uint

		amtToSend = uint(txStruct.SatsAmount.Uint64())

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

			for i := 0; (i < len(uTXOList)) && (uTXOSatAmount < amtToSend+gasFee); i++ {
				specificUTXO := bitcoindclient.CreateRawTransactionReqInputs{TxID: uTXOList[i].TxID, Vout: float64(uTXOList[i].Vout)}
				uTXO = append(uTXO, specificUTXO)
				uTXOSatAmount += uint(uTXOList[i].Amount * 100000000)

				krt := bitcoindclient.SignRawTransactionWithKeyReqPrevTxs{TxID: uTXOList[i].TxID, Vout: float64(uTXOList[i].Vout), ScriptPubkey: uTXOList[i].ScriptPubKey, RedeemScript: redeemScriptCoverter[multiSigDepositAddress]}
				svvr = append(svvr, krt)
			}

			uTXO_satoshiAmount := txStruct.SatsAmount.Uint64()
			if uTXO_satoshiAmount > 50000 {
				//TODO estimate
				gasFee := 6000
				rScript := redeemScriptCoverter[multiSigDepositAddress]

				log.Println("rscript", rScript)
				outpmap := make(map[string]float64)
				outpmap[decodedAddress] = float64(float64(int(uTXO_satoshiAmount)-gasFee) / 100000000)
				outpmap[multiSigDepositAddress] = float64(float64((uTXOSatAmount-uint(uTXO_satoshiAmount))-uint(gasFee)) / 100000000)
				crout := bitcoindclient.CreateRawTransactionReqOutputs{A: outpmap}
				thr := bitcoindclient.CreateRawTransactionReq{Inputs: uTXO, Outputs: []bitcoindclient.CreateRawTransactionReqOutputs{crout}, LockTime: float64(0), Replaceable: false}
				txres, erry := bc.CreateRawTransaction(context.Background(), thr)
				if erry != nil {
					// Handle err
					log.Println("creatingrawtx error:", erry)
				} else {
					log.Println("Success creating")
					log.Println(txres.Hex)
				}

				lpp := recWalletPrivateKey.Serialise()
				kj := secp256k1.PrivKeyFromBytes(lpp)
				frrq, _ := btcutil.NewWIF(kj, &secondCfg.MainNetParams, true)
				jreq := bitcoindclient.SignRawTransactionWithKeyReq{HexString: txres.Hex, Privkeys: []string{frrq.String()}, PrevTxs: svvr}
				refy, _ := bc.SignRawTransactionWithKey(context.Background(), jreq)

				recPSBTs[txres.Hex] = svvr

				if reflect.DeepEqual(earlyRec[txres.Hex], svvr) {
					nFSBT := holdingTable[txres.Hex]

					lpp1 := recWalletPrivateKey.Serialise()
					kj1 := secp256k1.PrivKeyFromBytes(lpp1)
					frrq, _ := btcutil.NewWIF(kj1, &secondCfg.MainNetParams, true)
					jreq2 := bitcoindclient.SignRawTransactionWithKeyReq{HexString: nFSBT, Privkeys: []string{frrq.String()}, PrevTxs: svvr}
					refy2, err := bc.SignRawTransactionWithKey(context.Background(), jreq2)

					if alreadyProcessed[txres.Hex] == false && err == nil {
						alreadyProcessed[txres.Hex] = true
						hrtt := 0.1
						var maxFee *float64 = &hrtt
						txHexToSend := bitcoindclient.SendRawTransactionReq{HexString: refy2.Hex, MaxFeeRate: maxFee}
						rawSent, err := bc.SendRawTransaction(context.Background(), txHexToSend)
						if err != nil {
							log.Println("Error broadcasting raw BTC transaction")
						} else {
							log.Println("Successfully broadcast raw BTC transaction: ", rawSent.Hex)
						}
					}

				} else {
					//if not already recieved, pass it to other nodes
					tempStruct := psbtFullStruct{refy.Hex, uint(999), svvr, txres.Hex}
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
	var decodedBody psbtFullStruct
	body, err := io.ReadAll(req.Body)
	err = json.Unmarshal(body, &decodedBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	var sPK uint = decodedBody.Spk
	psbtHex := decodedBody.Psbt
	txvh := decodedBody.TxVirginHex
	if err != nil {
		panic(err)
	}
	svvr := decodedBody.Svvra
	respBytes, _ := hex.DecodeString("success")
	rw.Write(respBytes)
	log.Println(reflect.DeepEqual(recPSBTs[txvh], decodedBody.Svvra))

	log.Println("hexVirgin: ", txvh, "psbtHex: ", psbtHex, recPSBTs[txvh], decodedBody.Svvra, sPK)
	if reflect.DeepEqual(recPSBTs[txvh], decodedBody.Svvra) {
		//make sure result is good
		lpp := recWalletPrivateKey.Serialise()
		if sPK == 999 {
			lpp = relativePrivKeyArray[sPK].Serialise()
		}
		kj := secp256k1.PrivKeyFromBytes(lpp)
		frrq, _ := btcutil.NewWIF(kj, &secondCfg.MainNetParams, true)
		jreq := bitcoindclient.SignRawTransactionWithKeyReq{HexString: psbtHex, Privkeys: []string{frrq.String()}, PrevTxs: svvr}
		refy, err := bc.SignRawTransactionWithKey(context.Background(), jreq)
		if err != nil {
			log.Println(err)
		}
		log.Println(988, alreadyProcessed[txvh])
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
					submtxop, errop := polygonBTCOverallContract.OraclePing(auth, uint8(7))
					if errop != nil {
						log.Println("Error making Oracle Ping")
					} else {
						log.Println("Successfully swapmined: ", submtxop.Hash())
					}

				}
			}

		})
	} else {
		earlyRec[txvh] = svvr
		holdingTable[txvh] = psbtHex
	}

}

//==================================SHARE PUBLICKEYS AND DEAL WITH MULTISIGS BEING SUBMITTED TO POLYGON==================

//subscribe to Polygon
//make recvnodepubkeys valid endpoint temporarily

func subscribeToPolygon() {
	sentSigBool, err := preInteractionContract.SignalSent(&bind.CallOpts{Pending: false, From: polygonAccountType, BlockNumber: nil, Context: nil}, big.NewInt(int64(SigNumber)))
	if err != nil {
		log.Println(err)
	}
	var inputsK uint
	inputsK = uint(sentSigBool)
	if inputsK == 2 || inputsK == 3 {
		SigNumber++
		openEndpoint = true
		time.AfterFunc(300*time.Second, func() {
			openEndpoint = false
		})
		cNID_pubKeys = make(map[uint][]string)
		listenGL = make(map[string]bool)
		relativePrivKeyArray = make([]*bec.PrivateKey, 0)
		cNID_pubKeys[uint(localNodeID)] = make([]string, 0)

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
		relativePubKeyArray = []string{receiveWalletPubKeyString, pubKey1, pubKey2, pubKey3}
		cNID_pubKeys[uint(localNodeID)] = []string{receiveWalletPubKeyString, pubKey1, pubKey2, pubKey3}
		for i := 3; uint(i) < recommendedAddressAmount; i++ {
			var pathY string = "2147483648/0/10" + fmt.Sprint(uint(personalRoot)+uint(i))
			child, _ := extKey.DeriveChildFromPath(pathY)
			privKey, _ := child.ECPrivKey()
			pubKey, _ := child.ECPubKey()
			pubKeyStr := hex.EncodeToString((pubKey.SerialiseCompressed()[:]))
			relativePrivKeyArray = append(relativePrivKeyArray, privKey)
			relativePubKeyArray = append(relativePubKeyArray, pubKeyStr)
			cNID_pubKeys[uint(localNodeID)] = append(cNID_pubKeys[uint(localNodeID)], pubKeyStr)
		}
		validNodesIP = make([]string, 0)

		//query (ip+port*)s from smart contract *if necessary
		for u := 0; u < len(currentNodeIDs); u++ {
			retIP, err := preInteractionContract.CnidIp(&bind.CallOpts{Pending: false, From: polygonAccountType, BlockNumber: nil, Context: nil}, currentNodeIDs[u])
			if err != nil {
				log.Printf("Error retrieving IP for Node %d", currentNodeIDs[u])
			}
			validNodesIP = append(validNodesIP, retIP)
			ip_cNID[retIP] = currentNodeIDs[u]
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
			if ip_cNID[validNodesIP[j]] != localNodeID {
				if inputsK == 3 {

					jsonBytes, _ := json.Marshal(recvFullStruct{localPortNum, relativePubKeyArray[(len(relativePubKeyArray) - 4):], true})
					payload = bytes.NewBuffer(jsonBytes)

				} else {
					jsonBytes, _ := json.Marshal(recvFullStruct{localPortNum, relativePubKeyArray, false})
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
		_, qErr := mySQLClient.Query(`UPDATE `+configValues.DB_NAME+` SET cval = ? WHERE name = 'signum'`, SigNumber)
		if qErr != nil {
			log.Println(qErr)
		} else {
			log.Printf("Updated signum %d", SigNumber)
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
	var dataSendFPK dataTransmiion
	dataSendFPK = dataTransmiion{addressConverter, receiveWalletPubKey}
	jsonBytes, _ := json.Marshal(dataSendFPK)
	rw.Write(jsonBytes)
}
func recvnodepubkeysPOST(rw http.ResponseWriter, req *http.Request) {
	if openEndpoint == true {
		tbSent, _ := hex.DecodeString("success")
		rw.Write(tbSent)
		var decodedBody recvFullStruct
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
		/*
			//Removed as this is vulnerable to spoofing
			userRemoteIP := req.Header.Get("X-Real-Ip")
			if userRemoteIP == "" {
				userRemoteIP = req.Header.Get("X-Forwarded-For")
			}
			if userRemoteIP == "" {
				userRemoteIP = req.RemoteAddr
			}
		*/
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
			if ip_cNID[userRemoteIP+":"+fmt.Sprint(portTEST)] != 9 {
				cNID_pubKeys[uint(ip_cNID[userRemoteIP+":"+fmt.Sprint(portTEST)])] = pubKeyArray
			}
			var checkFlag bool = true
			var lowestNum uint = 0
			for i := 0; i < len(currentNodeIDs); i++ {

				if cNID_pubKeys[uint(currentNodeIDs[i])] == nil || len(cNID_pubKeys[uint(currentNodeIDs[i])]) < 3 {
					checkFlag = false
				} else {
					if lowestNum == 0 {
						lowestNum = uint(len(cNID_pubKeys[uint(currentNodeIDs[i])]))
					} else if uint(len(cNID_pubKeys[uint(currentNodeIDs[i])])) < lowestNum {
						lowestNum = uint(len(cNID_pubKeys[uint(currentNodeIDs[i])]))
					}
				}
			}

			if checkFlag == true && lowestNum != 0 {
				for k := 0; uint(k) < lowestNum; k++ {

					publicKeysNew := make([]*bec.PublicKey, 0)
					for j := 0; j < len(currentNodeIDs); j++ {
						decPubKey := fromStringToPubKey(cNID_pubKeys[uint(j)][k])
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
						privKeyConverter[multisigaddrRecNew] = relativePrivKeyArray[k]
						pubKeyConverter[multisigaddrRecNew] = relativePubKeyArray[k]
						addressRecNew, _ := btcutil.DecodeAddress(multisigaddrRecNew, &secondCfg.MainNetParams)
						validMultiSigAddrs = append(validMultiSigAddrs, addressRecNew)
						validMultiSigs[uint(len(validMultiSigAddrs)-1)] = multisigaddrRecNew
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
			//TODO send TX
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
