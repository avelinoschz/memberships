package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Server ...
type Server struct {
	router    *mux.Router
	startupAt time.Time
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// NewServer ...
func NewServer() *Server {
	s := &Server{
		router:    mux.NewRouter(),
		startupAt: time.Now(),
	}
	s.routes()
	return s
}
