package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ChristinaKozi/go-gin-gonic/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

func GetSongs(c *gin.Context) {
	cursor, err := songCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var songs []models.Song
	if err := cursor.All(context.Background(), &songs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, songs)
}

func GetSongByID(c *gin.Context) {
	id := c.Param("id")
	var song models.Song
	err := songCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&song)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
		return
	}

	c.JSON(http.StatusOK, song)
}
