package main

import (
	"os"
)

func WriteFile(filename, message string) {
	bytes := []byte(message)
	os.WriteFile(filename, bytes, 0644)
	pf("Successfully creating file %s\n", filename)
}

func ReadFile(filename string) string {
	data, _ := os.ReadFile(filename)
	// p("File content")
	// p(string(data))
	return string(data)
}
