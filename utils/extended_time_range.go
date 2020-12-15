package utils

import "time"

type ExtendedTimeRange struct {
	base TimeRange
}

func (extendedTimeRange ExtendedTimeRange) Start() time.Time {
	return extendedTimeRange.base.Start()
}

func (extendedTimeRange ExtendedTimeRange) End() time.Time {
	return extendedTimeRange.base.End()
}

func ExtendTimeRange(base TimeRange) ExtendedTimeRange {
	return ExtendedTimeRange{
		base: base,
	}
}

func (extendedTimeRange ExtendedTimeRange) IsConflict(target TimeRange) bool {
	onOrAfterStartTime := target.Start().Equal(extendedTimeRange.Start()) || target.Start().After(extendedTimeRange.Start())
	onOrBeforeEndTime := target.End().Equal(extendedTimeRange.End()) || target.End().Before(extendedTimeRange.End())

	return onOrAfterStartTime && onOrBeforeEndTime
}

func (extendedTimeRange ExtendedTimeRange) IsTimeBetween(target time.Time) bool {
	onOrAfterStartTime := target.Equal(extendedTimeRange.Start()) || target.After(extendedTimeRange.Start())
	onOrBeforeEndTime := target.Equal(extendedTimeRange.End()) || target.Before(extendedTimeRange.End())

	return onOrAfterStartTime && onOrBeforeEndTime
}
