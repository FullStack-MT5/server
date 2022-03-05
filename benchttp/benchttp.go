package benchttp

import (
	"net/http"
	"net/url"
	"time"
)

// Report represents the result of a Benchttp benchmark run.
type Report struct {
	// Benchmark represents the detailed collection of requests done
	// during a Benchttp benchmark run.
	Benchmark struct {
		// Record represents the summary of a HTTP response.
		Records []struct {
			Time  time.Duration `firestore:"time" json:"time"`
			Code  int           `firestore:"code" json:"code"`
			Bytes int           `firestore:"bytes" json:"bytes"`
			Error string        `firestore:"error,omitempty" json:"error,omitempty"`

			// Event is a stage of an outgoing HTTP request associated with a timestamp.
			Events []struct {
				Name string        `firestore:"name" json:"name"`
				Time time.Duration `firestore:"time" json:"time"`
			} `firestore:"events" json:"events"`
		} `firestore:"records" json:"records"`

		Length   int           `firestore:"length" json:"length"`
		Success  int           `firestore:"success" json:"success"`
		Fail     int           `firestore:"fail" json:"fail"`
		Duration time.Duration `firestore:"duration" json:"duration"`
	} `firestore:"benchmark" json:"benchmark"`

	Metadata struct {
		// Config represents the global configuration of the runner.
		Config struct {
			// Request contains the confing options relative to a single request.
			Request struct {
				Method string      `firestore:"method" json:"method"`
				URL    *url.URL    `firestore:"url" json:"url"`
				Header http.Header `firestore:"header" json:"header"`

				Body struct {
					Type    string `firestore:"type" json:"type"`
					Content []byte `firestore:"content" json:"content"`
				} `firestore:"body" json:"body"`
			} `firestore:"request" json:"request"`

			// Runner contains options relative to the runner.
			Runner struct {
				Requests       int           `firestore:"requests" json:"requests"`
				Concurrency    int           `firestore:"concurrency" json:"concurrency"`
				Interval       time.Duration `firestore:"interval" json:"interval"`
				RequestTimeout time.Duration `firestore:"requestTimeout" json:"requestTimeout"`
				GlobalTimeout  time.Duration `firestore:"globalTimeout" json:"globalTimeout"`
			} `firestore:"runner" json:"runner"`

			// Output contains options relative to the output.
			Output struct {
				Out      []string
				Silent   bool
				Template string
			} `firestore:"-" json:"-"`
		} `firestore:"config" json:"config"`

		FinishedAt time.Time `firestore:"finishedAt" json:"finishedAt"`
	} `firestore:"metadata" json:"metadata"`
}
