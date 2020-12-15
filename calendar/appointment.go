package calendar

import "time"

type Appointment struct {
	Client    string    `json:"client"`
	Purpose   string    `json:"purpose"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

func (appointment Appointment) Start() time.Time {
	return appointment.StartTime
}

func (appointment Appointment) End() time.Time {
	return appointment.EndTime
}
