// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: driver.sql

package db

import (
	"context"
	"database/sql"
)

const createDriver = `-- name: CreateDriver :one
INSERT INTO driver (
  phone,
  name
) VALUES (
  $1, $2
)
RETURNING id, phone, name, date_of_birth, avatar_url, created_at, status, latitude, longitude
`

type CreateDriverParams struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

func (q *Queries) CreateDriver(ctx context.Context, arg CreateDriverParams) (Driver, error) {
	row := q.db.QueryRowContext(ctx, createDriver, arg.Phone, arg.Name)
	var i Driver
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Name,
		&i.DateOfBirth,
		&i.AvatarUrl,
		&i.CreatedAt,
		&i.Status,
		&i.Latitude,
		&i.Longitude,
	)
	return i, err
}

const deleteDriver = `-- name: DeleteDriver :exec
DELETE FROM driver
WHERE phone = $1
`

func (q *Queries) DeleteDriver(ctx context.Context, phone string) error {
	_, err := q.db.ExecContext(ctx, deleteDriver, phone)
	return err
}

const getDriver = `-- name: GetDriver :one
SELECT id, phone, name, date_of_birth, avatar_url, created_at, status, latitude, longitude FROM driver
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetDriver(ctx context.Context, id int64) (Driver, error) {
	row := q.db.QueryRowContext(ctx, getDriver, id)
	var i Driver
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Name,
		&i.DateOfBirth,
		&i.AvatarUrl,
		&i.CreatedAt,
		&i.Status,
		&i.Latitude,
		&i.Longitude,
	)
	return i, err
}

const getDriverByPhone = `-- name: GetDriverByPhone :one
SELECT id, phone, name, date_of_birth, avatar_url, created_at, status, latitude, longitude FROM driver
WHERE phone = $1 LIMIT 1
`

func (q *Queries) GetDriverByPhone(ctx context.Context, phone string) (Driver, error) {
	row := q.db.QueryRowContext(ctx, getDriverByPhone, phone)
	var i Driver
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Name,
		&i.DateOfBirth,
		&i.AvatarUrl,
		&i.CreatedAt,
		&i.Status,
		&i.Latitude,
		&i.Longitude,
	)
	return i, err
}

const getDriverForUpdate = `-- name: GetDriverForUpdate :one
SELECT id, phone, name, date_of_birth, avatar_url, created_at, status, latitude, longitude FROM driver
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetDriverForUpdate(ctx context.Context, id int64) (Driver, error) {
	row := q.db.QueryRowContext(ctx, getDriverForUpdate, id)
	var i Driver
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Name,
		&i.DateOfBirth,
		&i.AvatarUrl,
		&i.CreatedAt,
		&i.Status,
		&i.Latitude,
		&i.Longitude,
	)
	return i, err
}

const getDriverNearby = `-- name: GetDriverNearby :many
SELECT phone, latitude, longitude, SQRT(
    POW(69.1 * (latitude - $2::float8), 2) +
    POW(69.1 * ($3::float8 - longitude) * COS(latitude / 57.3), 2))::float8 AS distance
FROM driver 
WHERE status = 'finding'
AND latitude IS NOT NULL
AND longitude IS NOT NULL
ORDER BY distance LIMIT $1
`

