package main

import (
	"fmt"
	"time"
)

func goroutine() {
	index := 0
	channel := make(chan string)

	// run in backgroud - not working
	go func() {
		for index < 10 {
			fmt.Println("Hello ", index)

			time.Sleep(500 * time.Millisecond)
			index++

			channel <- fmt.Sprintf("Hello %d", index)
		}
	}()
	Call()
	value := <-channel
	fmt.Println(value)
}
