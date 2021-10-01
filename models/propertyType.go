package model

import "gorm.io/gorm"

type PropertyType struct {
	gorm.Model
	ID   uint   `gorm:"primarykey"`
	Name string `json:"name"`
}
