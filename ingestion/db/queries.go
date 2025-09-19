package db

const rawOceanDataQuery = `
CREATE TABLE IF NOT EXISTS raw_ocean_data (
    id SERIAL PRIMARY KEY,
    source_observation_id TEXT,
    data_source TEXT,
    date VARCHAR(50),
    station_id TEXT,
    latitude TEXT,
    longitude TEXT,
    depth_m TEXT,
    temperature_surface_c TEXT,
    temperature_100m_c TEXT,
    salinity_psu TEXT,
    dissolved_oxygen_mg_l TEXT,
    ph  TEXT,
    chlorophyll_a_mg_m3 TEXT,
    wave_height_m TEXT,
    current_speed_m_s TEXT,
    region TEXT,
    data_quality TEXT,
    UNIQUE (source_observation_id)
)
`

const rawFisheriesDataQuery = `
CREATE TABLE IF NOT EXISTS raw_fisheries_data (
    id SERIAL PRIMARY KEY,
    catch_id TEXT UNIQUE,
    data_source TEXT,
    date VARCHAR(50),
    vessel_id TEXT,
    port TEXT,
    species_common TEXT,
    species_scientific TEXT,
    catch_weight_kg DOUBLE PRECISION,
    market_price_per_kg TEXT,
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
