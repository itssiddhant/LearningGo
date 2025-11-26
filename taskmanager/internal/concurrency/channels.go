package concurrency

import "fmt"

func BasicChannel() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from Channel!"
	}()

	msg := <-ch
	fmt.Println("Received: ", msg)
}
