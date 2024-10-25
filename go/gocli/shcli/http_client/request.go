package http_client

import (
	"bytes"
	"net"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second, // global request timeout
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 2 * time.Second, // dial timeout
		}).DialContext,
		TLSHandshakeTimeout:   2 * time.Second, // tls handshake timeout
		ResponseHeaderTimeout: 2 * time.Second, // response header timeout
	},
}

func request(method string, header map[string]string, payload []byte, url string) (http.Response, error) {
	var (
		req *http.Request
		err error
	)
	req, err = http.NewRequest(method, url, nil)
	if method == "POST" || method == "PUT" {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(payload))
	}
	if err != nil {
		Log(err)
		return http.Response{}, err
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		Log(err)
		return http.Response{}, err
	}
	if resp.StatusCode != 200 {
		Log("Error to get request", url)
		return http.Response{}, ErrFailRequests
	}
	defer resp.Body.Close()
	return *resp, nil
}
