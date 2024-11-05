package otp

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func validateFlow(totp *Totp) {
	url, key, err := totp.GenerateKey("hello")
	convey.So(err, convey.ShouldBeNil)
	convey.So(url, convey.ShouldNotBeEmpty)
	convey.So(key, convey.ShouldNotBeEmpty)

	convey.So(totp._conf.AccountName, convey.ShouldBeEmpty)

	code, err := totp.GenerateCode(key)
	convey.So(err, convey.ShouldBeNil)
	convey.So(code, convey.ShouldNotBeEmpty)

	convey.So(totp.Validate(key, code), convey.ShouldBeTrue)
}

func TestTotp(t *testing.T) {
	convey.Convey("测试Totp", t, func() {

		convey.Convey("不带选项", func() {

			otp := NewTotp("tset")
			convey.So(otp, convey.ShouldNotBeNil)

			validateFlow(otp)
		})

		convey.Convey("带选项", func() {

			otp := NewTotp("tset",
				WithPeriod(10),
				WithSecretSize(32),
				WithDigits(DigitsEight),
				WithAlgorithm(AlgorithmSHA512),
			)
			convey.So(otp, convey.ShouldNotBeNil)

			validateFlow(otp)
		})

		convey.Convey("构建失败的", func() {
			otp := NewTotp("tset")
			convey.So(otp, convey.ShouldNotBeNil)

			_, _, err := otp.GenerateKey("")
			convey.So(err, convey.ShouldNotBeNil)
		})

	})
}
