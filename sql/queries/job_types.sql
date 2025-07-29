-- name: CreateJobType :many
INSERT INTO job_types(name)
VALUES
('send_email'), 
('send_push_notification'), 
('send_sms')
RETURNING id, name;  


-- name: ListJobTypes :many
SELECT id, name 
FROM job_types
ORDER BY name;


-- name: GetJobTypeById :one
SELECT id, name 
FROM job_types
WHERE id = $1;