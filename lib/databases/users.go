package database

import (
	"ca-myproperty/config"
	"ca-myproperty/middlewares"

	model "ca-myproperty/models"
)

func GetAllUsers() []model.User {
	var users []model.User
	config.DB.Find(&users)
	return users
}

func GetUserByID(id string) model.User {
	var user model.User
	config.DB.Where("id = ?", id).Preload("Wishlist").Find(&user)
	return user
}

func CreateUser(user model.User) model.User {
	config.DB.Create(&user)
	return user
}

func UpdateUser(id string, user model.User) model.User {
	config.DB.Where("id = ?", id).Updates(&user)
	return user
}

func GetDetailUser(userId int) (interface{}, error) {
	var user model.User

	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}
func LoginUser(user *model.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	// if err := config.DB.Save(user).Error; err != nil {
	// 	return nil, err
	// }
	return user, nil
}

// func IsValidBasic(u string) bool {
// 	return true
// }

// func IsValid(u string, p string) bool {
// 	var user model.User
// 	if err := config.DB.Where("email = ?", u).Find(&user).Error; err != nil {
// 		return false
// 	}

// 	return p == user.Password
// }
