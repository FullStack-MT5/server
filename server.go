package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/benchttp/server/benchttp"
	"github.com/benchttp/server/httplog"
)

type Server struct {
	*http.Server
	router *mux.Router

	ReportService benchttp.ReportService
	StatsService  benchttp.StatsService
}

// New returns a Server with specified configuration parameters.
func New(port string, rs benchttp.ReportService, s benchttp.StatsService) *Server {
	return &Server{
		Server:        &http.Server{Addr: ":" + port},
		ReportService: rs,
		StatsService:  s,
	}
}

func (s *Server) Start() error {
	s.init()

	log.Printf("Server listening at %s\n", s.localAddr())

	return s.ListenAndServe()
}

func (s *Server) init() {
	s.router = mux.NewRouter().StrictSlash(true)

	s.router.Use(httplog.Request)

	s.registerRoutes()

	s.Handler = s.router
}

func (s *Server) localAddr() string {
	return fmt.Sprintf("http://localhost%s", s.Addr)
}
