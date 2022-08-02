// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
	"time"
)

type Address struct {
	ID          int64   `json:"id"`
	DriverID    int64   `json:"driver_id"`
	Detail      string  `json:"detail"`
	HouseNumber string  `json:"house_number"`
	Street      string  `json:"street"`
	Ward        string  `json:"ward"`
	District    string  `json:"district"`
	City        string  `json:"city"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type Driver struct {
	ID             int64          `json:"id"`
	Phone          string         `json:"phone"`
	HashedPassword string         `json:"hashed_password"`
	Name           string         `json:"name"`
	DateOfBirth    sql.NullTime   `json:"date_of_birth"`
	AvatarUrl      sql.NullString `json:"avatar_url"`
	Verified       bool           `json:"verified"`
	CreatedAt      time.Time      `json:"created_at"`
}
