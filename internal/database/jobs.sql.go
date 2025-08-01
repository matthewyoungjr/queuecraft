// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: jobs.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createJob = `-- name: CreateJob :one
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

RETURNING id, type, payload, status, run_at, created_at
`

type CreateJobParams struct {
	Column1 pgtype.Int4
	Column2 []byte
	Column3 pgtype.Text
	Column4 pgtype.Int2
	Column5 pgtype.Int2
	Column6 pgtype.Int4
	Column7 pgtype.Timestamp
	Column8 pgtype.Timestamp
}

type CreateJobRow struct {
	ID        pgtype.UUID
	Type      pgtype.Int4
	Payload   []byte
	Status    pgtype.Text
	RunAt     pgtype.Timestamp
	CreatedAt pgtype.Timestamp
}

func (q *Queries) CreateJob(ctx context.Context, arg CreateJobParams) (CreateJobRow, error) {
	row := q.db.QueryRow(ctx, createJob,
		arg.Column1,
		arg.Column2,
		arg.Column3,
		arg.Column4,
		arg.Column5,
		arg.Column6,
		arg.Column7,
		arg.Column8,
	)
	var i CreateJobRow
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Payload,
		&i.Status,
		&i.RunAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteJobById = `-- name: DeleteJobById :exec
DELETE FROM jobs 
WHERE id = $1
`

func (q *Queries) DeleteJobById(ctx context.Context, dollar_1 pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteJobById, dollar_1)
	return err
}

const getJobById = `-- name: GetJobById :one
SELECT id, type, payload, status, retries, run_at, created_at 
FROM jobs
WHERE id = $1
`

type GetJobByIdRow struct {
	ID        pgtype.UUID
	Type      pgtype.Int4
	Payload   []byte
	Status    pgtype.Text
	Retries   pgtype.Int2
	RunAt     pgtype.Timestamp
	CreatedAt pgtype.Timestamp
}

func (q *Queries) GetJobById(ctx context.Context, dollar_1 pgtype.UUID) (GetJobByIdRow, error) {
	row := q.db.QueryRow(ctx, getJobById, dollar_1)
	var i GetJobByIdRow
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Payload,
		&i.Status,
		&i.Retries,
		&i.RunAt,
		&i.CreatedAt,
	)
	return i, err
}

const listJobs = `-- name: ListJobs :many
SELECT id, type, payload, status, run_at, created_at
FROM jobs 
ORDER BY created_at DESC
`

type ListJobsRow struct {
	ID        pgtype.UUID
	Type      pgtype.Int4
	Payload   []byte
	Status    pgtype.Text
	RunAt     pgtype.Timestamp
	CreatedAt pgtype.Timestamp
}

func (q *Queries) ListJobs(ctx context.Context) ([]ListJobsRow, error) {
	rows, err := q.db.Query(ctx, listJobs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListJobsRow
	for rows.Next() {
		var i ListJobsRow
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Payload,
			&i.Status,
			&i.RunAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listJobsPaginated = `-- name: ListJobsPaginated :many
SELECT id, type, payload, status, run_at, created_at
FROM jobs 
ORDER BY created_at DESC
LIMIT $1 OFFSET $2
`

type ListJobsPaginatedParams struct {
	Column1 pgtype.Int8
	Column2 pgtype.Int8
}

type ListJobsPaginatedRow struct {
	ID        pgtype.UUID
	Type      pgtype.Int4
	Payload   []byte
	Status    pgtype.Text
	RunAt     pgtype.Timestamp
	CreatedAt pgtype.Timestamp
}

func (q *Queries) ListJobsPaginated(ctx context.Context, arg ListJobsPaginatedParams) ([]ListJobsPaginatedRow, error) {
	rows, err := q.db.Query(ctx, listJobsPaginated, arg.Column1, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListJobsPaginatedRow
	for rows.Next() {
		var i ListJobsPaginatedRow
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Payload,
			&i.Status,
			&i.RunAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateJobDuration = `-- name: UpdateJobDuration :exec
UPDATE jobs
SET duration_ms = $2
WHERE id = $1
`

type UpdateJobDurationParams struct {
	Column1 pgtype.UUID
	Column2 pgtype.Int4
}

func (q *Queries) UpdateJobDuration(ctx context.Context, arg UpdateJobDurationParams) error {
	_, err := q.db.Exec(ctx, updateJobDuration, arg.Column1, arg.Column2)
	return err
}

const updateJobStatus = `-- name: UpdateJobStatus :exec
UPDATE jobs
SET status = $1
WHERE id = $2
`

type UpdateJobStatusParams struct {
	Column1 pgtype.Text
	Column2 pgtype.UUID
}

func (q *Queries) UpdateJobStatus(ctx context.Context, arg UpdateJobStatusParams) error {
	_, err := q.db.Exec(ctx, updateJobStatus, arg.Column1, arg.Column2)
	return err
}
