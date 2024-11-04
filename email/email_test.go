package email

import (
	"github.com/lucky-finger/core/env"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEmail(t *testing.T) {
	convey.Convey("测试邮箱", t, func() {
		convey.Convey("创建账户", func() {
			account := New("", "")
			convey.So(account.address, convey.ShouldBeEmpty)
			convey.So(account.password, convey.ShouldBeEmpty)
			convey.So(account.smtpHost, convey.ShouldBeEmpty)
			convey.So(account.smtpPort, convey.ShouldEqual, defaultSmtpPort)

			account = New("xxxx@qq.com", "123", WithSmtpPort(123))
			convey.So(account.smtpHost, convey.ShouldEqual, "smtp.qq.com")
			convey.So(account.smtpPort, convey.ShouldEqual, 123)

			account = New("xxx@qq.com", "123", WithSmtpHost("smtp.xxx.com"))
			convey.So(account.smtpHost, convey.ShouldEqual, "smtp.xxx.com")
		})

		convey.Convey("发送测试邮件", func() {
			testEmailAddressEnv := "TEST_EMAIL_ADDRESS"
			testEmailPasswordEnv := "TEST_EMAIL_PASSWORD"
			testEmailSendToEnv := "TEST_EMAIL_SEND_TO"

			address := env.Get[string](testEmailAddressEnv)
			password := env.Get[string](testEmailPasswordEnv)
			sendTo := env.Get[string](testEmailSendToEnv)

			if address == "" || password == "" || sendTo == "" {
				t.Skip("请设置环境变量: " + testEmailAddressEnv + " 、 " + testEmailPasswordEnv + "和" + testEmailSendToEnv + " 以测试邮件发送")
				return
			}

			emailAccount := New(address, password)
			err := emailAccount.SendTo(
				sendTo,
				MsgWithSubject("测试邮件"),
				MsgWithTextContent("这是一封文本测试邮件"),
			)
			convey.So(err, convey.ShouldBeNil)

			err = emailAccount.SendTo(
				sendTo,
				MsgWithSubject("测试邮件"),
				MsgWithHtmlContent("<h1>这是一封html测试邮件</h1>"),
			)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
