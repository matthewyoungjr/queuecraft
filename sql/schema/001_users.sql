-- +goose up
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    full_name VARCHAR(70) NOT NULL,
    email VARCHAR(40) NOT NULL,
    password VARCHAR(60) NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);


-- +goose down 
DROP TABLE IF EXISTS users;