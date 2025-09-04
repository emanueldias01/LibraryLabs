package main

import (
	"log"

	"github.com/emanueldias01/LibraryLabs/db"
	"github.com/emanueldias01/LibraryLabs/router"
)

func main() {
	conn, err := db.OpenConnection()

	if err != nil{
		log.Fatal("Erro to open connection to db")
	}


	db.RunMigrations(conn)

	router.InitializeRoutes()
}