package controllers

import (
	"encoding/json"
	"fmt"
	"librarylabs/models"
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
	list := models.GetAllBooks()

	json.NewEncoder(w).Encode(list)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	idInt := GetIdInd(r)

	book, err := models.GetBookById(idInt)

	if err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	var bookCreate models.Book
	json.NewDecoder(r.Body).Decode(&bookCreate)

	models.CreateBook(bookCreate)
	json.NewEncoder(w).Encode(bookCreate)

	location := fmt.Sprintf("/books/%d", bookCreate.Id)
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	idInt := GetIdInd(r)

	var bodyBook models.Book

	json.NewDecoder(r.Body).Decode(&bodyBook)

	bookUpdate, err := models.UpdateBook(bodyBook, idInt)

	if err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(bookUpdate)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	idInt := GetIdInd(r)

	models.DeleteBook(idInt)
	w.WriteHeader(http.StatusNoContent)
}

func SetBookUnavailable(w http.ResponseWriter, r *http.Request){
	idInt := GetIdInd(r)
	var(
		book models.Book
		err error
	)
	book, err = models.SetBookUnavailable(idInt)

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

	
	book,err = models.SetBookAvailable(idInt)

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

func GetBooksByGenre(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	genre := vars["genre"]

	list := models.GetBooksByGenre(genre)

	json.NewEncoder(w).Encode(list)
}

func GetBooksByName(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	name := vars["name"]

	list := models.GetBooksByName(name)

	json.NewEncoder(w).Encode(list)
}