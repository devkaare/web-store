package handler

import (
	"net/http"
	"time"

	"github.com/devkaare/web-store/hash"
	"github.com/devkaare/web-store/model"
	"github.com/google/uuid"
)

type Authentication struct{}

func (a *Authentication) SignUp(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	verifPass := r.FormValue("verif_password")

	if checkValues([]string{firstName, lastName, email, password, verifPass}) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if password != verifPass {
		w.Write([]byte("<p>Please make sure that both passwords match!</p>"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	existingUser, err := userHandler.Repo.GetUserByEmail(email)
	if checkIfNotErrNoRows(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !checkValues([]string{existingUser.FirstName, existingUser.LastName, existingUser.Email, existingUser.Password}) {
		w.WriteHeader(http.StatusConflict)
		return
	}

	hashedPassword, err := hash.HashPassword(password)
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &model.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hashedPassword,
	}

	_, err = userHandler.Repo.CreateUser(user)
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *Authentication) SignIn(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	expectedPassword := r.FormValue("password")

	if checkValues([]string{email, expectedPassword}) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	existingUser, err := userHandler.Repo.GetUserByEmail(email)
	if checkIfNotErrNoRows(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !hash.CheckPasswordHash(expectedPassword, existingUser.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionID := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	session := &model.Session{
		SessionID: sessionID,
		UserID:    existingUser.UserID,
		Expiry:    expiresAt,
	}

	err = sessionHandler.Repo.CreateSession(session)
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionID,
		Expires: expiresAt,
	})

	w.WriteHeader(http.StatusOK)
}
