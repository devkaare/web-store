package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/devkaare/web-store/auth"
	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository/query"
	"github.com/google/uuid"
)

type Session struct {
	Repo *query.PostgresRepo
}

func isExpired(s *model.Session) bool {
	return s.Expiry.Before(time.Now())
}

func (s *Session) SignIn(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	expectedPassword := r.FormValue("password")

	user, ok, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !ok || auth.CheckPasswordHash(expectedPassword, user.Password) {
		fmt.Fprintln(w, "invalid credentials")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionID := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	session := &model.Session{
		SessionID: sessionID,
		UserID:    user.UserID,
		Expiry:    expiresAt,
	}

	if err := s.Repo.CreateSession(session); err != nil {
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

func (s *Session) Welcome(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sessionID := c.Value

	session, ok, err := s.Repo.GetSessionBySessionID(sessionID)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if isExpired(session) {
		if err := s.Repo.DeleteSessionBySessionID(sessionID); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "welcome user: %d", session.UserID)
}

func (s *Session) Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sessionID := c.Value

	session, ok, err := s.Repo.GetSessionBySessionID(sessionID)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if isExpired(session) {
		if err := s.Repo.DeleteSessionBySessionID(sessionID); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	newSessionID := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	newSession := &model.Session{
		SessionID: newSessionID,
		UserID:    session.UserID,
		Expiry:    expiresAt,
	}

	if err := s.Repo.CreateSession(newSession); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := s.Repo.DeleteSessionBySessionID(sessionID); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionID,
		Expires: time.Now().Add(120 * time.Second),
	})
}

func (s *Session) LogOut(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sessionID := c.Value

	if err := s.Repo.DeleteSessionBySessionID(sessionID); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
