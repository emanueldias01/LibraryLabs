package routes

import (
	"librarylabs/controllers"
	"net/http"

	"github.com/gorilla/mux"
)


var r = mux.NewRouter()

func HandleRequest(){
	r.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/books/create", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/update/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/delete/{id}", controllers.DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/unavailable/{id}", controllers.SetBookUnavailable).Methods("PUT")
	r.HandleFunc("/books/available/{id}", controllers.SetBookAvailable).Methods("PUT")
	http.ListenAndServe(":8000", r)
}