package utils

import "database/sql"

type Repo struct {
	Client *sql.DB
}
