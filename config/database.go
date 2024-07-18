package config

import (
	"context"
    "log"
    "os"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDatabase() {
    mongoURI := os.Getenv("MONGO_URI")
    databaseName := os.Getenv("DATABASE_NAME")

    if mongoURI == "" || databaseName == "" {
        log.Fatal("FATAL: MONGO_URI and DATABASE_NAME must be set")
    }

    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    DB = client.Database(databaseName)
    log.Println("Connected to MongoDB!")
}
