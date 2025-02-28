package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/devkaare/web-store/hash"
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

func (s *Session) SignUp(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	_, err := s.Repo.GetUserByEmail(email)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &model.User{
		Email:    email,
		Password: password,
	}

	if _, err := s.Repo.CreateUser(user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Session) SignIn(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	expectedPassword := r.FormValue("password")

	existingUser, err := s.Repo.GetUserByEmail(email)
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

	session, err := s.Repo.GetSessionBySessionID(sessionID)
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
		if err := s.Repo.DeleteSessionBySessionID(sessionID); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// w.Write([]byte("User is authorized"))
	http.Redirect(w, r, "/listings", http.StatusSeeOther)
}

func (s *Session) Refresh(w http.ResponseWriter, r *http.Request) {
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

	session, err := s.Repo.GetSessionBySessionID(sessionID)
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

	if err := s.Repo.DeleteSessionBySessionID(sessionID); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Session) GetSessions(w http.ResponseWriter, r *http.Request) {
	sessions, err := s.Repo.GetSessions()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(sessions)
	_, _ = w.Write(jsonResp)
}
