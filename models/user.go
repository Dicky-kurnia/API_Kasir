package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primariKey"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"Updated_at"`
}
