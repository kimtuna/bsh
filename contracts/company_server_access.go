// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"os"
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
	_ = abi.ConvertType
)

// CompanyServerAccessMetaData contains all meta data concerning the CompanyServerAccess contract.
var CompanyServerAccessMetaData = &bind.MetaData{
	ABI: func() string {
		abiBytes, err := os.ReadFile("abi.json")
		if err != nil {
			panic(err)
		}
		return string(abiBytes)
	}(),
}

// CompanyServerAccessABI is the input ABI used to generate the binding from.
// Deprecated: Use CompanyServerAccessMetaData.ABI instead.
var CompanyServerAccessABI = CompanyServerAccessMetaData.ABI

// CompanyServerAccess is an auto generated Go binding around an Ethereum contract.
type CompanyServerAccess struct {
	CompanyServerAccessCaller     // Read-only binding to the contract
	CompanyServerAccessTransactor // Write-only binding to the contract
	CompanyServerAccessFilterer   // Log filterer for contract events
}

// CompanyServerAccessCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompanyServerAccessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompanyServerAccessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompanyServerAccessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompanyServerAccessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompanyServerAccessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompanyServerAccessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompanyServerAccessSession struct {
	Contract     *CompanyServerAccess // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CompanyServerAccessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompanyServerAccessCallerSession struct {
	Contract *CompanyServerAccessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// CompanyServerAccessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompanyServerAccessTransactorSession struct {
	Contract     *CompanyServerAccessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// CompanyServerAccessRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompanyServerAccessRaw struct {
	Contract *CompanyServerAccess // Generic contract binding to access the raw methods on
}

// CompanyServerAccessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompanyServerAccessCallerRaw struct {
	Contract *CompanyServerAccessCaller // Generic read-only contract binding to access the raw methods on
}

// CompanyServerAccessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompanyServerAccessTransactorRaw struct {
	Contract *CompanyServerAccessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompanyServerAccess creates a new instance of CompanyServerAccess, bound to a specific deployed contract.
