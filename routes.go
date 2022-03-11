package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type param string

const (
	// alphanum20 is a regular expression matching an sequence of
	// 1 to 20 alphanumeric characters.
	alphanum20 = "[a-zA-Z0-9]{1,20}"

	// idParam is the parameter name for the path variable to identify
	// one resource of a collection of resources, typically in RESTful
	// APIs: "api/resources/{idParam}".
	idParam param = "id"
)

// pathParam returns the value of the path parameter p from the request
// context. Returns an error if the parameter is not found in the context.
func pathParam(r *http.Request, p param) (string, error) {
	v, ok := mux.Vars(r)[string(p)]
	if !ok {
		return "", fmt.Errorf("missing path parameter \"%s\"", p)
	}

	return v, nil
}

func (s *Server) registerRoutes() {
	idPathVar := fmt.Sprintf("{%s:%s}", idParam, alphanum20)

	s.router.HandleFunc("/", handleRoot)

	v1 := s.router.PathPrefix("/v1").Subrouter()

	// Auth
	v1.HandleFunc("/signin", s.handleSignin).Methods("POST")
	v1.HandleFunc("/token", s.mustAuth(s.handleCreateAccessToken)).Methods("GET")

	// Reports
	v1.HandleFunc("/reports", s.mustAuth(s.createReport)).Methods("POST")
	v1.HandleFunc("/reports/"+idPathVar, s.mustAuth(s.retrieveReport)).Methods("GET")

	// Stats
	v1.HandleFunc("/stats", s.mustAuth(s.retrieveAllStats)).Methods("GET")
	v1.HandleFunc("/stats/"+idPathVar, s.mustAuth(s.retrieveStatsByID)).Methods("GET")
}

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("⚡")) //nolint
}
