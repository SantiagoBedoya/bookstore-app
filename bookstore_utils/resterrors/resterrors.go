package resterrors

import "net/http"

type RestError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

func NewRestError(message string, statusCode int, err string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: statusCode,
		Error:      err,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Error:      "not_found",
	}
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "bad_request",
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Error:      "internal_server_error",
	}
}

func NewUnauthorizedError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
		Error:      "unauthorized",
	}
}
