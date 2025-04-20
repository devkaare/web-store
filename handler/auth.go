package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/devkaare/web-store/hash"
	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository"
	"github.com/devkaare/web-store/repository/session"
	"github.com/devkaare/web-store/repository/shopping_session"
	"github.com/devkaare/web-store/repository/user"
	"github.com/google/uuid"
)

type Authentication struct {
	UserRepo            *user.Repo
	SessionRepo         *session.Repo
	ShoppingSessionRepo *shoppingsession.Repo
}

var authenticationHandler = &Authentication{
	UserRepo:            &user.Repo{},
	SessionRepo:         &session.Repo{},
	ShoppingSessionRepo: &shoppingsession.Repo{},
}

func NewAuthenticationHandler(db *sql.DB) *Authentication {
	authenticationHandler.UserRepo = repository.GetUser(func() *sql.DB {
		return db
	})
	authenticationHandler.SessionRepo = repository.GetSession(func() *sql.DB {
		return db
	})
	authenticationHandler.ShoppingSessionRepo = repository.GetShoppingSession(func() *sql.DB {
		return db
	})
	return authenticationHandler
}

func (a *Authentication) SignUp(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	verifPass := r.FormValue("verif_password")

	if firstName == "" || lastName == "" || email == "" || password == "" || verifPass == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("<p>Missing required fields</p>"))
		return
	}

	if password != verifPass {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("<p>Please make sure that both passwords match!</p>"))
		return
	}

	existingUser, err := a.UserRepo.GetUserByEmail(email)
	if err != nil && err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !(existingUser.FirstName == "" || existingUser.LastName == "" || existingUser.Email == "" || existingUser.Password == "") {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("<p>User with email already exists</p>"))
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

	_, err = a.UserRepo.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/signin", http.StatusSeeOther)
}

func (a *Authentication) SignIn(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	expectedPassword := r.FormValue("password")

	if email == "" || expectedPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("<p>Missing required fields</p>"))
		return
	}

	existingUser, err := a.UserRepo.GetUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid email or password"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !hash.CheckPasswordHash(expectedPassword, existingUser.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid email or password"))
		return
	}

	sessionID := uuid.NewString()

	session := &model.Session{
		SessionID: sessionID,
		UserID:    existingUser.UserID,
	}

	err = a.SessionRepo.CreateSession(session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	expiry := time.Now().Add(20 * time.Minute)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionID,
		Path:     "/",
		Expires:  expiry,
		MaxAge:   86400,
		HttpOnly: true,
		Secure:   true,
	})

	w.Write([]byte("<p>Successfully logged in! Go to products <a href=\"/\">here</a></p>"))
}
