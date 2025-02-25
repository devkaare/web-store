package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Health() map[string]string

	Close() error
}

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
	schema   = os.Getenv("DB_SCHEMA")
)

func New() *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS users (user_id SERIAL PRIMARY KEY, email TEXT NOT NULL, password TEXT NOT NULL)"); err != nil {
		log.Fatal(err)
	}

	// TODO: Change sizes to use postgres arrays
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS products (product_id SERIAL PRIMARY KEY, name TEXT NOT NULL, price INT NOT NULL, sizes TEXT NOT NULL, image_path TEXT NOT NULL)"); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS cart_items (user_id INT PRIMARY KEY, product_id INT NOT NULL, sizes TEXT NOT NULL, quantity INT NOT NULL)"); err != nil {
		log.Fatal(err)
	}

	return db
}
