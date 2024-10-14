package main

import (
	"fmt"
	"log"

	"github.com/go/basics"
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
		// panic(err) - will give you the same message
	}
}

func main() {
	basics.AdvancedDefine()
	basics.GoRoutine()
	// GetRequest()
	// DecodeCsv("decode_example.csv", ';')
}
