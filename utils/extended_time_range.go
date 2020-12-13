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
	onOrAfterStartTime := target.StartTime().Equal(extendedTimeRange.StartTime()) || target.StartTime().After(extendedTimeRange.StartTime())
	onOrBeforeEndTime := target.EndTime().Equal(extendedTimeRange.EndTime()) || target.EndTime().Before(extendedTimeRange.EndTime())

	return onOrAfterStartTime && onOrBeforeEndTime
}

func (extendedTimeRange ExtendedTimeRange) IsTimeBetween(target time.Time) bool {
	onOrAfterStartTime := target.Equal(extendedTimeRange.StartTime()) || target.After(extendedTimeRange.StartTime())
	onOrBeforeEndTime := target.Equal(extendedTimeRange.EndTime()) || target.Before(extendedTimeRange.EndTime())

	return onOrAfterStartTime && onOrBeforeEndTime
}
