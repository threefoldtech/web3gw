package goethclient

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateKeypair() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

func KeyFromSecret(secret string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(secret)
}
