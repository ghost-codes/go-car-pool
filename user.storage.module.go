package main

import (
	"fmt"
)

type UserType int

const (
	User UserType = iota
	Driver
)

type UserModel struct {
	ID             int      `json:"id"`
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	UserName       string   `json:"userName"`
	Email          string   `json:"email"`
	Phone          *string  `json:"phone"`
	Country        string   `json:"country"`
	HashedPassword string   `json:"-"`
	Type           UserType `json:"-"`
}

type UserStorageModule interface {
	CreateUser(data *SignUpData) (*UserModel, error)
	UpdateUser() error
	DeleteUser() error
	GetUserByIdentifier(Identifier string) (*UserModel, error)
	GetUserByID(ID int) error
	GetAllUser() error
}

func (pg *Postgres) createUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS passanger(
		id serial primary key,
	   	firstName varchar(50),
	   	lastName varchar(50),
	   	userName varchar(100),
	   	email varchar(255),
	   	phone varchar(15),
	   	country varchar(50),
	   	hashedPassword varchar(255),
	   	type serial
) `

	_, err := pg.db.Exec(query)

	return err
}

func (pg Postgres) CreateUser(data *SignUpData) (*UserModel, error) {
	query := `INSERT INTO passanger (
	   	firstName ,
	   	lastName ,
	   	userName ,
	   	email ,
	   	country ,
	   	hashedPassword ,
	   	type
	) VALUES($1,$2,$3,$4,$5,$6,$7)`

	_, err := pg.db.Query(query, data.FirstName, data.LastName, data.UserName, data.Email, "Ghana", data.Password, User)
	if err != nil {
		return nil, err
	}
	return pg.GetUserByIdentifier(data.Email)

}
func (pg Postgres) UpdateUser() error {
	return nil
}
func (pg Postgres) DeleteUser() error {
	return nil
}
func (pg Postgres) GetUser() error {
	return nil
}
func (pg Postgres) GetAllUser() error {
	return nil
}
func (pg Postgres) GetUserByID(ID int) error {
	return nil
}

func (pg Postgres) GetUserByIdentifier(Identifier string) (*UserModel, error) {
	query := `SELECT * FROM passanger WHERE email=$1 OR userName=$1`

	user := new(UserModel)

	rows, err := pg.db.Query(query, Identifier)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.UserName, &user.Email, &user.Phone, &user.Country, &user.HashedPassword, &user.Type)
		if err != nil {
			return nil, err
		}
		fmt.Print(user.ID)
		return user, nil
	}
	return nil, nil
}
