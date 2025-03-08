package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/devkaare/web-store/model"
	"github.com/devkaare/web-store/repository"
	"github.com/devkaare/web-store/repository/session"
	"github.com/google/uuid"
)

type Session struct {
	Repo *session.Repo
}

var sessionHandler = &Session{
	Repo: &session.Repo{},
}

func NewSessionHandler(db *sql.DB) *Session {
	sessionHandler.Repo = repository.GetSession(func() *sql.DB {
		return db
	})
	return sessionHandler
}

func isExpired(s *model.Session) bool {
	return s.Expiry.Before(time.Now())
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

func (s *Session) GetAllSessions(w http.ResponseWriter, r *http.Request) {
	sessions, err := s.Repo.GetAllSessions()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(sessions)
	_, _ = w.Write(jsonResp)
}
