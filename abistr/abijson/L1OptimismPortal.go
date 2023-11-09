// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abijson

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

// TypesOutputRootProof is an auto generated low-level Go binding around an user-defined struct.
type TypesOutputRootProof struct {
	Version                  [32]byte
	StateRoot                [32]byte
	MessagePasserStorageRoot [32]byte
	LatestBlockhash          [32]byte
}

// TypesWithdrawalTransaction is an auto generated low-level Go binding around an user-defined struct.
type TypesWithdrawalTransaction struct {
	Nonce    *big.Int
	Sender   common.Address
	Target   common.Address
	MntValue *big.Int
	EthValue *big.Int
	GasLimit *big.Int
	Data     []byte
}

// L1OptimismPortalMetaData contains all meta data concerning the L1OptimismPortal contract.
var L1OptimismPortalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractL2OutputOracle\",\"name\":\"_l2Oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"},{\"internalType\":\"contractSystemConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1MNT\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"opaqueData\",\"type\":\"bytes\"}],\"name\":\"TransactionDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"withdrawalHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"withdrawalHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"WithdrawalProven\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GUARDIAN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L1_MNT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_ORACLE\",\"outputs\":[{\"internalType\":\"contractL2OutputOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_CONFIG\",\"outputs\":[{\"internalType\":\"contractSystemConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mntValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mntTxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_isCreation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donateETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mntValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.WithdrawalTransaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"finalizeWithdrawalTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"finalizedWithdrawals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2OutputIndex\",\"type\":\"uint256\"}],\"name\":\"isOutputFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_byteCount\",\"type\":\"uint64\"}],\"name\":\"minimumGasLimit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"params\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"prevBaseFee\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"prevBoughtGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"prevBlockNum\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mntValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.WithdrawalTransaction\",\"name\":\"_tx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"messagePasserStorageRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"latestBlockhash\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.OutputRootProof\",\"name\":\"_outputRootProof\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_withdrawalProof\",\"type\":\"bytes[]\"}],\"name\":\"proveWithdrawalTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"provenWithdrawals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2OutputIndex\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// L1OptimismPortalABI is the input ABI used to generate the binding from.
// Deprecated: Use L1OptimismPortalMetaData.ABI instead.
var L1OptimismPortalABI = L1OptimismPortalMetaData.ABI

// L1OptimismPortal is an auto generated Go binding around an Ethereum contract.
type L1OptimismPortal struct {
	L1OptimismPortalCaller     // Read-only binding to the contract
	L1OptimismPortalTransactor // Write-only binding to the contract
	L1OptimismPortalFilterer   // Log filterer for contract events
}

// L1OptimismPortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1OptimismPortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1OptimismPortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1OptimismPortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1OptimismPortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1OptimismPortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1OptimismPortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1OptimismPortalSession struct {
	Contract     *L1OptimismPortal // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1OptimismPortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1OptimismPortalCallerSession struct {
	Contract *L1OptimismPortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// L1OptimismPortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1OptimismPortalTransactorSession struct {
	Contract     *L1OptimismPortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// L1OptimismPortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1OptimismPortalRaw struct {
	Contract *L1OptimismPortal // Generic contract binding to access the raw methods on
}

// L1OptimismPortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1OptimismPortalCallerRaw struct {
	Contract *L1OptimismPortalCaller // Generic read-only contract binding to access the raw methods on
}

// L1OptimismPortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1OptimismPortalTransactorRaw struct {
	Contract *L1OptimismPortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1OptimismPortal creates a new instance of L1OptimismPortal, bound to a specific deployed contract.
func NewL1OptimismPortal(address common.Address, backend bind.ContractBackend) (*L1OptimismPortal, error) {
	contract, err := bindL1OptimismPortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1OptimismPortal{L1OptimismPortalCaller: L1OptimismPortalCaller{contract: contract}, L1OptimismPortalTransactor: L1OptimismPortalTransactor{contract: contract}, L1OptimismPortalFilterer: L1OptimismPortalFilterer{contract: contract}}, nil
}

// NewL1OptimismPortalCaller creates a new read-only instance of L1OptimismPortal, bound to a specific deployed contract.
func NewL1OptimismPortalCaller(address common.Address, caller bind.ContractCaller) (*L1OptimismPortalCaller, error) {
	contract, err := bindL1OptimismPortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1OptimismPortalCaller{contract: contract}, nil
}

// NewL1OptimismPortalTransactor creates a new write-only instance of L1OptimismPortal, bound to a specific deployed contract.
func NewL1OptimismPortalTransactor(address common.Address, transactor bind.ContractTransactor) (*L1OptimismPortalTransactor, error) {
	contract, err := bindL1OptimismPortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1OptimismPortalTransactor{contract: contract}, nil
}

