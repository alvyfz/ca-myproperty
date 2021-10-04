package database

import (
	"ca-myproperty/config"
	model "ca-myproperty/models"
)

func CreateChat(chat model.Chat) model.Chat {
	config.DB.Create(&chat).Joins("Developer")
	return chat
}

func DeleteChatByID(id string) model.Chat {
	var chat model.Chat
	config.DB.Where("id = ?", id).Delete(&chat)
	return chat
}
