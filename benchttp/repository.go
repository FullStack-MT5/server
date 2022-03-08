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

// ComputedStatsService defines the interface to implement by a
// service facade inside the application.
type ComputedStatsService interface {
	// ListMetadataByUserID retrieves all the metadata (Tag and
	// FInishedAt) related to a User with its user id.
	ListMetadataByUserID(id string) ([]Metadata, error)
	// FindComputedStatsByID retrieves computed stats,
	// when provided with a metadata id.
	FindComputedStatsByID(id string) (ComputedStats, error)
}
