package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-agent/entity"
	"net/http"
)

func MakeRequest(client *http.Client, host *string, data *entity.SystemInfo) error {
	jsonByte, _ := json.Marshal(data)
	_, err := client.Post(fmt.Sprintf("http://%s/devices", *host), "application/json", bytes.NewBuffer(jsonByte))
	if err != nil {
		return err
	}
	return nil
}
