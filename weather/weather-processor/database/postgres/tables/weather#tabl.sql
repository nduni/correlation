CREATE TABLE weather (
    id bigserial PRIMARY KEY,
    created_at timestamp with time zone DEFAULT NOW() NOT NULL,
    updated_at timestamp with time zone DEFAULT NOW() NOT NULL,
    longitutde decimal(7,4),
    latitude decimal(7,4),
    UNIQUE (longitutde, latitude)
)