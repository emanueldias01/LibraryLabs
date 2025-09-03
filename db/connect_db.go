package db

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5"
)

func OpenConnection() (*pgx.Conn, error){
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	return conn, nil
}

func CloseConnection(conn *pgx.Conn){
	conn.Close(context.Background())
}

func RunMigrations(conn *pgx.Conn) error {
	query := `
	CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		name VARCHAR NOT NULL,
		author VARCHAR NOT NULL,
		year_publication INTEGER,
		publisher VARCHAR,
		pages_number INTEGER,
		language VARCHAR
	);
	`

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("Erro in migrations: %w", err)
	}

	defer conn.Close(context.Background())

	return nil
}
