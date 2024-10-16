package main

import (
	"os"
	"testing"
)

func TestFilterOut(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		ext      string
		minSize  int64
		expected bool
	}{
		{"FilterNoExtension", "testdata/dir.log", "", 0, false},
		{"FilterExtensionMatch", "testdata/dir.log", ".log", 0, false},
		{"FilterExtensionNoMatch", "testdata/dir.log", ".sh", 0, true},
		{"FilterExtensionSizeMatch", "testdata/dir.log", ".log", 10, false},
		{"FilterExtensionSizeNoMatch", "testdata/dir.log", ".log", 20, true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			info, err := os.Stat(testCase.file)
			if err != nil {
				t.Fatal(err)
			}
			file := filterOut(testCase.file, testCase.ext, testCase.minSize, info)
			if file != testCase.expected {
				t.Errorf("Expected '%t', got '%t' instead\n", testCase.expected, file)
			}
		})
	}
}
