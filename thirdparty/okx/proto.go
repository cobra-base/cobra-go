package okx

type GasLevel string

const GasLevelSlow = "slow"
const GasLevelAverage = "average"
const GasLevelFast = "fast"

type ApproveTransactionResult struct {
	Data               string `json:"data"`
	DexContractAddress string `json:"dexContractAddress"`
	GasLimit           string `json:"gasLimit"`
	GasPrice           string `json:"gasPrice"`
}

type ApproveTransactionRsp struct {
	Code string                      `json:"code"`
	Data []*ApproveTransactionResult `json:"data"`
	Msg  string                      `json:"msg"`
}

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

type SwapTx struct {
	Data                 string `json:"data"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	Gas                  string `json:"gas"`
	GasPrice             string `json:"gasPrice"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"` // EIP-1559:每单位 gas 优先费用的推荐值， BSC未空值
	Slippage             string `json:"slippage"`
	Value                string `json:"value"`
}

type SwapCallDataResult struct {
	Tx *SwapTx `json:"tx"`
}

type SwapCallDataRsp struct {
	Code string                `json:"code"`
	Data []*SwapCallDataResult `json:"data"`
	Msg  string                `json:"msg"`
}
