package repository

import (
	"github.com/benchttp/server/internal"
)

// Repository exposes the available operations to access the data layer.
type Repository struct {
	Report *internal.Report
}

// New returns a new repository given configuration parameters.
func New() (*Repository, error) {
	return &Repository{}, nil
}

// StoreReport stores a new report in the data layer.
func (r *Repository) StoreReport(data internal.Report) error {
	r.Report = &data
	return nil
}

// RetrieveReport retrieves a report given its id from the data layer.
func (r *Repository) RetrieveReport(id int) (*internal.Report, error) {
	return r.Report, nil
}
