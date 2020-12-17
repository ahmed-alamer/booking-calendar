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
	api.POST("/:providerId", func(context echo.Context) error {
		request, parseError := rpc.ParseRequest(context)
		if parseError != nil {
			return errorResponse(context, rpc.ErrorResponse(request, rpc.InvalidRequest, parseError))
		}

		providerId := context.Param("providerId")
		if providerId == "" {
			return errorResponse(context, rpc.ErrorResponse(request, rpc.RequestMissingProviderId, errors.New("missing provider id in the request path")))
		}

		provider, exists := providers[providerId]
		if !exists {
			return errorResponse(context, rpc.ErrorResponse(request, rpc.ProviderNotFound, errors.New("provider was not found")))
		}

		providerCalendar := calendars[provider.Id]
		response := providerCalendar.Execute(request)

		return context.JSON(http.StatusOK, response)
	})

	api.Logger.Fatal(api.Start(":8080"))
}

func errorResponse(context echo.Context, response rpc.Response) error {
	return context.JSON(http.StatusInternalServerError, response)
}
