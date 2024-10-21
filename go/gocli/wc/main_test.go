package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	testCases := [3]struct {
		name          string
		word          string
		expectedCount int
		countLine     bool
		countBytes    bool
	}{
		{"CountWord", "word1 word2 word3 work4\n", 4, false, false},
		{"CountLine", "word1\nword2\nword3", 3, true, false},
		{"CountBytes", "word1\nword2\nword3", 17, false, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := bytes.NewBufferString(tc.word)
			wc := count(res, tc.countLine, tc.countBytes)
			if wc != tc.expectedCount {
				t.Errorf("Expected %d, got %d instead.\n", tc.expectedCount, wc)
			}
		})
	}
}
