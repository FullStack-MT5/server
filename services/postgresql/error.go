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
	// ErrGettingIDInsertion is returned when server fails to get the ID
	// of an object that has just been inserted into the database.
	ErrGettingIDInsertion = errors.New("error retrieving ID of the row inserted")
)
