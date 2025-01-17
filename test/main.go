package main

import (
    "encoding/json"
    "fmt"
    "github.com/cobra-base/cobra-go/ethers"
    "github.com/cobra-base/cobra-go/glog"
    dexScreenerApi "github.com/cobra-base/cobra-go/thirdparty/dexscreener"
    "github.com/cobra-base/cobra-go/thirdparty/gate"
    "github.com/cobra-base/cobra-go/thirdparty/okx"
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

    endpoint := "https://bnb-mainnet.g.alchemy.com/v2/xxxxx"

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

    endpoint := "https://bnb-mainnet.g.alchemy.com/v2/xxxx"

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
    os.Setenv("HTTP_PROXY", "http://127.0.0.1:7897")
    os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7897")

    logConf := &glog.Config{}
    logConf.LogName = "test"
    glog.Init(logConf)

    conf := &okx.Conf{}
    conf.ApiKey = "xxxx"
    conf.SecretKey = "xxxxxx"
    conf.Passphrase = "xxxxxx"
    conf.ProjectId = "xxxxx"

    user := "0xE62d922D195A853f03211Ce69636D5A90288bFe2"
    usdt := "0x55d398326f99059fF775485246999027B3197955"

    // chainId := 56

    aggregator := okx.GetAggregator()
    aggregator.Init(conf)

    // token := "0xadcdbcb0db9edf31509971f64f0a8e0fc53b384d"
    // decimals := 18
    token := "0xc748673057861a797275CD8A068AbB95A902e8de"

    endpoint := "https://bnb-mainnet.g.alchemy.com/v2/xxxxx"
    v, e := ethers.AllowanceForERC20(common.HexToAddress(user), common.HexToAddress(token), common.HexToAddress(usdt), endpoint)
    fmt.Println(e, v)

    ////----
    signer, _ := ethers.NewSigner("xxxx")
    ////----
    txHash, e := ethers.ApproveForERC20(signer, common.HexToAddress("0x9Dd569c36610aF0b02c5D44b2427b922552E84c9"), common.HexToAddress(token), big.NewInt(1000),
        endpoint, 56, big.NewInt(1500000000), big.NewInt(2000000000))
    fmt.Println(e, txHash)
    /*
    	v, _ := ethers.ParseUnits("2000", 18)
    	// r, _ := aggregator.Quote(chainId, usdt, token, v)
    	cd, _ := aggregator.GenerateErc20SwapCallData(chainId, usdt, token, user, v, 0.01, okx.GasLevelSlow)
    	//fmt.Println(e)
    	// s, _ := json.Marshal(cd)
    	fmt.Println(cd.Tx.Gas, cd.Tx.GasPrice, cd.Tx.MaxPriorityFeePerGas)

    	cd, _ = aggregator.GenerateErc20SwapCallData(chainId, usdt, token, user, v, 0.01, okx.GasLevelAverage)
    	fmt.Println(cd.Tx.Gas, cd.Tx.GasPrice, cd.Tx.MaxPriorityFeePerGas)

    	cd, _ = aggregator.GenerateErc20SwapCallData(chainId, usdt, token, user, v, 0.01, okx.GasLevelFast)
    	fmt.Println(cd.Tx.Gas, cd.Tx.GasPrice, cd.Tx.MaxPriorityFeePerGas)
    */
    // fmt.Printf("%v", r)
    /*
       bp, sp, err := aggregator.QuoteUsdtPrice(chainId, token, usdt, 9, 18, "5000")
       fmt.Println(bp, sp, err)

       bp, sp, err = aggregator.QuoteUsdtPrice(chainId, token, usdt, 9, 18, "10000")

       fmt.Println(bp, sp, err)
    */
}

func main333() {
    os.Setenv("HTTP_PROXY", "http://127.0.0.1:7897")
    os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7897")

    conf := &gate.Conf{
        ApiKey:    "xxxx",
        ApiSecret: "xxxx",
    }
    gate.GetClient().Init(conf)
    b, e := gate.GetClient().GetSpotOrderBook("BABYDOGE_USDT")

    fmt.Println(e)
    i, _ := json.Marshal(b)
    fmt.Println(i)

}
