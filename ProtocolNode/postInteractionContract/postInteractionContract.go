// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package postInteractionContract

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

// PostInteractionContractexitmultisigObj is an auto generated low-level Go binding around an user-defined struct.
type PostInteractionContractexitmultisigObj struct {
	RecipientBtcAddr string
	SatsAmount       *big.Int
	Procd            bool
}

// PostInteractionContractMetaData contains all meta data concerning the PostInteractionContract contract.
var PostInteractionContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_msbtc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delegc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wrappingcontract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_obt\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"btcAmount\",\"type\":\"uint256\"}],\"name\":\"InteractionNotification\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_np\",\"type\":\"address\"}],\"name\":\"addPools\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recip\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sats\",\"type\":\"uint256\"}],\"name\":\"addToList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newadm\",\"type\":\"address\"}],\"name\":\"changeadmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_obtc\",\"type\":\"address\"}],\"name\":\"changeobtcc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nwc\",\"type\":\"address\"}],\"name\":\"changewc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txid\",\"type\":\"uint256\"}],\"name\":\"gettxs\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientBtcAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"satsAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"procd\",\"type\":\"bool\"}],\"internalType\":\"structPostInteractionContract.exitmultisigObj\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txid\",\"type\":\"uint256\"}],\"name\":\"marktxscomplete\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_btcAddress\",\"type\":\"string\"}],\"name\":\"populateShaTable\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newdeleg\",\"type\":\"address\"}],\"name\":\"setde\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txid\",\"type\":\"uint256\"}],\"name\":\"specadmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PostInteractionContractABI is the input ABI used to generate the binding from.
// Deprecated: Use PostInteractionContractMetaData.ABI instead.
var PostInteractionContractABI = PostInteractionContractMetaData.ABI

// PostInteractionContract is an auto generated Go binding around an Ethereum contract.
type PostInteractionContract struct {
	PostInteractionContractCaller     // Read-only binding to the contract
	PostInteractionContractTransactor // Write-only binding to the contract
	PostInteractionContractFilterer   // Log filterer for contract events
}

// PostInteractionContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PostInteractionContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostInteractionContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PostInteractionContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostInteractionContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PostInteractionContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostInteractionContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PostInteractionContractSession struct {
	Contract     *PostInteractionContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PostInteractionContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PostInteractionContractCallerSession struct {
	Contract *PostInteractionContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// PostInteractionContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PostInteractionContractTransactorSession struct {
	Contract     *PostInteractionContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// PostInteractionContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PostInteractionContractRaw struct {
	Contract *PostInteractionContract // Generic contract binding to access the raw methods on
}

// PostInteractionContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PostInteractionContractCallerRaw struct {
	Contract *PostInteractionContractCaller // Generic read-only contract binding to access the raw methods on
}

// PostInteractionContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PostInteractionContractTransactorRaw struct {
	Contract *PostInteractionContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPostInteractionContract creates a new instance of PostInteractionContract, bound to a specific deployed contract.
func NewPostInteractionContract(address common.Address, backend bind.ContractBackend) (*PostInteractionContract, error) {
	contract, err := bindPostInteractionContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PostInteractionContract{PostInteractionContractCaller: PostInteractionContractCaller{contract: contract}, PostInteractionContractTransactor: PostInteractionContractTransactor{contract: contract}, PostInteractionContractFilterer: PostInteractionContractFilterer{contract: contract}}, nil
}

// NewPostInteractionContractCaller creates a new read-only instance of PostInteractionContract, bound to a specific deployed contract.
func NewPostInteractionContractCaller(address common.Address, caller bind.ContractCaller) (*PostInteractionContractCaller, error) {
	contract, err := bindPostInteractionContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PostInteractionContractCaller{contract: contract}, nil
}

// NewPostInteractionContractTransactor creates a new write-only instance of PostInteractionContract, bound to a specific deployed contract.
func NewPostInteractionContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PostInteractionContractTransactor, error) {
	contract, err := bindPostInteractionContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PostInteractionContractTransactor{contract: contract}, nil
}

