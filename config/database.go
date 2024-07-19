package config

import (
    "context"
    "log"
    "os"
	
    "go.mongodb.org/mongo-driver/bson"
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
        log.Fatal("FATAL: Cannot connect to MongoDB with error: ", err)
    }

    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal("FATAL: Cannot ping MongoDB with error: ", err)
    }

    DB = client.Database(databaseName)
    log.Println("Connected to MongoDB!")
	
    setupIndices()
}

// Create indices on first and last name in order to speed up search queries
func setupIndices() {
    collection := DB.Collection("contacts")

    indexModel := mongo.IndexModel{
        Keys: bson.D{
			{Key: "first_name", Value: 1},
			{Key: "last_name", Value: 1},
		},
        Options: options.Index().SetUnique(false), // We can have duplicate named contacts. Let's say, several Eli Cohen-s.
    }

    _, err := collection.Indexes().CreateOne(context.Background(), indexModel)
    if err != nil {
        log.Fatal("FATAL: Failed to create name indices ", err)
    }
	
	// Indexes for individual fields
    indexModelPhone := mongo.IndexModel{
        Keys: bson.D{{Key: "phone", Value: 1}},
        Options: options.Index().SetUnique(false),
    }
    _, err = collection.Indexes().CreateOne(context.Background(), indexModelPhone)
    if err != nil {
        log.Fatal("FATAL: Failed to create phone index ", err)
    }

    indexModelAddress := mongo.IndexModel{
        Keys: bson.D{{Key: "address", Value: 1}},
        Options: options.Index().SetUnique(false),
    }
    _, err = collection.Indexes().CreateOne(context.Background(), indexModelAddress)
    if err != nil {
        log.Fatal("FATAL: Failed to create address index ", err)
    }

    log.Println("Indices created successfully")
}
