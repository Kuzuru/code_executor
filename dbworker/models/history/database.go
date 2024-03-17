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
	CollectionName = "history"
)

var ErrHistoryPointNotFound = errors.New("history point not found")

func CreateHistoryPoint(history History) error {
	collection := database.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	_, err := collection.InsertOne(ctx, history)

	return err
}

func FindHistoryBySourceID(sourceID int64) (History, error) {
	collection := database.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var historyPoint History

	err := collection.FindOne(ctx, bson.M{"source_id": sourceID}).Decode(&historyPoint)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return historyPoint, ErrHistoryPointNotFound
		}

		return historyPoint, err
	}

	return historyPoint, nil
}

func GetUserHistoryPoints(sourceIDs []int64, limit int64, offset int64) ([]History, error) {
	collection := database.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)

	filter := bson.M{
		"_id": bson.M{
			"$in": sourceIDs,
		},
	}

	var history []History
	cur, err := collection.Find(ctx, filter, findOptions)

	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Printf("[ERR] [models/sources/database.go] %+v", err)
		}
	}(cur, ctx)

	if err != nil {
		return history, err
	}

	return history, nil
}
