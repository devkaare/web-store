package utils

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type UtilsRepo struct {
	Client *sql.DB
}

func (r *UtilsRepo) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	err := r.Client.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err)
		return stats
	}

	stats["status"] = "up"
	stats["message"] = "It's healthy"

	return stats
}

func (r *UtilsRepo) Close() error {
	log.Println("Disconnected from database")
	return r.Client.Close()
}
