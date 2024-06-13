package main

import "net/http"

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, req *http.Request) {
	users, err := cfg.db.GetUsers()
	if err != nil {

	}
}