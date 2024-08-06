package main

import "fmt"

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
	for key, value := range data {
		fmt.Println(key)
		fmt.Println(value)
	}
}
