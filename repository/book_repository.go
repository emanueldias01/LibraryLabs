package repository

import (
	"context"
	"fmt"
	"github.com/emanueldias01/LibraryLabs/db"
	"github.com/emanueldias01/LibraryLabs/model"
)

func AllBooks() (*[]model.Book, error){
	conn, err := db.OpenConnection()

	if err != nil{
		return nil, fmt.Errorf("falied to connect database")
	}

	rows, err := conn.Query(context.Background(),"SELECT id, name, author, year_publication, publisher, pages_number, language FROM books")

	if err != nil{
		return nil, err
	}

	defer db.CloseConnection(conn)

	var books []model.Book
	for rows.Next(){
		var b model.Book
		if err := rows.Scan(&b.ID,
			&b.Name,
			&b.Author,
			&b.YearPublication,
			&b.Publisher,
			&b.PagesNumber,
			&b.Language); err != nil{
				return nil, err
			}
		books = append(books, b)
	}

	return &books, nil
}

func GetBookById(id int)(*model.Book, error){
	conn, err := db.OpenConnection()

	if err != nil{
		return nil, fmt.Errorf("falied to connect database")
	}

	defer db.CloseConnection(conn)

	var b model.Book
	err = conn.QueryRow(
		context.Background(),
		"SELECT id, name, author, year_publication, publisher, pages_number, language FROM books WHERE id=$1",
		id,
	).Scan(
		&b.ID,
		&b.Name,
		&b.Author,
		&b.YearPublication,
		&b.Publisher,
		&b.PagesNumber,
		&b.Language,
	)

	if err != nil {
		return nil, fmt.Errorf("book not found: %w", err)
	}

	return &b, nil
}

func CreateBook(b model.Book)(*model.Book, error){
	conn, err := db.OpenConnection()

	if err != nil{
		return nil, fmt.Errorf("falied to connect database")
	}

	defer db.CloseConnection(conn)

	err = conn.QueryRow(context.Background(), 
	"INSERT INTO books (name, author, year_publication, publisher, pages_number, language) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
	b.Name, b.Author, b.YearPublication, b.Publisher, b.PagesNumber, b.Language).Scan(&b.ID)

	if err != nil {
		return nil, err
	}

	return &b, nil
}

func UpdateBook(b model.Book) error{
	conn, err := db.OpenConnection()
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}
	defer db.CloseConnection(conn)

	cmdTag, err := conn.Exec(
		context.Background(),
		`UPDATE books
		 SET name=$1, author=$2, year_publication=$3, publisher=$4, pages_number=$5, language=$6
		 WHERE id=$7`,
		b.Name, b.Author, b.YearPublication, b.Publisher, b.PagesNumber, b.Language, b.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("book with ID %d not found", b.ID)
	}

	return nil
}

func DeleteBook(id uint) error {
	conn, err := db.OpenConnection()
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}
	defer db.CloseConnection(conn)

	cmdTag, err := conn.Exec(
		context.Background(),
		"DELETE FROM books WHERE id=$1",
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("book with ID %d not found", id)
	}

	return nil
}