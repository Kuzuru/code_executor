package models

import (
	"context"
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
	CollectionName = "users"
)

var (
	ErrUserNotFound      = errors.New("пользователь не найден")
	ErrUserAlreadyExists = errors.New("пользователь с таким именем уже существует")
)

func CreateUser(user *User) error {
	collection := database.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	id := uuid.New()
	user.ID = id.String()

	_, err := collection.InsertOne(ctx, user)

	return err
}

func FindUserByName(name string) (User, error) {
	collection := database.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var user User

	err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, ErrUserNotFound
		}

		return user, err
	}

	return user, nil
}

func FindUserByID(id int64) (User, error) {
	collection := database.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var user User

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, ErrUserNotFound
		}

		return user, err
	}

	return user, nil
}

func GetUsers(limit int64) ([]User, error) {
	collection := database.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	findOptions := options.Find()
	findOptions.SetLimit(limit)

	var users []User
	cur, err := collection.Find(ctx, bson.D{}, findOptions)

	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Printf("[ERR] [models/user/database.go] %+v", err)
		}
	}(cur, ctx)

	if err != nil {
		return users, err
	}

	for cur.Next(ctx) {
		var user User
		err := cur.Decode(&user)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func GetAllUsers(pageSize int64) ([]User, error) {
	collection := database.Client.Database(os.Getenv("MONGO_DB_NAME")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	findOptions := options.Find()
	var allUsers []User

	var skip int64
	for {
		findOptions.SetLimit(pageSize)
		findOptions.SetSkip(skip)

		cur, err := collection.Find(ctx, bson.D{}, findOptions)

		defer func(cur *mongo.Cursor, ctx context.Context) {
			err := cur.Close(ctx)
			if err != nil {
				log.Printf("[ERR] [models/user/database.go] %+v", err)
			}
		}(cur, ctx)

		if err != nil {
			return allUsers, err
		}

		usersFetched := false
		for cur.Next(ctx) {
			usersFetched = true
			var user User
			err := cur.Decode(&user)
			if err != nil {
				return allUsers, err
			}

			allUsers = append(allUsers, user)
		}

		if err := cur.Err(); err != nil {
			return allUsers, err
		}

		// If no users were fetched in the last iteration, break the loop
		if !usersFetched {
			break
		}

		skip += pageSize
	}

	return allUsers, nil
}
