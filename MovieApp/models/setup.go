package models

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString string

func init() {
	user := os.Getenv("USER_NAME")
	pass := os.Getenv("USER_PWD")
	if user == "" || pass == "" {
		log.Fatal("MongoDB credentials not set in environment variables")
	}
	connectionString = fmt.Sprintf("mongodb+srv://%s:%s@cluster0.i5dznvm.mongodb.net/?appName=Cluster0", user, pass)
}

const db = "movies"
const collName = "movies"

var mongoClient *mongo.Client

func ConnectDatabase() {

	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		panic(err)
	}

	log.Println("Connected to Database")
	mongoClient = client

}
