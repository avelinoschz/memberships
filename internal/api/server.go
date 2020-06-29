package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
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

// respond wraps the logic of marshaling the payload and the response
func (s *Server) respond(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", jsonapi.MediaType)

	// only transform to jsonapi format
	jsonapiPayload, err := jsonapi.Marshal(data)
	if err != nil {
		log.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		errs := []*jsonapi.ErrorObject{
			&jsonapiParseErr,
		}

		s.respondWithErrors(w, errs, http.StatusInternalServerError)
		return
	}

	// does the actual marshaling of the payload
	payload, err := json.Marshal(jsonapiPayload)
	if err != nil {
		log.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		errs := []*jsonapi.ErrorObject{
			&jsonParseErr,
		}

		s.respondWithErrors(w, errs, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(payload)
}

func (s *Server) respondWithErrors(w http.ResponseWriter, errs []*jsonapi.ErrorObject, status int) {
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(status)

	if err := jsonapi.MarshalErrors(w, errs); err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
