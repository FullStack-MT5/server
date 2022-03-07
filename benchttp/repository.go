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

// BenchmarkService defines the interface to implement by a
// service facade inside the application.
type BenchmarkService interface {
	// Close closes the connection to the database storing the
	// benchmarks data.
	Close()
	// ListMetadataByUserID retrieves all the metadata (Tag and
	// FInishedAt) related to a User with the User id.
	ListMetadataByUserID(id int) ([]*Metadata, error)
	// FindBenchmarkDetailByID retrieves a benchmark and all
	// its related details (metadata, codestats, timestats),
	// provided a metadata id.
	FindBenchmarkDetailByID(id int) (*Benchmark, error)
}
