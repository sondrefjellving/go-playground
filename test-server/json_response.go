package main

import (
	"encoding/json"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	payloadAsJson, err := json.Marshal(payload)
	if err != nil {
		jsonErrorResponse(w, 500, "Error decoding payload to json")
		return
	}

	w.WriteHeader(status)
	w.Write(payloadAsJson)
}

func jsonErrorResponse(w http.ResponseWriter, status int, errorMessage string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	errorData := errorResponse{
		Error: errorMessage,
	}
	errorRes, err := json.Marshal(errorData)
	if err != nil {
		w.WriteHeader(500)	
		w.Write([]byte("Error converting error response to json"))
		return
	}
	
	w.WriteHeader(status)
	w.Write(errorRes)
}