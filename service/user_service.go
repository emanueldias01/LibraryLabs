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

func LoginUser(u auth.User) (string, error){
	u.Password = auth.PasswordEncode(u.Password)

	err := repository.ComparatorUser(u)

	if err != nil{
		return  "",err
	}

	token, err := u.GenerateToken()

	if err != nil{
		return "", err
	}

	return token,nil
}