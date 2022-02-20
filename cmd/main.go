package main

import (
	"flag"
	"log"

	"github.com/benchttp/server/http"
)

const defaultPort = "9998"

func main() {
	port := flag.String("port", defaultPort, "Address for the server to listen on.")
	flag.Parse()
	addr := ":" + *port

	srv := http.NewServer(addr)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
