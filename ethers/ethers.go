package ethers

import (
    "context"
    "errors"
    "fmt"
    "github.com/cobra-base/cobra-go/ethers/binding"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "math/big"
    "strings"
    "time"
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
        if unit >= len(s) {
            part := strings.Repeat("0", unit-len(s))
            return fmt.Sprintf("0.%s%s", part, s)
        } else {
            part0 := s[0 : len(s)-unit]
            part1 := s[len(s)-unit:]
            return fmt.Sprintf("%s.%s", part0, part1)
        }
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
        return nil, fmt.Errorf("failed to integer part:%s,%s", s, parts[0])
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

func ToReadableAmount(v *big.Int, decimals int, precision int) string {
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
func BalanceAt(address common.Address, endpoint string) (*big.Int, error) {
    client, err := ethclient.Dial(endpoint)
    if err != nil {
        return nil, err
    }
    defer client.Close()

    balance, err := client.BalanceAt(context.Background(), address, nil)
    if err != nil {
        return nil, err
    }

    return balance, nil
}

func GetNameERC20(address common.Address, endpoint string) (string, error) {
    client, err := ethclient.Dial(endpoint)
    if err != nil {
        return "", err
    }
    defer client.Close()

    contract, err := binding.NewERC20(address, client)
    if err != nil {
        return "", err
    }
    name, err := contract.Name(&bind.CallOpts{})
    return name, err
}

func GetSymbolERC20(address common.Address, endpoint string) (string, error) {
    client, err := ethclient.Dial(endpoint)
    if err != nil {
        return "", err
    }
    defer client.Close()

    contract, err := binding.NewERC20(address, client)
    if err != nil {
        return "", err
    }
    name, err := contract.Symbol(&bind.CallOpts{})
    return name, err
}

func GetSupplyERC20(address common.Address, endpoint string) (*big.Int, error) {
    client, err := ethclient.Dial(endpoint)
    if err != nil {
        return nil, err
    }
    defer client.Close()

    contract, err := binding.NewERC20(address, client)
    if err != nil {
        return nil, err
    }
    supply, err := contract.TotalSupply(&bind.CallOpts{})
    return supply, err
}

func GetDecimalsERC20(address common.Address, endpoint string) (uint8, error) {
    client, err := ethclient.Dial(endpoint)
    if err != nil {
        return 0, err
    }
    defer client.Close()

    contract, err := binding.NewERC20(address, client)
    if err != nil {
        return 0, err
    }

    decimals, err := contract.Decimals(&bind.CallOpts{})
    return decimals, err
}

func GetBalanceERC20(owner common.Address, address common.Address, endpoint string) (*big.Int, error) {
    client, err := ethclient.Dial(endpoint)
    if err != nil {
        return nil, err
    }
    defer client.Close()

    contract, err := binding.NewERC20(address, client)
    if err != nil {
        return nil, err
    }
    bal, err := contract.BalanceOf(&bind.CallOpts{}, owner)
    return bal, err
}

func AllowanceERC20(owner common.Address, spender common.Address, token common.Address, endpoint string) (*big.Int, error) {
    client, err := ethclient.Dial(endpoint)
    if err != nil {
        return nil, err
    }
    defer client.Close()

    contract, err := binding.NewERC20(token, client)
    if err != nil {
        return nil, err
    }

    amount, err := contract.Allowance(&bind.CallOpts{}, owner, spender)
    if err != nil {
        return nil, err
    }
    return amount, err
}

type GasFeeParams struct {
    GasTipCap *big.Int
    GasFeeCap *big.Int
    GasLimit  uint64
}

func ApproveERC20(signer *Signer, spender common.Address, token common.Address, amount *big.Int, endpoint string, chainId int64, feeParams *GasFeeParams) (string, error) {
    client, err := ethclient.Dial(endpoint)
    if err != nil {
        return "", err
    }
    defer client.Close()

    contract, err := binding.NewERC20(token, client)
    if err != nil {
        return "", err
    }

    nonce, err := client.PendingNonceAt(context.Background(), signer.address)
    if err != nil {
        return "", fmt.Errorf("pending nonce fail,address %s:%s", signer.address, err.Error())
    }

    signing := types.LatestSignerForChainID(big.NewInt(chainId))
    opts := &bind.TransactOpts{
        From:  signer.address,
        Nonce: big.NewInt(int64(nonce)),
        Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
            return types.SignTx(transaction, signing, signer.privateKey)
        },
        GasTipCap: feeParams.GasTipCap, // EIP 1559
        GasFeeCap: feeParams.GasFeeCap,
        GasLimit:  feeParams.GasLimit,
        Context:   context.Background(),
    }

    tx, err := contract.Approve(opts, spender, amount)
    if err != nil {
        return "", fmt.Errorf("approve fail,owner %s,spender %s,token %s:%s", signer.address, spender.Hex(), token.Hex(), err.Error())
    }

    txHash := tx.Hash().Hex()

    return txHash, nil
}

func PendingApproveERC20Adequate(owner *Signer, spender common.Address, token common.Address, endpoint string, chainId int64, feeData *GasFeeParams) error {

    allowance, err := AllowanceERC20(owner.Address(), spender, token, endpoint)
    if err != nil {
        return err
    }

    if allowance.Cmp(new(big.Int).Div(MaxUint256, big.NewInt(10))) >= 0 {
        return nil
    }

    txHash, err := ApproveERC20(owner, spender, token, MaxUint256, endpoint, chainId, feeData)
    if err != nil {
        return err
    }

    status, _ := WaitTxStatus(txHash, 30*time.Second, 3*time.Second, endpoint)
    if status == TxStatusSuccess {
        return nil
    }

    return fmt.Errorf("approve adequate except,spender %s,tx %s,status %d", spender, txHash, status)
}

func TransferERC20(signer *Signer, toAddress common.Address, tokenAddress common.Address, amount *big.Int, endpoint string, chainId int64, feeParams *GasFeeParams) (string, error) {
    client, err := ethclient.Dial(endpoint)
    if err != nil {
        return "", err
    }
    defer client.Close()

    contract, err := binding.NewERC20(tokenAddress, client)
    if err != nil {
        return "", err
    }

    nonce, err := client.PendingNonceAt(context.Background(), signer.address)
    if err != nil {
        return "", fmt.Errorf("pending nonce fail,address %s:%s", signer.address, err.Error())
    }

    signing := types.LatestSignerForChainID(big.NewInt(chainId))
    opts := &bind.TransactOpts{
        From:  signer.address,
        Nonce: big.NewInt(int64(nonce)),
        Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
            return types.SignTx(transaction, signing, signer.privateKey)
        },
        GasTipCap: feeParams.GasTipCap, // EIP 1559
        GasFeeCap: feeParams.GasFeeCap,
        GasLimit:  feeParams.GasLimit,
        Context:   context.Background(),
    }

    tx, err := contract.Transfer(opts, toAddress, amount)
    if err != nil {
        return "", fmt.Errorf("transfer fail,owner %s,to %s,token %s:%s", signer.address, toAddress.Hex(), tokenAddress.Hex(), err.Error())
    }

    txHash := tx.Hash().Hex()

    return txHash, nil
}

