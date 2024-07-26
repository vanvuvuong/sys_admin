package main

import (
	"fmt"
	"time"
)

func main() {
	// channel := make(chan string)

	// run in backgroud - not working
	go func() {
		for index := 0; index < 3; index++ {
			fmt.Println("Hello ", index)
			index++
			// channel <- fmt.Sprintf("Hello %d", index)
		}
	}()
	go Call()
	// value := <-channel
	// fmt.Println(value)
	time.Sleep(time.Second)
	fmt.Println("done")
}
