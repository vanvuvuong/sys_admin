package main

import (
	"fmt"
	"log"
)

var (
	p  = fmt.Println
	pf = fmt.Printf
	sf = fmt.Sprintf
	l  = log.Println
	lf = log.Fatalln
)

func Check(prefix string, err error) {
	if err != nil {
		l(prefix, ": ", err)
		panic(err)
	}
}

func main() {
	// GetRequest()
	// DecodeCsv("decode_example.csv", ';')
}
