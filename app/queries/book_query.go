package queries

import (
	"context"

	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// BookQueries struct for queries from Book model.
type BookQueries struct {
	Coll *mongo.Collection
}

// GetBooks method for getting all books.
func (q *BookQueries) GetBooks() ([]models.Book, error) {
	// Define books variable.
	books := []models.Book{}

	// Send query to database.
	cursor, err := q.Coll.Find(context.Background(), bson.M{})
	if err != nil {
		return books, err
	}
	defer cursor.Close(context.Background())

	// Decode all books.
	if err := cursor.All(context.Background(), &books); err != nil {
		return books, err
	}

	// Return query result.
	return books, nil
}

// GetBooksByAuthor method for getting all books by given author.
func (q *BookQueries) GetBooksByAuthor(author string) ([]models.Book, error) {
	// Define books variable.
	books := []models.Book{}

	// Send query to database.
	cursor, err := q.Coll.Find(context.Background(), bson.M{"author": author})
	if err != nil {
		return books, err
	}
	defer cursor.Close(context.Background())

	// Decode all books.
	if err := cursor.All(context.Background(), &books); err != nil {
		return books, err
	}

	// Return query result.
	return books, nil
}

// GetBook method for getting one book by given ID.
func (q *BookQueries) GetBook(id uuid.UUID) (models.Book, error) {
	// Define book variable.
	book := models.Book{}

	// Define filter.
	filter := bson.M{"_id": id}

	// Send query to database.
	err := q.Coll.FindOne(context.Background(), filter).Decode(&book)
	if err != nil {
		// Return empty object and error.
		return book, err
	}

	// Return query result.
	return book, nil
}

// CreateBook method for creating book by given Book object.
func (q *BookQueries) CreateBook(b *models.Book) error {
	// Send query to database.
	_, err := q.Coll.InsertOne(context.Background(), b)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// UpdateBook method for updating book by given Book object.
func (q *BookQueries) UpdateBook(id uuid.UUID, b *models.Book) error {
	// Define filter and update document.
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"updated_at":  b.UpdatedAt,
			"title":       b.Title,
			"author":      b.Author,
			"book_status": b.BookStatus,
			"book_attrs":  b.BookAttrs,
		},
	}

	// Send query to database.
	_, err := q.Coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// DeleteBook method for delete book by given ID.
func (q *BookQueries) DeleteBook(id uuid.UUID) error {
	// Define filter.
	filter := bson.M{"_id": id}

	// Send query to database.
	_, err := q.Coll.DeleteOne(context.Background(), filter)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
