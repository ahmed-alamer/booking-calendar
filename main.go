package main

import (
	"booking-calendar/calendar"
	"booking-calendar/metadata"
	"booking-calendar/rpc"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	providers := make(map[uuid.UUID]metadata.Provider)
	calendars := make(map[uuid.UUID]calendar.Calendar)

	api := echo.New()

	api.GET("/providers", func(context echo.Context) error {
		if len(providers) == 0 {
			return context.JSON(http.StatusNotFound, map[string]string{"error": "no providers"})
		} else {
			return context.JSON(http.StatusOK, providers)
		}
	})

	api.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, calendar.ApiDocumentation())
	})

	api.POST("/:providerId", func(context echo.Context) error {
		request, parseError := rpc.ParseRequest(context)
		if parseError != nil {
			return context.JSON(http.StatusInternalServerError, rpc.ErrorResponse(request, rpc.InvalidRequest, parseError))
		}

		providerIdRaw := context.Param("providerId")
		if providerIdRaw == "" {
			return context.JSON(http.StatusInternalServerError, rpc.ErrorResponse(request, rpc.RequestMissingProviderId, errors.New("missing provider id in the request path")))
		}

		providerId, parseError := uuid.Parse(providerIdRaw)
		if parseError != nil {
			return errorResponse(context, rpc.ErrorResponse(request, rpc.ParseError, parseError))
		}

		provider, exists := providers[providerId]
		if !exists {
			return context.JSON(http.StatusNotFound, rpc.ErrorResponse(request, rpc.ProviderNotFound, errors.New("provider was not found")))
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
