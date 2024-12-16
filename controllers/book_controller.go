package controllers

import (
	"encoding/json"
	"fmt"
	"librarylabs/models"
	"librarylabs/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetIdInd(r *http.Request) int{
	vars := mux.Vars(r)

	id := vars["id"]

	idInt, err := strconv.Atoi(id)

	if err != nil{
		panic(err.Error())
	}

	return idInt
}

func GetAllBooks(w http.ResponseWriter, r *http.Request){
	list := service.GetAllBooks()

	json.NewEncoder(w).Encode(list)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	idInt := GetIdInd(r)

	book, err := service.GetBookById(idInt)

	if err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	var bookCreate models.Book
	json.NewDecoder(r.Body).Decode(&bookCreate)

	book, err := service.CreateBook(bookCreate)
	
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(book)

	location := fmt.Sprintf("/books/%d", bookCreate.Id)
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	idInt := GetIdInd(r)

	var bodyBook models.Book

	json.NewDecoder(r.Body).Decode(&bodyBook)

	bookUpdate, err := service.UpdateBook(bodyBook, idInt)

	if err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(bookUpdate)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	idInt := GetIdInd(r)

	service.DeleteBook(idInt)
	w.WriteHeader(http.StatusNoContent)
}

func SetBookUnavailable(w http.ResponseWriter, r *http.Request){
	idInt := GetIdInd(r)
	var(
		book models.Book
		err error
	)
	book, err = service.SetBookUnavailable(idInt)

	if err != nil{
		if(err.Error() == "Book not found"){
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if(err.Error() == "This book is already Unavailable"){
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		
	}

	json.NewEncoder(w).Encode(book)
}

func SetBookAvailable(w http.ResponseWriter, r *http.Request){
	idInt := GetIdInd(r)

	var(
		book models.Book
		err error
	)

	
	book,err = service.SetBookAvailable(idInt)


	if err != nil{
		if(err.Error() == "Book not found"){
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if(err.Error() == "This book is already Avaliable"){
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		
	}

	json.NewEncoder(w).Encode(book)
}

func GetBooksByGenre(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	genre := vars["genre"]

	list := service.GetBooksByGenre(genre)

	json.NewEncoder(w).Encode(list)
}

func GetBooksByName(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	name := vars["name"]

	list := service.GetBooksByName(name)

	json.NewEncoder(w).Encode(list)
}