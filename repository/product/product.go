package product

import "database/sql"

type ProductRepo struct {
	Client *sql.DB
}
