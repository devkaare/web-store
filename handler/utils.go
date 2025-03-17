package handler

import (
	"database/sql"
	"encoding/json"
	"log"
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

func check(err error) bool {
	if err != nil {
		log.Printf("check: %v", err)
		return true
	}
	return false
}

func checkIfErrNoRows(err error) bool {
	if err != nil {
		if err == sql.ErrNoRows {
			return true
		}
		log.Printf("check: %v", err)
		return true
	}
	return false
}

func checkIfNotErrNoRows(err error) bool {
	if err != nil && err != sql.ErrNoRows {
		log.Printf("check: %v", err)
		return true
	}
	return false
}

func checkIfErrNoCookie(err error) bool {
	if err != nil {
		if err == http.ErrNoCookie {
			return true
		}
		log.Printf("check: %v", err)
		return true
	}
	return false
}

func checkForm(w http.ResponseWriter, fields []string) bool {
	for _, v := range fields {
		if v == "" {
			w.WriteHeader(http.StatusBadRequest)
			return true
		}
	}
	return false
}
