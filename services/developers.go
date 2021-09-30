package service

import (
	database "ca-myproperty/lib/databases"
	model "ca-myproperty/models"
)

func CreateDeveloper(developer model.Developer) model.Developer {
	return database.CreateDeveloper(developer)
}

func UpdateDeveloper(id string, developer model.Developer) model.Developer {
	return database.UpdateDeveloper(id, developer)
}
