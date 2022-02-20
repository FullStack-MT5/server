package http

import (
	"encoding/gob"
	"fmt"
	"io"
	"net/http"

	"github.com/benchttp/server"
)

func (s *Server) handlePostReport(rw http.ResponseWriter, r *http.Request) {
	b := server.Benchmark{}

	err := gob.NewDecoder(r.Body).Decode(&b)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		respondHTTPError(rw, errBadRequest)
		return
	}

	_, err = s.Repository.Create(r.Context(), b)
	if err != nil {
		respondHTTPError(rw, errInternal)
		return
	}
	respondJSON(rw, 201, nil)
}

func (s *Server) handleGetReport(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	report, err := s.Repository.Retrieve(r.Context(), id)
	if err != nil {
		respondHTTPError(rw, errNotFound) // TODO differentiate not found and decoding
		return
	}
	respondJSON(rw, 200, report)
}
