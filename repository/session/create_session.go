package session

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) CreateSession(session *model.Session) error {
	_, err := r.Client.Exec(
		"INSERT INTO sessions (session_id, expiry) VALUES ($1, $2)",
		session.SessionID, session.Expiry,
	)
	if err != nil {
		return fmt.Errorf("CreateSession: %v", err)
	}

	return nil
}
