package service

import (
	"librarylabs/auth"
	"librarylabs/repository"
)

func CreateUser(u auth.User)(error){
	u.Password = auth.PasswordEncode(u.Password)
	err := repository.CreateUser(u)

	if err != nil{
		return err
	}

	return nil
}