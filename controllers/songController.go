package controllers

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var songCollection *mongo.Collection

func ConnectToDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("MongoDB connection error:", err)
		return
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("MongoDB ping error:", err)
		return
	}

	fmt.Println("Connected to MongoDB")
	songCollection = client.Database("musicDB").Collection("songs")
}
