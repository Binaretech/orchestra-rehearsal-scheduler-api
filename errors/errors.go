package errors

import "net/http"

const (
	INVALID_CREDENTIALS = "auth/invalid-credentials"
	FORBIDDEN           = "auth/forbidden"
	UNAUTHORIZED        = "auth/unauthorized"
	CONFLICT            = "error/conflict"
	NOT_FOUND           = "error/not-found"

	SECTION_NOT_FOUND = "section/not-found"

	INVALID_DATA_TYPE = "validation/invalid-data-type"

	CONCERT_PAST_DATE           = "concert/past-date"
	CONCERT_PAST_REHEARSAL_DATE = "concert/past-rehearsal-date"

	INTERNAL_ERROR = "internal/error"
)

type AppError interface {
	Code() int
	Error() string
	Message() string
}

type BadRequestError struct {
	message string
}

func NewBadRequestError(message string) BadRequestError {
	return BadRequestError{message: message}
}

func (e BadRequestError) Code() int {
	return http.StatusBadRequest
}

func (e BadRequestError) Error() string {
	return e.message
}

func (e BadRequestError) Message() string {
	return e.message
}
