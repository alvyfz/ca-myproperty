package model

import "gorm.io/gorm"

type PropertyType struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
}
