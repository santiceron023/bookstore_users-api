package errors

import (
	"errors"
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

func NewError(msg string) error{
	return errors.New(msg)
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Error:   "not_found",
		Code:    http.StatusNotFound,
	}
}
func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Error:   "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}
