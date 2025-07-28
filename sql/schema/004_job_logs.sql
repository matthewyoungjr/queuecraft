-- +goose up
CREATE TABLE IF NOT EXISTS job_logs(
    id SERIAL PRIMARY KEY, 
    job_id UUID REFERENCES jobs(id),
    user_id UUID REFERENCES users(id), 
    message TEXT DEFAULT NULL, 
    attempts SMALLINT DEFAULT NULL, 
    duration_ms INTEGER DEFAULT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);


-- +goose down 
DROP TABLE IF EXISTS job_logs;