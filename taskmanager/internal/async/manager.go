package async

import (
	"context"
	"fmt"
)

type Manager struct {
	Queue   *JobQueue
	Results chan string
}

func NewManager() *Manager {
	return &Manager{
		Queue:   NewJobQueue(50),
		Results: make(chan string, 50),
	}
}
func (m *Manager) StartWorkers(ctx context.Context, numWorkers int) {
	for i := 1; i <= numWorkers; i++ {
		go StartWorker(ctx, i, m.Queue.Queue, m.Results)
	}
	fmt.Printf("Started %d workers\n", numWorkers)
}
func (m *Manager) HandleResults() {
	go func() {
		for res := range m.Results {
			fmt.Printf("Result received: %s\n", res)
		}
	}()
}
