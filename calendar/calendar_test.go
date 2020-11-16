package calendar

import (
	"booking-calendar/utils"
	"log"
	"testing"
	"time"
)

func TestCalendar_CheckAvailability(t *testing.T) {
	operatingSchedule := utils.CompileBusinessWeekSchedule(parseTime("9:00AM"), parseTime("05:00PM"))
	providerCalendar := NewCalendar(operatingSchedule)

	availability := providerCalendar.CheckAvailability(CalendarQuery{
		Date:     parseDate("2020-11-16"),
		Duration: time.Hour,
	})

	for _, simpleTimeRange := range availability {
		log.Printf("%v to %v", simpleTimeRange.StartTime(), simpleTimeRange.EndTime())
	}

	if len(availability) != 9 {
		t.Error("Schedule should fully available for the day")
	}
}

func TestCalendar_BookAppointment(t *testing.T) {
	operatingSchedule := utils.CompileBusinessWeekSchedule(parseTime("9:00AM"), parseTime("05:00PM"))
	providerCalendar := NewCalendar(operatingSchedule)

	providerCalendar.BookAppointment(Appointment{
		Client:  "C1",
		Purpose: "Test",
		Start:   parseDateTime("2020-11-16T11:00"),
		End:     parseDateTime("2020-11-16T12:00"),
	})

	if len(providerCalendar.appointments) != 1 {
		t.Fail()
	}
}

func parseTime(timeStr string) time.Time {
	if parsedTime, err := time.Parse(time.Kitchen, timeStr); err == nil {
		return parsedTime
	} else {
		panic(err)
	}
}

func parseDate(dateStr string) time.Time {
	if parsedDate, err := time.Parse("2006-01-02", dateStr); err == nil {
		return parsedDate
	} else {
		panic(err)
	}
}

func parseDateTime(dateTimeStr string) time.Time {
	if parsedDate, err := time.Parse("2006-01-02T15:04", dateTimeStr); err == nil {
		return parsedDate
	} else {
		panic(err)
	}
}
