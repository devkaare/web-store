package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

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

	_, err = s.Repo.GetSessionBySessionID(sessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User is authorized"))
}

func (s *Session) Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sessionID := cookie.Value

	session, err := s.Repo.GetSessionBySessionID(sessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newSessionID := uuid.NewString()

	newSession := &model.Session{
		SessionID: newSessionID,
		UserID:    session.UserID,
	}

	err = s.Repo.CreateSession(newSession)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.Repo.DeleteSessionBySessionID(sessionID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session_token",
		Value: newSessionID,
	})
}

func (s *Session) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sessionID := cookie.Value

	err = s.Repo.DeleteSessionBySessionID(sessionID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Session) GetAllSessions(w http.ResponseWriter, r *http.Request) {
	sessions, err := s.Repo.GetAllSessions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(sessions)
	_, _ = w.Write(jsonResp)
}
