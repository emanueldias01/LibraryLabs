package service_test

import (
	"librarylabs/database"
	"librarylabs/models"
)

var ID int

func SetupDB(){
	database.ConnectDB()
}

func MockBook() models.Book{
	var book models.Book = models.Book{Name: "Name book", Author: "Emanuel Dias", Genre: "Fantasy", Description: "Any description Any description Any description Any description Any description Any description Any description"}
	database.DB.Save(&book)
	ID = book.Id

	return book
}

func MockDeleteBook(){
	var book models.Book
	database.DB.Delete(&book, ID)
}