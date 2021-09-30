package database

import (
	"ca-myproperty/config"
	model "ca-myproperty/models"
)

func CreateDeveloper(developer model.Developer) model.Developer {
	config.DB.Create(&developer)
	return developer
}

func UpdateDeveloper(id string, developer model.Developer) model.Developer {
	config.DB.Where("id = ?", id).Updates(&developer)
	return developer
}
