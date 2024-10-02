package main

import (
	"fmt"
	"strconv"
)

func stringBasic() {
	// declare string variable
	var string1 string
	string1 = "hi"
	string2 := "fork"

	// concate strings
	string3 := fmt.Sprintf("%s %s", string1, string2)
	fmt.Println(string3)

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
		fmt.Println(int_value_1)
	} else {
		fmt.Println(err)
	}

	float_value_1, err = strconv.ParseFloat(string_5, 64) // 64size
	if err == nil {
		fmt.Println(float_value_1)
	} else {
		fmt.Println(err)
	}

	// convert numberic to string
	number_1 := 8
	number_2 := 8.8

	var string_number_1, string_number_2 string
	string_number_1 = fmt.Sprintf("%d", number_1)
	fmt.Println(string_number_1)
	string_number_2 = fmt.Sprintf("%f", number_2)
	fmt.Println(string_number_2)
}
