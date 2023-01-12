package main

type UserStorageModule interface {
	CreateUser() error
	UpdateUser() error
	DeleteUser() error
	GetUser() error
	GetAllUser() error
}

func (pg Postgres) CreateUser() error {
	return nil
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
