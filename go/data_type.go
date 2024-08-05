package main

import "fmt"

func define() {
	exampleChar := 'a'
	fmt.Print(exampleChar)

	exampleString := "string"
	fmt.Print(exampleString)

	exampleFloat := 1.11111
	fmt.Print(exampleFloat)
	exampleInt := 1
	fmt.Print(exampleInt)

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

	// struct ~ equal to dictionary in other language
	type example_struct_type struct {
		Id   int32
		Name string
	}
	exampleStruct := new(example_struct_type)
	fmt.Print(exampleStruct)
}