func SendRawTransaction(signer *Signer, toAddress common.Address, txData []byte, endpoint string, chainId int64, feeParams *GasFeeParams) (string, error) {
    client, err := ethclient.Dial(endpoint)
    if err != nil {
        return "", err
    }
    defer client.Close()

    nonce, err := client.PendingNonceAt(context.Background(), signer.address)
    if err != nil {
        return "", fmt.Errorf("pending nonce fail,address %s:%s", signer.address, err.Error())
    }

    rawTx := types.NewTx(&types.DynamicFeeTx{
        ChainID:   big.NewInt(chainId),
        Nonce:     nonce,
        GasTipCap: feeParams.GasTipCap,
        GasFeeCap: feeParams.GasFeeCap,
        Gas:       feeParams.GasLimit,
        To:        &toAddress,
        Value:     big.NewInt(0),
        Data:      txData,
    })
    // signedTx, err := types.SignTx(rawTx, types.NewEIP155Signer(big.NewInt(chainId)), signer.privateKey)
    signedTx, err := types.SignTx(rawTx, types.NewLondonSigner(big.NewInt(chainId)), signer.privateKey)
    if err != nil {
        return "", fmt.Errorf("sign tx fail,address %s:%s", signer.address, err.Error())
    }

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        return "", fmt.Errorf("send tx fail,address %s:%s", signer.address, err.Error())
    }

    txHash := signedTx.Hash().Hex()

    return txHash, nil
}
