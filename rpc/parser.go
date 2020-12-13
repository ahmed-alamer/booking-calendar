package rpc

import "github.com/labstack/echo/v4"

func ParseRequest(context echo.Context) (Request, error) {
	var request Request
	if jsonError := context.Bind(request); jsonError == nil {
		return request, nil
	} else {
		return request, jsonError
	}
}
