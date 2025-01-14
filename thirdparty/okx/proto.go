package okx

type TokenInfo struct {
	Decimal              string `json:"decimal"`
	IsHoneyPot           bool   `json:"isHoneyPot"`
	TokenContractAddress string `json:"tokenContractAddress"`
	TokenSymbol          string `json:"tokenSymbol"`
	TokenUnitPrice       string `json:"tokenUnitPrice"`
}

type QuoteResult struct {
	FromToken             *TokenInfo `json:"fromToken"`
	ToToken               *TokenInfo `json:"toToken"`
	FromTokenAmount       string     `json:"fromTokenAmount"`
	ToTokenAmount         string     `json:"toTokenAmount"`
	PriceImpactPercentage string     `json:"priceImpactPercentage"`
}

type QuoteRsp struct {
	Code string         `json:"code"`
	Data []*QuoteResult `json:"data"`
	Msg  string         `json:"msg"`
}
