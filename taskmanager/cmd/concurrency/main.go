package main

import (
	"github.com/itssiddhant/taskmanager/internal/concurrency"
)

func main() {
	concurrency.BasicChannel()
	concurrency.BufferedChannel()
	concurrency.SelectDemo()
	concurrency.RunGoRoutines()
	concurrency.WorkerPool()
}
