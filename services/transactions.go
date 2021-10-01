package service

import (
	database "ca-myproperty/lib/databases"
	model "ca-myproperty/models"
)

func GetAllTransactions() []model.Transaction {
	return database.GetAllTransactions()
}

func CreateTransaction(transaction model.Transaction) model.Transaction {
	return database.CreateTransaction(transaction)
}

func UpdateTransaction(id string, transaction model.Transaction) model.Transaction {
	return database.UpdateTransaction(id, transaction)
}

func GetTransactionByID(id string) model.Transaction {
	return database.GetTransactionByID(id)
}

func DeleteTransactionByID(id string) model.Transaction {
	return database.DeleteTransactionByID(id)
}
