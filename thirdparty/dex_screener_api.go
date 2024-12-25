package dexScreenerApi

import (
    "encoding/json"
    "github.com/cobra-base/cobra-go/utils"
)

// GetTokenLiquidity
// Get one or multiple pairs by token address (rate-limit 300 requests per minute)
// https://api.dexscreener.com/latest/dex/tokens/{tokenAddresses}
func GetTokenLiquidity(address string) ([]*TokenPairs, error) {
    webUrl := "https://api.dexscreener.com/latest/dex/tokens/" + address
    data, err := utils.HttpGet(webUrl)
    if err != nil {
        return nil, err
    }
    rsp := &TokenLiquidityRsp{}
    err = json.Unmarshal(data, rsp)
    if err != nil {
        return nil, err
    }
    return rsp.Pairs, nil
}
