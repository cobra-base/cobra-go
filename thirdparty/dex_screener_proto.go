package dexScreenerApi

type TokenMeta struct {
    Address string `json:"address"`
    Name    string `json:"name"`
    Symbol  string `json:"symbol"`
}

type NumberOfTx struct {
    Buys  int64 `json:"buys"`
    Sells int64 `json:"sells"`
}

type TimeTx struct {
    M5  *NumberOfTx `json:"m5"`
    H1  *NumberOfTx `json:"h1"`
    H6  *NumberOfTx `json:"h6"`
    H24 *NumberOfTx `json:"h24"`
}

type TimeVolume struct {
    M5  float64 `json:"m1"`
    H1  float64 `json:"h1"`
    H6  float64 `json:"h6"`
    H24 float64 `json:"h24"`
}

type TimePriceChange struct {
    M5  float64 `json:"m1"`
    H1  float64 `json:"h1"`
    H6  float64 `json:"h6"`
    H24 float64 `json:"h24"`
}

type Liquidity struct {
    Usd   float64 `json:"usd"`
    Base  float64 `json:"base"`
    Quote float64 `json:"quote"`
}

type TokenPairs struct {
    ChainId      string           `json:"chainId"` // e.g. ethereum
    DexId        string           `json:"dexId"`   // e.g. uniswap, pancakeswap
    PairAddress  string           `json:"pairAddress"`
    BaseToken    *TokenMeta       `json:"baseToken"`
    QuoteToken   *TokenMeta       `json:"quoteToken"`
    PriceNative  string           `json:"priceNative"`
    PriceUsd     string           `json:"priceUsd"`
    Txns         *TimeTx          `json:"txns"`
    Volume       *TimeVolume      `json:"volume"`
    PriceChange  *TimePriceChange `json:"priceChange"`
    Liquidity    *Liquidity       `json:"liquidity"`
    Fdv          float64          `json:"fdv"`
    MarketCap    float64          `json:"marketCap"`
    PairCreateAt float64          `json:"pairCreateAt"`
}

type TokenLiquidityRsp struct {
    SchemaVersion string        `json:"schemaVersion"`
    Pairs         []*TokenPairs `json:"pairs"`
}
