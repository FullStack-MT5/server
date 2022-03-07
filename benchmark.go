package server

import (
	"net/http"
	"strconv"
)

func (s *Server) ListMetadataByUserID(w http.ResponseWriter, r *http.Request) {
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
	report, err := s.BenchmarkService.ListMetadataByUserID(idInt)
	if err != nil {
		writeError(w, ErrNotFound.Wrap(err)) // TODO differentiate not found and decoding
		return
	}

	writeJSON(w, report, 200)
}

func (s *Server) FindBenchmarkDetailByID(w http.ResponseWriter, r *http.Request) {
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
	report, err := s.BenchmarkService.FindBenchmarkDetailByID(idInt)
	if err != nil {
		writeError(w, ErrNotFound.Wrap(err)) // TODO differentiate not found and decoding
		return
	}

	writeJSON(w, report, 200)
}
