package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres" // blank import
	"github.com/joho/godotenv"
	"github.com/lib/pq"

	"github.com/benchttp/server/benchttp"
)

////////////////////////////////////////////////// computedStatsService //////////////////////////////////////////////////

type computedStatsService struct {
	db *sql.DB
}

func NewComputedStatsService(idleConn, maxConn int) (benchttp.ComputedStatsService, error) {
	connectionInfo, err := getConnectionInfoFromEnvVariables()
	if err != nil {
		return nil, err
	}

	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		connectionInfo.host,
		connectionInfo.user,
		connectionInfo.password,
		connectionInfo.dbName)

	db, err := sql.Open("cloudsqlpostgres", dbURI)
	if err != nil {
		return nil, ErrDatabaseConnection
	}

	err = db.Ping()
	if err != nil {
		return nil, ErrDatabasePing
	}

	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	return &computedStatsService{db}, nil
}

func (b *computedStatsService) Close() {
	b.db.Close()
}

////////////////////////////////////////////////// Connection //////////////////////////////////////////////////

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

	host := os.Getenv("HOST")
	if host == "" {
		appendError(errors.New("could not find 'HOST' environment variable"))
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		appendError(errors.New("could not find 'DB_USER' environment variable"))
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		appendError(errors.New("could not find 'DB_PASSWORD' environment variable"))
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		appendError(errors.New("could not find 'DB_NAME' environment variable"))
	}

	if len(envVariablesErrors) > 0 {
		return connectionInfo{}, &ErrNotFoundEnvVariables{envVariablesErrors}
	}

	return connectionInfo{host, user, password, dbName}, nil
}

////////////////////////////////////////////////// GET //////////////////////////////////////////////////

func (b *computedStatsService) ListMetadataByUserID(userID int) ([]*benchttp.Metadata, error) {
	metadataList := make([]*benchttp.Metadata, 0)

	stmt, err := b.db.Prepare("SELECT tag, finished_at FROM metadata WHERE user_id = $1 ORDER BY finished_at DESC")
	if err != nil {
		return []*benchttp.Metadata{}, ErrPreparingStmt
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return []*benchttp.Metadata{}, ErrExecutingPreparedStmt
	}
	defer rows.Close()

	for rows.Next() {
		metadata := new(benchttp.Metadata)
		err = rows.Scan(
			&metadata.Tag,
			&metadata.FinishedAt,
		)
		if err != nil {
			return nil, err
		}
		metadataList = append(metadataList, metadata)
	}

	return metadataList, nil
}

func (b *computedStatsService) FindComputedStatsByID(metadataID int) (*benchttp.ComputedStats, error) {
	computedStats := &benchttp.ComputedStats{}

	stmt := "SELECT m.tag, m.finished_at, c.code_1xx, c.code_2xx, c.code_3xx, c.code_4xx, c.code_5xx, t.min, t.max, t.mean, t.median, t.variance, t.deciles " +
		"FROM public.metadata AS m " +
		"INNER JOIN public.codestats AS c ON c.metadata_id = m.id " +
		"INNER JOIN public.timestats AS t ON t.metadata_id = m.id " +
		"WHERE m.id = $1 " +
		"ORDER BY m.finished_at DESC"

	row := b.db.QueryRow(stmt, metadataID)
	err := row.Scan(&computedStats.Metadata.Tag, &computedStats.Metadata.FinishedAt, &computedStats.Codestats.Code1xx, &computedStats.Codestats.Code2xx, &computedStats.Codestats.Code3xx, &computedStats.Codestats.Code4xx, &computedStats.Codestats.Code5xx, &computedStats.Timestats.Min, &computedStats.Timestats.Max, &computedStats.Timestats.Mean, &computedStats.Timestats.Median, &computedStats.Timestats.Variance, (*pq.Float64Array)(&computedStats.Timestats.Deciles))
	if err != nil {
		return computedStats, ErrScanningRows
	}

	return computedStats, nil
}
