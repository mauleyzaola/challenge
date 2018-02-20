package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	client  *http.Client
	port    int
	baseUrl string = "http://localhost:%d"
)

func newClient() *http.Client {
	return &http.Client{Timeout: time.Millisecond * 100}
}

func request(method, endpoint string, message interface{}) ([]byte, error) {
	var (
		reader *bytes.Buffer
		req    *http.Request
		err    error
	)
	if message != nil {
		data, err := json.Marshal(message)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewBuffer(data)
	}
	uri := fmt.Sprintf(baseUrl, port)
	uri += endpoint

	if reader != nil {
		req, err = http.NewRequest(method, uri, reader)
	} else {
		req, err = http.NewRequest(method, uri, nil)
	}
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {

	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
