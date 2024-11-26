package controllers

import (
	"encoding/json"
	"librarylabs/database"
	"librarylabs/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request){
	var list []models.Book
	database.DB.Find(&list)

	json.NewEncoder(w).Encode(list)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	var bookCreate models.Book
	json.NewDecoder(r.Body).Decode(&bookCreate)

	database.DB.Create(&bookCreate)

	json.NewEncoder(w).Encode(bookCreate)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	id := vars["id"]

	idInt, err := strconv.Atoi(id)

	if err != nil{
		panic(err.Error())
	}

	//decode bodyBook
	var bookBody models.Book
	json.NewDecoder(r.Body).Decode(&bookBody)

	//Find the book i want to update
	var bookUpdate models.Book
	database.DB.First(&bookUpdate, idInt)

	//Change atributes book
	if bookBody.Name != "" {
		bookUpdate.Name = bookBody.Name
	}

	if bookBody.Author != "" {
		bookUpdate.Author = bookBody.Author
	}

	if bookBody.Description != "" {
		bookUpdate.Description = bookBody.Description
	}

	if bookBody.Genre.String() != "" {
		bookUpdate.Genre = bookBody.Genre
	}

	//Save new changes
	database.DB.Save(&bookUpdate)
	json.NewEncoder(w).Encode(bookUpdate)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	id := vars["id"]

	idInt, err := strconv.Atoi(id)

	if err != nil{
		panic(err.Error())
	}

	var bookDelete models.Book
	database.DB.Delete(&bookDelete, idInt)
}