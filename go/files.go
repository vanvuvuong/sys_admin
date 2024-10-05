package main

import (
	"os"
)

func OpenFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		p(err)
		os.Exit(1)
	}
	return file
}

func ReadFile(filename string) string {
	data, _ := os.ReadFile(filename)
	// p("File content")
	// p(string(data))
	return string(data)
}

func WriteFile(filename, message string) {
	bytes := []byte(message)
	os.WriteFile(filename, bytes, 0644)
	pf("Successfully creating file %s\n", filename)
}

func CreateFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		p(err)
		os.Exit(1)
	}
	return file
}
