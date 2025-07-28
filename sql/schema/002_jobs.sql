-- +goose up
CREATE TABLE IF NOT EXISTS jobs(
    id UUID PRIMARY KEY, 
    type INTEGER REFERENCES job_types(id), 
    payload JSONB NOT NULL,
    status VARCHAR(100), 
    retries SMALLINT DEFAULT NULL, 
    max_retries SMALLINT DEFAULT NULL,
    duration_ms INTEGER DEFAULT NULL,   
    run_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose down
DROP TABLE IF EXISTS jobs;