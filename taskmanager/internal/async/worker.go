package async

import (
	"context"
	"fmt"
	"time"
)

func StartWorker(ctx context.Context, id int, jobs <-chan Job, results chan<- string) {
	fmt.Printf("Worker %d: Starting\n", id)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Stopping\n", id)
			return
		case job := <-jobs:
			fmt.Printf("Worker %d: Processing job %s\n", id, job.ID)
			// Simulate job processing time
			time.Sleep(2 * time.Second)
			fmt.Printf("Worker %d: Finished job %s\n", id, job.ID)

			results <- "done: " + job.ID
		}
	}
}
