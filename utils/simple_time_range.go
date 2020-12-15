package utils

import (
	"bytes"
	"encoding/json"
	"time"
)

type SimpleTimeRange struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

func (simpleTimeRange *SimpleTimeRange) UnmarshalJSON(rawBytes []byte) error {
	var rawString map[string]string

	decodeError := json.NewDecoder(bytes.NewBuffer(rawBytes)).Decode(&rawString)
	if decodeError != nil {
		return decodeError
	}

	for key, value := range rawString {
		if key == "startTime" {
			if startTime, timeParseError := time.Parse(time.RFC3339, value); timeParseError == nil {
				simpleTimeRange.StartTime = startTime
			}
		}

		if key == "endTime" {
			if endTime, timeParseError := time.Parse(time.RFC3339Nano, value); timeParseError == nil {
				simpleTimeRange.EndTime = endTime
			}
		}
	}

	return nil
}

func NewSimpleTimeRange(startTime time.Time, endTime time.Time) SimpleTimeRange {
	return SimpleTimeRange{
		StartTime: startTime,
		EndTime:   endTime,
	}
}

func (simpleTimeRange SimpleTimeRange) Equals(timeRange TimeRange) bool {
	return simpleTimeRange.Start().Equal(timeRange.Start()) && simpleTimeRange.End().Equal(timeRange.End())
}

func (simpleTimeRange SimpleTimeRange) Start() time.Time {
	return simpleTimeRange.StartTime
}

func (simpleTimeRange SimpleTimeRange) End() time.Time {
	return simpleTimeRange.EndTime
}
