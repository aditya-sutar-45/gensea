package db

import (
	"context"
	"database/sql"
)

func EnsureTables(ctx context.Context, db *sql.DB) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, rawOceanDataQuery); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.ExecContext(ctx, rawFisheriesDataQuery); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
