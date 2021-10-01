package database

import (
	"ca-myproperty/config"
	model "ca-myproperty/models"
)

func GetAllTransactions() []model.Transaction {
	var transaction []model.Transaction
	config.DB.Joins("User", "Property").Find(&transaction)
	return transaction
}

func GetTransactionByID(id string) model.Transaction {
	var transaction model.Transaction
	config.DB.Where("id = ?", id).Joins("User").Joins("Property").Find(&transaction)
	return transaction
}

func CreateTransaction(transaction model.Transaction) model.Transaction {
	config.DB.Create(&transaction).Joins("User").Joins("Property")
	return transaction
}

func UpdateTransaction(id string, transaction model.Transaction) model.Transaction {
	config.DB.Where("id = ?", id).Updates(&transaction).Joins("User").Joins("Property")
	return transaction
}
func DeleteTransactionByID(id string) model.Transaction {
	var transaction model.Transaction
	config.DB.Where("id = ?", id).Delete(&transaction)
	return transaction
}
