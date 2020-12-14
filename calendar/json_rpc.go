package calendar

import (
	"booking-calendar/rpc"
	"encoding/json"
)

const (
	BookAppointment   rpc.Method = "bookAppointment"
	CancelAppointment rpc.Method = "cancelAppointment"
	CheckAvailability rpc.Method = "checkAvailability"
	GetAppointments   rpc.Method = "getAppointments"
)

func (calendar Calendar) Execute(request rpc.Request) rpc.Response {
	switch request.Method {
	case BookAppointment:
		appointment := &Appointment{}
		if parseError := json.Unmarshal([]byte(request.Method), appointment); parseError == nil {
			return rpc.Success(request, appointment)
		} else {
			return rpc.ErrorResponse(request, rpc.InvalidParams, parseError)
		}
	case CancelAppointment:
		return rpc.MethodNotFoundError(request)
	case CheckAvailability:
		return rpc.MethodNotFoundError(request)
	case GetAppointments:
		return rpc.MethodNotFoundError(request)
	default:
		return rpc.MethodNotFoundError(request)
	}
}
