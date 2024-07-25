package main

import (
	"fmt"
	"strconv"
	"strings"
)

func string_manipulator() {
	// declare string variable
	var string_1 string
	string_2 := "hi"
	string_3 := "folk"

	// concate strings
	string_4 := fmt.Sprintf("%s %s", string_2, string_3)
	fmt.Println(string_4)

	// convert string to number
	var (
		err           error
		int_value_1   int64
		float_value_1 float64
	)
	string_1 = "5"
	string_5 := "5.5"
	int_value_1, err = strconv.ParseInt(string_1, 10, 32) // base10, 32 size
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
	fmt.Print(string_number_2)
	fmt.Print()

	// split string
	data := "name;color;address"
	nouns := strings.Split(data, ";")
	for _, noun := range nouns {
		fmt.Println(noun)
	}

	// check data len
	len_data := len(data)
	fmt.Printf("len=%d \n", len_data)

	// upper
	fmt.Println(strings.ToUpper(data))
}
