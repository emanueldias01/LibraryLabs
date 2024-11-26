package models

import(
	"librarylabs/database"
)

type Genre int

const (
	Fantasy Genre = iota
	ScienceFiction
	Drama
	Romance
	Adventure
)

func (g Genre) String() string {
	switch g {
	case Fantasy:
		return "Fantasy"
	case ScienceFiction:
		return "Science Fiction"
	case Drama:
		return "Drama"
	case Romance:
		return "Romance"
	case Adventure:
		return "Adventure"
	default:
		return "Unknown"
	}
}

type Book struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	Genre Genre `json:"genre"`
	Description string `json:"description"`
	Avaliable bool `json:"avaliable"`
}

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

	if bodyBook.Genre.String() != "" {
		bookUpdate.Genre = bodyBook.Genre
	}

	//Save new changes
	database.DB.Save(&bookUpdate)

	return bookUpdate
}

func DeleteBook(id int){
	var bookDelete Book
	database.DB.Delete(&bookDelete, id)
}