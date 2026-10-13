package models

import (
	"time"

	"github.com/google/uuid"
)

// Book struct to describe book object.
type Book struct {
	ID         uuid.UUID `bson:"_id" json:"id" validate:"required,uuid"`
	CreatedAt  time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at" json:"updated_at"`
	UserID     uuid.UUID `bson:"user_id" json:"user_id" validate:"required,uuid"`
	Title      string    `bson:"title" json:"title" validate:"required,lte=255"`
	Author     string    `bson:"author" json:"author" validate:"required,lte=255"`
	BookStatus int       `bson:"book_status" json:"book_status" validate:"required,len=1"`
	BookAttrs  BookAttrs `bson:"book_attrs" json:"book_attrs" validate:"required,dive"`
}

// BookAttrs struct to describe book attributes.
type BookAttrs struct {
	Picture     string `bson:"picture" json:"picture"`
	Description string `bson:"description" json:"description"`
	Rating      int    `bson:"rating" json:"rating" validate:"min=1,max=10"`
}
