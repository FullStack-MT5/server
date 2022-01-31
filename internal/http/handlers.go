package http

import (
	"encoding/json"
	"net/http"

	"github.com/benchttp/server/internal"
)

func (s *Server) handlePostReport(rw http.ResponseWriter, r *http.Request) {
	report := &internal.Report{}
	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		respondHTTPError(rw, errNotFound)
		return
	}

	s.Repository.Report = report

	respondJSON(rw, 201, nil)
}

func (s *Server) handleGetReport(rw http.ResponseWriter, r *http.Request) {
	if s.Repository.Report == nil {
		respondHTTPError(rw, errNotFound)
		return
	}

	respondJSON(rw, 200, s.Repository.Report)
}
