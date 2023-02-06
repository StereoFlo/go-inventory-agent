package repository

import (
	"errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"net"
	"os"
)

type SystemInfoRepository struct {
}

func NewSystemInfoRepository() *SystemInfoRepository {
	return &SystemInfoRepository{}
}

func (SystemInfoRepository) GetHost() (*string, error) {
	hostStat, err := host.Info()
	if err != nil {
		return nil, err
	}
	return &hostStat.Hostname, nil
}

func (SystemInfoRepository) GetOsName() (*string, error) {
	hostStat, err := host.Info()
	if err != nil {
		return nil, err
	}
	return &hostStat.Platform, nil
}

func (SystemInfoRepository) GetCpuName() (*string, error) {
	cpuStat, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	return &cpuStat[0].ModelName, nil
}

func (SystemInfoRepository) GetIp() (*string, error) {
	thisHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(thisHost)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ip := ipv4.String()
			return &ip, nil
		}
	}
	return nil, errors.New("IP not found")
}

func (SystemInfoRepository) GetRam() (*uint64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	ret := vmStat.Total / 1024 / 1024
	return &ret, err
}
