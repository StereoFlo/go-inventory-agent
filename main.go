package main

import (
	"go-agent/infrastructure"
	"go-agent/repository"
	"log"
	"net/http"
	"time"
)

func main() {
	address, err := infrastructure.GetServerAddress()
	if err != nil {
		log.Fatal(err)
	}
	partitionRepo := repository.NewPartitionRepository()
	userRepo := repository.NewUserRepository()
	sysInfo := repository.NewSystemInfoRepository()
	client := &http.Client{}
	for {
		info, err := infrastructure.GetInfo(sysInfo, userRepo, partitionRepo)
		_, err = infrastructure.MakeRequest(client, address, info)
		if err != nil {
			log.Println(err)
		}
		log.Println("OK")
		time.Sleep(time.Second * 10)
	}

}
