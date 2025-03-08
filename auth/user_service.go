package auth

import (
	"errors"
	"strings"
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

func VerifyToken(tokenVerify string) (User, error) {
    secret := "secret-key"
    
    token, err := jwt.Parse(removePrefixBearer(tokenVerify), func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
            return []byte(secret), nil
        }
        return nil, errors.New("invalid token")
    })

    if err != nil {
        return User{}, errors.New("invalid token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return User{}, errors.New("invalid token")
    }

    idFloat, ok := claims["id"].(float64)
    if !ok {
        return User{}, errors.New("invalid token: Id is missing or incorrect type")
    }

    login, ok := claims["login"].(string)
    if !ok || login == "" {
        return User{}, errors.New("invalid token: login is missing or incorrect type")
    }

    user := User{
        Id:       uint(idFloat),
        Login:    login,
    }

    return user, nil
}


func removePrefixBearer(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer ")
	}
	return token
}