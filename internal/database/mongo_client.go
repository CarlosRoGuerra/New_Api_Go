package database

import (
	"context"
	"fmt"
	"time"

	_ "github.com/CarlosRoGuerra/New_Api_Go/v1/configs"
	"github.com/CarlosRoGuerra/New_Api_Go/v1/pkg/types"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	client *mongo.Client
}

func NewDefaultMongoClient() (*MongoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	mongoHost := viper.GetString("mongo_host")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoHost))
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

func (m *MongoClient) DeleteUser(collection string, user types.User) error {
	var err error
	ctx := context.Background()
	filter := bson.M{"id": user.Id}
	coll, err := m.getCollection(collection)
	if err != nil {
		return err
	}
	_, err = coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoClient) UpdateUser(collection string, user types.User) (types.User, error) {
	var err error
	ctx := context.Background()
	filter := bson.M{"id": user.Id}
	coll, err := m.getCollection(collection)
	if err != nil {
		return user, err
	}
	update := bson.M{
		"$set": bson.M{
			"name":     user.Name,
			"password": user.Password,
		},
	}
	_, err = coll.UpdateOne(ctx, filter, update)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (m *MongoClient) getCollection(collection string) (*mongo.Collection, error) {
	return m.client.Database(viper.GetString("mongo_database")).Collection(collection), nil
}
