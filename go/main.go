package main

import (
	"github.com/gosysadmin/basics"
	"github.com/gosysadmin/gosysadmin"
)

func main() {
	basics.AdvancedDefine()
	basics.GoRoutine()
	gosysadmin.DecodeCsv("example.csv", ';')
}
