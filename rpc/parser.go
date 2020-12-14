package rpc

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"strings"
)

func ParseRequest(context echo.Context) (Request, error) {
	var request Request
	if jsonError := context.Bind(request); jsonError == nil {
		return request, nil
	} else {
		return request, jsonError
	}
}

func (request Request) ParseParams(target interface{}) error {
	return json.NewDecoder(strings.NewReader(request.Params)).Decode(target)
}
