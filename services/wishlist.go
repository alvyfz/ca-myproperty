package service

import (
	database "ca-myproperty/lib/databases"
	model "ca-myproperty/models"
)

func CreateWishlist(wishlist model.Wishlist) model.Wishlist {
	return database.CreateWishlist(wishlist)
}

func DeleteWishlistByID(id string) model.Wishlist {
	return database.DeleteWishlistByID(id)
}
