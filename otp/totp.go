package otp

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"time"
)

// Totp 时间同步的动态口令
type Totp struct {
	_conf totp.GenerateOpts
}

// GenerateKey 生成key
func (t *Totp) GenerateKey(accountName string) (string, string, error) {
	conf := t._conf
	conf.AccountName = accountName

	generate, err := totp.Generate(conf)
	if err != nil {
		return "", "", err
	}

	return generate.URL(), generate.Secret(), nil
}

// GenerateCode 生成动态口令
func (t *Totp) GenerateCode(key string) (string, error) {
	return totp.GenerateCode(key, time.Now())
}

// Validate 验证动态口令
func (t *Totp) Validate(key, code string) bool {
	return totp.Validate(code, key)
}

// TotpOpt 配置选项
type TotpOpt func(*totp.GenerateOpts)

// WithPeriod 设置时间周期
func WithPeriod(period uint) TotpOpt {
	return func(o *totp.GenerateOpts) {
		o.Period = period
	}
}

// WithSecretSize 设置密钥大小
func WithSecretSize(size uint) TotpOpt {
	return func(o *totp.GenerateOpts) {
		o.SecretSize = size
	}
}

type Digits otp.Digits

const (
	DigitsSix   Digits = 6
	DigitsEight Digits = 8
)

// WithDigits 设置动态口令的位数
func WithDigits(digits Digits) TotpOpt {
	return func(o *totp.GenerateOpts) {
		o.Digits = otp.Digits(digits)
	}
}

type Algorithm otp.Algorithm

const (
	// AlgorithmSHA1 should be used for compatibility with Google Authenticator.
	//
	// See https://github.com/pquerna/otp/issues/55 for additional details.
	AlgorithmSHA1 Algorithm = iota
	AlgorithmSHA256
	AlgorithmSHA512
	AlgorithmMD5
)

func WithAlgorithm(algorithm Algorithm) TotpOpt {
	return func(o *totp.GenerateOpts) {
		o.Algorithm = otp.Algorithm(algorithm)
	}
}

// NewTotp 创建一个新的时间同步的动态口令对象
func NewTotp(issuer string, opts ...TotpOpt) *Totp {

	_base := &totp.GenerateOpts{
		Issuer: issuer,
	}

	for _, opt := range opts {
		opt(_base)
	}

	return &Totp{*_base}

}
