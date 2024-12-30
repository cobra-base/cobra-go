// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancakeswap_binding

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
	_ = abi.ConvertType
)

// IQuoterV2QuoteExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuoterV2QuoteExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	AmountIn          *big.Int
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IQuoterV2QuoteExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuoterV2QuoteExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Amount            *big.Int
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
}

// PancakeswapV3QuoterMetaData contains all meta data concerning the PancakeswapV3Quoter contract.
var PancakeswapV3QuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_deployer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"pancakeV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"quoteExactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PancakeswapV3QuoterABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakeswapV3QuoterMetaData.ABI instead.
var PancakeswapV3QuoterABI = PancakeswapV3QuoterMetaData.ABI

// PancakeswapV3Quoter is an auto generated Go binding around an Ethereum contract.
type PancakeswapV3Quoter struct {
	PancakeswapV3QuoterCaller     // Read-only binding to the contract
	PancakeswapV3QuoterTransactor // Write-only binding to the contract
	PancakeswapV3QuoterFilterer   // Log filterer for contract events
}

// PancakeswapV3QuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeswapV3QuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeswapV3QuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeswapV3QuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeswapV3QuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeswapV3QuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeswapV3QuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeswapV3QuoterSession struct {
	Contract     *PancakeswapV3Quoter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PancakeswapV3QuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeswapV3QuoterCallerSession struct {
	Contract *PancakeswapV3QuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PancakeswapV3QuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeswapV3QuoterTransactorSession struct {
	Contract     *PancakeswapV3QuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PancakeswapV3QuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeswapV3QuoterRaw struct {
	Contract *PancakeswapV3Quoter // Generic contract binding to access the raw methods on
}

// PancakeswapV3QuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeswapV3QuoterCallerRaw struct {
	Contract *PancakeswapV3QuoterCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeswapV3QuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeswapV3QuoterTransactorRaw struct {
	Contract *PancakeswapV3QuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeswapV3Quoter creates a new instance of PancakeswapV3Quoter, bound to a specific deployed contract.
func NewPancakeswapV3Quoter(address common.Address, backend bind.ContractBackend) (*PancakeswapV3Quoter, error) {
	contract, err := bindPancakeswapV3Quoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakeswapV3Quoter{PancakeswapV3QuoterCaller: PancakeswapV3QuoterCaller{contract: contract}, PancakeswapV3QuoterTransactor: PancakeswapV3QuoterTransactor{contract: contract}, PancakeswapV3QuoterFilterer: PancakeswapV3QuoterFilterer{contract: contract}}, nil
}

// NewPancakeswapV3QuoterCaller creates a new read-only instance of PancakeswapV3Quoter, bound to a specific deployed contract.
func NewPancakeswapV3QuoterCaller(address common.Address, caller bind.ContractCaller) (*PancakeswapV3QuoterCaller, error) {
	contract, err := bindPancakeswapV3Quoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeswapV3QuoterCaller{contract: contract}, nil
}

// NewPancakeswapV3QuoterTransactor creates a new write-only instance of PancakeswapV3Quoter, bound to a specific deployed contract.
func NewPancakeswapV3QuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeswapV3QuoterTransactor, error) {
	contract, err := bindPancakeswapV3Quoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeswapV3QuoterTransactor{contract: contract}, nil
}

// NewPancakeswapV3QuoterFilterer creates a new log filterer instance of PancakeswapV3Quoter, bound to a specific deployed contract.
func NewPancakeswapV3QuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeswapV3QuoterFilterer, error) {
	contract, err := bindPancakeswapV3Quoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeswapV3QuoterFilterer{contract: contract}, nil
}

// bindPancakeswapV3Quoter binds a generic wrapper to an already deployed contract.
func bindPancakeswapV3Quoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PancakeswapV3QuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeswapV3Quoter *PancakeswapV3QuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeswapV3Quoter.Contract.PancakeswapV3QuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeswapV3Quoter *PancakeswapV3QuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeswapV3Quoter.Contract.PancakeswapV3QuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeswapV3Quoter *PancakeswapV3QuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeswapV3Quoter.Contract.PancakeswapV3QuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeswapV3Quoter *PancakeswapV3QuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeswapV3Quoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeswapV3Quoter *PancakeswapV3QuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeswapV3Quoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeswapV3Quoter *PancakeswapV3QuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeswapV3Quoter.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeswapV3Quoter.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterSession) WETH9() (common.Address, error) {
	return _PancakeswapV3Quoter.Contract.WETH9(&_PancakeswapV3Quoter.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterCallerSession) WETH9() (common.Address, error) {
	return _PancakeswapV3Quoter.Contract.WETH9(&_PancakeswapV3Quoter.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterCaller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeswapV3Quoter.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterSession) Deployer() (common.Address, error) {
	return _PancakeswapV3Quoter.Contract.Deployer(&_PancakeswapV3Quoter.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterCallerSession) Deployer() (common.Address, error) {
	return _PancakeswapV3Quoter.Contract.Deployer(&_PancakeswapV3Quoter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeswapV3Quoter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterSession) Factory() (common.Address, error) {
	return _PancakeswapV3Quoter.Contract.Factory(&_PancakeswapV3Quoter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterCallerSession) Factory() (common.Address, error) {
	return _PancakeswapV3Quoter.Contract.Factory(&_PancakeswapV3Quoter.CallOpts)
}

// PancakeV3SwapCallback is a free data retrieval call binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_PancakeswapV3Quoter *PancakeswapV3QuoterCaller) PancakeV3SwapCallback(opts *bind.CallOpts, amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	var out []interface{}
	err := _PancakeswapV3Quoter.contract.Call(opts, &out, "pancakeV3SwapCallback", amount0Delta, amount1Delta, path)

	if err != nil {
		return err
	}

	return err

}

// PancakeV3SwapCallback is a free data retrieval call binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_PancakeswapV3Quoter *PancakeswapV3QuoterSession) PancakeV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _PancakeswapV3Quoter.Contract.PancakeV3SwapCallback(&_PancakeswapV3Quoter.CallOpts, amount0Delta, amount1Delta, path)
}

// PancakeV3SwapCallback is a free data retrieval call binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_PancakeswapV3Quoter *PancakeswapV3QuoterCallerSession) PancakeV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _PancakeswapV3Quoter.Contract.PancakeV3SwapCallback(&_PancakeswapV3Quoter.CallOpts, amount0Delta, amount1Delta, path)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterTransactor) QuoteExactInput(opts *bind.TransactOpts, path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _PancakeswapV3Quoter.contract.Transact(opts, "quoteExactInput", path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _PancakeswapV3Quoter.Contract.QuoteExactInput(&_PancakeswapV3Quoter.TransactOpts, path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterTransactorSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _PancakeswapV3Quoter.Contract.QuoteExactInput(&_PancakeswapV3Quoter.TransactOpts, path, amountIn)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterTransactor) QuoteExactInputSingle(opts *bind.TransactOpts, params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _PancakeswapV3Quoter.contract.Transact(opts, "quoteExactInputSingle", params)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterSession) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _PancakeswapV3Quoter.Contract.QuoteExactInputSingle(&_PancakeswapV3Quoter.TransactOpts, params)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterTransactorSession) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _PancakeswapV3Quoter.Contract.QuoteExactInputSingle(&_PancakeswapV3Quoter.TransactOpts, params)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterTransactor) QuoteExactOutput(opts *bind.TransactOpts, path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _PancakeswapV3Quoter.contract.Transact(opts, "quoteExactOutput", path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _PancakeswapV3Quoter.Contract.QuoteExactOutput(&_PancakeswapV3Quoter.TransactOpts, path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterTransactorSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _PancakeswapV3Quoter.Contract.QuoteExactOutput(&_PancakeswapV3Quoter.TransactOpts, path, amountOut)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterTransactor) QuoteExactOutputSingle(opts *bind.TransactOpts, params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _PancakeswapV3Quoter.contract.Transact(opts, "quoteExactOutputSingle", params)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterSession) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _PancakeswapV3Quoter.Contract.QuoteExactOutputSingle(&_PancakeswapV3Quoter.TransactOpts, params)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_PancakeswapV3Quoter *PancakeswapV3QuoterTransactorSession) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _PancakeswapV3Quoter.Contract.QuoteExactOutputSingle(&_PancakeswapV3Quoter.TransactOpts, params)
}
