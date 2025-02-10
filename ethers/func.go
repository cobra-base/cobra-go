package ethers

import (
	"context"
	"fmt"
	"github.com/cobra-base/cobra-go/glog"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"time"
)

type TxStatus int

const (
	TxStatusUnknown  = TxStatus(0)
	TxStatusSuccess  = TxStatus(1)
	TxStatusFail     = TxStatus(2)
	TxStatusPending  = TxStatus(3)
	TxStatusNotFound = TxStatus(4)
)

func ParsePathForUniswapV3(path []byte) ([]common.Address, []int, error) {
	var addresses []common.Address
	var fees []int

	for i := 0; i < len(path); i += 23 {
		// 从 path 中解析出代币地址
		addresses = append(addresses, common.BytesToAddress(path[i:i+20]))

		// 如果这不是最后一个地址，则还有一个费用字段
		if i+23 < len(path) {
			// 从 path 中解析出费用
			fee := new(big.Int).SetBytes(path[i+20 : i+23]).Int64()
			fees = append(fees, int(fee))
		}
	}

	return addresses, fees, nil
}

func BuildPathForUniswapV3(addresses []common.Address, fees []int) ([]byte, error) {
	var path []byte
	for i, address := range addresses {
		// 将地址添加到路径中
		path = append(path, address.Bytes()...)

		// 如果这不是最后一个地址，则还有一个费用字段
		if i < len(fees) {
			// 将费用转换为字节数组并添加到路径中
			feeBytes := big.NewInt(int64(fees[i])).Bytes()
			fmt.Println(len(feeBytes))
			// 如果费用字节数组的长度小于3，我们需要在前面添加零字节
			if len(feeBytes) < 3 {
				feeBytes = append([]byte{0}, feeBytes...)
			}
			path = append(path, feeBytes...)
		}
	}

	return path, nil
}

func GetReadableTxStatus(status TxStatus) string {
	switch status {
	case TxStatusUnknown:
		return "Unknown"
	case TxStatusSuccess:
		return "Success"
	case TxStatusFail:
		return "Fail"
	case TxStatusPending:
		return "Pending"
	case TxStatusNotFound:
		return "NotFound"
	default:
		return fmt.Sprintf("Invalid_%d", status)
	}
}

func GetRecentBlockNumber(endpoint string) (int64, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		glog.Warnw("ethclient dail except", "endpoint", endpoint)
		return 0, err
	}
	defer client.Close()
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}
	return int64(blockNumber), nil
}

func GetTxStatus(txHash string, endpoint string) (TxStatus, int64) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		glog.Warnw("ethclient dail except", "endpoint", endpoint)
		return TxStatusUnknown, 0
	}
	defer client.Close()
	return GetTxStatusWithClient(txHash, client)
}

func GetTxStatusWithClient(txHash string, client *ethclient.Client) (TxStatus, int64) {
	_, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return TxStatusNotFound, 0
		}
		glog.Warnw("transaction by hash except", "txHash", txHash, "err", err)
		return TxStatusUnknown, 0
	}

	if isPending {
		return TxStatusPending, 0
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return TxStatusNotFound, 0
		}
		glog.Warnw("transaction receipt except", "txHash", txHash, "err", err)
		return TxStatusUnknown, 0
	}

	if receipt.Status == types.ReceiptStatusFailed {
		return TxStatusFail, 0
	} else if receipt.Status == types.ReceiptStatusSuccessful {
		return TxStatusSuccess, receipt.BlockNumber.Int64()
	} else {
		glog.Warnw("unrecognized transaction receipt status", "txHash", txHash, "status", receipt.Status)
		return TxStatusUnknown, 0
	}
}

func WaitTxStatus(txHash string, expired time.Duration, interval time.Duration, endpoint string) (TxStatus, int64) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		glog.Warnw("ethclient dail except", "endpoint", endpoint)
		return TxStatusUnknown, 0
	}
	defer client.Close()
	return WaitTxStatusWithClient(txHash, expired, interval, client)
}

func WaitTxStatusWithClient(txHash string, expired time.Duration, interval time.Duration, client *ethclient.Client) (TxStatus, int64) {
	txStatus := TxStatusUnknown
	blockNumber := int64(0)
	startTime := time.Now()
	for {
		time.Sleep(interval)
		txStatus, blockNumber = GetTxStatusWithClient(txHash, client)
		if txStatus == TxStatusSuccess || txStatus == TxStatusFail {
			break
		}
		if time.Now().Sub(startTime).Milliseconds() > expired.Milliseconds() {
			break
		}
	}
	return txStatus, blockNumber
}
