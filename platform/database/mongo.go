package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

// MongoConnection func for connection to MongoDB database.
func MongoConnection() (*mongo.Client, error) {
	// Define MongoDB connection URI.
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	// Define a context with timeout for connection.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define database connection for MongoDB.
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// Try to ping database.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		defer client.Disconnect(ctx) // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return client, nil
}
