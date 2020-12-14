package main

import (
	"booking-calendar/calendar"
	"booking-calendar/metadata"
	"booking-calendar/rpc"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	providers := make(map[string]metadata.Provider)
	calendars := make(map[string]calendar.Calendar)

	api := echo.New()
	api.GET("/", func(context echo.Context) error {
		if request, parseError := rpc.ParseRequest(context); parseError == nil {
			if providerId := context.QueryParam("provider"); providerId != "" {
				if provider, exists := providers[providerId]; exists {
					providerCalendar := calendars[provider.Id]
					response := providerCalendar.Execute(request)

					return context.JSON(http.StatusOK, response)
				} else {
					return context.JSON(http.StatusNotFound, rpc.ErrorResponse(request, rpc.ProviderNotFound, errors.New("provider was not found")))
				}
			} else {
				return context.JSON(http.StatusInternalServerError, rpc.ErrorResponse(request, rpc.RequestMissingProviderId, errors.New("missing provider id in the request path")))
			}
		} else {
			return context.JSON(http.StatusInternalServerError, rpc.ErrorResponse(request, rpc.InvalidRequest, parseError))
		}
	})

	api.Logger.Fatal(api.Start(":8080"))
}
