package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetMongoURI is a function that returns the MongoDB URI
func GetMongoURI() string {
	godotenv.Load(".env")
	return fmt.Sprintf("mongodb://%s:%s@%s:%s",
		os.Getenv("MONGO_USER"),
		os.Getenv("MONGO_PASSWORD"),
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
	)
}

// GetMongoClient is a function that returns the MongoDB client
func GetMongoClient() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(GetMongoURI())
	clientOptions.SetMaxPoolSize(200)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// check connections
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	collections := client.Database("Attendance")

	return collections, nil
}
