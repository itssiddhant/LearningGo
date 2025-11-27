package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/itssiddhant/taskmanager/internal/async"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	manager := async.NewManager()
	manager.StartWorkers(ctx, 3)
	manager.HandleResults()

	// Create jobs
	for i := 1; i <= 10; i++ {
		manager.Queue.Push(async.Job{
			ID:        fmt.Sprintf("job-%d", i),
			Payload:   "data",
			CreatedAt: time.Now(),
			MaxRetry:  3,
		})
	}

	// Listen for Ctrl+C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	cancel()

	time.Sleep(1 * time.Second) // let workers finish
}
