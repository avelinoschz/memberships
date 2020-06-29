package api

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/avelinoschz/yofio/internal/auth"
	"github.com/avelinoschz/yofio/internal/member"
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
func (s *Server) handleAuthLogin(w http.ResponseWriter, r *http.Request) {
	var member member.Member
	err := json.NewDecoder(r.Body).Decode(&member)
	if err != nil {
		log.Printf("Error reading user: %s\n", err)
	}

	if member.Name == "avelino" && member.Password == "password" {
		member.Password = "" // clean password to re-use model

		token := auth.GenerateJWT(string(member.ID))
		respToken := auth.ResponseToken{
			Token: token,
		}

		jsonResp, err := json.Marshal(respToken)
		if err != nil {
			log.Println("Error marshaling json response token")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
		return
	}

	w.WriteHeader(http.StatusForbidden)
	fmt.Fprintln(w, "Invalid user or password")
	return
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
