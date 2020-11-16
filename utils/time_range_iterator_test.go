package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestSimpleTimeRangeIterator(t *testing.T) {
	startTime, err := time.Parse("2006-01-02", "2020-01-01")
	if err != nil {
		panic(err)
	}

	endTime, err := time.Parse("2006-01-02", "2020-01-10")
	if err != nil {
		panic(err)
	}

	delta, err := time.ParseDuration("24h")
	if err != nil {
		panic(err)
	}

	timeRange := NewSimpleTimeRange(startTime, endTime)

	actual := make([]SimpleTimeRange, 0)
	iterator := NewTimeRangeIterator(timeRange, delta)
	for iterator.HasNext() {
		slot := iterator.Next()

		fmt.Printf("%v to %v\n", slot.StartTime(), slot.EndTime())

		actual = append(actual, NewSimpleTimeRange(slot.StartTime(), slot.EndTime()))
	}

	expected := getExpected(startTime, endTime, delta)

	if len(expected) != len(actual) {
		t.Error("Actual is not the same size of the expected list")
	}

	for i:= 0; i < len(expected); i++ {
		expectedSlot := expected[i]
		actualSlot := actual[i]

		if !expectedSlot.Equals(actualSlot) {
			t.Fail()
		}
	}
}

func getExpected(starTime time.Time, endTime time.Time, duration time.Duration) []SimpleTimeRange {
	result := make([]SimpleTimeRange, 0)
	for iterator := starTime; iterator.Before(endTime); iterator = iterator.Add(duration) {
		result = append(result, NewSimpleTimeRange(iterator, iterator.Add(duration)))
	}

	return result
}
