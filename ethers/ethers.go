package ethers

import (
	"context"
	"errors"
	"fmt"
	"github.com/cobra-base/cobra-go/ethers/binding"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

var (
	weiPerGO   = big.NewInt(1e18)
	weiPerGwei = big.NewInt(1e9)
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
	if unit < 0 {
		panic(fmt.Sprintf("invalid unit %d", unit))
	}
	s := v.String()
	if unit == 0 {
		return s
	} else {
		part := strings.Repeat("0", unit)
		return fmt.Sprintf("%s%s", s, part)
	}
}

func ParseEther(s string) (*big.Int, error) {
	return ParseUnits(s, 18)
}

func ParseGWei(s string) (*big.Int, error) {
	return ParseUnits(s, 9)
}

func ParseWei(s string) (*big.Int, error) {
	return ParseUnits(s, 0)
}

func ParseUnits(s string, unit int) (*big.Int, error) {
	parts := strings.Split(s, ".")

	whole, ok := new(big.Int).SetString(parts[0], 10)
	if !ok {
		return nil, fmt.Errorf("failed to integer part:%s", whole)
	}
	whole = new(big.Int).Mul(whole, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(unit)), nil))

	if len(parts) == 1 {
		return whole, nil
	}

	if len(parts) > 2 {
		return nil, errors.New("invalid value: more than one decimal point")
	}

	decStr := parts[1]
	if len(decStr) > unit {
		return nil, fmt.Errorf("too many decimal digits %d: limit %d", len(decStr), unit)
	}

	dec, ok := new(big.Int).SetString(decStr+strings.Repeat("0", unit-len(decStr)), 10)
	if !ok {
		return nil, fmt.Errorf("failed to decimal part: %s", decStr)
	}

	return whole.Add(whole, dec), nil
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
	contract, err := binding.NewERC20(tokenAddress, client)
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
	contract, err := binding.NewERC20(tokenAddress, client)
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
	contract, err := binding.NewERC20(tokenAddress, client)
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
	contract, err := binding.NewERC20(tokenAddress, client)
	if err != nil {
		return nil, err
	}
	bal, err := contract.BalanceOf(&bind.CallOpts{}, ownerAddress)
	return bal, err
}
