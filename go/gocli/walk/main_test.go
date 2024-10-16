package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name     string
		root     string
		cfg      config
		expected string
	}{
		{
			name: "NoFilter", root: "testdata",
			cfg:      config{ext: "", size: 0, list: true},
			expected: "testdata/dir.log\ntestdata/dir2/script.sh\n",
		},
		{
			name: "FilterExtensionMatch", root: "testdata",
			cfg:      config{ext: ".log", size: 0, list: true},
			expected: "testdata/dir.log\n",
		},
		{
			name: "FilterExtensionSizeMatch", root: "testdata",
			cfg:      config{ext: ".log", size: 10, list: true},
			expected: "testdata/dir.log\n",
		},
		{
			name: "FilterExtensionSizeNoMatch", root: "testdata",
			cfg:      config{ext: ".log", size: 20, list: true},
			expected: "",
		},
		{
			name: "FilterExtensionNoMatch", root: "testdata",
			cfg:      config{ext: ".gz", size: 0, list: true},
			expected: "",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var buffer bytes.Buffer
			if err := run(testCase.root, &buffer, testCase.cfg); err != nil {
				t.Fatal(err)
			}
			res := buffer.String()
			if testCase.expected != res {
				t.Errorf("Expected '%q', got '%q' instead\n", testCase.expected, res)
			}
		})
	}
}

func TestRunDelExtension(t *testing.T) {
	testCases := []struct {
		name        string
		cfg         config
		extNoDelete string
		nDelete     int
		nNoDelete   int
		expected    string
	}{{
		name:        "DeleteExtensionNoMatch",
		cfg:         config{ext: ".log", del: true},
		extNoDelete: ".gz", nDelete: 0, nNoDelete: 10,
		expected: "",
	}, {
		name:        "DeleteExtensionMatch",
		cfg:         config{ext: ".log", del: true},
		extNoDelete: "", nDelete: 10, nNoDelete: 0,
		expected: "",
	}, {
		name:        "DeleteExtensionMixed",
		cfg:         config{ext: ".log", del: true},
		extNoDelete: ".gz", nDelete: 5, nNoDelete: 5,
		expected: "",
	}}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var (
				buffer    bytes.Buffer
				logBuffer bytes.Buffer
			)
			testCase.cfg.wLog = &logBuffer

			tempDir, cleanup := createTempDir(t, map[string]int{
				testCase.cfg.ext:     testCase.nDelete,
				testCase.extNoDelete: testCase.nNoDelete,
			})
			defer cleanup()
			if err := run(tempDir, &buffer, testCase.cfg); err != nil {
				t.Fatal(err)
			}
			res := buffer.String()
			if testCase.expected != res {
				t.Errorf("Expected '%q', got '%q' instead\n", testCase.expected, res)
			}
			fileLeft, err := os.ReadDir(tempDir)
			if err != nil {
				t.Error(err)
			}
			if len(fileLeft) != testCase.nNoDelete {
				t.Errorf("Expected '%d' files left, got '%d' instead\n", testCase.nNoDelete, len(fileLeft))
			}
			expLogLines := testCase.nDelete + 1
			lines := bytes.Split(logBuffer.Bytes(), []byte("\n"))
			if len(lines) != expLogLines {
				t.Errorf("Expected '%d' log lines, got '%d' instead\n", expLogLines, len(lines))
			}
		})
	}
}

func createTempDir(t *testing.T, files map[string]int) (dirname string, cleanup func()) {
	t.Helper()
	err := os.MkdirAll("walktest", 0o755)
	if err != nil {
		t.Fatal(err)
	}
	for k, n := range files {
		for j := 1; j <= n; j++ {
			fname := fmt.Sprintf("files%d%s", j, k)
			fpath := filepath.Join("walktest", fname)
			if err := os.WriteFile(fpath, []byte("dummy"), 0o644); err != nil {
				t.Fatal(err)
			}
		}
	}
	return "walktest", func() { os.RemoveAll("walktest") }
}
