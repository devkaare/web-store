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
	if err != nil {
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
	password := r.FormValue("password")

	if firstName == "" || lastName == "" || email == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("<p>Missing required fields</p>"))
		return
	}

	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &model.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hashedPassword,
	}

	userID, err := u.Repo.CreateUser(user)
	if err != nil {
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
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
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
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = u.Repo.DeleteUserByUserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (u *User) UpdateUserByUserID(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	_, err := u.Repo.GetUserByUserID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &model.User{
		UserID:   userID,
		Email:    email,
		Password: hashedPassword,
	}

	err = u.Repo.UpdateUserByUserID(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
