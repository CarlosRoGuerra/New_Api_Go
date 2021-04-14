package database

import (
	"context"
	"fmt"
	"time"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const database = "carlos"
const mongohost = "mongodb://localhost:27017"

type MongoClient struct {
	client *mongo.Client
}

func NewDefaultMongoClient() (*MongoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	fmt.Println("connected to db")
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

	return m.client.Database(database).Collection(collection), nil
}
