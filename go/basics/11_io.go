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
	err := input.Err()
	Check("Error when get user input", err)
}
