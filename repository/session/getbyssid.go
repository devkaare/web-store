package session

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *Repo) GetSessionBySessionID(sessionID string) (*model.Session, error) {
	session := &model.Session{}

	row := r.Client.QueryRow("SELECT * FROM sessions WHERE session_id = $1", sessionID)
	if err := row.Scan(&session.SessionID, &session.UserID, &session.Expiry); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("GetSessionBySessionID %s: %v", sessionID, err)
	}

	return session, nil
}
