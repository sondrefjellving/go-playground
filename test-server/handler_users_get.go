package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, req *http.Request) {
	users, err := cfg.db.GetUsers()
	if err != nil {
		jsonErrorResponse(w, 500, "Couldn't find users")
		return
	}

	jsonResponse(w, http.StatusOK, users)
}