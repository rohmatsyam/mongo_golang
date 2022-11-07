package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectDB() (client *mongo.Client, err error) {
	client, err = mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	//	ping db
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB")
	return client, nil
}

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(EnvMongoDB()).Collection(collectionName)
	return collection
}
