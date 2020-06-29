package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/go-validator/validator"
	"github.com/google/jsonapi"
	log "github.com/sirupsen/logrus"
)

// HandleAlive ...
func (s *Server) HandleAlive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := struct {
			CurrentTime time.Time `jsonapi:"attr,currentTime"`
			Uptime      float64   `jsonapi:"attr,uptime"`
		}{
			CurrentTime: time.Now().UTC(),
			Uptime:      math.Round(time.Since(s.startupAt).Seconds()*100) / 100,
		}

		s.respond(w, &payload, http.StatusOK)
	}
}

// HandleMemberCreate ...
func (s *Server) HandleMemberCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var member Member

		// decode helper

		// TODO fix to accept jsonapi request format
		if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
			log.Error(err.Error())

			errObjects := []*jsonapi.ErrorObject{
				&jsonParseErr,
			}

			s.respondWithErrors(w, errObjects, http.StatusBadRequest)
			return
		}

		if err := validator.Validate(member); err != nil {
			log.Error(err.Error())

			var errObjects []*jsonapi.ErrorObject

			vErrs := err.(validator.ErrorMap)
			for f, e := range vErrs {
				errObjects = append(errObjects, &jsonapi.ErrorObject{
					Title:  fmt.Sprintf("Invalid %s", f),
					Detail: fmt.Sprintf("%v", e),
				})
			}

			s.respondWithErrors(w, errObjects, http.StatusUnprocessableEntity)
			return
		}

		s.respond(w, &member, http.StatusOK)
	}
}
