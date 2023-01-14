package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddr string
	store      Storage
}

func NewServer(portNumber string, pg *Postgres) *ApiServer {
	store := Storage{pg}

	return &ApiServer{
		portNumber,
		store,
	}
}

// Type Def for api functions
type apiFunc func(w http.ResponseWriter, r *http.Request) ErrorI

// Api Error struct
type ApiError struct {
	code int    `jsont:"-"`
	Err  string `json:"error"`
}

type ErrorI interface {
	Error() string
	StatusCode() int
}

func (a ApiError) Error() string {
	return a.Err
}
func (a ApiError) StatusCode() int {
	return a.code
}

// json parser for
func WriteJson(w http.ResponseWriter, statusCode int, data any) ErrorI {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {

		return ApiError{
			Err:  err.Error(),
			code: http.StatusInternalServerError,
		}
	}
	return nil
}

// handler wrapper
func MakeHttpHanlder(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, err.StatusCode(), err)

		}
	}
}

// StartServer
func (api *ApiServer) StartServer() {
	router := mux.NewRouter()
	api.InitAuthHandlers(router)
	// router.HandleFunc("/login", MakeHttpHanlder(api.userLogin))

	log.Println("Json Api server running on port", api.listenAddr)

	err := http.ListenAndServe(api.listenAddr, router)
	if err != nil {
		log.Fatal(err)
	}

}
