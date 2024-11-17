package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

// EncryptPKCS1v15 encrypts the given message with the public key using PKCS1v15 padding.
func (p *PublicKey) EncryptPKCS1v15(raw []byte) ([]byte, error) {
	if p.publicKey == nil {
		return nil, fmt.Errorf("public key is nil")
	}
	return rsa.EncryptPKCS1v15(rand.Reader, p.publicKey, raw)
}

// EncryptOAEP encrypts the given message with the public key using OAEP padding.
func (p *PublicKey) EncryptOAEP(raw []byte) ([]byte, error) {
	if p.publicKey == nil {
		return nil, fmt.Errorf("public key is nil")
	}
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, p.publicKey, raw, nil)
}
