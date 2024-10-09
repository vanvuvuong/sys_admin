package main

import (
	"encoding/json"
	"reflect"
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

// func DecodeJson() {
// 	var newData Data
// 	message := `{"Id":1,"Category":"ProgrammingLanguage","Name":"Golang"}`
// 	p(">>Decoding Json to struct")
// 	json.Unmarshal([]byte(message), &newData)
// 	pf("Id: %d\n", newData.Id)
// 	pf("Category: %s\n", newData.Category)
// 	pf("Name: %s\n", newData.Name)
// }

func stringify(data map[string]interface{}) map[string]interface{} {
	for key, value := range data {
		stringValue, err := value.(string)
		if err {
			jsonType := reflect.TypeOf(value)
			switch sf("%s", jsonType) {
			case "map[string]interface {}":
				return stringify(value.(map[string]interface{}))
			default:
				continue
			}
		} else {
			value = stringValue
		}
		data[key] = value
	}
	return data
}

func sliceStringify(data []interface{}) []map[string]interface{} {
	jsonType := reflect.TypeOf(data)
	var newData []map[string]interface{}
	switch sf("%s", jsonType) {
	case "[]interface {}":
		for _, row := range data.([]interface{}) {
			newData = append(newData, stringify(row.(map[string]interface{})))
		}
		return newData
	case "[]map[string]interface {}":
		for _, row := range data.(map[string]interface{}) {
			newData = append(newData, stringify(row.(map[string]interface{})))
		}
		return newData
	case "[]map":
		return newData
	default:
		return newData
	}
}

func DecodeJson(data []byte) map[string]any {
	var jsonData map[string]interface{}
	err := json.Unmarshal(data, &jsonData)
	jsonData = stringify(jsonData)
	Check("Error decode json", err)
	return jsonData
}
