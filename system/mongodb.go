package system

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection URI
var uri string
var client *mongo.Client
var System_Logs_DB *mongo.Collection
var Yahoo_Stocks_DB *mongo.Collection

var MONGO_INITIALISED = false

func InitMongo() {
	var err error

	if GetConfig("mongo-password-req") == "1" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s",
			GetConfig("mongo-user"),
			GetConfig("mongo-password"),
			GetConfig("mongo-ip"),
			GetConfig("mongo-port"))
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s",
			GetConfig("mongo-ip"),
			GetConfig("mongo-port"))
	}

	Log("Attempting to connect MongoDB to: "+uri, true)

	// Create a new client and connect to the server
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	Fatal("MongoDB couldn't make initial connection to "+uri, err)

	// Ping the primary
	err = client.Ping(context.TODO(), readpref.Primary())
	Fatal("Can't reach mongo at: "+uri, err)

	Log("Successfully connected MongoDB to: "+uri, false)

	System_Logs_DB = client.Database("system").Collection("logs")
	Log("Successfully connected to yahoo stocks mongodb", false)

	Yahoo_Stocks_DB = client.Database("yahoo").Collection("stocks")
	Log("Successfully connected to yahoo stocks mongodb", false)

	MONGO_INITIALISED = true
}
