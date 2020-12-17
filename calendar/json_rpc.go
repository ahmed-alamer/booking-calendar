package calendar

import (
	"booking-calendar/rpc"
	"booking-calendar/utils"
	"errors"
	"github.com/google/uuid"
)

const (
	BookAppointment   rpc.Method = "bookAppointment"
	CancelAppointment rpc.Method = "cancelAppointment"
	CheckAvailability rpc.Method = "checkAvailability"
	GetAppointments   rpc.Method = "getAppointments"
)

func (calendar Calendar) Execute(request rpc.Request) rpc.Response {
	switch request.Method {
	case GetAppointments:
		return getAppointments(request, calendar)
	case CheckAvailability:
		return checkAvailability(request, calendar)
	case BookAppointment:
		return bookAppointment(request)
	case CancelAppointment:
		return cancelAppointment(request, calendar)
	default:
		return rpc.MethodNotFoundError(request)
	}
}

func getAppointments(request rpc.Request, calendar Calendar) rpc.Response {
	var timeRange utils.SimpleTimeRange
	if jsonParseError := request.ParseParams(&timeRange); jsonParseError != nil {
		return rpc.ErrorResponse(request, rpc.ParseError, jsonParseError)
	}

	appointments := calendar.GetAppointments(timeRange)
	return rpc.Response{
		Id:     request.Id,
		Result: appointments,
	}
}

func checkAvailability(request rpc.Request, calendar Calendar) rpc.Response {
	var query Query
	jsonParseError := request.ParseParams(&query)
	if jsonParseError != nil {
		return rpc.ErrorResponse(request, rpc.ParseError, jsonParseError)
	}

	availability := calendar.CheckAvailability(query)

	return rpc.Response{
		Id:     request.Id,
		Result: availability,
	}
}

func bookAppointment(request rpc.Request) rpc.Response {
	appointment := &Appointment{}
	if parseError := request.ParseParams(&appointment); parseError != nil {
		return rpc.ErrorResponse(request, rpc.InvalidParams, parseError)
	}

	return rpc.Success(request, appointment)
}

func cancelAppointment(request rpc.Request, calendar Calendar) rpc.Response {
	params := &struct {
		Id string `json:"id"`
	}{}

	if jsonParseError := request.ParseParams(&params); jsonParseError != nil {
		return rpc.ErrorResponse(request, rpc.ParseError, jsonParseError)
	}

	id, uuidParseError := uuid.Parse(params.Id)
	if uuidParseError != nil {
		return rpc.ErrorResponse(request, rpc.ParseError, uuidParseError)
	}

	if removed := calendar.CancelAppointment(id); !removed {
		return rpc.ErrorResponse(request, rpc.AppointmentWasNotFound, errors.New("appointment was not found"))
	}

	return rpc.Response{
		Id:     request.Id,
		Result: "Appointment was cancelled",
	}
}
