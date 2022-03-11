package server

import (
	"net/http"
)

func (s *Server) selfUser(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r.Context())
	if user == nil {
		writeError(w, &ErrInternal)
		return
	}

	writeJSON(w, user, 200)
}
