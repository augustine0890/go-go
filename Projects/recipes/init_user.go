package main

import (
	"context"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Load the .env file in the current working directory
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	users := map[string]string{
		"admin":     "NTz8NgXN47knNZwk",
		"supertree": "TjyGqzQc5aGRMaR4",
		"test":      "PxYR7AzuwB5V2BKd",
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("users")

	for username, password := range users {
		passHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}

		collection.InsertOne(ctx, bson.M{
			"username": username,
			"password": string(passHashed),
		})
	}
}
