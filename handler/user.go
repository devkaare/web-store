// TODO: Add protected routes using `req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")`
package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/devkaare/web-store/hash"
	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository/query"
	"github.com/go-chi/chi/v5"
)

type User struct {
	Repo *query.PostgresRepo
}

func (u *User) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.Repo.GetUsers()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(users)
	_, _ = w.Write(jsonResp)
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	passwordHash, err := hash.HashPassword(password)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &model.User{
		Email:    email,
		Password: passwordHash,
	}

	userID, err := u.Repo.CreateUser(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.UserID = userID

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(user)
	_, _ = w.Write(jsonResp)
}

func (u *User) GetUserByUserID(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "ID"))

	user, err := u.Repo.GetUserByUserID(uint32(userID))
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(user)
	_, _ = w.Write(jsonResp)
}

func (u *User) DeleteUserByUserID(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "ID"))

	if _, err := u.Repo.GetUserByUserID(uint32(userID)); err != nil && err != sql.ErrNoRows {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := u.Repo.DeleteUserByUserID(uint32(userID)); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (u *User) UpdateUserByUserID(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "ID"))

	if _, err := u.Repo.GetUserByUserID(uint32(userID)); err != nil && err != sql.ErrNoRows {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	passwordHash, err := hash.HashPassword(password)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &model.User{
		UserID:   uint32(userID),
		Email:    email,
		Password: passwordHash,
	}

	if err := u.Repo.UpdateUserByUserID(user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