// NewPostInteractionContractFilterer creates a new log filterer instance of PostInteractionContract, bound to a specific deployed contract.
func NewPostInteractionContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PostInteractionContractFilterer, error) {
	contract, err := bindPostInteractionContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PostInteractionContractFilterer{contract: contract}, nil
}

// bindPostInteractionContract binds a generic wrapper to an already deployed contract.
func bindPostInteractionContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PostInteractionContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PostInteractionContract *PostInteractionContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PostInteractionContract.Contract.PostInteractionContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PostInteractionContract *PostInteractionContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.PostInteractionContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PostInteractionContract *PostInteractionContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.PostInteractionContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PostInteractionContract *PostInteractionContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PostInteractionContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PostInteractionContract *PostInteractionContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PostInteractionContract *PostInteractionContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.contract.Transact(opts, method, params...)
}

// Gettxs is a free data retrieval call binding the contract method 0x7cfa6453.
//
// Solidity: function gettxs(uint256 _txid) view returns((string,uint256,bool))
func (_PostInteractionContract *PostInteractionContractCaller) Gettxs(opts *bind.CallOpts, _txid *big.Int) (PostInteractionContractexitmultisigObj, error) {
	var out []interface{}
	err := _PostInteractionContract.contract.Call(opts, &out, "gettxs", _txid)

	if err != nil {
		return *new(PostInteractionContractexitmultisigObj), err
	}

	out0 := *abi.ConvertType(out[0], new(PostInteractionContractexitmultisigObj)).(*PostInteractionContractexitmultisigObj)

	return out0, err

}

// Gettxs is a free data retrieval call binding the contract method 0x7cfa6453.
//
// Solidity: function gettxs(uint256 _txid) view returns((string,uint256,bool))
func (_PostInteractionContract *PostInteractionContractSession) Gettxs(_txid *big.Int) (PostInteractionContractexitmultisigObj, error) {
	return _PostInteractionContract.Contract.Gettxs(&_PostInteractionContract.CallOpts, _txid)
}

// Gettxs is a free data retrieval call binding the contract method 0x7cfa6453.
//
// Solidity: function gettxs(uint256 _txid) view returns((string,uint256,bool))
func (_PostInteractionContract *PostInteractionContractCallerSession) Gettxs(_txid *big.Int) (PostInteractionContractexitmultisigObj, error) {
	return _PostInteractionContract.Contract.Gettxs(&_PostInteractionContract.CallOpts, _txid)
}

// AddPools is a paid mutator transaction binding the contract method 0x96e6bc29.
//
// Solidity: function addPools(address _np) returns(bool)
func (_PostInteractionContract *PostInteractionContractTransactor) AddPools(opts *bind.TransactOpts, _np common.Address) (*types.Transaction, error) {
	return _PostInteractionContract.contract.Transact(opts, "addPools", _np)
}

// AddPools is a paid mutator transaction binding the contract method 0x96e6bc29.
//
// Solidity: function addPools(address _np) returns(bool)
func (_PostInteractionContract *PostInteractionContractSession) AddPools(_np common.Address) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.AddPools(&_PostInteractionContract.TransactOpts, _np)
}

// AddPools is a paid mutator transaction binding the contract method 0x96e6bc29.
//
// Solidity: function addPools(address _np) returns(bool)
func (_PostInteractionContract *PostInteractionContractTransactorSession) AddPools(_np common.Address) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.AddPools(&_PostInteractionContract.TransactOpts, _np)
}

// AddToList is a paid mutator transaction binding the contract method 0xda2dca10.
//
// Solidity: function addToList(address _recip, uint256 _sats) returns(bool)
func (_PostInteractionContract *PostInteractionContractTransactor) AddToList(opts *bind.TransactOpts, _recip common.Address, _sats *big.Int) (*types.Transaction, error) {
	return _PostInteractionContract.contract.Transact(opts, "addToList", _recip, _sats)
}

// AddToList is a paid mutator transaction binding the contract method 0xda2dca10.
//
// Solidity: function addToList(address _recip, uint256 _sats) returns(bool)
func (_PostInteractionContract *PostInteractionContractSession) AddToList(_recip common.Address, _sats *big.Int) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.AddToList(&_PostInteractionContract.TransactOpts, _recip, _sats)
}

