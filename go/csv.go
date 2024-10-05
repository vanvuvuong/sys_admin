package main

import (
	"encoding/csv"
	"os"
)

func DecodeCsv(filename string, delimiter ...rune) [][]string {
	comma := ','
	if len(delimiter) > 0 {
		comma = delimiter[0]
	}
	p(">>Decoding CSV data")
	file := OpenFile(filename)
	reader := csv.NewReader(file)
	defer file.Close()
	reader.Comma = comma
	csvData, err := reader.ReadAll()
	if err != nil {
		l("DecodeCsv Err: ", err)
		os.Exit(1)
	}
	// p(csvData)
	// p(reflect.TypeOf(csvData))
	return csvData
}

func EncodeCsv(filename string, data [][]string) {
	file := CreateFile(filename)
	writer := csv.NewWriter(file)
	writer.WriteAll(data)

	if err := writer.Error(); err != nil {
		lf(err)
	}
}
