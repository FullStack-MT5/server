package benchttp

import "time"

// Metadata contains a computed stats group metadata
type Metadata struct {
	Tag        string    `json:"tag"`
	FinishedAt time.Time `json:"finishedAt"`
}

// Codestats represents the code stats related to a computed stats group
type Codestats struct {
	Code1xx int `json:"code1xx"`
	Code2xx int `json:"code2xx"`
	Code3xx int `json:"code3xx"`
	Code4xx int `json:"code4xx"`
	Code5xx int `json:"code5xx"`
}

// Stats represents the time stats related to a computed stats group
type Timestats struct {
	Min      float64   `json:"min"`
	Max      float64   `json:"max"`
	Mean     float64   `json:"mean"`
	Median   float64   `json:"median"`
	Variance float64   `json:"average"`
	Deciles  []float64 `json:"deciles"`
}

// ComputedStats contains Metadata, Codestats and Stats of a given computed stats group
type ComputedStats struct {
	Metadata  `json:"metadata"`
	Codestats `json:"codestats"`
	Timestats `json:"timestats"`
}
