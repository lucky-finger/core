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
		t.Log(systemAboutInfo.String())
	})
}
