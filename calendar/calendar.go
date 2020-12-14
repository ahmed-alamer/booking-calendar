package calendar

import (
	"booking-calendar/schedule"
	"booking-calendar/utils"
	"errors"
	"github.com/google/uuid"
	"time"
)

type Calendar struct {
	operatingSchedule map[string]schedule.DaySchedule
	appointments      map[uuid.UUID]Appointment
}

func NewCalendar(operatingSchedule map[string]schedule.DaySchedule) Calendar {
	return Calendar{
		operatingSchedule: operatingSchedule,
		appointments:      make(map[uuid.UUID]Appointment),
	}
}

func (calendar Calendar) GetAppointments(timeRange utils.TimeRange) []Appointment {
	result := make([]Appointment, 0)
	for _, appointment := range calendar.appointments {
		if utils.ExtendTimeRange(timeRange).IsTimeBetween(appointment.StartTime) {
			result = append(result, appointment)
		}
	}

	return result
}

func (calendar *Calendar) BookAppointment(appointment Appointment) (uuid.UUID, error) {
	for _, existingAppointment := range calendar.appointments {
		if utils.ExtendTimeRange(existingAppointment).IsConflict(appointment) {
			return uuid.UUID{}, errors.New("conflict")
		}
	}

	if id, uuidError := uuid.NewRandom(); uuidError == nil {
		calendar.appointments[id] = appointment
		return id, nil
	} else {
		return uuid.UUID{}, uuidError
	}
}

func (calendar *Calendar) CancelAppointment(id uuid.UUID) bool {
	if _, exists := calendar.appointments[id]; exists {
		delete(calendar.appointments, id) // I know this is a no-op if the key doesn't exist, returning an error is more helpful in debugging
		return true
	} else {
		return false
	}
}

func (calendar Calendar) CheckAvailability(query Query) []utils.SimpleTimeRange {
	if daySchedule, isOperatingDay := calendar.GetDaySchedule(query.Date); isOperatingDay {
		return calendar.compileAvailability(query, daySchedule)
	} else {
		return make([]utils.SimpleTimeRange, 0)
	}
}

func (calendar Calendar) GetDaySchedule(date time.Time) (schedule.DaySchedule, bool) {
	if daySchedule, isOperatingDay := calendar.operatingSchedule[date.Weekday().String()]; isOperatingDay {
		startTime := utils.JustifyTime(daySchedule.Start(), date)
		endTime := utils.JustifyTime(daySchedule.End(), date)

		return schedule.NewDaySchedule(startTime, endTime), true
	} else {
		return schedule.DaySchedule{}, false
	}
}

func (calendar Calendar) compileAvailability(query Query, daySchedule schedule.DaySchedule) []utils.SimpleTimeRange {
	availability := make([]utils.SimpleTimeRange, 0)

	timeSlotIterator := utils.NewTimeRangeIterator(daySchedule, query.Duration)
	for timeSlotIterator.HasNext() {
		timeRange := timeSlotIterator.Next()

		if !calendar.isConflict(timeRange) {
			simpleTimeRange := utils.NewSimpleTimeRange(timeRange.Start(), timeRange.End())
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
