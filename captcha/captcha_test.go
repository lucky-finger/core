package captcha

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCaptcha(t *testing.T) {

	convey.Convey("测试Captcha", t, func() {

		convey.Convey("初始化", func() {
			err := Init()
			convey.So(err, convey.ShouldBeNil)
		})

		convey.Convey("生成", func() {
			data, err := Generate()
			convey.So(err, convey.ShouldBeNil)
			convey.So(data, convey.ShouldNotBeNil)

			base64Data, err := data.GetMasterImage().ToBase64()
			t.Log(base64Data)
		})
	})
}
