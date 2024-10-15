package basics

import (
	"strconv"
)

func StringBasic() {
	// declare string variable
	var string1 string
	string1 = "hi"
	string2 := "fork"

	// concate strings
	string3 := Sf("%s %s", string1, string2)
	P("String:", string3)

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
		P("String to Int:", int_value_1)
	} else {
		P(err)
	}

	float_value_1, err = strconv.ParseFloat(string_5, 64) // 64size
	if err == nil {
		P("String to Float:", float_value_1)
	} else {
		P(err)
	}

	// convert numberic to string
	number_1 := 8
	number_2 := 8.8

	var string_number_1, string_number_2 string
	string_number_1 = Sf("%d", number_1)
	P("String:", string_number_1)
	string_number_2 = Sf("%f", number_2)
	P("String:", string_number_2)
}
