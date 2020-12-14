package utils

import "time"

type TimeRange interface {
	Start() time.Time
	End() time.Time
}
