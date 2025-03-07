package user

import "database/sql"

type UserRepo struct {
	Client *sql.DB
}
