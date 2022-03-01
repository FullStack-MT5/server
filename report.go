package server

import (
	"encoding/gob"
	"net/http"

	"github.com/benchttp/server/benchttp"
)

func (s *Server) createReport(w http.ResponseWriter, r *http.Request) {
	rep := benchttp.Report{}

	err := gob.NewDecoder(r.Body).Decode(&rep)
	if err != nil {
		writeError(w, &ErrBadRequest)
		return
	}

	_, err = s.ReportService.Create(r.Context(), rep)
	if err != nil {
		writeError(w, &ErrInternal)
		return
	}

	w.WriteHeader(201)
}

func (s *Server) retrieveReport(w http.ResponseWriter, r *http.Request) {
	id, err := pathParam(r, idParam)
	if err != nil {
		writeError(w, ErrBadRequest.Wrap(err))
		return
	}

	report, err := s.ReportService.Retrieve(r.Context(), id)
	if err != nil {
		writeError(w, ErrNotFound.Wrap(err)) // TODO differentiate not found and decoding
		return
	}

	writeJSON(w, report, 200)
}
