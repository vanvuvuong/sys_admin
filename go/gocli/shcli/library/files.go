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

func BufferReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", Err("Error to open file: %w", err)
	}
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
	return strings.Join(lines, "\n"), nil
}

func BufferWriteFile(filename, message string) error {
	var err error
	file, err := os.Create(filename)
	if err != nil {
		return Err("Error to create file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(message)
	writer.Flush()
	return err
}
