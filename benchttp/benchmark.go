package benchttp

import "time"

// Metadata contains a benchmark metadata
type Metadata struct {
	Tag        string
	FinishedAt time.Time
}

// Codestats represents the code stats related to a benchmark
type Codestats struct {
	Code1xx int
	Code2xx int
	Code3xx int
	Code4xx int
	Code5xx int
}

// Stats represents the time stats related to a benchmark
type Timestats struct {
	Min      float64
	Max      float64
	Mean     float64
	Median   float64
	Variance float64
	Deciles  []float64
}

// Benchmark contains Metadata, Codestats and Stats of a given Benchmark
type Benchmark struct {
	Metadata  Metadata
	Codestats Codestats
	Timestats Timestats
}
