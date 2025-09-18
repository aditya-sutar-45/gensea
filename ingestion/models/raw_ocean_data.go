package models

import (
	"context"
	"database/sql"
	"strings"
)

type RawOceanData struct {
	ID                  int
	SourceObservationID string  `csv:"observation_id"`
	Date                string  `csv:"date"`
	StationID           string  `csv:"station_id"`
	Latitude            float64 `csv:"latitude"`
	Longitude           float64 `csv:"longitude"`
	DepthM              float64 `csv:"depth_m"`
	TemperatureSurfaceC float64 `csv:"temperature_surface_C"`
	Temperature100mC    float64 `csv:"temperature_100m_C"`
	SalinityPSU         float64 `csv:"salinity_psu"`
	DissolvedOxygenMgL  float64 `csv:"dissolved_oxygen_mg_l"`
	PH                  float64 `csv:"pH"`
	ChlorophyllAMgM3    float64 `csv:"chlorophyll_a_mg_m3"`
	WaveHeightM         float64 `csv:"wave_height_m"`
	CurrentSpeedMS      float64 `csv:"current_speed_m_s"`
	Region              string  `csv:"region"`
	DataQuality         string  `csv:"data_quality"`
}

const insertQuery = `
INSERT INTO raw_ocean_data (
    source_observation_id, date, station_id, latitude, longitude, depth_m,
    temperature_surface_c, temperature_100m_c, salinity_psu,
    dissolved_oxygen_mg_l, ph, chlorophyll_a_mg_m3, wave_height_m,
    current_speed_m_s, region, data_quality
)
VALUES ($1,$2,$3,$4,$5,$6,
        $7,$8,$9,$10,$11,$12,
        $13,$14,$15,$16)
ON CONFLICT (source_observation_id) DO NOTHING
RETURNING id;
`

func (r *RawOceanData) Insert(ctx context.Context, db *sql.DB) error {
	date := strings.Trim(r.Date, " ")
	err := db.QueryRowContext(
		ctx, insertQuery,
		r.SourceObservationID, date, r.StationID, r.Latitude, r.Longitude,
		r.DepthM, r.TemperatureSurfaceC, r.Temperature100mC, r.SalinityPSU,
		r.DissolvedOxygenMgL, r.PH, r.ChlorophyllAMgM3, r.WaveHeightM,
		r.CurrentSpeedMS, r.Region, r.DataQuality,
	).Scan(&r.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			// conflict happened, row was not inserted
			return nil
		}
		return err
	}

	return nil
}
