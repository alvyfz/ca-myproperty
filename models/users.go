package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone_number"`
	Gender      string    `json:"gender"`
	City        string    `json:"city"`
	WishlistID  uint      `json:"wishlist_id"`
	Wishlist    *Wishlist `json:"wishlist,omitempty"`
}
