package database

import (
	"ca-myproperty/config"
	model "ca-myproperty/models"
)

func CreateDeveloper(developer model.Developer) model.Developer {
	config.DB.Create(&developer).Joins("User")
	return developer
}

func UpdateDeveloper(id string, developer model.Developer) model.Developer {
	config.DB.Where("id = ?", id).Updates(&developer).Joins("User")
	return developer
}
func GetDeveloperByID(id string) model.Developer {
	var developer model.Developer
	config.DB.Where("id = ?", id).Find(&developer).Joins("User")
	return developer
}
