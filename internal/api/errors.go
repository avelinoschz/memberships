package api

import "github.com/google/jsonapi"

var (
	jsonParseErr = jsonapi.ErrorObject{
		Title:  "JSON Parse error",
		Detail: "There was an errors parsing the JSON response",
		Status: "500",
	}
)
