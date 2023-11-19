// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package preInteraction

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PreInteractionContracttxobj is an auto generated low-level Go binding around an user-defined struct.
type PreInteractionContracttxobj struct {
	MethodId           uint8
	InternalEndChainid uint8
	C2                 *big.Int
	Multichainhop      bool
	Refbool            bool
	RefAddr            common.Address
	PairId             uint64
	Sentam             *big.Int
	FinalchainWallet   common.Address
	SecondpairId       uint64
	FirstchainAsset    common.Address
	FinalchainAsset    common.Address
	TipAm              *big.Int
	AltFee             bool
	Fsigid             uint8
	ValidatedCount     uint8
	Qtxhash            string
	Sendrar            common.Address
	Blockcc            *big.Int
}

// PreInteractionMetaData contains all meta data concerning the PreInteraction contract.
var PreInteractionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_msbtc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_specialAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pic\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_psbtc_oc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wmsbtc\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fs_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"btcaddrstr\",\"type\":\"string\"}],\"name\":\"Failed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"btcAmount\",\"type\":\"uint256\"}],\"name\":\"PostInteractionNotification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"methodId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"btcAmount\",\"type\":\"uint256\"}],\"name\":\"PreInteractionNotification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_sigtype\",\"type\":\"uint256\"}],\"name\":\"newNodesList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_sigtype\",\"type\":\"uint256\"}],\"name\":\"populateMSL\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vn\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_cn\",\"type\":\"uint8\"}],\"name\":\"addNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txhash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"c2\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"method_id\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sentam\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"internal_end_chainid\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"multichainhop\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"refbool\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"refAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pair_id\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"finalchain_wallet\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"secondpair_id\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"firstchain_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"finalchain_asset\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"tipAm\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"fsigid\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"alt_fee\",\"type\":\"bool\"}],\"name\":\"addToPreList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"btcrarray\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"hyr\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_pbb\",\"type\":\"address\"}],\"name\":\"changebtcoc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cnidArr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"cnid_ip\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"errorCodeStorage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"errorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getinids\",\"outputs\":[{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getmultisigamount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_txaddr\",\"type\":\"string\"}],\"name\":\"gettxs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"method_id\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"internal_end_chainid\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"c2\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"multichainhop\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"refbool\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"refAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pair_id\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sentam\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"finalchain_wallet\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"secondpair_id\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"firstchain_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"finalchain_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tipAm\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"alt_fee\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"fsigid\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"validatedCount\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"qtxhash\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sendrar\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockcc\",\"type\":\"uint256\"}],\"internalType\":\"structPreInteractionContract.txobj\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"hash_btcr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r_am\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_btcAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"piddy\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"uniqha\",\"type\":\"string\"}],\"name\":\"indirectRedeemLiq\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inidArr\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lowLevelDataStorage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"btcr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"vnum\",\"type\":\"uint256\"}],\"name\":\"populateRecBTCaddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_btcrarray\",\"type\":\"string[]\"}],\"name\":\"populateRecBTCaddresses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reasonCodeStorage\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetBTCfromOC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vnum\",\"type\":\"uint256\"}],\"name\":\"resetbtcr\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vn\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_cn\",\"type\":\"uint8\"}],\"name\":\"rmNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sigtype\",\"type\":\"uint256\"}],\"name\":\"sendSignal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dc\",\"type\":\"address\"}],\"name\":\"setDelegatec\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"intid\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"ipstr\",\"type\":\"string\"}],\"name\":\"setcnipip\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newdeleg\",\"type\":\"address\"}],\"name\":\"setde\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_msa\",\"type\":\"uint256\"}],\"name\":\"setmultisigamount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signalSent\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"hymn\",\"type\":\"uint8\"}],\"name\":\"specialadmitfunction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_btcAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_btcAddressR\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"satsam\",\"type\":\"uint256\"}],\"name\":\"submitConsensus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txha\",\"type\":\"string\"}],\"name\":\"trackdest\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"txtracker\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawequal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PreInteractionABI is the input ABI used to generate the binding from.
// Deprecated: Use PreInteractionMetaData.ABI instead.
var PreInteractionABI = PreInteractionMetaData.ABI

// PreInteraction is an auto generated Go binding around an Ethereum contract.
type PreInteraction struct {
	PreInteractionCaller     // Read-only binding to the contract
	PreInteractionTransactor // Write-only binding to the contract
	PreInteractionFilterer   // Log filterer for contract events
}

