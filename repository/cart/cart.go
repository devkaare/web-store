package cart

import "database/sql"

type CartRepo struct {
	Client *sql.DB
}
