package database

import (
	"ca-myproperty/config"
	model "ca-myproperty/models"
)

func GetAllTransactions() []model.Transaction {
	var transaction []model.Transaction
	config.DB.Find(&transaction).Joins("User", "Property")
	return transaction
}

func GetTransactionByID(id string) model.Transaction {
	var transaction model.Transaction
	config.DB.Where("id = ?", id).Find(&transaction).Joins("Property")
	return transaction
}

func CreateTransaction(transaction model.Transaction) model.Transaction {
	config.DB.Create(&transaction)
	config.DB.Where("transactions.id = ?", transaction.ID).Joins("Property").Find(&transaction)
	return transaction
}

func UpdateTransaction(id string, transaction model.Transaction) model.Transaction {
	config.DB.Where("id = ?", id).Updates(&transaction).Joins("Property")
	return transaction
}
func DeleteTransactionByID(id string) model.Transaction {
	var transaction model.Transaction
	config.DB.Where("id = ?", id).Delete(&transaction)
	return transaction
}