// PreInteractionCaller is an auto generated read-only Go binding around an Ethereum contract.
type PreInteractionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreInteractionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PreInteractionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreInteractionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PreInteractionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreInteractionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PreInteractionSession struct {
	Contract     *PreInteraction   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PreInteractionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PreInteractionCallerSession struct {
	Contract *PreInteractionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PreInteractionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PreInteractionTransactorSession struct {
	Contract     *PreInteractionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PreInteractionRaw is an auto generated low-level Go binding around an Ethereum contract.
type PreInteractionRaw struct {
	Contract *PreInteraction // Generic contract binding to access the raw methods on
}

// PreInteractionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PreInteractionCallerRaw struct {
	Contract *PreInteractionCaller // Generic read-only contract binding to access the raw methods on
}

// PreInteractionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PreInteractionTransactorRaw struct {
	Contract *PreInteractionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPreInteraction creates a new instance of PreInteraction, bound to a specific deployed contract.
func NewPreInteraction(address common.Address, backend bind.ContractBackend) (*PreInteraction, error) {
	contract, err := bindPreInteraction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PreInteraction{PreInteractionCaller: PreInteractionCaller{contract: contract}, PreInteractionTransactor: PreInteractionTransactor{contract: contract}, PreInteractionFilterer: PreInteractionFilterer{contract: contract}}, nil
}

// NewPreInteractionCaller creates a new read-only instance of PreInteraction, bound to a specific deployed contract.
func NewPreInteractionCaller(address common.Address, caller bind.ContractCaller) (*PreInteractionCaller, error) {
	contract, err := bindPreInteraction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PreInteractionCaller{contract: contract}, nil
}

// NewPreInteractionTransactor creates a new write-only instance of PreInteraction, bound to a specific deployed contract.
func NewPreInteractionTransactor(address common.Address, transactor bind.ContractTransactor) (*PreInteractionTransactor, error) {
	contract, err := bindPreInteraction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PreInteractionTransactor{contract: contract}, nil
}

// NewPreInteractionFilterer creates a new log filterer instance of PreInteraction, bound to a specific deployed contract.
func NewPreInteractionFilterer(address common.Address, filterer bind.ContractFilterer) (*PreInteractionFilterer, error) {
	contract, err := bindPreInteraction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PreInteractionFilterer{contract: contract}, nil
}

// bindPreInteraction binds a generic wrapper to an already deployed contract.
func bindPreInteraction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PreInteractionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PreInteraction *PreInteractionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PreInteraction.Contract.PreInteractionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PreInteraction *PreInteractionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreInteraction.Contract.PreInteractionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PreInteraction *PreInteractionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PreInteraction.Contract.PreInteractionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PreInteraction *PreInteractionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PreInteraction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PreInteraction *PreInteractionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreInteraction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PreInteraction *PreInteractionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PreInteraction.Contract.contract.Transact(opts, method, params...)
}

// Btcrarray is a free data retrieval call binding the contract method 0x89663e1c.
//
// Solidity: function btcrarray(uint256 ) view returns(string)
func (_PreInteraction *PreInteractionCaller) Btcrarray(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _PreInteraction.contract.Call(opts, &out, "btcrarray", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Btcrarray is a free data retrieval call binding the contract method 0x89663e1c.
//
// Solidity: function btcrarray(uint256 ) view returns(string)
func (_PreInteraction *PreInteractionSession) Btcrarray(arg0 *big.Int) (string, error) {
	return _PreInteraction.Contract.Btcrarray(&_PreInteraction.CallOpts, arg0)
}

// Btcrarray is a free data retrieval call binding the contract method 0x89663e1c.
//
// Solidity: function btcrarray(uint256 ) view returns(string)
func (_PreInteraction *PreInteractionCallerSession) Btcrarray(arg0 *big.Int) (string, error) {
	return _PreInteraction.Contract.Btcrarray(&_PreInteraction.CallOpts, arg0)
}

// CnidArr is a free data retrieval call binding the contract method 0x69243207.
//
// Solidity: function cnidArr(uint256 ) view returns(address)
func (_PreInteraction *PreInteractionCaller) CnidArr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PreInteraction.contract.Call(opts, &out, "cnidArr", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CnidArr is a free data retrieval call binding the contract method 0x69243207.
//
// Solidity: function cnidArr(uint256 ) view returns(address)
func (_PreInteraction *PreInteractionSession) CnidArr(arg0 *big.Int) (common.Address, error) {
	return _PreInteraction.Contract.CnidArr(&_PreInteraction.CallOpts, arg0)
}

// CnidArr is a free data retrieval call binding the contract method 0x69243207.
//
// Solidity: function cnidArr(uint256 ) view returns(address)
func (_PreInteraction *PreInteractionCallerSession) CnidArr(arg0 *big.Int) (common.Address, error) {
	return _PreInteraction.Contract.CnidArr(&_PreInteraction.CallOpts, arg0)
}

// CnidIp is a free data retrieval call binding the contract method 0x379337e0.
//
// Solidity: function cnid_ip(uint8 ) view returns(string)
func (_PreInteraction *PreInteractionCaller) CnidIp(opts *bind.CallOpts, arg0 uint8) (string, error) {
	var out []interface{}
	err := _PreInteraction.contract.Call(opts, &out, "cnid_ip", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CnidIp is a free data retrieval call binding the contract method 0x379337e0.
//
// Solidity: function cnid_ip(uint8 ) view returns(string)
func (_PreInteraction *PreInteractionSession) CnidIp(arg0 uint8) (string, error) {
	return _PreInteraction.Contract.CnidIp(&_PreInteraction.CallOpts, arg0)
}

// CnidIp is a free data retrieval call binding the contract method 0x379337e0.
//
// Solidity: function cnid_ip(uint8 ) view returns(string)
func (_PreInteraction *PreInteractionCallerSession) CnidIp(arg0 uint8) (string, error) {
	return _PreInteraction.Contract.CnidIp(&_PreInteraction.CallOpts, arg0)
}

// ErrorCodeStorage is a free data retrieval call binding the contract method 0xe926881d.
//
// Solidity: function errorCodeStorage(uint256 ) view returns(uint256)
func (_PreInteraction *PreInteractionCaller) ErrorCodeStorage(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PreInteraction.contract.Call(opts, &out, "errorCodeStorage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ErrorCodeStorage is a free data retrieval call binding the contract method 0xe926881d.
//
// Solidity: function errorCodeStorage(uint256 ) view returns(uint256)
func (_PreInteraction *PreInteractionSession) ErrorCodeStorage(arg0 *big.Int) (*big.Int, error) {
	return _PreInteraction.Contract.ErrorCodeStorage(&_PreInteraction.CallOpts, arg0)
}

// ErrorCodeStorage is a free data retrieval call binding the contract method 0xe926881d.
//
// Solidity: function errorCodeStorage(uint256 ) view returns(uint256)
func (_PreInteraction *PreInteractionCallerSession) ErrorCodeStorage(arg0 *big.Int) (*big.Int, error) {
	return _PreInteraction.Contract.ErrorCodeStorage(&_PreInteraction.CallOpts, arg0)
}

// ErrorCount is a free data retrieval call binding the contract method 0xd38dbee6.
//
// Solidity: function errorCount() view returns(uint256)
func (_PreInteraction *PreInteractionCaller) ErrorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PreInteraction.contract.Call(opts, &out, "errorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ErrorCount is a free data retrieval call binding the contract method 0xd38dbee6.
//
// Solidity: function errorCount() view returns(uint256)
func (_PreInteraction *PreInteractionSession) ErrorCount() (*big.Int, error) {
	return _PreInteraction.Contract.ErrorCount(&_PreInteraction.CallOpts)
}

// ErrorCount is a free data retrieval call binding the contract method 0xd38dbee6.
//
// Solidity: function errorCount() view returns(uint256)
func (_PreInteraction *PreInteractionCallerSession) ErrorCount() (*big.Int, error) {
	return _PreInteraction.Contract.ErrorCount(&_PreInteraction.CallOpts)
}

// Getinids is a free data retrieval call binding the contract method 0xd596ebe0.
//
// Solidity: function getinids() view returns(uint8[])
func (_PreInteraction *PreInteractionCaller) Getinids(opts *bind.CallOpts) ([]uint8, error) {
	var out []interface{}
	err := _PreInteraction.contract.Call(opts, &out, "getinids")

	if err != nil {
		return *new([]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint8)).(*[]uint8)

	return out0, err

}

// Getinids is a free data retrieval call binding the contract method 0xd596ebe0.
//
// Solidity: function getinids() view returns(uint8[])
func (_PreInteraction *PreInteractionSession) Getinids() ([]uint8, error) {
	return _PreInteraction.Contract.Getinids(&_PreInteraction.CallOpts)
}

// Getinids is a free data retrieval call binding the contract method 0xd596ebe0.
//
// Solidity: function getinids() view returns(uint8[])
func (_PreInteraction *PreInteractionCallerSession) Getinids() ([]uint8, error) {
	return _PreInteraction.Contract.Getinids(&_PreInteraction.CallOpts)
}

// Getmultisigamount is a free data retrieval call binding the contract method 0xf7df79cb.
//
// Solidity: function getmultisigamount() view returns(uint256)
func (_PreInteraction *PreInteractionCaller) Getmultisigamount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PreInteraction.contract.Call(opts, &out, "getmultisigamount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Getmultisigamount is a free data retrieval call binding the contract method 0xf7df79cb.
//
// Solidity: function getmultisigamount() view returns(uint256)
func (_PreInteraction *PreInteractionSession) Getmultisigamount() (*big.Int, error) {
	return _PreInteraction.Contract.Getmultisigamount(&_PreInteraction.CallOpts)
}

// Getmultisigamount is a free data retrieval call binding the contract method 0xf7df79cb.
//
// Solidity: function getmultisigamount() view returns(uint256)
func (_PreInteraction *PreInteractionCallerSession) Getmultisigamount() (*big.Int, error) {
	return _PreInteraction.Contract.Getmultisigamount(&_PreInteraction.CallOpts)
}

// HashBtcr is a free data retrieval call binding the contract method 0x803c4480.
//
// Solidity: function hash_btcr(string ) view returns(uint256)
func (_PreInteraction *PreInteractionCaller) HashBtcr(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _PreInteraction.contract.Call(opts, &out, "hash_btcr", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashBtcr is a free data retrieval call binding the contract method 0x803c4480.
//
// Solidity: function hash_btcr(string ) view returns(uint256)
func (_PreInteraction *PreInteractionSession) HashBtcr(arg0 string) (*big.Int, error) {
	return _PreInteraction.Contract.HashBtcr(&_PreInteraction.CallOpts, arg0)
}

// HashBtcr is a free data retrieval call binding the contract method 0x803c4480.
//
// Solidity: function hash_btcr(string ) view returns(uint256)
func (_PreInteraction *PreInteractionCallerSession) HashBtcr(arg0 string) (*big.Int, error) {
	return _PreInteraction.Contract.HashBtcr(&_PreInteraction.CallOpts, arg0)
}

// InidArr is a free data retrieval call binding the contract method 0xfb6e70e7.
//
// Solidity: function inidArr(uint256 ) view returns(uint8)
func (_PreInteraction *PreInteractionCaller) InidArr(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _PreInteraction.contract.Call(opts, &out, "inidArr", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// InidArr is a free data retrieval call binding the contract method 0xfb6e70e7.
//
// Solidity: function inidArr(uint256 ) view returns(uint8)
func (_PreInteraction *PreInteractionSession) InidArr(arg0 *big.Int) (uint8, error) {
	return _PreInteraction.Contract.InidArr(&_PreInteraction.CallOpts, arg0)
}

// InidArr is a free data retrieval call binding the contract method 0xfb6e70e7.
//
// Solidity: function inidArr(uint256 ) view returns(uint8)
func (_PreInteraction *PreInteractionCallerSession) InidArr(arg0 *big.Int) (uint8, error) {
	return _PreInteraction.Contract.InidArr(&_PreInteraction.CallOpts, arg0)
}

// LowLevelDataStorage is a free data retrieval call binding the contract method 0xedf42a1d.
//
// Solidity: function lowLevelDataStorage(uint256 ) view returns(bytes)
func (_PreInteraction *PreInteractionCaller) LowLevelDataStorage(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _PreInteraction.contract.Call(opts, &out, "lowLevelDataStorage", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// LowLevelDataStorage is a free data retrieval call binding the contract method 0xedf42a1d.
//
// Solidity: function lowLevelDataStorage(uint256 ) view returns(bytes)
func (_PreInteraction *PreInteractionSession) LowLevelDataStorage(arg0 *big.Int) ([]byte, error) {
	return _PreInteraction.Contract.LowLevelDataStorage(&_PreInteraction.CallOpts, arg0)
}

// LowLevelDataStorage is a free data retrieval call binding the contract method 0xedf42a1d.
//
// Solidity: function lowLevelDataStorage(uint256 ) view returns(bytes)
func (_PreInteraction *PreInteractionCallerSession) LowLevelDataStorage(arg0 *big.Int) ([]byte, error) {
	return _PreInteraction.Contract.LowLevelDataStorage(&_PreInteraction.CallOpts, arg0)
}

// ReasonCodeStorage is a free data retrieval call binding the contract method 0xc6a3f5e0.
//
// Solidity: function reasonCodeStorage(uint256 ) view returns(string)
func (_PreInteraction *PreInteractionCaller) ReasonCodeStorage(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _PreInteraction.contract.Call(opts, &out, "reasonCodeStorage", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReasonCodeStorage is a free data retrieval call binding the contract method 0xc6a3f5e0.
//
// Solidity: function reasonCodeStorage(uint256 ) view returns(string)
func (_PreInteraction *PreInteractionSession) ReasonCodeStorage(arg0 *big.Int) (string, error) {
	return _PreInteraction.Contract.ReasonCodeStorage(&_PreInteraction.CallOpts, arg0)
}

// ReasonCodeStorage is a free data retrieval call binding the contract method 0xc6a3f5e0.
//
// Solidity: function reasonCodeStorage(uint256 ) view returns(string)
func (_PreInteraction *PreInteractionCallerSession) ReasonCodeStorage(arg0 *big.Int) (string, error) {
	return _PreInteraction.Contract.ReasonCodeStorage(&_PreInteraction.CallOpts, arg0)
}

// SignalSent is a free data retrieval call binding the contract method 0xa91d0ca3.
//
// Solidity: function signalSent(uint256 ) view returns(uint8)
func (_PreInteraction *PreInteractionCaller) SignalSent(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _PreInteraction.contract.Call(opts, &out, "signalSent", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SignalSent is a free data retrieval call binding the contract method 0xa91d0ca3.
//
// Solidity: function signalSent(uint256 ) view returns(uint8)
func (_PreInteraction *PreInteractionSession) SignalSent(arg0 *big.Int) (uint8, error) {
	return _PreInteraction.Contract.SignalSent(&_PreInteraction.CallOpts, arg0)
}

// SignalSent is a free data retrieval call binding the contract method 0xa91d0ca3.
//
// Solidity: function signalSent(uint256 ) view returns(uint8)
func (_PreInteraction *PreInteractionCallerSession) SignalSent(arg0 *big.Int) (uint8, error) {
	return _PreInteraction.Contract.SignalSent(&_PreInteraction.CallOpts, arg0)
}

// Trackdest is a free data retrieval call binding the contract method 0x7636a854.
//
// Solidity: function trackdest(string txha) view returns(string)
func (_PreInteraction *PreInteractionCaller) Trackdest(opts *bind.CallOpts, txha string) (string, error) {
	var out []interface{}
	err := _PreInteraction.contract.Call(opts, &out, "trackdest", txha)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Trackdest is a free data retrieval call binding the contract method 0x7636a854.
//
// Solidity: function trackdest(string txha) view returns(string)
func (_PreInteraction *PreInteractionSession) Trackdest(txha string) (string, error) {
	return _PreInteraction.Contract.Trackdest(&_PreInteraction.CallOpts, txha)
}

// Trackdest is a free data retrieval call binding the contract method 0x7636a854.
//
// Solidity: function trackdest(string txha) view returns(string)
func (_PreInteraction *PreInteractionCallerSession) Trackdest(txha string) (string, error) {
	return _PreInteraction.Contract.Trackdest(&_PreInteraction.CallOpts, txha)
}

// Txtracker is a free data retrieval call binding the contract method 0x300c331d.
//
// Solidity: function txtracker(address , string ) view returns(string)
func (_PreInteraction *PreInteractionCaller) Txtracker(opts *bind.CallOpts, arg0 common.Address, arg1 string) (string, error) {
	var out []interface{}
	err := _PreInteraction.contract.Call(opts, &out, "txtracker", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Txtracker is a free data retrieval call binding the contract method 0x300c331d.
//
// Solidity: function txtracker(address , string ) view returns(string)
func (_PreInteraction *PreInteractionSession) Txtracker(arg0 common.Address, arg1 string) (string, error) {
	return _PreInteraction.Contract.Txtracker(&_PreInteraction.CallOpts, arg0, arg1)
}

// Txtracker is a free data retrieval call binding the contract method 0x300c331d.
//
// Solidity: function txtracker(address , string ) view returns(string)
func (_PreInteraction *PreInteractionCallerSession) Txtracker(arg0 common.Address, arg1 string) (string, error) {
	return _PreInteraction.Contract.Txtracker(&_PreInteraction.CallOpts, arg0, arg1)
}

// AddNode is a paid mutator transaction binding the contract method 0xb6292695.
//
// Solidity: function addNode(address _vn, uint8 _cn) returns(bool)
func (_PreInteraction *PreInteractionTransactor) AddNode(opts *bind.TransactOpts, _vn common.Address, _cn uint8) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "addNode", _vn, _cn)
}

// AddNode is a paid mutator transaction binding the contract method 0xb6292695.
//
// Solidity: function addNode(address _vn, uint8 _cn) returns(bool)
func (_PreInteraction *PreInteractionSession) AddNode(_vn common.Address, _cn uint8) (*types.Transaction, error) {
	return _PreInteraction.Contract.AddNode(&_PreInteraction.TransactOpts, _vn, _cn)
}

// AddNode is a paid mutator transaction binding the contract method 0xb6292695.
//
// Solidity: function addNode(address _vn, uint8 _cn) returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) AddNode(_vn common.Address, _cn uint8) (*types.Transaction, error) {
	return _PreInteraction.Contract.AddNode(&_PreInteraction.TransactOpts, _vn, _cn)
}

// AddToPreList is a paid mutator transaction binding the contract method 0x57c871b8.
//
// Solidity: function addToPreList(string txhash, uint256 c2, uint8 method_id, uint256 sentam, uint8 internal_end_chainid, bool multichainhop, bool refbool, address refAddr, uint64 pair_id, address finalchain_wallet, uint64 secondpair_id, address firstchain_asset, address finalchain_asset, uint128 tipAm, uint8 fsigid, bool alt_fee) returns(bool)
func (_PreInteraction *PreInteractionTransactor) AddToPreList(opts *bind.TransactOpts, txhash string, c2 *big.Int, method_id uint8, sentam *big.Int, internal_end_chainid uint8, multichainhop bool, refbool bool, refAddr common.Address, pair_id uint64, finalchain_wallet common.Address, secondpair_id uint64, firstchain_asset common.Address, finalchain_asset common.Address, tipAm *big.Int, fsigid uint8, alt_fee bool) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "addToPreList", txhash, c2, method_id, sentam, internal_end_chainid, multichainhop, refbool, refAddr, pair_id, finalchain_wallet, secondpair_id, firstchain_asset, finalchain_asset, tipAm, fsigid, alt_fee)
}

// AddToPreList is a paid mutator transaction binding the contract method 0x57c871b8.
//
// Solidity: function addToPreList(string txhash, uint256 c2, uint8 method_id, uint256 sentam, uint8 internal_end_chainid, bool multichainhop, bool refbool, address refAddr, uint64 pair_id, address finalchain_wallet, uint64 secondpair_id, address firstchain_asset, address finalchain_asset, uint128 tipAm, uint8 fsigid, bool alt_fee) returns(bool)
func (_PreInteraction *PreInteractionSession) AddToPreList(txhash string, c2 *big.Int, method_id uint8, sentam *big.Int, internal_end_chainid uint8, multichainhop bool, refbool bool, refAddr common.Address, pair_id uint64, finalchain_wallet common.Address, secondpair_id uint64, firstchain_asset common.Address, finalchain_asset common.Address, tipAm *big.Int, fsigid uint8, alt_fee bool) (*types.Transaction, error) {
	return _PreInteraction.Contract.AddToPreList(&_PreInteraction.TransactOpts, txhash, c2, method_id, sentam, internal_end_chainid, multichainhop, refbool, refAddr, pair_id, finalchain_wallet, secondpair_id, firstchain_asset, finalchain_asset, tipAm, fsigid, alt_fee)
}

// AddToPreList is a paid mutator transaction binding the contract method 0x57c871b8.
//
// Solidity: function addToPreList(string txhash, uint256 c2, uint8 method_id, uint256 sentam, uint8 internal_end_chainid, bool multichainhop, bool refbool, address refAddr, uint64 pair_id, address finalchain_wallet, uint64 secondpair_id, address firstchain_asset, address finalchain_asset, uint128 tipAm, uint8 fsigid, bool alt_fee) returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) AddToPreList(txhash string, c2 *big.Int, method_id uint8, sentam *big.Int, internal_end_chainid uint8, multichainhop bool, refbool bool, refAddr common.Address, pair_id uint64, finalchain_wallet common.Address, secondpair_id uint64, firstchain_asset common.Address, finalchain_asset common.Address, tipAm *big.Int, fsigid uint8, alt_fee bool) (*types.Transaction, error) {
	return _PreInteraction.Contract.AddToPreList(&_PreInteraction.TransactOpts, txhash, c2, method_id, sentam, internal_end_chainid, multichainhop, refbool, refAddr, pair_id, finalchain_wallet, secondpair_id, firstchain_asset, finalchain_asset, tipAm, fsigid, alt_fee)
}

// Changebtcoc is a paid mutator transaction binding the contract method 0xeee8f9a7.
//
// Solidity: function changebtcoc(uint8 hyr, address _pbb) returns(bool)
func (_PreInteraction *PreInteractionTransactor) Changebtcoc(opts *bind.TransactOpts, hyr uint8, _pbb common.Address) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "changebtcoc", hyr, _pbb)
}

// Changebtcoc is a paid mutator transaction binding the contract method 0xeee8f9a7.
//
// Solidity: function changebtcoc(uint8 hyr, address _pbb) returns(bool)
func (_PreInteraction *PreInteractionSession) Changebtcoc(hyr uint8, _pbb common.Address) (*types.Transaction, error) {
	return _PreInteraction.Contract.Changebtcoc(&_PreInteraction.TransactOpts, hyr, _pbb)
}

// Changebtcoc is a paid mutator transaction binding the contract method 0xeee8f9a7.
//
// Solidity: function changebtcoc(uint8 hyr, address _pbb) returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) Changebtcoc(hyr uint8, _pbb common.Address) (*types.Transaction, error) {
	return _PreInteraction.Contract.Changebtcoc(&_PreInteraction.TransactOpts, hyr, _pbb)
}

// Gettxs is a paid mutator transaction binding the contract method 0xaa8d1610.
//
// Solidity: function gettxs(string _txaddr) returns((uint8,uint8,uint256,bool,bool,address,uint64,uint256,address,uint64,address,address,uint256,bool,uint8,uint8,string,address,uint256))
func (_PreInteraction *PreInteractionTransactor) Gettxs(opts *bind.TransactOpts, _txaddr string) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "gettxs", _txaddr)
}

// Gettxs is a paid mutator transaction binding the contract method 0xaa8d1610.
//
// Solidity: function gettxs(string _txaddr) returns((uint8,uint8,uint256,bool,bool,address,uint64,uint256,address,uint64,address,address,uint256,bool,uint8,uint8,string,address,uint256))
func (_PreInteraction *PreInteractionSession) Gettxs(_txaddr string) (*types.Transaction, error) {
	return _PreInteraction.Contract.Gettxs(&_PreInteraction.TransactOpts, _txaddr)
}

// Gettxs is a paid mutator transaction binding the contract method 0xaa8d1610.
//
// Solidity: function gettxs(string _txaddr) returns((uint8,uint8,uint256,bool,bool,address,uint64,uint256,address,uint64,address,address,uint256,bool,uint8,uint8,string,address,uint256))
func (_PreInteraction *PreInteractionTransactorSession) Gettxs(_txaddr string) (*types.Transaction, error) {
	return _PreInteraction.Contract.Gettxs(&_PreInteraction.TransactOpts, _txaddr)
}

// IndirectRedeemLiq is a paid mutator transaction binding the contract method 0xcbe203df.
//
// Solidity: function indirectRedeemLiq(uint256 r_am, string _btcAddress, uint64 piddy, string uniqha) returns(bool)
func (_PreInteraction *PreInteractionTransactor) IndirectRedeemLiq(opts *bind.TransactOpts, r_am *big.Int, _btcAddress string, piddy uint64, uniqha string) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "indirectRedeemLiq", r_am, _btcAddress, piddy, uniqha)
}

// IndirectRedeemLiq is a paid mutator transaction binding the contract method 0xcbe203df.
//
// Solidity: function indirectRedeemLiq(uint256 r_am, string _btcAddress, uint64 piddy, string uniqha) returns(bool)
func (_PreInteraction *PreInteractionSession) IndirectRedeemLiq(r_am *big.Int, _btcAddress string, piddy uint64, uniqha string) (*types.Transaction, error) {
	return _PreInteraction.Contract.IndirectRedeemLiq(&_PreInteraction.TransactOpts, r_am, _btcAddress, piddy, uniqha)
}

// IndirectRedeemLiq is a paid mutator transaction binding the contract method 0xcbe203df.
//
// Solidity: function indirectRedeemLiq(uint256 r_am, string _btcAddress, uint64 piddy, string uniqha) returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) IndirectRedeemLiq(r_am *big.Int, _btcAddress string, piddy uint64, uniqha string) (*types.Transaction, error) {
	return _PreInteraction.Contract.IndirectRedeemLiq(&_PreInteraction.TransactOpts, r_am, _btcAddress, piddy, uniqha)
}

// PopulateRecBTCaddress is a paid mutator transaction binding the contract method 0xf7d66113.
//
// Solidity: function populateRecBTCaddress(string btcr, uint256 vnum) returns(bool)
func (_PreInteraction *PreInteractionTransactor) PopulateRecBTCaddress(opts *bind.TransactOpts, btcr string, vnum *big.Int) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "populateRecBTCaddress", btcr, vnum)
}

// PopulateRecBTCaddress is a paid mutator transaction binding the contract method 0xf7d66113.
//
// Solidity: function populateRecBTCaddress(string btcr, uint256 vnum) returns(bool)
func (_PreInteraction *PreInteractionSession) PopulateRecBTCaddress(btcr string, vnum *big.Int) (*types.Transaction, error) {
	return _PreInteraction.Contract.PopulateRecBTCaddress(&_PreInteraction.TransactOpts, btcr, vnum)
}

// PopulateRecBTCaddress is a paid mutator transaction binding the contract method 0xf7d66113.
//
// Solidity: function populateRecBTCaddress(string btcr, uint256 vnum) returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) PopulateRecBTCaddress(btcr string, vnum *big.Int) (*types.Transaction, error) {
	return _PreInteraction.Contract.PopulateRecBTCaddress(&_PreInteraction.TransactOpts, btcr, vnum)
}

// PopulateRecBTCaddresses is a paid mutator transaction binding the contract method 0x1e305add.
//
// Solidity: function populateRecBTCaddresses(string[] _btcrarray) returns(bool)
func (_PreInteraction *PreInteractionTransactor) PopulateRecBTCaddresses(opts *bind.TransactOpts, _btcrarray []string) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "populateRecBTCaddresses", _btcrarray)
}

// PopulateRecBTCaddresses is a paid mutator transaction binding the contract method 0x1e305add.
//
// Solidity: function populateRecBTCaddresses(string[] _btcrarray) returns(bool)
func (_PreInteraction *PreInteractionSession) PopulateRecBTCaddresses(_btcrarray []string) (*types.Transaction, error) {
	return _PreInteraction.Contract.PopulateRecBTCaddresses(&_PreInteraction.TransactOpts, _btcrarray)
}

// PopulateRecBTCaddresses is a paid mutator transaction binding the contract method 0x1e305add.
//
// Solidity: function populateRecBTCaddresses(string[] _btcrarray) returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) PopulateRecBTCaddresses(_btcrarray []string) (*types.Transaction, error) {
	return _PreInteraction.Contract.PopulateRecBTCaddresses(&_PreInteraction.TransactOpts, _btcrarray)
}

// ResetBTCfromOC is a paid mutator transaction binding the contract method 0xf161a20c.
//
// Solidity: function resetBTCfromOC() returns(bool)
func (_PreInteraction *PreInteractionTransactor) ResetBTCfromOC(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "resetBTCfromOC")
}

// ResetBTCfromOC is a paid mutator transaction binding the contract method 0xf161a20c.
//
// Solidity: function resetBTCfromOC() returns(bool)
func (_PreInteraction *PreInteractionSession) ResetBTCfromOC() (*types.Transaction, error) {
	return _PreInteraction.Contract.ResetBTCfromOC(&_PreInteraction.TransactOpts)
}

// ResetBTCfromOC is a paid mutator transaction binding the contract method 0xf161a20c.
//
// Solidity: function resetBTCfromOC() returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) ResetBTCfromOC() (*types.Transaction, error) {
	return _PreInteraction.Contract.ResetBTCfromOC(&_PreInteraction.TransactOpts)
}

// Resetbtcr is a paid mutator transaction binding the contract method 0x34143913.
//
// Solidity: function resetbtcr(uint256 vnum) returns(bool)
func (_PreInteraction *PreInteractionTransactor) Resetbtcr(opts *bind.TransactOpts, vnum *big.Int) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "resetbtcr", vnum)
}

