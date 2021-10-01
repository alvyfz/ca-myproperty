package database

import (
	"ca-myproperty/config"
	model "ca-myproperty/models"
)

func GetAllPropertyTypes() []model.PropertyType {
	var PropertyTypes []model.PropertyType
	config.DB.Find(&PropertyTypes)
	return PropertyTypes
}

func GetPropertyTypeByID(id string) model.PropertyType {
	var PropertyType model.PropertyType
	config.DB.Where("id = ?", id).Find(&PropertyType)
	return PropertyType
}

func CreatePropertyType(PropertyType model.PropertyType) model.PropertyType {
	config.DB.Create(&PropertyType)
	return PropertyType
}

func UpdatePropertyType(id string, PropertyType model.PropertyType) model.PropertyType {
	config.DB.Where("id = ?", id).Updates(&PropertyType)
	return PropertyType
}
