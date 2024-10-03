package basics

import (
	"fmt"
	s "strings"
)

func stringFunction() {
	// split string
	data := "name;color;address"
	nouns := s.Split(data, ";")
	for _, noun := range nouns {
		fmt.Println(noun)
	}

	// string functions
	p("Contains: ", s.Contains(data, ";"))
	p("Count: ", s.Count(data, ";"))
	p("HasPrefix: ", s.HasPrefix(data, "na"))
	p("HasSuffix: ", s.HasSuffix(data, "na"))
	p("Index: ", s.Index(data, "na"))
	p("Join: ", s.Join(nouns[1:], ","))
	p("Replace: ", s.Replace(data, ";", ",", -1))
	p("Split: ", s.Split(data, ";"))
	p("ToUpper: ", s.ToUpper(data))
	p("ToLower: ", s.ToLower(data))
}

func stringEncoding() {
}
