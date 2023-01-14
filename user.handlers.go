package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitUserHandlers(router *mux.Router) {
	router.HandleFunc("/user", MakeHttpHanlder(CreateUserHandler))
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) ErrorI {
	return nil
}