// AddToList is a paid mutator transaction binding the contract method 0xda2dca10.
//
// Solidity: function addToList(address _recip, uint256 _sats) returns(bool)
func (_PostInteractionContract *PostInteractionContractTransactorSession) AddToList(_recip common.Address, _sats *big.Int) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.AddToList(&_PostInteractionContract.TransactOpts, _recip, _sats)
}

// Changeadmin is a paid mutator transaction binding the contract method 0x26ea7ab8.
//
// Solidity: function changeadmin(address newadm) returns(bool)
func (_PostInteractionContract *PostInteractionContractTransactor) Changeadmin(opts *bind.TransactOpts, newadm common.Address) (*types.Transaction, error) {
	return _PostInteractionContract.contract.Transact(opts, "changeadmin", newadm)
}

// Changeadmin is a paid mutator transaction binding the contract method 0x26ea7ab8.
//
// Solidity: function changeadmin(address newadm) returns(bool)
func (_PostInteractionContract *PostInteractionContractSession) Changeadmin(newadm common.Address) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.Changeadmin(&_PostInteractionContract.TransactOpts, newadm)
}

// Changeadmin is a paid mutator transaction binding the contract method 0x26ea7ab8.
//
// Solidity: function changeadmin(address newadm) returns(bool)
func (_PostInteractionContract *PostInteractionContractTransactorSession) Changeadmin(newadm common.Address) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.Changeadmin(&_PostInteractionContract.TransactOpts, newadm)
}

// Changeobtcc is a paid mutator transaction binding the contract method 0xf896e32c.
//
// Solidity: function changeobtcc(address _obtc) returns(bool)
func (_PostInteractionContract *PostInteractionContractTransactor) Changeobtcc(opts *bind.TransactOpts, _obtc common.Address) (*types.Transaction, error) {
	return _PostInteractionContract.contract.Transact(opts, "changeobtcc", _obtc)
}

// Changeobtcc is a paid mutator transaction binding the contract method 0xf896e32c.
//
// Solidity: function changeobtcc(address _obtc) returns(bool)
func (_PostInteractionContract *PostInteractionContractSession) Changeobtcc(_obtc common.Address) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.Changeobtcc(&_PostInteractionContract.TransactOpts, _obtc)
}

// Changeobtcc is a paid mutator transaction binding the contract method 0xf896e32c.
//
// Solidity: function changeobtcc(address _obtc) returns(bool)
func (_PostInteractionContract *PostInteractionContractTransactorSession) Changeobtcc(_obtc common.Address) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.Changeobtcc(&_PostInteractionContract.TransactOpts, _obtc)
}

// Changewc is a paid mutator transaction binding the contract method 0x0a9b6f4f.
//
// Solidity: function changewc(address _nwc) returns(bool)
func (_PostInteractionContract *PostInteractionContractTransactor) Changewc(opts *bind.TransactOpts, _nwc common.Address) (*types.Transaction, error) {
	return _PostInteractionContract.contract.Transact(opts, "changewc", _nwc)
}

// Changewc is a paid mutator transaction binding the contract method 0x0a9b6f4f.
//
// Solidity: function changewc(address _nwc) returns(bool)
func (_PostInteractionContract *PostInteractionContractSession) Changewc(_nwc common.Address) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.Changewc(&_PostInteractionContract.TransactOpts, _nwc)
}

// Changewc is a paid mutator transaction binding the contract method 0x0a9b6f4f.
//
// Solidity: function changewc(address _nwc) returns(bool)
func (_PostInteractionContract *PostInteractionContractTransactorSession) Changewc(_nwc common.Address) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.Changewc(&_PostInteractionContract.TransactOpts, _nwc)
}

// Marktxscomplete is a paid mutator transaction binding the contract method 0xf33c07ae.
//
// Solidity: function marktxscomplete(uint256 _txid) returns(bool)
func (_PostInteractionContract *PostInteractionContractTransactor) Marktxscomplete(opts *bind.TransactOpts, _txid *big.Int) (*types.Transaction, error) {
	return _PostInteractionContract.contract.Transact(opts, "marktxscomplete", _txid)
}

