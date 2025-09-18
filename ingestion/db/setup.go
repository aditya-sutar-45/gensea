package db

import (
	"context"
	"database/sql"
)

const rawOceanDataQuery = `
CREATE TABLE IF NOT EXISTS raw_ocean_data (
    id SERIAL PRIMARY KEY,
    source_observation_id TEXT,
    date VARCHAR(50),
    station_id TEXT,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    depth_m DOUBLE PRECISION,
    temperature_surface_c DOUBLE PRECISION,
    temperature_100m_c DOUBLE PRECISION,
    salinity_psu DOUBLE PRECISION,
    dissolved_oxygen_mg_l DOUBLE PRECISION,
    ph DOUBLE PRECISION,
    chlorophyll_a_mg_m3 DOUBLE PRECISION,
    wave_height_m DOUBLE PRECISION,
    current_speed_m_s DOUBLE PRECISION,
    region TEXT,
    data_quality TEXT,
    UNIQUE (source_observation_id)
)
`

func EnsureTables(ctx context.Context, db *sql.DB) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, rawOceanDataQuery); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
