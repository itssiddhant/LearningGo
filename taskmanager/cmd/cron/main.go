package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/itssiddhant/taskmanager/internal/async"
	"github.com/itssiddhant/taskmanager/internal/cron"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	manager := async.NewManager()
	manager.StartWorkers(ctx, 3)
	manager.HandleResults()

	scheduler := cron.NewCronScheduler(manager)

	scheduler.AddCron("heartbeat", "*", "send-heartbeat")
	scheduler.AddCron("db-cleaner", "30", "clean-database")

	go scheduler.Start(ctx)

	// handle shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	cancel()
	time.Sleep(1 * time.Second)
}
