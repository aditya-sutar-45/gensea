package db

import (
	"context"
	"database/sql"

	"github.com/aditya-sutar-45/gensea/ingestion/models"
)

func ImportRecords[T models.DBModel](ctx context.Context, db *sql.DB, records []T) error {

	for i := range records {
		if err := records[i].Insert(ctx, db); err != nil {
			return err
		}
	}
	return nil
}
