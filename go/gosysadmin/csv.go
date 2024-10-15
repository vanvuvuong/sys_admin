package gosysadmin

import (
	"encoding/csv"

	"github.com/gosysadmin/basics"
)

func DecodeCsv(filename string, delimiter ...rune) [][]string {
	comma := ','
	if len(delimiter) > 0 {
		comma = delimiter[0]
	}
	basics.P(">>Decoding CSV data")
	file := OpenFile(filename)
	reader := csv.NewReader(file)
	defer file.Close()
	reader.Comma = comma
	csvData, err := reader.ReadAll()
	basics.Check("Error when decode CSV file", err)
	// p(csvData)
	// p(reflect.TypeOf(csvData))
	return csvData
}

func EncodeCsv(filename string, data [][]string) {
	file := CreateFile(filename)
	writer := csv.NewWriter(file)
	writer.WriteAll(data)
	basics.Check("Error when encode CSV file", writer.Error())
}