func NewCompanyServerAccess(address common.Address, backend bind.ContractBackend) (*CompanyServerAccess, error) {
	contract, err := bindCompanyServerAccess(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompanyServerAccess{CompanyServerAccessCaller: CompanyServerAccessCaller{contract: contract}, CompanyServerAccessTransactor: CompanyServerAccessTransactor{contract: contract}, CompanyServerAccessFilterer: CompanyServerAccessFilterer{contract: contract}}, nil
}

// NewCompanyServerAccessCaller creates a new read-only instance of CompanyServerAccess, bound to a specific deployed contract.
func NewCompanyServerAccessCaller(address common.Address, caller bind.ContractCaller) (*CompanyServerAccessCaller, error) {
	contract, err := bindCompanyServerAccess(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompanyServerAccessCaller{contract: contract}, nil
}

// NewCompanyServerAccessTransactor creates a new write-only instance of CompanyServerAccess, bound to a specific deployed contract.
func NewCompanyServerAccessTransactor(address common.Address, transactor bind.ContractTransactor) (*CompanyServerAccessTransactor, error) {
	contract, err := bindCompanyServerAccess(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompanyServerAccessTransactor{contract: contract}, nil
}

// NewCompanyServerAccessFilterer creates a new log filterer instance of CompanyServerAccess, bound to a specific deployed contract.
func NewCompanyServerAccessFilterer(address common.Address, filterer bind.ContractFilterer) (*CompanyServerAccessFilterer, error) {
	contract, err := bindCompanyServerAccess(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompanyServerAccessFilterer{contract: contract}, nil
}

// bindCompanyServerAccess binds a generic wrapper to an already deployed contract.
func bindCompanyServerAccess(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CompanyServerAccessMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompanyServerAccess *CompanyServerAccessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompanyServerAccess.Contract.CompanyServerAccessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompanyServerAccess *CompanyServerAccessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.CompanyServerAccessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompanyServerAccess *CompanyServerAccessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.CompanyServerAccessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompanyServerAccess *CompanyServerAccessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompanyServerAccess.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompanyServerAccess *CompanyServerAccessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompanyServerAccess *CompanyServerAccessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.contract.Transact(opts, method, params...)
}

// CheckSubscription is a free data retrieval call binding the contract method 0x4f62647e.
//
// Solidity: function checkSubscription(address _company) view returns(bool)
func (_CompanyServerAccess *CompanyServerAccessCaller) CheckSubscription(opts *bind.CallOpts, _company common.Address) (bool, error) {
	var out []interface{}
	err := _CompanyServerAccess.contract.Call(opts, &out, "checkSubscription", _company)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckSubscription is a free data retrieval call binding the contract method 0x4f62647e.
//
// Solidity: function checkSubscription(address _company) view returns(bool)
func (_CompanyServerAccess *CompanyServerAccessSession) CheckSubscription(_company common.Address) (bool, error) {
	return _CompanyServerAccess.Contract.CheckSubscription(&_CompanyServerAccess.CallOpts, _company)
}

// CheckSubscription is a free data retrieval call binding the contract method 0x4f62647e.
//
// Solidity: function checkSubscription(address _company) view returns(bool)
func (_CompanyServerAccess *CompanyServerAccessCallerSession) CheckSubscription(_company common.Address) (bool, error) {
	return _CompanyServerAccess.Contract.CheckSubscription(&_CompanyServerAccess.CallOpts, _company)
}

// Companies is a free data retrieval call binding the contract method 0x355e6ce8.
//
// Solidity: function companies(address ) view returns(string name, string ceoName, string email, uint256 subscriptionEnd, bool isActive)
func (_CompanyServerAccess *CompanyServerAccessCaller) Companies(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name            string
	CeoName         string
	Email           string
	SubscriptionEnd *big.Int
	IsActive        bool
}, error) {
	var out []interface{}
	err := _CompanyServerAccess.contract.Call(opts, &out, "companies", arg0)

	outstruct := new(struct {
		Name            string
		CeoName         string
		Email           string
		SubscriptionEnd *big.Int
		IsActive        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.CeoName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Email = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.SubscriptionEnd = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Companies is a free data retrieval call binding the contract method 0x355e6ce8.
//
// Solidity: function companies(address ) view returns(string name, string ceoName, string email, uint256 subscriptionEnd, bool isActive)
func (_CompanyServerAccess *CompanyServerAccessSession) Companies(arg0 common.Address) (struct {
	Name            string
	CeoName         string
	Email           string
	SubscriptionEnd *big.Int
	IsActive        bool
}, error) {
	return _CompanyServerAccess.Contract.Companies(&_CompanyServerAccess.CallOpts, arg0)
}

// Companies is a free data retrieval call binding the contract method 0x355e6ce8.
//
// Solidity: function companies(address ) view returns(string name, string ceoName, string email, uint256 subscriptionEnd, bool isActive)
func (_CompanyServerAccess *CompanyServerAccessCallerSession) Companies(arg0 common.Address) (struct {
	Name            string
	CeoName         string
	Email           string
	SubscriptionEnd *big.Int
	IsActive        bool
}, error) {
	return _CompanyServerAccess.Contract.Companies(&_CompanyServerAccess.CallOpts, arg0)
}

// GetCompanyInfo is a free data retrieval call binding the contract method 0x42e12c83.
//
// Solidity: function getCompanyInfo(address _company) view returns(string name, string ceoName, string email, uint256 subscriptionEnd, bool isActive)
func (_CompanyServerAccess *CompanyServerAccessCaller) GetCompanyInfo(opts *bind.CallOpts, _company common.Address) (struct {
	Name            string
	CeoName         string
	Email           string
	SubscriptionEnd *big.Int
	IsActive        bool
}, error) {
	var out []interface{}
	err := _CompanyServerAccess.contract.Call(opts, &out, "getCompanyInfo", _company)

	outstruct := new(struct {
		Name            string
		CeoName         string
		Email           string
		SubscriptionEnd *big.Int
		IsActive        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.CeoName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Email = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.SubscriptionEnd = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// GetCompanyInfo is a free data retrieval call binding the contract method 0x42e12c83.
//
// Solidity: function getCompanyInfo(address _company) view returns(string name, string ceoName, string email, uint256 subscriptionEnd, bool isActive)
func (_CompanyServerAccess *CompanyServerAccessSession) GetCompanyInfo(_company common.Address) (struct {
	Name            string
	CeoName         string
	Email           string
	SubscriptionEnd *big.Int
	IsActive        bool
}, error) {
	return _CompanyServerAccess.Contract.GetCompanyInfo(&_CompanyServerAccess.CallOpts, _company)
}

// GetCompanyInfo is a free data retrieval call binding the contract method 0x42e12c83.
//
// Solidity: function getCompanyInfo(address _company) view returns(string name, string ceoName, string email, uint256 subscriptionEnd, bool isActive)
func (_CompanyServerAccess *CompanyServerAccessCallerSession) GetCompanyInfo(_company common.Address) (struct {
	Name            string
	CeoName         string
	Email           string
	SubscriptionEnd *big.Int
	IsActive        bool
}, error) {
	return _CompanyServerAccess.Contract.GetCompanyInfo(&_CompanyServerAccess.CallOpts, _company)
}

// GetSubscriptionPrice is a free data retrieval call binding the contract method 0xad644ccb.
//
// Solidity: function getSubscriptionPrice(uint256 _type) view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessCaller) GetSubscriptionPrice(opts *bind.CallOpts, _type *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CompanyServerAccess.contract.Call(opts, &out, "getSubscriptionPrice", _type)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSubscriptionPrice is a free data retrieval call binding the contract method 0xad644ccb.
//
// Solidity: function getSubscriptionPrice(uint256 _type) view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessSession) GetSubscriptionPrice(_type *big.Int) (*big.Int, error) {
	return _CompanyServerAccess.Contract.GetSubscriptionPrice(&_CompanyServerAccess.CallOpts, _type)
}

// GetSubscriptionPrice is a free data retrieval call binding the contract method 0xad644ccb.
//
// Solidity: function getSubscriptionPrice(uint256 _type) view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessCallerSession) GetSubscriptionPrice(_type *big.Int) (*big.Int, error) {
	return _CompanyServerAccess.Contract.GetSubscriptionPrice(&_CompanyServerAccess.CallOpts, _type)
}

// MonthlyPrice is a free data retrieval call binding the contract method 0xa06c5a24.
//
// Solidity: function monthlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessCaller) MonthlyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CompanyServerAccess.contract.Call(opts, &out, "monthlyPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MonthlyPrice is a free data retrieval call binding the contract method 0xa06c5a24.
//
// Solidity: function monthlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessSession) MonthlyPrice() (*big.Int, error) {
	return _CompanyServerAccess.Contract.MonthlyPrice(&_CompanyServerAccess.CallOpts)
}

// MonthlyPrice is a free data retrieval call binding the contract method 0xa06c5a24.
//
// Solidity: function monthlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessCallerSession) MonthlyPrice() (*big.Int, error) {
	return _CompanyServerAccess.Contract.MonthlyPrice(&_CompanyServerAccess.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CompanyServerAccess *CompanyServerAccessCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CompanyServerAccess.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CompanyServerAccess *CompanyServerAccessSession) Owner() (common.Address, error) {
	return _CompanyServerAccess.Contract.Owner(&_CompanyServerAccess.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CompanyServerAccess *CompanyServerAccessCallerSession) Owner() (common.Address, error) {
	return _CompanyServerAccess.Contract.Owner(&_CompanyServerAccess.CallOpts)
}

// QuarterlyPrice is a free data retrieval call binding the contract method 0x6d2db884.
//
// Solidity: function quarterlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessCaller) QuarterlyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CompanyServerAccess.contract.Call(opts, &out, "quarterlyPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuarterlyPrice is a free data retrieval call binding the contract method 0x6d2db884.
//
// Solidity: function quarterlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessSession) QuarterlyPrice() (*big.Int, error) {
	return _CompanyServerAccess.Contract.QuarterlyPrice(&_CompanyServerAccess.CallOpts)
}

// QuarterlyPrice is a free data retrieval call binding the contract method 0x6d2db884.
//
// Solidity: function quarterlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessCallerSession) QuarterlyPrice() (*big.Int, error) {
	return _CompanyServerAccess.Contract.QuarterlyPrice(&_CompanyServerAccess.CallOpts)
}

// YearlyPrice is a free data retrieval call binding the contract method 0xff82286b.
//
// Solidity: function yearlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessCaller) YearlyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CompanyServerAccess.contract.Call(opts, &out, "yearlyPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YearlyPrice is a free data retrieval call binding the contract method 0xff82286b.
//
// Solidity: function yearlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessSession) YearlyPrice() (*big.Int, error) {
	return _CompanyServerAccess.Contract.YearlyPrice(&_CompanyServerAccess.CallOpts)
}

// YearlyPrice is a free data retrieval call binding the contract method 0xff82286b.
//
// Solidity: function yearlyPrice() view returns(uint256)
func (_CompanyServerAccess *CompanyServerAccessCallerSession) YearlyPrice() (*big.Int, error) {
	return _CompanyServerAccess.Contract.YearlyPrice(&_CompanyServerAccess.CallOpts)
}

// ExpireSubscription is a paid mutator transaction binding the contract method 0x18a36b09.
//
// Solidity: function expireSubscription(address _company) returns()
func (_CompanyServerAccess *CompanyServerAccessTransactor) ExpireSubscription(opts *bind.TransactOpts, _company common.Address) (*types.Transaction, error) {
	return _CompanyServerAccess.contract.Transact(opts, "expireSubscription", _company)
}

// ExpireSubscription is a paid mutator transaction binding the contract method 0x18a36b09.
//
// Solidity: function expireSubscription(address _company) returns()
func (_CompanyServerAccess *CompanyServerAccessSession) ExpireSubscription(_company common.Address) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.ExpireSubscription(&_CompanyServerAccess.TransactOpts, _company)
}

// ExpireSubscription is a paid mutator transaction binding the contract method 0x18a36b09.
//
// Solidity: function expireSubscription(address _company) returns()
func (_CompanyServerAccess *CompanyServerAccessTransactorSession) ExpireSubscription(_company common.Address) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.ExpireSubscription(&_CompanyServerAccess.TransactOpts, _company)
}

// ExtendSubscription is a paid mutator transaction binding the contract method 0xfc59282d.
//
// Solidity: function extendSubscription(uint256 _subscriptionType) payable returns()
func (_CompanyServerAccess *CompanyServerAccessTransactor) ExtendSubscription(opts *bind.TransactOpts, _subscriptionType *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.contract.Transact(opts, "extendSubscription", _subscriptionType)
}

// ExtendSubscription is a paid mutator transaction binding the contract method 0xfc59282d.
//
// Solidity: function extendSubscription(uint256 _subscriptionType) payable returns()
func (_CompanyServerAccess *CompanyServerAccessSession) ExtendSubscription(_subscriptionType *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.ExtendSubscription(&_CompanyServerAccess.TransactOpts, _subscriptionType)
}

// ExtendSubscription is a paid mutator transaction binding the contract method 0xfc59282d.
//
// Solidity: function extendSubscription(uint256 _subscriptionType) payable returns()
func (_CompanyServerAccess *CompanyServerAccessTransactorSession) ExtendSubscription(_subscriptionType *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.ExtendSubscription(&_CompanyServerAccess.TransactOpts, _subscriptionType)
}

// RegisterCompany is a paid mutator transaction binding the contract method 0xfbb9de45.
//
// Solidity: function registerCompany(string _name, string _ceoName, string _email, uint256 _subscriptionType) payable returns()
func (_CompanyServerAccess *CompanyServerAccessTransactor) RegisterCompany(opts *bind.TransactOpts, _name string, _ceoName string, _email string, _subscriptionType *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.contract.Transact(opts, "registerCompany", _name, _ceoName, _email, _subscriptionType)
}

// RegisterCompany is a paid mutator transaction binding the contract method 0xfbb9de45.
//
// Solidity: function registerCompany(string _name, string _ceoName, string _email, uint256 _subscriptionType) payable returns()
func (_CompanyServerAccess *CompanyServerAccessSession) RegisterCompany(_name string, _ceoName string, _email string, _subscriptionType *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.RegisterCompany(&_CompanyServerAccess.TransactOpts, _name, _ceoName, _email, _subscriptionType)
}

// RegisterCompany is a paid mutator transaction binding the contract method 0xfbb9de45.
//
// Solidity: function registerCompany(string _name, string _ceoName, string _email, uint256 _subscriptionType) payable returns()
func (_CompanyServerAccess *CompanyServerAccessTransactorSession) RegisterCompany(_name string, _ceoName string, _email string, _subscriptionType *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.RegisterCompany(&_CompanyServerAccess.TransactOpts, _name, _ceoName, _email, _subscriptionType)
}

// SetPrices is a paid mutator transaction binding the contract method 0xa88fe42d.
//
// Solidity: function setPrices(uint256 _monthly, uint256 _quarterly, uint256 _yearly) returns()
func (_CompanyServerAccess *CompanyServerAccessTransactor) SetPrices(opts *bind.TransactOpts, _monthly *big.Int, _quarterly *big.Int, _yearly *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.contract.Transact(opts, "setPrices", _monthly, _quarterly, _yearly)
}

// SetPrices is a paid mutator transaction binding the contract method 0xa88fe42d.
//
// Solidity: function setPrices(uint256 _monthly, uint256 _quarterly, uint256 _yearly) returns()
func (_CompanyServerAccess *CompanyServerAccessSession) SetPrices(_monthly *big.Int, _quarterly *big.Int, _yearly *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.SetPrices(&_CompanyServerAccess.TransactOpts, _monthly, _quarterly, _yearly)
}

// SetPrices is a paid mutator transaction binding the contract method 0xa88fe42d.
//
// Solidity: function setPrices(uint256 _monthly, uint256 _quarterly, uint256 _yearly) returns()
func (_CompanyServerAccess *CompanyServerAccessTransactorSession) SetPrices(_monthly *big.Int, _quarterly *big.Int, _yearly *big.Int) (*types.Transaction, error) {
	return _CompanyServerAccess.Contract.SetPrices(&_CompanyServerAccess.TransactOpts, _monthly, _quarterly, _yearly)
}

// CompanyServerAccessCompanyRegisteredIterator is returned from FilterCompanyRegistered and is used to iterate over the raw logs and unpacked data for CompanyRegistered events raised by the CompanyServerAccess contract.
type CompanyServerAccessCompanyRegisteredIterator struct {
	Event *CompanyServerAccessCompanyRegistered // Event containing the contract specifics and raw log

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
func (it *CompanyServerAccessCompanyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompanyServerAccessCompanyRegistered)
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
		it.Event = new(CompanyServerAccessCompanyRegistered)
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
func (it *CompanyServerAccessCompanyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompanyServerAccessCompanyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompanyServerAccessCompanyRegistered represents a CompanyRegistered event raised by the CompanyServerAccess contract.
type CompanyServerAccessCompanyRegistered struct {
	Company         common.Address
	Name            string
	Email           string
	SubscriptionEnd *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCompanyRegistered is a free log retrieval operation binding the contract event 0x17b5e49d185ce3d6009efaa3bbe4d93fbd4da1148be97921127f15c8e7d3b713.
//
// Solidity: event CompanyRegistered(address indexed company, string name, string email, uint256 subscriptionEnd)
func (_CompanyServerAccess *CompanyServerAccessFilterer) FilterCompanyRegistered(opts *bind.FilterOpts, company []common.Address) (*CompanyServerAccessCompanyRegisteredIterator, error) {

	var companyRule []interface{}
	for _, companyItem := range company {
		companyRule = append(companyRule, companyItem)
	}

	logs, sub, err := _CompanyServerAccess.contract.FilterLogs(opts, "CompanyRegistered", companyRule)
	if err != nil {
		return nil, err
	}
	return &CompanyServerAccessCompanyRegisteredIterator{contract: _CompanyServerAccess.contract, event: "CompanyRegistered", logs: logs, sub: sub}, nil
}

// WatchCompanyRegistered is a free log subscription operation binding the contract event 0x17b5e49d185ce3d6009efaa3bbe4d93fbd4da1148be97921127f15c8e7d3b713.
//
// Solidity: event CompanyRegistered(address indexed company, string name, string email, uint256 subscriptionEnd)
func (_CompanyServerAccess *CompanyServerAccessFilterer) WatchCompanyRegistered(opts *bind.WatchOpts, sink chan<- *CompanyServerAccessCompanyRegistered, company []common.Address) (event.Subscription, error) {

	var companyRule []interface{}
	for _, companyItem := range company {
		companyRule = append(companyRule, companyItem)
	}

	logs, sub, err := _CompanyServerAccess.contract.WatchLogs(opts, "CompanyRegistered", companyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompanyServerAccessCompanyRegistered)
				if err := _CompanyServerAccess.contract.UnpackLog(event, "CompanyRegistered", log); err != nil {
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

// ParseCompanyRegistered is a log parse operation binding the contract event 0x17b5e49d185ce3d6009efaa3bbe4d93fbd4da1148be97921127f15c8e7d3b713.
//
// Solidity: event CompanyRegistered(address indexed company, string name, string email, uint256 subscriptionEnd)
func (_CompanyServerAccess *CompanyServerAccessFilterer) ParseCompanyRegistered(log types.Log) (*CompanyServerAccessCompanyRegistered, error) {
	event := new(CompanyServerAccessCompanyRegistered)
	if err := _CompanyServerAccess.contract.UnpackLog(event, "CompanyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompanyServerAccessServiceExpiredIterator is returned from FilterServiceExpired and is used to iterate over the raw logs and unpacked data for ServiceExpired events raised by the CompanyServerAccess contract.
type CompanyServerAccessServiceExpiredIterator struct {
	Event *CompanyServerAccessServiceExpired // Event containing the contract specifics and raw log

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
func (it *CompanyServerAccessServiceExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompanyServerAccessServiceExpired)
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
		it.Event = new(CompanyServerAccessServiceExpired)
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
func (it *CompanyServerAccessServiceExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompanyServerAccessServiceExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompanyServerAccessServiceExpired represents a ServiceExpired event raised by the CompanyServerAccess contract.
type CompanyServerAccessServiceExpired struct {
	Company common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterServiceExpired is a free log retrieval operation binding the contract event 0xf62bf3d3cce73f50844c8aa76ac538f0ef87fafdb91b66248483849f9fb78fe3.
//
// Solidity: event ServiceExpired(address indexed company)
func (_CompanyServerAccess *CompanyServerAccessFilterer) FilterServiceExpired(opts *bind.FilterOpts, company []common.Address) (*CompanyServerAccessServiceExpiredIterator, error) {

	var companyRule []interface{}
	for _, companyItem := range company {
		companyRule = append(companyRule, companyItem)
	}

	logs, sub, err := _CompanyServerAccess.contract.FilterLogs(opts, "ServiceExpired", companyRule)
	if err != nil {
		return nil, err
	}
	return &CompanyServerAccessServiceExpiredIterator{contract: _CompanyServerAccess.contract, event: "ServiceExpired", logs: logs, sub: sub}, nil
}

// WatchServiceExpired is a free log subscription operation binding the contract event 0xf62bf3d3cce73f50844c8aa76ac538f0ef87fafdb91b66248483849f9fb78fe3.
//
// Solidity: event ServiceExpired(address indexed company)
func (_CompanyServerAccess *CompanyServerAccessFilterer) WatchServiceExpired(opts *bind.WatchOpts, sink chan<- *CompanyServerAccessServiceExpired, company []common.Address) (event.Subscription, error) {

	var companyRule []interface{}
	for _, companyItem := range company {
		companyRule = append(companyRule, companyItem)
	}

	logs, sub, err := _CompanyServerAccess.contract.WatchLogs(opts, "ServiceExpired", companyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompanyServerAccessServiceExpired)
				if err := _CompanyServerAccess.contract.UnpackLog(event, "ServiceExpired", log); err != nil {
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

// ParseServiceExpired is a log parse operation binding the contract event 0xf62bf3d3cce73f50844c8aa76ac538f0ef87fafdb91b66248483849f9fb78fe3.
//
// Solidity: event ServiceExpired(address indexed company)
func (_CompanyServerAccess *CompanyServerAccessFilterer) ParseServiceExpired(log types.Log) (*CompanyServerAccessServiceExpired, error) {
	event := new(CompanyServerAccessServiceExpired)
	if err := _CompanyServerAccess.contract.UnpackLog(event, "ServiceExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompanyServerAccessSubscriptionExtendedIterator is returned from FilterSubscriptionExtended and is used to iterate over the raw logs and unpacked data for SubscriptionExtended events raised by the CompanyServerAccess contract.
type CompanyServerAccessSubscriptionExtendedIterator struct {
	Event *CompanyServerAccessSubscriptionExtended // Event containing the contract specifics and raw log

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
func (it *CompanyServerAccessSubscriptionExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompanyServerAccessSubscriptionExtended)
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
		it.Event = new(CompanyServerAccessSubscriptionExtended)
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
func (it *CompanyServerAccessSubscriptionExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompanyServerAccessSubscriptionExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompanyServerAccessSubscriptionExtended represents a SubscriptionExtended event raised by the CompanyServerAccess contract.
type CompanyServerAccessSubscriptionExtended struct {
	Company    common.Address
	NewEndDate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionExtended is a free log retrieval operation binding the contract event 0x6be44fac34dd0021d70df8aa95fb0461c302ee380857ffacd9c764afb2cb2554.
//
// Solidity: event SubscriptionExtended(address indexed company, uint256 newEndDate)
func (_CompanyServerAccess *CompanyServerAccessFilterer) FilterSubscriptionExtended(opts *bind.FilterOpts, company []common.Address) (*CompanyServerAccessSubscriptionExtendedIterator, error) {

	var companyRule []interface{}
	for _, companyItem := range company {
		companyRule = append(companyRule, companyItem)
	}

	logs, sub, err := _CompanyServerAccess.contract.FilterLogs(opts, "SubscriptionExtended", companyRule)
	if err != nil {
		return nil, err
	}
	return &CompanyServerAccessSubscriptionExtendedIterator{contract: _CompanyServerAccess.contract, event: "SubscriptionExtended", logs: logs, sub: sub}, nil
}

// WatchSubscriptionExtended is a free log subscription operation binding the contract event 0x6be44fac34dd0021d70df8aa95fb0461c302ee380857ffacd9c764afb2cb2554.
//
// Solidity: event SubscriptionExtended(address indexed company, uint256 newEndDate)
func (_CompanyServerAccess *CompanyServerAccessFilterer) WatchSubscriptionExtended(opts *bind.WatchOpts, sink chan<- *CompanyServerAccessSubscriptionExtended, company []common.Address) (event.Subscription, error) {

	var companyRule []interface{}
	for _, companyItem := range company {
		companyRule = append(companyRule, companyItem)
	}

	logs, sub, err := _CompanyServerAccess.contract.WatchLogs(opts, "SubscriptionExtended", companyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompanyServerAccessSubscriptionExtended)
				if err := _CompanyServerAccess.contract.UnpackLog(event, "SubscriptionExtended", log); err != nil {
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

// ParseSubscriptionExtended is a log parse operation binding the contract event 0x6be44fac34dd0021d70df8aa95fb0461c302ee380857ffacd9c764afb2cb2554.
//
// Solidity: event SubscriptionExtended(address indexed company, uint256 newEndDate)
func (_CompanyServerAccess *CompanyServerAccessFilterer) ParseSubscriptionExtended(log types.Log) (*CompanyServerAccessSubscriptionExtended, error) {
	event := new(CompanyServerAccessSubscriptionExtended)
	if err := _CompanyServerAccess.contract.UnpackLog(event, "SubscriptionExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
