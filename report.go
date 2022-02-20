package server

import (
	"net/url"
	"time"
)

type Report struct {
	Config  Config   `json:"config"`
	Records []Record `json:"records"`
	Length  int      `json:"length"`
	Success int      `json:"success"`
	Fail    int      `json:"fail"`
}

type Record struct {
	Time  time.Duration `json:"time"`
	Code  int           `json:"code"`
	Bytes int           `json:"bytes"`
	Error error         `json:"error,omitempty"`
}

type Config struct {
	Request       Request
	RunnerOptions RunnerOptions
}

type Request struct {
	Method  string
	URL     *url.URL
	Timeout time.Duration
}

// RunnerOptions contains options relative to the runner.
type RunnerOptions struct {
	Requests      int
	Concurrency   int
	GlobalTimeout time.Duration
}
