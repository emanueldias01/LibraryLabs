package routes

import (
	"librarylabs/controllers"
	"net/http"

	"github.com/gorilla/mux"
)


var r = mux.NewRouter()

func HandleRequest(){
	r.HandleFunc("/books", controllers.GetAllBooks)
	r.HandleFunc("/books/create", controllers.CreateBook)
	r.HandleFunc("/books/update/{id}", controllers.UpdateBook)
	r.HandleFunc("/books/delete/{id}", controllers.DeleteBook)
	r.HandleFunc("/books/unavailable/{id}", controllers.SetBookUnavailable)
	r.HandleFunc("/books/available/{id}", controllers.SetBookAvailable)
	http.ListenAndServe(":8000", r)
}