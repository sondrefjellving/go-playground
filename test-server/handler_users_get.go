package main

import (
	"net/http"
	"strconv"
)

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, req *http.Request) {
	users, err := cfg.db.GetUsers()
	if err != nil {
		jsonErrorResponse(w, 500, "Couldn't find users")
		return
	}

	jsonResponse(w, http.StatusOK, users)
}

func (cfg *apiConfig) handlerUsersGetById(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("userId"))
	if err != nil {
		jsonErrorResponse(w, http.StatusBadRequest, "Not valid id")
		return
	}

	user, err := cfg.db.GetUserById(id)
	if err != nil {
		jsonErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, user)
}