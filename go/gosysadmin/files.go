package gosysadmin

import (
	"bufio"
	"os"
	"strings"

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

func BufferReadFile(filename string) string {
	file, err := os.Open(filename)
	basics.Check("Error open file", err)
	defer file.Close()

	reader := bufio.NewReader(file)
	lines := make([]string, 100)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func WriteFile(filename, message string) {
	bytes := []byte(message)
	os.WriteFile(filename, bytes, 0o644)
	basics.Pf("Successfully creating file %s\n", filename)
}

func BufferWriteFile(filename, message string) {
	file, err := os.Create(filename)
	if err != nil {
		basics.Check("Error:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(message)
	writer.Flush()
}

func CreateFile(filename string) *os.File {
	file, err := os.Create(filename)
	basics.Check("Error create file", err)
	return file
}
