package server

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var r http.Handler
var req *http.Request
var err error
var respRec *httptest.ResponseRecorder
var db *sql.DB

func setup() {
	testServer := &Server{
		port: 3000,
		db:   db,
	}
	r = testServer.RegisterRoutes()
	respRec = httptest.NewRecorder()
}

func TestConnectToDB(t *testing.T) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", "kaare", "password", "localhost", "5432", "webstore", "public")
	db, err = sql.Open("pgx", connStr)
	if err != nil {
		t.Fatalf("TestConnectToDB: %v", err)
	}
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
