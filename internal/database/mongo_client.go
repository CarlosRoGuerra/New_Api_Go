package database

import (
	"context"
	"fmt"
	"time"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Collection string
	Database   string
}

func (m *MongoClient) CreateUser(collection string) (types.User, error) {
	user := types.User{Id: "01", Name: "Carlos", Password: "456"}
	ctx := context.Background()
	coll := GetCollection("users")
	_, nil := coll.InsertOne(ctx, user)
	return types.User{}, nil
}

func (m *MongoClient) GetUsers(collection string) ([]types.User, error) {
	var users types.User
	var err error
	filter := bson.D{}
	ctx := context.Background()
	cur, err := GetCollection("users").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		err = cur.Decode(&users)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

	}
	return nil, nil
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

	return client.Database("carlos").Collection(collection)
}
