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
func IsValid(u string, p string) bool {
	return database.IsValid(u, p)
}

// func LoginUser(user *model.User) (interface{}, error) {
// 	return database.LoginUser(user)
// }

// func GetDetailUser(id int) (interface{}, error) {
// 	return database.GetDetailUser(id)
// }

// func isValid(u  , p string) model.User {
// 	return database.isValid(u, p)
// }
