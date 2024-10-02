package main

import (
	"strconv"
)

func StringBasic() {
	// declare string variable
	var string1 string
	string1 = "hi"
	string2 := "fork"

	// concate strings
	string3 := sf("%s %s", string1, string2)
	p("String:", string3)

	// convert string to number
	var (
		err           error
		int_value_1   int64
		float_value_1 float64
	)
	string1 = "5"
	string_5 := "5.5"
	int_value_1, err = strconv.ParseInt(string1, 10, 32) // base10, 32 size
	if err == nil {
		p("String to Int:", int_value_1)
	} else {
		p(err)
	}

	float_value_1, err = strconv.ParseFloat(string_5, 64) // 64size
	if err == nil {
		p("String to Float:", float_value_1)
	} else {
		p(err)
	}

	// convert numberic to string
	number_1 := 8
	number_2 := 8.8

	var string_number_1, string_number_2 string
	string_number_1 = sf("%d", number_1)
	p("String:", string_number_1)
	string_number_2 = sf("%f", number_2)
	p("String:", string_number_2)
}
