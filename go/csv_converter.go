package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func OpenFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return file
}

func CreateFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return file
}

func DecodeCsv() {
	fmt.Println(">>Decoding struct to csv")
	file := OpenFile("decode_example.csv")
	reader := csv.NewReader(file)
	defer file.Close()
	reader.FieldsPerRecord = 3
	reader.Comma = ';'
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var (
		data       Data
		dataBucket []Data
	)

	for _, item := range csvData {
		data.Id, err = strconv.ParseInt(item[0], 10, 32)
		data.Name = item[1]
		data.Category = item[2]
		dataBucket = append(dataBucket, data)
		fmt.Println(data)
	}
	fmt.Println(dataBucket)
}

func main() {
	DecodeCsv()
}
