package main

import (
	"fmt"
)

func define() {
	// char variable example
	exampleChar := 'a'
	fmt.Print(exampleChar)

	// string variable example
	exampleString := "string"
	fmt.Print(exampleString)

	// number variable example
	exampleFloat := 1.11111
	fmt.Print(exampleFloat)
	exampleInt := 1
	fmt.Print(exampleInt)

	// multiple variables declaration
	exampleInt1, exampleInt2 := 1, 2
	fmt.Print(exampleInt1, exampleInt2)

	// array: fixed type & length
	var exampleArray1 [5]int
	exampleArray2 := [...]int{1: 5}
	fmt.Print(exampleArray1, exampleArray2)

	// two-dimentional array
	exampleArray3 := [2][3]int{
		{1: 3},
		{1: 3},
	}
	fmt.Print(exampleArray3)

	// slice: fixed type, dynamic length
	var exampleSlice []string
	fmt.Print(exampleSlice)

	// map: like dictionary in Python
	exampleMap1 := make(map[string]string)
	exampleMap1["key"] = "value"
	fmt.Print(exampleMap1)
	exampleMap2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Print(exampleMap2)
}
