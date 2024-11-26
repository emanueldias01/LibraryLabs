package database

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var(
	DB *gorm.DB
	err error
)

func ConnectDB(){
	strConn := "host=localhost user=root password=root dbname=librarylabs port=9920 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strConn))

	if err != nil{
		log.Fatal("fail to connect to the database")
	}
}