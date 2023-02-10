package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-agent/entity"
	repository2 "go-agent/repository"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	host, err := getServerIp()
	if err != nil {
		panic(err)
	}
	partitionRepo := repository2.NewPartitionRepository()
	userRepo := repository2.NewUserRepository()
	sysInfo := repository2.NewSystemInfoRepository()
	client := &http.Client{}
	for {
		info := getInfo(sysInfo, userRepo, partitionRepo)
		err := makeRequest(client, host, info)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("OK")
		time.Sleep(time.Second * 10)
	}

}

func getInfo(sysInfo *repository2.SystemInfoRepository, userRepo *repository2.UserRepository, partitionRepo *repository2.PartitionRepository) *entity.SystemInfo {
	hostStat, err := sysInfo.GetHost()
	if err != nil {
		log.Fatalf(err.Error())
	}
	cpuName, err := sysInfo.GetCpuName()
	if err != nil {
		log.Fatalf(err.Error())
	}
	osName, err := sysInfo.GetOsName()
	if err != nil {
		log.Fatalf(err.Error())
	}
	ram, err := sysInfo.GetRam()
	if err != nil {
		log.Fatalf(err.Error())
	}
	ip, err := sysInfo.GetIp()
	if err != nil {
		log.Fatalf(err.Error())
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

	return info
}

func makeRequest(client *http.Client, host *string, data *entity.SystemInfo) error {
	jsonByte, _ := json.Marshal(data)
	_, err := client.Post(fmt.Sprintf("http://%s/devices", *host), "application/json", bytes.NewBuffer(jsonByte))
	if err != nil {
		return err
	}
	return nil
}

func getServerIp() (*string, error) {
	log.Println("start listening the server on 2712")
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
	log.Println("Success! Received service address: " + r)

	return &r, nil
}
