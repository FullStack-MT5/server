package server

import (
	"net/http"
	"strconv"
)

func (s *Server) ListMetadataByUserID(w http.ResponseWriter, r *http.Request) {
	// TO DO: get userID from authentication to use it here instead of "1"
	stats, err := s.ComputedStatsService.ListMetadataByUserID(1)
	if err != nil {
		writeError(w, ErrNotFound.Wrap(err)) // TODO differentiate not found and decoding
		return
	}

	writeJSON(w, stats, 200)
}

func (s *Server) FindComputedStatsByID(w http.ResponseWriter, r *http.Request) {
	idInt, err := getStringParamAndConvertToInt(r, idParam)
	if err != nil {
		writeError(w, ErrBadRequest.Wrap(err))
		return
	}

	stats, err := s.ComputedStatsService.FindComputedStatsByID(idInt)
	if err != nil {
		writeError(w, ErrNotFound.Wrap(err)) // TODO differentiate not found and decoding
		return
	}

	writeJSON(w, stats, 200)
}

func getStringParamAndConvertToInt(r *http.Request, idParam param) (int, error) {
	idString, err := pathParam(r, idParam)
	if err != nil {
		return 0, err
	}
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		return 0, err
	}
	return idInt, nil
}
