package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)
	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	}
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}
	if countLines {
		scanner.Split(bufio.ScanLines)
	}
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}

func main() {
	lines := flag.Bool("l", false, "Count lines")
	wordBytes := flag.Bool("b", false, "Count numbers of bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *wordBytes))
}
