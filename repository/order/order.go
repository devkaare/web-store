package order

import (
	"database/sql"
)

type Repo struct {
	Client *sql.DB
}
