package db

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

const rawFisheriesDataQuery = `
CREATE TABLE IF NOT EXISTS raw_fisheries_data (
    id SERIAL PRIMARY KEY,
    catch_id TEXT UNIQUE,
    date VARCHAR(50),
    vessel_id TEXT,
    port TEXT,
    species_common TEXT,
    species_scientific TEXT,
    catch_weight_kg DOUBLE PRECISION,
    market_price_per_kg DOUBLE PRECISION,
    fishing_method TEXT,
    depth_fished_m DOUBLE PRECISION,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    effort_hours DOUBLE PRECISION,
    crew_size INT,
    weather_condition TEXT,
    bycatch_kg DOUBLE PRECISION
)
`
