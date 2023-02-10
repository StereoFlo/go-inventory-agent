package infrastructure

import (
	"log"
	"net"
)

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
