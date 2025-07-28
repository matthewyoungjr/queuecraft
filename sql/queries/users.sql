-- name: CreateUser :one
INSERT INTO users(id, full_name, email, password, created_at)
VALUES(
    gen_random_uuid(), 
    $1, -- full_name
    $2, -- email
    $3, -- password
    NOW()
)

RETURNING id, full_name, email, created_at;

-- name: GetUserById :one
SELECT id, full_name, email, created_at
FROM users
WHERE id = $1;


-- name: GetUserByEmail :one 
SELECT id, full_name, email, password
FROM users 
WHERE email = $1;


-- name: DeleteUserById :one
DELETE FROM users 
WHERE id = $1
RETURNING id, full_name, email, created_at;