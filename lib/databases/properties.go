package database

import (
	"ca-myproperty/config"
	model "ca-myproperty/models"
)

func GetAllProperties() []model.Property {
	var property []model.Property
	config.DB.Joins("PropertyType", "User", "Developer").Find(&property)
	return property
}

func GetPropertyByID(id string) model.Property {
	var property model.Property
	config.DB.Where("id = ?", id).Joins("PropertyType", "User", "Developer").Find(&property)
	return property
}

func CreateProperty(Property model.Property) model.Property {
	config.DB.Create(&Property).Joins("PropertyType", "User", "Developer")
	return Property
}

func UpdateProperty(id string, property model.Property) model.Property {
	config.DB.Where("id = ?", id).Updates(&property).Joins("PropertyType", "User", "Developer")
	return property
}
func DeletePropertyByID(id string) model.Property {
	var property model.Property
	config.DB.Where("id = ?", id).Delete(&property)
	return property
}
