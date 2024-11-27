package models

import (
	"fmt"
	"librarylabs/database"
)

func GetAllBooks() []Book{
	var list []Book
	database.DB.Find(&list)

	return list
}

func GetBookById(id int) Book{
	var book Book
	database.DB.First(&book, id)

	return book
}

func CreateBook(bookCreate Book){
	bookCreate.SelectGenre(bookCreate.Genre)
	database.DB.Create(&bookCreate)
}

func UpdateBook(bodyBook Book, id int) Book{
	//Find the book i want to update
	var bookUpdate Book
	database.DB.First(&bookUpdate, id)

	//Change atributes book
	if bodyBook.Name != "" {
		bookUpdate.Name = bodyBook.Name
	}

	if bodyBook.Author != "" {
		bookUpdate.Author = bodyBook.Author
	}

	if bodyBook.Description != "" {
		bookUpdate.Description = bodyBook.Description
	}

	if bodyBook.Genre != "" {
		bookUpdate.SelectGenre(bodyBook.Genre)
	}

	//Save new changes
	database.DB.Save(&bookUpdate)

	return bookUpdate
}

func DeleteBook(id int){
	var bookDelete Book
	database.DB.Delete(&bookDelete, id)
}

func SetBookAvailable(id int) (Book, error){

	var(
		book Book
		err error
	)

	database.DB.Find(&book, id)

	if !book.Available{
		book.Available = true
		database.DB.Save(&book)
	}else{
		err = fmt.Errorf("This book is already Avaliable")
	}

	

	return book, err
}

func SetBookUnavailable(id int) (Book, error){
	var(
		book Book
		err error
	)
	database.DB.Find(&book, id)

	if book.Available{
		book.Available = false
		database.DB.Save(&book)
	}else{
		err = fmt.Errorf("This book is already Unavailable")
	}
	
	

	return book,err
}

func GetBooksByGenre(genre string)[]Book{
	var list[]Book

	database.DB.Raw("SELECT * FROM books WHERE genre = ?", genre).Scan(&list)

	return list
}

func GetBooksByName(name string)[]Book{
	var list[]Book

	database.DB.Raw("SELECT * FROM books WHERE name LIKE ?", "%"+name+"%").Scan(&list)

	return list
}