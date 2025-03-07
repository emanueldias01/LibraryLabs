package auth

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func(u* User) GenerateToken() (string, error){
	secret := "secret-key"

	claims := jwt.MapClaims{
		"id" : u.Id,
		"login" : u.Login,
		"exp" : time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStrig, err := token.SignedString([]byte(secret))

	if err != nil{
		return "", err
	}

	return tokenStrig, nil
}