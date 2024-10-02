package main

import (
	"time"
)

// Condition controller
func Condition() {
	if 8%2 == 0 || 7%2 == 0 {
		p("Either 8 or 7 are even")
	}

	if 7%2 == 0 {
		p("7 is even")
	} else {
		p("7 is odd")
	}

	num := 9
	if num < 0 {
		p(num, "is negative")
	} else if num < 10 {
		p(num, "has 1 digit")
	} else {
		p(num, "has multiple digits")
	}
}

// Loop controller
func Loop() {
	for j := 0; j < 3; j++ {
		p("Oh, a for loop", j)
	}

	i := 1
	for i <= 3 {
		p("Oh, another for loop", i)
		i = i + 1
	}

	var data []map[string]string
	data = append(data, map[string]string{"key1": "value1", "key2": "value2"})
	for _, item := range data {
		for key, value := range item {
			pf("%s -> %s\n", key, value)
		}
	}

	today := time.Now().Weekday().String()
	switch today {
	case "Sunday":
		p("Here we go")
	case "Monday":
		p("Just started")
	default:
		p("A switch example:", today)
	}
}
