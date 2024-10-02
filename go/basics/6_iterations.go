package main

import "fmt"

func rangeOverData() {
	// loop over array
	nums := [...]int{1: 10}
	for index, num := range nums {
		fmt.Printf("%3d: %-3d\n", index, num)
	}

	// loop over map
	exampleMap := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range exampleMap {
		fmt.Printf("%s -> %10s\n", key, value)
	}
	for key := range exampleMap {
		fmt.Printf("key: %s\n", key)
	}
}
