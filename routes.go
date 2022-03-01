package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type param string

const (
	// alphanum20 is a regular expression matching an sequence of
	// 20 alphanumeric characters.
	alphanum20 = "[a-zA-Z0-9]{2,20}"

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

	v1.HandleFunc("/reports", s.createReport).Methods("POST")

	v1.HandleFunc("/reports/"+idPathVar, s.retrieveReport).Methods("GET")
}

func handleRoot(rw http.ResponseWriter, _ *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	rw.WriteHeader(200)
	rw.Write([]byte("⚡")) //nolint
}
