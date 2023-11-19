package setup

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type ConfigValues struct {
	PORT              uint
	EXTERNALIP        string
	ADDR              string
	PRIV_KEY          string
	HTTP_URL          string
	CHAIN_ID          uint
	NODE_ID           uint
	DB_URI            string
	DB_NAME           string
	CONTRACT_ADDR     string
	CONTRACTBTC_ADDR  string
	CONTRACTPOST_ADDR string
	CONTRACTPRE_ADDR  string
	BTC_MNEUMONIC     string
	BTC_WALLET_ROOT   uint
	BTCNODE_HOST      string
	BTCNODE_USER      string
	BTCNODE_PASS      string
}

func SetupWalkthrough() ConfigValues {
	var currentlyConfig ConfigValues
	fmt.Print("This is the initial set up process when first starting up a Morphswap node. \n If you would like to instead populate the config file directly rather than following these prompts, feel free to exit. Otherwise, please proceed. \n What is the port the Morphswap node will be accessbile from? ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input")
		os.Exit(4)
	}

	input = strings.TrimSuffix(input, "\n")
	halpy, _ := new(big.Int).SetString(input, 10)
	currentlyConfig.PORT = uint(halpy.Uint64())

	fmt.Print("Please enter the IP address or URL of this node: ")
	reader1 := bufio.NewReader(os.Stdin)
	input1, err := reader1.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		os.Exit(4)
	}

	currentlyConfig.EXTERNALIP = strings.TrimSuffix(input1, "\n")
	fmt.Print("What is the on-chain central-chain address of this Morphswap node?: ")
	reader2 := bufio.NewReader(os.Stdin)
	input2, err := reader2.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		os.Exit(4)
	}

	currentlyConfig.ADDR = strings.TrimSuffix(input2, "\n")

	fmt.Print("What is that address's private key (NOT 0x-prefixed): ")
	reader3 := bufio.NewReader(os.Stdin)
	input3, err := reader3.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		os.Exit(4)
	}

	currentlyConfig.PRIV_KEY = strings.TrimSuffix(input3, "\n")

	fmt.Print("What is the HTTP URL of the RPC you will be using? : ")
	reader4 := bufio.NewReader(os.Stdin)
	input4, err := reader4.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		os.Exit(4)
	}

	currentlyConfig.HTTP_URL = strings.TrimSuffix(input4, "\n")

	fmt.Print("What is the chain ID you will be connecting to as the central chain (this is should usually be 137 by default)? : ")
	reader5 := bufio.NewReader(os.Stdin)
	input5, err := reader5.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		os.Exit(4)
	}

	input5 = strings.TrimSuffix(input5, "\n")
	cId, _ := new(big.Int).SetString(input5, 10)
	currentlyConfig.CHAIN_ID = uint(cId.Uint64())

	fmt.Print("What is the ID of this node? : ")
	reader6 := bufio.NewReader(os.Stdin)
	input6, err := reader6.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		os.Exit(4)
	}

	// remove the delimeter from the string
	input6 = strings.TrimSuffix(input6, "\n")
	//convert from string to bigInt
	nodeId, _ := new(big.Int).SetString(input6, 10)
	//convert from bigInt to uint
	currentlyConfig.NODE_ID = uint(nodeId.Uint64())

	fmt.Print("What is the database URI?: ")
	reader8 := bufio.NewReader(os.Stdin)
	input8, err := reader8.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		os.Exit(4)
	}

	currentlyConfig.DB_URI = strings.TrimSuffix(input8, "\n")

	fmt.Print("What is the overall contract address for chain " + strconv.FormatUint(uint64(currentlyConfig.CHAIN_ID), 10) + "?: ")
	reader9 := bufio.NewReader(os.Stdin)
	input9, err := reader9.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		os.Exit(4)
	}

	currentlyConfig.CONTRACT_ADDR = strings.TrimSuffix(input9, "\n")

	fmt.Print("What is the mneumonic for the node's bitcoin wallet? : ")
	reader15 := bufio.NewReader(os.Stdin)
	input15, err := reader15.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		os.Exit(4)
	}

	currentlyConfig.BTC_MNEUMONIC = strings.TrimSuffix(input15, "\n")

	fmt.Print("What is the ID of this node? : ")
	reader14 := bufio.NewReader(os.Stdin)
	input14, err := reader14.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		os.Exit(4)
	}

	// remove the delimeter from the string
	input14 = strings.TrimSuffix(input14, "\n")
	//convert from string to bigInt
	rootNumber, _ := new(big.Int).SetString(input14, 10)
	//convert from bigInt to uint
	currentlyConfig.BTC_WALLET_ROOT = uint(rootNumber.Uint64())

	fmt.Print("What is the address of the PreInteraction Contract? : ")
	reader = bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		os.Exit(4)
	}
	currentlyConfig.CONTRACTPRE_ADDR = strings.TrimSuffix(input, "\n")
	fmt.Print("What is the address of the PostInteraction Contract? : ")
	reader = bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		os.Exit(4)
	}
	currentlyConfig.CONTRACTPOST_ADDR = strings.TrimSuffix(input, "\n")
	fmt.Print("What is the address of the Bitcoin Central Contract (deployed on the central chain)? : ")
	reader = bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		os.Exit(4)
	}
	currentlyConfig.CONTRACTBTC_ADDR = strings.TrimSuffix(input, "\n")

	fmt.Print("What is the hostname or URL of the Bitcoin Node this Morphswap node will be integrated with? : ")
	reader = bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		os.Exit(4)
	}
	currentlyConfig.BTCNODE_HOST = strings.TrimSuffix(input, "\n")
	fmt.Print("What is the username of the Bitcoin Node? : ")
	reader = bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		os.Exit(4)
	}
	currentlyConfig.BTCNODE_USER = strings.TrimSuffix(input, "\n")
	fmt.Print("What is the password of the Bitcoin Node? : ")
	reader = bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		os.Exit(4)
	}
	currentlyConfig.BTCNODE_PASS = strings.TrimSuffix(input, "\n")

	tbwString := []byte("EXTERNAL_PORT: " + input + "\n" +
		"EXTERNAL_IP: " + ("\"" + currentlyConfig.EXTERNALIP + "\"") + "\n" +
		"ADDR: " + "\"" + currentlyConfig.ADDR + "\"" + "\n" +
		"PRIV_KEY: " + "\"" + currentlyConfig.PRIV_KEY + "\"" + "\n" +
		"HTTP_URL: " + "\"" + currentlyConfig.HTTP_URL + "\"" + "\n" +
		"CHAIN_ID: " + fmt.Sprint(currentlyConfig.CHAIN_ID) + "\n" +
		"NODE_ID: " + fmt.Sprint(currentlyConfig.NODE_ID) + "\n" +
		"DB_URI: " + "\"" + currentlyConfig.DB_URI + "\"" + "\n" +
		"DB_NAME: " + "\"" + currentlyConfig.DB_NAME + "\"" + "\n" +
		"CONTRACT_ADDR: " + "\"" + currentlyConfig.CONTRACT_ADDR + "\"" + "\n" +
		"CONTRACTBTC_ADDR: " + "\"" + currentlyConfig.CONTRACTBTC_ADDR + "\"" + "\n" +
		"CONTRACTPOST_ADDR: " + "\"" + currentlyConfig.CONTRACTPOST_ADDR + "\"" + "\n" +
		"CONTRACTPRE_ADDR: " + "\"" + currentlyConfig.CONTRACTPRE_ADDR + "\"" + "\n" +
		"BTC_MNEUMONIC: " + "\"" + currentlyConfig.BTC_MNEUMONIC + "\"" + "\n" +
		"BTC_WALLET_ROOT: " + "\"" + fmt.Sprint(currentlyConfig.BTC_WALLET_ROOT) + "\"" + "\n" +
		"BTCNODE_HOST: " + "\"" + currentlyConfig.BTCNODE_HOST + "\"" + "\n" +
		"BTCNODE_USER: " + "\"" + currentlyConfig.BTCNODE_USER + "\"" + "\n" +
		"BTCNODE_PASS: " + "\"" + currentlyConfig.BTCNODE_PASS + "\"" + "\n" +
		"\n")
	err = os.WriteFile("config.yaml", tbwString, 0644)
	if err != nil {
		panic(err)
	}
	return currentlyConfig
}
func ReadConfig() ConfigValues {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var currentlyConfig ConfigValues
	// populate global variables with values from config file (config.yaml)
	currentlyConfig.EXTERNALIP = viper.GetString("EXTERNAL_IP")
	currentlyConfig.ADDR = viper.GetString("ADDR")
	currentlyConfig.PRIV_KEY = viper.GetString("PRIV_KEY")
	currentlyConfig.HTTP_URL = viper.GetString("HTTP_URL")
	currentlyConfig.CHAIN_ID = viper.GetUint("CHAIN_ID")
	currentlyConfig.NODE_ID = viper.GetUint("NODE_ID")
	currentlyConfig.DB_URI = viper.GetString("DB_URI")
	currentlyConfig.DB_NAME = viper.GetString("DB_NAME")
	currentlyConfig.CONTRACT_ADDR = viper.GetString("CONTRACT_ADDR")
	currentlyConfig.CONTRACTBTC_ADDR = viper.GetString("CONTRACTBTC_ADDR")
	currentlyConfig.CONTRACTPOST_ADDR = viper.GetString("CONTRACTPOST_ADDR")
	currentlyConfig.CONTRACTPRE_ADDR = viper.GetString("CONTRACTPRE_ADDR")
	currentlyConfig.BTC_MNEUMONIC = viper.GetString("BTC_MNEUMONIC")
	currentlyConfig.BTC_WALLET_ROOT = viper.GetUint("BTC_WALLET_ROOT")
	currentlyConfig.BTCNODE_HOST = viper.GetString("BTCNODE_HOST")
	currentlyConfig.BTCNODE_USER = viper.GetString("BTCNODE_USER")
	currentlyConfig.BTCNODE_PASS = viper.GetString("BTCNODE_PASS")
	return currentlyConfig
}
