package email

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"strings"
)

const (
	// defaultSmtpPort 默认smtp协议端口
	defaultSmtpPort int = 587
)

// Account 邮箱账户
type Account struct {
	// address 邮箱地址
	address string
	// password 邮箱密码
	password string
	// SmtpHost SMTP服务器地址
	smtpHost string
	// SmtpPort SMTP服务器端口
	smtpPort int
	// ssl 是否使用ssl
	ssl bool
	// tlsConfig TLS配置
	tlsConfig *tls.Config
	// _senderDialer 邮件发送对象
	_senderDialer *gomail.Dialer
}

// SendTo 发送邮件
func (a *Account) SendTo(address string, opts ...MsgOption) error {
	if a._senderDialer == nil {
		return fmt.Errorf("email account not initialized")
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", a.address)
	msg.SetHeader("To", address)
	for _, opt := range opts {
		opt(msg)
	}

	return a._senderDialer.DialAndSend(msg)
}

// Option 选项
type Option func(*Account)

// WithSmtpHost 设置SMTP服务器地址
func WithSmtpHost(host string) Option {
	return func(account *Account) {
		account.smtpHost = host
	}
}

// WithSmtpPort 设置SMTP服务器端口
func WithSmtpPort(port int) Option {
	return func(account *Account) {
		account.smtpPort = port
	}
}

// WithDisabledSSL 禁用SSL
func WithDisabledSSL() Option {
	return func(account *Account) {
		account.ssl = false
	}
}

// WithSkipTLSVerify 跳过TLS验证
func WithSkipTLSVerify() Option {
	return func(account *Account) {
		account.tlsConfig = &tls.Config{InsecureSkipVerify: true}
	}
}

// extractSmtpHost 从邮箱地址中提取SMTP服务器地址
func extractSmtpHost(email string) (string, error) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid email address: %s", email)
	}
	return "smtp." + parts[1], nil
}

// New 创建邮箱账户
func New(address, password string, opts ...Option) *Account {
	smtpHost, _ := extractSmtpHost(address)
	account := &Account{
		address:  address,
		password: password,
		smtpHost: smtpHost,
		smtpPort: defaultSmtpPort,
		ssl:      true,
	}

	for _, opt := range opts {
		opt(account)
	}

	dialer := gomail.NewDialer(account.smtpHost, account.smtpPort, account.address, account.password)
	dialer.SSL = account.ssl
	dialer.TLSConfig = account.tlsConfig
	account._senderDialer = dialer
	return account
}
