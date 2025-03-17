package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/devkaare/web-store/hash"
	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository"
	"github.com/devkaare/web-store/repository/user"
	"github.com/go-chi/chi/v5"
)

type User struct {
	Repo *user.Repo
}

var userHandler = &User{
	Repo: &user.Repo{},
}

func NewUserHandler(db *sql.DB) *User {
	userHandler.Repo = repository.GetUser(func() *sql.DB {
		return db
	})
	return userHandler
}

func (u *User) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.Repo.GetAllUsers()
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(users)
	_, _ = w.Write(jsonResp)
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	pass := r.FormValue("password")

	if checkForm(w, []string{firstName, lastName, email, pass}) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	passHash, err := hash.HashPass(pass)
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &model.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  passHash,
	}

	userID, err := u.Repo.CreateUser(user)
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.UserID = userID

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(user)
	_, _ = w.Write(jsonResp)
}

func (u *User) GetUserByUserID(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	user, err := u.Repo.GetUserByUserID(userID)
	if checkIfNotErrNoRows(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(user)
	_, _ = w.Write(jsonResp)
}

func (u *User) DeleteUserByUserID(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	_, err := u.Repo.GetUserByUserID(userID)
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = u.Repo.DeleteUserByUserID(userID)
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (u *User) UpdateUserByUserID(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	_, err := u.Repo.GetUserByUserID(userID)
	if checkIfNotErrNoRows(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	email := r.FormValue("email")
	pass := r.FormValue("password")

	passHash, err := hash.HashPass(pass)
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &model.User{
		UserID:   userID,
		Email:    email,
		Password: passHash,
	}

	err = u.Repo.UpdateUserByUserID(user)
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
