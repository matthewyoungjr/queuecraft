-- +goose up
CREATE TABLE IF NOT EXISTS job_types(
    id SERIAL PRIMARY KEY,
    name VARCHAR(25)
);


-- +goose down
DROP TABLE IF EXISTS job_types