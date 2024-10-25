package library

import (
	"bufio"
	"os"
	"strings"
)

func OpenFile(filename string) *os.File {
	file, err := os.Open(filename)
	Log("Error open file", err)
	return file
}

func BufferReadFile(filename string) string {
	file, err := os.Open(filename)
	Log("Error open file", err)
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

func BufferWriteFile(filename, message string) {
	file, err := os.Create(filename)
	if err != nil {
		Log("Error:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(message)
	writer.Flush()
}
