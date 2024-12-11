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

func FindByGenre(l *[]models.Book, genre string){
	database.DB.Where(&models.Book{Genre: genre}).Find(&l)
}

func FindByName(l *[]models.Book, name string){
	database.DB.Raw("SELECT * FROM books WHERE name LIKE ?", "%"+name+"%").Scan(&l)
}