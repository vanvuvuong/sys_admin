package library

import (
	"encoding/csv"
	"os"
)

func DecodeCsv(filename string, delimiter ...rune) ([][]string, error) {
	var err error
	comma := ','
	if len(delimiter) > 0 {
		comma = delimiter[0]
	}
	file, err := os.Open(filename)
	if err != nil {
		return [][]string{}, Err("Error to open CSV file", err)
	}
	reader := csv.NewReader(file)
	defer file.Close()
	reader.Comma = comma
	csvData, err := reader.ReadAll()
	if err != nil {
		Err("Error to decode CSV file", err)
		return [][]string{}, err
	}
	// p(csvData)
	// p(reflect.TypeOf(csvData))
	return csvData, err
}
