package model

import "time"

type Session struct {
	SessionID string
	UserID    uint32
	Expiry    time.Time
}
