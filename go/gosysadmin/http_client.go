package gosysadmin

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gosysadmin/basics"
)

var client = http.Client{}

func GetRequest(url string, headers map[string]string) string {
	req, err := http.NewRequest("GET", url, nil)
	basics.P(">>Send GET request to ", url)
	basics.Check("Error create request", err)
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	basics.Check("Error send request", err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	basics.Check("Error parse request body", err)
	return string(body)
}

func PostRequest(url string, headers map[string]string, payload []byte) string {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	basics.Check("Error create request", err)
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	basics.Check("Error send request", err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	basics.Check("Error parse request body", err)
	return string(body)
}
