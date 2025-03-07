package session

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *SessionRepo) CreateSession(session *model.Session) error {
	_, err := r.Client.Exec(
		"INSERT INTO sessions (session_id, user_id, expiry) VALUES ($1, $2, $3)",
		session.SessionID, session.UserID, session.Expiry,
	)
	if err != nil {
		return fmt.Errorf("CreateSession: %v", err)
	}

	return nil
}
