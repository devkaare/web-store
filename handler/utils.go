package handler

import (
	"encoding/json"
	"net/http"

	"github.com/devkaare/web-store/repository/query"
)

type Utils struct {
	Repo *query.PostgresRepo
}

func (u *Utils) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(u.Repo.Health())
	_, _ = w.Write(jsonResp)
}
