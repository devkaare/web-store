package utils

import "log"

func (r *Repo) Close() error {
	log.Println("Disconnected from database")
	return r.Client.Close()
}
