package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/benchttp/server"
	"github.com/benchttp/server/firestore"
	"github.com/benchttp/server/postgresql"
)

const defaultPort = "9998"

func main() {
	port := flag.String("port", defaultPort, "Address for the server to listen on.")
	flag.Parse()
	addr := ":" + *port

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	projectID := os.Getenv("GOOGLE_PROJECT_ID")
	if projectID == "" {
		log.Fatalf("GOOGLE_PROJECT_ID variable is not defined")
	}

	collectionID := os.Getenv("FIRESTORE_COLLECTION_ID")
	if projectID == "" {
		log.Fatalf("FIRESTORE_COLLECTION_ID variable is not defined")
	}

	rs, err := firestore.NewReportService(context.Background(), projectID, collectionID)
	if err != nil {
		log.Fatal(err)
	}

	bs, err := postgresql.NewBenchmarkService(10, 10)
	if err != nil {
		log.Fatal(err)
	}

	srv := server.New(addr, rs, bs)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
