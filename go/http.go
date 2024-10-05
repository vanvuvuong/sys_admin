package main

import (
	"bufio"
	"net/http"
)

func GetRequest() {
	url := "https://google.com"
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		p(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
