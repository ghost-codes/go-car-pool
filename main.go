package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Yeah Buddy!")
	store, err := NewStorate()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.intStorage(); err != nil {
		log.Fatal(err)
	}
	server := NewServer(":8989", store)
	server.StartServer()
}
