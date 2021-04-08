package database

import (
	"context"
	"time"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/pkg/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	OnGetUsers   func(tableName string) ([]types.User, error)
	OnCreateUser func(tableName string) (types.User, error)
}

func (m *MongoClient) GetUsers(tableName string) ([]types.User, error) {
	if m.OnGetUsers != nil {
		return m.OnGetUsers(tableName)
	}
	return nil, nil
}

func (m *MongoClient) CreateUser(tableName string) (types.User, error) {
	if m.OnCreateUser != nil {
		return m.OnCreateUser(tableName)
	}
	return types.User{}, nil
}

func GetCollection(collection string) *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database("users").Collection(collection)
}
