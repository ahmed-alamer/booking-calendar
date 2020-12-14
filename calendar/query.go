package calendar

import (
	"bytes"
	"encoding/json"
	"time"
)

type Query struct {
	Date     time.Time     `json:"date"`
	Duration time.Duration `json:"duration"`
}

func (query *Query) UnmarshalJSON(rawBytes []byte) error {
	var rawString = make(map[string]string)

	decodeError := json.NewDecoder(bytes.NewBuffer(rawBytes)).Decode(&rawString)
	if decodeError != nil {
		return decodeError
	}

	for key, value := range rawString {
		if key == "date" {
			if date, timeParseError := time.Parse(time.RFC3339, value); timeParseError == nil {
				query.Date = date
			} else {
				return timeParseError
			}
		} else if key == "duration" {
			if duration, durationParseError := time.ParseDuration(value); durationParseError == nil {
				query.Duration = duration
			} else {
				return durationParseError
			}
		}
	}

	return nil
}
