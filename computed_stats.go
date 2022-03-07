package server

import (
	"net/http"
	"strconv"
)

func (s *Server) ListMetadataByUserID(w http.ResponseWriter, r *http.Request) {
	// TO DO: get userID from authentication to use it here instead of "1"
	report, err := s.ComputedStatsService.ListMetadataByUserID(1)
	if err != nil {
		writeError(w, ErrNotFound.Wrap(err)) // TODO differentiate not found and decoding
		return
	}

	writeJSON(w, report, 200)
}

func (s *Server) FindComputedStatsByID(w http.ResponseWriter, r *http.Request) {
	idString, err := pathParam(r, idParam)
	if err != nil {
		writeError(w, ErrBadRequest.Wrap(err))
		return
	}

	idInt, err := strconv.Atoi(idString)
	if err != nil {
		writeError(w, ErrBadRequest.Wrap(err))
		return
	}
	report, err := s.ComputedStatsService.FindComputedStatsByID(idInt)
	if err != nil {
		writeError(w, ErrNotFound.Wrap(err)) // TODO differentiate not found and decoding
		return
	}

	writeJSON(w, report, 200)
}
