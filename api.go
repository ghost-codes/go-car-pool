package main

type ApiServer struct {
	listenAddr string
	store      Storage
}

func NewServer(portNumber string, pg Postgres) *ApiServer {
	store := Storage{pg}

	return &ApiServer{
		portNumber,
		store,
	}
}
