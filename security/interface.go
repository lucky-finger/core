package security

// IKeyToPem 密钥转pem
type IKeyToPem interface {
	// ToPem 密钥转pem
	ToPem() ([]byte, error)
	// ToPemString 密钥转pem字符串
	ToPemString() (string, error)
}

// IKeyPairToPem 密钥对转pem
type IKeyPairToPem interface {
	// PrivateKey 获取私钥
	PrivateKey() IKeyToPem
	// PublicKey 获取公钥
	PublicKey() IKeyToPem
}