// Marktxscomplete is a paid mutator transaction binding the contract method 0xf33c07ae.
//
// Solidity: function marktxscomplete(uint256 _txid) returns(bool)
func (_PostInteractionContract *PostInteractionContractSession) Marktxscomplete(_txid *big.Int) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.Marktxscomplete(&_PostInteractionContract.TransactOpts, _txid)
}

// Marktxscomplete is a paid mutator transaction binding the contract method 0xf33c07ae.
//
// Solidity: function marktxscomplete(uint256 _txid) returns(bool)
func (_PostInteractionContract *PostInteractionContractTransactorSession) Marktxscomplete(_txid *big.Int) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.Marktxscomplete(&_PostInteractionContract.TransactOpts, _txid)
}

// PopulateShaTable is a paid mutator transaction binding the contract method 0x842ace9c.
//
// Solidity: function populateShaTable(string _btcAddress) returns(address)
func (_PostInteractionContract *PostInteractionContractTransactor) PopulateShaTable(opts *bind.TransactOpts, _btcAddress string) (*types.Transaction, error) {
	return _PostInteractionContract.contract.Transact(opts, "populateShaTable", _btcAddress)
}

// PopulateShaTable is a paid mutator transaction binding the contract method 0x842ace9c.
//
// Solidity: function populateShaTable(string _btcAddress) returns(address)
func (_PostInteractionContract *PostInteractionContractSession) PopulateShaTable(_btcAddress string) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.PopulateShaTable(&_PostInteractionContract.TransactOpts, _btcAddress)
}

// PopulateShaTable is a paid mutator transaction binding the contract method 0x842ace9c.
//
// Solidity: function populateShaTable(string _btcAddress) returns(address)
func (_PostInteractionContract *PostInteractionContractTransactorSession) PopulateShaTable(_btcAddress string) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.PopulateShaTable(&_PostInteractionContract.TransactOpts, _btcAddress)
}

// Setde is a paid mutator transaction binding the contract method 0xb07c410a.
//
// Solidity: function setde(address _newdeleg) returns(bool)
func (_PostInteractionContract *PostInteractionContractTransactor) Setde(opts *bind.TransactOpts, _newdeleg common.Address) (*types.Transaction, error) {
	return _PostInteractionContract.contract.Transact(opts, "setde", _newdeleg)
}

// Setde is a paid mutator transaction binding the contract method 0xb07c410a.
//
// Solidity: function setde(address _newdeleg) returns(bool)
func (_PostInteractionContract *PostInteractionContractSession) Setde(_newdeleg common.Address) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.Setde(&_PostInteractionContract.TransactOpts, _newdeleg)
}

// Setde is a paid mutator transaction binding the contract method 0xb07c410a.
//
// Solidity: function setde(address _newdeleg) returns(bool)
func (_PostInteractionContract *PostInteractionContractTransactorSession) Setde(_newdeleg common.Address) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.Setde(&_PostInteractionContract.TransactOpts, _newdeleg)
}

// Specadmin is a paid mutator transaction binding the contract method 0x4c115041.
//
// Solidity: function specadmin(uint256 _txid) returns(bool)
func (_PostInteractionContract *PostInteractionContractTransactor) Specadmin(opts *bind.TransactOpts, _txid *big.Int) (*types.Transaction, error) {
	return _PostInteractionContract.contract.Transact(opts, "specadmin", _txid)
}

// Specadmin is a paid mutator transaction binding the contract method 0x4c115041.
//
// Solidity: function specadmin(uint256 _txid) returns(bool)
func (_PostInteractionContract *PostInteractionContractSession) Specadmin(_txid *big.Int) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.Specadmin(&_PostInteractionContract.TransactOpts, _txid)
}

