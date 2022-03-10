package benchttp

import "time"

// StatsDescriptor contains a computed stats group description information
type StatsDescriptor struct {
	ID         string    `json:"id"`
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

// Timestats represents the time stats related to a computed stats group
type Timestats struct {
	Min     float64   `json:"min"`
	Max     float64   `json:"max"`
	Mean    float64   `json:"mean"`
	Median  float64   `json:"median"`
	StdDev  float64   `json:"standardDeviation"`
	Deciles []float64 `json:"deciles"`
}

// Stats contains StatsDescriptor, Codestats and Timestats of a given computed stats group
type Stats struct {
	Descriptor StatsDescriptor `json:"descriptor"`
	Code       Codestats       `json:"code"`
	Time       Timestats       `json:"time"`
}
