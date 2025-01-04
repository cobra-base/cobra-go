package ethers

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
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
