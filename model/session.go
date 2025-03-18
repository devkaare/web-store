package model

import "time"

type Session struct {
	SessionID string
	UserID    int
	Expiry    time.Time
}
