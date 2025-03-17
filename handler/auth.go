package handler

import (
	"context"
	"database/sql"
	"log"
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
	verifPassword := r.FormValue("verif_password")

	if firstName == "" || lastName == "" || email == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if password != verifPassword {
		w.Write([]byte("Please make sure that both passwords match!"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := userHandler.Repo.GetUserByEmail(email)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &model.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hashedPassword,
	}

	if _, err := userHandler.Repo.CreateUser(user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a *Authentication) SignIn(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	expectedPassword := r.FormValue("password")

	if email == "" || expectedPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	existingUser, err := userHandler.Repo.GetUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if hash.CheckPasswordHash(expectedPassword, existingUser.Password) {
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

	if err := sessionHandler.Repo.CreateSession(session); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionID,
		Expires: expiresAt,
	})
}

func (a *Authentication) SetUserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		sessionID := cookie.Value

		session, err := sessionHandler.Repo.GetSessionBySessionID(sessionID)
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if isExpired(session) {
			if err := sessionHandler.Repo.DeleteSessionBySessionID(sessionID); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", session.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
