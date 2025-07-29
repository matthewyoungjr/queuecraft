-- name: CreateJobLog :one
INSERT INTO job_logs(
    id, job_id, user_id, message, attempts, 
    duration_ms, created_at
)
VALUES(
    DEFAULT, 
    $1, -- job_id 
    $2, -- user id 
    $3, -- message 
    $4, -- attempts
    $5, -- duration_ms 
    $6 -- created_at 
)
RETURNING *;


-- name: ListAllLogs :many
SELECT * 
FROM job_logs
ORDER BY id;


-- name: ListLogsPaginated :many 
SELECT *
FROM job_logs
LIMIT $1 OFFSET $2;