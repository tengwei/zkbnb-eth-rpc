package _rpc

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/bnb-chain/zkbnb-eth-rpc/_utils"
)

type AuthClient struct {
	PrivateKey *ecdsa.PrivateKey `json:"private_key"`
	PublicKey  *ecdsa.PublicKey  `json:"public_key"`
	Address    common.Address    `json:"address"`
	ChainId    *big.Int          `json:"chain_id"`
}

func NewAuthClient(priKey string, chainId *big.Int) (authCli *AuthClient, err error) {
	// validate private key
	if !_utils.IsValidPrivateKey(priKey) {
		return nil, ErrInvalidPrivateKey
	}
	// transfer private key str to private key
	privateKey, err := _utils.DecodePrivateKey(priKey)
	if err != nil {
		return nil, err
	}
	// get public key
	publicKey := privateKey.PublicKey
	address := crypto.PubkeyToAddress(publicKey)
	return &AuthClient{PrivateKey: privateKey, PublicKey: &publicKey, Address: address, ChainId: chainId}, nil
}
