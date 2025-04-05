package product

import (
	"database/sql"
)

type Repo struct {
	Client *sql.DB
}

const limit = 4

func getOffset(page int) int {
	return (page - 1) * limit
}
