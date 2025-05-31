package library

import (
	"bufio"
	"os"
)

func init() {
	os.MkdirAll("logs", os.ModePerm)
}

func BufferReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", Err("Error to open file: %w", err)
	}
	return string(content), nil
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
