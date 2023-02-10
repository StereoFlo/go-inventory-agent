package infrastructure

import (
	"go-agent/entity"
	repository2 "go-agent/repository"
)

func GetInfo(sysInfo *repository2.SystemInfoRepository, userRepo *repository2.UserRepository, partitionRepo *repository2.PartitionRepository) (*entity.SystemInfo, error) {
	hostStat, err := sysInfo.GetHost()
	if err != nil {
		return nil, err
	}
	cpuName, err := sysInfo.GetCpuName()
	if err != nil {
		return nil, err
	}
	osName, err := sysInfo.GetOsName()
	if err != nil {
		return nil, err
	}
	ram, err := sysInfo.GetRam()
	if err != nil {
		return nil, err
	}
	ip, err := sysInfo.GetIp()
	if err != nil {
		return nil, err
	}
	partitions, err := partitionRepo.GetPartitions()
	muser, err := userRepo.GetUser()
	info := new(entity.SystemInfo)
	info.User = muser
	info.Hostname = *hostStat
	info.Platform = *osName
	info.CPU = *cpuName
	info.RAM = *ram
	info.Partitions = partitions
	info.IP = *ip

	return info, nil
}
