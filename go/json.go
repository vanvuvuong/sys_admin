package main

import (
	fjson "github.com/valyala/fastjson"
)

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
