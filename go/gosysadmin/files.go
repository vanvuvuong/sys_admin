package gosysadmin

import (
	"os"

	"github.com/gosysadmin/basics"
)

func OpenFile(filename string) *os.File {
	file, err := os.Open(filename)
	basics.Check("Error open file", err)
	return file
}

func ReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	basics.Check("Error read file", err)
	// p("File content")
	// p(string(data))
	return string(data)
}

func WriteFile(filename, message string) {
	bytes := []byte(message)
	os.WriteFile(filename, bytes, 0o644)
	basics.Pf("Successfully creating file %s\n", filename)
}

func CreateFile(filename string) *os.File {
	file, err := os.Create(filename)
	basics.Check("Error create file", err)
	return file
}
