package database

import (
	"ca-myproperty/config"
	model "ca-myproperty/models"
)

func GetAllUsers() []model.User {
	var users []model.User
	config.DB.Find(&users)
	return users
}

func CreateUser(user model.User) model.User {
	config.DB.Create(&user)
	return user
}

func UpdateUser(id string, user model.User) model.User {
	config.DB.Where("id = ?", id).Updates(&user)
	return user
}