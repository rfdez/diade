COPY celebrations(id, date, name, status, type) FROM '/docker-entrypoint-initdb.d/celebrations-2022.csv' WITH DELIMITER ',' CSV HEADER;
