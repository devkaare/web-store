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
	pass := r.FormValue("password")
	verifPass := r.FormValue("verif_password")

	if checkForm(w, []string{firstName, lastName, email, pass, verifPass}) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if pass != verifPass {
		w.Write([]byte("Please make sure that both passwords match!"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := userHandler.Repo.GetUserByEmail(email)
	if checkIfNotErrNoRows(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hashedPass, err := hash.HashPass(pass)
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &model.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hashedPass,
	}

	_, err = userHandler.Repo.CreateUser(user)
	if check(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a *Authentication) SignIn(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	expectedPass := r.FormValue("password")

	if checkForm(w, []string{email, expectedPass}) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	existingUser, err := userHandler.Repo.GetUserByEmail(email)
	if checkIfErrNoRows(err) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if hash.CheckPassHash(expectedPass, existingUser.Password) {
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
}
