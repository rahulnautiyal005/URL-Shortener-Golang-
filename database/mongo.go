package database

import (
	"context"
	"errors"
	"time"
	"url-shortener/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoStore(uri string) (*MongoStore, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	collection := client.Database("urlshortener").Collection("urls")

	return &MongoStore{
		client:     client,
		collection: collection,
	}, nil
}

func (s *MongoStore) GetNextID() (int64, error) {
	// MongoDB doesn't have auto-increment IDs like SQL. 
	// We can use a counter collection or just use the current timestamp in nanoseconds for a unique ID.
	return time.Now().UnixNano(), nil
}

func (s *MongoStore) Save(url *models.URL) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.collection.InsertOne(ctx, url)
	return err
}

func (s *MongoStore) GetByCode(code string) (*models.URL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var url models.URL
	err := s.collection.FindOne(ctx, bson.M{"short_code": code}).Decode(&url)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("URL not found")
		}
		return nil, err
	}
	return &url, nil
}

func (s *MongoStore) IncrementClick(code string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.collection.UpdateOne(
		ctx,
		bson.M{"short_code": code},
		bson.M{"$inc": bson.M{"click_count": 1}},
	)
	return err
}
