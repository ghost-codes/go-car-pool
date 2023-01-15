package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Yeah Buddy!")
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
