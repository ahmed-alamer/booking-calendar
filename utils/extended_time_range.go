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
	onOrAfterStartTime := target.StartTime().After(extendedTimeRange.StartTime()) || target.StartTime().Equal(extendedTimeRange.StartTime())
	onOrBeforeEndTime := target.EndTime().Before(extendedTimeRange.EndTime()) || target.EndTime().Equal(extendedTimeRange.EndTime())

	return onOrAfterStartTime && onOrBeforeEndTime
}
