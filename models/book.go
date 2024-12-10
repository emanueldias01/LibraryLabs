package models

import "gopkg.in/validator.v2"


type Book struct {
	Id int `json:"id"`
	Name string `json:"name" validate:"max=50, nonzero"`
	Author string `json:"author" validate:"max=30, nonzero"`
	Genre string `json:"genre" validate:"max=15", nonzero`
	Description string `json:"description" validate:"min=50, max=255, nonzero"`
	Available bool `json:"available"`
}

func ValidateFields(b *Book) error{
	if err := validator.Validate(b); err != nil{
		return err
	}
	return nil
}