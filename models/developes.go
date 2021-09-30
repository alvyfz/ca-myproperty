package model

import "gorm.io/gorm"

type Developer struct {
	gorm.Model
	ID          uint   `gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	City        string `json:"city"`
	UserID      uint   `json:"user_id"`
	User        *User  `json:"user,omitempty"`
}