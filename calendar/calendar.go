package calendar

import (
	"booking-calendar/schedule"
	"booking-calendar/utils"
	"time"
)

type Calendar struct {
	operatingSchedule map[string]schedule.DaySchedule
	appointments      []Appointment
}

func NewCalendar(operatingSchedule map[string]schedule.DaySchedule) Calendar {
	return Calendar{
		operatingSchedule: operatingSchedule,
		appointments:      make([]Appointment, 0),
	}
}

type CalendarQuery struct {
	Date     time.Time
	Duration time.Duration
}

func (query CalendarQuery) Day() string {
	return query.Date.Weekday().String()
}

func (calendar *Calendar) BookAppointment(appointment Appointment) bool {
	for _, existingAppointment := range calendar.appointments {
		if utils.ExtendTimeRange(existingAppointment).IsConflict(appointment) {
			return false
		}
	}

	calendar.appointments = append(calendar.appointments, appointment)

	return true
}

func (calendar *Calendar) CancelAppointment(startTime time.Time) {
	appointments := make([]Appointment, 0)
	for _, appointment := range calendar.appointments {
		if !appointment.StartTime().Equal(startTime) {
			appointments = append(appointments, appointment)
		}
	}

	calendar.appointments = appointments
}

func (calendar Calendar) CheckAvailability(query CalendarQuery) []utils.SimpleTimeRange {
	if daySchedule, isOperatingDay := calendar.operatingSchedule[query.Day()]; isOperatingDay {
		return calendar.compileAvailability(query, daySchedule)
	} else {
		return make([]utils.SimpleTimeRange, 0)
	}
}

func (calendar Calendar) compileAvailability(query CalendarQuery, daySchedule schedule.DaySchedule) []utils.SimpleTimeRange {
	availability := make([]utils.SimpleTimeRange, 0)

	timeSlotIterator := daySchedule.Iterator(query.Date, query.Duration)
	for timeSlotIterator.HasNext() {
		timeRange := timeSlotIterator.Next()

		if !calendar.isConflict(timeRange) {
			simpleTimeRange := utils.NewSimpleTimeRange(timeRange.StartTime(), timeRange.EndTime())
			availability = append(availability, simpleTimeRange)
		}
	}

	return availability
}

func (calendar Calendar) isConflict(timeRange utils.TimeRange) bool {
	for _, appointment := range calendar.appointments {
		if utils.ExtendTimeRange(appointment).IsConflict(timeRange) {
			return true
		}
	}

	return false
}
