package api

import (
	"math"
	"net/http"
	"time"

	"github.com/avelinoschz/memberships/internal/auth"
	"github.com/avelinoschz/memberships/internal/member"
	"github.com/google/jsonapi"
	log "github.com/sirupsen/logrus"
)

// HandleAlive is a health check endpoint.
// Returns current time of the request and the server's uptime.
func (s *Server) handleAlive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := struct {
			CurrentTime time.Time `jsonapi:"attr,currentTime"`
			Uptime      float64   `jsonapi:"attr,uptime"`
		}{
			CurrentTime: time.Now().UTC(),
			Uptime:      math.Round(time.Since(s.startupAt).Seconds()*100) / 100,
		}

		respond(w, &payload, http.StatusOK)
	}
}

// Login returns a JWT with the same user info received and standard claims
func (s *Server) handleAuthLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var member member.Member
		if errs := decode(r, &member); errs != nil {
			respondWithErrors(w, errs, http.StatusBadRequest)
			return
		}

		if member.Name == "avelino" && member.Password == "password" {
			token := auth.GenerateJWT(string(member.ID))
			respToken := auth.ResponseToken{
				Token: token,
			}

			respond(w, &respToken, http.StatusOK)
			return
		}

		log.Error("Invalid user or password")
		errs := []*jsonapi.ErrorObject{
			&invalidUser,
		}

		respondWithErrors(w, errs, http.StatusForbidden)
		return
	}
}

// HandleMembersCreate registers a new member in the API.
func (s *Server) handleMembersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var member member.Member

		if errs := decode(r, &member); errs != nil {
			respondWithErrors(w, errs, http.StatusBadRequest)
			return
		}

		if errs := validate(&member); errs != nil {
			respondWithErrors(w, errs, http.StatusBadRequest)
			return
		}

		s.db.NewRecord(member)

		respond(w, &member, http.StatusOK)
	}
}

// HandlePaymentsGet ...
func (s *Server) handlePaymentsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
	}
}

// HandlePaymentsCreate ...
func (s *Server) handlePaymentsCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
	}
}
