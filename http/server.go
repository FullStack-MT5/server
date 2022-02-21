package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/benchttp/server"
	"github.com/benchttp/server/http/httplog"
)

type Server struct {
	*http.Server
	router     *mux.Router
	Repository server.Repository
}

// NewServer returns a Server with specified configuration parameters.
func NewServer(addr string, repo server.Repository) *Server {
	return &Server{
		Server:     &http.Server{Addr: addr},
		Repository: repo,
	}
}

func (s *Server) Start() error {
	s.init()

	log.Printf("Server listening at http://localhost%s\n", s.Addr)
	return s.ListenAndServe()
}

func (s *Server) init() {
	s.router = mux.NewRouter().StrictSlash(true)

	s.router.Use(httplog.Request)

	s.registerRoutes()

	s.Handler = s.router
}

const alphanum20 = "[a-zA-Z0-9]{2,20}"

func (s *Server) registerRoutes() {
	s.router.HandleFunc("/", handleRoot)

	s.router.HandleFunc("/benchmark", s.handleRetrieveBenchmark).Methods("POST")

	s.router.HandleFunc("/benchmark", s.handleCreateBenchmark).Methods("GET").
		Queries("id", fmt.Sprintf("{id:%s}", alphanum20))
}

func handleRoot(rw http.ResponseWriter, _ *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	rw.WriteHeader(200)
	rw.Write([]byte("⚡")) //nolint
}
