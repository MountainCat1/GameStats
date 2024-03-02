package repositories

import (
	"context"
	domain "gamestats-domain"
	"gamestats-domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var _ domain.MatchRepo = MongoDbMatchRepo{}

type MongoDbMatchRepo struct {
	collection *mongo.Collection
}

func MewMatchRepository(db *mongo.Database) domain.MatchRepo {
	return &MongoDbMatchRepo{
		collection: db.Collection("matches"),
	}
}

func (repo MongoDbMatchRepo) GetMatchDetails(matchID string) (*entities.MatchDetails, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var entity entities.MatchDetails
	filter := bson.M{"id": matchID}
	err := repo.collection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo MongoDbMatchRepo) AddMatchDetails(matchDetails *entities.MatchDetails) (error, *entities.MatchDetails) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.InsertOne(ctx, matchDetails)
	if err != nil {
		return err, nil
	}

	return nil, matchDetails
}
