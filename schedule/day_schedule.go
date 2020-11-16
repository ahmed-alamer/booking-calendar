package schedule

import (
	"booking-calendar/utils"
	"time"
)

type DaySchedule struct {
	startTime time.Time
	endTime   time.Time
}

func (daySchedule DaySchedule) StartTime() time.Time {
	return daySchedule.startTime
}

func (daySchedule DaySchedule) EndTime() time.Time {
	return daySchedule.endTime
}

func (daySchedule DaySchedule) Iterator(startTime time.Time, duration time.Duration) utils.TimeRangeIterator {
	// Adjust to the target day
	daySchedule.startTime = utils.AdjustTime(startTime, daySchedule.startTime)
	daySchedule.endTime = utils.AdjustTime(startTime, daySchedule.endTime)

	return utils.NewTimeRangeIterator(daySchedule, duration)
}


func NewDaySchedule(startTime time.Time, endTime time.Time) DaySchedule {
	return DaySchedule{
		startTime: startTime,
		endTime:   endTime,
	}
}
