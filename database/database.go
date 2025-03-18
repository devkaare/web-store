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

func createTables(db *sql.DB) {
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			user_id SERIAL PRIMARY KEY,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS sessions (
			session_id TEXT PRIMARY KEY,
			user_id INT NOT NULL,
			expiry TIMESTAMP WITH TIME ZONE NOT NULL
		)
	`); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			category_id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
		)
	`); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			product_id SERIAL PRIMARY KEY,
			category_id INT NOT NULL,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			price INT NOT NULL,
			color TEXT NOT NULL,
			size TEXT NOT NULL,
			image_path TEXT NOT NULL
		)
	`); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS shopping_sessions (
			session_id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			total INT NOT NULL
		)
	`); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS cart_items (
			cart_item_id SERIAL PRIMARY KEY,
			shopping_session_id INT NOT NULL,
			product_id INT NOT NULL,
			quantity INT NOT NULL
		)
	`); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS order_details (
			order_details_id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			payment_id TEXT NOT NULL,
			total INT NOT NULL
		)
	`); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS order_items (
			order_item_id SERIAL PRIMARY KEY,
			order_details_id INT NOT NULL,
			product_id INT NOT NULL,
			quantity INT NOT NULL
		)
	`); err != nil {
		log.Fatal(err)
	}
}

func New() *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=UTC&search_path=%s", username, password, host, port, database, schema)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	createTables(db)

	return db
}
