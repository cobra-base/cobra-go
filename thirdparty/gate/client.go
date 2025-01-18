package gate

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/gateio/gateapi-go/v6"
	"strings"
)

type Conf struct {
	ApiKey    string `json:"apiKey"`
	ApiSecret string `json:"apiSecret"`
}

type Client struct {
	conf *Conf
	api  *gateapi.APIClient
	ctx  context.Context
}

var clientInstance = &Client{}

func GetClient() *Client {
	return clientInstance
}

func (s *Client) Init(conf *Conf) error {
	s.conf = conf

	s.api = gateapi.NewAPIClient(gateapi.NewConfiguration())
	s.ctx = context.WithValue(context.Background(), gateapi.ContextGateAPIV4, gateapi.GateAPIV4{
		Key:    conf.ApiKey,
		Secret: conf.ApiSecret,
	})

	return nil
}

// GetCurrencyChains 查询币种支持的链
func (s *Client) GetCurrencyChains(currency string) ([]gateapi.CurrencyChain, error) {
	c, _, err := s.api.WalletApi.ListCurrencyChains(context.Background(), currency)
	return c, err
}

// GetSpotCurrenciesSingle 查询单个币种信息
func (s *Client) GetSpotCurrenciesSingle(currency string) (*gateapi.Currency, error) {
	c, _, err := s.api.SpotApi.GetCurrency(context.Background(), currency)
	if err != nil {
		return nil, err
	}
	return &c, err
}

// GetSpotOrderBook 获取市场深度信息
func (s *Client) GetSpotOrderBook(currencyPair string) (*gateapi.OrderBook, error) {
	opts := &gateapi.ListOrderBookOpts{}
	opts.Limit = optional.NewInt32(50)
	b, _, err := s.api.SpotApi.ListOrderBook(context.Background(), currencyPair, opts)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

// GetSpotAccounts 获取现货交易账户列表
func (s *Client) GetSpotAccounts(currency string) (*gateapi.SpotAccount, error) {
	opts := &gateapi.ListSpotAccountsOpts{}
	opts.Currency = optional.NewString(currency)
	v, _, err := s.api.SpotApi.ListSpotAccounts(s.ctx, opts)
	if err != nil {
		return nil, err
	}
	if len(v) != 1 {
		return nil, fmt.Errorf("spot accounts length except,currency %s,len %d", currency, len(v))
	}
	return &v[0], nil
}

// GetWalletFee 查询个人交易费率
func (s *Client) GetWalletFee(currencyPair string) (*gateapi.TradeFee, error) {
	opts := &gateapi.GetTradeFeeOpts{}
	opts.CurrencyPair = optional.NewString(currencyPair)
	te, _, err := s.api.WalletApi.GetTradeFee(s.ctx, opts)
	if err != nil {
		return nil, err
	}
	return &te, nil
}

// GetWalletDepositAddress 获取币种充值地址
// chain: BSC, ETH
func (s *Client) GetWalletDepositAddress(currency string, chain string) (string, error) {
	da, _, err := s.api.WalletApi.GetDepositAddress(s.ctx, currency)
	if err != nil {
		return "", err
	}
	for _, v := range da.MultichainAddresses {
		if strings.EqualFold(chain, v.Chain) {
			return v.Address, nil
		}
	}
	return "", fmt.Errorf("not found deposit address,currency %s,chain %s", currency, chain)
}

// GetWalletDepositRecords 查询充值记录
func (s *Client) GetWalletDepositRecords(currency string, from int64, to int64, limit int32, offset int32) ([]gateapi.LedgerRecord, error) {
	opts := &gateapi.ListDepositsOpts{}
	opts.Currency = optional.NewString(currency)
	opts.From = optional.NewInt64(from)
	opts.To = optional.NewInt64(to)
	opts.Limit = optional.NewInt32(limit)
	opts.Offset = optional.NewInt32(offset)
	ld, _, err := s.api.WalletApi.ListDeposits(s.ctx, opts)
	if err != nil {
		return nil, err
	}
	return ld, nil
}

// GetWalletWithdrawalRecords 查询提现记录
func (s *Client) GetWalletWithdrawalRecords(currency string, from int64, to int64, limit int32, offset int32) ([]gateapi.WithdrawalRecord, error) {
	opts := &gateapi.ListWithdrawalsOpts{}
	opts.Currency = optional.NewString(currency)
	opts.From = optional.NewInt64(from)
	opts.To = optional.NewInt64(to)
	opts.Limit = optional.NewInt32(limit)
	opts.Offset = optional.NewInt32(offset)
	wr, _, err := s.api.WalletApi.ListWithdrawals(s.ctx, opts)
	if err != nil {
		return nil, err
	}
	return wr, nil
}

// GetWalletWithdrawStatus 查询币种体现状态
func (s *Client) GetWalletWithdrawStatus(currency string) ([]gateapi.WithdrawStatus, error) {
	opts := &gateapi.ListWithdrawStatusOpts{}
	opts.Currency = optional.NewString(currency)
	ws, _, err := s.api.WalletApi.ListWithdrawStatus(s.ctx, opts)
	if err != nil {
		return nil, err
	}
	return ws, nil
}
