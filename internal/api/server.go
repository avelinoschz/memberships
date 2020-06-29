package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Server ...
type Server struct {
	db        *gorm.DB
	router    *mux.Router
	startupAt time.Time
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// NewServer ...
func NewServer(db *gorm.DB) (*Server, error) {

	if db == nil {
		return nil, errors.New("database is missing")
	}

	s := &Server{
		db:        db,
		router:    mux.NewRouter(),
		startupAt: time.Now(),
	}
	s.routes()
	return s, nil
}
