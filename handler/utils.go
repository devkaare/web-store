package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/devkaare/web-store/repository"
	"github.com/devkaare/web-store/repository/utils"
	_ "github.com/joho/godotenv"
)

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
