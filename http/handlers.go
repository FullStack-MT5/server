package http

import (
	"encoding/gob"
	"io"
	"net/http"

	"github.com/benchttp/server"
)

func (s *Server) handleCreate(rw http.ResponseWriter, r *http.Request) {
	b := server.Benchmark{}

	err := gob.NewDecoder(r.Body).Decode(&b)
	if err != nil && err != io.EOF {
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

func (s *Server) handleRetrieve(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	report, err := s.Repository.Retrieve(r.Context(), id)
	if err != nil {
		respondHTTPError(rw, errNotFound) // TODO differentiate not found and decoding
		return
	}
	respondJSON(rw, 200, report)
}
