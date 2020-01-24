CREATE TABLE scooters (
	id SERIAL PRIMARY KEY,
    lat FLOAT8 NOT NULL,
    lng FLOAT8 NOT NULL,
    is_reserved BOOLEAN NOT NULL DEFAULT FALSE
);
CREATE INDEX ON scooters USING gist (ll_to_earth(lat, lng));
INSERT INTO scooters (id, lat, lng) VALUES (10, 37.788548, -122.411548), (8, 37.783223, -122.398630);
