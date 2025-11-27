package async

type JobQueue struct {
	Queue chan Job
}

func NewJobQueue(size int) *JobQueue {
	return &JobQueue{
		Queue: make(chan Job, size),
	}
}

func (q *JobQueue) Push(job Job) {
	q.Queue <- job
}
