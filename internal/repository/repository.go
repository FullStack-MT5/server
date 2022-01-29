package repository

import "errors"

// Repository exposes the available operations to access the data layer.
type Repository struct{}

// New returns a new repository given configuration parameters.
func New(cfg interface{}) (*Repository, error) {
	return nil, errors.New("not implemented")
}

// StoreReport stores a new report in the data layer.
func (r *Repository) StoreReport(data interface{}) error {
	return errors.New("not implemented")
}

// RetrieveReport retrieves a report given its id from the data layer.
func (r *Repository) RetrieveReport(id int) (interface{}, error) {
	return nil, errors.New("not implemented")
}
