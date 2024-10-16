package basics

import (
	"fmt"
	s "strings"
)

func StringFunction() {
	// split string
	data := "name;color;address"
	nouns := s.Split(data, ";")
	for _, noun := range nouns {
		fmt.Println(noun)
	}

	// string functions
	P("Contains: ", s.Contains(data, ";"))
	P("Count: ", s.Count(data, ";"))
	P("HasPrefix: ", s.HasPrefix(data, "na"))
	P("HasSuffix: ", s.HasSuffix(data, "na"))
	P("Index: ", s.Index(data, "na"))
	P("Join: ", s.Join(nouns[1:], ","))
	P("Replace: ", s.Replace(data, ";", ",", -1))
	P("Split: ", s.Split(data, ";"))
	P("ToUpper: ", s.ToUpper(data))
	P("ToLower: ", s.ToLower(data))
}

func StringEncoding() {
}
