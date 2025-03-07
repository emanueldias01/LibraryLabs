package auth

type User struct {
	Id uint `json:"id"`
	Login string `json:"login"`
	Password string `json:"password"`
}