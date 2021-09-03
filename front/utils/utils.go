package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func RequestToAPI(tpe, path string, data interface{}, timeOut ...time.Duration) ([]byte, bool) {
	recive, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error marshal")
	}
	req, err := http.NewRequest(tpe, path, bytes.NewBuffer(recive))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	if len(timeOut) > 0 {
		client.Timeout = timeOut[0]
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("front utils.go RequestToApi() error request:", err.Error())
		return nil, false
	}
	defer resp.Body.Close()
	answerByte, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		return answerByte, true
	}

	return answerByte, false
}
