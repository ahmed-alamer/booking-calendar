package utils

import "time"

type ExtendedTimeRange struct {
	base TimeRange
}

func (extendedTimeRange ExtendedTimeRange) StartTime() time.Time {
	return extendedTimeRange.base.StartTime()
}

func (extendedTimeRange ExtendedTimeRange) EndTime() time.Time {
	return extendedTimeRange.base.EndTime()
}

func ExtendTimeRange(base TimeRange) ExtendedTimeRange {
	return ExtendedTimeRange{
		base: base,
	}
}

func (extendedTimeRange ExtendedTimeRange) IsConflict(target TimeRange) bool {
	return target.StartTime().After(extendedTimeRange.StartTime()) && target.EndTime().Before(extendedTimeRange.EndTime())
}