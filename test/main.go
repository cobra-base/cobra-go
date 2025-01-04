package main

import (
	"encoding/json"
	"fmt"
	"github.com/cobra-base/cobra-go/ethers"
	"github.com/cobra-base/cobra-go/glog"
	dexScreenerApi "github.com/cobra-base/cobra-go/thirdparty/dexscreener"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"os"
)

func main2() {
	fmt.Println("Start")

	logConf := &glog.Config{}
	logConf.LogDir = "."
	logConf.LogLevel = "debug"
	logConf.LogName = "test"
	logConf.WriteConsole = true
	glog.Init(logConf)

	os.Setenv("HTTP_PROXY", "http://127.0.0.1:7897")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7897")

	tokenAddress := "0x111111111117dC0aa78b770fA6A738034120C302"
	v, err := dexScreenerApi.GetTokenLiquidity(tokenAddress)
	fmt.Println(err, len(v))
	for _, e := range v {
		i, _ := json.Marshal(e)
		fmt.Println(string(i))
	}

}

func main_uniwap_v2() {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:7897")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7897")

	endpoint := "https://bnb-mainnet.g.alchemy.com/v2/R_dtWtvB3kAeG5ErH0CMtXi37rtJwhZW"

	routerAddress := common.HexToAddress(ethers.Chains["bsc"].UniswapV2RouterAddress)

	tokenIn := common.HexToAddress(ethers.Chains["bsc"].WETHAddress)
	tokenOut := common.HexToAddress(ethers.Chains["bsc"].USDTAddress)

	exp := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)
	amountIn := big.NewInt(0).Mul(big.NewInt(1), exp)

	quoter := ethers.GetQuoter()
	v, e := quoter.GetAmountsOutForUniswapV2(routerAddress, []common.Address{tokenIn, tokenOut}, amountIn, endpoint)

	fmt.Println(e)
	fmt.Println(v)
}

func main_uniwap_v3() {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:7897")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7897")

	endpoint := "https://bnb-mainnet.g.alchemy.com/v2/R_dtWtvB3kAeG5ErH0CMtXi37rtJwhZW"

	routerAddress := ethers.Chains["bsc"].UniswapV3QuoterAddress

	tokenIn := common.HexToAddress(ethers.Chains["bsc"].WETHAddress)
	tokenOut := common.HexToAddress(ethers.Chains["bsc"].USDTAddress)
	fee := big.NewInt(100)

	exp := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)
	amountIn := big.NewInt(0).Mul(big.NewInt(1), exp)

	quoter := ethers.GetQuoter()
	v, e := quoter.QuoteExactInputSingleForUniswapV3(common.HexToAddress(routerAddress), tokenIn, tokenOut, amountIn, fee, endpoint)

	fmt.Println(e)
	fmt.Println(v)
	fmt.Println(ethers.FormatEther(v.AmountOut))
}

func main() {
	i := big.NewInt(3520430602)
	s := ethers.FormatUnits(i, 18)
	fmt.Println(s)
}
