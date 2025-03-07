package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/devkaare/web-store/hash"
	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository/user"
	"github.com/go-chi/chi/v5"
)

type User struct {
	UserRepo *user.UserRepo
}

func (u *User) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.UserRepo.GetAllUsers()
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
	reqApiKey := r.URL.Query().Get("api_key")
	if reqApiKey != apiKey {
		w.WriteHeader(http.StatusUnauthorized)
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
		Email:    email,
		Password: passwordHash,
	}

	userID, err := u.UserRepo.CreateUser(user)
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
	userID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	user, err := u.UserRepo.GetUserByUserID(userID)
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
	reqApiKey := r.URL.Query().Get("api_key")
	if reqApiKey != apiKey {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	if _, err := u.UserRepo.GetUserByUserID(userID); err != nil && err != sql.ErrNoRows {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := u.UserRepo.DeleteUserByUserID(userID); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (u *User) UpdateUserByUserID(w http.ResponseWriter, r *http.Request) {
	reqApiKey := r.URL.Query().Get("api_key")
	if reqApiKey != apiKey {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	if _, err := u.UserRepo.GetUserByUserID(userID); err != nil && err != sql.ErrNoRows {
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
		UserID:   userID,
		Email:    email,
		Password: passwordHash,
	}

	if err := u.UserRepo.UpdateUserByUserID(user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
