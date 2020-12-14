package rpc

import "errors"

func Success(request Request, result interface{}) Response {
	return Response{
		Id:     request.Id,
		Result: result,
	}
}

func ErrorResponse(request Request, code ErrorCode, err error) Response {
	return Response{
		Id: request.Id,
		Error: Error{
			Code:    code,
			Message: err.Error(),
		},
	}
}

func MethodNotFoundError(request Request) Response {
	return ErrorResponse(request, MethodNotFound, errors.New("methodNotFound"))
}
