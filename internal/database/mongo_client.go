package database

import (
	"context"
	"fmt"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const database = "carlos"
const mongohost = "mongodb://localhost:27017"

type MongoClient struct {
	client *mongo.Client
}

func NewDefaultMongoClient() (*MongoClient, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongohost))
	if err != nil {
		return nil, err
	}

	return &MongoClient{client: client}, nil
}

func (m *MongoClient) CreateUser(collection string, user types.User) (types.User, error) {
	ctx := context.Background()
	coll, err := m.getCollection(collection)
	if err != nil {
		return types.User{}, err
	}
	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		return types.User{}, err
	}
	return user, nil
}

func (m *MongoClient) GetUsers(collection string) ([]types.User, error) {
	var err error
	ctx := context.Background()
	coll, err := m.getCollection(collection)
	if err != nil {
		return nil, err
	}
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	// find a way to return users xD
	currentUser := types.User{}
	users := []types.User{}
	for cursor.Next(ctx) {
		err = cursor.Decode(&currentUser)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		users = append(users, currentUser)
	}
	return users, nil
}

func (m *MongoClient) DeleteUser(tableName string, user types.User) error {
	return nil
}

func (m *MongoClient) getCollection(collection string) (*mongo.Collection, error) {
	ctx := context.Background()
	err := m.client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return m.client.Database(database).Collection(collection), nil
}
