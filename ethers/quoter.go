package ethers

import (
	pancakeswap_binding "github.com/cobra-base/cobra-go/ethers/binding/pancakeswap"
	uniswap_binding "github.com/cobra-base/cobra-go/ethers/binding/uniswap"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type QuoteSingleResult struct {
	AmountOut               *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}

type QuoteResult struct {
	AmountOut                   *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}

type Quoter struct {
}

var quoterInstance = &Quoter{}

func GetQuoter() *Quoter {
	return quoterInstance
}

func (s *Quoter) GetAmountsOutForUniswapV2(routerAddress common.Address, path []common.Address, amountIn *big.Int, endpoint string) ([]*big.Int, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	router, err := uniswap_binding.NewUniswapV2Router(routerAddress, client)
	if err != nil {
		return nil, err
	}

	amountsOut, err := router.GetAmountsOut(nil, amountIn, path)
	if err != nil {
		return nil, err
	}

	return amountsOut, nil
}

func (s *Quoter) QuoteExactInputSingleForUniswapV3(quoterAddress common.Address, tokenIn common.Address, tokenOut common.Address,
	amountIn *big.Int, fee *big.Int, endpoint string) (*QuoteSingleResult, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	quoter, err := uniswap_binding.NewUniswapV3Quoter(quoterAddress, client)
	if err != nil {
		return nil, err
	}

	params := uniswap_binding.IQuoterV2QuoteExactInputSingleParams{}
	params.TokenIn = tokenIn
	params.TokenOut = tokenOut
	params.AmountIn = amountIn
	params.Fee = fee
	params.SqrtPriceLimitX96 = big.NewInt(0)

	var out []interface{}
	rawCaller := &uniswap_binding.UniswapV3QuoterRaw{Contract: quoter}
	err = rawCaller.Call(nil, &out, "quoteExactInputSingle", params)
	if err != nil {
		return nil, err
	}

	// uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate
	qr := &QuoteSingleResult{}
	qr.AmountOut = out[0].(*big.Int)
	qr.SqrtPriceX96After = out[1].(*big.Int)
	qr.InitializedTicksCrossed = out[2].(uint32)
	qr.GasEstimate = out[3].(*big.Int)

	return qr, nil
}

func (s *Quoter) QuoteExactInputForUniswapV3(quoterAddress common.Address, path []byte, amountIn *big.Int, endpoint string) (*QuoteResult, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	quoter, err := uniswap_binding.NewUniswapV3Quoter(quoterAddress, client)
	if err != nil {
		return nil, err
	}

	var outs []interface{}
	rawCaller := &uniswap_binding.UniswapV3QuoterRaw{Contract: quoter}
	err = rawCaller.Call(nil, &outs, "quoteExactInput", path, amountIn)
	if err != nil {
		return nil, err
	}

	// uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate
	qr := &QuoteResult{}
	qr.AmountOut = outs[0].(*big.Int)
	qr.SqrtPriceX96AfterList = outs[1].([]*big.Int)
	qr.InitializedTicksCrossedList = outs[2].([]uint32)
	qr.GasEstimate = outs[3].(*big.Int)

	return qr, nil
}

func (s *Quoter) GetAmountsOutForPancakeswapV2(routerAddress common.Address, path []common.Address, amountIn *big.Int, endpoint string) ([]*big.Int, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	router, err := pancakeswap_binding.NewPancakeswapV2Router(routerAddress, client)
	if err != nil {
		return nil, err
	}

	amountsOut, err := router.GetAmountsOut(nil, amountIn, path)
	if err != nil {
		return nil, err
	}

	return amountsOut, nil
}

func (s *Quoter) QuoteExactInputSingleForPancakeswapV3(quoterAddress common.Address, tokenIn common.Address, tokenOut common.Address,
	amountIn *big.Int, fee *big.Int, endpoint string) (*QuoteSingleResult, error) {

	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	quoter, err := pancakeswap_binding.NewPancakeswapV3Quoter(quoterAddress, client)
	if err != nil {
		return nil, err
	}

	params := pancakeswap_binding.IQuoterV2QuoteExactInputSingleParams{}
	params.TokenIn = tokenIn
	params.TokenOut = tokenOut
	params.AmountIn = amountIn
	params.Fee = fee
	params.SqrtPriceLimitX96 = big.NewInt(0)

	var outs []interface{}
	rawCaller := &pancakeswap_binding.PancakeswapV3QuoterRaw{Contract: quoter}
	err = rawCaller.Call(nil, &outs, "quoteExactInputSingle", params)
	if err != nil {
		return nil, err
	}

	// uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate
	qr := &QuoteSingleResult{}
	qr.AmountOut = outs[0].(*big.Int)
	qr.SqrtPriceX96After = outs[1].(*big.Int)
	qr.InitializedTicksCrossed = outs[2].(uint32)
	qr.GasEstimate = outs[3].(*big.Int)

	return qr, nil
}

func (s *Quoter) QuoteExactInputForPancakeV3(quoterAddress common.Address, path []byte, amountIn *big.Int, endpoint string) (*QuoteResult, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	quoter, err := pancakeswap_binding.NewPancakeswapV3Quoter(quoterAddress, client)
	if err != nil {
		return nil, err
	}

	var outs []interface{}
	rawCaller := &pancakeswap_binding.PancakeswapV3QuoterRaw{Contract: quoter}
	err = rawCaller.Call(nil, &outs, "quoteExactInput", path, amountIn)
	if err != nil {
		return nil, err
	}

	// uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate
	qr := &QuoteResult{}
	qr.AmountOut = outs[0].(*big.Int)
	qr.SqrtPriceX96AfterList = outs[1].([]*big.Int)
	qr.InitializedTicksCrossedList = outs[2].([]uint32)
	qr.GasEstimate = outs[3].(*big.Int)

	return qr, nil
}

func (s *Quoter) GetFeeForUniswapV3(pairAddress common.Address, endpoint string) (int64, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return 0, err
	}
	defer client.Close()

	pair, err := uniswap_binding.NewUniswapV3Pair(pairAddress, client)
	if err != nil {
		return 0, err
	}

	fee, err := pair.Fee(nil)
	if err != nil {
		return 0, err
	}

	return fee.Int64(), nil
}

func (s *Quoter) GetFeeForPancakeswapV3(pairAddress common.Address, endpoint string) (int64, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return 0, err
	}
	defer client.Close()

	pair, err := pancakeswap_binding.NewPancakeswapV3Pair(pairAddress, client)
	if err != nil {
		return 0, err
	}

	fee, err := pair.Fee(nil)
	if err != nil {
		return 0, err
	}

	return fee.Int64(), nil
}