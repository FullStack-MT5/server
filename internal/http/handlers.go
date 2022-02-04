package http

import (
	"encoding/json"
	"net/http"

	"github.com/benchttp/server/internal"
)

func (s *Server) handlePostReport(rw http.ResponseWriter, r *http.Request) {
	report := internal.Report{}
	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		respondHTTPError(rw, errBadRequest)
		return
	}

	id, err := s.Repository.StoreReport(report)
	if err != nil {
		respondHTTPError(rw, errInternal)
		return
	}
	respondJSON(rw, 201, struct {
		Id string `json:"id"`
	}{Id: id})
}

func (s *Server) handleGetReport(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	report, err := s.Repository.RetrieveReport(id)
	if err != nil {
		respondHTTPError(rw, errNotFound) // TODO differenciate not found and decoding
		return
	}
	respondJSON(rw, 200, report)
}
