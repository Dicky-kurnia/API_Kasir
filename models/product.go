package models

import (
	"time"
)

type Product struct {
	ID           uint      `json:"id" gorm:"primariKey"`
	Name         string    `json:"name"`
	SerialNumber string    `json:"serial_number"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated-at"`
}
