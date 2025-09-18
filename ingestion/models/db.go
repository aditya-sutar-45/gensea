package models

import (
	"context"
	"database/sql"
)

type DBModel interface {
	Insert(ctx context.Context, db *sql.DB) error
}
