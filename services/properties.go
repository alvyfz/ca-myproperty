package service

import (
	database "ca-myproperty/lib/databases"
	model "ca-myproperty/models"
)

func GetAllProperties() []model.Property {
	return database.GetAllProperties()
}

func CreateProperty(Property model.Property) model.Property {
	return database.CreateProperty(Property)
}

func UpdateProperty(id string, Property model.Property) model.Property {
	return database.UpdateProperty(id, Property)
}

func GetPropertyByID(id string) model.Property {
	return database.GetPropertyByID(id)
}

func DeletePropertyByID(id string) model.Property {
	return database.DeletePropertyByID(id)
}
