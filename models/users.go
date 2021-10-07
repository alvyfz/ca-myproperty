package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	City        string `json:"city"`
}
