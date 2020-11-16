package utils

import "time"

type TimeRange interface {
	StartTime() time.Time
	EndTime() time.Time
}