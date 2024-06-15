package main

import (
	"io"
	"net/http"
)

func (cfg *apiConfig) handlerUsersPost(w http.ResponseWriter, req *http.Request) {
	dat, err := io.ReadAll(req.Body)
	if err != nil {
		jsonErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}

	savedUser, err := cfg.db.CreateUser(dat)
	if err != nil {
		jsonErrorResponse(w, http.StatusInternalServerError, "Error saving user to db")
		return
	}

	jsonResponse(w, http.StatusCreated, savedUser)
}