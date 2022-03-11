package server

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrBadRequest = httpError{
		Code:    http.StatusBadRequest,
		Message: http.StatusText(http.StatusBadRequest),
	}
	ErrNotFound = httpError{
		Code:    http.StatusNotFound,
		Message: http.StatusText(http.StatusNotFound),
	}
	ErrInternal = httpError{
		Code:    http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
	}
	ErrUnauthorized = httpError{
		Code:    http.StatusUnauthorized,
		Message: http.StatusText(http.StatusUnauthorized),
	}
)

type httpError struct {
	Code    int    `json:"-"`
	Message string `json:"error"`

	inner error
}

func (e *httpError) Error() string {
	if e.inner != nil {
		return fmt.Sprintf("%s: %s", e.Message, e.inner)
	}
	return e.Message
}

func (e *httpError) Wrap(err error) error {
	return &httpError{
		Code:    e.Code,
		Message: e.Message,
		inner:   err,
	}
}

func (e *httpError) Unwrap() error {
	return e.inner
}

// httpErrorOf tries to read err as httpError.
// Defaults to ErrInternal if err is not an httpError.
func httpErrorOf(err error) *httpError {
	var e *httpError
	if errors.As(err, &e) && e != nil {
		return e
	}
	return &ErrInternal
}
