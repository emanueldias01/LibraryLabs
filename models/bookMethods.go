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

func GetBookById(id int) (Book, error){
	var(
		book Book
		err error
	)
	database.DB.First(&book, id)

	if book.Id == 0{
		err = fmt.Errorf("Book not found")
	}

	return book, err
}

func CreateBook(bookCreate Book){
	ValidateFields(&bookCreate)
	bookCreate.SelectGenre(bookCreate.Genre)
	database.DB.Create(&bookCreate)
}

func UpdateBook(bodyBook Book, id int) (Book, error){
	var (
		book Book
		err error
	)

	database.DB.First(&book, id)

	if book.Id == 0{
		err = fmt.Errorf("Book not found")
	}

	database.DB.Model(&book).UpdateColumns(bodyBook)


	database.DB.First(&book, id)
	return book, err
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

	if book.Id == 0{
		err = fmt.Errorf("Book not found")
	}

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

	if book.Id == 0{
		err = fmt.Errorf("Book not found")
	}

	if book.Available{
		book.Available = false
		database.DB.Save(&book)
	}
	if !book.Available && book.Id != 0{
		err = fmt.Errorf("This book is already Unavailable")
	}
	

	return book,err
}

func GetBooksByGenre(genre string)[]Book{
	var list[]Book

	database.DB.Where(Book{Genre: genre}).Find(&list)

	return list
}

func GetBooksByName(name string)[]Book{
	var list[]Book

	database.DB.Where(&Book{Name: name}).Find(&list)

	return list
}