package ali

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/lucky-finger/core/sms"
)

type smsImpl struct {
	// client 客户端
	client *client.Client
	// rawConf 原始配置
	rawConf *sms.Config
}

func (s *smsImpl) SendToPhoneWithTemplate(phone, signName, templateCode string, templateParam map[string]string) error {
	param := ""
	if len(templateParam) > 0 {
		marshal, err := json.Marshal(templateParam)
		if err != nil {
			return fmt.Errorf("error while marshaling template param: %w", err)
		}
		param = string(marshal)
	}
	request := &client.SendSmsRequest{
		PhoneNumbers:  &phone,
		SignName:      &signName,
		TemplateCode:  &templateCode,
		TemplateParam: &param,
	}

	response, err := s.client.SendSmsWithOptions(request, &service.RuntimeOptions{})
	if err != nil {
		return err
	}

	fmt.Println(response.GoString())
	return nil
}

func (s *smsImpl) SendToPhone(phone, signName string) error {
	return s.SendToPhoneWithTemplate(phone, signName, s.rawConf.DefaultTemplateCode, s.rawConf.DefaultTemplateParam())
}

func init() {
	sms.Injection(func(conf *sms.Config) (sms.ISMS, error) {
		openapiConf := openapi.Config{
			AccessKeyId:     &conf.AccessKeyID,
			AccessKeySecret: &conf.AccessKeySecret,
			Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
		}

		cli, err := client.NewClient(&openapiConf)
		if err != nil {
			return nil, err
		}

		return &smsImpl{client: cli, rawConf: conf}, nil
	})
}
