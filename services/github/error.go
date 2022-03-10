package github

import "errors"

var (
	errRequest    = errors.New("error sending request")
	errStatusCode = errors.New("unexpected status (expected 200)")
	errParse      = errors.New("error parsing response")
)
