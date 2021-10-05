package config

import (
	model "ca-myproperty/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var DBLog *mongo.Client

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Name     string
}

func InitDB() {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "root",
		"DB_Port":     "3306",
		"DB_Host":     "localhost",
		"DB_Name":     "myproperty_new",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

// if wanna using mongoDB incomment this
func InitLog() {
	var err error
	DBLog, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/myproperty"))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err = DBLog.Connect(ctx)
	if err != nil {
		panic(err)
	}

	DBLog.ListDatabaseNames(ctx, bson.M{})
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{}, &model.Chat{}, &model.Developer{}, &model.PropertyType{}, &model.Wishlist{}, &model.Property{}, &model.Transaction{})
}
