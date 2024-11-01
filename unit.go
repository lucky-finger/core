package core

import (
	"fmt"
	"time"
)

// Unit 数据单位
type Unit uint64

const (
	Byte Unit = 1
	// KB KB单位
	KB = Byte * 1024
	// MB MB单位
	MB = KB * 1024
	// GB GB单位
	GB = MB * 1024
	// TB TB单位
	TB = GB * 1024
)

// ByteTo 将字节转换为指定单位
func (u Unit) ByteTo(unit Unit) float64 {
	fu := float64(u)
	switch unit {
	case KB:
		return fu / float64(KB)
	case MB:
		return fu / float64(MB)
	case GB:
		return fu / float64(GB)
	case TB:
		return fu / float64(TB)
	default:
		return fu
	}
}

// UnitString 返回单位的字符串表示形式
func (u Unit) UnitString() string {
	switch {
	case u >= TB:
		return "TB"
	case u >= GB:
		return "GB"
	case u >= MB:
		return "MB"
	case u >= KB:
		return "KB"
	default:
		return "B"
	}
}

// String 返回单位的字符串表示形式, 自动转换为合适的单位，保留两位小数
func (u Unit) String() string {
	time.Second
	switch {
	case u >= TB:
		return fmt.Sprintf("%.2fTB", u.ByteTo(TB))
	case u >= GB:
		return fmt.Sprintf("%.2fGB", u.ByteTo(GB))
	case u >= MB:
		return fmt.Sprintf("%.2fMB", u.ByteTo(MB))
	case u >= KB:
		return fmt.Sprintf("%.2fKB", u.ByteTo(KB))
	default:
		return fmt.Sprintf("%dB", u)
	}
}

// TimeUnit 时间单位
type TimeUnit uint64
