package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres" // blank import
	"github.com/joho/godotenv"
)

type ComputedStatsService struct {
	db *sql.DB
}

func NewComputedStatsService(idleConn, maxConn int) (ComputedStatsService, error) {
	connectionInfo, err := getConnectionInfoFromEnvVariables()
	if err != nil {
		return ComputedStatsService{}, err
	}

	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		connectionInfo.host,
		connectionInfo.user,
		connectionInfo.password,
		connectionInfo.dbName)

	db, err := sql.Open("cloudsqlpostgres", dbURI)
	if err != nil {
		return ComputedStatsService{}, ErrDatabaseConnection
	}

	err = db.Ping()
	if err != nil {
		return ComputedStatsService{}, ErrDatabasePing
	}

	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	return ComputedStatsService{db}, nil
}

type connectionInfo struct {
	host     string
	user     string
	password string
	dbName   string
}

func getConnectionInfoFromEnvVariables() (connectionInfo, error) {
	envVariablesErrors := []error{}

	appendError := func(err error) {
		envVariablesErrors = append(envVariablesErrors, err)
	}

	err := godotenv.Load(".env")
	if err != nil {
		appendError(errors.New("could not load '.env' file"))
	}

	host := os.Getenv("PSQL_HOST")
	if host == "" {
		appendError(errors.New("could not find 'HOST' environment variable"))
	}

	user := os.Getenv("PSQL_USER")
	if user == "" {
		appendError(errors.New("could not find 'DB_USER' environment variable"))
	}

	password := os.Getenv("PSQL_PASSWORD")
	if password == "" {
		appendError(errors.New("could not find 'DB_PASSWORD' environment variable"))
	}

	dbName := os.Getenv("PSQL_NAME")
	if dbName == "" {
		appendError(errors.New("could not find 'DB_NAME' environment variable"))
	}

	if len(envVariablesErrors) > 0 {
		return connectionInfo{}, &ErrNotFoundEnvVariables{envVariablesErrors}
	}

	return connectionInfo{host, user, password, dbName}, nil
}
