package model

import "time"

type Session struct {
	SessionID string    `json:"session_id"`
	UserID    uint32    `json:"user_id"`
	Expiry    time.Time `json:"expiry"`
}
