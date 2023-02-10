package main

import (
	"fmt"
	"go-agent/infrastructure"
	"go-agent/repository"
	"log"
	"net/http"
	"time"
)

func main() {
	host, err := infrastructure.GetServerAddress()
	if err != nil {
		log.Fatal(err)
	}
	partitionRepo := repository.NewPartitionRepository()
	userRepo := repository.NewUserRepository()
	sysInfo := repository.NewSystemInfoRepository()
	client := &http.Client{}
	for {
		info, err := infrastructure.GetInfo(sysInfo, userRepo, partitionRepo)
		err = infrastructure.MakeRequest(client, host, info)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("OK")
		time.Sleep(time.Second * 10)
	}

}
