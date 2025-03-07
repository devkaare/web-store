package session

import "database/sql"

type PostgresRepo struct {
	Client *sql.DB
}
