import pandas as pd
import numpy as np
from sqlalchemy import create_engine
from sklearn.impute import KNNImputer
from pyod.models.iforest import IForest
from fuzzywuzzy import process, fuzz
from datetime import datetime, timezone
from dateutil import parser

# Database connection params
db_config = {
    'user': 'postgres',
    'password': 'Sushant',
    'host': 'localhost',
    'port': '5432',
    'database': 'GenSea'
}

# Create SQLAlchemy engine and connect
conn_string = (
    f"postgresql://{db_config['user']}:{db_config['password']}"
    f"@{db_config['host']}:{db_config['port']}/{db_config['database']}"
)
engine = create_engine(conn_string)
conn = engine.connect()

# Load raw tables into DataFrames
ocean_df = pd.read_sql("SELECT * FROM raw_ocean_data", conn)
fish_df  = pd.read_sql("SELECT * FROM raw_fisheries_data", conn)

# ISO date format
def to_iso8601(x):
    try:
        dt = parser.parse(str(x), dayfirst=True)
        return dt.astimezone(timezone.utc).isoformat()
    except Exception:
        return str(x)

for df in (ocean_df, fish_df):
    df['date'] = df['date'].apply(to_iso8601)

# Automated imputation preprocessing for empty strings and enforce numeric dtype

num_cols_ocean = [
    'temperature_surface_c', 'temperature_100m_c', 'salinity_psu',
    'dissolved_oxygen_mg_l', 'ph', 'chlorophyll_a_mg_m3',
    'wave_height_m', 'current_speed_m_s'
]
num_cols_fish = [
    'catch_weight_kg', 'market_price_per_kg', 'depth_fished_m',
    'latitude', 'longitude', 'effort_hours', 'bycatch_kg'
]

# Replace empty strings in numeric columns by np.nan and ensure float dtype
ocean_df[num_cols_ocean] = ocean_df[num_cols_ocean].replace(r'^\s*$', np.nan, regex=True).astype(float)
fish_df[num_cols_fish] = fish_df[num_cols_fish].replace(r'^\s*$', np.nan, regex=True).astype(float)

imputer_ocean = KNNImputer(n_neighbors=5)
ocean_df[num_cols_ocean] = imputer_ocean.fit_transform(ocean_df[num_cols_ocean])

imputer_fish = KNNImputer(n_neighbors=5)
fish_df[num_cols_fish] = imputer_fish.fit_transform(fish_df[num_cols_fish])

# ANOMALY DETECTION using Isolation Forest
clf_ocean = IForest(contamination=0.01)
clf_ocean.fit(ocean_df[num_cols_ocean])
ocean_df['anomaly'] = clf_ocean.predict(ocean_df[num_cols_ocean])
ocean_df = ocean_df[ocean_df['anomaly'] == 0].drop(columns=['anomaly'])

clf_fish = IForest(contamination=0.01)
clf_fish.fit(fish_df[num_cols_fish])
fish_df['anomaly'] = clf_fish.predict(fish_df[num_cols_fish])
fish_df = fish_df[fish_df['anomaly'] == 0].drop(columns=['anomaly'])

# TEXT STANDARDIZATION & NORMALIZATION
for col in ['region', 'data_quality']:
    ocean_df[col] = ocean_df[col].astype(str).str.lower().str.strip()

for col in [
    'port', 'species_common', 'species_scientific',
    'fishing_method', 'weather_condition'
]:
    fish_df[col] = fish_df[col].astype(str).str.lower().str.strip()

# Fuzzy normalize scientific names against known list
known_species = [
    "thunnus albacares", "scomberomorus commerson",
    "lutjanus malabaricus"
]
def normalize_name(name):
    best, score = process.extractOne(name, known_species, scorer=fuzz.token_sort_ratio)
    return best if score >= 80 else name

fish_df['species_scientific'] = fish_df['species_scientific'].apply(normalize_name)

# ADD METADATA COLUMNS
for df in (ocean_df, fish_df):
    df['ingested_at'] = datetime.now(timezone.utc)
    df['pipeline_version'] = 'v1.0_ai_clean'

ocean_df.to_sql(
    'cleaned_ocean_data',
    engine,
    schema='public',
    if_exists='replace',
    index=False,
    method='multi'
)

fish_df.to_sql(
    'cleaned_fisheries_data',
    engine,
    schema='public',
    if_exists='replace',
    index=False,
    method='multi'
)

engine.dispose()
conn.close()
