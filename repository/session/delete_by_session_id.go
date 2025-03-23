package session

import (
	"fmt"
)

func (r *Repo) DeleteSessionBySessionID(sessionID string) error {
	result, err := r.Client.Exec("DELETE FROM sessions WHERE session_id = $1", sessionID)
	if err != nil {
		return fmt.Errorf("DeleteSessionBySessionID %s: %v", sessionID, err)
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
