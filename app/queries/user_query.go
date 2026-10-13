package queries

import (
	"context"

	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// UserQueries struct for queries from User model.
type UserQueries struct {
	Coll *mongo.Collection
}

// GetUserByID query for getting one User by given ID.
func (q *UserQueries) GetUserByID(id uuid.UUID) (models.User, error) {
	// Define User variable.
	user := models.User{}

	// Define filter.
	filter := bson.M{"_id": id}

	// Send query to database.
	err := q.Coll.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		// Return empty object and error.
		return user, err
	}

	// Return query result.
	return user, nil
}

// GetUserByEmail query for getting one User by given Email.
func (q *UserQueries) GetUserByEmail(email string) (models.User, error) {
	// Define User variable.
	user := models.User{}

	// Define filter.
	filter := bson.M{"email": email}

	// Send query to database.
	err := q.Coll.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		// Return empty object and error.
		return user, err
	}

	// Return query result.
	return user, nil
}

// CreateUser query for creating a new user by given email and password hash.
func (q *UserQueries) CreateUser(u *models.User) error {
	// Send query to database.
	_, err := q.Coll.InsertOne(context.Background(), u)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
