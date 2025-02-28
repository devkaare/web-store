package handler

import (
	"encoding/json"
	"net/http"
	// "os"

	"github.com/devkaare/web-store/repository/query"
	_ "github.com/joho/godotenv"
)

// var apiKey = os.Getenv("API_KEY")
var apiKey = "81566e986cf8cc685a05ac5b634af7f8"

type Utils struct {
	Repo *query.PostgresRepo
}

func (u *Utils) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(u.Repo.Health())
	_, _ = w.Write(jsonResp)
}
