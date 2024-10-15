package gosysadmin

import (
	"github.com/gosysadmin/basics"
	fjson "github.com/valyala/fastjson"
)

func DecodeJson(data []byte) *fjson.Value {
	jsonData := fjson.MustParse(string(data))
	var newValue fjson.Arena

	// convert to to object
	objectJsonData, err := jsonData.Object()
	basics.Check("Error convert JSON to Object", err)

	// loop over object
	objectJsonData.Visit(func(key []byte, value *fjson.Value) {
		basics.Pf("%v: %v", key, value)
	})
	// update new value for objectJsonData["example"]
	objectJsonData.Set("example", newValue.NewString("value"))

	// loop over array
	arrayJsonData, err := jsonData.Array()
	basics.Check("Error convert JSON to array", err)
	// clone new JSON array
	newArray := newValue.NewArray()
	for index, object := range arrayJsonData {
		newArray.SetArrayItem(index, object)
	}

	return jsonData
}