// Resetbtcr is a paid mutator transaction binding the contract method 0x34143913.
//
// Solidity: function resetbtcr(uint256 vnum) returns(bool)
func (_PreInteraction *PreInteractionSession) Resetbtcr(vnum *big.Int) (*types.Transaction, error) {
	return _PreInteraction.Contract.Resetbtcr(&_PreInteraction.TransactOpts, vnum)
}

// Resetbtcr is a paid mutator transaction binding the contract method 0x34143913.
//
// Solidity: function resetbtcr(uint256 vnum) returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) Resetbtcr(vnum *big.Int) (*types.Transaction, error) {
	return _PreInteraction.Contract.Resetbtcr(&_PreInteraction.TransactOpts, vnum)
}

// RmNode is a paid mutator transaction binding the contract method 0xa6d06d15.
//
// Solidity: function rmNode(address _vn, uint8 _cn) returns(bool)
func (_PreInteraction *PreInteractionTransactor) RmNode(opts *bind.TransactOpts, _vn common.Address, _cn uint8) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "rmNode", _vn, _cn)
}

// RmNode is a paid mutator transaction binding the contract method 0xa6d06d15.
//
// Solidity: function rmNode(address _vn, uint8 _cn) returns(bool)
func (_PreInteraction *PreInteractionSession) RmNode(_vn common.Address, _cn uint8) (*types.Transaction, error) {
	return _PreInteraction.Contract.RmNode(&_PreInteraction.TransactOpts, _vn, _cn)
}

