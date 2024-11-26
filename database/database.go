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
	strConn := "host=localhost user=root password=root dbname=LibraryLabs port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strConn))

	if err != nil{
		log.Fatal("fail to connect to the database")
	}
}