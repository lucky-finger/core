package ali

import (
	"github.com/lucky-finger/core/env"
	"testing"
)

func TestSend(t *testing.T) {
	accessKey := env.Get[string]("SMS_ALI_ACCESS_KEY")
	accessSecret := env.Get[string]("SMS_ALI_ACCESS_SECRET")
	signName := env.Get[string]("SMS_ALI_SIGN_NAME")
	templateCode := env.Get[string]("SMS_ALI_TEMPLATE_CODE")
	phone := env.Get[string]("SMS_ALI_PHONE")

	if accessKey == "" || accessSecret == "" || signName == "" || templateCode == "" || phone == "" {
		t.Skip()
		return
	}
}
