package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

type Storage struct {
	userStore UserStorageModule
}

func NewStorate() (*Postgres, error) {
	password := os.Getenv("password")
	connStr := fmt.Sprintf("user=postgres dbname=go_car_pool password=%s sslmode=disable", password)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Postgres{db}, nil
}

func (pg *Postgres) intStorage() error {
	return pg.createUserTable()
}
