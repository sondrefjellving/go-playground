package main

import (
	"encoding/json"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	payloadAsJson, err := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Write(payloadAsJson)
}

func jsonErrorResponse(w http.ResponseWriter, status int, errorMessage string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	jsonResponse(w, status, errorResponse{
		Error: errorMessage,
	})
}