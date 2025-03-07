package session

import "database/sql"

type SessionRepo struct {
	Client *sql.DB
}
