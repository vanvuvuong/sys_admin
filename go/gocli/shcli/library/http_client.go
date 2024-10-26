package library

import (
	"bytes"
	"net"
	"net/http"
	"slices"
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

func Request(method string, header map[string]string, payload []byte, url string) (http.Response, error) {
	var (
		req   *http.Request
		err   error
		delay = 4
	)
	req, err = http.NewRequest(method, url, nil)
	if method == "POST" || method == "PUT" {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(payload))
	}
	if err != nil {
		return http.Response{}, Err("Error to create new request: %w", err)
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return http.Response{}, Err("Error to send request: %w", err)
	}
	successResp := []string{"200", "201", "202"}
	if !slices.Contains(successResp, Sf("%d", resp.StatusCode)) {
		Pf("Error to get desired response.\n\tURL:%s\n\tStatus code: %d\n", url, resp.StatusCode)
		for index := 0; index < 5; index++ {
			Pf("Retry: %d - URL: %s - Delay: %ds\n", index+1, url, delay)
			time.Sleep(time.Duration(delay) * time.Second)
			resp, _ = client.Do(req)
			if slices.Contains(successResp, Sf("%d", resp.StatusCode)) {
				break
			}
		}
		if !slices.Contains(successResp, Sf("%d", resp.StatusCode)) {
			return http.Response{}, Err("Error to get desired response.\n\tURL %s\n\tStatus code: %d", url, resp.StatusCode)
		}
	}
	return *resp, nil
}
