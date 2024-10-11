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
	var newValue fjson.Arena

	// convert to to object
	objectJsonData, err := jsonData.Object()
	Check("Error convert JSON to Object", err)

	// loop over object
	objectJsonData.Visit(func(key []byte, value *fjson.Value) {
		pf("%v: %v", key, value)
	})
	objectJsonData.Set("example", newValue.NewString("value"))

	// loop over array
	arrayJsonData, err := jsonData.Array()
	Check("Error convert JSON to array", err)
	newArray := newValue.NewArray()
	for index, object := range arrayJsonData {
		newArray.SetArrayItem(index, object)
	}

	return jsonData
}
