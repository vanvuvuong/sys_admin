package basics

import (
	"bufio"
	"os"
)

func InputScanner() {
	pf("Get user input\n->")
	input := bufio.NewScanner(os.Stdin)
	lines := 0
	for input.Scan() {
		lines++
	}
	pf("There are %d line(s)\n", lines)
	if err := input.Err(); err != nil {
		p("Error when get input", err)
	}
}
