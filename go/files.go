package main

import (
	"os"
)

func OpenFile(filename string) *os.File {
	file, err := os.Open(filename)
	Check("Error open file", err)
	return file
}

func ReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	Check("Error read file", err)
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
	Check("Error create file", err)
	return file
}
