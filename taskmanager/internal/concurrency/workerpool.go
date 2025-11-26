package concurrency

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
		results <- job * 2
	}
}

func WorkerPool() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start 3 workers
	for w := 1; w < 4; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for r := 1; r <= 5; r++ {
		fmt.Println("Result:", <-results)
	}
}
