CREATE OR REPLACE FUNCTION create_weather (latitude decimal, --$1
longitude decimal --$2
)
    RETURNS bigint
    AS $$
BEGIN
    INSERT INTO weather (longitude, latitude)
        VALUES ($1, $2)
    RETURNING id;
END
$$
LANGUAGE plpgsql;
