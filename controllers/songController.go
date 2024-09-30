package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ChristinaKozi/go-gin-gonic/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func CreateSong(c *gin.Context) {
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	song.ID = primitive.NewObjectID().Hex()
	result, err := songCollection.InsertOne(context.Background(), song)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
