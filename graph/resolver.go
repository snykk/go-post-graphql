package graph

import (
	"github.com/jackc/pgx/v5"
)

type Resolver struct {
	DB *pgx.Conn
}
