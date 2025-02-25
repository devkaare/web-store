package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/devkaare/web-store/auth"
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

	// jsonResp, _ := json.Marshal(users)
	// _, _ = w.Write(jsonResp)

	fmt.Fprintln(w, users)
}

// curl -X POST localhost:3000/users -d "email=johndoe@gmail.com&password=secretpass8888"
func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	passwordHash, err := auth.HashPassword(password)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &model.User{
		Email:    email,
		Password: passwordHash,
	}

	// TODO: Use returned userID for auth
	if _, err := u.Repo.CreateUser(user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (u *User) GetUserByUserID(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "ID"))

	user, ok, err := u.Repo.GetUserByUserID(uint32(userID))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !ok {
		fmt.Fprintf(w, "user with user_id: %d does not exist", userID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, user)
}

func (u *User) DeleteUserByUserID(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "ID"))

	_, ok, err := u.Repo.GetUserByUserID(uint32(userID))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !ok {
		fmt.Fprintf(w, "user with user_id: %d does not exist", userID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := u.Repo.DeleteUserByUserID(uint32(userID)); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// curl -X PUT localhost:3000/users/1 -d "email=newjohndoe@gmail.com&password=newPassword"
func (u *User) UpdateUserByUserID(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "ID"))

	_, ok, err := u.Repo.GetUserByUserID(uint32(userID))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !ok {
		fmt.Fprintf(w, "user with user_id: %d does not exist", userID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	passwordHash, err := auth.HashPassword(password)
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
