package main

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	overallContractMask "morphswap/node/overallContractMask"
	postInteractionContract "morphswap/node/postInteractionContract"
	"morphswap/node/setup"
	"os"
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	secondCfg "github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	bitcoindclient "github.com/joakimofv/go-bitcoindclient/v23"
	"github.com/libsv/go-bk/bec"
	"github.com/spf13/viper"
)

// =========================================================================
// Utility functions
// =========================================================================
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
func (s rPubKeys) Len() int {
	return len(s)
}
func (s rPubKeys) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s rPubKeys) Less(i, j int) bool {
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

// ======================================================================
// Config setup
// ======================================================================
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

//===================================================================================
//Connection functions
//===================================================================================

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
			SignalNumber = idb
		}
	}

}
