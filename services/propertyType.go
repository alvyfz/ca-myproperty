package service

import (
	database "ca-myproperty/lib/databases"
	model "ca-myproperty/models"
)

func GetAllPropertyTypes() []model.PropertyType {
	return database.GetAllPropertyTypes()
}

func CreatePropertyType(PropertyType model.PropertyType) model.PropertyType {
	return database.CreatePropertyType(PropertyType)
}

func UpdatePropertyType(id string, PropertyType model.PropertyType) model.PropertyType {
	return database.UpdatePropertyType(id, PropertyType)
}

func GetPropertyTypeByID(id string) model.PropertyType {
	return database.GetPropertyTypeByID(id)
}
