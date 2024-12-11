package repository

import (
	"librarylabs/database"
	"librarylabs/models"
)

func Create(b *models.Book){
	database.DB.Create(&b)
}

func Find(b *models.Book,id int){
	database.DB.First(&b, id)
}

func ListAll() []models.Book{
	var list []models.Book
	database.DB.Find(&list)
	return list
}

func Update(bref models.Book, bbody models.Book){
	database.DB.Model(&bref).UpdateColumns(bbody)
}

func Delete(id int){
	var b *models.Book
	database.DB.Delete(&b, id)
}

func Save(b *models.Book){
	database.DB.Save(&b)
}

func FindByGenre(genre string) []models.Book{
	var l []models.Book
	database.DB.Where(&models.Book{Genre: genre}).Find(&l)
	return l
}

func FindByName(name string) []models.Book{
	var l []models.Book
	database.DB.Raw("SELECT * FROM books WHERE name LIKE ?", "%"+name+"%").Scan(&l)
	return l
}