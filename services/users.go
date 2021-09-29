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
