package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-agent/entity"
	"io"
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
