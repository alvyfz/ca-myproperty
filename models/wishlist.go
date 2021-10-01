package model

import "gorm.io/gorm"

type Wishlist struct {
	gorm.Model
	ID         uint      `gorm:"primarykey"`
	UserId     uint      `json:"user_id"`
	User       *User     `json:"user"`
	PropertyID uint      `json:"property_id"`
	Property   *Property `json:"property"`
}
