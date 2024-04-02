package models

import (
	"context"
	models "dbworker/models/users"
	"errors"
	"log"
	"os"
	"time"

	"dbworker/database"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CollectionName = "sources"
)

var ErrSourceNotFound = errors.New("source not found")

func CreateSource(source SourceDTO) error {
	collection := database.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	preparedSource := new(Source)
	currentTimeDate := time.Now()

	preparedSource.CreatedAt = &currentTimeDate
	preparedSource.UpdatedAt = &currentTimeDate
	preparedSource.LastRunAt = &currentTimeDate

	id := uuid.New()

	preparedSource.ID = id.String()
	preparedSource.UserId = source.UserId

	_, err := models.FindUserByID(preparedSource.UserId)
	if err != nil {
		return errors.New("пользователя не существует")
	}

	_, err = collection.InsertOne(ctx, preparedSource)

	return err
}

func FindSourceByID(id string) (Source, error) {
	collection := database.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var source Source

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&source)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return source, ErrSourceNotFound
		}

		return source, err
	}

	return source, nil
}

func FindSourceByUserID(userId string) (Source, error) {
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

func GetUserSources(userId string, limit int64, offset int64) ([]Source, error) {
	collection := database.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)

	filter := bson.M{
		"user_id": bson.M{
			"$eq": userId,
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

	for cur.Next(ctx) {
		var source Source
		err := cur.Decode(&source)
		if err != nil {
			return sources, err
		}

		sources = append(sources, source)
	}

	if err := cur.Err(); err != nil {
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
