package models

import (
	"context"
	"dbworker/database"
	"errors"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CollectionName = "sources"
)

var ErrSourceNotFound = errors.New("sources not found")

func CreateSource(source Source) error {
	collection := database.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	_, err := collection.InsertOne(ctx, source)

	return err
}

func FindSourceByUserID(userId int64) (Source, error) {
	collection := database.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var source Source

	err := collection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&source)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return source, ErrSourceNotFound
		}

		return source, err
	}

	return source, nil
}

func GetUserSources(user_id int64, limit int64, offset int64) ([]Source, error) {
	collection := database.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)

	filter := bson.M{
		"user_id": bson.M{
			"$eq": user_id,
		},
	}

	var sources []Source
	cur, err := collection.Find(ctx, filter, findOptions)

	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Printf("[ERR] [models/sources/database.go] %+v", err)
		}
	}(cur, ctx)

	if err != nil {
		return sources, err
	}

	return sources, nil
}

func (source *Source) UpdateSourceUpdatetAt(updatedAt time.Time) error {
	collection := database.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	filter := bson.M{
		"user_id": bson.M{
			"$eq": source.UserId,
		},
	}

	update := bson.M{
		"$set": bson.M{
			"updated_at": updatedAt,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		log.Println("[ERR] [models/sources/database.go] No matching document found or updated_at already exists")
	}

	return nil
}
