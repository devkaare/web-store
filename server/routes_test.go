package server

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var (
	port = 3000

	r       http.Handler
	req     *http.Request
	err     error
	respRec *httptest.ResponseRecorder
	db      *sql.DB

	testDatabase string
	testPassword string
	testUsername string
	testDBPort   string
	testHost     string
	testSchema   string
	testApiKey   string
)

func setup() {
	err := godotenv.Load("./../.env")
	if err != nil {
		panic(fmt.Errorf("setup: %v", err))
	}

	testDatabase = os.Getenv("DB_DATABASE")
	testPassword = os.Getenv("DB_PASSWORD")
	testUsername = os.Getenv("DB_USERNAME")
	testDBPort = os.Getenv("DB_PORT")
	testHost = os.Getenv("DB_HOST")
	testSchema = os.Getenv("DB_SCHEMA")
	testApiKey = os.Getenv("API_KEY")

	testServer := &Server{
		port: port,
		db:   db,
	}

	r = testServer.RegisterRoutes()
	respRec = httptest.NewRecorder()
}

func TestConnectToDB(t *testing.T) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=UTC&search_path=%s", testUsername, testPassword, testHost, testDBPort, testDatabase, testSchema)

	db, err = sql.Open("pgx", connStr)
	if err != nil {
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
