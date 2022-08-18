// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
	"time"
)

type Driver struct {
	ID          int64           `json:"id"`
	Phone       string          `json:"phone"`
	Name        string          `json:"name"`
	DateOfBirth sql.NullTime    `json:"date_of_birth"`
	AvatarUrl   sql.NullString  `json:"avatar_url"`
	CreatedAt   time.Time       `json:"created_at"`
	Status      string          `json:"status"`
	Latitude    sql.NullFloat64 `json:"latitude"`
	Longitude   sql.NullFloat64 `json:"longitude"`
}
