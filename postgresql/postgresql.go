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

////////////////////////////////////////////////// benchmarkService //////////////////////////////////////////////////

type benchmarkService struct {
	db *sql.DB
}

func NewBenchmarkService(idleConn, maxConn int) (benchttp.BenchmarkService, error) {
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

	return &benchmarkService{db}, nil
}

func (b *benchmarkService) Close() {
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

func (b *benchmarkService) ListMetadataByUserID(userID int) ([]*benchttp.Metadata, error) {
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

func (b *benchmarkService) FindBenchmarkDetailByID(metadataID int) (*benchttp.Benchmark, error) {
	benchmark := &benchttp.Benchmark{}

	stmt := "SELECT m.tag, m.finished_at, c.code_1xx, c.code_2xx, c.code_3xx, c.code_4xx, c.code_5xx, t.min, t.max, t.mean, t.median, t.variance, t.deciles " +
		"FROM public.metadata AS m " +
		"INNER JOIN public.codestats AS c ON c.metadata_id = m.id " +
		"INNER JOIN public.timestats AS t ON t.metadata_id = m.id " +
		"WHERE m.id = $1 " +
		"ORDER BY m.finished_at DESC"

	row := b.db.QueryRow(stmt, metadataID)
	err := row.Scan(&benchmark.Metadata.Tag, &benchmark.Metadata.FinishedAt, &benchmark.Codestats.Code1xx, &benchmark.Codestats.Code2xx, &benchmark.Codestats.Code3xx, &benchmark.Codestats.Code4xx, &benchmark.Codestats.Code5xx, &benchmark.Timestats.Min, &benchmark.Timestats.Max, &benchmark.Timestats.Mean, &benchmark.Timestats.Median, &benchmark.Timestats.Variance, (*pq.Float64Array)(&benchmark.Timestats.Deciles))
	if err != nil {
		return benchmark, ErrScanningRows
	}

	return benchmark, nil
}
