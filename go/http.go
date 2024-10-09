package main

import (
	"bytes"
	"io"
	"net/http"
)

var client = http.Client{}

func GetRequest(url, token string) string {
	req, err := http.NewRequest("GET", url, nil)
	Check("Error create request", err)
	req.Header.Add("user-agent:", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0")
	req.Header.Add("authorization:", sf("Bearer %s", token))
	resp, err := client.Do(req)
	Check("Error send request", err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	Check("Error parse request body", err)
	return string(body)
}

func PostRequest(url, token string, payload []byte) string {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	Check("Error create request", err)
	req.Header.Add("user-agent:", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0")
	req.Header.Add("authorization:", sf("Bearer %s", token))
	resp, err := client.Do(req)
	Check("Error send request", err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	Check("Error parse request body", err)
	return string(body)
}
