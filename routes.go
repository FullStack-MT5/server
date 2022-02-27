package server

import (
	"fmt"
	"net/http"
)

// alphanum20 is a regular expression matching an sequence of
// 20 alphanumeric characters.
const alphanum20 = "[a-zA-Z0-9]{2,20}"

func (s *Server) registerRoutes() {
	s.router.HandleFunc("/", handleRoot)

	s.router.HandleFunc("/report", s.createReport).Methods("POST")

	s.router.HandleFunc("/report", s.retrieveReport).Methods("GET").
		Queries("id", fmt.Sprintf("{id:%s}", alphanum20))
}

func handleRoot(rw http.ResponseWriter, _ *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	rw.WriteHeader(200)
	rw.Write([]byte("⚡")) //nolint
}
