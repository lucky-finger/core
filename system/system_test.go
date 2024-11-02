package system

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetSystemAboutInfo(t *testing.T) {
	convey.Convey("获取系统摘要信息", t, func() {
		systemAboutInfo, err := GetSystemAboutInfo()
		convey.So(err, convey.ShouldBeNil)
		convey.So(systemAboutInfo, convey.ShouldNotBeNil)
		convey.So(systemAboutInfo.CPU.Get(systemAboutInfo.CPU.Len-1), convey.ShouldNotBeNil)

		convey.Convey("错误的索引获取CPU", func() {
			convey.So(systemAboutInfo.CPU.Get(-1), convey.ShouldBeNil)
			convey.So(systemAboutInfo.CPU.Get(systemAboutInfo.CPU.Len), convey.ShouldBeNil)
		})
	})
}
