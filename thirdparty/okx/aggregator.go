package okx

import (
    "crypto/hmac"
    "crypto/sha256"
    "crypto/tls"
    "encoding/base64"
    "encoding/json"
    "errors"
    "fmt"
    "github.com/cobra-base/cobra-go/ethers"
    "github.com/cobra-base/cobra-go/utils"
    "io"
    "math/big"
    "net/http"
    "net/url"
    "strings"
    "time"
)

const okxHost = "https://www.okx.com"

type Conf struct {
    ApiKey     string `json:"apiKey"`
    SecretKey  string `json:"secretKey"`
    Passphrase string `json:"passphrase"`
    ProjectId  string `json:"projectId"`
}

type Aggregator struct {
    conf *Conf
}

var aggregatorInstance = &Aggregator{}

func GetAggregator() *Aggregator {
    return aggregatorInstance
}

func (s *Aggregator) Init(conf *Conf) error {
    s.conf = conf
    return nil
}

func (s *Aggregator) preHash(timestamp string, method string, requestPath string, params map[string]string) string {

    values := url.Values{}
    for k, v := range params {
        values.Add(k, v)
    }

    var queryString string

    if method == "GET" {
        queryString = "?" + values.Encode()
    } else if method == "POST" {
        queryString = values.Encode()
    }

    h := fmt.Sprintf("%s%s%s%s", timestamp, method, requestPath, queryString)

    return h
}

func (s *Aggregator) sign(message string, secretKey string) string {
    hmac := hmac.New(sha256.New, []byte(secretKey))
    hmac.Write([]byte(message))
    byteData := hmac.Sum(nil)
    signature := base64.StdEncoding.EncodeToString(byteData)
    return signature
}

func (s *Aggregator) createSignature(method string, requestPath string, params map[string]string) (string, string) {
    timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
    message := s.preHash(timestamp, method, requestPath, params)
    signature := s.sign(message, s.conf.SecretKey)
    return signature, timestamp
}

