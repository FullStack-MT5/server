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

// errorCode tries to read err as httpError to extract
// its code. Default to 500 if err is not httpError or nil.
func errorCode(err error) int {
	var e *httpError

	if err == nil {
		return 500
	} else if errors.As(err, &e) {
		return e.Code
	}
	return 500
}

// errorCode tries to read err as httpError to extract
// its message. Default to "Internal Server Error" if
// err is not httpError or nil.
func errorMessage(err error) string {
	var e *httpError

	if err == nil {
		return http.StatusText(500)
	} else if errors.As(err, &e) {
		return e.Message
	}
	return http.StatusText(500)
}
