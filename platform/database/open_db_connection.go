package database

import (
	"os"

	"github.com/create-go-app/fiber-go-template/app/queries"
)

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries // load queries from User model
	*queries.BookQueries // load queries from Book model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new MongoDB connection.
	client, err := MongoConnection()
	if err != nil {
		return nil, err
	}

	// Define database name from .env file.
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "fiber_go_template"
	}

	// Select database.
	db := client.Database(dbName)

	return &Queries{
		// Set queries from models:
		UserQueries: &queries.UserQueries{Coll: db.Collection("users")}, // from User model
		BookQueries: &queries.BookQueries{Coll: db.Collection("books")}, // from Book model
	}, nil
}
