package schedule

import (
	"booking-calendar/utils"
	"time"
)

type DaySchedule struct {
	startTime time.Time
	endTime   time.Time
}

func (daySchedule DaySchedule) Start() time.Time {
	return daySchedule.startTime
}

func (daySchedule DaySchedule) End() time.Time {
	return daySchedule.endTime
}

func (daySchedule DaySchedule) IteratorForDate(date time.Time, duration time.Duration) utils.TimeRangeIterator {
	// Adjust to the target day
	daySchedule.startTime = utils.JustifyTime(daySchedule.startTime, date)
	daySchedule.endTime = utils.JustifyTime(daySchedule.endTime, date)

	return utils.NewTimeRangeIterator(daySchedule, duration)
}

func NewDaySchedule(startTime time.Time, endTime time.Time) DaySchedule {
	return DaySchedule{
		startTime: startTime,
		endTime:   endTime,
	}
}

func CompileBusinessWeekSchedule(startTime time.Time, endTime time.Time) map[string]DaySchedule {
	return map[string]DaySchedule{
		time.Monday.String():    NewDaySchedule(startTime, endTime),
		time.Tuesday.String():   NewDaySchedule(startTime, endTime),
		time.Wednesday.String(): NewDaySchedule(startTime, endTime),
		time.Thursday.String():  NewDaySchedule(startTime, endTime),
		time.Friday.String():    NewDaySchedule(startTime, endTime),
		time.Saturday.String():  NewDaySchedule(startTime, endTime),
	}
}
