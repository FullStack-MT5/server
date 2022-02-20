package http

import (
	"net/http"
)

func (s *Server) handlePostReport(rw http.ResponseWriter, r *http.Request) {
	respondHTTPError(rw, errInternal)
}

func (s *Server) handleGetReport(rw http.ResponseWriter, r *http.Request) {
	respondHTTPError(rw, errInternal)
}
