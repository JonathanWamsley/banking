package errs

import "net/http"

// AppError returns errors relevant for a http response, a status and error message
type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

// AsMessage returns only the string
func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

// NewNotFoundError returns status not found(404) error + msg
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

// NewUnexpectedError returns internal sever(500) error + msg
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

// NewValidationError returns status unprocessable(422) error + msg
func NewValidationError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}
