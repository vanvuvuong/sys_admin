package basics

import (
	"fmt"
	"log"
)

var (
	P  = fmt.Println
	Pf = fmt.Printf
	Sf = fmt.Sprintf
	L  = log.Println
)

func Check(prefix string, err error) {
	if err != nil {
		L(prefix, ": ", err)
		// panic(err) - will give you the same message
	}
}
