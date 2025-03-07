package session

import (
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *SessionRepo) GetAllSessions() ([]model.Session, error) {
	var sessions []model.Session

	rows, err := r.Client.Query("SELECT * FROM sessions")
	if err != nil {
		return sessions, err
	}
	defer rows.Close()

	for rows.Next() {
		var session model.Session
		if err := rows.Scan(&session.SessionID, &session.UserID, &session.Expiry); err != nil {
			return sessions, fmt.Errorf("GetAllSessions %s: %v", session.SessionID, err)
		}
		sessions = append(sessions, session)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllSessions %v:", err)
	}

	return sessions, nil
}
