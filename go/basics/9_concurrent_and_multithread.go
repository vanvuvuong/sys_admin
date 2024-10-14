package basics

import "time"

func process(index int) {
	pf("%d processing\n", index)
}

// Multi
func GoRoutine() {
	for index := 0; index < 10; index++ {
		go process(index)
	}
	p("Sleep 4 second.")
	time.Sleep(time.Second * 4)
}
