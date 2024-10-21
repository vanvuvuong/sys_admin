package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	testCases := [2]struct {
		name      string
		word      string
		count     int
		countLine bool
	}{
		{"CountWord", "word1 word2 word3 work4\n", 4, false},
		{"CountLine", "word1\nword2\nword3", 3, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := bytes.NewBufferString(tc.word)
			wc := count(res, tc.countLine)
			if wc != tc.count {
				t.Errorf("Expected %d, got %d instead.\n", tc.count, wc)
			}
		})
	}
}
