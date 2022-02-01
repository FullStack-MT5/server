package repository

import (
	"github.com/benchttp/server/internal"
	"github.com/benchttp/server/internal/store"
	"github.com/google/uuid"
)

// Repository exposes the available operations to access the data layer.
type Repository struct {
	store store.Store
}

// New returns a new repository.
// For now it uses store.NewDefault as the internal store.
func New() (*Repository, error) {
	return &Repository{
		store: store.NewDefault(),
	}, nil
}

// StoreReport stores a report in the data layer
// and returns the key to retrieve it later on.
func (r *Repository) StoreReport(data internal.Report) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	idStr := id.String()

	if err := r.store.Store(idStr, data); err != nil {
		return "", err
	}
	return idStr, nil
}

// RetrieveReport retrieves a report by id from the data layer.
func (r *Repository) RetrieveReport(id string) (internal.Report, error) {
	report := internal.Report{}

	if err := r.store.Load(id, &report); err != nil {
		return internal.Report{}, err
	}
	return report, nil
}
