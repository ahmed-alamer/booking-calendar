package calendar

import "time"

type Appointment struct {
	Client  string    `json:"client"`
	Purpose string    `json:"purpose"`
	Start   time.Time `json:"startTime"`
	End     time.Time `json:"endTime"`
}

func (appointment Appointment) StartTime() time.Time {
	return appointment.Start
}

func (appointment Appointment) EndTime() time.Time {
	return appointment.End
}
