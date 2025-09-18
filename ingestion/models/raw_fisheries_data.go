package models

import (
	"context"
	"database/sql"
	"strings"
)

type RawFisheriesData struct {
	ID                int
	CatchID           string  `csv:"catch_id"`
	Date              string  `csv:"date"`
	VesselID          string  `csv:"vessel_id"`
	Port              string  `csv:"port"`
	SpeciesCommon     string  `csv:"species_common"`
	SpeciesScientific string  `csv:"species_scientific"`
	CatchWeightKg     float64 `csv:"catch_weight_kg"`
	MarketPricePerKg  float64 `csv:"market_price_per_kg"`
	FishingMethod     string  `csv:"fishing_method"`
	DepthFishedM      float64 `csv:"depth_fished_m"`
	Latitude          float64 `csv:"latitude"`
	Longitude         float64 `csv:"longitude"`
	EffortHours       float64 `csv:"effort_hours"`
	CrewSize          int     `csv:"crew_size"`
	WeatherCondition  string  `csv:"weather_condition"`
	BycatchKg         float64 `csv:"bycatch_kg"`
}

const insertFisheriesQuery = `
INSERT INTO raw_fisheries_data (
    catch_id, date, vessel_id, port, species_common, species_scientific,
    catch_weight_kg, market_price_per_kg, fishing_method, depth_fished_m,
    latitude, longitude, effort_hours, crew_size, weather_condition, bycatch_kg
)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)
ON CONFLICT (catch_id) DO NOTHING
RETURNING id;
`

func (r *RawFisheriesData) Insert(ctx context.Context, db *sql.DB) error {
	r.Date = strings.TrimSpace(r.Date)

	err := db.QueryRowContext(
		ctx, insertFisheriesQuery,
		r.CatchID, r.Date, r.VesselID, r.Port, r.SpeciesCommon, r.SpeciesScientific,
		r.CatchWeightKg, r.MarketPricePerKg, r.FishingMethod, r.DepthFishedM,
		r.Latitude, r.Longitude, r.EffortHours, r.CrewSize, r.WeatherCondition, r.BycatchKg,
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
