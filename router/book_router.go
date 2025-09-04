package router

import (
	"net/http"

	"github.com/emanueldias01/LibraryLabs/handler"
)

func booksHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        handler.GetAllBooks(w, r)
    case http.MethodPost:
        handler.CreateBook(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func bookByIDHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        handler.GetBookById(w, r)
    case http.MethodPut:
        handler.UpdateBook(w, r)
    case http.MethodDelete:
        handler.DeleteBook(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func InitializeRoutes(){
	http.HandleFunc("/api/v1/books", booksHandler)
    http.HandleFunc("/api/v1/books/", bookByIDHandler)

    http.ListenAndServe(":8000", nil)
}