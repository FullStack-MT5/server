package server

import (
	"net/http"
	"net/url"
	"time"
)

// Benchmark represents the result of a Benchttp benchmark run.
type Benchmark struct {
	// Report represents the detailed collection of requests done
	// during a Benchttp benchmark run.
	Report struct {
		// Record represents the summary of a HTTP response.
		Records []struct {
			Time  time.Duration
			Code  int
			Bytes int
			Error string

			// Event is a stage of an outgoing HTTP request associated with a timestamp.
			Events []struct {
				Name string
				Time time.Duration
			}
		}

		Length   int
		Success  int
		Fail     int
		Duration time.Duration
	}

	Metadata struct {
		// Config represents the global configuration of the runner.
		Config struct {
			// Request contains the confing options relative to a single request.
			Request struct {
				Method string
				URL    *url.URL
				Header http.Header

				Body struct {
					Type    string
					Content []byte
				}
			}

			// Runner contains options relative to the runner.
			Runner struct {
				Requests       int
				Concurrency    int
				Interval       time.Duration
				RequestTimeout time.Duration
				GlobalTimeout  time.Duration
			}

			// Output contains options relative to the output.
			Output struct {
				Out      []string
				Silent   bool
				Template string
			}
		}

		FinishedAt time.Time
	}
}
