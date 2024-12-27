package main

import (
	"encoding/json"
	"fmt"
	"github.com/cobra-base/cobra-go/ethers"
	"github.com/cobra-base/cobra-go/glog"
	dexScreenerApi "github.com/cobra-base/cobra-go/thirdparty"
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

func main() {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:7897")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7897")

	endpoint := "https://bnb-mainnet.g.alchemy.com/v2/R_dtWtvB3kAeG5ErH0CMtXi37rtJwhZW"
	address := "0xE62d922D195A853f03211Ce69636D5A90288bFe2"
	v, _ := ethers.BalanceAt(address, endpoint)
	fmt.Println(ethers.FormatEther(v))
	fmt.Println(ethers.FormatGWei(v))
	fmt.Println(ethers.FormatWei(v))
	fmt.Println(ethers.ToFriendlyAmount(v, 18, 4))
	fmt.Println(ethers.ToFriendlyAmount(v, 9, 4))

	t := "0x111111111117dC0aa78b770fA6A738034120C302"
	fmt.Println(ethers.GetERC20Name(t, endpoint))
	fmt.Println(ethers.GetERC20Symbol(t, endpoint))
	fmt.Println(ethers.GetERC20Decimals(t, endpoint))
	fmt.Println(ethers.GetERC20Balance("0x111111111117dC0aa78b770fA6A738034120C302", t, endpoint))
}
