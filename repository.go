package server

import "context"

// Repository defines the interface to implement by a
// data layer facade inside the application.
type Repository interface {
	// Create creates and stores a Benchmark in the data layer
	// and returns its ID.
	Create(ctx context.Context, data Benchmark) (string, error)

	// Retrieve retrieves a Benchmark by ID from the data layer.
	Retrieve(ctx context.Context, id string) (Benchmark, error)
}
