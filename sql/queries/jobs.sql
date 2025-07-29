-- name: CreateJob :one
INSERT INTO jobs(
    id, type, payload, status, retries, 
    max_retries, duration_ms, run_at, created_at
)

VALUES(
    gen_random_uuid(),
    $1, -- type
    $2, -- payload
    $3, -- status
    $4, -- retries
    $5, -- max_retries
    $6, -- duration_ms
    $7, -- run_at 
    $8  -- created_at
)

RETURNING id, type, payload, status, run_at, created_at;

-- name: GetJobById :one 
SELECT id, type, payload, status, retries, run_at, created_at 
FROM jobs
RETURNING id, type, payload, status, retries, run_at, created_at;


-- name: DeleteJobById :exec
DELETE FROM jobs 
WHERE id = $1;


-- name: ListJobs :many
SELECT id, type, payload, status, run_at, created_at
FROM jobs 
ORDER BY created_at DESC;

-- name: ListJobsPaginated :many
SELECT id, type, payload, status, run_at, created_at
FROM jobs 
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;


-- name: UpdateJobStatus :exec
UPDATE jobs
SET status = $1
WHERE id = $2;


-- name: UpdateJobDuration :exec
UPDATE jobs
SET duration_ms = $2
WHERE id = $1;
