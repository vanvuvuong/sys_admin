package main

import "fmt"

// Condition controller
func condition() {
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	num := 9
	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Loop controller
func loop() {
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	var data []map[string]string
	data = append(data, map[string]string{"key1": "value1", "key2": "value2"})
	for key, value := range data {
		fmt.Println(key)
		fmt.Println(value)
	}
}
