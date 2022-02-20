package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/benchttp/server/http/httplog"
)

type Server struct {
	*http.Server
	router *mux.Router
}

// NewServer returns a Server with specified configuration parameters.
func NewServer(addr string) *Server {
	return &Server{
		Server: &http.Server{Addr: addr},
	}
}

func (s *Server) Start() error {
	s.init()

	log.Printf("Server listening at http://localhost%s\n", s.Addr)
	return s.ListenAndServe()
}

func (s *Server) init() {
	s.router = mux.NewRouter().StrictSlash(true)
	s.registerRoutes()
	s.router.Use(httplog.Request)
	s.Handler = s.router
}

func (s *Server) registerRoutes() {
	s.router.HandleFunc("/", handleRoot)
	s.router.HandleFunc("/report", s.handleGetReport).Methods("GET")
	s.router.HandleFunc("/report", s.handlePostReport).Methods("POST")
}

func handleRoot(rw http.ResponseWriter, _ *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	rw.WriteHeader(200)
	rw.Write([]byte("🖕")) //nolint
}
