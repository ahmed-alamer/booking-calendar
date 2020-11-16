package utils

import "time"

type SimpleTimeRange struct {
	startTime time.Time
	endTime   time.Time
}

func NewSimpleTimeRange(startTime time.Time, endTime time.Time) SimpleTimeRange {
	return SimpleTimeRange{
		startTime: startTime,
		endTime:   endTime,
	}
}

func (simpleTimeRange SimpleTimeRange) StartTime() time.Time {
	return simpleTimeRange.startTime
}

func (simpleTimeRange SimpleTimeRange) EndTime() time.Time {
	return simpleTimeRange.endTime
}

func (simpleTimeRange SimpleTimeRange) Equals(timeRange TimeRange) bool {
	return simpleTimeRange.StartTime().Equal(timeRange.StartTime()) && simpleTimeRange.EndTime().Equal(timeRange.EndTime())
}
