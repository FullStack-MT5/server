package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres" // blank import
)

type ComputedStatsService struct {
	db *sql.DB
}

func NewComputedStatsService(config Config) (ComputedStatsService, error) {
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.User,
		config.Password,
		config.DBName)

	db, err := sql.Open("cloudsqlpostgres", dbURI)
	if err != nil {
		return ComputedStatsService{}, ErrDatabaseConnection
	}

	err = db.Ping()
	if err != nil {
		return ComputedStatsService{}, ErrDatabasePing
	}

	db.SetMaxIdleConns(config.IdleConn)
	db.SetMaxOpenConns(config.MaxConn)

	return ComputedStatsService{db}, nil
}

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	IdleConn int
	MaxConn  int
}