// RmNode is a paid mutator transaction binding the contract method 0xa6d06d15.
//
// Solidity: function rmNode(address _vn, uint8 _cn) returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) RmNode(_vn common.Address, _cn uint8) (*types.Transaction, error) {
	return _PreInteraction.Contract.RmNode(&_PreInteraction.TransactOpts, _vn, _cn)
}

// SendSignal is a paid mutator transaction binding the contract method 0x54fd35bf.
//
// Solidity: function sendSignal(uint256 sigtype) returns(bool)
func (_PreInteraction *PreInteractionTransactor) SendSignal(opts *bind.TransactOpts, sigtype *big.Int) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "sendSignal", sigtype)
}

// SendSignal is a paid mutator transaction binding the contract method 0x54fd35bf.
//
// Solidity: function sendSignal(uint256 sigtype) returns(bool)
func (_PreInteraction *PreInteractionSession) SendSignal(sigtype *big.Int) (*types.Transaction, error) {
	return _PreInteraction.Contract.SendSignal(&_PreInteraction.TransactOpts, sigtype)
}

// SendSignal is a paid mutator transaction binding the contract method 0x54fd35bf.
//
// Solidity: function sendSignal(uint256 sigtype) returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) SendSignal(sigtype *big.Int) (*types.Transaction, error) {
	return _PreInteraction.Contract.SendSignal(&_PreInteraction.TransactOpts, sigtype)
}

