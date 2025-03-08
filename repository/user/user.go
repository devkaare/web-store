package user

import "database/sql"

type Repo struct {
	Client *sql.DB
}

func GetUser(userRepoGetter func() *sql.DB) *Repo {
	userRepo := userRepoGetter()
	return &Repo{Client: userRepo}
}
