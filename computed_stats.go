package server

import (
	"net/http"
)

func (s *Server) ListMetadataByUserID(w http.ResponseWriter, r *http.Request) {
	// TO DO: get userID from authentication to use it here instead of "1"
	stats, err := s.ComputedStatsService.ListMetadataByUserID("1")
	if err != nil {
		writeError(w, ErrNotFound.Wrap(err)) // TODO differentiate not found and decoding
		return
	}

	writeJSON(w, stats, 200)
}

func (s *Server) FindComputedStatsByID(w http.ResponseWriter, r *http.Request) {
	id, err := pathParam(r, idParam)
	if err != nil {
		writeError(w, ErrBadRequest.Wrap(err))
		return
	}

	stats, err := s.ComputedStatsService.FindComputedStatsByID(id)
	if err != nil {
		writeError(w, ErrNotFound.Wrap(err)) // TODO differentiate not found and decoding
		return
	}

	writeJSON(w, stats, 200)
}
