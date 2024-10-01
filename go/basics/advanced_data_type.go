package main

import "fmt"

func advancedDefine() {
	// pointer:
	// struct: more likely to class in OOP
	type example_struct_type struct {
		Id   int32
		Name string
	}
	exampleStruct := new(example_struct_type)
	fmt.Print(exampleStruct)
	// method:
	// interface:
	// enum:
	// struct embedding:
	// generic:
}
