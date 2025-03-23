package session

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) CreateSession(session *model.Session) error {
	_, err := r.Client.Exec(
		// "INSERT INTO sessions (session_id, user_id) VALUES ($1, $2)",
		// session.SessionID, session.UserID,
		"INSERT INTO sessions (session_id) VALUES ($1)",
		session.SessionID,
	)
	if err != nil {
		return fmt.Errorf("CreateSession: %v", err)
	}

	return nil
}
