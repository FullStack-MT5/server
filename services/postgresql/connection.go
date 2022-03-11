package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres" // blank import
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	IdleConn int
	MaxConn  int
}

// Connection holds the connection to a PostgreSQL database.
// It must be passed to a service constructor.
type Connection struct {
	db *sql.DB
}

// Connect opens a connection with a PostgreSQL database and
// returns a Connection to utilize it.
func Connect(config Config) (Connection, error) {
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.User,
		config.Password,
		config.DBName)

	db, err := sql.Open("cloudsqlpostgres", dbURI)
	if err != nil {
		return Connection{}, ErrDatabaseConnection
	}

	err = db.Ping()
	if err != nil {
		return Connection{}, ErrDatabasePing
	}

	db.SetMaxIdleConns(config.IdleConn)
	db.SetMaxOpenConns(config.MaxConn)

	return Connection{db}, nil
}
