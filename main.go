package main

import (
	"fmt"
	"librarylabs/database"
	"librarylabs/routes"
)

func main() {
	fmt.Println("Initialize API LibraryLabs")
	database.ConnectDB()

	defer routes.HandleRequest()
}