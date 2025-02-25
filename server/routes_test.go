package server

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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

func TestGetUsers(t *testing.T) {
	setup()

	req, err = http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatalf("TestGetUsers: %v", err)
	}

	r.ServeHTTP(respRec, req)

	if respRec.Code != http.StatusOK {
		t.Fatalf("TestGetUsers: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}

	result := respRec.Result().Body
	data, err := io.ReadAll(result)
	if err != nil {
		t.Fatalf("TestGetUsers: %v", err)
	}

	fmt.Println(string(data))
}

func TestCreateUser(t *testing.T) {
	setup()

	resource := "/users/"
	rawData := url.Values{}
	rawData.Set("email", "willsmith@gmail.com")
	rawData.Set("password", "secret123")

	u, _ := url.ParseRequestURI("http://localhost:3000")
	u.Path = resource
	urlStr := u.String()

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(rawData.Encode()))
	// req.Header.Add("Authorization")
	if err != nil {
		t.Fatalf("TestCreateUser: %v", err)
	}

	r.ServeHTTP(respRec, req)

	if respRec.Code != http.StatusOK {
		t.Fatalf("TestCreateUser: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}

	result := respRec.Result().Body
	data, err := io.ReadAll(result)
	if err != nil {
		t.Fatalf("TestCreateUser: %v", err)
	}

	fmt.Println(string(data))
}