// NewL1OptimismPortalFilterer creates a new log filterer instance of L1OptimismPortal, bound to a specific deployed contract.
func NewL1OptimismPortalFilterer(address common.Address, filterer bind.ContractFilterer) (*L1OptimismPortalFilterer, error) {
	contract, err := bindL1OptimismPortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1OptimismPortalFilterer{contract: contract}, nil
}

// bindL1OptimismPortal binds a generic wrapper to an already deployed contract.
func bindL1OptimismPortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1OptimismPortalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1OptimismPortal *L1OptimismPortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1OptimismPortal.Contract.L1OptimismPortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1OptimismPortal *L1OptimismPortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.L1OptimismPortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1OptimismPortal *L1OptimismPortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.L1OptimismPortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1OptimismPortal *L1OptimismPortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1OptimismPortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1OptimismPortal *L1OptimismPortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1OptimismPortal *L1OptimismPortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.contract.Transact(opts, method, params...)
}

// GUARDIAN is a free data retrieval call binding the contract method 0x724c184c.
//
// Solidity: function GUARDIAN() view returns(address)
func (_L1OptimismPortal *L1OptimismPortalCaller) GUARDIAN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1OptimismPortal.contract.Call(opts, &out, "GUARDIAN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GUARDIAN is a free data retrieval call binding the contract method 0x724c184c.
//
// Solidity: function GUARDIAN() view returns(address)
func (_L1OptimismPortal *L1OptimismPortalSession) GUARDIAN() (common.Address, error) {
	return _L1OptimismPortal.Contract.GUARDIAN(&_L1OptimismPortal.CallOpts)
}

// GUARDIAN is a free data retrieval call binding the contract method 0x724c184c.
//
// Solidity: function GUARDIAN() view returns(address)
func (_L1OptimismPortal *L1OptimismPortalCallerSession) GUARDIAN() (common.Address, error) {
	return _L1OptimismPortal.Contract.GUARDIAN(&_L1OptimismPortal.CallOpts)
}

// L1MNTADDRESS is a free data retrieval call binding the contract method 0xac6986c5.
//
// Solidity: function L1_MNT_ADDRESS() view returns(address)
func (_L1OptimismPortal *L1OptimismPortalCaller) L1MNTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1OptimismPortal.contract.Call(opts, &out, "L1_MNT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1MNTADDRESS is a free data retrieval call binding the contract method 0xac6986c5.
//
// Solidity: function L1_MNT_ADDRESS() view returns(address)
func (_L1OptimismPortal *L1OptimismPortalSession) L1MNTADDRESS() (common.Address, error) {
	return _L1OptimismPortal.Contract.L1MNTADDRESS(&_L1OptimismPortal.CallOpts)
}

// L1MNTADDRESS is a free data retrieval call binding the contract method 0xac6986c5.
//
// Solidity: function L1_MNT_ADDRESS() view returns(address)
func (_L1OptimismPortal *L1OptimismPortalCallerSession) L1MNTADDRESS() (common.Address, error) {
	return _L1OptimismPortal.Contract.L1MNTADDRESS(&_L1OptimismPortal.CallOpts)
}

// L2ORACLE is a free data retrieval call binding the contract method 0x001c2ff6.
//
// Solidity: function L2_ORACLE() view returns(address)
func (_L1OptimismPortal *L1OptimismPortalCaller) L2ORACLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1OptimismPortal.contract.Call(opts, &out, "L2_ORACLE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ORACLE is a free data retrieval call binding the contract method 0x001c2ff6.
//
// Solidity: function L2_ORACLE() view returns(address)
func (_L1OptimismPortal *L1OptimismPortalSession) L2ORACLE() (common.Address, error) {
	return _L1OptimismPortal.Contract.L2ORACLE(&_L1OptimismPortal.CallOpts)
}

// L2ORACLE is a free data retrieval call binding the contract method 0x001c2ff6.
//
// Solidity: function L2_ORACLE() view returns(address)
func (_L1OptimismPortal *L1OptimismPortalCallerSession) L2ORACLE() (common.Address, error) {
	return _L1OptimismPortal.Contract.L2ORACLE(&_L1OptimismPortal.CallOpts)
}

// SYSTEMCONFIG is a free data retrieval call binding the contract method 0xf0498750.
//
// Solidity: function SYSTEM_CONFIG() view returns(address)
func (_L1OptimismPortal *L1OptimismPortalCaller) SYSTEMCONFIG(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1OptimismPortal.contract.Call(opts, &out, "SYSTEM_CONFIG")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMCONFIG is a free data retrieval call binding the contract method 0xf0498750.
//
// Solidity: function SYSTEM_CONFIG() view returns(address)
func (_L1OptimismPortal *L1OptimismPortalSession) SYSTEMCONFIG() (common.Address, error) {
	return _L1OptimismPortal.Contract.SYSTEMCONFIG(&_L1OptimismPortal.CallOpts)
}

// SYSTEMCONFIG is a free data retrieval call binding the contract method 0xf0498750.
//
// Solidity: function SYSTEM_CONFIG() view returns(address)
func (_L1OptimismPortal *L1OptimismPortalCallerSession) SYSTEMCONFIG() (common.Address, error) {
	return _L1OptimismPortal.Contract.SYSTEMCONFIG(&_L1OptimismPortal.CallOpts)
}

// FinalizedWithdrawals is a free data retrieval call binding the contract method 0xa14238e7.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (_L1OptimismPortal *L1OptimismPortalCaller) FinalizedWithdrawals(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L1OptimismPortal.contract.Call(opts, &out, "finalizedWithdrawals", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FinalizedWithdrawals is a free data retrieval call binding the contract method 0xa14238e7.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (_L1OptimismPortal *L1OptimismPortalSession) FinalizedWithdrawals(arg0 [32]byte) (bool, error) {
	return _L1OptimismPortal.Contract.FinalizedWithdrawals(&_L1OptimismPortal.CallOpts, arg0)
}

// FinalizedWithdrawals is a free data retrieval call binding the contract method 0xa14238e7.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (_L1OptimismPortal *L1OptimismPortalCallerSession) FinalizedWithdrawals(arg0 [32]byte) (bool, error) {
	return _L1OptimismPortal.Contract.FinalizedWithdrawals(&_L1OptimismPortal.CallOpts, arg0)
}

// IsOutputFinalized is a free data retrieval call binding the contract method 0x6dbffb78.
//
// Solidity: function isOutputFinalized(uint256 _l2OutputIndex) view returns(bool)
func (_L1OptimismPortal *L1OptimismPortalCaller) IsOutputFinalized(opts *bind.CallOpts, _l2OutputIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _L1OptimismPortal.contract.Call(opts, &out, "isOutputFinalized", _l2OutputIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOutputFinalized is a free data retrieval call binding the contract method 0x6dbffb78.
//
// Solidity: function isOutputFinalized(uint256 _l2OutputIndex) view returns(bool)
func (_L1OptimismPortal *L1OptimismPortalSession) IsOutputFinalized(_l2OutputIndex *big.Int) (bool, error) {
	return _L1OptimismPortal.Contract.IsOutputFinalized(&_L1OptimismPortal.CallOpts, _l2OutputIndex)
}

// IsOutputFinalized is a free data retrieval call binding the contract method 0x6dbffb78.
//
// Solidity: function isOutputFinalized(uint256 _l2OutputIndex) view returns(bool)
func (_L1OptimismPortal *L1OptimismPortalCallerSession) IsOutputFinalized(_l2OutputIndex *big.Int) (bool, error) {
	return _L1OptimismPortal.Contract.IsOutputFinalized(&_L1OptimismPortal.CallOpts, _l2OutputIndex)
}

// L2Sender is a free data retrieval call binding the contract method 0x9bf62d82.
//
// Solidity: function l2Sender() view returns(address)
func (_L1OptimismPortal *L1OptimismPortalCaller) L2Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1OptimismPortal.contract.Call(opts, &out, "l2Sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Sender is a free data retrieval call binding the contract method 0x9bf62d82.
//
// Solidity: function l2Sender() view returns(address)
func (_L1OptimismPortal *L1OptimismPortalSession) L2Sender() (common.Address, error) {
	return _L1OptimismPortal.Contract.L2Sender(&_L1OptimismPortal.CallOpts)
}

// L2Sender is a free data retrieval call binding the contract method 0x9bf62d82.
//
// Solidity: function l2Sender() view returns(address)
func (_L1OptimismPortal *L1OptimismPortalCallerSession) L2Sender() (common.Address, error) {
	return _L1OptimismPortal.Contract.L2Sender(&_L1OptimismPortal.CallOpts)
}

// MinimumGasLimit is a free data retrieval call binding the contract method 0xa35d99df.
//
// Solidity: function minimumGasLimit(uint64 _byteCount) pure returns(uint64)
func (_L1OptimismPortal *L1OptimismPortalCaller) MinimumGasLimit(opts *bind.CallOpts, _byteCount uint64) (uint64, error) {
	var out []interface{}
	err := _L1OptimismPortal.contract.Call(opts, &out, "minimumGasLimit", _byteCount)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinimumGasLimit is a free data retrieval call binding the contract method 0xa35d99df.
//
// Solidity: function minimumGasLimit(uint64 _byteCount) pure returns(uint64)
func (_L1OptimismPortal *L1OptimismPortalSession) MinimumGasLimit(_byteCount uint64) (uint64, error) {
	return _L1OptimismPortal.Contract.MinimumGasLimit(&_L1OptimismPortal.CallOpts, _byteCount)
}

// MinimumGasLimit is a free data retrieval call binding the contract method 0xa35d99df.
//
// Solidity: function minimumGasLimit(uint64 _byteCount) pure returns(uint64)
func (_L1OptimismPortal *L1OptimismPortalCallerSession) MinimumGasLimit(_byteCount uint64) (uint64, error) {
	return _L1OptimismPortal.Contract.MinimumGasLimit(&_L1OptimismPortal.CallOpts, _byteCount)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(uint128 prevBaseFee, uint64 prevBoughtGas, uint64 prevBlockNum)
func (_L1OptimismPortal *L1OptimismPortalCaller) Params(opts *bind.CallOpts) (struct {
	PrevBaseFee   *big.Int
	PrevBoughtGas uint64
	PrevBlockNum  uint64
}, error) {
	var out []interface{}
	err := _L1OptimismPortal.contract.Call(opts, &out, "params")

	outstruct := new(struct {
		PrevBaseFee   *big.Int
		PrevBoughtGas uint64
		PrevBlockNum  uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PrevBaseFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PrevBoughtGas = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.PrevBlockNum = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(uint128 prevBaseFee, uint64 prevBoughtGas, uint64 prevBlockNum)
func (_L1OptimismPortal *L1OptimismPortalSession) Params() (struct {
	PrevBaseFee   *big.Int
	PrevBoughtGas uint64
	PrevBlockNum  uint64
}, error) {
	return _L1OptimismPortal.Contract.Params(&_L1OptimismPortal.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(uint128 prevBaseFee, uint64 prevBoughtGas, uint64 prevBlockNum)
func (_L1OptimismPortal *L1OptimismPortalCallerSession) Params() (struct {
	PrevBaseFee   *big.Int
	PrevBoughtGas uint64
	PrevBlockNum  uint64
}, error) {
	return _L1OptimismPortal.Contract.Params(&_L1OptimismPortal.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1OptimismPortal *L1OptimismPortalCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L1OptimismPortal.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1OptimismPortal *L1OptimismPortalSession) Paused() (bool, error) {
	return _L1OptimismPortal.Contract.Paused(&_L1OptimismPortal.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1OptimismPortal *L1OptimismPortalCallerSession) Paused() (bool, error) {
	return _L1OptimismPortal.Contract.Paused(&_L1OptimismPortal.CallOpts)
}

// ProvenWithdrawals is a free data retrieval call binding the contract method 0xe965084c.
//
// Solidity: function provenWithdrawals(bytes32 ) view returns(bytes32 outputRoot, uint128 timestamp, uint128 l2OutputIndex)
func (_L1OptimismPortal *L1OptimismPortalCaller) ProvenWithdrawals(opts *bind.CallOpts, arg0 [32]byte) (struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2OutputIndex *big.Int
}, error) {
	var out []interface{}
	err := _L1OptimismPortal.contract.Call(opts, &out, "provenWithdrawals", arg0)

	outstruct := new(struct {
		OutputRoot    [32]byte
		Timestamp     *big.Int
		L2OutputIndex *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutputRoot = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.L2OutputIndex = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProvenWithdrawals is a free data retrieval call binding the contract method 0xe965084c.
//
// Solidity: function provenWithdrawals(bytes32 ) view returns(bytes32 outputRoot, uint128 timestamp, uint128 l2OutputIndex)
func (_L1OptimismPortal *L1OptimismPortalSession) ProvenWithdrawals(arg0 [32]byte) (struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2OutputIndex *big.Int
}, error) {
	return _L1OptimismPortal.Contract.ProvenWithdrawals(&_L1OptimismPortal.CallOpts, arg0)
}

// ProvenWithdrawals is a free data retrieval call binding the contract method 0xe965084c.
//
// Solidity: function provenWithdrawals(bytes32 ) view returns(bytes32 outputRoot, uint128 timestamp, uint128 l2OutputIndex)
func (_L1OptimismPortal *L1OptimismPortalCallerSession) ProvenWithdrawals(arg0 [32]byte) (struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2OutputIndex *big.Int
}, error) {
	return _L1OptimismPortal.Contract.ProvenWithdrawals(&_L1OptimismPortal.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L1OptimismPortal *L1OptimismPortalCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L1OptimismPortal.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L1OptimismPortal *L1OptimismPortalSession) Version() (string, error) {
	return _L1OptimismPortal.Contract.Version(&_L1OptimismPortal.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L1OptimismPortal *L1OptimismPortalCallerSession) Version() (string, error) {
	return _L1OptimismPortal.Contract.Version(&_L1OptimismPortal.CallOpts)
}

// DepositTransaction is a paid mutator transaction binding the contract method 0x85945feb.
//
// Solidity: function depositTransaction(uint256 _mntValue, address _to, uint256 _mntTxValue, uint64 _gasLimit, bool _isCreation, bytes _data) payable returns()
func (_L1OptimismPortal *L1OptimismPortalTransactor) DepositTransaction(opts *bind.TransactOpts, _mntValue *big.Int, _to common.Address, _mntTxValue *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _L1OptimismPortal.contract.Transact(opts, "depositTransaction", _mntValue, _to, _mntTxValue, _gasLimit, _isCreation, _data)
}

// DepositTransaction is a paid mutator transaction binding the contract method 0x85945feb.
//
// Solidity: function depositTransaction(uint256 _mntValue, address _to, uint256 _mntTxValue, uint64 _gasLimit, bool _isCreation, bytes _data) payable returns()
func (_L1OptimismPortal *L1OptimismPortalSession) DepositTransaction(_mntValue *big.Int, _to common.Address, _mntTxValue *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.DepositTransaction(&_L1OptimismPortal.TransactOpts, _mntValue, _to, _mntTxValue, _gasLimit, _isCreation, _data)
}

// DepositTransaction is a paid mutator transaction binding the contract method 0x85945feb.
//
// Solidity: function depositTransaction(uint256 _mntValue, address _to, uint256 _mntTxValue, uint64 _gasLimit, bool _isCreation, bytes _data) payable returns()
func (_L1OptimismPortal *L1OptimismPortalTransactorSession) DepositTransaction(_mntValue *big.Int, _to common.Address, _mntTxValue *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.DepositTransaction(&_L1OptimismPortal.TransactOpts, _mntValue, _to, _mntTxValue, _gasLimit, _isCreation, _data)
}

// DonateETH is a paid mutator transaction binding the contract method 0x8b4c40b0.
//
// Solidity: function donateETH() payable returns()
func (_L1OptimismPortal *L1OptimismPortalTransactor) DonateETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1OptimismPortal.contract.Transact(opts, "donateETH")
}

// DonateETH is a paid mutator transaction binding the contract method 0x8b4c40b0.
//
// Solidity: function donateETH() payable returns()
func (_L1OptimismPortal *L1OptimismPortalSession) DonateETH() (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.DonateETH(&_L1OptimismPortal.TransactOpts)
}

// DonateETH is a paid mutator transaction binding the contract method 0x8b4c40b0.
//
// Solidity: function donateETH() payable returns()
func (_L1OptimismPortal *L1OptimismPortalTransactorSession) DonateETH() (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.DonateETH(&_L1OptimismPortal.TransactOpts)
}

// FinalizeWithdrawalTransaction is a paid mutator transaction binding the contract method 0x2e71d4a4.
//
// Solidity: function finalizeWithdrawalTransaction((uint256,address,address,uint256,uint256,uint256,bytes) _tx) returns()
func (_L1OptimismPortal *L1OptimismPortalTransactor) FinalizeWithdrawalTransaction(opts *bind.TransactOpts, _tx TypesWithdrawalTransaction) (*types.Transaction, error) {
	return _L1OptimismPortal.contract.Transact(opts, "finalizeWithdrawalTransaction", _tx)
}

// FinalizeWithdrawalTransaction is a paid mutator transaction binding the contract method 0x2e71d4a4.
//
// Solidity: function finalizeWithdrawalTransaction((uint256,address,address,uint256,uint256,uint256,bytes) _tx) returns()
func (_L1OptimismPortal *L1OptimismPortalSession) FinalizeWithdrawalTransaction(_tx TypesWithdrawalTransaction) (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.FinalizeWithdrawalTransaction(&_L1OptimismPortal.TransactOpts, _tx)
}

// FinalizeWithdrawalTransaction is a paid mutator transaction binding the contract method 0x2e71d4a4.
//
// Solidity: function finalizeWithdrawalTransaction((uint256,address,address,uint256,uint256,uint256,bytes) _tx) returns()
func (_L1OptimismPortal *L1OptimismPortalTransactorSession) FinalizeWithdrawalTransaction(_tx TypesWithdrawalTransaction) (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.FinalizeWithdrawalTransaction(&_L1OptimismPortal.TransactOpts, _tx)
}

// Initialize is a paid mutator transaction binding the contract method 0xd53a822f.
//
// Solidity: function initialize(bool _paused) returns()
func (_L1OptimismPortal *L1OptimismPortalTransactor) Initialize(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _L1OptimismPortal.contract.Transact(opts, "initialize", _paused)
}

// Initialize is a paid mutator transaction binding the contract method 0xd53a822f.
//
// Solidity: function initialize(bool _paused) returns()
func (_L1OptimismPortal *L1OptimismPortalSession) Initialize(_paused bool) (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.Initialize(&_L1OptimismPortal.TransactOpts, _paused)
}

// Initialize is a paid mutator transaction binding the contract method 0xd53a822f.
//
// Solidity: function initialize(bool _paused) returns()
func (_L1OptimismPortal *L1OptimismPortalTransactorSession) Initialize(_paused bool) (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.Initialize(&_L1OptimismPortal.TransactOpts, _paused)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1OptimismPortal *L1OptimismPortalTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1OptimismPortal.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1OptimismPortal *L1OptimismPortalSession) Pause() (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.Pause(&_L1OptimismPortal.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1OptimismPortal *L1OptimismPortalTransactorSession) Pause() (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.Pause(&_L1OptimismPortal.TransactOpts)
}

// ProveWithdrawalTransaction is a paid mutator transaction binding the contract method 0xd69b2b1b.
//
// Solidity: function proveWithdrawalTransaction((uint256,address,address,uint256,uint256,uint256,bytes) _tx, uint256 _l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes[] _withdrawalProof) returns()
func (_L1OptimismPortal *L1OptimismPortalTransactor) ProveWithdrawalTransaction(opts *bind.TransactOpts, _tx TypesWithdrawalTransaction, _l2OutputIndex *big.Int, _outputRootProof TypesOutputRootProof, _withdrawalProof [][]byte) (*types.Transaction, error) {
	return _L1OptimismPortal.contract.Transact(opts, "proveWithdrawalTransaction", _tx, _l2OutputIndex, _outputRootProof, _withdrawalProof)
}

// ProveWithdrawalTransaction is a paid mutator transaction binding the contract method 0xd69b2b1b.
//
// Solidity: function proveWithdrawalTransaction((uint256,address,address,uint256,uint256,uint256,bytes) _tx, uint256 _l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes[] _withdrawalProof) returns()
func (_L1OptimismPortal *L1OptimismPortalSession) ProveWithdrawalTransaction(_tx TypesWithdrawalTransaction, _l2OutputIndex *big.Int, _outputRootProof TypesOutputRootProof, _withdrawalProof [][]byte) (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.ProveWithdrawalTransaction(&_L1OptimismPortal.TransactOpts, _tx, _l2OutputIndex, _outputRootProof, _withdrawalProof)
}

// ProveWithdrawalTransaction is a paid mutator transaction binding the contract method 0xd69b2b1b.
//
// Solidity: function proveWithdrawalTransaction((uint256,address,address,uint256,uint256,uint256,bytes) _tx, uint256 _l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes[] _withdrawalProof) returns()
func (_L1OptimismPortal *L1OptimismPortalTransactorSession) ProveWithdrawalTransaction(_tx TypesWithdrawalTransaction, _l2OutputIndex *big.Int, _outputRootProof TypesOutputRootProof, _withdrawalProof [][]byte) (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.ProveWithdrawalTransaction(&_L1OptimismPortal.TransactOpts, _tx, _l2OutputIndex, _outputRootProof, _withdrawalProof)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1OptimismPortal *L1OptimismPortalTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1OptimismPortal.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1OptimismPortal *L1OptimismPortalSession) Unpause() (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.Unpause(&_L1OptimismPortal.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1OptimismPortal *L1OptimismPortalTransactorSession) Unpause() (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.Unpause(&_L1OptimismPortal.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1OptimismPortal *L1OptimismPortalTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1OptimismPortal.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1OptimismPortal *L1OptimismPortalSession) Receive() (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.Receive(&_L1OptimismPortal.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1OptimismPortal *L1OptimismPortalTransactorSession) Receive() (*types.Transaction, error) {
	return _L1OptimismPortal.Contract.Receive(&_L1OptimismPortal.TransactOpts)
}

// L1OptimismPortalInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1OptimismPortal contract.
type L1OptimismPortalInitializedIterator struct {
	Event *L1OptimismPortalInitialized // Event containing the contract specifics and raw log

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
func (it *L1OptimismPortalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1OptimismPortalInitialized)
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
		it.Event = new(L1OptimismPortalInitialized)
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
func (it *L1OptimismPortalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1OptimismPortalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1OptimismPortalInitialized represents a Initialized event raised by the L1OptimismPortal contract.
type L1OptimismPortalInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1OptimismPortal *L1OptimismPortalFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1OptimismPortalInitializedIterator, error) {

	logs, sub, err := _L1OptimismPortal.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1OptimismPortalInitializedIterator{contract: _L1OptimismPortal.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1OptimismPortal *L1OptimismPortalFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1OptimismPortalInitialized) (event.Subscription, error) {

	logs, sub, err := _L1OptimismPortal.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1OptimismPortalInitialized)
				if err := _L1OptimismPortal.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1OptimismPortal *L1OptimismPortalFilterer) ParseInitialized(log types.Log) (*L1OptimismPortalInitialized, error) {
	event := new(L1OptimismPortalInitialized)
	if err := _L1OptimismPortal.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1OptimismPortalPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the L1OptimismPortal contract.
type L1OptimismPortalPausedIterator struct {
	Event *L1OptimismPortalPaused // Event containing the contract specifics and raw log

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
func (it *L1OptimismPortalPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1OptimismPortalPaused)
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
		it.Event = new(L1OptimismPortalPaused)
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
func (it *L1OptimismPortalPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1OptimismPortalPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1OptimismPortalPaused represents a Paused event raised by the L1OptimismPortal contract.
type L1OptimismPortalPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1OptimismPortal *L1OptimismPortalFilterer) FilterPaused(opts *bind.FilterOpts) (*L1OptimismPortalPausedIterator, error) {

	logs, sub, err := _L1OptimismPortal.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &L1OptimismPortalPausedIterator{contract: _L1OptimismPortal.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1OptimismPortal *L1OptimismPortalFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *L1OptimismPortalPaused) (event.Subscription, error) {

	logs, sub, err := _L1OptimismPortal.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1OptimismPortalPaused)
				if err := _L1OptimismPortal.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1OptimismPortal *L1OptimismPortalFilterer) ParsePaused(log types.Log) (*L1OptimismPortalPaused, error) {
	event := new(L1OptimismPortalPaused)
	if err := _L1OptimismPortal.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1OptimismPortalTransactionDepositedIterator is returned from FilterTransactionDeposited and is used to iterate over the raw logs and unpacked data for TransactionDeposited events raised by the L1OptimismPortal contract.
type L1OptimismPortalTransactionDepositedIterator struct {
	Event *L1OptimismPortalTransactionDeposited // Event containing the contract specifics and raw log

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
func (it *L1OptimismPortalTransactionDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1OptimismPortalTransactionDeposited)
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
		it.Event = new(L1OptimismPortalTransactionDeposited)
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
func (it *L1OptimismPortalTransactionDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1OptimismPortalTransactionDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1OptimismPortalTransactionDeposited represents a TransactionDeposited event raised by the L1OptimismPortal contract.
type L1OptimismPortalTransactionDeposited struct {
	From       common.Address
	To         common.Address
	Version    *big.Int
	OpaqueData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransactionDeposited is a free log retrieval operation binding the contract event 0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32.
//
// Solidity: event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData)
func (_L1OptimismPortal *L1OptimismPortalFilterer) FilterTransactionDeposited(opts *bind.FilterOpts, from []common.Address, to []common.Address, version []*big.Int) (*L1OptimismPortalTransactionDepositedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _L1OptimismPortal.contract.FilterLogs(opts, "TransactionDeposited", fromRule, toRule, versionRule)
	if err != nil {
		return nil, err
	}
	return &L1OptimismPortalTransactionDepositedIterator{contract: _L1OptimismPortal.contract, event: "TransactionDeposited", logs: logs, sub: sub}, nil
}

// WatchTransactionDeposited is a free log subscription operation binding the contract event 0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32.
//
// Solidity: event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData)
func (_L1OptimismPortal *L1OptimismPortalFilterer) WatchTransactionDeposited(opts *bind.WatchOpts, sink chan<- *L1OptimismPortalTransactionDeposited, from []common.Address, to []common.Address, version []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _L1OptimismPortal.contract.WatchLogs(opts, "TransactionDeposited", fromRule, toRule, versionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1OptimismPortalTransactionDeposited)
				if err := _L1OptimismPortal.contract.UnpackLog(event, "TransactionDeposited", log); err != nil {
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

// ParseTransactionDeposited is a log parse operation binding the contract event 0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32.
//
// Solidity: event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData)
func (_L1OptimismPortal *L1OptimismPortalFilterer) ParseTransactionDeposited(log types.Log) (*L1OptimismPortalTransactionDeposited, error) {
	event := new(L1OptimismPortalTransactionDeposited)
	if err := _L1OptimismPortal.contract.UnpackLog(event, "TransactionDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1OptimismPortalUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the L1OptimismPortal contract.
type L1OptimismPortalUnpausedIterator struct {
	Event *L1OptimismPortalUnpaused // Event containing the contract specifics and raw log

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
func (it *L1OptimismPortalUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1OptimismPortalUnpaused)
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
		it.Event = new(L1OptimismPortalUnpaused)
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
func (it *L1OptimismPortalUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1OptimismPortalUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1OptimismPortalUnpaused represents a Unpaused event raised by the L1OptimismPortal contract.
type L1OptimismPortalUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1OptimismPortal *L1OptimismPortalFilterer) FilterUnpaused(opts *bind.FilterOpts) (*L1OptimismPortalUnpausedIterator, error) {

	logs, sub, err := _L1OptimismPortal.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &L1OptimismPortalUnpausedIterator{contract: _L1OptimismPortal.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1OptimismPortal *L1OptimismPortalFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *L1OptimismPortalUnpaused) (event.Subscription, error) {

	logs, sub, err := _L1OptimismPortal.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1OptimismPortalUnpaused)
				if err := _L1OptimismPortal.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1OptimismPortal *L1OptimismPortalFilterer) ParseUnpaused(log types.Log) (*L1OptimismPortalUnpaused, error) {
	event := new(L1OptimismPortalUnpaused)
	if err := _L1OptimismPortal.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1OptimismPortalWithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the L1OptimismPortal contract.
type L1OptimismPortalWithdrawalFinalizedIterator struct {
	Event *L1OptimismPortalWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *L1OptimismPortalWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1OptimismPortalWithdrawalFinalized)
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
		it.Event = new(L1OptimismPortalWithdrawalFinalized)
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
func (it *L1OptimismPortalWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1OptimismPortalWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1OptimismPortalWithdrawalFinalized represents a WithdrawalFinalized event raised by the L1OptimismPortal contract.
type L1OptimismPortalWithdrawalFinalized struct {
	WithdrawalHash [32]byte
	Success        bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalized is a free log retrieval operation binding the contract event 0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (_L1OptimismPortal *L1OptimismPortalFilterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, withdrawalHash [][32]byte) (*L1OptimismPortalWithdrawalFinalizedIterator, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}

	logs, sub, err := _L1OptimismPortal.contract.FilterLogs(opts, "WithdrawalFinalized", withdrawalHashRule)
	if err != nil {
		return nil, err
	}
	return &L1OptimismPortalWithdrawalFinalizedIterator{contract: _L1OptimismPortal.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (_L1OptimismPortal *L1OptimismPortalFilterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *L1OptimismPortalWithdrawalFinalized, withdrawalHash [][32]byte) (event.Subscription, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}

	logs, sub, err := _L1OptimismPortal.contract.WatchLogs(opts, "WithdrawalFinalized", withdrawalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1OptimismPortalWithdrawalFinalized)
				if err := _L1OptimismPortal.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
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

// ParseWithdrawalFinalized is a log parse operation binding the contract event 0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (_L1OptimismPortal *L1OptimismPortalFilterer) ParseWithdrawalFinalized(log types.Log) (*L1OptimismPortalWithdrawalFinalized, error) {
	event := new(L1OptimismPortalWithdrawalFinalized)
	if err := _L1OptimismPortal.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1OptimismPortalWithdrawalProvenIterator is returned from FilterWithdrawalProven and is used to iterate over the raw logs and unpacked data for WithdrawalProven events raised by the L1OptimismPortal contract.
type L1OptimismPortalWithdrawalProvenIterator struct {
	Event *L1OptimismPortalWithdrawalProven // Event containing the contract specifics and raw log

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
func (it *L1OptimismPortalWithdrawalProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1OptimismPortalWithdrawalProven)
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
		it.Event = new(L1OptimismPortalWithdrawalProven)
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
func (it *L1OptimismPortalWithdrawalProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1OptimismPortalWithdrawalProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1OptimismPortalWithdrawalProven represents a WithdrawalProven event raised by the L1OptimismPortal contract.
type L1OptimismPortalWithdrawalProven struct {
	WithdrawalHash [32]byte
	From           common.Address
	To             common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalProven is a free log retrieval operation binding the contract event 0x67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f62.
//
// Solidity: event WithdrawalProven(bytes32 indexed withdrawalHash, address indexed from, address indexed to)
func (_L1OptimismPortal *L1OptimismPortalFilterer) FilterWithdrawalProven(opts *bind.FilterOpts, withdrawalHash [][32]byte, from []common.Address, to []common.Address) (*L1OptimismPortalWithdrawalProvenIterator, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1OptimismPortal.contract.FilterLogs(opts, "WithdrawalProven", withdrawalHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1OptimismPortalWithdrawalProvenIterator{contract: _L1OptimismPortal.contract, event: "WithdrawalProven", logs: logs, sub: sub}, nil
}

// WatchWithdrawalProven is a free log subscription operation binding the contract event 0x67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f62.
//
// Solidity: event WithdrawalProven(bytes32 indexed withdrawalHash, address indexed from, address indexed to)
func (_L1OptimismPortal *L1OptimismPortalFilterer) WatchWithdrawalProven(opts *bind.WatchOpts, sink chan<- *L1OptimismPortalWithdrawalProven, withdrawalHash [][32]byte, from []common.Address, to []common.Address) (event.Subscription, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1OptimismPortal.contract.WatchLogs(opts, "WithdrawalProven", withdrawalHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1OptimismPortalWithdrawalProven)
				if err := _L1OptimismPortal.contract.UnpackLog(event, "WithdrawalProven", log); err != nil {
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

// ParseWithdrawalProven is a log parse operation binding the contract event 0x67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f62.
//
// Solidity: event WithdrawalProven(bytes32 indexed withdrawalHash, address indexed from, address indexed to)
func (_L1OptimismPortal *L1OptimismPortalFilterer) ParseWithdrawalProven(log types.Log) (*L1OptimismPortalWithdrawalProven, error) {
	event := new(L1OptimismPortalWithdrawalProven)
	if err := _L1OptimismPortal.contract.UnpackLog(event, "WithdrawalProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
