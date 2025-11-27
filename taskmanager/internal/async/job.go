package async

import "time"

type Job struct {
	ID        string
	Payload   string
	CreatedAt time.Time
	Attempts  int
	MaxRetry  int
}
