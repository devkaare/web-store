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

	apiUrl := "http://localhost:3000"
	resource := "/users/"
	rawData := url.Values{}
	rawData.Set("email", "willsmithspersonalemail@gmail.com")
	rawData.Set("password", "supersecret123")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	fmt.Println(strings.NewReader(rawData.Encode()))
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(rawData.Encode()))
	if err != nil {
		t.Fatalf("TestCreateUser: %v", err)
	}

	// req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(respRec, req)

	fmt.Println(respRec.Result().Status)
}

func TestUpdateUser(t *testing.T) {
	setup()

	apiUrl := "http://localhost:3000"
	resource := "/users/1"
	rawData := url.Values{}
	rawData.Set("email", "willsmithspersonalemail@gmail.com")
	rawData.Set("password", "supersecret123")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	fmt.Println(strings.NewReader(rawData.Encode()))
	req, err := http.NewRequest("PUT", urlStr, strings.NewReader(rawData.Encode()))
	if err != nil {
		t.Fatalf("TestUpdateUser: %v", err)
	}

	// req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(respRec, req)

	fmt.Println(respRec.Result().Status)
}

func TestDeleteUser(t *testing.T) {
	setup()

	req, err = http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatalf("TestDeleteUser: %v", err)
	}

	r.ServeHTTP(respRec, req)

	fmt.Println(respRec.Result().Status)
}
