CREATE TABLE measurment (
    id bigserial PRIMARY KEY,
    created_at timestamp with time zone DEFAULT NOW() NOT NULL,
    updated_at timestamp with time zone DEFAULT NOW() NOT NULL,
    measurement_time timestamp with time zone DEFAULT NOW() NOT NULL,
    temperature_2m decimal(10,2),
    showers decimal(10,2),
    surface_pressure decimal(10,2),
    windspeed_10m decimal(10,2),
    winddirection_10m decimal (5,2),
    weather_id bigint REFERENCES weather (id),
    UNIQUE (measurement_time)
)