// SetDelegatec is a paid mutator transaction binding the contract method 0x17bea24f.
//
// Solidity: function setDelegatec(address _dc) returns(bool)
func (_PreInteraction *PreInteractionTransactor) SetDelegatec(opts *bind.TransactOpts, _dc common.Address) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "setDelegatec", _dc)
}

// SetDelegatec is a paid mutator transaction binding the contract method 0x17bea24f.
//
// Solidity: function setDelegatec(address _dc) returns(bool)
func (_PreInteraction *PreInteractionSession) SetDelegatec(_dc common.Address) (*types.Transaction, error) {
	return _PreInteraction.Contract.SetDelegatec(&_PreInteraction.TransactOpts, _dc)
}

// SetDelegatec is a paid mutator transaction binding the contract method 0x17bea24f.
//
// Solidity: function setDelegatec(address _dc) returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) SetDelegatec(_dc common.Address) (*types.Transaction, error) {
	return _PreInteraction.Contract.SetDelegatec(&_PreInteraction.TransactOpts, _dc)
}

// Setcnipip is a paid mutator transaction binding the contract method 0xc44c7fb4.
//
// Solidity: function setcnipip(uint8 intid, string ipstr) returns(bool)
func (_PreInteraction *PreInteractionTransactor) Setcnipip(opts *bind.TransactOpts, intid uint8, ipstr string) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "setcnipip", intid, ipstr)
}

