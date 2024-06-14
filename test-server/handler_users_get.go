package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, req *http.Request) {
	users, err := cfg.db.GetUsers()
	w.Header().Add("Text-Content", "application/json")
	if err != nil {
		jsonErrorResponse(w, 500, "Couldn't find users")
		return
	}

	usersRes, err := json.Marshal(users)
	if err != nil {
		jsonErrorResponse(w, 500, "Error converting response to json")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(usersRes)
}