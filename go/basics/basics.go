package basics

import (
	"fmt"
	"log"
)

var (
	p  = fmt.Println
	pf = fmt.Printf
	sf = fmt.Sprintf
	l  = log.Println
)

func Check(prefix string, err error) {
	if err != nil {
		l(prefix, ": ", err)
		panic(err)
	}
}

func main() {
	// Define()
	// AdvancedDefine()
	// StringBasic()
	// Condition()
	// Loop()
	// RangeOverData()
	// GoUnsynchronize()
	// InputScanner()
}
