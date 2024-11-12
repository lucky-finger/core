package rsa

import (
	"crypto/rand"
	"crypto/rsa"
)

// GenerateKeyPair 生成rsa密钥对
func GenerateKeyPair(bits int) (*KeyPair, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		privateKey: &PrivateKey{privateKey: privateKey},
		publicKey:  &PublicKey{publicKey: &privateKey.PublicKey},
	}, nil
}
