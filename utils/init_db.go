package utils

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

// InitDB initializes a PostgreSQL connection using pgx
func InitDB(dsn string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Println("dsn: ", dsn)
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	return conn
}
