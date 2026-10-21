package apperr

import "net/http"

type Error struct {
	Status  int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func New(status int, msg string) *Error {
	return &Error{
		Status:  status,
		Message: msg,
	}
}

func Invalid(msg string) *Error      { return New(http.StatusBadRequest, msg) }
func Unauthorized(msg string) *Error { return New(http.StatusUnauthorized, msg) }
func Forbidden(msg string) *Error    { return New(http.StatusForbidden, msg) }
func NotFound(msg string) *Error     { return New(http.StatusNotFound, msg) }
func Conflict(msg string) *Error     { return New(http.StatusConflict, msg) }
