package main

import (
	"flag"
	"log"

	"github.com/benchttp/server/internal/http"
)

func main() {
	port := flag.String("port", "9000", "Address for the server to listen on.")
	flag.Parse()
	addr := ":" + *port

	srv := http.NewServer(addr)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
