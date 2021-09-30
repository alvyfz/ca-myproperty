package model

import "gorm.io/gorm"

type PropertyType struct {
	gorm.Model
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