func (s *Aggregator) GetReq(reqPath string, params map[string]string) ([]byte, error) {
    signature, timestamp := s.createSignature("GET", reqPath, params)

    webUrl := okxHost + reqPath
    if len(params) > 0 {
        values := url.Values{}
        for k, v := range params {
            values.Add(k, v)
        }
        webUrl = webUrl + "?" + values.Encode()
    }

    req, err := http.NewRequest("GET", webUrl, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("OK-ACCESS-KEY", s.conf.ApiKey)
    req.Header.Set("OK-ACCESS-PASSPHRASE", s.conf.Passphrase)
    req.Header.Set("OK-ACCESS-PROJECT", s.conf.ProjectId)
    req.Header.Set("OK-ACCESS-SIGN", signature)
    req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

    client := &http.Client{
        Transport: tr,
        Timeout:   time.Second * 5,
    }

    stopWatch := utils.NewStopWatch()
    rsp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer rsp.Body.Close()

    stopWatch.Stop()
    // spendTime := stopWatch.Duration()

    statusCode := rsp.StatusCode
    statusMessage := rsp.Status

    // glog.Debugw("ok aggregator get", "url", webUrl, "statusCode", statusCode, "statusMessage", statusMessage, "spendTime(Ms)", spendTime)

    // v, _ := json.Marshal(rsp)
    // fmt.Println(v)

    if statusCode != 200 {
        return nil, errors.New(fmt.Sprintf("%s,%s", statusMessage, webUrl))
    }

    body, err := io.ReadAll(rsp.Body)
    if err != nil {
        return nil, err
    }

    return body, nil
}

func (s *Aggregator) GetApproveTransaction(chainId int64, tokenContractAddress string, approveAmount *big.Int) (*ApproveTransactionResult, error) {
    requestPath := "/api/v5/dex/aggregator/approve-transaction"
    params := map[string]string{
        "chainId":              fmt.Sprintf("%d", chainId),
        "tokenContractAddress": tokenContractAddress,
        "approveAmount":        ethers.FormatWei(approveAmount),
    }
    data, err := s.GetReq(requestPath, params)
    if err != nil {
        return nil, err
    }

    rsp := &ApproveTransactionRsp{}
    err = json.Unmarshal(data, rsp)
    if err != nil {
        return nil, err
    }

    if !strings.EqualFold(rsp.Code, "0") {
        return nil, fmt.Errorf("approve transaction code except,code %s,msg %s", rsp.Code, rsp.Msg)
    }

    if len(rsp.Code) != 1 {
        return nil, fmt.Errorf("approve transaction data except,len %d", len(rsp.Code))
    }

    ar := rsp.Data[0]

    return ar, err
}

func (s *Aggregator) Quote(chainId int64, fromTokenAddress string, toTokenAddress string, fromTokenAmount *big.Int) (*QuoteResult, error) {
    requestPath := "/api/v5/dex/aggregator/quote"
    fromAmount := ethers.FormatWei(fromTokenAmount)
    params := map[string]string{
        "amount":           fromAmount,
        "chainId":          fmt.Sprintf("%d", chainId),
        "fromTokenAddress": fromTokenAddress,
        "toTokenAddress":   toTokenAddress,
    }

    data, err := s.GetReq(requestPath, params)
    if err != nil {
        return nil, err
    }
    ////----
    // fmt.Println(string(data))
    // return nil, nil
    ////----
    rsp := &QuoteRsp{}
    err = json.Unmarshal(data, rsp)
    if err != nil {
        return nil, err
    }

    if !strings.EqualFold(rsp.Code, "0") {
        return nil, fmt.Errorf("quote code except,code %s,msg %s", rsp.Code, rsp.Msg)
    }

    if len(rsp.Code) != 1 {
        return nil, fmt.Errorf("quote data except,len %d", len(rsp.Code))
    }

    qr := rsp.Data[0]

    return qr, err
}

func (s *Aggregator) QuoteUsdtPrice(chainId int64, tokenAddress string, usdtAddress string, tokenDecimals int, usdtDecimals int, usdtAmountIn string) (float64, float64, error) {

    // 查询买价

    usdtAmount, _ := ethers.ParseUnits(usdtAmountIn, usdtDecimals)

    qr, err := s.Quote(chainId, usdtAddress, tokenAddress, usdtAmount)
    if err != nil {
        return 0, 0, err
    }

    tokenAmount, err := ethers.ParseWei(qr.ToTokenAmount)
    if err != nil {
        return 0, 0, err
    }

    priceOfWei := new(big.Int).Div(new(big.Int).Mul(usdtAmount, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(tokenDecimals)), nil)), tokenAmount)

    m := big.NewFloat(0).SetInt(priceOfWei)
    n := big.NewFloat(0).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(usdtDecimals)), nil))
    buyPrice, _ := big.NewFloat(0).Quo(m, n).Float64()

    // 查询卖价

    qr, err = s.Quote(chainId, tokenAddress, usdtAddress, tokenAmount)
    if err != nil {
        return 0, 0, err
    }

    usdtAmount, err = ethers.ParseWei(qr.ToTokenAmount)
    if err != nil {
        return 0, 0, err
    }

    priceOfWei = new(big.Int).Div(new(big.Int).Mul(usdtAmount, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(tokenDecimals)), nil)), tokenAmount)

    m = big.NewFloat(0).SetInt(priceOfWei)
    n = big.NewFloat(0).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(usdtDecimals)), nil))
    sellPrice, _ := big.NewFloat(0).Quo(m, n).Float64()

    return buyPrice, sellPrice, nil
}

// GenerateErc20SwapCallData 生成代币交换调用数据
// slippage 滑点最小值为 0，最大值为 1 （如：0.005代表这笔交易的最大滑点为0.5%，1代表这笔交易的最大滑点为 100%））
func (s *Aggregator) GenerateErc20SwapCallData(chainId int64, fromTokenAddress string, toTokenAddress string, userWalletAddress string,
    fromTokenAmount *big.Int, slippage float64, gasLevel GasLevel) (*SwapCallDataResult, error) {
    params := map[string]string{
        "chainId":           fmt.Sprintf("%d", chainId),
        "amount":            ethers.FormatWei(fromTokenAmount),
        "fromTokenAddress":  fromTokenAddress,
        "toTokenAddress":    toTokenAddress,
        "userWalletAddress": userWalletAddress,
        "slippage":          fmt.Sprintf("%.4f", slippage),
        "gasLevel":          string(gasLevel),
    }
    requestPath := "/api/v5/dex/aggregator/swap"
    data, err := s.GetReq(requestPath, params)
    if err != nil {
        return nil, err
    }

    rsp := &SwapCallDataRsp{}
    err = json.Unmarshal(data, rsp)
    if err != nil {
        return nil, err
    }

    if !strings.EqualFold(rsp.Code, "0") {
        return nil, fmt.Errorf("generate swap data except,code %s,msg %s", rsp.Code, rsp.Msg)
    }

    if len(rsp.Code) != 1 {
        return nil, fmt.Errorf("generate swap data except,len %d", len(rsp.Code))
    }

    sr := rsp.Data[0]

    return sr, err
}
