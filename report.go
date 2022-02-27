package server

import (
	"encoding/gob"
	"io"
	"net/http"

	"github.com/benchttp/server/benchttp"
)

func (s *Server) createReport(rw http.ResponseWriter, r *http.Request) {
	rep := benchttp.Report{}

	err := gob.NewDecoder(r.Body).Decode(&r)
	if err != nil && err != io.EOF {
		respondHTTPError(rw, errBadRequest)
		return
	}

	_, err = s.ReportService.Create(r.Context(), rep)
	if err != nil {
		respondHTTPError(rw, errInternal)
		return
	}

	respondJSON(rw, 201, nil)
}

func (s *Server) retrieveReport(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	report, err := s.ReportService.Retrieve(r.Context(), id)
	if err != nil {
		respondHTTPError(rw, errNotFound) // TODO differentiate not found and decoding
		return
	}

	respondJSON(rw, 200, report)
}
