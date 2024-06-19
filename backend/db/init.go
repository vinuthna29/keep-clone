package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() error {
	var err error
	// opening a connection
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	// connStr := "postgres://vinuthnavenigalla:vinu2907@localhost:5432/postgres?sslmode=disable"

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	//creating a table in db if not present
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS notes(
			id SERIAL PRIMARY KEY,
			title TEXT,
			content TEXT
		)
	`)
	return err
}

func CloseDB() {
	db.Close()
}

func CreateNote(title, content string) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO notes(title, content) VALUES ($1, $2) RETURNING id", title, content).Scan(&id)
	return id, err
}
