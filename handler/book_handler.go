package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/emanueldias01/LibraryLabs/dto"
	"github.com/emanueldias01/LibraryLabs/service"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	b, err := service.GetAllBooks()
	if err != nil{
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(b)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := parseId(r)

	if err != nil{
		writeError(w, http.StatusBadRequest, "ID is Invalid")
		return
	}

	b, err := service.GetBookById(&id)

	if err != nil{
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(b)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var book dto.BookRequest
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil{
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	b, err := service.CreateBook(&book)

	if err != nil{
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(b)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPut{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := parseId(r)

	if err != nil{
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	var book dto.BookRequest
	err = json.NewDecoder(r.Body).Decode(&book)

	if err != nil{
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	b, err := service.UpdateBook(book, id)

	if err != nil{
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(b)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodDelete{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := parseId(r)

	if err != nil{
		writeError(w, http.StatusBadRequest, "ID is invalid")
		return
	}

	err = service.DeleteBookById(id)

	if err != nil{
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func parseId(r *http.Request) (int, error){
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
    if len(parts) < 2 {
        return 0, fmt.Errorf("ID is not present")
    }
    id, err := strconv.Atoi(parts[len(parts)-1])
    if err != nil {
        return 0, fmt.Errorf("invalid ID")
    }
    return id, nil
}

func writeError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

