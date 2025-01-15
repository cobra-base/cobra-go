package gate

import (
    "context"
    "github.com/antihax/optional"
    "github.com/gateio/gateapi-go/v6"
)

type Conf struct {
    ApiKey    string `json:"apiKey"`
    ApiSecret string `json:"apiSecret"`
}

type Client struct {
    conf       *Conf
    gateClient *gateapi.APIClient
    gateCtx    context.Context
}

var clientInstance = &Client{}

func GetClient() *Client {
    return clientInstance
}

func (s *Client) Init(conf *Conf) error {
    s.conf = conf

    s.gateClient = gateapi.NewAPIClient(gateapi.NewConfiguration())
    s.gateCtx = context.WithValue(context.Background(), gateapi.ContextGateAPIV4, gateapi.GateAPIV4{
        Key:    conf.ApiKey,
        Secret: conf.ApiSecret,
    })

    return nil
}

// GetCurrencyChains 查询币种支持的链
func (s *Client) GetCurrencyChains(currency string) ([]gateapi.CurrencyChain, error) {
    c, _, err := s.gateClient.WalletApi.ListCurrencyChains(context.Background(), currency)
    return c, err
}

// GetSpotCurrenciesSingle 查询单个币种信息
func (s *Client) GetSpotCurrenciesSingle(currency string) (*gateapi.Currency, error) {
    c, _, err := s.gateClient.SpotApi.GetCurrency(context.Background(), currency)
    if err != nil {
        return nil, err
    }
    return &c, err
}

// GetSpotOrderBook 获取市场深度信息
func (s *Client) GetSpotOrderBook(currencyPair string) (*gateapi.OrderBook, error) {
    opts := &gateapi.ListOrderBookOpts{}
    opts.Limit = optional.NewInt32(50)
    b, _, err := s.gateClient.SpotApi.ListOrderBook(context.Background(), currencyPair, opts)
    if err != nil {
        return nil, err
    }
    return &b, nil
}
