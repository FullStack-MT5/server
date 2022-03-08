package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/benchttp/server"
	"github.com/benchttp/server/services/firestore"
	"github.com/benchttp/server/services/postgresql"
)

const defaultPort = "9998"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	port := flag.String("port", defaultPort, "Address for the server to listen on.")
	flag.Parse()
	addr := ":" + *port

	err := godotenv.Load(".env")
	// no error returned here because .env is not deployed
	if err != nil {
		fmt.Println("no .env file found")
	}

	projectID := os.Getenv("GOOGLE_PROJECT_ID")
	if projectID == "" {
		return errors.New("GOOGLE_PROJECT_ID variable is not defined")
	}

	collectionID := os.Getenv("FIRESTORE_COLLECTION_ID")
	if projectID == "" {
		return errors.New("FIRESTORE_COLLECTION_ID variable is not defined")
	}

	var psqlConfig postgresql.Config

	psqlConfig.Host = os.Getenv("PSQL_HOST")
	if psqlConfig.Host == "" {
		return errors.New("could not find 'PSQL_HOST' environment variable")
	}

	psqlConfig.User = os.Getenv("PSQL_USER")
	if psqlConfig.User == "" {
		return errors.New("could not find 'PSQL_USER' environment variable")
	}

	psqlConfig.Password = os.Getenv("PSQL_PASSWORD")
	if psqlConfig.Password == "" {
		return errors.New("could not find 'PSQL_PASSWORD' environment variable")
	}

	psqlConfig.DBName = os.Getenv("PSQL_NAME")
	if psqlConfig.DBName == "" {
		return errors.New("could not find 'PSQL_NAME' environment variable")
	}

	psqlConfig.IdleConn = 10
	psqlConfig.MaxConn = 25

	rs, err := firestore.NewReportService(context.Background(), projectID, collectionID)
	if err != nil {
		return err
	}

	s, err := postgresql.NewStatsService(psqlConfig)
	if err != nil {
		return err
	}

	srv := server.New(addr, rs, s)
	return srv.Start()
}
