package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-agent/entity"
	"go-agent/repository"
	"io"
	"log"
	"net"
	"net/http"
)

func MakeRequest(client *http.Client, host *string, data *entity.SystemInfo) ([]byte, error) {
	jsonByte, _ := json.Marshal(data)
	resp, err := client.Post(fmt.Sprintf("http://%s/devices", *host), "application/json", bytes.NewBuffer(jsonByte))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetServerAddress() (*string, error) {
	log.Println("Trying to find a server")
	pc, err := net.ListenPacket("udp4", ":2712")
	if err != nil {
		return nil, err
	}
	defer pc.Close()

	buf := make([]byte, 1024)
	n, _, err := pc.ReadFrom(buf)
	if err != nil {
		return nil, err
	}

	r := string(buf[:n])
	log.Println("Success! Received server address: " + r)

	return &r, nil
}

func GetInfo(
	sysInfo *repository.SystemInfoRepository,
	userRepo *repository.UserRepository,
	partitionRepo *repository.PartitionRepository) (*entity.SystemInfo, error) {
	host, err := sysInfo.GetHost()
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
	osUser, err := userRepo.GetUser()
	info := new(entity.SystemInfo)
	info.User = osUser
	info.Hostname = *host
	info.Platform = *osName
	info.CPU = *cpuName
	info.RAM = *ram
	info.Partitions = partitions
	info.IP = *ip

	return info, nil
}
