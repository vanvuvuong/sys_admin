package main

import (
	"encoding/json"
	"reflect"
	"strings"
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

func mapStringify(data map[string]interface{}) map[string]any {
	for key, value := range data {
		stringValue, err := value.(string)
		if err {
			l(err)
			jsonType := reflect.TypeOf(value)
			switch jsonType.String() {
			case "map[string]interface {}":
				tempValue, _ := value.(map[string]interface{})
				Check("Error parse nested JSON", &json.MarshalerError{})
				value = mapStringify(tempValue)
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
	switch jsonType.String() {
	case "[]interface {}":
		for _, row := range data {
			tempRow := row.(map[string]interface{})
			Check("Error parse nested JSON", &json.MarshalerError{})
			newData = append(newData, mapStringify(tempRow))
		}
		return newData
	default:
		return newData
	}
}

func DecodeJson(data []byte) map[string]any {
	var jsonData map[string]interface{}
	err := json.Unmarshal(data, &jsonData)
	for key, value := range jsonData {
		switch strings.HasPrefix(reflect.TypeOf(value).String(), "[]") {
		case true:
			tempValue, _ := value.([]interface{})
			Check("Error parse nested JSON", &json.MarshalerError{})
			value = sliceStringify(tempValue)
		default:
			tempValue, _ := value.(map[string]interface{})
			Check("Error parse nested JSON", &json.MarshalerError{})
			value = mapStringify(tempValue)
		}
		jsonData[key] = value
	}
	Check("Error decode json", err)
	return jsonData
}

// func process(in interface{}) {
//   switch v := in.(type) {
//   case map[string]*Book:
//      for s, b := range v {
//          // b has type *Book
//          fmt.Printf("%s: book=%v\n" s, b)
//      }
//   case map[string]*Author:
//      for s, a := range v {
//          // a has type *Author
//          fmt.Printf("%s: author=%v\n" s, a)
//      }
//    case []*Book:
//      for i, b := range v {
//          fmt.Printf("%d: book=%v\n" i, b)
//      }
//    case []*Author:
//      for i, a := range v {
//          fmt.Printf("%d: author=%v\n" i, a)
//      }
//    case *Book:
//      fmt.Ptintf("book=%v\n", v)
//    case *Author:
//      fmt.Printf("author=%v\n", v)
//    default:
//      // handle unknown type
//    }
// }
