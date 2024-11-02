package system

import (
	"context"
	"encoding/json"
	"github.com/lucky-finger/core"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"time"
)

// AboutInfo SystemAboutInfo 一个包含有关系统信息的结构体
type AboutInfo struct {
	// HostName 主机名
	HostName string
	// UpTime 运行时间
	Uptime time.Duration
	// BootTime `time.Time` 表示系统启动时间
	BootTime time.Time
	// ProcessNumber 进程数量
	ProcessNumber uint64
	// OS 系统名称 ex: freebsd, linux
	OS string
	// Platform 平台名称 ex: ubuntu, linuxmint
	Platform string
	// PlatformFamily 平台家族 ex: debian, rhel
	PlatformFamily string
	// PlatformVersion 平台版本
	PlatformVersion string
	// KernelVersion 内核版本
	KernelVersion string
	// KernelArch 内核架构
	KernelArch string
	// CPU CPU信息
	CPU *CpuInfo
	// Mem 内存信息
	Mem *MemoryInfo
}

// CpuInfo CPU信息
type CpuInfo struct {
	// Len CPU数量
	Len int
	// Infos CPU信息
	Infos []cpu.InfoStat
}

type MemoryInfo struct {
	// Total 总内存
	Total core.Unit
	// Available 可用内存
	Available core.Unit
	// Used 已经使用了的内存
	Used core.Unit
	// UsedPercent 已经使用了的内存百分比
	UsedPercent float64
	// Free 空闲内存
	Free core.Unit
}

// Get 获取CPU信息
func (c *CpuInfo) Get(index int) *cpu.InfoStat {
	if index < 0 || index >= c.Len {
		return nil
	}
	return &c.Infos[index]
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
		HostName:        hostInfo.Hostname,
		Uptime:          time.Second * time.Duration(hostInfo.Uptime),
		BootTime:        time.Unix(int64(hostInfo.BootTime), 0),
		ProcessNumber:   hostInfo.Procs,
		OS:              hostInfo.OS,
		Platform:        hostInfo.Platform,
		PlatformFamily:  hostInfo.PlatformFamily,
		PlatformVersion: hostInfo.PlatformVersion,
		KernelVersion:   hostInfo.KernelVersion,
		KernelArch:      hostInfo.KernelArch,
		CPU: &CpuInfo{
			Len:   len(cpuInfo),
			Infos: cpuInfo,
		},
		Mem: &MemoryInfo{
			Total:       core.Unit(memInfo.Total),
			Available:   core.Unit(memInfo.Available),
			Used:        core.Unit(memInfo.Used),
			UsedPercent: memInfo.UsedPercent,
			Free:        core.Unit(memInfo.Free),
		},
	}, nil

}
