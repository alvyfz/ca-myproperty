package service

import (
	database "ca-myproperty/lib/databases"
	model "ca-myproperty/models"
)

func CreateChat(chat model.Chat) model.Chat {
	return database.CreateChat(chat)
}

func DeleteChatByID(id string) model.Chat {
	return database.DeleteChatByID(id)
}
