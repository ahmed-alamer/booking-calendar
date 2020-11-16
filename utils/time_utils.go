package utils

import "time"

func AdjustTime(source time.Time, adjustment time.Time) time.Time {
	return time.Date(adjustment.Year(), adjustment.Month(), adjustment.Day(), source.Hour(), source.Minute(), 0, 0, time.UTC)
}
