package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
	log "github.com/sirupsen/logrus"
	"gopkg.in/validator.v2"
)

// decode wraps the logic and errors from every endpoint repeatead decode action
func decode(r *http.Request, data interface{}) []*jsonapi.ErrorObject {
	// TODO fix to accept jsonapi request format
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		log.Error(err.Error())

		errObjects := []*jsonapi.ErrorObject{
			&jsonParseErr,
		}

		return errObjects
	}
	return nil
}

func validate(data interface{}) []*jsonapi.ErrorObject {
	if err := validator.Validate(data); err != nil {
		log.Error(err.Error())

		var errObjects []*jsonapi.ErrorObject

		vErrs := err.(validator.ErrorMap)
		for f, e := range vErrs {
			errObjects = append(errObjects, &jsonapi.ErrorObject{
				Title:  fmt.Sprintf("Invalid %s", f),
				Detail: fmt.Sprintf("%v", e),
			})
		}

		return errObjects
	}

	return nil
}

// respond wraps the logic of marshaling the payload and the response
func respond(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", jsonapi.MediaType)

	// only transform to jsonapi format
	jsonapiPayload, err := jsonapi.Marshal(data)
	if err != nil {
		log.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		errs := []*jsonapi.ErrorObject{
			&jsonapiParseErr,
		}

		respondWithErrors(w, errs, http.StatusInternalServerError)
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

		respondWithErrors(w, errs, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(payload)
}

// respond wraps the logic of marshaling errors in jsonapi format
func respondWithErrors(w http.ResponseWriter, errs []*jsonapi.ErrorObject, status int) {
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(status)

	if err := jsonapi.MarshalErrors(w, errs); err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
