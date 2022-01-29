package http

import (
	"log"
	"net/http"

	"github.com/benchttp/server/pkg/httplog"
	"github.com/gorilla/mux"
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
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("👋 🌍"))
}
