package sms

import "errors"

type Config struct {
	// AccessKeyID 认证ID
	AccessKeyID string
	// AccessKeySecret 认证密钥
	AccessKeySecret string
	// DefaultTemplateCode 默认模板
	DefaultTemplateCode string
	// DefaultTemplateParam 默认模板参数
	DefaultTemplateParam func() map[string]string
}

// ISMS 短信接口
type ISMS interface {
	// SendToPhone 发送短信
	SendToPhone(phone, signName string) error
	// SendToPhoneWithTemplate 发送短信伴随模板
	SendToPhoneWithTemplate(phone, sigName, templateCode string, templateParam map[string]string) error
}

type ISMSFactory func(conf *Config) (ISMS, error)

// iSMSFactory 工厂
var iSMSFactory ISMSFactory

// Injection 注入驱动
func Injection(factory ISMSFactory) {
	iSMSFactory = factory
}

// globalInstance 全局实例
var globalInstance ISMS

// Init 初始化
func Init(conf *Config) (err error) {
	if iSMSFactory == nil {
		return errors.New("sms driver is nil")
	}
	globalInstance, err = iSMSFactory(conf)
	return err
}
