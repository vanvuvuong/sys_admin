package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Id       int64  `json."id"`
	Category string `json."category"`
	Name     string `json."name"`
}

func EncodeJson() {
	fmt.Println(">>Encoding struct to Json")
	data := &Data{Id: 1, Category: "ProgrammingLanguage", Name: "Golang"}
	encodingData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(encodingData))
}

func DecodeJson() {
	var newData Data
	message := `{"Id":1,"Category":"ProgrammingLanguage","Name":"Golang"}`
	fmt.Println(">>Decoding Json to struct")
	json.Unmarshal([]byte(message), &newData)
	fmt.Printf("Id: %d\n", newData.Id)
	fmt.Printf("Category: %s\n", newData.Category)
	fmt.Printf("Name: %s\n", newData.Name)
}
