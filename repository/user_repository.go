package repository

import (
	"errors"

	"librarylabs/auth"
	"librarylabs/database"
)

func CreateUser(u auth.User) error{
	db := database.DB

	var user auth.User
	err := db.Where("login = ?", u.Login).First(&user).Error

	if err == nil {
		return errors.New("this user exists")
	}

	if err := db.Create(&u).Error; err != nil{
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

func ComparatorUser(u auth.User) error{

	var user auth.User
	db := database.DB
	err := db.Where("login = ?", u.Login).First(&user).Error

	if err != nil{
		return err
	}

	if u.Password != user.Password{
		return errors.New("password invalid")
	}

	return nil
	
}