package controllers

import (
	"encoding/json"
	"librarylabs/auth"
	"librarylabs/service"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
	var user auth.User
	json.NewDecoder(r.Body).Decode(&user)

	err := service.CreateUser(user)
	
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request){
	var user auth.User
	json.NewDecoder(r.Body).Decode(&user)

	token, err := service.LoginUser(user)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

