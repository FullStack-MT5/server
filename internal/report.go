package internal

import (
	"net/http"
	"net/url"
	"time"
)

type Report struct {
	Target  Target   `json:"target"`
	Records []Record `json:"records"`
	Length  int      `json:"length"`
	Success int      `json:"success"`
	Fail    int      `json:"fail"`
}

type Record struct {
	Time  time.Duration `json:"time"`
	Code  int           `json:"code"`
	Bytes int           `json:"bytes"`
	Error error         `json:"error"`
}

type Target struct {
	Method string      `json:"method"`
	URL    *url.URL    `json:"url"`
	Body   []byte      `json:"body,omitempty"`   // Body is not used for now.
	Header http.Header `json:"header,omitempty"` // Header is not used for now.
}
