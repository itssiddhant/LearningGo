package concurrency

import "fmt"

func SelectDemo() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { ch1 <- "From 1" }()
	go func() { ch2 <- "From 2" }()

	select {
	case msg := <-ch1:
		fmt.Println("Got:", msg)
	case msg := <-ch2:
		fmt.Println("Got:", msg)
	}
}
