package basics

import "regexp"

func RegexExample() {
	// find AB000 pattern
	regexPattern, _ := regexp.Compile("[a-z][a-z][0-9]+")
	stringExample1 := "VASL908-S"
	pf("%s", regexPattern.FindString(stringExample1))

	// find word with -ing suffix
	regexPattern, _ = regexp.Compile(`(\w+)ing`)
	stringExample2 := "I am doing a good job. Please not inviting me to go out."
	pf("%s", regexPattern.FindString(stringExample2))
}
