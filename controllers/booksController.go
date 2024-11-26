package controllers

import (
	"encoding/json"
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

	book := models.GetBookById(idInt)

	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	var bookCreate models.Book
	json.NewDecoder(r.Body).Decode(&bookCreate)

	models.CreateBook(bookCreate)
	json.NewEncoder(w).Encode(bookCreate)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	idInt := GetIdInd(r)

	var bodyBook models.Book
	json.NewDecoder(r.Body).Decode(&bodyBook)

	bookUpdate := models.UpdateBook(bodyBook, idInt)

	json.NewEncoder(w).Encode(bookUpdate)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	idInt := GetIdInd(r)

	models.DeleteBook(idInt)
}