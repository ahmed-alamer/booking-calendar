package utils

import (
	"booking-calendar/schedule"
	"time"
)

func CompileBusinessWeekSchedule(startTime time.Time, endTime time.Time) map[string]schedule.DaySchedule {
	return map[string]schedule.DaySchedule{
		time.Monday.String(): schedule.NewDaySchedule(startTime, endTime),
		time.Tuesday.String(): schedule.NewDaySchedule(startTime, endTime),
		time.Wednesday.String(): schedule.NewDaySchedule(startTime, endTime),
		time.Thursday.String(): schedule.NewDaySchedule(startTime, endTime),
		time.Friday.String(): schedule.NewDaySchedule(startTime, endTime),
		time.Saturday.String(): schedule.NewDaySchedule(startTime, endTime),
	}
}
