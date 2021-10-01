package database

import (
	"ca-myproperty/config"
	model "ca-myproperty/models"
)

func CreateWishlist(wishlist model.Wishlist) model.Wishlist {
	config.DB.Create(&wishlist).Joins("User", "Property")
	return wishlist
}

func DeleteWishlistByID(id string) model.Wishlist {
	var wishlist model.Wishlist
	config.DB.Where("id = ?", id).Delete(&wishlist)
	return wishlist
}
