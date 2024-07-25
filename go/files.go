package main

import (
	"fmt"
	"os"
)

func writeFile(filename, message string) {
	bytes := []byte(message)
	os.WriteFile(filename, bytes, 0644)
	fmt.Printf("Successfully creating file %s\n", filename)
}

func readFile(filename string) {
	data, _ := os.ReadFile(filename)
	fmt.Println("File content")
	fmt.Println(string(data))
}
