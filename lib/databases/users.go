package database

import (
	"ca-myproperty/config"

	model "ca-myproperty/models"
)

func GetAllUsers() []model.User {
	var users []model.User
	config.DB.Find(&users).Joins("Wishlist")
	return users
}

func GetUserByID(id string) model.User {
	var user model.User
	config.DB.Where("id = ?", id).Joins("Wishlist").Find(&user)
	return user
}

func CreateUser(user model.User) model.User {
	config.DB.Create(&user).Joins("Wishlist")
	return user
}

func UpdateUser(id string, user model.User) model.User {
	config.DB.Where("id = ?", id).Updates(&user).Joins("Wishlist")
	return user
}

func IsValid(u string, p string) bool {
	var user model.User
	if err := config.DB.Where("email = ?", u).Find(&user).Error; err != nil {
		return false
	}

	return p == user.Password
}
