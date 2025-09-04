package service

import (
	"github.com/emanueldias01/LibraryLabs/dto"
	"github.com/emanueldias01/LibraryLabs/model"
	"github.com/emanueldias01/LibraryLabs/repository"
)

func GetAllBooks()(*[]dto.BookResponse, error){
	b, err := repository.AllBooks()
	if err != nil{
		return nil, err
	}

	var books []dto.BookResponse
	for _, e := range *b{
		b := dto.BookResponse{e.ID, e.Name, e.Author, e.YearPublication, e.Publisher, e.PagesNumber, e.Language}
		books = append(books, b)
	}

	return &books, nil
}

func GetBookById(id *int)(*dto.BookResponse, error){
	e, err := repository.GetBookById(*id)
	if err != nil{
		return nil, err
	}

	book := dto.BookResponse{e.ID, e.Name, e.Author, e.YearPublication, e.Publisher, e.PagesNumber, e.Language}

	return &book, nil
}

func CreateBook(dtoReq *dto.BookRequest)(*dto.BookResponse, error){
	b := model.Book{0, dtoReq.Name, dtoReq.Author, dtoReq.YearPublication, dtoReq.Publisher, dtoReq.PagesNumber, dtoReq.Language}
	e, err := repository.CreateBook(&b)

	if err != nil{
		return nil, err
	}

	book := dto.BookResponse{e.ID, e.Name, e.Author, e.YearPublication, e.Publisher, e.PagesNumber, e.Language}

	return &book, nil
}

func UpdateBook(dtoReq dto.BookRequest, id int)(*dto.BookResponse, error){
	b, err := repository.GetBookById(id)

	if err != nil{
		return nil, err
	}

	b.UpdateInfo(dtoReq)

	err = repository.UpdateBook(b)

	if err != nil{
		return nil, err
	}

	return GetBookById(&id)
}

func DeleteBookById(id int) error{
	err := repository.DeleteBook(&id)

	if err != nil{
		return err
	}

	return nil
}