package model

import "github.com/emanueldias01/LibraryLabs/dto"

type Book struct {
	ID uint
	Name string
	Author string
	YearPublication uint
	Publisher string
	PagesNumber uint
	Language string
}

func(b *Book) UpdateInfo(dto dto.BookRequest){
	if dto.Name != ""{
		b.Name = dto.Name
	}
	if dto.Author != ""{
		b.Author = dto.Author
	}
	if dto.YearPublication != 0{
		b.YearPublication = dto.YearPublication
	}
	if dto.Publisher != ""{
		b.Publisher = dto.Publisher
	}
	if dto.PagesNumber != 0{
		b.PagesNumber = dto.PagesNumber
	}
	if dto.Language != ""{
		b.Language = dto.Language
	}
}