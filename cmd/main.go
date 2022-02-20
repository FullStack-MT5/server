package main

import (
	"flag"
	"log"

	"github.com/benchttp/server/http"
	"github.com/benchttp/server/internal/repository"
)

const defaultPort = "9998"

func main() {
	port := flag.String("port", defaultPort, "Address for the server to listen on.")
	flag.Parse()
	addr := ":" + *port

	repo, err := repository.New()
	if err != nil {
		log.Fatal(err)
	}

	srv := http.NewServer(addr, repo)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
