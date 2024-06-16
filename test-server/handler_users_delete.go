package main

import (
	"net/http"
	"strconv"
)

func (cfg *apiConfig) handlerUsersDeleteById(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("userId"))	
	if err != nil {
		jsonErrorResponse(w, http.StatusBadRequest, "Invalid id")
		return
	}

	err = cfg.db.DeleteUserById(id)
	if err != nil {
		jsonErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	jsonResponse(w, http.StatusNoContent, struct{}{})
}