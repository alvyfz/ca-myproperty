package database

import (
	"ca-myproperty/config"
	"errors"

	model "ca-myproperty/models"
)

// func GetAllUsers() []model.User {
// 	var users []model.User
// 	config.DB.Find(&users)
// 	return users
// }

func GetUserByID(id string) model.User {
	var user model.User
	config.DB.Where("id = ?", id).Find(&user)
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

func IsValid(u string, p string) (model.User, error) {
	var user model.User
	if err := config.DB.Where("email = ?", u).Find(&user).Error; err != nil {
		// fmt.Println(err)
		return model.User{}, err

	}
	if user.Password != p {
		return model.User{}, errors.New("email or password incorrect")
	}

	return user, nil
}
