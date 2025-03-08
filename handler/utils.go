package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/devkaare/web-store/repository"
	"github.com/devkaare/web-store/repository/utils"
	_ "github.com/joho/godotenv"
)

// var apiKey = os.Getenv("API_KEY")
var apiKey = "81566e986cf8cc685a05ac5b634af7f8"

type Utils struct {
	Repo *utils.Repo
}

var utilsHandler = &Utils{
	Repo: &utils.Repo{},
}

func NewUtilsHandler(db *sql.DB) *Utils {
	utilsHandler.Repo = repository.GetUtils(func() *sql.DB {
		return db
	})
	return utilsHandler
}

func (u *Utils) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(u.Repo.Health())
	_, _ = w.Write(jsonResp)
}
