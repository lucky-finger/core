package rsa

import (
	"crypto/rand"
	"crypto/rsa"
)

type BitSize int

const (
	Bit1024 BitSize = 1024
	Bit2048 BitSize = 2048
)

// GenerateKeyPair 生成密钥对
func GenerateKeyPair() (*KeyPair, error) {
	return GenerateKeyPairWithBitSize(Bit2048)
}

// GenerateKeyPairWithBitSize 生成密钥对
func GenerateKeyPairWithBitSize(biteSize BitSize) (*KeyPair, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, int(biteSize))
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		privateKey: &PrivateKey{privateKey: privateKey},
		publicKey:  &PublicKey{publicKey: &privateKey.PublicKey},
	}, nil
}
