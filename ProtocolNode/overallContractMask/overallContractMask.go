// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package overallContractMask

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

// OverallContractMasktxobj is an auto generated low-level Go binding around an user-defined struct.
type OverallContractMasktxobj struct {
	MethodId             uint8
	InternalStartChainid uint8
	InternalEndChainid   uint8
	PairId               uint64
	FinalchainWallet     common.Address
	SecondpairId         uint64
	FirstchainAsset      common.Address
	FinalchainAsset      common.Address
	Quadrillionratio     uint64
	Quadrilliontipratio  uint64
	Rtxnum               *big.Int
	AltFee               bool
}

// OverallContractMaskMetaData contains all meta data concerning the OverallContractMask contract.
var OverallContractMaskMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"reqchainid\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"rtx_number\",\"type\":\"uint128\"}],\"name\":\"getTxByRTxNumber\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"method_id\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"internal_start_chainid\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"internal_end_chainid\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"pair_id\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"finalchain_wallet\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"secondpair_id\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"firstchain_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"finalchain_asset\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"quadrillionratio\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"quadrilliontipratio\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"rtxnum\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"alt_fee\",\"type\":\"bool\"}],\"internalType\":\"structOverallContractMask.txobj\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_icid\",\"type\":\"uint8\"}],\"name\":\"oraclePing\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// OverallContractMaskABI is the input ABI used to generate the binding from.
// Deprecated: Use OverallContractMaskMetaData.ABI instead.
var OverallContractMaskABI = OverallContractMaskMetaData.ABI

// OverallContractMask is an auto generated Go binding around an Ethereum contract.
type OverallContractMask struct {
	OverallContractMaskCaller     // Read-only binding to the contract
	OverallContractMaskTransactor // Write-only binding to the contract
	OverallContractMaskFilterer   // Log filterer for contract events
}

