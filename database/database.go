package database

import (
	"log"
	//"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var(
	DB *gorm.DB
	err error
)

func ConnectDB(){
	//wait 10 seconds for connect to db
	//time.Sleep(10 * time.Second)
	strConn := "host=localhost user=root password=root dbname=LibraryLabs port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strConn))

	if err != nil{
		log.Fatal("fail to connect to the database")
	}
}