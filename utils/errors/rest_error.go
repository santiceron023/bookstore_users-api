package errors

import (
	"net/http"
)

type RestError struct {
	Message string `json: "message"`
	Error   string `json: "error"`
	Code    int    `json: "code"`
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Error:   "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Error:   "not_found",
		Code:    http.StatusNotFound,
	}
}