// OverallContractMaskCaller is an auto generated read-only Go binding around an Ethereum contract.
type OverallContractMaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OverallContractMaskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OverallContractMaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OverallContractMaskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OverallContractMaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OverallContractMaskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OverallContractMaskSession struct {
	Contract     *OverallContractMask // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OverallContractMaskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OverallContractMaskCallerSession struct {
	Contract *OverallContractMaskCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OverallContractMaskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OverallContractMaskTransactorSession struct {
	Contract     *OverallContractMaskTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OverallContractMaskRaw is an auto generated low-level Go binding around an Ethereum contract.
type OverallContractMaskRaw struct {
	Contract *OverallContractMask // Generic contract binding to access the raw methods on
}

// OverallContractMaskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OverallContractMaskCallerRaw struct {
	Contract *OverallContractMaskCaller // Generic read-only contract binding to access the raw methods on
}

// OverallContractMaskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OverallContractMaskTransactorRaw struct {
	Contract *OverallContractMaskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOverallContractMask creates a new instance of OverallContractMask, bound to a specific deployed contract.
func NewOverallContractMask(address common.Address, backend bind.ContractBackend) (*OverallContractMask, error) {
	contract, err := bindOverallContractMask(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OverallContractMask{OverallContractMaskCaller: OverallContractMaskCaller{contract: contract}, OverallContractMaskTransactor: OverallContractMaskTransactor{contract: contract}, OverallContractMaskFilterer: OverallContractMaskFilterer{contract: contract}}, nil
}

// NewOverallContractMaskCaller creates a new read-only instance of OverallContractMask, bound to a specific deployed contract.
func NewOverallContractMaskCaller(address common.Address, caller bind.ContractCaller) (*OverallContractMaskCaller, error) {
	contract, err := bindOverallContractMask(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OverallContractMaskCaller{contract: contract}, nil
}

// NewOverallContractMaskTransactor creates a new write-only instance of OverallContractMask, bound to a specific deployed contract.
func NewOverallContractMaskTransactor(address common.Address, transactor bind.ContractTransactor) (*OverallContractMaskTransactor, error) {
	contract, err := bindOverallContractMask(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OverallContractMaskTransactor{contract: contract}, nil
}

// NewOverallContractMaskFilterer creates a new log filterer instance of OverallContractMask, bound to a specific deployed contract.
func NewOverallContractMaskFilterer(address common.Address, filterer bind.ContractFilterer) (*OverallContractMaskFilterer, error) {
	contract, err := bindOverallContractMask(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OverallContractMaskFilterer{contract: contract}, nil
}

// bindOverallContractMask binds a generic wrapper to an already deployed contract.
func bindOverallContractMask(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OverallContractMaskABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OverallContractMask *OverallContractMaskRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OverallContractMask.Contract.OverallContractMaskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OverallContractMask *OverallContractMaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OverallContractMask.Contract.OverallContractMaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OverallContractMask *OverallContractMaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OverallContractMask.Contract.OverallContractMaskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OverallContractMask *OverallContractMaskCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OverallContractMask.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OverallContractMask *OverallContractMaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OverallContractMask.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OverallContractMask *OverallContractMaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OverallContractMask.Contract.contract.Transact(opts, method, params...)
}

// GetTxByRTxNumber is a free data retrieval call binding the contract method 0x841573f9.
//
// Solidity: function getTxByRTxNumber(uint8 reqchainid, uint128 rtx_number) view returns((uint8,uint8,uint8,uint64,address,uint64,address,address,uint64,uint64,uint128,bool))
func (_OverallContractMask *OverallContractMaskCaller) GetTxByRTxNumber(opts *bind.CallOpts, reqchainid uint8, rtx_number *big.Int) (OverallContractMasktxobj, error) {
	var out []interface{}
	err := _OverallContractMask.contract.Call(opts, &out, "getTxByRTxNumber", reqchainid, rtx_number)

	if err != nil {
		return *new(OverallContractMasktxobj), err
	}

	out0 := *abi.ConvertType(out[0], new(OverallContractMasktxobj)).(*OverallContractMasktxobj)

	return out0, err

}

// GetTxByRTxNumber is a free data retrieval call binding the contract method 0x841573f9.
//
// Solidity: function getTxByRTxNumber(uint8 reqchainid, uint128 rtx_number) view returns((uint8,uint8,uint8,uint64,address,uint64,address,address,uint64,uint64,uint128,bool))
func (_OverallContractMask *OverallContractMaskSession) GetTxByRTxNumber(reqchainid uint8, rtx_number *big.Int) (OverallContractMasktxobj, error) {
	return _OverallContractMask.Contract.GetTxByRTxNumber(&_OverallContractMask.CallOpts, reqchainid, rtx_number)
}

// GetTxByRTxNumber is a free data retrieval call binding the contract method 0x841573f9.
//
// Solidity: function getTxByRTxNumber(uint8 reqchainid, uint128 rtx_number) view returns((uint8,uint8,uint8,uint64,address,uint64,address,address,uint64,uint64,uint128,bool))
func (_OverallContractMask *OverallContractMaskCallerSession) GetTxByRTxNumber(reqchainid uint8, rtx_number *big.Int) (OverallContractMasktxobj, error) {
	return _OverallContractMask.Contract.GetTxByRTxNumber(&_OverallContractMask.CallOpts, reqchainid, rtx_number)
}

// OraclePing is a paid mutator transaction binding the contract method 0xd10d7f49.
//
// Solidity: function oraclePing(uint8 _icid) returns(uint128, bool)
func (_OverallContractMask *OverallContractMaskTransactor) OraclePing(opts *bind.TransactOpts, _icid uint8) (*types.Transaction, error) {
	return _OverallContractMask.contract.Transact(opts, "oraclePing", _icid)
}

// OraclePing is a paid mutator transaction binding the contract method 0xd10d7f49.
//
// Solidity: function oraclePing(uint8 _icid) returns(uint128, bool)
func (_OverallContractMask *OverallContractMaskSession) OraclePing(_icid uint8) (*types.Transaction, error) {
	return _OverallContractMask.Contract.OraclePing(&_OverallContractMask.TransactOpts, _icid)
}

// OraclePing is a paid mutator transaction binding the contract method 0xd10d7f49.
//
// Solidity: function oraclePing(uint8 _icid) returns(uint128, bool)
func (_OverallContractMask *OverallContractMaskTransactorSession) OraclePing(_icid uint8) (*types.Transaction, error) {
	return _OverallContractMask.Contract.OraclePing(&_OverallContractMask.TransactOpts, _icid)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_OverallContractMask *OverallContractMaskTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _OverallContractMask.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_OverallContractMask *OverallContractMaskSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _OverallContractMask.Contract.Fallback(&_OverallContractMask.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_OverallContractMask *OverallContractMaskTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _OverallContractMask.Contract.Fallback(&_OverallContractMask.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OverallContractMask *OverallContractMaskTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OverallContractMask.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OverallContractMask *OverallContractMaskSession) Receive() (*types.Transaction, error) {
	return _OverallContractMask.Contract.Receive(&_OverallContractMask.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OverallContractMask *OverallContractMaskTransactorSession) Receive() (*types.Transaction, error) {
	return _OverallContractMask.Contract.Receive(&_OverallContractMask.TransactOpts)
}
