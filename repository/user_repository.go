package repository

import (
	"errors"
	"librarylabs/auth"
	"librarylabs/database"
)

func CreateUser(u auth.User) error{
	db := database.DB
	err := db.Create(&u).Error

	if err != nil{
		return err
	}

	return nil
}

func GetUserById(id uint) (auth.User, error){
	var user auth.User
	db := database.DB
	db.Find(&user, id)

	if user.Login == ""{
		return auth.User{}, errors.New("user not found")
	}

	return user, nil
}