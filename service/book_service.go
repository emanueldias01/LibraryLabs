package service

import (
	"fmt"
	"librarylabs/database"
	"librarylabs/models"
	"librarylabs/repository"
)

func GetAllBooks() []models.Book{
	var list []models.Book
	database.DB.Find(&list)

	return list
}

func GetBookById(id int) (models.Book, error){
	var(
		book models.Book
		err error
	)
	database.DB.First(&book, id)

	if book.Id == 0{
		err = fmt.Errorf("Book not found")
	}

	return book, err
}

func CreateBook(bookCreate models.Book) error{
	if err := models.ValidateFields(&bookCreate); err != nil{
		return err
	}
	bookCreate.SelectGenre(bookCreate.Genre)
	repository.Create(&bookCreate)

	return nil
}

func UpdateBook(bodyBook models.Book, id int) (models.Book, error){
	var (
		book models.Book
		err error
	)

	repository.Find(&book,id)

	if book.Id == 0{
		err = fmt.Errorf("Book not found")
	}

	repository.Update(book, bodyBook)


	repository.Find(&book,id)
	return book, err
}

func DeleteBook(id int){
	repository.Delete(id)
}

func SetBookAvailable(id int) (models.Book, error){

	var(
		book models.Book
		err error
	)

	repository.Find(&book,id)

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

func SetBookUnavailable(id int) (models.Book, error){
	var(
		book models.Book
		err error
	)
	repository.Find(&book, id)

	if book.Id == 0{
		err = fmt.Errorf("Book not found")
	}

	if book.Available{
		book.Available = false
		repository.Save(&book)
	}
	if !book.Available && book.Id != 0{
		err = fmt.Errorf("This book is already Unavailable")
	}
	

	return book,err
}

func GetBooksByGenre(genre string)[]models.Book{
	var list[]models.Book

	repository.FindByGenre(&list, genre)

	return list
}

func GetBooksByName(name string)[]models.Book{
	var list[]models.Book
	repository.FindByName(&list, name)
	return list
}