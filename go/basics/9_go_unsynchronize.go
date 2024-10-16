package basics

import (
	"time"
)

func GoUnsynchronize() {
	// channel := make(chan string)

	// run in backgroud - not working
	go func() {
		for index := 0; index < 8; index++ {
			P("Hello ", index)
			index++
			// channel <- fmt.Sprintf("Hello %d", index)
		}
	}()
	go Add(1, 5)
	// value := <-channel
	// fmt.Println(value)
	time.Sleep(time.Second)
	P("done")
}
