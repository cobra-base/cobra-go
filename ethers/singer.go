package ethers

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
)
import "github.com/ethereum/go-ethereum/common"

type Signer struct {
	hexKey     string
	privateKey *ecdsa.PrivateKey
	address    common.Address
}

func NewSigner(hexPriKey string) (*Signer, error) {
	privateKey, err := crypto.HexToECDSA(hexPriKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	pubKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("public key cant convert ecdsa")
	}

	address := crypto.PubkeyToAddress(*pubKeyECDSA)

	s := &Signer{}
	s.hexKey = hexPriKey
	s.privateKey = privateKey
	s.address = address

	return s, nil
}

func (s *Signer) Address() common.Address {
	return s.address
}
