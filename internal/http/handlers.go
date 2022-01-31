package http

import (
	"encoding/json"
	"net/http"

	"github.com/benchttp/server/internal"
)

func (s *Server) handlePostReport(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	report := &internal.Report{}
	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		rw.WriteHeader(400)
		return
	}

	s.Repository.Report = report

	rw.WriteHeader(201)
}

func (s *Server) handleGetReport(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	if s.Repository.Report == nil {
		rw.WriteHeader(404)
		return
	}

	resp, err := json.Marshal(s.Repository.Report)

	if err != nil {
		rw.WriteHeader(500)
		return
	}

	rw.Write(resp)
	rw.WriteHeader(200)
}
