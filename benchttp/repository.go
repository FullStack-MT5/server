package benchttp

import "context"

// ReportService defines the interface to implement by a
// service facade inside the application.
type ReportService interface {
	// Create creates and stores a Report in the data layer
	// and returns its ID.
	Create(ctx context.Context, data Report) (string, error)

	// Retrieve retrieves a Report by ID from the data layer.
	Retrieve(ctx context.Context, id string) (Report, error)
}

type BenchmarkService interface {
	Close()
	ListMetadataByUserID(id int) ([]*Metadata, error)
	FindBenchmarkDetailByID(id int) (*Benchmark, error)
}
