package utils

import (
	"time"
)


type TimeRangeIterator struct {
	pointer time.Time
	endTime time.Time
	delta   time.Duration
}

func NewTimeRangeIterator(timeRange TimeRange, delta time.Duration) TimeRangeIterator {
	return TimeRangeIterator{
		pointer: timeRange.StartTime(),
		endTime: timeRange.EndTime(),
		delta:   delta,
	}
}

func (iterator TimeRangeIterator) HasNext() bool {
	return iterator.pointer.Before(iterator.endTime)
}

func (iterator *TimeRangeIterator) Next() TimeRange {
	timeRange := NewSimpleTimeRange(iterator.pointer, iterator.pointer.Add(iterator.delta))

	iterator.pointer = iterator.pointer.Add(iterator.delta)

	return timeRange
}
