package models


type Book struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	Genre string `json:"genre"`
	Description string `json:"description"`
	Available bool `json:"available"`
}