// Setcnipip is a paid mutator transaction binding the contract method 0xc44c7fb4.
//
// Solidity: function setcnipip(uint8 intid, string ipstr) returns(bool)
func (_PreInteraction *PreInteractionSession) Setcnipip(intid uint8, ipstr string) (*types.Transaction, error) {
	return _PreInteraction.Contract.Setcnipip(&_PreInteraction.TransactOpts, intid, ipstr)
}

// Setcnipip is a paid mutator transaction binding the contract method 0xc44c7fb4.
//
// Solidity: function setcnipip(uint8 intid, string ipstr) returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) Setcnipip(intid uint8, ipstr string) (*types.Transaction, error) {
	return _PreInteraction.Contract.Setcnipip(&_PreInteraction.TransactOpts, intid, ipstr)
}

// Setde is a paid mutator transaction binding the contract method 0xb07c410a.
//
// Solidity: function setde(address _newdeleg) returns(bool)
func (_PreInteraction *PreInteractionTransactor) Setde(opts *bind.TransactOpts, _newdeleg common.Address) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "setde", _newdeleg)
}

// Setde is a paid mutator transaction binding the contract method 0xb07c410a.
//
// Solidity: function setde(address _newdeleg) returns(bool)
func (_PreInteraction *PreInteractionSession) Setde(_newdeleg common.Address) (*types.Transaction, error) {
	return _PreInteraction.Contract.Setde(&_PreInteraction.TransactOpts, _newdeleg)
}

// Setde is a paid mutator transaction binding the contract method 0xb07c410a.
//
// Solidity: function setde(address _newdeleg) returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) Setde(_newdeleg common.Address) (*types.Transaction, error) {
	return _PreInteraction.Contract.Setde(&_PreInteraction.TransactOpts, _newdeleg)
}

// Setmultisigamount is a paid mutator transaction binding the contract method 0xd9ab23b3.
//
// Solidity: function setmultisigamount(uint256 _msa) returns(bool)
func (_PreInteraction *PreInteractionTransactor) Setmultisigamount(opts *bind.TransactOpts, _msa *big.Int) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "setmultisigamount", _msa)
}

// Setmultisigamount is a paid mutator transaction binding the contract method 0xd9ab23b3.
//
// Solidity: function setmultisigamount(uint256 _msa) returns(bool)
func (_PreInteraction *PreInteractionSession) Setmultisigamount(_msa *big.Int) (*types.Transaction, error) {
	return _PreInteraction.Contract.Setmultisigamount(&_PreInteraction.TransactOpts, _msa)
}

// Setmultisigamount is a paid mutator transaction binding the contract method 0xd9ab23b3.
//
// Solidity: function setmultisigamount(uint256 _msa) returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) Setmultisigamount(_msa *big.Int) (*types.Transaction, error) {
	return _PreInteraction.Contract.Setmultisigamount(&_PreInteraction.TransactOpts, _msa)
}

// Specialadmitfunction is a paid mutator transaction binding the contract method 0x78f12169.
//
// Solidity: function specialadmitfunction(uint8 hymn) returns(bool)
func (_PreInteraction *PreInteractionTransactor) Specialadmitfunction(opts *bind.TransactOpts, hymn uint8) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "specialadmitfunction", hymn)
}

// Specialadmitfunction is a paid mutator transaction binding the contract method 0x78f12169.
//
// Solidity: function specialadmitfunction(uint8 hymn) returns(bool)
func (_PreInteraction *PreInteractionSession) Specialadmitfunction(hymn uint8) (*types.Transaction, error) {
	return _PreInteraction.Contract.Specialadmitfunction(&_PreInteraction.TransactOpts, hymn)
}

// Specialadmitfunction is a paid mutator transaction binding the contract method 0x78f12169.
//
// Solidity: function specialadmitfunction(uint8 hymn) returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) Specialadmitfunction(hymn uint8) (*types.Transaction, error) {
	return _PreInteraction.Contract.Specialadmitfunction(&_PreInteraction.TransactOpts, hymn)
}

// SubmitConsensus is a paid mutator transaction binding the contract method 0xd5c9f297.
//
// Solidity: function submitConsensus(string _btcAddress, string _btcAddressR, uint256 satsam) returns(bool)
func (_PreInteraction *PreInteractionTransactor) SubmitConsensus(opts *bind.TransactOpts, _btcAddress string, _btcAddressR string, satsam *big.Int) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "submitConsensus", _btcAddress, _btcAddressR, satsam)
}

// SubmitConsensus is a paid mutator transaction binding the contract method 0xd5c9f297.
//
// Solidity: function submitConsensus(string _btcAddress, string _btcAddressR, uint256 satsam) returns(bool)
func (_PreInteraction *PreInteractionSession) SubmitConsensus(_btcAddress string, _btcAddressR string, satsam *big.Int) (*types.Transaction, error) {
	return _PreInteraction.Contract.SubmitConsensus(&_PreInteraction.TransactOpts, _btcAddress, _btcAddressR, satsam)
}

// SubmitConsensus is a paid mutator transaction binding the contract method 0xd5c9f297.
//
// Solidity: function submitConsensus(string _btcAddress, string _btcAddressR, uint256 satsam) returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) SubmitConsensus(_btcAddress string, _btcAddressR string, satsam *big.Int) (*types.Transaction, error) {
	return _PreInteraction.Contract.SubmitConsensus(&_PreInteraction.TransactOpts, _btcAddress, _btcAddressR, satsam)
}

// Withdrawequal is a paid mutator transaction binding the contract method 0xffc3916d.
//
// Solidity: function withdrawequal() returns(bool)
func (_PreInteraction *PreInteractionTransactor) Withdrawequal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreInteraction.contract.Transact(opts, "withdrawequal")
}

// Withdrawequal is a paid mutator transaction binding the contract method 0xffc3916d.
//
// Solidity: function withdrawequal() returns(bool)
func (_PreInteraction *PreInteractionSession) Withdrawequal() (*types.Transaction, error) {
	return _PreInteraction.Contract.Withdrawequal(&_PreInteraction.TransactOpts)
}

