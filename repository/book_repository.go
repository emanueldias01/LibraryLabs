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

func Update(br models.Book, bb models.Book){
	database.DB.Model(&br).UpdateColumns(bb)
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