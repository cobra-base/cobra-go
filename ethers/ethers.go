package ethers

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
	"strings"
)

func FormatEther(v *big.Int) string {
	return FormatUnits(v, 18)
}

func FormatGWei(v *big.Int) string {
	return FormatUnits(v, 9)
}

func FormatWei(v *big.Int) string {
	return FormatUnits(v, 0)
}

func FormatUnits(v *big.Int, unit int) string {
	f := new(big.Float)
	f.SetInt(v)
	q := new(big.Float).Quo(f, big.NewFloat(math.Pow10(unit)))
	return q.String()
}

func ToFriendlyAmount(v *big.Int, decimals int, precision int) string {
	s := FormatUnits(v, decimals)
	dot := strings.Index(s, ".")
	if dot > 0 {
		if dot+precision+1 < len(s) {
			return s[0 : dot+precision+1]
		}
	}
	return s
}

// IsValidERC20Address 目标地址是否符合ERC20标准
func IsValidERC20Address(address string) bool {
	return common.IsHexAddress(address)
}

// BalanceAt 获取账户ETH余额
func BalanceAt(address string, endpoint string) (*big.Int, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func GetERC20Name(address string, endpoint string) (string, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return "", err
	}
	defer client.Close()

	tokenAddress := common.HexToAddress(address)
	contract, err := NewERC20(tokenAddress, client)
	if err != nil {
		return "", err
	}
	name, err := contract.Name(&bind.CallOpts{})
	return name, err
}

func GetERC20Symbol(address string, endpoint string) (string, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return "", err
	}
	defer client.Close()

	tokenAddress := common.HexToAddress(address)
	contract, err := NewERC20(tokenAddress, client)
	if err != nil {
		return "", err
	}
	name, err := contract.Symbol(&bind.CallOpts{})
	return name, err
}

func GetERC20Decimals(address string, endpoint string) (uint8, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return 0, err
	}
	defer client.Close()

	tokenAddress := common.HexToAddress(address)
	contract, err := NewERC20(tokenAddress, client)
	if err != nil {
		return 0, err
	}
	decimals, err := contract.Decimals(&bind.CallOpts{})
	return decimals, err
}

func GetERC20Balance(owner string, address string, endpoint string) (*big.Int, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ownerAddress := common.HexToAddress(owner)
	tokenAddress := common.HexToAddress(address)
	contract, err := NewERC20(tokenAddress, client)
	if err != nil {
		return nil, err
	}
	bal, err := contract.BalanceOf(&bind.CallOpts{}, ownerAddress)
	return bal, err
}