// Withdrawequal is a paid mutator transaction binding the contract method 0xffc3916d.
//
// Solidity: function withdrawequal() returns(bool)
func (_PreInteraction *PreInteractionTransactorSession) Withdrawequal() (*types.Transaction, error) {
	return _PreInteraction.Contract.Withdrawequal(&_PreInteraction.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PreInteraction *PreInteractionTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _PreInteraction.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PreInteraction *PreInteractionSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PreInteraction.Contract.Fallback(&_PreInteraction.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PreInteraction *PreInteractionTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PreInteraction.Contract.Fallback(&_PreInteraction.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PreInteraction *PreInteractionTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreInteraction.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PreInteraction *PreInteractionSession) Receive() (*types.Transaction, error) {
	return _PreInteraction.Contract.Receive(&_PreInteraction.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PreInteraction *PreInteractionTransactorSession) Receive() (*types.Transaction, error) {
	return _PreInteraction.Contract.Receive(&_PreInteraction.TransactOpts)
}

// PreInteractionFailedIterator is returned from FilterFailed and is used to iterate over the raw logs and unpacked data for Failed events raised by the PreInteraction contract.
type PreInteractionFailedIterator struct {
	Event *PreInteractionFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreInteractionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreInteractionFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreInteractionFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreInteractionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreInteractionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreInteractionFailed represents a Failed event raised by the PreInteraction contract.
type PreInteractionFailed struct {
	Pid        *big.Int
	FsId       *big.Int
	Btcaddrstr common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFailed is a free log retrieval operation binding the contract event 0x2ec2ffd3b6f68867bfd45d0e542c4eea1fd7b9163ad9d172a3b71f50bae3df1e.
//
// Solidity: event Failed(uint256 indexed pid, uint256 indexed fs_id, string indexed btcaddrstr)
func (_PreInteraction *PreInteractionFilterer) FilterFailed(opts *bind.FilterOpts, pid []*big.Int, fs_id []*big.Int, btcaddrstr []string) (*PreInteractionFailedIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var fs_idRule []interface{}
	for _, fs_idItem := range fs_id {
		fs_idRule = append(fs_idRule, fs_idItem)
	}
	var btcaddrstrRule []interface{}
	for _, btcaddrstrItem := range btcaddrstr {
		btcaddrstrRule = append(btcaddrstrRule, btcaddrstrItem)
	}

	logs, sub, err := _PreInteraction.contract.FilterLogs(opts, "Failed", pidRule, fs_idRule, btcaddrstrRule)
	if err != nil {
		return nil, err
	}
	return &PreInteractionFailedIterator{contract: _PreInteraction.contract, event: "Failed", logs: logs, sub: sub}, nil
}

// WatchFailed is a free log subscription operation binding the contract event 0x2ec2ffd3b6f68867bfd45d0e542c4eea1fd7b9163ad9d172a3b71f50bae3df1e.
//
// Solidity: event Failed(uint256 indexed pid, uint256 indexed fs_id, string indexed btcaddrstr)
func (_PreInteraction *PreInteractionFilterer) WatchFailed(opts *bind.WatchOpts, sink chan<- *PreInteractionFailed, pid []*big.Int, fs_id []*big.Int, btcaddrstr []string) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var fs_idRule []interface{}
	for _, fs_idItem := range fs_id {
		fs_idRule = append(fs_idRule, fs_idItem)
	}
	var btcaddrstrRule []interface{}
	for _, btcaddrstrItem := range btcaddrstr {
		btcaddrstrRule = append(btcaddrstrRule, btcaddrstrItem)
	}

	logs, sub, err := _PreInteraction.contract.WatchLogs(opts, "Failed", pidRule, fs_idRule, btcaddrstrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreInteractionFailed)
				if err := _PreInteraction.contract.UnpackLog(event, "Failed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFailed is a log parse operation binding the contract event 0x2ec2ffd3b6f68867bfd45d0e542c4eea1fd7b9163ad9d172a3b71f50bae3df1e.
//
// Solidity: event Failed(uint256 indexed pid, uint256 indexed fs_id, string indexed btcaddrstr)
func (_PreInteraction *PreInteractionFilterer) ParseFailed(log types.Log) (*PreInteractionFailed, error) {
	event := new(PreInteractionFailed)
	if err := _PreInteraction.contract.UnpackLog(event, "Failed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PreInteractionPostInteractionNotificationIterator is returned from FilterPostInteractionNotification and is used to iterate over the raw logs and unpacked data for PostInteractionNotification events raised by the PreInteraction contract.
type PreInteractionPostInteractionNotificationIterator struct {
	Event *PreInteractionPostInteractionNotification // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreInteractionPostInteractionNotificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreInteractionPostInteractionNotification)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreInteractionPostInteractionNotification)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreInteractionPostInteractionNotificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreInteractionPostInteractionNotificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreInteractionPostInteractionNotification represents a PostInteractionNotification event raised by the PreInteraction contract.
type PreInteractionPostInteractionNotification struct {
	RecipAddress common.Address
	BtcAmount    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPostInteractionNotification is a free log retrieval operation binding the contract event 0x7d27ffa7a49067c28cb4dab9352e662a00381e02edd475356a7c2fc16f520258.
//
// Solidity: event PostInteractionNotification(address indexed recipAddress, uint256 indexed btcAmount)
func (_PreInteraction *PreInteractionFilterer) FilterPostInteractionNotification(opts *bind.FilterOpts, recipAddress []common.Address, btcAmount []*big.Int) (*PreInteractionPostInteractionNotificationIterator, error) {

	var recipAddressRule []interface{}
	for _, recipAddressItem := range recipAddress {
		recipAddressRule = append(recipAddressRule, recipAddressItem)
	}
	var btcAmountRule []interface{}
	for _, btcAmountItem := range btcAmount {
		btcAmountRule = append(btcAmountRule, btcAmountItem)
	}

	logs, sub, err := _PreInteraction.contract.FilterLogs(opts, "PostInteractionNotification", recipAddressRule, btcAmountRule)
	if err != nil {
		return nil, err
	}
	return &PreInteractionPostInteractionNotificationIterator{contract: _PreInteraction.contract, event: "PostInteractionNotification", logs: logs, sub: sub}, nil
}

// WatchPostInteractionNotification is a free log subscription operation binding the contract event 0x7d27ffa7a49067c28cb4dab9352e662a00381e02edd475356a7c2fc16f520258.
//
// Solidity: event PostInteractionNotification(address indexed recipAddress, uint256 indexed btcAmount)
func (_PreInteraction *PreInteractionFilterer) WatchPostInteractionNotification(opts *bind.WatchOpts, sink chan<- *PreInteractionPostInteractionNotification, recipAddress []common.Address, btcAmount []*big.Int) (event.Subscription, error) {

	var recipAddressRule []interface{}
	for _, recipAddressItem := range recipAddress {
		recipAddressRule = append(recipAddressRule, recipAddressItem)
	}
	var btcAmountRule []interface{}
	for _, btcAmountItem := range btcAmount {
		btcAmountRule = append(btcAmountRule, btcAmountItem)
	}

	logs, sub, err := _PreInteraction.contract.WatchLogs(opts, "PostInteractionNotification", recipAddressRule, btcAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreInteractionPostInteractionNotification)
				if err := _PreInteraction.contract.UnpackLog(event, "PostInteractionNotification", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePostInteractionNotification is a log parse operation binding the contract event 0x7d27ffa7a49067c28cb4dab9352e662a00381e02edd475356a7c2fc16f520258.
//
// Solidity: event PostInteractionNotification(address indexed recipAddress, uint256 indexed btcAmount)
func (_PreInteraction *PreInteractionFilterer) ParsePostInteractionNotification(log types.Log) (*PreInteractionPostInteractionNotification, error) {
	event := new(PreInteractionPostInteractionNotification)
	if err := _PreInteraction.contract.UnpackLog(event, "PostInteractionNotification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PreInteractionPreInteractionNotificationIterator is returned from FilterPreInteractionNotification and is used to iterate over the raw logs and unpacked data for PreInteractionNotification events raised by the PreInteraction contract.
type PreInteractionPreInteractionNotificationIterator struct {
	Event *PreInteractionPreInteractionNotification // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreInteractionPreInteractionNotificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreInteractionPreInteractionNotification)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreInteractionPreInteractionNotification)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreInteractionPreInteractionNotificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreInteractionPreInteractionNotificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreInteractionPreInteractionNotification represents a PreInteractionNotification event raised by the PreInteraction contract.
type PreInteractionPreInteractionNotification struct {
	MethodId  *big.Int
	BtcAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPreInteractionNotification is a free log retrieval operation binding the contract event 0xac0f5426a2d5150cff42daf93eeb82bdfd64d81f7b892886031d55266f233298.
//
// Solidity: event PreInteractionNotification(uint256 indexed methodId, uint256 indexed btcAmount)
func (_PreInteraction *PreInteractionFilterer) FilterPreInteractionNotification(opts *bind.FilterOpts, methodId []*big.Int, btcAmount []*big.Int) (*PreInteractionPreInteractionNotificationIterator, error) {

	var methodIdRule []interface{}
	for _, methodIdItem := range methodId {
		methodIdRule = append(methodIdRule, methodIdItem)
	}
	var btcAmountRule []interface{}
	for _, btcAmountItem := range btcAmount {
		btcAmountRule = append(btcAmountRule, btcAmountItem)
	}

	logs, sub, err := _PreInteraction.contract.FilterLogs(opts, "PreInteractionNotification", methodIdRule, btcAmountRule)
	if err != nil {
		return nil, err
	}
	return &PreInteractionPreInteractionNotificationIterator{contract: _PreInteraction.contract, event: "PreInteractionNotification", logs: logs, sub: sub}, nil
}

// WatchPreInteractionNotification is a free log subscription operation binding the contract event 0xac0f5426a2d5150cff42daf93eeb82bdfd64d81f7b892886031d55266f233298.
//
// Solidity: event PreInteractionNotification(uint256 indexed methodId, uint256 indexed btcAmount)
func (_PreInteraction *PreInteractionFilterer) WatchPreInteractionNotification(opts *bind.WatchOpts, sink chan<- *PreInteractionPreInteractionNotification, methodId []*big.Int, btcAmount []*big.Int) (event.Subscription, error) {

	var methodIdRule []interface{}
	for _, methodIdItem := range methodId {
		methodIdRule = append(methodIdRule, methodIdItem)
	}
	var btcAmountRule []interface{}
	for _, btcAmountItem := range btcAmount {
		btcAmountRule = append(btcAmountRule, btcAmountItem)
	}

	logs, sub, err := _PreInteraction.contract.WatchLogs(opts, "PreInteractionNotification", methodIdRule, btcAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreInteractionPreInteractionNotification)
				if err := _PreInteraction.contract.UnpackLog(event, "PreInteractionNotification", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePreInteractionNotification is a log parse operation binding the contract event 0xac0f5426a2d5150cff42daf93eeb82bdfd64d81f7b892886031d55266f233298.
//
// Solidity: event PreInteractionNotification(uint256 indexed methodId, uint256 indexed btcAmount)
func (_PreInteraction *PreInteractionFilterer) ParsePreInteractionNotification(log types.Log) (*PreInteractionPreInteractionNotification, error) {
	event := new(PreInteractionPreInteractionNotification)
	if err := _PreInteraction.contract.UnpackLog(event, "PreInteractionNotification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PreInteractionNewNodesListIterator is returned from FilterNewNodesList and is used to iterate over the raw logs and unpacked data for NewNodesList events raised by the PreInteraction contract.
type PreInteractionNewNodesListIterator struct {
	Event *PreInteractionNewNodesList // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreInteractionNewNodesListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreInteractionNewNodesList)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreInteractionNewNodesList)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreInteractionNewNodesListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreInteractionNewNodesListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreInteractionNewNodesList represents a NewNodesList event raised by the PreInteraction contract.
type PreInteractionNewNodesList struct {
	Sigtype *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewNodesList is a free log retrieval operation binding the contract event 0xdc9584bd9cf06f9f9c3d9de83b5bd718c4db744c92a0dc843c2bf17bba626dbb.
//
// Solidity: event newNodesList(uint256 indexed _sigtype)
func (_PreInteraction *PreInteractionFilterer) FilterNewNodesList(opts *bind.FilterOpts, _sigtype []*big.Int) (*PreInteractionNewNodesListIterator, error) {

	var _sigtypeRule []interface{}
	for _, _sigtypeItem := range _sigtype {
		_sigtypeRule = append(_sigtypeRule, _sigtypeItem)
	}

	logs, sub, err := _PreInteraction.contract.FilterLogs(opts, "newNodesList", _sigtypeRule)
	if err != nil {
		return nil, err
	}
	return &PreInteractionNewNodesListIterator{contract: _PreInteraction.contract, event: "newNodesList", logs: logs, sub: sub}, nil
}

// WatchNewNodesList is a free log subscription operation binding the contract event 0xdc9584bd9cf06f9f9c3d9de83b5bd718c4db744c92a0dc843c2bf17bba626dbb.
//
// Solidity: event newNodesList(uint256 indexed _sigtype)
func (_PreInteraction *PreInteractionFilterer) WatchNewNodesList(opts *bind.WatchOpts, sink chan<- *PreInteractionNewNodesList, _sigtype []*big.Int) (event.Subscription, error) {

	var _sigtypeRule []interface{}
	for _, _sigtypeItem := range _sigtype {
		_sigtypeRule = append(_sigtypeRule, _sigtypeItem)
	}

	logs, sub, err := _PreInteraction.contract.WatchLogs(opts, "newNodesList", _sigtypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreInteractionNewNodesList)
				if err := _PreInteraction.contract.UnpackLog(event, "newNodesList", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewNodesList is a log parse operation binding the contract event 0xdc9584bd9cf06f9f9c3d9de83b5bd718c4db744c92a0dc843c2bf17bba626dbb.
//
// Solidity: event newNodesList(uint256 indexed _sigtype)
func (_PreInteraction *PreInteractionFilterer) ParseNewNodesList(log types.Log) (*PreInteractionNewNodesList, error) {
	event := new(PreInteractionNewNodesList)
	if err := _PreInteraction.contract.UnpackLog(event, "newNodesList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PreInteractionPopulateMSLIterator is returned from FilterPopulateMSL and is used to iterate over the raw logs and unpacked data for PopulateMSL events raised by the PreInteraction contract.
type PreInteractionPopulateMSLIterator struct {
	Event *PreInteractionPopulateMSL // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreInteractionPopulateMSLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreInteractionPopulateMSL)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreInteractionPopulateMSL)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreInteractionPopulateMSLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreInteractionPopulateMSLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreInteractionPopulateMSL represents a PopulateMSL event raised by the PreInteraction contract.
type PreInteractionPopulateMSL struct {
	Sigtype *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPopulateMSL is a free log retrieval operation binding the contract event 0x626551c9eae307eb3cdda0b42dd957f9b3b1f19f26b295a63773fb2662f66e97.
//
// Solidity: event populateMSL(uint256 indexed _sigtype)
func (_PreInteraction *PreInteractionFilterer) FilterPopulateMSL(opts *bind.FilterOpts, _sigtype []*big.Int) (*PreInteractionPopulateMSLIterator, error) {

	var _sigtypeRule []interface{}
	for _, _sigtypeItem := range _sigtype {
		_sigtypeRule = append(_sigtypeRule, _sigtypeItem)
	}

	logs, sub, err := _PreInteraction.contract.FilterLogs(opts, "populateMSL", _sigtypeRule)
	if err != nil {
		return nil, err
	}
	return &PreInteractionPopulateMSLIterator{contract: _PreInteraction.contract, event: "populateMSL", logs: logs, sub: sub}, nil
}

// WatchPopulateMSL is a free log subscription operation binding the contract event 0x626551c9eae307eb3cdda0b42dd957f9b3b1f19f26b295a63773fb2662f66e97.
//
// Solidity: event populateMSL(uint256 indexed _sigtype)
func (_PreInteraction *PreInteractionFilterer) WatchPopulateMSL(opts *bind.WatchOpts, sink chan<- *PreInteractionPopulateMSL, _sigtype []*big.Int) (event.Subscription, error) {

	var _sigtypeRule []interface{}
	for _, _sigtypeItem := range _sigtype {
		_sigtypeRule = append(_sigtypeRule, _sigtypeItem)
	}

	logs, sub, err := _PreInteraction.contract.WatchLogs(opts, "populateMSL", _sigtypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreInteractionPopulateMSL)
				if err := _PreInteraction.contract.UnpackLog(event, "populateMSL", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePopulateMSL is a log parse operation binding the contract event 0x626551c9eae307eb3cdda0b42dd957f9b3b1f19f26b295a63773fb2662f66e97.
//
// Solidity: event populateMSL(uint256 indexed _sigtype)
func (_PreInteraction *PreInteractionFilterer) ParsePopulateMSL(log types.Log) (*PreInteractionPopulateMSL, error) {
	event := new(PreInteractionPopulateMSL)
	if err := _PreInteraction.contract.UnpackLog(event, "populateMSL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
