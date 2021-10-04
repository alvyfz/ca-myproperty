package service

import (
	database "ca-myproperty/lib/databases"
	model "ca-myproperty/models"
)

func GetAllUsers() []model.User {
	return database.GetAllUsers()
}

func CreateUser(user model.User) model.User {
	return database.CreateUser(user)
}

func UpdateUser(id string, user model.User) model.User {
	return database.UpdateUser(id, user)
}

func GetUserByID(id string) model.User {
	return database.GetUserByID(id)
}
func IsValid(u string, p string) (model.User, error) {
	return database.IsValid(u, p)
}

