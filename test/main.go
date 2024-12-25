package main

import (
    "encoding/json"
    "fmt"
    "github.com/cobra-base/cobra-go/glog"
    dexScreenerApi "github.com/cobra-base/cobra-go/thirdparty"
    "os"
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

    tokenAddress := "0x111111111117dC0aa78b770fA6A738034120C302"
    v, err := dexScreenerApi.GetTokenLiquidity(tokenAddress)
    fmt.Println(err, len(v))
    for _, e := range v {
        i, _ := json.Marshal(e)
        fmt.Println(string(i))
    }

}
