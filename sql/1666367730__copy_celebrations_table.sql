COPY celebrations(id, date, name, status, type) FROM '/docker-entrypoint-initdb.d/celebrations.csv' WITH DELIMITER ',' CSV HEADER;
