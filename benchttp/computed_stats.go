package benchttp

import "time"

// Metadata contains a computed stats group metadata
type Metadata struct {
	Tag        string
	FinishedAt time.Time
}

// Codestats represents the code stats related to a computed stats group
type Codestats struct {
	Code1xx int
	Code2xx int
	Code3xx int
	Code4xx int
	Code5xx int
}

// Stats represents the time stats related to a computed stats group
type Timestats struct {
	Min      float64
	Max      float64
	Mean     float64
	Median   float64
	Variance float64
	Deciles  []float64
}

// ComputedStats contains Metadata, Codestats and Stats of a given computed stats group
type ComputedStats struct {
	Metadata  Metadata
	Codestats Codestats
	Timestats Timestats
}
