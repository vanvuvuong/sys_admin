package main

import (
	"encoding/json"
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
	if err != nil {
		p(err)
		return
	}
	p(string(encodingData))
}

func DecodeJson() {
	var newData Data
	message := `{"Id":1,"Category":"ProgrammingLanguage","Name":"Golang"}`
	p(">>Decoding Json to struct")
	json.Unmarshal([]byte(message), &newData)
	pf("Id: %d\n", newData.Id)
	pf("Category: %s\n", newData.Category)
	pf("Name: %s\n", newData.Name)
}
