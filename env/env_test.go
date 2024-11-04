package env

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEnv(t *testing.T) {
	convey.Convey("测试环境变量获取", t, func() {
		t.Setenv("TEST_INT", "123")
		t.Setenv("TEST_INT8", "123")
		t.Setenv("TEST_INT32", "123")
		t.Setenv("TEST_INT64", "123")
		t.Setenv("TEST_UINT", "123")
		t.Setenv("TEST_UINT8", "123")
		t.Setenv("TEST_UINT32", "123")
		t.Setenv("TEST_UINT64", "123")
		t.Setenv("TEST_FLOAT32", "123.123")
		t.Setenv("TEST_FLOAT64", "123.123")
		t.Setenv("TEST_BOOL", "true")
		t.Setenv("TEST_STRING", "test")

		convey.Convey("正确获取", func() {
			convey.So(Get[int]("TEST_INT"), convey.ShouldEqual, 123)
			convey.So(Get[int8]("TEST_INT8"), convey.ShouldEqual, int8(123))
			convey.So(Get[int32]("TEST_INT32"), convey.ShouldEqual, int32(123))
			convey.So(Get[int64]("TEST_INT64"), convey.ShouldEqual, int64(123))
			convey.So(Get[uint]("TEST_UINT"), convey.ShouldEqual, uint(123))
			convey.So(Get[uint8]("TEST_UINT8"), convey.ShouldEqual, uint8(123))
			convey.So(Get[uint32]("TEST_UINT32"), convey.ShouldEqual, uint32(123))
			convey.So(Get[uint64]("TEST_UINT64"), convey.ShouldEqual, uint64(123))
			convey.So(Get[float32]("TEST_FLOAT32"), convey.ShouldEqual, float32(123.123))
			convey.So(Get[float64]("TEST_FLOAT64"), convey.ShouldEqual, 123.123)
			convey.So(Get[bool]("TEST_BOOL"), convey.ShouldEqual, true)
			convey.So(Get[string]("TEST_STRING"), convey.ShouldEqual, "test")
		})

		convey.Convey("获取不存在的环境变量返回默认值", func() {
			convey.So(GetWithDefault[int]("NON_EXISTENT_INT", 456), convey.ShouldEqual, 456)
			convey.So(GetWithDefault[string]("NON_EXISTENT_STRING", "default"), convey.ShouldEqual, "default")
		})

		convey.Convey("获取环境变量类型转换失败返回默认值", func() {
			t.Setenv("INVALID_INT", "invalid")
			t.Setenv("INVALID_INT8", "invalid")
			t.Setenv("INVALID_INT32", "invalid")
			t.Setenv("INVALID_INT64", "invalid")
			t.Setenv("INVALID_UINT", "invalid")
			t.Setenv("INVALID_UINT8", "invalid")
			t.Setenv("INVALID_UINT32", "invalid")
			t.Setenv("INVALID_UINT64", "invalid")
			t.Setenv("INVALID_FLOAT32", "invalid")
			t.Setenv("INVALID_FLOAT64", "invalid")
			t.Setenv("INVALID_BOOL", "invalid")
			convey.So(GetWithDefault[int]("INVALID_INT", 789), convey.ShouldEqual, 789)
			convey.So(GetWithDefault[int8]("INVALID_INT8", int8(123)), convey.ShouldEqual, int8(123))
			convey.So(GetWithDefault[int32]("INVALID_INT32", int32(789)), convey.ShouldEqual, int32(789))
			convey.So(GetWithDefault[int64]("INVALID_INT64", int64(789)), convey.ShouldEqual, int64(789))
			convey.So(GetWithDefault[uint]("INVALID_UINT", uint(789)), convey.ShouldEqual, uint(789))
			convey.So(GetWithDefault[uint8]("INVALID_UINT8", uint8(123)), convey.ShouldEqual, uint8(123))
			convey.So(GetWithDefault[uint32]("INVALID_UINT32", uint32(789)), convey.ShouldEqual, uint32(789))
			convey.So(GetWithDefault[uint64]("INVALID_UINT64", uint64(789)), convey.ShouldEqual, uint64(789))
			convey.So(GetWithDefault[float32]("INVALID_FLOAT32", float32(789.789)), convey.ShouldEqual, float32(789.789))
			convey.So(GetWithDefault[float64]("INVALID_FLOAT64", 789.789), convey.ShouldEqual, 789.789)
			convey.So(GetWithDefault[bool]("INVALID_BOOL", false), convey.ShouldEqual, false)
		})

		convey.Convey("获取环境变量为空字符串返回默认值", func() {
			t.Setenv("EMPTY_STRING", "")
			convey.So(GetWithDefault[string]("EMPTY_STRING", "default"), convey.ShouldEqual, "")
		})
	})
}
