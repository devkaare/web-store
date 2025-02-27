package server

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var port = 3000

var (
	r       http.Handler
	req     *http.Request
	err     error
	respRec *httptest.ResponseRecorder
	db      *sql.DB
)

func setup() {
	testServer := &Server{
		port: port,
		db:   db,
	}
	r = testServer.RegisterRoutes()
	respRec = httptest.NewRecorder()
}

func TestConnectToDB(t *testing.T) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", "kaare", "password", "localhost", "5432", "dbwebstore", "public")
	db, err = sql.Open("pgx", connStr)
	if err != nil {
		t.Fatalf("TestConnectToDB: %v", err)
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS users (user_id SERIAL PRIMARY KEY, email TEXT NOT NULL, password TEXT NOT NULL)"); err != nil {
		t.Fatalf("TestConnectToDB: %v", err)
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS products (product_id SERIAL PRIMARY KEY, name TEXT NOT NULL, price INT NOT NULL, sizes TEXT NOT NULL, image_path TEXT NOT NULL)"); err != nil {
		t.Fatalf("TestConnectToDB: %v", err)
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS cart_items (user_id INT PRIMARY KEY, product_id INT NOT NULL, size TEXT NOT NULL, quantity INT NOT NULL)"); err != nil {
		t.Fatalf("TestConnectToDB: %v", err)
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS sessions (session_id TEXT NOT NULL, user_id INT NOT NULL, expiry TIMESTAMP WITH TIME ZONE NOT NULL)"); err != nil {
		t.Fatalf("TestConnectToDB: %v", err)
	}
}

func TestMain(m *testing.M) {
	setup()
	TestConnectToDB(&testing.T{})
	m.Run()
}

func TestHealth(t *testing.T) {
	setup()

	req, err = http.NewRequest("GET", "/utils/health", nil)
	if err != nil {
		t.Fatalf("TestHealth: %v", err)
	}

	r.ServeHTTP(respRec, req)

	if respRec.Code != http.StatusOK {
		t.Fatalf("TestHealth: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}

	result := respRec.Result().Body
	data, err := io.ReadAll(result)
	if err != nil {
		t.Fatalf("TestHealth: %v", err)
	}

	fmt.Println(string(data))
}
