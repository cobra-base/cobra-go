package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cobra-base/cobra-go/ethers"
	"github.com/cobra-base/cobra-go/glog"
	"github.com/cobra-base/cobra-go/thirdparty/gate"
	"github.com/cobra-base/cobra-go/thirdparty/okx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
	"time"
)

func main() {
	fmt.Println("Start")

	logConf := &glog.Config{}
	logConf.LogDir = "."
	logConf.LogLevel = "debug"
	logConf.LogName = "test"
	logConf.WriteConsole = true
	glog.Init(logConf)

	os.Setenv("HTTP_PROXY", "http://127.0.0.1:7897")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7897")

	i := big.NewInt(123456)
	fmt.Println(ethers.FormatUnits(i, 0))
	fmt.Println(ethers.FormatUnits(i, 1))
	fmt.Println(ethers.FormatUnits(i, 5))
	fmt.Println(ethers.FormatUnits(i, 10))
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

func main5555() {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:7897")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7897")

	logConf := &glog.Config{}
	logConf.LogName = "test"
	glog.Init(logConf)

	conf := &okx.Conf{}
	conf.ApiKey = "e5be2397-669c-47f7-9ce0-722030f89854"
	conf.SecretKey = "A1BE68AF3925CFFE531FFBBDDFA24605"
	conf.Passphrase = "bUwac#?3H8CyKuq10cA"
	conf.ProjectId = "c002833d9ee1950b0939324f365e8c52"

	// wallet := "0x9C79073774f97F1C2406d5c1eC7CDC08d5240d19"
	// usdt := "0x55d398326f99059fF775485246999027B3197955"
	// token := "0xc748673057861a797275CD8A068AbB95A902e8de"

	// chainId := int64(56)

	aggregator := okx.GetAggregator()
	aggregator.Init(conf)

	endpoint := "https://bnb-mainnet.g.alchemy.com/v2/R_dtWtvB3kAeG5ErH0CMtXi37rtJwhZW"
	client, _ := ethclient.Dial(endpoint)

	t, _ := client.SuggestGasPrice(context.Background())
	fmt.Println(t)

	t, _ = client.SuggestGasTipCap(context.Background())
	fmt.Println(t)

}

func main3333() {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:7897")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7897")

	conf := &gate.Conf{
		ApiKey:    "710e192cca360ea2bddeea9dbdcc8e05",                                 ////----
		ApiSecret: "90cc594b504dc7ac0efe368c21ea793ca3f42483a33ade6bf9b31ebf1b7e53a2", ////----
	}
	gate.GetClient().Init(conf)

	// b, e := gate.GetClient().CreateSpotOrder("MARS_USDT", "buy", "3.0174795")
	// b, e := gate.GetClient().Withdraw(common.HexToAddress("0x174AB062927df9E3949B0283FEbd98FcB29552Fe"), "USDT", "1", "BSC")
	// s, _ := json.Marshal(b)
	// fmt.Println("*****11", e, string(s))

	status, e := gate.GetClient().GetWalletDepositAddress("BABYDOGE", "BSC")
	// a, e := gate.GetClient().GetWalletDepositAddress("USDT", "BSC")
	s, _ := json.Marshal(status)
	fmt.Println("****1", string(s), e)

	return
	to := time.Now().Unix()
	from := to - 60*60*24
	for {
		time.Sleep(2 * time.Second)
		// r, e := gate.GetClient().GetWalletWithdrawalRecords("USDT", from, to, 10, 0)
		r, e := gate.GetClient().GetWalletDepositRecords("USDT", from, to, 10, 0)
		s, _ := json.Marshal(r)
		fmt.Println("*********22", time.Now().String(), e, string(s))
	}

}
