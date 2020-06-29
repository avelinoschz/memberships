package api

import "github.com/google/jsonapi"

var (
	jsonParseErr = jsonapi.ErrorObject{
		Title:  "JSON parse",
		Detail: "There was an error parsing the JSON response",
		Status: "500",
	}

	jsonapiParseErr = jsonapi.ErrorObject{
		Title:  "JSON API parse",
		Detail: "There was an error parsing to JSON API response",
		Status: "500",
	}

	jsonAPIRepErr = jsonapi.ErrorObject{
		Title:  "JSONAPI representation",
		Detail: "Data is not a jsonapi representation",
		Status: "400",
	}

	invalidUser = jsonapi.ErrorObject{
		Title:  "Invalid User",
		Detail: "Member name or/and password were wrong",
		Status: "403",
	}
)
