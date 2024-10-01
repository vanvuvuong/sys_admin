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

	// array: fixed type & len
	var exampleArray1 [5]int
	exampleArray2 := [...]int{1: 5}
	fmt.Print(exampleArray1, exampleArray2)

	// two-dimentional array
	exampleArray3 := [2][3]int{
		{1: 3},
		{1: 3},
	}
	fmt.Print(exampleArray3)

	// slice: fixed type
	var exampleSlice []string
	fmt.Print(exampleSlice)

	// map: dict in Python
	exampleMap1 := make(map[string]string)
	exampleMap1["key"] = "value"
	fmt.Print(exampleMap1)
	exampleMap2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Print(exampleMap2)

	// struct ~ equal to dictionary in other language
	type example_struct_type struct {
		Id   int32
		Name string
	}
	exampleStruct := new(example_struct_type)
	fmt.Print(exampleStruct)
}
