package infrastructure

import (
	"gamestats-domain/entities"
	"gamestats-domain/repositories"
)

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

func (r *userRepository) GetUserLoginDetails(username string) (*entities.LoginDetails, error) {
	user, err := r.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return &user.LoginDetails, err
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *mongo.Database) repositories.IUserRepo {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

// AddUser inserts a new User document into the MongoDB
func (r *userRepository) AddUser(user *entities.User) (error, *entities.User) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err, nil
	}

	return nil, user
}

// GetUserByUsername finds a User document by username
func (r *userRepository) GetUserByUsername(username string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user entities.User
	filter := bson.M{"username": username}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
