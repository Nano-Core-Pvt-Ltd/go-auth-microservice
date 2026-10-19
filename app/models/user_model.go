package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe User object.
type User struct {
	ID           uuid.UUID `bson:"_id" json:"id" validate:"required,uuid"`
	CreatedAt    time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at" json:"updated_at"`
	Name         string    `bson:"name" json:"name"`
	Email        string    `bson:"email" json:"email" validate:"required,email,lte=255"`
	PasswordHash string    `bson:"password_hash" json:"password_hash,omitempty" validate:"required,lte=255"`
	UserStatus   int       `bson:"user_status" json:"user_status" validate:"required,len=1"`
	UserRole     string    `bson:"user_role" json:"user_role" validate:"required,lte=25"`
}
