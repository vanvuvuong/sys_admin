package basics

import "time"

func process(index int) {
	Pf("%d processing\n", index)
}

// Multi
func GoRoutine() {
	for index := 0; index < 10; index++ {
		go process(index)
	}
	P("Sleep 4 second.")
	time.Sleep(time.Second * 4)
}
