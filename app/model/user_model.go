package model

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	Email     string     `json:"email" db:"email"`
	SureName  string     `json:"sure_name" db:"sure_name"`
	FullName  string     `json:"full_name" db:"full_name"`
	Password  string     `json:"password" db:"password"`
	GoogleId  *string    `json:"google_id" db:"google_id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type CreateUserModel struct {
	Email    string `json:"email" db:"email"`
	SureName string `json:"sure_name" db:"sure_name"`
	FullName string `json:"full_name" db:"full_name"`
	Password string `json:"password" db:"password"`
}
