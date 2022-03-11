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

// StatsService defines the interface to implement by a
// service facade inside the application.
type StatsService interface {
	// ListAvailable retrieves all the StatsDescriptor (Tag and
	// FinishedAt) concerning the reports of a given user.
	ListAvailable(id string) ([]StatsDescriptor, error)
	// GetByID retrieves computed stats,
	// when provided with a StatsDescriptor id.
	GetByID(id string) (Stats, error)
}

type UserService interface {
	// Create creates and stores a User in the data layer
	// and returns its ID.
	Create(name, email string) (string, error)

	// Retrieve retrieves a User by ID from the data layer.
	Retrieve(id string) (User, error)

	// Exists returns true if a user with name ane email
	// already exists in the data layer.
	Exists(name, email string) bool
}
