package infrastructure

import (
	domain "gamestats-domain"
	"gamestats-domain/entities"
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

func (repo *userRepository) GetUserLoginDetails(username string) (*entities.LoginDetails, error) {
	user, err := repo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return &user.LoginDetails, err
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *mongo.Database) domain.UserRepo {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

// AddUser inserts a new User document into the MongoDB
func (repo *userRepository) AddUser(user *entities.User) (error, *entities.User) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.InsertOne(ctx, user)
	if err != nil {
		return err, nil
	}

	return nil, user
}

// GetUserByUsername finds a User document by username
func (repo *userRepository) GetUserByUsername(username string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user entities.User
	filter := bson.M{"username": username}
	err := repo.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
