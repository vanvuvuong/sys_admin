package main

import (
	"encoding/json"

	fjson "github.com/valyala/fastjson"
)

type Data struct {
	Id       int64  `json."id"`
	Category string `json."category"`
	Name     string `json."name"`
}

func EncodeJson() {
	p(">>Encoding struct to Json")
	data := &Data{Id: 1, Category: "ProgrammingLanguage", Name: "Golang"}
	encodingData, err := json.Marshal(data)
	Check("Error encoding JSON", err)
	p(string(encodingData))
}

func DecodeJson(data []byte) *fjson.Value {
	jsonData := fjson.MustParse(string(data))
	return jsonData
}
