package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `json:"user_id"`
	User       *User     `json:"user"`
	PropertyID uint      `json:"property_id,"`
	Property   *Property `json:"property,omitempty"`
}
