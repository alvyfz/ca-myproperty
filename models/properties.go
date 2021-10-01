package model

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	ID             uint          `gorm:"primaryKey"`
	Name           string        `json:"name"`
	Price          int           `json:"price"`
	Description    string        `json:"description"`
	Bedroom        int           `json:"bedroom"`
	Bathroom       int           `json:"bathroom"`
	SurfaceArea    int           `json:"surface_area"`
	BuildingArea   int           `json:"building_area"`
	Interior       string        `json:"interior"`
	Location       string        `json:"location"`
	PropertyTypeID uint          `json:"property_type_id"`
	PropertyType   *PropertyType `json:"property_type,omitempty"`
	UserID         uint          `json:"user_id"`
	User           *User         `json:"user,omitempty"`
	DeveloperID    uint          `json:"developer_id"`
	Developer      *Developer    `json:"developer,omitempty"`
}