type GetDriverNearbyParams struct {
	Limit     int32   `json:"limit"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GetDriverNearbyRow struct {
	Phone     string          `json:"phone"`
	Latitude  sql.NullFloat64 `json:"latitude"`
	Longitude sql.NullFloat64 `json:"longitude"`
	Distance  float64         `json:"distance"`
}

func (q *Queries) GetDriverNearby(ctx context.Context, arg GetDriverNearbyParams) ([]GetDriverNearbyRow, error) {
	rows, err := q.db.QueryContext(ctx, getDriverNearby, arg.Limit, arg.Latitude, arg.Longitude)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetDriverNearbyRow
	for rows.Next() {
		var i GetDriverNearbyRow
		if err := rows.Scan(
			&i.Phone,
			&i.Latitude,
			&i.Longitude,
			&i.Distance,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listDrivers = `-- name: ListDrivers :many
SELECT id, phone, name, date_of_birth, avatar_url, created_at, status, latitude, longitude FROM driver
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListDriversParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListDrivers(ctx context.Context, arg ListDriversParams) ([]Driver, error) {
	rows, err := q.db.QueryContext(ctx, listDrivers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Driver
	for rows.Next() {
		var i Driver
		if err := rows.Scan(
			&i.ID,
			&i.Phone,
			&i.Name,
			&i.DateOfBirth,
			&i.AvatarUrl,
			&i.CreatedAt,
			&i.Status,
			&i.Latitude,
			&i.Longitude,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDriver = `-- name: UpdateDriver :one

UPDATE driver
SET name = $2,
  date_of_birth = $3
WHERE phone = $1
RETURNING id, phone, name, date_of_birth, avatar_url, created_at, status, latitude, longitude
`

type UpdateDriverParams struct {
	Phone       string       `json:"phone"`
	Name        string       `json:"name"`
	DateOfBirth sql.NullTime `json:"date_of_birth"`
}

// pagination: offset: skip many rows
func (q *Queries) UpdateDriver(ctx context.Context, arg UpdateDriverParams) (Driver, error) {
	row := q.db.QueryRowContext(ctx, updateDriver, arg.Phone, arg.Name, arg.DateOfBirth)
	var i Driver
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Name,
		&i.DateOfBirth,
		&i.AvatarUrl,
		&i.CreatedAt,
		&i.Status,
		&i.Latitude,
		&i.Longitude,
	)
	return i, err
}

const updateLocation = `-- name: UpdateLocation :one
UPDATE driver
SET latitude = $2,
  longitude = $3
WHERE phone = $1
RETURNING id, phone, name, date_of_birth, avatar_url, created_at, status, latitude, longitude
`

type UpdateLocationParams struct {
	Phone     string          `json:"phone"`
	Latitude  sql.NullFloat64 `json:"latitude"`
	Longitude sql.NullFloat64 `json:"longitude"`
}

func (q *Queries) UpdateLocation(ctx context.Context, arg UpdateLocationParams) (Driver, error) {
	row := q.db.QueryRowContext(ctx, updateLocation, arg.Phone, arg.Latitude, arg.Longitude)
	var i Driver
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Name,
		&i.DateOfBirth,
		&i.AvatarUrl,
		&i.CreatedAt,
		&i.Status,
		&i.Latitude,
		&i.Longitude,
	)
	return i, err
}

const updateStatus = `-- name: UpdateStatus :one
UPDATE driver
SET status = $2
WHERE phone = $1
RETURNING id, phone, name, date_of_birth, avatar_url, created_at, status, latitude, longitude
`

type UpdateStatusParams struct {
	Phone  string `json:"phone"`
	Status string `json:"status"`
}

func (q *Queries) UpdateStatus(ctx context.Context, arg UpdateStatusParams) (Driver, error) {
	row := q.db.QueryRowContext(ctx, updateStatus, arg.Phone, arg.Status)
	var i Driver
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Name,
		&i.DateOfBirth,
		&i.AvatarUrl,
		&i.CreatedAt,
		&i.Status,
		&i.Latitude,
		&i.Longitude,
	)
	return i, err
}

const verify = `-- name: Verify :one
UPDATE driver
SET verified = true
WHERE phone = $1
RETURNING id, phone, name, date_of_birth, avatar_url, created_at, status, latitude, longitude
`

func (q *Queries) Verify(ctx context.Context, phone string) (Driver, error) {
	row := q.db.QueryRowContext(ctx, verify, phone)
	var i Driver
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Name,
		&i.DateOfBirth,
		&i.AvatarUrl,
		&i.CreatedAt,
		&i.Status,
		&i.Latitude,
		&i.Longitude,
	)
	return i, err
}
