package cron

import (
	"context"
	"fmt"
	"time"

	"github.com/itssiddhant/taskmanager/internal/async"
)

type CronScheduler struct {
	CronJobs []CronJob
	AsyncMgr *async.Manager
}

func NewCronScheduler(asyncMgr *async.Manager) *CronScheduler {
	return &CronScheduler{
		AsyncMgr: asyncMgr,
	}
}

func (c *CronScheduler) AddCron(id, expr, task string) {
	c.CronJobs = append(c.CronJobs, CronJob{
		ID:       id,
		Expr:     Parse(expr),
		TaskName: task,
		Enabled:  true,
	})
}

func (c *CronScheduler) Start(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cron scheduler shutting down")
			return

		case now := <-ticker.C:
			c.tick(now)
		}
	}
}

func (c *CronScheduler) tick(now time.Time) {
	for i := range c.CronJobs {
		job := &c.CronJobs[i]

		if !job.Enabled {
			continue
		}

		if job.Expr.Matches(now) && now.Sub(job.LastRun) > time.Minute {
			job.LastRun = now

			c.AsyncMgr.Queue.Push(async.Job{
				ID:        fmt.Sprintf("cron-%s-%d", job.ID, now.Unix()),
				Payload:   job.TaskName,
				CreatedAt: now,
				MaxRetry:  2,
			})

			fmt.Println("Triggered cron:", job.ID)
		}
	}
}
