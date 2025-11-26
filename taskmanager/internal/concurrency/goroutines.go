package concurrency

import (
	"fmt"
	"time"
)

func RunGoRoutines() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Goroutine1 Count: %d\n", i)
			time.Sleep(300 * time.Millisecond)
		}
	}()
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Goroutine2 Count: %d\n", i)
			time.Sleep(300 * time.Millisecond)
		}
	}()
	time.Sleep(2 * time.Second)
}
