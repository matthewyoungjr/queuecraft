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

-- name: DeleteJobById :exec
DELETE FROM jobs 
WHERE id = $1;