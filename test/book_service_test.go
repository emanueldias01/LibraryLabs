package service_test

import (
	"librarylabs/database"
	"librarylabs/models"
	"librarylabs/service"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestListAllBookSucess(t *testing.T){
	SetupDB()
	book := MockBook()
	defer MockDeleteBook()

	list := service.GetAllBooks()

	assert.Contains(t, list, book, "Return mock book")
}

func TestFindBookByIdSucess(t *testing.T){
	SetupDB()
	book := MockBook()
	defer MockDeleteBook()

	result, _ := service.GetBookById(ID)

	assert.Equal(t, book, result, "Must return mock book")
}

func TestSetBookAvaliableSucess(t *testing.T){
	SetupDB()
	MockBook()
	defer MockDeleteBook()

	service.SetBookAvailable(ID)
	result, _ := service.GetBookById(ID)

	assert.Equal(t, true, result.Available, "Set book available")

}

func TestSetBookUnavaliableSucess(t *testing.T){
	SetupDB()
	MockBook()
	defer MockDeleteBook()

	service.SetBookAvailable(ID)
	service.SetBookUnavailable(ID)
	result, _ := service.GetBookById(ID)

	assert.Equal(t, false, result.Available, "Set book unavailable")

}

func TestGetBooksByGenreContainsBookSucess(t *testing.T){
	SetupDB()
	book := MockBook()
	defer MockDeleteBook()

	list := service.GetBooksByGenre("Fantasy")

	assert.Contains(t, list, book, "The mock book is present in list because your genre is fantasy")
}

func TestGetBooksByGenreNotContainsBookSucess(t *testing.T){
	SetupDB()
	book := MockBook()
	defer MockDeleteBook()

	list := service.GetBooksByGenre("Science Fiction")

	assert.NotContains(t, list, book, "The mock book is not present in list, because your genre is not Science Fiction")
}

func TestGetBooksByNameSucess(t *testing.T){
	SetupDB()
	book := MockBook()
	defer MockDeleteBook()

	list := service.GetBooksByName("Name book")

	assert.Contains(t, list, book)
}

func TestCreateBookSucess(t *testing.T){
	SetupDB()
	var book models.Book = models.Book{Name: "Name book", Author: "Emanuel Dias", Genre: "Fantasy", Description: "Any description Any description Any description Any description Any description Any description Any description"}

	book ,err := service.CreateBook(book)

	assert.Equal(t, "Name book", book.Name)
	assert.Equal(t, "Emanuel Dias", book.Author)
	assert.Equal(t, "Fantasy", book.Genre)
	assert.Equal(t, "Any description Any description Any description Any description Any description Any description Any description", book.Description)

	assert.Equal(t, false, book.Available)
	assert.Equal(t, nil, err)
}

func TestUpdateBookSucess(t *testing.T){
	SetupDB()
	MockBook()
	defer MockDeleteBook()

	var book models.Book = models.Book{Name: "Name edit", Author: "Emanuel", Genre: "Science Fiction", Description: "Any description Any description Any description Any description Any description Any description Any description"}

	result, _ := service.UpdateBook(book, ID)

	assert.Equal(t, "Name edit", result.Name)
	assert.Equal(t, "Emanuel", result.Author)
	assert.Equal(t, "Science Fiction", result.Genre)
	assert.Equal(t, "Any description Any description Any description Any description Any description Any description Any description", book.Description)
}

func TestDeleteBookSucess(t *testing.T){
	SetupDB()
	MockBook()

	service.DeleteBook(ID)

	_, err := service.GetBookById(ID)

	assert.Equal(t, "Book not found", err.Error(), "Delete book, search this book and not found")
}