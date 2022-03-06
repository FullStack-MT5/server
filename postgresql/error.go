package postgresql

import (
	"errors"
)

var (
	// ErrDatabaseConnection is returned when server fails to connect to
	// the database.
	ErrDatabaseConnection = errors.New("database connection error")
	// ErrDatabasePing is returned when server fails to ping the
	// database.
	ErrDatabasePing = errors.New("database ping error")
	// ErrPreparingStmt is returned when server fails to prepare
	// a prepared statement.
	ErrPreparingStmt = errors.New("error executing prepared statement")
	// ErrExecutingPreparedStmt is returned when server fails to execute
	// a query with a prepared statement.
	ErrExecutingPreparedStmt = errors.New("error executing prepared statement")
	// ErrScanningRows is returned when server fails to scan the rows
	// returned by a query.
	ErrScanningRows = errors.New("error trying to scan result rows")
)

// ErrNotFoundEnvVariables regroups all environment variables not found
type ErrNotFoundEnvVariables struct {
	notFoundEnvVariables []error
}

// ErrNotFoundEnvVariables.Error() lists all environment variables not found
func (e *ErrNotFoundEnvVariables) Error() string {
	message := "Environment variables not found:\n"
	for _, err := range e.notFoundEnvVariables {
		message += err.Error() + "\n"
	}
	return message
}
