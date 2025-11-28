package cron

import "time"

type CronJob struct {
	ID       string
	Expr     CronExpr
	TaskName string
	LastRun  time.Time
	Enabled  bool
}
