package system

import (
	"context"
	"encoding/json"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

// AboutInfo SystemAboutInfo 一个包含有关系统信息的结构体
type AboutInfo struct {
	// InfoStat 主机状态
	*host.InfoStat
	// CPUS CPU信息
	CPUS []cpu.InfoStat
	// Mem 内存信息
	Mem *mem.VirtualMemoryStat
}

// String 返回SystemAboutInfo的字符串表示形式
func (s *AboutInfo) String() string {
	marshal, _ := json.Marshal(s)
	return string(marshal)
}

// GetSystemAboutInfo 获取系统信息
func GetSystemAboutInfo() (*AboutInfo, error) {
	return GetSystemAboutInfoWithContext(context.Background())

}

// GetSystemAboutInfoWithContext 获取系统信息
func GetSystemAboutInfoWithContext(ctx context.Context) (*AboutInfo, error) {
	hostInfo, err := host.InfoWithContext(ctx)
	if err != nil {
		return nil, err
	}

	cpuInfo, err := cpu.InfoWithContext(ctx)
	if err != nil {
		return nil, err
	}

	memInfo, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return nil, err
	}

	return &AboutInfo{
		InfoStat: hostInfo,
		CPUS:     cpuInfo,
		Mem:      memInfo,
	}, nil

}
