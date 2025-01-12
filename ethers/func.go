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

func ParsePath(path []byte) ([]common.Address, []int, error) {
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

func BuildPath(addresses []common.Address, fees []int) ([]byte, error) {
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

func GetTxStatus(txHash string, endpoint string) TxStatus {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		glog.Warnw("ethclient dail except", "endpoint", endpoint)
		return TxStatusUnknown
	}
	defer client.Close()
	return GetTxStatusWithClient(txHash, client)
}

func GetTxStatusWithClient(txHash string, client *ethclient.Client) TxStatus {
	_, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return TxStatusNotFound
		}
		glog.Warnw("transaction by hash except", "txHash", txHash, "err", err)
		return TxStatusUnknown
	}

	if isPending {
		return TxStatusPending
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return TxStatusNotFound
		}
		glog.Warnw("transaction receipt except", "txHash", txHash, "err", err)
		return TxStatusUnknown
	}

	if receipt.Status == types.ReceiptStatusFailed {
		return TxStatusFail
	} else if receipt.Status == types.ReceiptStatusSuccessful {
		return TxStatusSuccess
	} else {
		glog.Warnw("unrecognized transaction receipt status", "txHash", txHash, "status", receipt.Status)
		return TxStatusUnknown
	}
}

func WaitForTxStatus(txHash string, expired time.Duration, interval time.Duration, client *ethclient.Client) TxStatus {
	txStatus := TxStatusUnknown
	startTime := time.Now()
	for {
		time.Sleep(interval)
		txStatus = GetTxStatusWithClient(txHash, client)
		if txStatus == TxStatusSuccess || txStatus == TxStatusFail {
			break
		}
		if time.Now().Sub(startTime).Milliseconds() > expired.Milliseconds() {
			break
		}
	}
	return txStatus
}
