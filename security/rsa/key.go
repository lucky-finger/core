package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type PublicKey struct {
	// publicKey 公钥
	publicKey *rsa.PublicKey
}

func (p *PublicKey) ToPemMust() []byte {
	r, _ := p.ToPem()
	return r
}

func (p *PublicKey) ToPemStringMust() string {
	r, _ := p.ToPemString()
	return r
}

func (p *PublicKey) ToPem() ([]byte, error) {
	if p.publicKey == nil {
		return nil, errors.New("public key is nil")
	}

	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(p.publicKey),
	}), nil
}

func (p *PublicKey) ToPemString() (string, error) {
	r, err := p.ToPem()
	if err != nil {
		return "", err
	}

	return string(r), nil
}

type PrivateKey struct {
	// privateKey 私钥
	privateKey *rsa.PrivateKey
}

func (p *PrivateKey) ToPemMust() []byte {
	r, _ := p.ToPem()
	return r
}

func (p *PrivateKey) ToPemStringMust() string {
	r, _ := p.ToPemString()
	return r
}

func (p *PrivateKey) ToPem() ([]byte, error) {
	if p.privateKey == nil {
		return nil, errors.New("private key is nil")
	}

	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(p.privateKey),
	}), nil
}

func (p *PrivateKey) ToPemString() (string, error) {
	r, err := p.ToPem()
	if err != nil {
		return "", err
	}

	return string(r), nil
}

// KeyPair rsa密钥对
type KeyPair struct {
	// publicKey 公钥
	publicKey *PublicKey
	// privateKey 私钥
	privateKey *PrivateKey
}

func (k *KeyPair) PrivateKey() IKeyToPem {
	return k.privateKey
}

func (k *KeyPair) PublicKey() IKeyToPem {
	return k.publicKey
}
