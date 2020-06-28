package api

import (
	"encoding/json"
	"math"
	"net/http"
	"time"

	"github.com/google/jsonapi"
	log "github.com/sirupsen/logrus"
)

// HandleAlive ...
func (s *Server) HandleAlive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("HandleAlive request")

		w.Header().Set("Content-Type", jsonapi.MediaType)

		alive := struct {
			CurrentTime time.Time `json:"currentTime"`
			Uptime      float64   `json:"uptime"`
		}{
			CurrentTime: time.Now().UTC(),
			Uptime:      math.Round(time.Since(s.startupAt).Seconds()*100) / 100,
		}

		payload, err := json.Marshal(alive)
		if err != nil {
			log.Error(err.Error())

			w.WriteHeader(http.StatusInternalServerError)

			jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
				Title:  "JSON Parse error",
				Detail: "There was an errors parsing the JSON response",
				Status: "500",
			}})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(payload)
	}
}
