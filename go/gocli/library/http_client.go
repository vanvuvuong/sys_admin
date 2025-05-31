package library

import (
	"crypto/tls"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/valyala/fastjson"
)

func Request(method string, header map[string]string, payload []byte, url string, sslCheck ...bool) (resty.Response, error) {
	var (
		resp                                  *resty.Response
		err                                   error
		skipSslCheck, debug                   = false, false
		retryCount, retryWait, requestTimeout = RetryCount, RetryWait, RequestTimeout
	)
	if len(sslCheck) > 0 {
		skipSslCheck = !sslCheck[0]
	} else if len(sslCheck) == 2 {
		debug = sslCheck[1]
	}
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: skipSslCheck}).
		SetRetryCount(retryCount).
		SetRetryWaitTime(retryWait).
		SetTimeout(requestTimeout).
		SetDebug(debug)
	client.SetHeaders(header)
	log.Debug("-------------------------------------")
	log.Debug("url: ", url)
	if method == "GET" {
		resp, err = client.R().Get(url)
	}
	if method == "BASIC_AUTH" {
		basicAuth := fastjson.MustParse(string(payload))
		client.SetBasicAuth(strings.Trim(basicAuth.Get("username").String(), "\""), strings.Trim(basicAuth.Get("password").String(), "\""))
		resp, err = client.R().Get(url)
	}
	if method == "POST" {
		header["content-type"] = "application/json"
		client.SetHeaders(header)
		resp, err = client.R().SetBody(payload).Post(url)
		log.Debug("payload: ", string(payload))
	}
	if method == "PUT" {
		header["content-type"] = "application/json"
		client.SetHeaders(header)
		resp, err = client.R().SetBody(payload).Put(url)
		log.Debug("payload: ", string(payload))
	}
	log.Debug("response: ", resp)
	log.Debug("error: ", err)
	log.Debug("-------------------------------------")
	if resp.IsError() {
		return *resp, Err("Error to get desired response: %s", string(resp.Body()))
	}
	return *resp, err
}

func AuthRequest(header map[string]string, username, password, url string, sslCheck ...bool) (resty.Response, error) {
	var (
		resp                                  *resty.Response
		err                                   error
		skipSslCheck, debug                   = false, false
		retryCount, retryWait, requestTimeout = 5, 5 * time.Second, 15 * time.Second
	)
	if len(sslCheck) > 0 {
		skipSslCheck = !sslCheck[0]
	} else if len(sslCheck) == 2 {
		debug = sslCheck[1]
	}
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: skipSslCheck}).SetRetryCount(retryCount).SetRetryWaitTime(retryWait).SetTimeout(requestTimeout).SetDebug(debug)
	client.SetHeaders(header)
	log.Debug("-------------------------------------")
	log.Debug("url: ", url)
	client.SetFormData(map[string]string{
		"client_id":  "account",
		"grant_type": "password",
		"username":   username,
		"password":   password,
	})
	resp, err = client.R().Post(url)
	if resp.IsError() {
		return *resp, Err("Error to get desired response: %s", string(resp.Body()))
	}
	return *resp, err
}

func Download(url, fileSavePath string, sslCheck ...bool) error {
	var (
		err                 error
		skipSslCheck, debug = false, false
	)
	if len(sslCheck) > 0 {
		skipSslCheck = !sslCheck[0]
	}
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: skipSslCheck}).
		SetRetryCount(RetryCount).
		SetRetryWaitTime(RetryWait).
		SetTimeout(DownloadRequestTimeout).
		SetDebug(debug)
	_, err = client.R().SetOutput(fileSavePath).Get(url)
	return err
}

func Upload(url, filePath string, header map[string]string, sslCheck ...bool) error {
	var (
		err                 error
		skipSslCheck, debug = false, false
	)
	if len(sslCheck) > 0 {
		skipSslCheck = !sslCheck[0]
	}
	// fileReader, _ := os.Open(filePath)
	// defer fileReader.Close()
	fileByte, _ := os.ReadFile(filePath)
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: skipSslCheck}).
		SetRetryCount(RetryCount).
		SetRetryWaitTime(RetryWait).
		SetTimeout(UploadRequestTimeout).
		SetDebug(debug)
	client.SetHeaders(header)
	log.Debug("-------------------------------------")
	resp, err := client.R().
		SetContentLength(true).
		// SetMultipartField("cell-check.zip", filePath, "gzip", fileReader).
		SetBody(map[string][]byte{"formData": fileByte}).
		Post(url)
	log.Debug("status: ", string(resp.Status()))
	log.Debug("header: ", resp.Request.Header)
	// log.Debug("payload: ", resp.Request.Body)
	log.Debug("response: ", resp)
	log.Debug("error: ", err)
	log.Debug("-------------------------------------")
	return err
}