// Specadmin is a paid mutator transaction binding the contract method 0x4c115041.
//
// Solidity: function specadmin(uint256 _txid) returns(bool)
func (_PostInteractionContract *PostInteractionContractTransactorSession) Specadmin(_txid *big.Int) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.Specadmin(&_PostInteractionContract.TransactOpts, _txid)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PostInteractionContract *PostInteractionContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _PostInteractionContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PostInteractionContract *PostInteractionContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.Fallback(&_PostInteractionContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PostInteractionContract *PostInteractionContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PostInteractionContract.Contract.Fallback(&_PostInteractionContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PostInteractionContract *PostInteractionContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PostInteractionContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PostInteractionContract *PostInteractionContractSession) Receive() (*types.Transaction, error) {
	return _PostInteractionContract.Contract.Receive(&_PostInteractionContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PostInteractionContract *PostInteractionContractTransactorSession) Receive() (*types.Transaction, error) {
	return _PostInteractionContract.Contract.Receive(&_PostInteractionContract.TransactOpts)
}

// PostInteractionContractInteractionNotificationIterator is returned from FilterInteractionNotification and is used to iterate over the raw logs and unpacked data for InteractionNotification events raised by the PostInteractionContract contract.
type PostInteractionContractInteractionNotificationIterator struct {
	Event *PostInteractionContractInteractionNotification // Event containing the contract specifics and raw log

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
func (it *PostInteractionContractInteractionNotificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostInteractionContractInteractionNotification)
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
		it.Event = new(PostInteractionContractInteractionNotification)
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
func (it *PostInteractionContractInteractionNotificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostInteractionContractInteractionNotificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostInteractionContractInteractionNotification represents a InteractionNotification event raised by the PostInteractionContract contract.
type PostInteractionContractInteractionNotification struct {
	RecipAddress common.Address
	BtcAmount    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInteractionNotification is a free log retrieval operation binding the contract event 0xaccabdb71f710dadad25f2cbb7fe4902dadc158c64363a7f4451bef865833edb.
//
// Solidity: event InteractionNotification(address indexed recipAddress, uint256 indexed btcAmount)
func (_PostInteractionContract *PostInteractionContractFilterer) FilterInteractionNotification(opts *bind.FilterOpts, recipAddress []common.Address, btcAmount []*big.Int) (*PostInteractionContractInteractionNotificationIterator, error) {

	var recipAddressRule []interface{}
	for _, recipAddressItem := range recipAddress {
		recipAddressRule = append(recipAddressRule, recipAddressItem)
	}
	var btcAmountRule []interface{}
	for _, btcAmountItem := range btcAmount {
		btcAmountRule = append(btcAmountRule, btcAmountItem)
	}

	logs, sub, err := _PostInteractionContract.contract.FilterLogs(opts, "InteractionNotification", recipAddressRule, btcAmountRule)
	if err != nil {
		return nil, err
	}
	return &PostInteractionContractInteractionNotificationIterator{contract: _PostInteractionContract.contract, event: "InteractionNotification", logs: logs, sub: sub}, nil
}

// WatchInteractionNotification is a free log subscription operation binding the contract event 0xaccabdb71f710dadad25f2cbb7fe4902dadc158c64363a7f4451bef865833edb.
//
// Solidity: event InteractionNotification(address indexed recipAddress, uint256 indexed btcAmount)
func (_PostInteractionContract *PostInteractionContractFilterer) WatchInteractionNotification(opts *bind.WatchOpts, sink chan<- *PostInteractionContractInteractionNotification, recipAddress []common.Address, btcAmount []*big.Int) (event.Subscription, error) {

	var recipAddressRule []interface{}
	for _, recipAddressItem := range recipAddress {
		recipAddressRule = append(recipAddressRule, recipAddressItem)
	}
	var btcAmountRule []interface{}
	for _, btcAmountItem := range btcAmount {
		btcAmountRule = append(btcAmountRule, btcAmountItem)
	}

	logs, sub, err := _PostInteractionContract.contract.WatchLogs(opts, "InteractionNotification", recipAddressRule, btcAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostInteractionContractInteractionNotification)
				if err := _PostInteractionContract.contract.UnpackLog(event, "InteractionNotification", log); err != nil {
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

// ParseInteractionNotification is a log parse operation binding the contract event 0xaccabdb71f710dadad25f2cbb7fe4902dadc158c64363a7f4451bef865833edb.
//
// Solidity: event InteractionNotification(address indexed recipAddress, uint256 indexed btcAmount)
func (_PostInteractionContract *PostInteractionContractFilterer) ParseInteractionNotification(log types.Log) (*PostInteractionContractInteractionNotification, error) {
	event := new(PostInteractionContractInteractionNotification)
	if err := _PostInteractionContract.contract.UnpackLog(event, "InteractionNotification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
