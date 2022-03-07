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

	ReportService    benchttp.ReportService
	BenchmarkService benchttp.BenchmarkService
}

// New returns a Server with specified configuration parameters.
func New(addr string, rs benchttp.ReportService, bs benchttp.BenchmarkService) *Server {
	return &Server{
		Server:           &http.Server{Addr: addr},
		ReportService:    rs,
		BenchmarkService: bs,
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
