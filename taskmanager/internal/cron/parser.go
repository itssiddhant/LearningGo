package cron

import (
	"strconv"
	"strings"
	"time"
)

type CronExpr struct {
	Minute int // -1 = wildcard
}

func Parse(expr string) CronExpr {
	parts := strings.Split(expr, " ")
	minPart := parts[0]

	if minPart == "*" || minPart == "*/1" {
		return CronExpr{Minute: -1}
	}

	v, _ := strconv.Atoi(minPart)
	return CronExpr{Minute: v}
}

func (c CronExpr) Matches(t time.Time) bool {
	if c.Minute == -1 {
		return true
	}
	return t.Minute() == c.Minute
}
