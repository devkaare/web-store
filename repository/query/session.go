package query

import (
	"database/sql"
	"fmt"

	"github.com/devkaare/web-store/model"
)

func (r *PostgresRepo) CreateSession(session *model.Session) error {
	_, err := r.Client.Exec(
		"INSERT INTO sessions (session_id, user_id, expiry) VALUES ($1, $2, $3)",
		session.SessionID, session.UserID, session.Expiry,
	)
	if err != nil {
		return fmt.Errorf("CreateSession: %v", err)
	}

	return nil
}

func (r *PostgresRepo) GetSessions() ([]model.Session, error) {
	var sessions []model.Session

	rows, err := r.Client.Query("SELECT * FROM sessions")
	if err != nil {
		return sessions, err
	}
	defer rows.Close()

	for rows.Next() {
		var session model.Session
		if err := rows.Scan(&session.SessionID, &session.UserID, &session.Expiry); err != nil {
			return sessions, fmt.Errorf("GetSessions %s: %v", session.SessionID, err)
		}
		sessions = append(sessions, session)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetSessions %v:", err)
	}

	return sessions, nil
}

func (r *PostgresRepo) GetSessionBySessionID(sessionID string) (*model.Session, error) {
	session := &model.Session{}

	row := r.Client.QueryRow("SELECT * FROM sessions WHERE session_id = $1", sessionID)
	if err := row.Scan(&session.SessionID, &session.UserID, &session.Expiry); err != nil {
		if err == sql.ErrNoRows {
			return session, nil
		}
		return session, fmt.Errorf("GetSessionBySessionID %s: %v", sessionID, err)
	}

	return session, nil
}

func (r *PostgresRepo) DeleteSessionBySessionID(sessionID string) error {
	result, err := r.Client.Exec("DELETE FROM sessions WHERE session_id = $1", sessionID)
	if err != nil {
		return fmt.Errorf("DeleteSessionBySessionID %s, %v", sessionID, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteSessionBySessionID %s: %v", sessionID, err)
	}
	if count < 1 {
		return fmt.Errorf("DeleteSessionBySessionID %s: session not found", sessionID)
	}

	return nil
}
