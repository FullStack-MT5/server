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
			Time  time.Duration `firestore:"time"`
			Code  int           `firestore:"code"`
			Bytes int           `firestore:"bytes"`
			Error string        `firestore:"error,omitempty"`

			// Event is a stage of an outgoing HTTP request associated with a timestamp.
			Events []struct {
				Name string        `firestore:"name"`
				Time time.Duration `firestore:"time"`
			} `firestore:"events"`
		} `firestore:"records"`

		Length   int           `firestore:"length"`
		Success  int           `firestore:"success"`
		Fail     int           `firestore:"fail"`
		Duration time.Duration `firestore:"duration"`
	} `firestore:"benchmark"`

	Metadata struct {
		// Config represents the global configuration of the runner.
		Config struct {
			// Request contains the confing options relative to a single request.
			Request struct {
				Method string      `firestore:"method"`
				URL    *url.URL    `firestore:"url"`
				Header http.Header `firestore:"header"`

				Body struct {
					Type    string `firestore:"type"`
					Content []byte `firestore:"content"`
				} `firestore:"body"`
			} `firestore:"request"`

			// Runner contains options relative to the runner.
			Runner struct {
				Requests       int           `firestore:"requests"`
				Concurrency    int           `firestore:"concurrency"`
				Interval       time.Duration `firestore:"interval"`
				RequestTimeout time.Duration `firestore:"requestTimeout"`
				GlobalTimeout  time.Duration `firestore:"globalTimeout"`
			} `firestore:"runner"`

			// Output contains options relative to the output.
			Output struct {
				Out      []string
				Silent   bool
				Template string
			} `firestore:"-"`
		} `firestore:"config"`

		FinishedAt time.Time `firestore:"finishedAt"`
	} `firestore:"metadata"`
}
