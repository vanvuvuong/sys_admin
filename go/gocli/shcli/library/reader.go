package library

import (
	"encoding/csv"
)

func DecodeCsv(filename string, delimiter ...rune) [][]string {
	comma := ','
	if len(delimiter) > 0 {
		comma = delimiter[0]
	}
	Log(">>Decoding CSV data")
	file := OpenFile(filename)
	reader := csv.NewReader(file)
	defer file.Close()
	reader.Comma = comma
	csvData, err := reader.ReadAll()
	Log("Error when decode CSV file", err)
	// p(csvData)
	// p(reflect.TypeOf(csvData))
	return csvData
}
